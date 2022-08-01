# Has path

Write a function, hasPath, that takes in a map representing the adjacency list of a directed acyclic graph and two nodes (src, dst). The function should return a boolean indicating whether or not there exists a directed path between the source and destination nodes.

**test_00:**
```go
var graph = graph.Graph[string]{
  "f": {"g", "i"},
  "g": {"h"},
  "h": {},
  "i": {"g", "k"},
  "j": {"i"},
  "k": {},
}

hasPath(graph, "f", "k") // true
```
**test_01:**
```go
var graph = graph.Graph[string]{
  "f": {"g", "i"},
  "g": {"h"},
  "h": {},
  "i": {"g", "k"},
  "j": {"i"},
  "k": {},
}

hasPath(graph, "f", "j") // false
```
**test_02:**
```go
var graph = graph.Graph[string]{
  "f": {"g", "i"},
  "g": {"h"},
  "h": {},
  "i": {"g", "k"},
  "j": {"i"},
  "k": {},
}

hasPath(graph, "i", "h") // true
```
**test_03:**
```go
var graph = graph.Graph[string]{
  "v": {"x", "w"},
  "w": {},
  "x": {},
  "y": {"z"},
  "z": {},
}

hasPath(graph, "v", "w") // true
```
**test_04:**
```go
var graph = graph.Graph[string]{
  "v": {"x", "w"},
  "w": {},
  "x": {},
  "y": {"z"},
  "z": {},
}

hasPath(graph, "v", "z") // false
```
