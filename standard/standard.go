package standard

import (
	"sort"
)

// chop performs a binary search on the given array looking for the given target.
//
// This implementation uses the Standard Library sort.Search function. I
// created this implementation not to be too smart, but for pragmatism, just to
// show that I love the standard library and It is my first reference when
// programming in Go.
func chop(target int, array []int) int {
	i := sort.Search(len(array), func(i int) bool { return array[i] >= target })
	if i < len(array) && array[i] == target {
		return i
	} else {
		return -1
	}
}
