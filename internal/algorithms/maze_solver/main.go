package maze_solver

type Point struct {
	X int
	Y int
}

var DIR = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func walk(maze [][]string, wall string, curr Point, end Point, seen [][]bool, path []Point) bool {
	// 1. Base case
	// off the map
	if curr.X < 0 || curr.X >= len(maze[0]) || curr.Y < 0 || curr.Y >= len(maze) {
		return false
	}

	// on a wall
	if maze[curr.Y][curr.X] == wall {
		return false
	}

	// at the end of the maze
	if curr.X == end.X && curr.Y == end.Y {
		path = append(path, end)
		return true
	}

	// we seen this specific tile
	if seen[curr.Y][curr.X] {
		return false
	}

	// 2. Recurse
	// Pre
	seen[curr.Y][curr.X] = true
	path = append(path, curr)

	// Recurse
	for i := 0; i < len(DIR); i++ {
		x, y := DIR[i][0], DIR[i][1]
		newPoint := Point{
			X: curr.X + x,
			Y: curr.Y + y,
		}

		if walk(maze, wall, newPoint, end, seen, path) {
			return true
		}
	}

	// Post
	path = path[:len(path)-1]

	return false
}

func Solve(maze [][]string, wall string, start Point, end Point) []Point {
	seen := make([][]bool, len(maze[0]))
	var path []Point = make([]Point, 1)

	walk(maze, wall, start, end, seen, path)

	return path
}
