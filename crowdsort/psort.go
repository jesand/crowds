package crowdsort

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

// A unique identifier for a batch of comparisons
type BatchId string

// PSortable describes the methods necessary to sort by performing comparisons
// in parallel. In addition to the standard sort.Interface methods, we have a
// methods ScheduleLess() and Compare() which allow the parallel sorter to
// schedule a batch of parallel comparisons and dispatch them to the comparer.
//
// A single round of comparisons has four steps:
// 1. ScheduleLess() is called for all i, j in the current batch.
// 2. Compare() is called.
// 3. WaitForBatch() is called to block until the batch is done.
// 4. Less() is called for i, j corresponding to objects for which
//    ScheduleLess() was called (though they may be moved into new positions
//    with Swap() before Less() is called).
//
// If Less() receives a call for some i, j which can not be answered because of
// a failure to previously pass the indices to ScheduleLess(), it should
// to panic.
type PSorter interface {
	sort.Interface

	// Register a comparison to be performed in parallel. This will always be
	// called before Less() for a given pair of values.
	// However, Less() should be prepared to compare the objects even if they
	// have subsequently been moved by Swap() to different indices. It should
	// also be prepared to ask the reverse question: both Less(i, j) and
	// Less(j, i).
	ScheduleLess(i, j int) error

	// Compare tells the sort algorithm that the local comparison batch is
	// finished with all calls to ScheduleLess(), and will next call
	// WaitForBatch() in preparation for calling Less(). Compare() returns
	// an identifier which is used in the call to WaitForBatch().
	Compare() (BatchId, error)

	// WaitForBatch() blocks until the comparison batch has completed. It
	// returns an error if the batch fails for some reason.
	// If an id of an old, already-complete batch is used then WaitForBatch must
	// return immediately.
	WaitForBatch(id BatchId) error

	// ChoosePivots() selects s optimal pivots, given a list of indexes of
	// elements which are already in their correct positions. Note that this is
	// only used if a maximum number of rounds is set. Otherwise, we always
	// select the median of each interval bounded by elements with known correct
	// positions.
	ChoosePivots(s int, correct CorrectList) []int
}

// Adapts a standard sort.Interface so it can be used for PSort() or PSelect().
type SortAdapter struct {
	sort.Interface
}

func (sorter SortAdapter) ScheduleLess(i, j int) error                   { return nil }
func (sorter SortAdapter) Compare() (BatchId, error)                     { return BatchId("dummyid"), nil }
func (sorter SortAdapter) WaitForBatch(id BatchId) error                 { return nil }
func (sorter SortAdapter) ChoosePivots(s int, correct CorrectList) []int { return nil }

// Select the kth-largest value (the smallest has k=0) in a list slice, and
// partitions the slice around it so that all smaller items are before it and
// all larger items are after it.
// You may pass -1 for left and right to select over the entire list.
// If left and right are at least 0, they are used to slice the list, and the
// kth largest slice element is found instead.
func PSelect(sorter PSorter, left, right, k int) error {
	var (
		alg *PSortAlg
		err error
	)
	func() {
		defer func() {
			if rerr := recover(); rerr != nil {
				err = rerr.(error)
			}
		}()
		alg = NewPSelectAlg(sorter, left, right, k)
	}()
	if err != nil {
		return err
	}
	return alg.Run()
}

// Constructs a PSortAlg instance to perform the specified selection.
func NewPSelectAlg(sorter PSorter, left, right, k int) *PSortAlg {
	if left < 0 {
		left = 0
	}
	if right < 0 || right > sorter.Len() {
		right = sorter.Len()
	}
	if left >= right {
		left = right - 1
	}
	if k < 0 || k+left >= right {
		panic(fmt.Errorf("Invalid PSelect call: k=%d should be between 0 and %d", k, right-left-1))
	}
	return &PSortAlg{
		Correct:     CorrectList{{0, left - 1}, {right, sorter.Len() - 1}},
		FixedPivots: []int{left + k},
		MaxRounds:   1,
		Requested:   make(map[int]map[int]bool),
		Sorter:      sorter,
	}
}

