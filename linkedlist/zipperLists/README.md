# Zipper lists

Write a function, zipperLists, that takes in the head of two linked lists as arguments. The function should zipper the two lists together into single linked list by alternating nodes. If one of the linked lists is longer than the other, the resulting list should terminate with the remaining nodes. The function should return the head of the zippered linked list.

Do this in-place, by mutating the original Nodes.

You may assume that both input lists are non-empty.

**test_00:**
```go
var a = Node{Val: "a"}
var b = Node{Val: "b"}
var c = Node{Val: "c"}

a.next = &b
b.next = &c

// a -> b -> c

var x = Node{Val: "x"}
var y = Node{Val: "y"}
var z = Node{Val: "z"}


x.next = &y
y.next = &z

// x -> y -> z

zipperLists(a, x) // a -> x -> b -> y -> c -> z
```
**test_01:**
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

var x = Node{Val: "x"}
var y = Node{Val: "y"}
var z = Node{Val: "z"}

x.next = &y
y.next = &z

// x -> y -> z

zipperLists(a, x) // a -> x -> b -> y -> c -> z -> d -> e -> f
```
**test_02:**
```go
var s = Node{Val: "s"}
var t = Node{Val: "t"}

s.next = &t

var one = Node{Val: "1"}
var two = Node{Val: "2"}
var three = Node{Val: "3"}
var four = Node{Val: "4"}

one.next = &two
two.next = &three
three.next = &four

// 1 -> 2 -> 3 -> 4

zipperLists(s, one) // s -> 1 -> t -> 2 -> 3 -> 4
```
**test_03:**
```go
var w = Node{Val: "w"}

// w

var one = Node{Val: "1"}
var two = Node{Val: "2"}
var three = Node{Val: "3"}

one.next = &two
two.next = &three

// 1 -> 2 -> 3 

zipperLists(w, one) // w -> 1 -> 2 -> 3
```
**test_04:**
```go
var one = Node{Val: "1"}
var two = Node{Val: "2"}
var three = Node{Val: "3"}

one.next = &two
two.next = &three

// 1 -> 2 -> 3 

var w = Node{Val: "w"}

// w

zipperLists(one, w) // 1 -> w -> 2 -> 3
```