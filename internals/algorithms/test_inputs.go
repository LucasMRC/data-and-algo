package algorithms

import "github.com/LucasMRC/kata-machine-go/internals/test_utils"

var search_test_cases = []test_utils.AlgorithmTestCase[int, bool]{
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

var crystall_ball_test_cases = []test_utils.AlgorithmTestCase[[]bool, int]{
	{
		Value:    []bool{false, false, false, false, false, false, false, true, true},
		Expected: 7,
	},
	{
		Value:    []bool{false, false, false, false, false, false, false, false, true},
		Expected: 8,
	},
	{
		Value:    []bool{false, false, false, false, false, false, false, false, false},
		Expected: -1,
	},
	{
		Value:    []bool{true, true, true, true, true, true, true, true, true},
		Expected: 0,
	},
	{
		Value:    []bool{false, false, true, true, true, true, true, true, true},
		Expected: 2,
	},
}

var sort_test_cases = []test_utils.AlgorithmTestCase[[]int, []int]{
	{
		Value:    []int{4, 1, 5, 12, 55, 23, 7, 18, 3},
		Expected: []int{1, 3, 4, 5, 7, 12, 18, 23, 55},
	},
	{
		Value:    []int{1, 4, 1, 61, 5, 2, 76, 7, 61, 6, 5, 12},
		Expected: []int{1, 1, 2, 4, 5, 5, 6, 7, 12, 61, 61, 76},
	},
}
