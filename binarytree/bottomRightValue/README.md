# Bottom right value

Write a function, bottomRightValue, that takes in the root of a binary tree. The function should return the right-most value in the bottom-most level of the tree.

You may assume that the input tree is non-empty.

**test_00:**
```go
var a = Node{Val: 3}
var b = Node{Val: 11}
var c = Node{Val: 10}
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
//   11     10
//  / \      \
// 4   -2     1

bottomRightValue(&a) // 1
```
**test_01:**
```go
var a = Node{Val: -1}
var b = Node{Val: -6}
var c = Node{Val: -5}
var d = Node{Val: -3}
var e = Node{Val: -4}
var f = Node{Val: -13}
var g = Node{Val: -2}
var h = Node{Val: 6}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f
e.Left = &g
e.Right = &h

//        -1
//      /   \
//    -6    -5
//   /  \     \
// -3   -4   -13
//     / \       
//    -2  6

bottomRightValue(&a) // 6
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
var h = Node{Val: 6}
var i = Node{Val: 7}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f
e.Left = &g
e.Right = &h
f.Left = &i

//        -1
//      /   \
//    -6    -5
//   /  \     \
// -3   -4   -13
//     / \    /   
//    -2  6  7 

bottomRightValue(&a) // -> 7
```
**test_03:**
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
b.Right = &d
d.Left = &e
e.Right = &f

//      a
//    /   \ 
//   b     c
//    \
//     d
//    /
//   e
//  /
// f

bottomRightValue(&a) // "f"
```
**test_04:**
```go
var a = Node{Val: "a"}

//      "a"

bottomRightValue(&a) // "a"
