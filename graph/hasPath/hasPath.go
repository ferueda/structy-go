package hasPath

import (
	// "github.com/ferueda/structy-go/ds"
	"github.com/ferueda/structy-go/graph"
)

// breath first (iterative)

// func hasPath(graph graph.Graph[string], src, dst string) bool {
// 	if src == dst {
// 		return true
// 	}

// 	queue := ds.NewQueue[string]()
// 	queue.Enqueue(src)

// 	for queue.Size() > 0 {
// 		curr, _ := queue.Dequeue()
// 		if curr == dst {
// 			return true
// 		}

// 		for _, neighbor := range graph[curr] {
// 			queue.Enqueue(neighbor)
// 		}
// 	}

// 	return false
// }

// depth first (recursive)

func hasPath(graph graph.Graph[string], src, dst string) bool {
	if src == dst {
		return true
	}

	for _, neighbor := range graph[src] {
		if hasPath(graph, neighbor, dst) {
			return true
		}
	}

	return false
}