// Partially sort a list using large batches of comparisons in parallel.
// If maxRounds iterations are performed, the sort operation will stop.
// Pass -1 for maxRounds to completely sort the list.
func PSort(sorter PSorter, maxRounds int) error {
	return NewPSortAlg(sorter, maxRounds).Run()
}

// Constructs a PSortAlg instance to perform the specified partial sort.
func NewPSortAlg(sorter PSorter, maxRounds int) *PSortAlg {
	return &PSortAlg{
		MaxRounds: maxRounds,
		Requested: make(map[int]map[int]bool),
		Sorter:    sorter,
	}
}

// PSortAlg implements parallel sorting and selection.
// This type is meant to be safe to Marshal/Unmarshal in your format of choice,
// so that program execution can be terminated and resumed in between comparison
// batches with very high latency.
type PSortAlg struct {

	// The ID of the most recently-scheduled batch of comparisons
	BatchId BatchId

	// The number of batches and comparisons we've requested
	NumBatches, NumComparisons int

	// Statistics about batch size
	BatchSize, MinBatch, MaxBatch int

	// The interface used for comparisons
	Sorter PSorter

	// The round currently being run
	Round int

	// The number of rounds of sorting remaining, or -1 to run to completion
	MaxRounds int

	// Inclusive index intervals of already-selected pivots
	Correct CorrectList

	// Optional indices to always use for Pivots
	FixedPivots []int

	// Indices of target pivots to select
	Pivots []int

	// Parallel stack frames for selectSort()
	CallStacks []*selectSortCallStack

	// Comparisons we have already requested
	Requested map[int]map[int]bool
}

// Choose a set of pivots to select in parallel this round, and place them in
// alg.Pivots.
func (alg *PSortAlg) ChoosePivots() {

	// Find the next pivots
	if len(alg.FixedPivots) > 0 {
		alg.Pivots = alg.FixedPivots
	} else if alg.MaxRounds > 0 {
		s := int(math.Floor(math.Log2(float64(alg.Round)) - 1))
		alg.Pivots = alg.Sorter.ChoosePivots(s, alg.Correct)
	} else {
		alg.Pivots = nil
	}
	if len(alg.Pivots) == 0 {
		left := 0
		for _, rng := range alg.Correct {
			if rng[0] == left {
				left = rng[1] + 1
			} else {
				right := rng[0] - 1
				median := left + (right-left)/2
				alg.Pivots = append(alg.Pivots, median)
				left = rng[1] + 1
			}
		}
		if left < alg.Sorter.Len() {
			right := alg.Sorter.Len() - 1
			median := left + (right-left)/2
			alg.Pivots = append(alg.Pivots, median)
		}
	}
}

// Run the configured sort/select algorithm to completion.
func (alg *PSortAlg) Run() (err error) {
	alg.Round = 1
	for err == nil && (alg.MaxRounds < 0 || alg.Round <= alg.MaxRounds) {
		err = alg.RunNextRound()
	}
	return
}

