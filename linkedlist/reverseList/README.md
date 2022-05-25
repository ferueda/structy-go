# Reverse list

Write a function, reverseList, that takes in the head of a linked list as an argument. The function should reverse the order of the nodes in the linked list in-place and return the new head of the reversed linked list.

**test_00:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "d"}
var e = Node{Val: "e"}
var f = Node{Val: "f"}

a.next = &b
b.next = &c
c.next = &d
d.next = &e
e.next = &f

// a -> b -> c -> d -> e -> f

reverseList(a) // f -> e -> d -> c -> b -> a
```
**test_01:**
```go
var x = Node{Val: "x"}
var y = Node{Val: "y"}

x.next = &y

// x -> y

reverseList(x) // -> y -> x
```
**test_02:**
```go
var p = Node{Val: "p"}

// p

reverseList(p) // -> p
```
