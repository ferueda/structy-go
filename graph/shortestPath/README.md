# Shortest path

Write a function, shortestPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB). The function should return the length of the shortest path between A and B. Consider the length as the number of edges in the path, not the number of nodes. If there is no path between A and B, then return -1.

**test_00:**
```go
var edges = [][]string{
  {"w", "x"},
  {"x", "y"},
  {"z", "y"},
  {"z", "v"},
  {"w", "v"},
}

shortestPath(edges, "w", "z") // 2
```
**test_01:**
```go
var edges = [][]string{
  {"w", "x"},
  {"x", "y"},
  {"z", "y"},
  {"z", "v"},
  {"w", "v"},
}

shortestPath(edges, "y", "x") // 1
```
**test_02:**
```go
var edges = [][]string{
  {"a", "c"},
  {"a", "b"},
  {"c", "b"},
  {"c", "d"},
  {"b", "d"},
  {"e", "d"},
  {"g", "f"},
}

shortestPath(edges, "a", "e") // 3
```
**test_03:**
```go
var edges = [][]string{
  {"a", "c"},
  {"a", "b"},
  {"c", "b"},
  {"c", "d"},
  {"b", "d"},
  {"e", "d"},
  {"g", "f"},
}

shortestPath(edges, "e", "c") // 2
```
**test_04:**
```go
var edges = [][]string{
  {"a", "c"},
  {"a", "b"},
  {"c", "b"},
  {"c", "d"},
  {"b", "d"},
  {"e", "d"},
  {"g", "f"},
}

shortestPath(edges, "b", "g") // -1
```
**test_05:**
```go
var edges = [][]string{
  {"c", "n"},
  {"c", "e"},
  {"c", "s"},
  {"c", "w"},
  {"w", "e"},
}

shortestPath(edges, "w", "e") // 1
```
**test_06:**
```go
var edges = [][]string{
  {"c", "n"},
  {"c", "e"},
  {"c", "s"},
  {"c", "w"},
  {"w", "e"},
}

shortestPath(edges, "n", "e") // 2
```
**test_07:**
```go
var edges = [][]string{
  {"m", "n"},
  {"n", "o"},
  {"o", "p"},
  {"p", "q"},
  {"t", "o"},
  {"r", "q"},
  {"r", "s"},
}

shortestPath(edges, "m", "s") // 6
```