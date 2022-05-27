# Max root to leaf path sum

Write a function, maxPathSum, that takes in the root of a binary tree that contains number values. The function should return the maximum sum of any root to leaf path within the tree.

You may assume that the input tree is non-empty.

**test_00:**
```go
var a = Node{Val: 3}
var b = Node{Val: 11}
var c = Node{Val: 4}
var d = Node{Val: 4}
var e = Node{Val: -2}
var f = Node{Val: 1}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f

//       3
//    /    \
//   11     4
//  / \      \
// 4   -2     1

maxPathSum(&a) // 18
```
**test_01:**
```go
var a = Node{Val: 5}
var b = Node{Val: 11}
var c = Node{Val: 54}
var d = Node{Val: 20}
var e = Node{Val: 15}
var f = Node{Val: 1}
var g = Node{Val: 3}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
e.Left = &f
e.Right = &g

//        5
//     /    \
//    11    54
//  /   \
// 20   15
//      / \
//     1  3

maxPathSum(&a) // 59
```
**test_02:**
```go
var a = Node{Val: -1}
var b = Node{Val: -6}
var c = Node{Val: -5}
var d = Node{Val: -3}
var e = Node{Val: 0}
var f = Node{Val: -13}
var g = Node{Val: -1}
var h = Node{Val: -2}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f
e.Left = &g
f.Right = &h

//        -1
//      /   \
//    -6    -5
//   /  \     \
// -3   0    -13
//     /       \
//    -1       -2

maxPathSum(&a) // -8
```
**test_03:**
```go
var a = Node{Val: 42}

// 42

maxPathSum(&a) // 42
```