package closestCarrot

import (
	"fmt"

	"github.com/ferueda/structy-go/ds"
)

type position struct {
	row  int
	col  int
	dist int
}

func (p position) string() string {
	return fmt.Sprintf("%v,%v", p.row, p.col)
}

// breadth first (iterative)

func closestCarrot(grid [][]string, row, col int) int {
	visited := make(map[string]bool)
	queue := ds.NewQueue[position]()
	queue.Enqueue(position{row: row, col: col, dist: 0})

	for queue.Size() > 0 {
		currPos, _ := queue.Dequeue()

		if currPos.row < 0 || currPos.row >= len(grid) {
			continue
		}
		if currPos.col < 0 || currPos.col >= len(grid[currPos.row]) {
			continue
		}
		if grid[currPos.row][currPos.col] == "X" {
			continue
		}
		if _, ok := visited[currPos.string()]; ok {
			continue
		}

		if grid[currPos.row][currPos.col] == "C" {
			return currPos.dist
		}

		visited[currPos.string()] = true

		queue.Enqueue(position{row: currPos.row + 1, col: currPos.col, dist: currPos.dist + 1})
		queue.Enqueue(position{row: currPos.row - 1, col: currPos.col, dist: currPos.dist + 1})
		queue.Enqueue(position{row: currPos.row, col: currPos.col + 1, dist: currPos.dist + 1})
		queue.Enqueue(position{row: currPos.row, col: currPos.col - 1, dist: currPos.dist + 1})
	}

	return -1
}
