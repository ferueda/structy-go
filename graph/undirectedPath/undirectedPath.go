package undirectedPath

import (
	"github.com/ferueda/structy-go/graph"
)

// depth first (recursive)

func undirectedPath[T comparable](edges [][]T, src, dst T) bool {
	visited := make(map[T]bool)
	g := make(graph.Graph[T])
	graph.BuildUndirectedGraph(g, edges)
	return hasPath(g, src, dst, visited)
}

func hasPath[T comparable](graph graph.Graph[T], src, dst T, visited map[T]bool) bool {
	if src == dst {
		return true
	}
	if _, ok := visited[src]; ok {
		return false
	}

	visited[src] = true

	for _, neighbor := range graph[src] {
		if _, ok := visited[neighbor]; ok {
			continue
		}
		if hasPath(graph, neighbor, dst, visited) {
			return true
		}
	}

	return false
}
