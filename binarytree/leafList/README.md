# Level averages

Write a function, leafList, that takes in the root of a binary tree and returns an array containing the values of all leaf nodes in left-to-right order.

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

leafList(&a) // [ "d", "e", "f" ] 
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
var h = Node{Val: "h"}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f
e.Left = &g
f.Right = &h

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f
//    /       \
//   g         h

leafList(&a) // [ "d", "g", "h" ]
```
**test_02:**
```go
var a = Node{Val: "5"}
var b = Node{Val: "11"}
var c = Node{Val: "54"}
var d = Node{Val: "20"}
var e = Node{Val: "15"}
var f = Node{Val: "1"}
var g = Node{Val: "3"}

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

leafList(&a) // [ "20", "1", "3", "54" ]
```
**test_03:**
```go
var x = Node{Val: "x"}

//      x

leafList(&x) // [ "x" ]
```
**test_04:**
```go
leafList(nil) // [ ]

