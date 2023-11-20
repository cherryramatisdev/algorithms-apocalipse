package maze_solver_test

import (
	"testing"

	"github.com/cherryramatisdev/algorithms-apocalipse/internal/algorithms/maze_solver"
	"gotest.tools/assert"
)

func drawPath(data [][]string, path []maze_solver.Point) [][]string {
	// Update the path in the 2D string slice
	for _, p := range path {
		if p.Y < len(data) && p.X < len(data[p.Y]) {
			data[p.Y][p.X] = "*"

		}

	}

	return data
}

func HappyPath(t *testing.T) {
	maze := [][]string{
		{"xxxxxxxxxx x"},
		{"x        x x"},
		{"x        x x"},
		{"x xxxxxxxx x"},
		{"x          x"},
		{"x xxxxxxxxxx"},
	}

	mazeResult := []maze_solver.Point{
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
	}

	result := maze_solver.Solve(maze, "x", maze_solver.Point{X: 10, Y: 0}, maze_solver.Point{X: 1, Y: 5})

	assert.Equal(t, drawPath(maze, result), drawPath(maze, mazeResult))
}