// Run the next step of the configured sort/select algorithm. Each step will
// start a batch of n comparisons, where the following approximately holds:
// alg.Sorter.Len() <= n <= 2 * alg.Sorter.Len().
func (alg *PSortAlg) RunNextRound() error {

	// Prepare to clean up the stacks
	defer func() {
		var nextStacks []*selectSortCallStack
		for _, stack := range alg.CallStacks {
			if len(*stack) > 0 {
				nextStacks = append(nextStacks, stack)
			}
		}
		alg.CallStacks = nextStacks
		if len(alg.CallStacks) == 0 {
			alg.Round++
		}
	}()

	// Wait for the most recent batch
	if alg.BatchId != "" {
		if err := alg.Sorter.WaitForBatch(alg.BatchId); err != nil {
			return err
		}
		alg.BatchId = ""
		for _, stack := range alg.CallStacks {
			frame := stack.NextFrame()
			frame.IsWaitingForBatch = false
		}
	}

	// Prepare for the next selection
	if len(alg.CallStacks) == 0 {
		alg.Correct.Compact(alg.Sorter.Len())
		if len(alg.Correct) == 1 && alg.Correct[0] == [2]int{0, alg.Sorter.Len() - 1} {
			alg.MaxRounds = alg.Round
			return nil
		}

		alg.ChoosePivots()
		for _, pivot := range alg.Pivots {
			if !alg.Correct.IsCorrect(pivot) {
				left, right := 0, alg.Sorter.Len()-1
				for _, rng := range alg.Correct {
					if rng[1] < pivot && rng[1] >= left {
						left = rng[1] + 1
					} else if rng[0] > pivot && rng[0] <= right {
						right = rng[0] - 1
					}
				}
				stack := &selectSortCallStack{}
				stack.Call(left, right, pivot-left)
				alg.CallStacks = append(alg.CallStacks, stack)
			}
		}
		if len(alg.CallStacks) == 0 {
			alg.MaxRounds = alg.Round
			return nil
		}
	}

	// Run until the next select batch
	var (
		err           error
		anyIsWaiting  = false
		anyNotWaiting = false
	)
	for i, stack := range alg.CallStacks {
		frame := stack.NextFrame()
		if !frame.IsWaitingForBatch {
			switch frame.NextPhase {
			case selectSortFindU:
				err = alg.selectSortFindU(i)
			case selectSortFindV:
				err = alg.selectSortFindV(i)
			case selectSortPartition1:
				err = alg.selectSortPartition1(i)
			case selectSortPartition2:
				err = alg.selectSortPartition2(i)
			case selectSortReduce:
				err = alg.selectSortReduce(i)
			default:
				err = fmt.Errorf("Invalid call stack phase: %v", frame.NextPhase)
			}
			if err != nil {
				return err
			}
		}
		if frame.IsWaitingForBatch {
			anyIsWaiting = true
		} else {
			anyNotWaiting = true
		}
	}

	// Kick off the comparison batch
	if anyIsWaiting && !anyNotWaiting {
		alg.NumBatches++
		if alg.MinBatch == 0 || alg.BatchSize < alg.MinBatch {
			alg.MinBatch = alg.BatchSize
		}
		if alg.BatchSize > alg.MaxBatch {
			alg.MaxBatch = alg.BatchSize
		}
		alg.BatchSize = 0
		alg.BatchId, err = alg.Sorter.Compare()
	}
	return err
}

func (alg *PSortAlg) selectSortFindU(stack int) error {
	frame := alg.CallStacks[stack].NextFrame()
	frame.N = frame.Right - frame.Left + 1
	const (
		alpha = 1.0 / 2
		beta  = 1.0 / 2
	)
	var (
		fn = float64(frame.N)
		f  = math.Pow(fn, 2/3) * math.Pow(math.Log(fn), 1/3)
	)
	frame.SampleSize = int(math.Min(math.Ceil(alpha*f), fn-1))
	var gapSize = math.Pow(beta*float64(frame.SampleSize)*math.Log(fn), 1/2)
	if fn <= 1 {
		alg.CallStacks[stack].SendReturn(frame.Left, frame.Left)
		return nil
	}

	// Sampling
	for i := 0; i < int(frame.SampleSize); i++ {
		alg.Sorter.Swap(frame.Left+i, frame.Left+i+rand.Intn(frame.N-i))
	}

	// Pivot selection
	var approxK = float64(frame.K*frame.SampleSize) / fn
	frame.IU = int(math.Max(math.Ceil(approxK-gapSize), 0))
	frame.IV = int(math.Min(math.Ceil(approxK+gapSize), float64(frame.SampleSize-1)))
	if frame.SampleSize > 1 {
		alg.CallStacks[stack].Recurse(frame.Left, frame.Left+frame.SampleSize-1, frame.IU)
		frame.NextPhase = selectSortFindV
		return nil
	} else {
		frame.NextPhase = selectSortPartition1
		return nil
	}
}

