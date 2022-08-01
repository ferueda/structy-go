# Tree levels

Write a function, treeLevels, that takes in the root of a binary tree. The function should return a 2-Dimensional array where each subarray represents a level of the tree.

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

//       a
//    /    \
//   b      c
//  / \      \
// d   e     f

treeLevels(&a)
// [
//   ["a"],
//   ["b", "c"],
//   ["d", "e", "f"]
// ]
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
var i = Node{Val: "i"}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f
e.Left = &g
e.Right = &h
f.Left = &i

//         a
//      /    \
//     b      c
//   /  \      \
//  d    e      f
//      / \    /   
//     g  h   i 

treeLevels(&a)
// [
//   ["a"],
//   ["b", "c"],
//   ["d", "e", "f"],
//   ["g", "h", "i"]
// ]
```
**test_02:**
```go
var q = Node{Val: "q"}
var r = Node{Val: "r"}
var s = Node{Val: "s"}
var t = Node{Val: "t"}
var u = Node{Val: "u"}
var v = Node{Val: "v"}

q.Left = &r
q.Right = &s
r.Right = &t
t.Left = &u
u.Right = &v

//      q
//    /   \ 
//   r     s
//    \
//     t
//    /
//   u
//  /
// v

treeLevels(&q)
// [
//   ["q"],
//   ["r", "s"],
//   ["t"],
//   ["u"],
//   ["v"]
// ]
```
**test_03:**
```go
treeLevels(nil) // []
