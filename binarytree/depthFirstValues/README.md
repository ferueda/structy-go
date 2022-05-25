# Depth first values

Write a function, depthFirstValues, that takes in the root of a binary tree. The function should return an array containing all values of the tree in depth-first order.

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

depthFirstValues(&a) // ["a", "b", "d", "e", "c", "f"]
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

depthFirstValues(&a) // ["a", "b", "d", "e", "g", "c", "f"]
```
**test_02:**
```go
var a = Node{Val: "a"}

// a

depthFirstValues(&a) // ["a"]
```
**test_03:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "d"}
var e = Node{Val: "e"}

a.Right = &b
b.Left = &c
c.Right = &d
d.Right = &e

//      a
//       \
//        b
//       /
//      c
//       \
//        d
//         \
//          e

depthFirstValues(&a) // ["a", "b", "c", "d", "e"]
```
**test_04:**
```go
depthFirstValues(nil) // []
```