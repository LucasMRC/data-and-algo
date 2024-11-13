package algorithms

import "github.com/LucasMRC/kata-machine-go/internals/test_utils"

var search_test_cases = []test_utils.BasicTestCase[int, bool]{
	{
		Value:    3,
		Expected: true,
	},
	{
		Value:    69,
		Expected: true,
	},
	{
		Value:    7,
		Expected: false,
	},
	{
		Value:    99,
		Expected: true,
	},
	{
		Value:    100,
		Expected: false,
	},
	{
		Value:    13371,
		Expected: false,
	},
	{
		Value:    69420,
		Expected: true,
	},
	{
		Value:    -1,
		Expected: false,
	},
}

var sort_test_cases = []test_utils.BasicTestCase[[]int, []int]{
	{
		Value:    []int{4, 1, 5, 12, 55, 23, 7, 18, 3},
		Expected: []int{1, 3, 4, 5, 7, 12, 18, 23, 55},
	},
	{
		Value:    []int{1, 4, 1, 61, 5, 2, 76, 7, 61, 6, 5, 12},
		Expected: []int{1, 1, 2, 4, 5, 5, 6, 7, 12, 61, 61, 76},
	},
}
