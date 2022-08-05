package longestPath

import (
	"github.com/ferueda/structy-go/graph"
	"github.com/ferueda/structy-go/utils"
)

type item[T comparable] struct {
	node T
	dist int
}

// depth first (recursive)

func longestPath[T comparable](graph graph.Graph[T]) int {
	longest := 0

	for node := range graph {
		longest = utils.Max(longest, traverse(graph, item[T]{node: node, dist: 0}))
	}

	return longest
}

func traverse[T comparable](graph graph.Graph[T], i item[T]) int {
	if len(graph[i.node]) == 0 {
		return i.dist
	}

	maxDist := 0

	for _, neighbor := range graph[i.node] {
		maxDist = utils.Max(maxDist, traverse(graph, item[T]{node: neighbor, dist: i.dist + 1}))
	}

	return maxDist
}
