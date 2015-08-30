package crowdsort

import (
	. "github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

const (
	tol    = 0.06
	trials = 10000
)

func TestPerfectCmp(t *testing.T) {
	Convey("Given an instance of PerfectCmp", t, func() {
		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		cmp := NewComparator(PerfectComparator,
			SortAdapter{sort.IntSlice(list)},
			IntSliceL1Dist(list))

		Convey("Comparisons are noise-free", func() {
			correct, incorrect := 0, 0
			for i := 0; i < trials; i++ {
				if cmp.Less(3, 7) {
					correct++
				} else {
					incorrect++
				}
			}
			So(correct, ShouldEqual, trials)
			So(incorrect, ShouldEqual, 0)
		})
	})
}

func TestConstCmp(t *testing.T) {
	Convey("Given an instance of ConstCmp", t, func() {
		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		cmp := NewComparator(ConstComparator,
			SortAdapter{sort.IntSlice(list)},
			IntSliceL1Dist(list))

		Convey("Comparison noise is consistent with the model", func() {
			correct, incorrect := 0, 0
			for i := 0; i < trials; i++ {
				if cmp.Less(3, 7) {
					correct++
				} else {
					incorrect++
				}
			}
			prob := cmp.(ConstCmp).PCorrect
			So(correct, ShouldBeBetween, trials*(prob-tol), trials*(prob+tol))
			So(incorrect, ShouldBeBetween, trials*(1-prob-tol), trials*(1-prob+tol))
		})
	})
}

func TestRoundRobinCmp(t *testing.T) {
	Convey("Given an instance of RoundRobinCmp", t, func() {
		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		cmp := NewComparator(RoundRobinComparator,
			SortAdapter{sort.IntSlice(list)},
			IntSliceL1Dist(list))

		Convey("Comparison noise is consistent with the model", func() {
			correct, incorrect := 0, 0
			for i := 0; i < trials; i++ {
				if cmp.Less(3, 7) {
					correct++
				} else {
					incorrect++
				}
			}
			rrcmp := cmp.(*RoundRobinCmp)
			prob := (rrcmp.Min + rrcmp.Max) / 2
			So(correct, ShouldBeBetween, trials*(prob-tol), trials*(prob+tol))
			So(incorrect, ShouldBeBetween, trials*(1-prob-tol), trials*(1-prob+tol))
		})
	})
}

func TestDistanceCmp(t *testing.T) {
	Convey("Given an instance of DistanceCmp", t, func() {
		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		cmp := NewComparator(DistanceComparator,
			SortAdapter{sort.IntSlice(list)},
			IntSliceL1Dist(list))

		Convey("Comparison noise is consistent with the model", func() {
			var (
				correct, incorrect int
				prob               float64
			)

			correct, incorrect = 0, 0
			for i := 0; i < trials; i++ {
				if cmp.Less(1, 2) {
					correct++
				} else {
					incorrect++
				}
			}
			prob = 0.5 + 0.5*1/9
			So(correct, ShouldBeBetween, trials*(prob-tol), trials*(prob+tol))
			So(incorrect, ShouldBeBetween, trials*(1-prob-tol), trials*(1-prob+tol))

			correct, incorrect = 0, 0
			for i := 0; i < trials; i++ {
				if cmp.Less(3, 7) {
					correct++
				} else {
					incorrect++
				}
			}
			prob = 0.5 + 0.5*4/9
			So(correct, ShouldBeBetween, trials*(prob-tol), trials*(prob+tol))
			So(incorrect, ShouldBeBetween, trials*(1-prob-tol), trials*(1-prob+tol))

			correct, incorrect = 0, 0
			for i := 0; i < trials; i++ {
				if cmp.Less(1, 8) {
					correct++
				} else {
					incorrect++
				}
			}
			prob = 0.5 + 0.5*7/9
			So(correct, ShouldBeBetween, trials*(prob-tol), trials*(prob+tol))
			So(incorrect, ShouldBeBetween, trials*(1-prob-tol), trials*(1-prob+tol))
		})
	})
}

func TestRoundRobinDistanceCmp(t *testing.T) {
	Convey("Given an instance of RoundRobinDistanceCmp", t, func() {
		list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		cmp := NewComparator(RoundRobinDistanceComparator,
			SortAdapter{sort.IntSlice(list)},
			IntSliceL1Dist(list))

		Convey("Comparison noise is consistent with the model", func() {
			var (
				correct, incorrect int
				prob               float64
				rrcmp              = cmp.(*RoundRobinDistanceCmp).RoundRobinCmp
				rrprob             = (rrcmp.Min + rrcmp.Max) / 2
			)

			correct, incorrect = 0, 0
			for i := 0; i < trials; i++ {
				if cmp.Less(1, 2) {
					correct++
				} else {
					incorrect++
				}
			}
			prob = rrprob * (0.5 + 0.5*1/9)
			So(correct, ShouldBeBetween, trials*(prob-tol), trials*(prob+tol))
			So(incorrect, ShouldBeBetween, trials*(1-prob-tol), trials*(1-prob+tol))

			correct, incorrect = 0, 0
			for i := 0; i < trials; i++ {
				if cmp.Less(3, 7) {
					correct++
				} else {
					incorrect++
				}
			}
			prob = rrprob * (0.5 + 0.5*4/9)
			So(correct, ShouldBeBetween, trials*(prob-tol), trials*(prob+tol))
			So(incorrect, ShouldBeBetween, trials*(1-prob-tol), trials*(1-prob+tol))

			correct, incorrect = 0, 0
			for i := 0; i < trials; i++ {
				if cmp.Less(1, 8) {
					correct++
				} else {
					incorrect++
				}
			}
			prob = rrprob * (0.5 + 0.5*7/9)
			So(correct, ShouldBeBetween, trials*(prob-tol), trials*(prob+tol))
			So(incorrect, ShouldBeBetween, trials*(1-prob-tol), trials*(1-prob+tol))
		})
	})
}
