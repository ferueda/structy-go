# Tree value count

Write a function, treeValueCount, that takes in the root of a binary tree and a target value. The function should return the number of times that the target occurs in the tree.

**test_00:**
```go
var a = Node{Val: 12}
var b = Node{Val: 6}
var c = Node{Val: 6}
var d = Node{Val: 4}
var e = Node{Val: 6}
var f = Node{Val: 12}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f

//      12
//    /   \
//   6     6
//  / \     \
// 4   6     12

treeValueCount(&a, 6) // 3
```
**test_01:**
```go
var a = Node{Val: 12}
var b = Node{Val: 6}
var c = Node{Val: 6}
var d = Node{Val: 4}
var e = Node{Val: 6}
var f = Node{Val: 12}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f

//      12
//    /   \
//   6     6
//  / \     \
// 4   6     12

treeValueCount(&a, 12) // 2
```
**test_02:**
```go
var a = Node{Val: 7}
var b = Node{Val: 5}
var c = Node{Val: 1}
var d = Node{Val: 8}
var e = Node{Val: 7}
var f = Node{Val: 7}
var g = Node{Val: 1}
var h = Node{Val: 1}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f
e.Left = &g
f.Right = &h

//      7
//    /   \
//   5     1
//  / \     \
// 1   8     7
//    /       \
//   1         1

treeValueCount(&a, 1) // -> 4
```
**test_03:**
```go
var a = Node{Val: 7}
var b = Node{Val: 5}
var c = Node{Val: 1}
var d = Node{Val: 8}
var e = Node{Val: 7}
var f = Node{Val: 7}
var g = Node{Val: 1}
var h = Node{Val: 1}

a.Left = &b
a.Right = &c
b.Left = &d
b.Right = &e
c.Right = &f
e.Left = &g
f.Right = &h

//      7
//    /   \
//   5     1
//  / \     \
// 1   8     7
//    /       \
//   1         1

treeValueCount(&a, 9) // 0
```
**test_04:**
```go
treeValueCount(nil, 42) // 0
