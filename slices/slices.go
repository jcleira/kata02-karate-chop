package slices

// chop performs a binary search on the given array looking for the given target.
//
// On this implementation I wanted to give a try manipulating the slice instead
// of boundaries updating, actually it's by little the most performant!,
// probably doesn't look as the simplest due the offset, but has less
// allocation involved.
func chop(target int, array []int) int {
	offset := 0

	for len(array) > 0 {
		h := len(array) / 2

		if array[h] == target {
			return h + offset
		}

		if array[h] > target {
			array = array[:h]
		} else {
			array = array[h+1:]
			offset += h + 1
		}
	}

	return -1
}
