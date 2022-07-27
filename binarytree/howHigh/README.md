# How high

Write a function, howHigh, that takes in the root of a binary tree. The function should return a number representing the height of the tree.

The height of a binary tree is defined as the maximal number of edges from the root node to any leaf node.

If the tree is empty, return -1.

**test_00:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "d"}
var e = Node{Val: "e"}
var f = Node{Val: "f"}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f

howHigh(&a) // 2
```
**test_01:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "d"}
var e = Node{Val: "e"}
var f = Node{Val: "f"}
var g = Node{Val: "g"}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f
e.Left = &g

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f
//    /
//   g

howHigh(&a) // 3
```
**test_02:**
```go
var a = Node{Val: "a"}
var c = Node{Val: "c"}

a.Right = &c

//      a
//       \
//        c

howHigh(&a) // -> 1
```
**test_03:**
```go
var a = Node{Val: "a"}

//      a

howHigh(&a) // 0
```
**test_04:**
```go
howHigh(nil) // -1
