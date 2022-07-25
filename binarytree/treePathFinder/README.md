# Tree path finder

Write a function, pathFinder, that takes in the root of a binary tree and a target value. The function should return an array representing a path to the target value. If the target value is not found in the tree, then return null.

You may assume that the tree contains unique values.

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

pathFinder(&a, "e") // [ 'a', 'b', 'e' ]
```
**test_01:**
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

pathFinder(&a, "p") // nil
```
**test_02:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "d"}
var e = Node{Val: "e"}
var f = Node{Val: "f"}
var f = Node{Val: "g"}
var f = Node{Val: "h"}

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

pathFinder(&a, "c") // -> ['a', 'c']
```
**test_03:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "d"}
var e = Node{Val: "e"}
var f = Node{Val: "f"}
var f = Node{Val: "g"}
var f = Node{Val: "h"}

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

pathFinder(&a, "h") // ['a', 'c', 'f', 'h']
```
**test_04:**
```go
var x = Node{Val: "x"}

//  x

pathFinder(&x, "x") // ['x']
```
**test_05:**
```go
pathFinder(nil, "x") // nil
```
**test_06:**
```go

var root = Node{Val: 0}
var curr = root
for  i := 1; i <= 6000; i++ {
  curr.Right = Node{Val: i}
  curr = curr.Right
}

//      0
//       \
//        1
//         \
//          2
//           \
//            3
//             .
//              .
//               .
//              5999
//                \
//                6000

pathFinder(root, 3451) // [0, 1, 2, 3, ..., 3450, 3451]
```