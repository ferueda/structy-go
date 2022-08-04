# Closest carrot

Write a function, closestCarrot, that takes in a grid, a starting row, and a starting column. In the grid, X's are walls, O's are open spaces, and C's are carrots. The function should return a number representing the length of the shortest path from the starting position to a carrot. You may move up, down, left, or right, but cannot pass through walls (X). If there is no possible path to a carrot, then return -1.

**test_00:**
```go
var grid = [][]string{
  {"O", "O", "O", "O", "O"},
  {"O", "X", "O", "O", "O"},
  {"O", "X", "X", "O", "O"},
  {"O", "X", "C", "O", "O"},
  {"O", "X", "X", "O", "O"},
  {"C", "O", "O", "O", "O"},
}

closestCarrot(grid, 1, 2) // 4
```
**test_01:**
```go
var grid = [][]string{
  {"O", "O", "O", "O", "O"},
  {"O", "X", "O", "O", "O"},
  {"O", "X", "X", "O", "O"},
  {"O", "X", "C", "O", "O"},
  {"O", "X", "X", "O", "O"},
  {"C", "O", "O", "O", "O"},
}

closestCarrot(grid, 0, 0) // 5
```
**test_02:**
```go
var grid = [][]string{
  {"O", "O", "X", "X", "X"},
  {"O", "X", "X", "X", "C"},
  {"O", "X", "O", "X", "X"},
  {"O", "O", "O", "O", "O"},
  {"O", "X", "X", "X", "X"},
  {"O", "O", "O", "O", "O"},
  {"O", "O", "C", "O", "O"},
  {"O", "O", "O", "O", "O"},
}

closestCarrot(grid, 3, 4) // 9
```
**test_03:**
```go
var grid = [][]string{
  {"O", "O", "X", "O", "O"},
  {"O", "X", "X", "X", "O"},
  {"O", "X", "C", "C", "O"},
}

closestCarrot(grid, 1, 4) // 2
```
**test_04:**
```go
var grid = [][]string{
  {"O", "O", "X", "O", "O"},
  {"O", "X", "X", "X", "O"},
  {"O", "X", "C", "C", "O"},
}

closestCarrot(grid, 2, 0) // -1
```
**test_05:**
```go
var grid = [][]string{
  {"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "X", "X"},
  {"O", "O", "O", "O", "O", "O", "O", "O", "X", "C"},
}

closestCarrot(grid, 0, 0) // -1
```