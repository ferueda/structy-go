# Insert node

Write a function, insertNode, that takes in the head of a linked list, a value, and an index. The function should insert a new node with the value into the list at the specified index. Consider the head of the linked list as index 0. The function should return the head of the resulting linked list.

Do this in-place.

You may assume that the input list is non-empty and the index is not greater than the length of the input list.

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

insertNode(a, "x", 2) // a -> b -> x -> c -> d
```
**test_01:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "d"}

a.next = &b
b.next = &c
c.next = &d

// a -> b -> c -> d 

insertNode(a, "v", 3) // a -> b -> c -> v -> d
```
**test_02:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "d"}

a.next = &b
b.next = &c
c.next = &d

// a -> b -> c -> d 

insertNode(a, "m", 4) // a -> b -> c -> d -> m
```
**test_03:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}

a.next = &b

// a -> b

insertNode(a, "z", 0) // z -> a -> b 
```