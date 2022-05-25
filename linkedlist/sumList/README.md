# Sum list

Write a function, sumList, that takes in the head of a linked list containing numbers as an argument. The function should return the total sum of all values in the linked list.

**test_00:**
```go
var a = Node{Val: 2}
var b = Node{Val: 8}
var c = Node{Val: 3}
var d = Node{Val: -1}
var e = Node{Val: 7}

a.next = &b
b.next = &c
c.next = &d
d.next = &e

// 2 -> 8 -> 3 -> -1 -> 7

sumList(a) // -> 19
```
**test_01:**
```go
var x = Node{Val: 38}
var y = Node{Val: 4}

x.next = &y

// 38 -> 4

sumList(x) // -> 42
```
**test_02:**
```go
var q = Node{Val: 100}

// 100

sumList(q) // -> 100
```
**test_03:**
```go
sumList(nil) // -> 0
```