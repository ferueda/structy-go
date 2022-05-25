# Create linkedlist

Write a function, createLinkedlist, that takes in an array of values as an argument. The function should create a linked list containing each element of the array as the values of the nodes. The function should return the head of the linked list.

**test_00:**
```go
createLinkedlist([]string{"h", "e", "y"}) // h -> e -> y
```
**test_01:**
```go
createLinkedlist([]string{"1", "7", "1", "8"}) // 1 -> 7 -> 1 -> 8
```
**test_02:**
```go
createLinkedlist([]string{"a"}) // a
```
**test_03:**
```go
createLinkedlist([]string{}) // nil
```