package exercises

import "github.com/LucasMRC/kata-machine-go/internals/test_utils"

var crystall_ball_test_cases = []test_utils.BasicTestCase[[]bool, int]{
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
