package algorithms

import (
	"fmt"
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
	case "quick":
		testQuickSort()
	case "all":
		testLinearSearch()
		testBinarySearch()
		testBubbleSort()
		testQuickSort()
	}
}
