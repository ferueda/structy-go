# Minimum island

Write a function, minimumIsland, that takes in a grid containing Ws and Ls. W represents water and L represents land. The function should return the size of the smallest island. An island is a vertically or horizontally connected region of land.

You may assume that the grid contains at least one island.

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

minimumIsland(grid) // 2
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

minimumIsland(grid) // 1
```
**test_02:**
```go
var grid = [][]string{
  {"L", "L", "L"},
  {"L", "L", "L"},
  {"L", "L", "L"},
}

minimumIsland(grid) // 9
```
**test_03:**
```go
var grid = [][]string{
  {"W", "W"},
  {"L", "L"},
  {"W", "W"},
  {"W", "L"},
}

minimumIsland(grid) // 1
```
