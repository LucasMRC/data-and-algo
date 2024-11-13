package algorithms

import (
	"fmt"

	"github.com/LucasMRC/kata-machine-go/internals/test_utils"
)

var test_results = test_utils.TestResults

func testLinearSearch() {
	fmt.Println("LinearSearch: running tests.")
	fmt.Println("----------------------------")
	input := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	test_cases := search_test_cases

	for i, test := range test_cases {
		result := linear_search(input, test.Value)
		test_status := test_utils.ToEqual(result, test.Expected)
		fmt.Printf("test case %d: %v\n", i, test_status)
		if test_status == test_results.FAILED {
			fmt.Printf("\texpected %v, result %v\n", test.Expected, result)
		}
	}
	fmt.Println("============================")
}

func testBinarySearch() {
	fmt.Println("BinarySearch: running tests.")
	fmt.Println("----------------------------")
	input := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	test_cases := search_test_cases

	for i, test := range test_cases {
		result := binary_search(input, test.Value)
		test_status := test_utils.ToEqual(result, test.Expected)
		fmt.Printf("test case %d: %v\n", i, test_status)
		if test_status == test_results.FAILED {
			fmt.Printf("\texpected %v, result %v\n", test.Expected, result)
		}
	}
	fmt.Println("============================")
}

func testBubbleSort() {
	fmt.Println("BubbleSort: running tests.")
	fmt.Println("-----------------------------")
	test_cases := sort_test_cases

	for i, test := range test_cases {
		result := bubble_sort(test.Value)
		test_status := test_utils.SlicesToEqual(result, test.Expected)
		fmt.Printf("test case %d: %v\n", i, test_status)
		if test_status == test_results.FAILED {
			fmt.Printf("\texpected %v, result %v\n", test.Expected, result)
		}
	}
	fmt.Println("============================")
}
