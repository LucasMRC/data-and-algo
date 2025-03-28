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

type maze_solver_args struct {
	maze  []string
	start Coord
	end   Coord
}

var maze_solver_test_cases = []test_utils.BasicTestCase[maze_solver_args, []Coord]{
	{
		Value: maze_solver_args{
			maze: []string{
				"########## #",
				"#        # #",
				"#        # #",
				"# ######## #",
				"#          #",
				"# ##########",
			},
			start: Coord{
				X: 10,
				Y: 0,
			},
			end: Coord{
				X: 1,
				Y: 5,
			},
		},
		Expected: []Coord{
			{X: 10, Y: 0},
			{X: 10, Y: 1},
			{X: 10, Y: 2},
			{X: 10, Y: 3},
			{X: 10, Y: 4},
			{X: 9, Y: 4},
			{X: 8, Y: 4},
			{X: 7, Y: 4},
			{X: 6, Y: 4},
			{X: 5, Y: 4},
			{X: 4, Y: 4},
			{X: 3, Y: 4},
			{X: 2, Y: 4},
			{X: 1, Y: 4},
			{X: 1, Y: 5},
		},
	},
	{
		Value: maze_solver_args{
			maze: []string{
				"## #########",
				"##   # ### #",
				"####       #",
				"# ###### ###",
				"#          #",
				"##### ######",
			},
			start: Coord{
				X: 2,
				Y: 0,
			},
			end: Coord{
				X: 5,
				Y: 5,
			},
		},
		Expected: []Coord{
			{X: 2, Y: 0},
			{X: 2, Y: 1},
			{X: 3, Y: 1},
			{X: 4, Y: 1},
			{X: 4, Y: 2},
			{X: 5, Y: 2},
			{X: 6, Y: 2},
			{X: 7, Y: 2},
			{X: 8, Y: 2},
			{X: 8, Y: 3},
			{X: 8, Y: 4},
			{X: 7, Y: 4},
			{X: 6, Y: 4},
			{X: 5, Y: 4},
			{X: 5, Y: 5},
		},
	},
}
