# Find value

Write a function, findValue, that takes in the head of a linked list and a target value. The function should return a boolean indicating whether or not the linked list contains the target.

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

findValue(a, "c") // -> true
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

findValue(a, "d") // -> true
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

findValue(a, "q") // -> false
```
**test_03:**
```go
var a = Node{Val: "jason"}
var b = Node{Val: "leneli"}

a.next = &b

// jason -> leneli

findValue(a, "jason") // -> true
```
**test_04:**
```go
var a = Node{Val: 42}

// 42

findValue(a, 42) // -> true
```
**test_05:**
```go
var a = Node{Val: 42}

// 42

findValue(a, 100) // -> false
```