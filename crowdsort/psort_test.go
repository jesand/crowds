package crowdsort

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type SortRecorder struct {
	sort.Interface

	List             *[]int
	Scheduled        [][2]int
	UsedComparisons  [][2]int
	ExtraComparisons [][2]int
}

func (sorter *SortRecorder) ScheduleLess(i, j int) error {
	sorter.Scheduled = append(sorter.Scheduled,
		[2]int{(*sorter.List)[i], (*sorter.List)[j]})
	sorter.UsedComparisons = append(sorter.UsedComparisons, [2]int{-1, -1})
	return nil
}
func (sorter *SortRecorder) Less(i, j int) bool {
	found := false
	fwd := [2]int{(*sorter.List)[i], (*sorter.List)[j]}
	bwd := [2]int{(*sorter.List)[j], (*sorter.List)[i]}
	for i, pair := range sorter.Scheduled {
		if pair == fwd || pair == bwd {
			sorter.UsedComparisons[i] = pair
			found = true
			break
		}
	}
	if !found {
		sorter.ExtraComparisons = append(sorter.ExtraComparisons, fwd)
	}
	return sorter.Interface.Less(i, j)
}
func (sorter *SortRecorder) Compare() (BatchId, error) {
	return BatchId("dummyid"), nil
}
func (sorter *SortRecorder) WaitForBatch(id BatchId) error {
	return nil
}
func (sorter *SortRecorder) ChoosePivots(s int, correct CorrectList) []int {
	return nil
}

