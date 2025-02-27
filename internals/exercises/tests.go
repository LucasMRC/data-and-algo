package exercises

import (
	"fmt"

	"github.com/LucasMRC/kata-machine-go/internals/test_utils"
)

var test_results = test_utils.TestResults

func testCrystallBalls() {
	fmt.Println("CrystallBalls: running tests.")
	fmt.Println("-----------------------------")
	test_cases := crystall_ball_test_cases

	for i, test := range test_cases {
		result := crystall_balls(test.Value)
		test_status := test_utils.ToEqual(result, test.Expected)
		fmt.Printf("test case %d: %v\n", i, test_status)
		if test_status == test_results.FAILED {
			fmt.Printf("\texpected %v, result %v\n", test.Expected, result)
		}
	}
}
