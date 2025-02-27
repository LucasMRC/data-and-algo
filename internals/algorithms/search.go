package algorithms

import (
	"math"
	"slices"
)

func linear_search[T comparable](haystack []T, needle T) bool {
	return slices.Contains(haystack, needle)
}

func binary_search(haystack []int, needle int) (bool, int) {
	for {
		checking := int(math.Floor(float64(len(haystack) / 2)))
		if len(haystack) == 0 {
			return false, 0
		} else if haystack[checking] == needle {
			return true, checking
		} else if haystack[checking] > needle {
			haystack = haystack[:checking]
		} else {
			haystack = haystack[checking+1:]
		}
	}
}
