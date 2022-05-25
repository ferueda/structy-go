# Linked list values

Write a function, linkedListValues, that takes in the head of a linked list as an argument. The function should return an array containing all values of the nodes in the linked list.

**test_00:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "d"}

a.next = &b
b.next = &c
c.next = &d

// a -> b -> c -> d

linkedListValues(a) // -> [ "a", "b", "c", "d" ]
```
**test_01:**
```go
var x = Node{Val: "x"}
var y = Node{Val: "y"}

x.next = &y

// x -> y

linkedListValues(x) // -> [ "x", "y" ]
```
**test_02:**
```go
var q = Node{Val: "q"}

// q

linkedListValues(q) // -> [ "q" ]
```
**test_03:**
```go
linkedListValues(nil) // -> [ ]
```