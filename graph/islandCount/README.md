# Island count

Write a function, islandCount, that takes in a grid containing Ws and Ls. W represents water and L represents land. The function should return the number of islands on the grid. An island is a vertically or horizontally connected region of land.

**test_00:**
```go
var grid = [][]string{
  {"W", "L", "W", "W", "W"},
  {"W", "L", "W", "W", "W"},
  {"W", "W", "W", "L", "W"},
  {"W", "W", "L", "L", "W"},
  {"L", "W", "W", "L", "L"},
  {"L", "L", "W", "W", "W"},
}

islandCount(grid) // 3
```
**test_01:**
```go
var grid = [][]string{
  {"L", "W", "W", "L", "W"},
  {"L", "W", "W", "L", "L"},
  {"W", "L", "W", "L", "W"},
  {"W", "W", "W", "W", "W"},
  {"W", "W", "L", "L", "L"},
}

islandCount(grid) // 4
```
**test_02:**
```go
var grid = [][]string{
  {"L", "L", "L"},
  {"L", "L", "L"},
  {"L", "L", "L"},
}

islandCount(grid) // 1
```
**test_03:**
```go
var grid = [][]string{
  {"W", "W"},
  {"W", "W"},
  {"W", "W"},
}

islandCount(grid) // 0
```