func (alg *PSortAlg) selectSortFindV(stack int) error {
	_, uPlus := alg.CallStacks[stack].ReceiveReturn()
	frame := alg.CallStacks[stack].NextFrame()
	if frame.IV > uPlus {
		alg.CallStacks[stack].Recurse(frame.RetMaxIdx, frame.Left+frame.SampleSize-1,
			frame.IV-frame.RetMaxIdx)
	}
	frame.NextPhase = selectSortPartition1
	return nil
}

func (alg *PSortAlg) selectSortPartition1(stack int) error {

	// Ignore the result from a possible recursive call to place v correctly
	if alg.CallStacks[stack].NextFrame().RetMaxIdx > 0 {
		alg.CallStacks[stack].ReceiveReturn()
	}

	frame := alg.CallStacks[stack].NextFrame()
	frame.NextPhase = selectSortPartition2
	if frame.IU == frame.IV || frame.K >= frame.N/2 {
		// Partition around u, then v
		requested, err := alg.schedulePartition(frame.Left, frame.Right, frame.IU)
		frame.IsWaitingForBatch = requested
		return err
	} else {
		// Partition around v, then u
		requested, err := alg.schedulePartition(frame.Left, frame.Right, frame.IV)
		frame.IsWaitingForBatch = requested
		return err
	}
}

func (alg *PSortAlg) selectSortPartition2(stack int) error {
	frame := alg.CallStacks[stack].NextFrame()
	frame.NextPhase = selectSortReduce
	if frame.IU == frame.IV || frame.K >= frame.N/2 {
		// Partition around u, then v
		frame.ULeft, frame.URight = alg.arrayPartition(frame.Left, frame.Right, frame.IU)
		if frame.IV > frame.IU {
			requested, err := alg.schedulePartition(frame.URight+1, frame.Right, frame.IV)
			frame.IsWaitingForBatch = requested
			return err
		} else {
			return nil
		}
	} else {
		// Partition around v, then u
		frame.VLeft, frame.VRight = alg.arrayPartition(frame.Left, frame.Right, frame.IV)
		requested, err := alg.schedulePartition(frame.Left, frame.VRight-1, frame.IU)
		frame.IsWaitingForBatch = requested
		return err
	}
}

func (alg *PSortAlg) selectSortReduce(stack int) error {
	frame := alg.CallStacks[stack].NextFrame()
	if frame.IU == frame.IV {
		frame.VLeft, frame.VRight = frame.ULeft, frame.URight
	} else if frame.K >= frame.N/2 {
		// Partition around u, then v
		frame.VLeft, frame.VRight = alg.arrayPartition(frame.URight+1, frame.Right, frame.IV)
	} else {
		// Partition around v, then u
		frame.ULeft, frame.URight = alg.arrayPartition(frame.Left, frame.VRight-1, frame.IU)
	}

	var (
		numL = frame.ULeft - frame.Left
		numU = frame.URight - frame.ULeft + 1
		// numM = int(math.Max(float64(frame.VLeft-frame.URight-1), 0))
		// numV = frame.VRight - frame.VLeft + 1
		numR = frame.Right - frame.VRight
	)
	alg.Correct = append(alg.Correct, [2]int{frame.ULeft, frame.URight})
	if frame.ULeft != frame.VLeft {
		alg.Correct = append(alg.Correct, [2]int{frame.VLeft, frame.VRight})
	}

	// Stopping test
	k := frame.Left + frame.K
	if frame.ULeft <= k && k <= frame.URight {
		alg.CallStacks[stack].SendReturn(frame.ULeft, frame.URight)
		return nil
	}
	if frame.VLeft <= k && k <= frame.VRight {
		alg.CallStacks[stack].SendReturn(frame.VLeft, frame.VRight)
		return nil
	}

	// Reduction
	if frame.K < frame.ULeft-frame.Left {
		// Left side
		alg.CallStacks[stack].TailRecurse(frame.Left, frame.ULeft-1, frame.K)
	} else if frame.K > frame.VRight-frame.Left {
		// Right side
		alg.CallStacks[stack].TailRecurse(frame.VRight+1, frame.Right, frame.K-frame.N+numR)
	} else {
		// Middle
		alg.CallStacks[stack].TailRecurse(frame.URight+1, frame.VLeft-1, frame.K-numL-numU)
	}
	return nil
}

