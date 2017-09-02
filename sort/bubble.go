// Package sort contains sort functions for working with arrays.
package sort

// BubbleSort takes an array and sort it using BubbleSort algorithm.
func BubbleSort(array []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				SwapElements(array, i, i+1)
				swapped = true
			}
		}
	}
	return array
}
