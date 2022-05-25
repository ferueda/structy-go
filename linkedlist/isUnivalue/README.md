# Is univalue

Write a function, isUnivalue, that takes in the head of a linked list as an argument. The function should return a boolean indicating whether or not the linked list contains exactly one unique value.

You may assume that the input list is non-empty.

**test_00:**
```go
var a = Node{Val: 7}
var b = Node{Val: 7}
var c = Node{Val: 7}

a.next = &b
b.next = &c


isUnivalue(a) // true
```
**test_01:**
```go
var a = Node{Val: 7}
var b = Node{Val: 7}
var c = Node{Val: 4}

a.next = &b
b.next = &c


isUnivalue(a) // false
```
**test_02:**
```go
var a = Node{Val: 2}
var b = Node{Val: 2}
var c = Node{Val: 2}
var d = Node{Val: 2}
var e = Node{Val: 2}

a.next = &b
b.next = &c
c.next = &d
d.next = &e


isUnivalue(a) // true
```
**test_03:**
```go
var a = Node{Val: 2}
var b = Node{Val: 2}
var c = Node{Val: 2}
var d = Node{Val: 3}
var e = Node{Val: 2}

a.next = &b
b.next = &c
c.next = &d
d.next = &e


isUnivalue(a) // false
```
**test_04:**
```go
var a = Node{Val: 1}

isUnivalue(a) // true
```