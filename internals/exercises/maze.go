package exercises

type Maze = []string
type Coord struct {
	X int
	Y int
}

func maze_solver(m Maze, start, end Coord) []Coord {
	path := []Coord{}
	seen := []Coord{}

	walk(&m, start, end, &path, &seen)
	return path
}

func walk(m *Maze, pos, end Coord, path, seen *[]Coord) bool {
	if pos.X == end.X && pos.Y == end.Y {
		*path = append(*path, pos)
		return true
	}

	flag := false
	if not_seen(seen, pos) {
		*path = append(*path, pos)
		var n_pos Coord
		if n_pos = (Coord{X: pos.X + 1, Y: pos.Y}); !flag && (*m)[n_pos.Y][n_pos.X] != '#' {
			flag = walk(m, n_pos, end, path, seen)
		}
		if n_pos = (Coord{X: pos.X - 1, Y: pos.Y}); !flag && n_pos.X >= 0 && (*m)[n_pos.Y][n_pos.X] != '#' {
			flag = walk(m, n_pos, end, path, seen)
		}
		if n_pos = (Coord{X: pos.X, Y: pos.Y + 1}); !flag && (*m)[n_pos.Y][n_pos.X] != '#' {
			flag = walk(m, n_pos, end, path, seen)
		}
		if n_pos = (Coord{X: pos.X, Y: pos.Y - 1}); !flag && n_pos.Y >= 0 && (*m)[n_pos.Y][n_pos.X] != '#' {
			flag = walk(m, n_pos, end, path, seen)
		}
		if !flag {
			*path = (*path)[:len(*path)-1]
		}
	}
	return flag
}

func not_seen(s *[]Coord, p Coord) bool {
	for _, c := range *s {
		if c.X == p.X && c.Y == p.Y {
			return false
		}
	}
	*s = append(*s, p)
	return true
}
