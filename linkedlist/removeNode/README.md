# Remove node

Write a function, removeNode, that takes in the head of a linked list and a target value as arguments. The function should delete the node containing the target value from the linked list and return the head of the resulting linked list. If the target appears multiple times in the linked list, only remove the first instance of the target in the list.

Do this in-place.

You may assume that the input list is non-empty.

You may assume that the input list contains the target.

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

removeNode(a, "c") // a -> b -> d -> e -> f
```
**test_01:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}

a.next = &b
b.next = &c

// a -> b -> c

removeNode(a, "c") // a -> b
```
**test_02:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}

a.next = &b
b.next = &c

// a -> b -> c

removeNode(a, "a") // b -> c
```
**test_03:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}
var d = Node{Val: "b"}

a.next = &b
b.next = &c
c.next = &d

// a -> b -> c -> b

removeNode(a, "b") // a -> c -> b
```
**test_04:**
```go
var a = Node{Val: "a"}

// a

removeNode(a, "a") // nil
```