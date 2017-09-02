// Package sort contains sort functions for working with arrays.
package sort

// SwapElements takes an array and swap elements on i and j positions.
func SwapElements(array []int, i, j int) {
	temp := array[j]
	array[j] = array[i]
	array[i] = temp
}

// MergeArrays takes two arrays and merge it maintaining order.
func MergeArrays(left, right []int) []int {
	results := make([]int, 0)

	for len(left) > 0 || len(right) > 0 {
		switch {
		// When both are non-empty arrays
		case len(left) > 0 && len(right) > 0:
			if left[0] <= right[0] {
				results = append(results, left[0])
				left = left[1:len(left)]
			} else {
				results = append(results, right[0])
				right = right[1:len(right)]
			}
		// When right is a empty array
		case len(left) > 0:
			results = append(results, left[0])
			left = left[1:len(left)]
		// When right is a non-empty array
		default:
			results = append(results, right[0])
			right = right[1:len(right)]
		}
	}

	return results
}

// Equals compare two arrays to verified equality between them.
func Equals(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
