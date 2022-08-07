package shortestPath

import (
	"github.com/ferueda/structy-go/ds"
	"github.com/ferueda/structy-go/graph"
)

type queueItem[T comparable] struct {
	node     T
	distance int
}

// breadth first (iterative)

func shortestPath[T comparable](edges [][]T, src, dst T) int {
	g := make(graph.Graph[T])
	graph.BuildUndirectedGraph(g, edges)
	visited := make(map[T]bool)
	queue := ds.NewQueue[queueItem[T]]()
	queue.Enqueue(queueItem[T]{node: src, distance: 0})

	for queue.Size() > 0 {
		item, _ := queue.Dequeue()
		if item.node == dst {
			return item.distance
		}

		visited[item.node] = true

		for _, neighbor := range g[item.node] {
			if _, ok := visited[neighbor]; !ok {
				queue.Enqueue(queueItem[T]{node: neighbor, distance: item.distance + 1})
			}
		}
	}

	return -1
}
