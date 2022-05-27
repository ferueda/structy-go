# Tree min value

Write a function, treeMinValue, that takes in the root of a binary tree that contains number values. The function should return the minimum value within the tree.

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

treeMinValue(&a) // -2
```
**test_01:**
```go
var a = Node{Val: 5}
var b = Node{Val: 11}
var c = Node{Val: 3}
var d = Node{Val: 4}
var e = Node{Val: 14}
var f = Node{Val: 12}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f

//       5
//    /    \
//   11     3
//  / \      \
// 4   14     12

treeMinValue(&a) // 3
```
**test_02:**
```go
var a = Node{Val: -1}
var b = Node{Val: -6}
var c = Node{Val: -5}
var d = Node{Val: -3}
var e = Node{Val: -4}
var f = Node{Val: -13}
var g = Node{Val: -2}
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
// -3   -4   -13
//     /       \
//    -2       -2

treeMinValue(&a) // -13
```
**test_03:**
```go
var a = Node{Val: 42}

// 42

treeMinValue(&a) // 42
```