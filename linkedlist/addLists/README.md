# Add lists

Write a function, addLists, that takes in the head of two linked lists, each representing a number. The nodes of the linked lists contain digits as values. The nodes in the input lists are reversed this means that the least significant digit of the number is the head. The function should return the head of a new linked listed representing the sum of the input lists. The output list should have its digits reversed as well.

```
Say we wanted to compute 621 + 354 normally. The sum is 975:

   621
 + 354
 -----
   975

Then, the reversed linked list format of this problem would appear as:

    1 -> 2 -> 6
 +  4 -> 5 -> 3
 --------------
    5 -> 7 -> 9
```

**test_00:**
```go
//   621
// + 354
// -----
//   975

var a1 = Node{Val: 1}
var a2 = Node{Val: 2}
var a3 = Node{Val: 6}

a1.next = &a2
a2.next = &a3

// 1 -> 2 -> 6

var b1 = Node{Val: 4}
var b2 = Node{Val: 5}
var b3 = Node{Val: 3}

b1.next = &b2
b2.next = &b3

// 4 -> 5 -> 3

addLists(a1, b1) // 5 -> 7 -> 9
```
**test_01:**
```go
//  7541
// +  32
// -----
//  7573

var a1 = Node{Val: 1}
var a2 = Node{Val: 4}
var a3 = Node{Val: 5}
var a4 = Node{Val: 7}

a1.next = &a2
a2.next = &a3
a3.next = &a4

// 1 -> 4 -> 5 -> 7

var b1 = Node{Val: 2}
var b2 = Node{Val: 3}

b1.next = &b2

// 2 -> 3 

addLists(a1, b1) // 3 -> 7 -> 5 -> 7
```
**test_02:**
```go
//   39
// + 47
// ----
//   86

var a1 = Node{Val: 9}
var a2 = Node{Val: 3}

a1.next = &a2

// 9 -> 3

var b1 = Node{Val: 7}
var b2 = Node{Val: 4}

b1.next = &b2

// 7 -> 4

addLists(a1, b1) // 6 -> 8
```
**test_03:**
```go
//   89
// + 47
// ----
//  136

var a1 = Node{Val: 9}
var a2 = Node{Val: 8}

a1.next = &a2

// 9 -> 8 

var b1 = Node{Val: 7}
var b2 = Node{Val: 4}

b1.next = &b2

// 7 -> 4

addLists(a1, b1) // 6 -> 3 -> 1
```
**test_04:**
```go
//   999
//  +  6
//  ----
//  1005

var a1 = Node{Val: 9}
var a2 = Node{Val: 9}
var a3 = Node{Val: 9}

a1.next = &a2
a2.next = &a3

// 9 -> 9 -> 9

var b1 = Node{Val: 6}

// 6

addLists(a1, b1) // 5 -> 0 -> 0 -> 1
```