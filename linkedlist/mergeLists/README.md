# Merge lists

Write a function, mergeLists, that takes in the head of two sorted linked lists as arguments. The function should merge the two lists together into single sorted linked list. The function should return the head of the merged linked list.

Do this in-place, by mutating the original Nodes.

You may assume that both input lists are non-empty and contain increasing sorted numbers.

**test_00:**
```go
var a = Node{Val: 5}
var b = Node{Val: 7}
var c = Node{Val: 10}
var d = Node{Val: 12}
var e = Node{Val: 20}
var f = Node{Val: 28}

a.next = &b
b.next = &c
c.next = &d
d.next = &e
e.next = &f

// 5 -> 7 -> 10 -> 12 -> 20 -> 28

var q = Node{Val: 6}
var r = Node{Val: 8}
var s = Node{Val: 9}
var t = Node{Val: 25}


q.next = &r
r.next = &s
s.next = &t

// 6 -> 8 -> 9 -> 25

mergeLists(a, q) // 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 12 -> 20 -> 25 -> 28
```
**test_01:**
```go
var a = Node{Val: 5}
var b = Node{Val: 7}
var c = Node{Val: 10}
var d = Node{Val: 12}
var e = Node{Val: 20}
var f = Node{Val: 28}

a.next = &b
b.next = &c
c.next = &d
d.next = &e
e.next = &f

// 5 -> 7 -> 10 -> 12 -> 20 -> 28

var q = Node{Val: 1}
var r = Node{Val: 8}
var s = Node{Val: 9}
var t = Node{Val: 10}


q.next = &r
r.next = &s
s.next = &t

// 1 -> 8 -> 9 -> 10

mergeLists(a, q) // 1 -> 5 -> 7 -> 8 -> 9 -> 10 -> 10 -> 12 -> 20 -> 28
```
**test_02:**
```go
var h = Node{Val: 30}

// 30

var p = Node{Val: 15}
var q = Node{Val: 67}

p.next = &q

// 15 -> 67

mergeLists(h, p) // 15 -> 30 -> 67
```