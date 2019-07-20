package interfaces

// BoundaryUpdater interface define the actions needed to become a Bynary
// Search Boundary for this implementation, Boundaries are the bootom and upper
// indexes on each bynary search step.
type BoundaryUpdater interface {
	Update(int)
}

// Boundary is a Custom type from int that will implement BoundaryUpdater.
type Boundary int

// Updates, updates the Boundary to the given value.
func (b *Boundary) Update(value int) {
	*b = Boundary(value)
}

// chop performs a binary search on the given array looking for the given target.
//
// This implementation uses interfaces, interfaces in Go is my favourite
// feature, more than concurrency. Probably this example is not the best case
// for interfaces, but I love developing Go with interfaces and wanted to give
// it a try.
func chop(target int, array []int) int {
	lower := Boundary(0)
	upper := Boundary(len(array))

	for lower < upper {
		h := int(lower + (upper-lower)/2)

		if array[h] == target {
			return h
		}

		if array[h] > target {
			upper.Update(h)
		} else {
			lower.Update(h + 1)
		}
	}

	return -1
}
