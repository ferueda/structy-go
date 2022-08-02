# Undirected path

Write a function, undirectedPath, that takes in an array of edges for an undirected graph and two nodes (nodeA, nodeB). The function should return a boolean indicating whether or not there exists a path between nodeA and nodeB.

**test_00:**
```go
var edges = [][]string{
  {"i", "j"},
  {"k", "i"},
  {"m", "k"},
  {"k", "l"},
  {"o", "n"},
}

undirectedPath(edges, "j", "m") // true
```
**test_01:**
```go
var edges = [][]string{
  {"i", "j"},
  {"k", "i"},
  {"m", "k"},
  {"k", "l"},
  {"o", "n"},
}

undirectedPath(edges, "m", "j") // true
```
**test_02:**
```go
var edges = [][]string{
  {"i", "j"},
  {"k", "i"},
  {"m", "k"},
  {"k", "l"},
  {"o", "n"},
}

undirectedPath(edges, "l", "j") // true
```
**test_03:**
```go
var edges = [][]string{
  {"i", "j"},
  {"k", "i"},
  {"m", "k"},
  {"k", "l"},
  {"o", "n"},
}

undirectedPath(edges, "k", "o") // false
```
**test_04:**
```go
var edges = [][]string{
  {"i", "j"},
  {"k", "i"},
  {"m", "k"},
  {"k", "l"},
  {"o", "n"},
}

undirectedPath(edges, "i", "o") // false
```
**test_05:**
```go
var edges = [][]string{
  {"b", "a"},
  {"c", "a"},
  {"b", "c"},
  {"q", "r"},
  {"q", "s"},
  {"q", "u"},
  {"q", "t"},
}

undirectedPath(edges, "a", "b") // true
```
**test_06:**
```go
var edges = [][]string{
  {"b", "a"},
  {"c", "a"},
  {"b", "c"},
  {"q", "r"},
  {"q", "s"},
  {"q", "u"},
  {"q", "t"},
}

undirectedPath(edges, "a", "c") // true
```
**test_07:**
```go
var edges = [][]string{
  {"b", "a"},
  {"c", "a"},
  {"b", "c"},
  {"q", "r"},
  {"q", "s"},
  {"q", "u"},
  {"q", "t"},
}

undirectedPath(edges, "r", "t") // true
```
**test_08:**
```go
var edges = [][]string{
  {"b", "a"},
  {"c", "a"},
  {"b", "c"},
  {"q", "r"},
  {"q", "s"},
  {"q", "u"},
  {"q", "t"},
}

undirectedPath(edges, "r", "b") // false
```
**test_09:**
```go
var edges = [][]string{
  {"s", "r"},
  {"t", "q"},
  {"q", "r"},
}

undirectedPath(edges, "r", "t") // true
```