// Schedules the comparisons needed to partition a sublist between left and
// right (inclusive) about the element at index left + k.
func (alg *PSortAlg) schedulePartition(left, right, k int) (bool, error) {
	var (
		j         = left + k
		requested bool
	)
	for i := left; i <= right; i++ {
		if i != j && !alg.Requested[j][i] && !alg.Requested[i][j] {
			alg.BatchSize++
			alg.NumComparisons++
			if alg.Requested[j] == nil {
				alg.Requested[j] = make(map[int]bool)
			}
			alg.Requested[j][i] = true
			requested = true
			if err := alg.Sorter.ScheduleLess(i, j); err != nil {
				return requested, err
			}
		}
	}
	return requested, nil
}

// Partitions an array, using previously-scheduled comparisons, about the
// element at index left + k.
// Returns indices lEq, rEq such that all elements with indices between
// lEq and rEq, inclusive, are equal to the item at index left + k.
func (alg *PSortAlg) arrayPartition(left, right, k int) (lEq, rEq int) {
	if right < left || k < 0 || k > right-left {
		return
	}

	// A1. Initialize. Should have x_left <= x_iv <= x_right.
	lEq, rEq = left, right
	var (
		i = lEq
		p = left + 1
		q = right - 1
		j = rEq
	)

	// Keep track of where v is: sort.Interface doesn't let us copy the value
	iv := left + k
	safeSwap := func(a, b int) {
		alg.Sorter.Swap(a, b)
		if a == iv {
			iv = b
		} else if b == iv {
			iv = a
		}
	}

	if left != iv {
		safeSwap(left, iv)
	}
	if alg.Sorter.Less(iv, right) {
		rEq = q
	} else if alg.Sorter.Less(right, iv) {
		safeSwap(left, right)
		lEq = p
	}

	// A2. Increase i until x_i >= x_iv
	partitioning := true
	for partitioning {
		i++
		for i != iv && alg.Sorter.Less(i, iv) {
			i++
		}

		// A3. Decrease j until x_j <= x_iv
		j--
		for j != iv && alg.Sorter.Less(iv, j) {
			j--
		}

		// A4. Exchange: x_j <= x_iv <= x_i
		if i < j {
			safeSwap(i, j)
			if !alg.Sorter.Less(i, iv) && !alg.Sorter.Less(iv, i) {
				safeSwap(i, p)
				p++
			}
			if !alg.Sorter.Less(j, iv) && !alg.Sorter.Less(iv, j) {
				safeSwap(j, q)
				q--
			}
		} else {
			partitioning = false
			if i == j {
				i++
				j--
			}
		}
	}

	// A5. Cleanup.
	alg.arraySwap(lEq, p-1, j)
	alg.arraySwap(i, q, right)
	lEq = lEq + j - p + 1
	rEq = rEq - q + i - 1
	return
}

// Swaps items between list[a:b] and list[b+1:c] such that the first
// d elements and last d elements are swapped in arbitrary order, where d
// is the smaller size of the two sets of ranges.
func (alg *PSortAlg) arraySwap(a, b, c int) {
	for d := 0; d < int(math.Min(float64(b-a+1), float64(c-b))); d++ {
		alg.Sorter.Swap(a+d, c-d)
	}
}

type selectSortPhase int

const (
	selectSortFindU selectSortPhase = iota
	selectSortFindV
	selectSortPartition1
	selectSortPartition2
	selectSortReduce
)

// A single call stack frame for the selectSort algorithm.
type selectSortFrame struct {

	// The (inclusive) boundaries of the interval to select within. It is
	// assumed that a previous partition operation has been run over these order
	// statistics.
	Left, Right int

	// The order statistic to select is Left + K.
	K int

	// The number of elements between Left and Right
	N int

	// The size of the subsample we're using in this iteration
	SampleSize int

	// The 'u' and 'v' order statistics to attempt to bracket k with
	IU, IV int

	// The minimum and maximum indices equal to IU and IV
	ULeft, URight, VLeft, VRight int

	// The phase of the algorithm to run next
	NextPhase selectSortPhase

	// Whether the frame needs to wait for a comparison batch to complete
	IsWaitingForBatch bool

	// Whether the frame is blocking on a return call
	WantReturn bool

	// Whether the frame has returned
	HasReturned bool

	// The minimum and maximum indices of any element equal to the
	// Kth order statistic, as returned from the prior stack frame.
	RetMinIdx, RetMaxIdx int
}

