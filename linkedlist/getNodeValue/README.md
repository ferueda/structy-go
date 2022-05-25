# Get node value

Write a function, getNodeValue, that takes in the head of a linked list and an index. The function should return the value of the linked list at the specified index.

If there is no node at the given index, then return an empty string.

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

getNodeValue(a, 2) // -> "c"
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

getNodeValue(a, 3) // -> "d"
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

getNodeValue(a, 7) // -> ""
```
**test_03:**
```go
var node1 = Node{Val: "banana"}
var node2 = Node{Val: "mango"}

node1.next = &node2

// jason -> leneli

getNodeValue(node1, 0) // -> "banana"
```
**test_04:**
```go
var node1 = Node{Val: "banana"}
var node2 = Node{Val: "mango"}

node1.next = &node2

// jason -> leneli

getNodeValue(node1, 1) // -> "mango"
```
