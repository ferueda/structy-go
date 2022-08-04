package minimumIsland

import (
	"fmt"
	"math"

	"github.com/ferueda/structy-go/utils"
	// "github.com/ferueda/structy-go/ds"
)

type position struct {
	row int
	col int
}

func (p position) string() string {
	return fmt.Sprintf("%v,%v", p.row, p.col)
}

func minimumIsland(grid [][]string) int {
	visited := make(map[string]bool)
	smallest := math.MaxInt

	for row := range grid {
		for col := range grid[row] {
			pos := position{row: row, col: col}
			if _, ok := visited[pos.string()]; !ok && grid[row][col] == "L" {
				smallest = utils.Min(smallest, traverse(grid, pos, visited))
			}
		}
	}

	return smallest
}

// depth first (recursive)

func traverse(grid [][]string, pos position, visited map[string]bool) int {
	rowInBounds := 0 <= pos.row && pos.row < len(grid)
	if !rowInBounds {
		return 0
	}
	colInBounds := 0 <= pos.col && pos.col < len(grid[pos.row])
	if !colInBounds {
		return 0
	}
	if grid[pos.row][pos.col] == "W" {
		return 0
	}
	if _, ok := visited[pos.string()]; ok {
		return 0
	}

	count := 1
	visited[pos.string()] = true

	count += traverse(grid, position{row: pos.row + 1, col: pos.col}, visited)
	count += traverse(grid, position{row: pos.row - 1, col: pos.col}, visited)
	count += traverse(grid, position{row: pos.row, col: pos.col + 1}, visited)
	count += traverse(grid, position{row: pos.row, col: pos.col - 1}, visited)

	return count
}

// breadth first (iterative)

// func traverse(grid [][]string, pos position, visited map[string]bool) int {
// 	queue := ds.NewQueue[position]()
// 	queue.Enqueue(pos)
// 	count := 0

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

// 		count += 1
// 		visited[currPos.string()] = true

// 		queue.Enqueue(position{row: currPos.row + 1, col: currPos.col})
// 		queue.Enqueue(position{row: currPos.row - 1, col: currPos.col})
// 		queue.Enqueue(position{row: currPos.row, col: currPos.col + 1})
// 		queue.Enqueue(position{row: currPos.row, col: currPos.col - 1})

// 	}

// 	return count
// }