// A call stack for selectSort
type selectSortCallStack []*selectSortFrame

// Get the next stack frame to execute
func (stack *selectSortCallStack) NextFrame() *selectSortFrame {
	for i := len(*stack) - 1; i >= 0; i-- {
		if !(*stack)[i].HasReturned {
			return (*stack)[i]
		}
	}
	return nil
}

// Make a recursive call on the stack
func (stack *selectSortCallStack) Call(left, right, k int) {
	*stack = append(*stack, &selectSortFrame{
		NextPhase: selectSortFindU,
		Left:      left,
		Right:     right,
		K:         k,
	})
}

// Make a recursive call to receive a return value
func (stack *selectSortCallStack) Recurse(left, right, k int) {
	stack.NextFrame().WantReturn = true
	stack.Call(left, right, k)
}

// Make a tail-recursive call on the stack
func (stack *selectSortCallStack) TailRecurse(left, right, k int) {
	frame := stack.NextFrame()
	*frame = selectSortFrame{
		NextPhase: selectSortFindU,
		Left:      left,
		Right:     right,
		K:         k,
	}
}

// Store the return value for the next recursive call
func (stack *selectSortCallStack) SendReturn(minIdx, maxIdx int) {
	var iFrame = len(*stack) - 1
	for iFrame >= 0 && (*stack)[iFrame].HasReturned {
		iFrame--
	}
	if iFrame < 1 {
		*stack = selectSortCallStack{}
	} else {
		frame := (*stack)[iFrame]
		frame.HasReturned = true
		if (*stack)[iFrame-1].WantReturn {
			frame.RetMinIdx, frame.RetMaxIdx = minIdx, maxIdx
		} else {
			// Nobody receives the last return value, so remove the unwanted frames
			*stack = (*stack)[:iFrame]
		}
	}
}

// Get the return value from the last recursive call
func (stack *selectSortCallStack) ReceiveReturn() (minIdx, maxIdx int) {
	frame := (*stack)[len(*stack)-1]
	minIdx, maxIdx = frame.RetMinIdx, frame.RetMaxIdx
	*stack = (*stack)[0 : len(*stack)-1]
	return
}

type CorrectList [][2]int

func (list CorrectList) Len() int           { return len(list) }
func (list CorrectList) Swap(i, j int)      { list[i], list[j] = list[j], list[i] }
func (list CorrectList) Less(i, j int) bool { return list[i][0] < list[j][0] }

// Compact the list by sorting and merging adjacent or overlapping intervals
func (list *CorrectList) Compact(n int) {
	if len(*list) > 0 {
		sort.Sort(*list)
		var newCorrect [][2]int
		var left, right = -1, -1
		for i, c := range *list {
			if i == 0 {
				left, right = c[0], c[1]
				if left <= 1 {
					left = 0
				}
			} else if right < c[0]-2 {
				newCorrect = append(newCorrect, [2]int{left, right})
				left, right = c[0], c[1]
			} else {
				right = c[1]
			}
		}
		if right >= n-2 {
			right = n - 1
		}
		newCorrect = append(newCorrect, [2]int{left, right})
		*list = newCorrect
	}
}

// Ask whether a given order statistic has been selected
func (list CorrectList) IsCorrect(k int) bool {
	for _, rng := range list {
		if rng[0] <= k && k <= rng[1] {
			return true
		}
	}
	return false
}

// Expand the list into all distinct correct indices. For testing purposes.
func (list CorrectList) Expand() []int {
	var expanded []int
	for _, rng := range list {
		for i := rng[0]; i <= rng[1]; i++ {
			expanded = append(expanded, i)
		}
	}
	return expanded
}
