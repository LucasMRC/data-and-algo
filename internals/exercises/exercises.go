package exercises

import "fmt"

func Run(ex string) {
	fmt.Println("============================")

	switch ex {
	case "crystall":
		testCrystallBalls()
	case "maze":
		testMazeSolver()
	case "all":
		testCrystallBalls()
		testMazeSolver()
	}
}
