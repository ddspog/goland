// Package sort contains sort functions for working with arrays.
package sort

// Parallel takes an array and sort using selected number of threads.
func Parallel(array []int, numberOfThreads int) []int {
	var step = len(array) / numberOfThreads
	var done = make(chan []int, numberOfThreads)
	var result = make([]int, 0)

	var numberOfUsedThreads = 0
	for i, j := 0, step; j < len(array); i, j = j, j+step {
		go sort(array[i:j], done)
		numberOfUsedThreads++
	}
	for i := 0; i < numberOfUsedThreads; i++ {
		arr := <-done
		result = MergeArrays(result, arr)
	}

	return result
}

func sort(array []int, done chan []int) {
	BubbleSort(array)
	done <- array
}
