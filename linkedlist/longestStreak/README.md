# Longest streak

Write a function, longestStreak, that takes in the head of a linked list as an argument. The function should return the length of the longest consecutive streak of the same value within the list.

**test_00:**
```go
var a = Node{Val: 5}
var b = Node{Val: 5}
var c = Node{Val: 7}
var d = Node{Val: 7}
var e = Node{Val: 7}
var f = Node{Val: 6}

a.next = &b
b.next = &c
c.next = &d
d.next = &e
e.next = &f

// 5 -> 5 -> 7 -> 7 -> 7 -> 6

longestStreak(a) // 3
```
**test_01:**
```go
var a = Node{Val: 3}
var b = Node{Val: 3}
var c = Node{Val: 3}
var d = Node{Val: 3}
var e = Node{Val: 9}
var f = Node{Val: 9}

a.next = &b
b.next = &c
c.next = &d
d.next = &e
e.next = &f

// 3 -> 3 -> 3 -> 3 -> 9 -> 9

longestStreak(a) // 4
```
**test_02:**
```go
var a = Node{Val: 9}
var b = Node{Val: 9}
var c = Node{Val: 1}
var d = Node{Val: 9}
var e = Node{Val: 9}
var f = Node{Val: 9}

a.next = &b
b.next = &c
c.next = &d
d.next = &e
e.next = &f

// 9 -> 9 -> 1 -> 9 -> 9 -> 9

longestStreak(a) // 3
```
**test_03:**
```go
var a = Node{Val: 5}
var b = Node{Val: 5}

a.next = &b

// 5 -> 5

longestStreak(a) // 2
```
**test_04:**
```go
var a = Node{Val: 4}

// 5

longestStreak(a) // 1
```
**test_05:**
```go
longestStreak(nil) // 0
```