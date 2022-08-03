package islandCount

import (
	"fmt"
	// "github.com/ferueda/structy-go/ds"
)

type position struct {
	row int
	col int
}

func (p position) string() string {
	return fmt.Sprintf("%v,%v", p.row, p.col)
}

func islandCount(grid [][]string) int {
	visited := make(map[string]bool)
	count := 0

	for row := range grid {
		for col := range grid[row] {
			pos := position{row: row, col: col}
			if _, ok := visited[pos.string()]; !ok && grid[row][col] == "L" {
				traverse(grid, pos, visited)
				count += 1
			}
		}
	}

	return count
}

// depth first (recursive)

func traverse(grid [][]string, pos position, visited map[string]bool) {
	rowInBounds := 0 <= pos.row && pos.row < len(grid)
	if !rowInBounds {
		return
	}
	colInBounds := 0 <= pos.col && pos.col < len(grid[pos.row])
	if !colInBounds {
		return
	}
	if grid[pos.row][pos.col] == "W" {
		return
	}
	if _, ok := visited[pos.string()]; ok {
		return
	}

	visited[pos.string()] = true

	traverse(grid, position{row: pos.row + 1, col: pos.col}, visited)
	traverse(grid, position{row: pos.row - 1, col: pos.col}, visited)
	traverse(grid, position{row: pos.row, col: pos.col + 1}, visited)
	traverse(grid, position{row: pos.row, col: pos.col - 1}, visited)
}

// breadth first (iterative)

// func traverse(grid [][]string, pos position, visited map[string]bool) {
// 	queue := ds.NewQueue[position]()
// 	queue.Enqueue(pos)

// 	for queue.Size() > 0 {
// 		currPos, _ := queue.Dequeue()
// 		if currPos.row < 0 || currPos.row >= len(grid) {
// 			continue
// 		}
// 		if currPos.col < 0 || currPos.col >= len(grid[currPos.row]) {
// 			continue
// 		}
// 		if grid[currPos.row][currPos.col] == "W" {
// 			continue
// 		}
// 		if _, ok := visited[currPos.string()]; ok {
// 			continue
// 		}

// 		visited[currPos.string()] = true

// 		queue.Enqueue(position{row: currPos.row + 1, col: currPos.col})
// 		queue.Enqueue(position{row: currPos.row - 1, col: currPos.col})
// 		queue.Enqueue(position{row: currPos.row, col: currPos.col + 1})
// 		queue.Enqueue(position{row: currPos.row, col: currPos.col - 1})

// 	}
// }
