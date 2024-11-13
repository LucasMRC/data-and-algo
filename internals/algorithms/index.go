package algorithms

import (
	"fmt"
	"math"
)

func Run(algo string) {
	fmt.Println("============================")

	switch algo {
	case "linear":
		testLinearSearch()
	case "binary":
		testBinarySearch()
	case "bubble":
		testBubbleSort()
	case "all":
		testLinearSearch()
		testBinarySearch()
		testBubbleSort()
	}
}

func linear_search[T comparable](haystack []T, needle T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func binary_search(haystack []int, needle int) bool {
	checking := func(h []int) int {
		return int(math.Floor(float64(len(h) / 2)))
	}
	for len(haystack) > 0 {
		index := checking(haystack)
		if index >= len(haystack) {
			break
		}
		value := haystack[index]
		if value == needle {
			return true
		}
		if value < needle {
			haystack = haystack[index+1:]
		} else {
			haystack = haystack[:index]
		}
	}
	return false
}

func bubble_sort(input []int) []int {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j] > input[j+1] {
				tmp := input[j]
				input[j] = input[j+1]
				input[j+1] = tmp
			}
		}
	}

	return input
}
