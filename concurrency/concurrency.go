package concurrency

// chop performs a binary search on the given array looking for the given target.
//
// This implementation uses concurrency composition with two go routines: the
// first will check for the value and set new boundaries, the second will
// recalculate new halves to check. They are related by three operational
// channels and a result channel.
//
// I don't personally like the solution for this implementation, It just don't
// look simple, and the performance after checking the benchmark is poor
// compared with the basic one.
//
// But I do like the defered cleanup for the channels (and goroutines), as once
// that we get a result on either goroutine (both may return a result), the
// defer will trigger each other channel closing.
func chop(target int, array []int) int {
	iC := make(chan int)
	jC := make(chan int)
	hC := make(chan int)
	result := make(chan int)

	go func() {
		defer close(iC)
		defer close(jC)

		for h := range hC {
			if array[h] == target {
				result <- h
				return
			}

			if array[h] > target {
				jC <- h
			} else {
				iC <- h + 1
			}
		}
	}()

	go func(i, j int) {
		defer close(hC)

		for {
			select {
			case i = <-iC:
			case j = <-jC:
			}

			if iC == nil && jC == nil {
				return
			}

			if i >= j {
				result <- -1
				return
			}

			hC <- i + (j-i)/2
		}
	}(0, len(array))

	jC <- len(array)

	return <-result
}
