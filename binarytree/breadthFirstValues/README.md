# Breadth first values

Write a function, breadthFirstValues, that takes in the root of a binary tree. The function should return an array containing all values of the tree in breadth-first order.

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

breadthFirstValues(&a) // ["a", "b", "c", "d", "e", "f"]
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

breadthFirstValues(&a) // ["a", "b", "c", "d", "e", "f", "g", "h"]
```
**test_02:**
```go
var a = Node{Val: "a"}

// a

breadthFirstValues(&a) // ["a"]
```
**test_03:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "d"}
var e = Node{Val: "e"}
var x = Node{Val: "x"}

a.Right = &b
b.Left = &c
c.Left = &x
c.Right = &d
d.Right = &e

//      a
//       \
//        b
//       /
//      c
//    /  \
//   x    d
//         \
//          e

breadthFirstValues(&a) // ["a", "b", "c", "x", "d", "e"]
```
**test_04:**
```go
breadthFirstValues(nil) // []
```