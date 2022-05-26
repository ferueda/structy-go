# Tree sum

Write a function, treeSum, that takes in the root of a binary tree that contains number values. The function should return the total sum of all values in the tree.

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

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f

treeSum(&a) // ["a", "b", "d", "e", "c", "f"]
```
**test_01:**
```go
var a = Node{Val: 1}
var b = Node{Val: 6}
var c = Node{Val: 0}
var d = Node{Val: 3}
var e = Node{Val: -6}
var f = Node{Val: 2}
var g = Node{Val: 2}
var h = Node{Val: 2}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f
e.Left = &g
f.Right = &h

//      1
//    /   \
//   6     0
//  / \     \
// 3   -6    2
//    /       \
//   2         2

treeSum(&a) // 10
```
**test_02:**
```go
treeSum(nil) // 0
```