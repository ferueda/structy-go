# Level averages

Write a function, levelAverages, that takes in the root of a binary tree that contains number values. The function should return an array containing the average value of each level.

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

levelAverages(&a) // [ 3, 7.5, 1 ] 
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

levelAverages(&a) // [ 5, 32.5, 17.5, 2 ] 
```
**test_02:**
```go
var a = Node{Val: -1}
var b = Node{Val: -6}
var c = Node{Val: -5}
var d = Node{Val: -3}
var e = Node{Val: 0}
var f = Node{Val: 45}
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
// -3   0     45
//     /       \
//    -1       -2

levelAverages(&a) // [ -1, -5.5, 14, -1.5 ]
```
**test_03:**
```go
var q = Node{Val: 13}
var r = Node{Val: 4}
var s = Node{Val: 2}
var t = Node{Val: 9}
var u = Node{Val: 2}
var v = Node{Val: 42}

q.Left = &r
q.Right = &s
r.Right = &t
t.Left = &u
u.Right = &v

//        13
//      /   \
//     4     2
//      \
//       9
//      /
//     2
//    /
//   42

levelAverages(&q) // [ 13, 3, 9, 2, 42 ]
```
**test_04:**
```go
levelAverages(nil) // [ ]

