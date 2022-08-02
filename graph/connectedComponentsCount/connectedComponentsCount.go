package connectedComponentsCount

import (
	"github.com/ferueda/structy-go/graph"
)

// depth first (recursive)

func connectedComponentsCount[T comparable](graph graph.Graph[T]) int {
	visited := make(map[T]bool)
	connected := 0

	for node := range graph {
		if _, ok := visited[node]; !ok {
			traverse(graph, node, visited)
			connected += 1
		}
	}

	return connected
}

func traverse[T comparable](graph graph.Graph[T], src T, visited map[T]bool) {
	if _, ok := visited[src]; ok {
		return
	}

	visited[src] = true

	for _, neighbor := range graph[src] {
		if _, ok := visited[neighbor]; !ok {
			traverse(graph, neighbor, visited)
		}
	}
}
