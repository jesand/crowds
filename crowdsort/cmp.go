package crowdsort

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type ComparatorType string

const (
	PerfectComparator            ComparatorType = "perfect"
	ConstComparator              ComparatorType = "const"
	RoundRobinComparator         ComparatorType = "round-robin"
	DistanceComparator           ComparatorType = "dist"
	RoundRobinDistanceComparator ComparatorType = "round-robin-dist"
)

// A function to provide the normalized distance between two list objects
type DistFn func(i, j int) float64

// Returns a DistFn to return the normalized L1 distance between values in a []int
func IntSliceL1Dist(slice []int) DistFn {
	s2 := make([]int, len(slice))
	copy(s2[:], slice[:])
	sort.Ints(s2)
	max := float64(s2[len(s2)-1] - s2[0])
	return func(i, j int) float64 {
		if max == 0 {
			return 0
		}
		return math.Abs(float64(slice[i]-slice[j])) / max
	}
}

// Returns a DistFn to return the normalized L1 distance between values in a []float64
func Float64SliceL1Dist(slice []float64) DistFn {
	s2 := make([]float64, len(slice))
	copy(s2[:], slice[:])
	sort.Float64s(s2)
	max := s2[len(s2)-1] - s2[0]
	return func(i, j int) float64 {
		if max == 0 {
			return 0
		}
		return math.Abs(slice[i]-slice[j]) / max
	}
}

// Generate a new comparator of the specified type with default parameters.
func NewComparator(id ComparatorType, cmp PSorter, dist DistFn) PSorter {
	switch id {
	case PerfectComparator:
		return PerfectCmp{PSorter: cmp}
	case ConstComparator:
		return ConstCmp{
			PSorter:  cmp,
			PCorrect: 0.7,
		}
	case RoundRobinComparator:
		return &RoundRobinCmp{
			PSorter:    cmp,
			NumWorkers: 20,
			Min:        0.55,
			Max:        0.9,
		}
	case DistanceComparator:
		return &DistanceCmp{
			PSorter: cmp,
			DistFn:  dist,
			MinDist: 0.1,
			MaxDist: 0.9,
		}
	case RoundRobinDistanceComparator:
		return &RoundRobinDistanceCmp{
			PSorter: cmp,
			DistanceCmp: DistanceCmp{
				DistFn:  dist,
				MinDist: 0.1,
				MaxDist: 0.9,
			},
			RoundRobinCmp: RoundRobinCmp{
				NumWorkers: 20,
				Min:        0.55,
				Max:        0.9,
			},
		}
	default:
		return nil
	}
}

// All comparisons are accurate.
type PerfectCmp struct {
	PSorter
}

// Make a random comparison with certain probability of correctness
func randLess(i, j int, cmp sort.Interface, pCorrect float64) bool {
	if rand.Float64() <= pCorrect {
		return cmp.Less(i, j)
	} else {
		return !cmp.Less(i, j)
	}
}

// Comparisons are correct with fixed probability PCorrect and incorrect
// otherwise.
type ConstCmp struct {
	PSorter
	PCorrect float64
}

func (cmp ConstCmp) Less(i, j int) bool {
	return randLess(i, j, cmp.PSorter, cmp.PCorrect)
}

// Comparisons are correct with worker-dependent probabilities drawn from
// Uniform(Min, Max). Comparisons alternate round-robin between workers.
type RoundRobinCmp struct {
	PSorter
	Min, Max   float64
	NumWorkers int
	Workers    []float64
	NextWorker int
}

func (cmp *RoundRobinCmp) PCorrect(i, j int) float64 {
	if len(cmp.Workers) == 0 {
		for i := 0; i < cmp.NumWorkers; i++ {
			pCorrect := cmp.Min + (cmp.Max-cmp.Min)*rand.Float64()
			cmp.Workers = append(cmp.Workers, pCorrect)
		}
	}
	pCorrect := cmp.Workers[cmp.NextWorker]
	cmp.NextWorker = (cmp.NextWorker + 1) % cmp.NumWorkers
	return pCorrect
}

func (cmp *RoundRobinCmp) Less(i, j int) bool {
	return randLess(i, j, cmp.PSorter, cmp.PCorrect(i, j))
}

// The probability of correct answers depends on the distance between the
// objects.
//
// The probability of correctness is piecewise linear, so:
// p = 0.5 when dist <= MinDist,
// p = 1.0 when dist >= MaxDist,
// and p is interpolated when MinDist < dist < MaxDist.
type DistanceCmp struct {
	PSorter
	DistFn           DistFn
	MinDist, MaxDist float64
}

func (cmp *DistanceCmp) PCorrect(i, j int) float64 {
	var (
		dist     = cmp.DistFn(i, j)
		pCorrect float64
	)
	if dist <= cmp.MinDist {
		pCorrect = 0.5
	} else if dist >= cmp.MaxDist {
		pCorrect = 1.0
	} else {
		pCorrect = 0.5 + 0.5*(dist-cmp.MinDist)/(cmp.MaxDist-cmp.MinDist)
	}
	fmt.Println("i: ", i, " j: ", j, " Dist: ", dist, " PCorrect: ", pCorrect)
	return pCorrect
}

func (cmp *DistanceCmp) Less(i, j int) bool {
	return randLess(i, j, cmp.PSorter, cmp.PCorrect(i, j))
}

// The probability of correct answers is the product of a round-robin model
// and a distance model.
type RoundRobinDistanceCmp struct {
	PSorter
	RoundRobinCmp
	DistanceCmp
}

func (cmp *RoundRobinDistanceCmp) PCorrect(i, j int) float64 {
	return cmp.RoundRobinCmp.PCorrect(i, j) * cmp.DistanceCmp.PCorrect(i, j)
}

func (cmp *RoundRobinDistanceCmp) Less(i, j int) bool {
	return randLess(i, j, cmp.PSorter, cmp.PCorrect(i, j))
}
