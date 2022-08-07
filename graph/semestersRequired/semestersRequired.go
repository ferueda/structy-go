package semestersRequired

import (
	"github.com/ferueda/structy-go/graph"
	"github.com/ferueda/structy-go/utils"
)

// depth first (recursive)

func semestersRequired(prereqs [][]int) int {
	g := make(graph.Graph[int])
	graph.BuildDirectedGraph(g, prereqs)
	max := 1

	for course := range g {
		max = utils.Max(max, traverse(g, course))
	}

	return max
}

func traverse(graph graph.Graph[int], src int) int {
	max := 0

	for _, nextCourse := range graph[src] {
		max = utils.Max(max, traverse(graph, nextCourse))
	}

	return max + 1
}
