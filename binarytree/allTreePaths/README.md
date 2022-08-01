# All tree paths

Write a function, allTreePaths, that takes in the root of a binary tree. The function should return a 2-Dimensional array where each subarray represents a root-to-leaf path in the tree.

The order within an individual path must start at the root and end at the leaf, but the relative order among paths in the outer array does not matter.

You may assume that the input tree is non-empty.

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

allTreePaths(&a)
// [ 
//   [ "a", "b", "d" ], 
//   [ "a", "b", "e" ], 
//   [ "a", "c", "f" ] 
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

allTreePaths(&a)
// [ 
//   [ "a", "b", "d" ], 
//   [ "a", "b", "e", "g" ], 
//   [ "a", "b", "e", "h" ], 
//   [ "a", "c", "f", "i" ] 
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

allTreePaths(&q)
// [ 
//   [ "q", "r", "t", "u", "v" ], 
//   [ "q", "s" ] 
// ] 
```
**test_03:**
```go
var z = Node{Val: "z"}

//      z

allTreePaths(&z)
// [
//   ["z"]
// ]