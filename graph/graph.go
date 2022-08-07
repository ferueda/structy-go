package graph

type Graph[T comparable] map[T][]T

func BuildUndirectedGraph[T comparable](graph Graph[T], edges [][]T) Graph[T] {
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return graph
}

func BuildDirectedGraph[T comparable](graph Graph[T], edges [][]T) Graph[T] {
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}
	return graph
}
