# Has path

Write a function, hasPath, that takes in a map representing the adjacency list of a directed acyclic graph and two nodes (src, dst). The function should return a boolean indicating whether or not there exists a directed path between the source and destination nodes.

**test_00:**
```go
var g = graph.Graph[string]{
  "f": {"g", "i"},
  "g": {"h"},
  "h": {},
  "i": {"g", "k"},
  "j": {"i"},
  "k": {},
}

hasPath(g, "f", "k") // true
```
**test_01:**
```go
var g = graph.Graph[string]{
  "f": {"g", "i"},
  "g": {"h"},
  "h": {},
  "i": {"g", "k"},
  "j": {"i"},
  "k": {},
}

hasPath(g, "f", "j") // false
```
**test_02:**
```go
var g = graph.Graph[string]{
  "f": {"g", "i"},
  "g": {"h"},
  "h": {},
  "i": {"g", "k"},
  "j": {"i"},
  "k": {},
}

hasPath(g, "i", "h") // true
```
**test_03:**
```go
var g = graph.Graph[string]{
  "v": {"x", "w"},
  "w": {},
  "x": {},
  "y": {"z"},
  "z": {},
}

hasPath(g, "v", "w") // true
```
**test_04:**
```go
var g = graph.Graph[string]{
  "v": {"x", "w"},
  "w": {},
  "x": {},
  "y": {"z"},
  "z": {},
}

hasPath(g, "v", "z") // false
```
