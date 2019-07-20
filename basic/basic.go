package basic

// chop performs a binary search on the given array looking for the given target.
//
// This is the simplest implementation that comes to my mind. The algorithm
// iterates setting the middle point for two min & max indexes (i, j) then
// updates those indexes to the higher or lower half.
func chop(target int, array []int) int {
	i := 0
	j := len(array)

	for i < j {
		h := i + (j-i)/2

		if array[h] == target {
			return h
		}

		if array[h] > target {
			j = h
		} else {
			i = h + 1
		}
	}

	return -1
}