func TestPSelect(t *testing.T) {
	Convey("Given a small list", t, func() {
		list := []int{5, 4, 2, 6, 8, 0, 3, 1, 7, 9}
		sorter := SortAdapter{sort.IntSlice(list)}

		Convey("Invalid calls are validated", func() {
			So(PSelect(sorter, 0, 0, 1).Error(), ShouldEqual, "Invalid PSelect call: k=1 should be between 0 and 0")
			So(PSelect(sorter, 0, 0, -1).Error(), ShouldEqual, "Invalid PSelect call: k=-1 should be between 0 and 0")
			So(PSelect(sorter, 5, 7, -1).Error(), ShouldEqual, "Invalid PSelect call: k=-1 should be between 0 and 1")
			So(PSelect(sorter, 5, 7, 2).Error(), ShouldEqual, "Invalid PSelect call: k=2 should be between 0 and 1")
		})

		Convey("Selecting the minimum works", func() {
			err := PSelect(sorter, -1, -1, 0)
			So(err, ShouldBeNil)
			So(list[0], ShouldEqual, 0)

			sort.Ints(list[1:])
			So(list, ShouldResemble, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})

		Convey("Selecting the maximum works", func() {
			err := PSelect(sorter, -1, -1, 9)
			So(err, ShouldBeNil)
			So(list[9], ShouldEqual, 9)

			sort.Ints(list[0:9])
			So(list, ShouldResemble, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})

		Convey("Selecting an inner item works", func() {
			err := PSelect(sorter, -1, -1, 5)
			So(err, ShouldBeNil)
			So(list[5], ShouldEqual, 5)

			sort.Ints(list[0:5])
			sort.Ints(list[6:])
			So(list, ShouldResemble, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		})

		Convey("Selecting the first of a range works", func() {
			list[9] = -9
			err := PSelect(sorter, 5, 9, 0)
			So(err, ShouldBeNil)
			So(list[5], ShouldEqual, 0)

			sort.Ints(list[6:9])
			So(list, ShouldResemble, []int{5, 4, 2, 6, 8, 0, 1, 3, 7, -9})
		})

		Convey("Selecting the last of a range works", func() {
			list[9] = -9
			err := PSelect(sorter, 5, 9, 3)
			So(err, ShouldBeNil)
			So(list[8], ShouldEqual, 7)

			sort.Ints(list[5:8])
			So(list, ShouldResemble, []int{5, 4, 2, 6, 8, 0, 1, 3, 7, -9})
		})

		Convey("Selecting an inner item of a range works", func() {
			list[9] = -9
			err := PSelect(sorter, 5, 9, 1)
			So(err, ShouldBeNil)
			So(list[6], ShouldEqual, 1)

			sort.Ints(list[7:9])
			So(list, ShouldResemble, []int{5, 4, 2, 6, 8, 0, 1, 3, 7, -9})
		})
	})

	Convey("Given a large random list", t, func() {
		const n = 1000
		var (
			orig        = rand.Perm(n)
			list        = make([]int, len(orig))
			sorted      = make([]int, len(orig))
			left, right = n / 3, 2 * n / 3
			sortedSlice = make([]int, right-left)
		)
		copy(list[:], orig[:])
		copy(sorted[:], orig[:])
		copy(sortedSlice[:], orig[left:right])
		sort.Ints(sorted)
		sorter := SortAdapter{sort.IntSlice(list)}

		Convey("Selecting the minimum works", func() {
			err := PSelect(sorter, -1, -1, 0)
			So(err, ShouldBeNil)
			So(list[0], ShouldEqual, 0)

			sort.Ints(list[1:])
			So(list, ShouldResemble, sorted)
		})

		Convey("Selecting the maximum works", func() {
			err := PSelect(sorter, -1, -1, n-1)
			So(err, ShouldBeNil)
			So(list[n-1], ShouldEqual, n-1)

			sort.Ints(list[0 : n-1])
			So(list, ShouldResemble, sorted)
		})

		Convey("Selecting an inner item works", func() {
			k := 2 * n / 3
			err := PSelect(sorter, -1, -1, k)
			So(err, ShouldBeNil)
			So(list[k], ShouldEqual, k)

			sort.Ints(list[0:k])
			sort.Ints(list[k+1:])
			So(list, ShouldResemble, sorted)
		})

		Convey("Selecting the first of a range works", func() {
			list[left+20] = -1
			sortedSlice[20] = -1
			err := PSelect(sorter, left, right, 0)
			So(err, ShouldBeNil)
			So(list[left], ShouldEqual, -1)

			sort.Ints(list[left+1 : right])
			sort.Ints(sortedSlice)
			So(list[:left], ShouldResemble, orig[:left])
			So(list[left:right], ShouldResemble, sortedSlice)
			So(list[right:], ShouldResemble, orig[right:])
		})

		Convey("Selecting the last of a range works", func() {
			list[left+20] = n * 2
			sortedSlice[20] = n * 2
			err := PSelect(sorter, left, right, right-left-1)
			So(err, ShouldBeNil)
			So(list[right-1], ShouldEqual, n*2)

			sort.Ints(list[left : right-1])
			sort.Ints(sortedSlice)
			So(list[:left], ShouldResemble, orig[:left])
			So(list[left:right], ShouldResemble, sortedSlice)
			So(list[right:], ShouldResemble, orig[right:])
		})

		Convey("Selecting an inner item of a range works", func() {
			for i, v := range rand.Perm(right - left) {
				list[left+i] = v
				sortedSlice[i] = v
			}
			k := 20
			err := PSelect(sorter, left, right, k)
			So(err, ShouldBeNil)
			So(list[left+k], ShouldEqual, k)

			sort.Ints(list[left : left+k])
			sort.Ints(list[left+k+1 : right])
			sort.Ints(sortedSlice)
			So(list[:left], ShouldResemble, orig[:left])
			So(list[left:right], ShouldResemble, sortedSlice)
			So(list[right:], ShouldResemble, orig[right:])
		})
	})
}

func TestPSort(t *testing.T) {
	Convey("Given a large random list", t, func() {
		const n = 1000
		var (
			list    = rand.Perm(n)
			dupList = make([]int, len(list))
			sorted  = make([]int, len(list))
		)
		copy(sorted[:], list[:])
		sort.Ints(sorted)
		sorter := SortAdapter{sort.IntSlice(list)}

		Convey("Sorting the list for 1 round partitions about the median", func() {
			err := PSort(sorter, 1)
			So(err, ShouldBeNil)

			k := (n - 1) / 2
			So(list[k], ShouldEqual, k)
			sort.Ints(list[:k])
			sort.Ints(list[k+1:])
			So(list, ShouldResemble, sorted)
		})

		Convey("Sorting the entire list works", func() {
			err := PSort(sorter, -1)
			So(err, ShouldBeNil)
			So(list, ShouldResemble, sorted)
		})

		Convey("Sorting the entire list by alg works", func() {
			alg := NewPSortAlg(sorter, 1000)
			err := alg.Run()
			So(err, ShouldBeNil)
			So(list, ShouldResemble, sorted)
			So(alg.Correct, ShouldResemble, CorrectList{{0, n - 1}})
			So(alg.Round, ShouldBeLessThanOrEqualTo, int(math.Ceil(math.Log2(n))))

			So(alg.NumComparisons, ShouldBeBetween, 5*n, 8*n) // tight for now...
			So(alg.MinBatch, ShouldBeGreaterThan, 0)          // too small...
			So(alg.MaxBatch, ShouldBeLessThan, n)
			So(alg.NumBatches, ShouldBeBetween, 3*alg.Round, 5*alg.Round) // arbitrary...
		})

		Convey("When I sort round by round", func() {
			var (
				alg      = NewPSortAlg(sorter, 2)
				err      error
				correct  CorrectList
				corrects []int
				pivots   []int
			)

			// Round 1: median
			Convey("Round 1 is correct", func() {
				err = alg.RunNextRound()
				So(alg.Pivots, ShouldResemble, []int{499})
				pivots = append(pivots, alg.Pivots...)
				sort.Ints(pivots)
				for err == nil && len(alg.CallStacks) > 0 {
					err = alg.RunNextRound()
				}
				So(err, ShouldBeNil)
				copy(dupList[:], list[:])
				correct = append(CorrectList{}, alg.Correct...)
				correct.Compact(n)
				corrects = correct.Expand()
				left := 0
				for _, pivot := range pivots {
					So(list[pivot], ShouldEqual, pivot)
					So(corrects, ShouldContain, pivot)
					sort.Ints(dupList[left:pivot])
					left = pivot + 1
				}
				if left < len(list)-1 {
					sort.Ints(dupList[left:len(list)])
				}
				So(dupList, ShouldResemble, sorted)

				// Round 2: median of each group defined by correct
				Convey("Round 2 is correct", func() {
					err = alg.RunNextRound()
					expPivots := len(correct) + 1
					if correct.IsCorrect(0) {
						expPivots--
					}
					if correct.IsCorrect(n - 1) {
						expPivots--
					}
					So(len(alg.Pivots), ShouldEqual, expPivots)
					left, iPivot := -1, 0
					for _, rng := range correct {
						if rng[0] > 0 {
							So(alg.Pivots[iPivot], ShouldBeBetween, left, rng[0])
							iPivot++
						}
						left = rng[1]
					}
					if !correct.IsCorrect(n - 1) {
						So(alg.Pivots[len(alg.Pivots)-1], ShouldBeBetween, left, len(list))
					}
					pivots = append(pivots, alg.Pivots...)
					sort.Ints(pivots)
					for err == nil && len(alg.CallStacks) > 0 {
						err = alg.RunNextRound()
					}
					So(err, ShouldBeNil)
					copy(dupList[:], list[:])
					correct = append(CorrectList{}, alg.Correct...)
					correct.Compact(n)
					corrects = correct.Expand()
					left = 0
					for _, pivot := range pivots {
						So(list[pivot], ShouldEqual, pivot)
						So(corrects, ShouldContain, pivot)
						sort.Ints(dupList[left:pivot])
						left = pivot + 1
					}
					if left < len(list)-1 {
						sort.Ints(dupList[left:len(list)])
					}
					So(dupList, ShouldResemble, sorted)

					// Round 3: median of each group defined by correct
					Convey("Round 3 is correct", func() {
						err = alg.RunNextRound()
						expPivots := len(correct) + 1
						if correct.IsCorrect(0) {
							expPivots--
						}
						if correct.IsCorrect(n - 1) {
							expPivots--
						}
						So(len(alg.Pivots), ShouldEqual, expPivots)
						left, iPivot := -1, 0
						for _, rng := range correct {
							if rng[0] > 0 {
								So(alg.Pivots[iPivot], ShouldBeBetween, left, rng[0])
								iPivot++
							}
							left = rng[1]
						}
						if !correct.IsCorrect(n - 1) {
							So(alg.Pivots[len(alg.Pivots)-1], ShouldBeBetween, left, len(list))
						}
						pivots = append(pivots, alg.Pivots...)
						sort.Ints(pivots)
						for err == nil && len(alg.CallStacks) > 0 {
							err = alg.RunNextRound()
						}
						So(err, ShouldBeNil)
						copy(dupList[:], list[:])
						correct = append(CorrectList{}, alg.Correct...)
						correct.Compact(n)
						corrects = correct.Expand()
						left = 0
						for _, pivot := range pivots {
							So(list[pivot], ShouldEqual, pivot)
							So(corrects, ShouldContain, pivot)
							sort.Ints(dupList[left:pivot])
							left = pivot + 1
						}
						if left < len(list)-1 {
							sort.Ints(dupList[left:len(list)])
						}
						So(dupList, ShouldResemble, sorted)
					})
				})
			})
		})
	})
}

func TestCallStack(t *testing.T) {
	Convey("Given an empty call stack", t, func() {
		var stack selectSortCallStack

		Convey("NextFrame() is nil", func() {
			So(stack.NextFrame(), ShouldBeNil)
		})

		Convey("When I make a call", func() {
			stack.Call(10, 100, 20)

			Convey("NextFrame() is a fresh call", func() {
				So(stack.NextFrame(), ShouldResemble, &selectSortFrame{
					Left:      10,
					Right:     100,
					K:         20,
					NextPhase: selectSortFindU,
				})
			})

			Convey("When I make a recursive call", func() {
				stack.Recurse(20, 40, 25)

				Convey("The stack is correct", func() {
					So(stack, ShouldResemble, selectSortCallStack{
						&selectSortFrame{
							Left:       10,
							Right:      100,
							K:          20,
							NextPhase:  selectSortFindU,
							WantReturn: true,
						},
						&selectSortFrame{
							Left:      20,
							Right:     40,
							K:         25,
							NextPhase: selectSortFindU,
						},
					})
				})

				Convey("NextFrame() is a fresh call", func() {
					So(stack.NextFrame(), ShouldResemble, &selectSortFrame{
						Left:      20,
						Right:     40,
						K:         25,
						NextPhase: selectSortFindU,
					})
				})

				Convey("When I send a return value", func() {
					stack.SendReturn(25, 35)

					Convey("NextFrame() is correct", func() {
						So(stack.NextFrame(), ShouldResemble, &selectSortFrame{
							Left:       10,
							Right:      100,
							K:          20,
							NextPhase:  selectSortFindU,
							WantReturn: true,
						})
					})

					Convey("The stack is correct", func() {
						So(stack, ShouldResemble, selectSortCallStack{
							&selectSortFrame{
								Left:       10,
								Right:      100,
								K:          20,
								NextPhase:  selectSortFindU,
								WantReturn: true,
							},
							&selectSortFrame{
								Left:        20,
								Right:       40,
								K:           25,
								NextPhase:   selectSortFindU,
								HasReturned: true,
								RetMinIdx:   25,
								RetMaxIdx:   35,
							},
						})
					})

					Convey("When I receive the return value", func() {
						min, max := stack.ReceiveReturn()
						Convey("The stack is correct", func() {
							So(stack, ShouldResemble, selectSortCallStack{
								&selectSortFrame{
									Left:       10,
									Right:      100,
									K:          20,
									NextPhase:  selectSortFindU,
									WantReturn: true,
								},
							})
						})
						Convey("The return values are correct", func() {
							So(min, ShouldEqual, 25)
							So(max, ShouldEqual, 35)
						})
					})
				})

				Convey("When I make a tail recursive call", func() {
					stack.TailRecurse(15, 75, 41)
					Convey("The stack is correct", func() {
						So(stack, ShouldResemble, selectSortCallStack{
							&selectSortFrame{
								Left:       10,
								Right:      100,
								K:          20,
								NextPhase:  selectSortFindU,
								WantReturn: true,
							},
							&selectSortFrame{
								Left:      15,
								Right:     75,
								K:         41,
								NextPhase: selectSortFindU,
							},
						})
					})
				})
			})

			Convey("When I send a return value", func() {
				stack.SendReturn(25, 35)

				Convey("Then the stack is empty", func() {
					So(stack, ShouldResemble, selectSortCallStack{})
				})
			})
		})
	})
}

func TestPSortAlgChoosePivots(t *testing.T) {
	Convey("Given an algorithm with fixed pivots", t, func() {
		alg := PSortAlg{
			FixedPivots: []int{4, 7},
			Sorter:      SortAdapter{sort.IntSlice([]int{5, 4, 2, 6, 8, 0, 3, 1, 7, 9})},
		}

		Convey("ChoosePivots picks the fixed pivots", func() {
			alg.ChoosePivots()
			So(alg.Pivots, ShouldResemble, []int{4, 7})
		})
	})

	Convey("Given an algorithm with no defined pivots", t, func() {
		alg := PSortAlg{
			Sorter: SortAdapter{sort.IntSlice([]int{5, 4, 2, 6, 8, 0, 3, 1, 7, 9})},
		}

		Convey("ChoosePivots picks the median", func() {
			alg.ChoosePivots()
			So(alg.Pivots, ShouldResemble, []int{4})
		})
	})

	Convey("Given an algorithm with no defined pivots and two correct indices", t, func() {
		alg := PSortAlg{
			Sorter:  SortAdapter{sort.IntSlice([]int{5, 4, 2, 6, 8, 0, 3, 1, 7, 9})},
			Correct: CorrectList{{3, 3}, {7, 7}},
		}

		Convey("ChoosePivots picks the median of each range", func() {
			alg.ChoosePivots()
			So(alg.Pivots, ShouldResemble, []int{1, 5, 8})
		})
	})
}

func TestPSortAlgSchedulePartition(t *testing.T) {
	Convey("Given a list of numbers", t, func() {
		list := []int{5, 4, 2, 6, 8, 0, 3, 1, 7, 9}
		sorter := &SortRecorder{Interface: sort.IntSlice(list), List: &list}
		alg := NewPSortAlg(sorter, -1)

		Convey("Scheduling a partition about the element at index 0 works", func() {
			alg.schedulePartition(0, len(list)-1, 0)
			So(sorter.Scheduled, ShouldResemble, [][2]int{
				{4, 5},
				{2, 5},
				{6, 5},
				{8, 5},
				{0, 5},
				{3, 5},
				{1, 5},
				{7, 5},
				{9, 5},
			})
		})

		Convey("Scheduling a sublist works", func() {
			alg.schedulePartition(3, 8, 2)
			So(sorter.Scheduled, ShouldResemble, [][2]int{
				{6, 0},
				{8, 0},
				{3, 0},
				{1, 0},
				{7, 0},
			})

			Convey("Partitioning uses only scheduled comparisons", func() {
				alg.arrayPartition(3, 8, 2)
				So(sorter.UsedComparisons, ShouldResemble, sorter.Scheduled)
				So(sorter.ExtraComparisons, ShouldBeEmpty)
			})
		})
	})
}

func TestPSortAlgArrayPartition(t *testing.T) {
	Convey("Given a permutation of 0:9", t, func() {
		list := []int{5, 4, 2, 6, 8, 0, 3, 1, 7, 9}
		alg := PSortAlg{Sorter: SortAdapter{sort.IntSlice(list)}}

		Convey("Partitioning about the element at index 0 works", func() {
			alg.arrayPartition(0, len(list)-1, 0)
			So(list, ShouldResemble, []int{0, 4, 2, 1, 3, 5, 9, 6, 7, 8})
		})

		Convey("Partitioning about the element at index 3 works", func() {
			alg.arrayPartition(0, len(list)-1, 3)
			So(list, ShouldResemble, []int{3, 4, 2, 5, 1, 0, 6, 9, 7, 8})
		})

		Convey("Partitioning about 0 works", func() {
			alg.arrayPartition(0, len(list)-1, 5)
			So(list, ShouldResemble, []int{0, 9, 2, 6, 8, 5, 3, 1, 7, 4})
		})

		Convey("Partitioning about 1 works", func() {
			alg.arrayPartition(0, len(list)-1, 7)
			So(list, ShouldResemble, []int{0, 1, 9, 6, 8, 4, 3, 5, 7, 2})
		})

		Convey("Partitioning about 8 works", func() {
			alg.arrayPartition(0, len(list)-1, 4)
			So(list, ShouldResemble, []int{7, 4, 2, 6, 5, 0, 3, 1, 8, 9})
		})

		Convey("Partitioning about 9 works", func() {
			alg.arrayPartition(0, len(list)-1, 9)
			So(list, ShouldResemble, []int{5, 4, 2, 6, 8, 0, 3, 1, 7, 9})
		})

		Convey("Partitioning a sub-range starting from the left works", func() {
			alg.arrayPartition(0, 3, 1)
			So(list, ShouldResemble, []int{2, 4, 6, 5, 8, 0, 3, 1, 7, 9})
		})

		Convey("Partitioning an inner sub-range from the left works", func() {
			alg.arrayPartition(3, 7, 0)
			So(list, ShouldResemble, []int{5, 4, 2, 1, 3, 0, 6, 8, 7, 9})
		})

		Convey("Partitioning an inner sub-range from the middle works", func() {
			alg.arrayPartition(3, 7, 3)
			So(list, ShouldResemble, []int{5, 4, 2, 1, 0, 3, 6, 8, 7, 9})
		})

		Convey("Partitioning an inner sub-range from the right works", func() {
			alg.arrayPartition(3, 7, 4)
			So(list, ShouldResemble, []int{5, 4, 2, 0, 1, 6, 3, 8, 7, 9})
		})

		Convey("Partitioning a sub-range ending on the right works", func() {
			alg.arrayPartition(5, 9, 1)
			So(list, ShouldResemble, []int{5, 4, 2, 6, 8, 1, 0, 3, 9, 7})
		})
	})
}

func TestPSortAlgArraySwap(t *testing.T) {
	Convey("Given a list of numbers", t, func() {
		list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		alg := PSortAlg{Sorter: SortAdapter{sort.IntSlice(list)}}

		Convey("Swapping 1:1 does nothing", func() {
			alg.arraySwap(1, 1, 1)
			So(list, ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		})

		Convey("Swapping 1:2 with 3:3 works", func() {
			alg.arraySwap(1, 2, 3)
			So(list, ShouldResemble, []int{1, 4, 3, 2, 5, 6, 7, 8, 9, 10})
		})

		Convey("Swapping 1:2 with 3:4 works", func() {
			alg.arraySwap(1, 2, 4)
			So(list, ShouldResemble, []int{1, 5, 4, 3, 2, 6, 7, 8, 9, 10})
		})

		Convey("Swapping 1:2 with 3:6 works", func() {
			alg.arraySwap(1, 2, 6)
			So(list, ShouldResemble, []int{1, 7, 6, 4, 5, 3, 2, 8, 9, 10})
		})

		Convey("Swapping 0:1 with 2:6 works", func() {
			alg.arraySwap(0, 1, 6)
			So(list, ShouldResemble, []int{7, 6, 3, 4, 5, 2, 1, 8, 9, 10})
		})

		Convey("Swapping 5:7 with 8:9 works", func() {
			alg.arraySwap(5, 7, 9)
			So(list, ShouldResemble, []int{1, 2, 3, 4, 5, 10, 9, 8, 7, 6})
		})
	})
}

func TestCorrectListCompact(t *testing.T) {
	Convey("Compacting an empty list works", t, func() {
		list := CorrectList{}
		list.Compact(0)
		So(list, ShouldResemble, CorrectList{})
	})

	Convey("Compacting a compacted list does nothing", t, func() {
		list := CorrectList{{2, 2}, {5, 5}, {8, 10}}
		list.Compact(20)
		So(list, ShouldResemble, CorrectList{{2, 2}, {5, 5}, {8, 10}})
	})

	Convey("Compacting adjacent elements works", t, func() {
		list := CorrectList{{2, 2}, {3, 4}, {7, 10}}
		list.Compact(20)
		So(list, ShouldResemble, CorrectList{{2, 4}, {7, 10}})
	})

	Convey("Compacting across skipped elements works", t, func() {
		list := CorrectList{{2, 2}, {4, 5}, {7, 10}, {13, 15}}
		list.Compact(20)
		So(list, ShouldResemble, CorrectList{{2, 10}, {13, 15}})
	})

	Convey("Compacting the first element works", t, func() {
		list := CorrectList{{1, 2}, {4, 5}, {7, 10}, {13, 15}}
		list.Compact(20)
		So(list, ShouldResemble, CorrectList{{0, 10}, {13, 15}})
	})

	Convey("Compacting near the last element works", t, func() {
		list := CorrectList{{2, 2}, {4, 5}, {7, 10}, {13, 15}}
		list.Compact(18)
		So(list, ShouldResemble, CorrectList{{2, 10}, {13, 15}})
	})

	Convey("Compacting the last element works", t, func() {
		list := CorrectList{{2, 2}, {4, 5}, {7, 10}, {13, 15}}
		list.Compact(17)
		So(list, ShouldResemble, CorrectList{{2, 10}, {13, 16}})
	})
}
