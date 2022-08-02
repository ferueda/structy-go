package largestComponent

import (
	// "github.com/ferueda/structy-go/ds"
	"github.com/ferueda/structy-go/graph"
	"github.com/ferueda/structy-go/utils"
)

func largestComponent[T comparable](graph graph.Graph[T]) int {
	visited := make(map[T]bool)
	largest := 0

	for node := range graph {
		if _, ok := visited[node]; !ok {
			size := traverse(graph, node, visited)
			largest = utils.Max(size, largest)
		}
	}

	return largest
}

// breadth first (iterative)

// func traverse[T comparable](graph graph.Graph[T], src T, visited map[T]bool) int {
// 	queue := ds.NewQueue[T]()
// 	queue.Enqueue(src)
// 	count := 0

// 	for queue.Size() > 0 {
// 		curr, _ := queue.Dequeue()
// 		if _, ok := visited[curr]; ok {
// 			continue
// 		}

// 		visited[curr] = true
// 		count += 1

// 		for _, neighbor := range graph[curr] {
// 			if _, ok := visited[neighbor]; !ok {
// 				queue.Enqueue(neighbor)
// 			}
// 		}
// 	}

// 	return count
// }

// depth first (recursive)

func traverse[T comparable](graph graph.Graph[T], src T, visited map[T]bool) int {
	if _, ok := visited[src]; ok {
		return 0
	}

	visited[src] = true
	size := 1

	for _, neighbor := range graph[src] {
		size += traverse(graph, neighbor, visited)
	}

	return size
}
