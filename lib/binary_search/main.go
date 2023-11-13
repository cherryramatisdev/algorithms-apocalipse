package binary_search

import "math"

func BinarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack)

	for low < high {
		middle := int(math.Floor(float64(low + (high-low)/2)))
		value := haystack[middle]

		if value == needle {
			return true
		}

		if value > needle {
			high = middle
		}

		if value < needle {
			low = middle + 1
		}
	}

	return false
}
