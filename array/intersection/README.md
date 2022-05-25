# Intersection

Write a function, intersection, that takes in two arrays, a,b, as arguments. The function should return a new array containing elements that are in both of the two arrays.

You may assume that each input array does not contain duplicate elements.

**test_00:**
```go
intersection([4,2,1,6], [3,6,9,2,10]) // -> [2,6]
```
**test_01:**
```go
intersection([2,4,6], [4,2]) // -> [2,4]
```
**test_02:**
```go
intersection([4,2,1], [1,2,4,6]) // -> [1,2,4]
```
**test_03:**
```go
intersection([0,1,2], [10,11]) // -> []
```
**test_04:**
```go
a, b := make([]int, 0, 50000), make([]int, 0, 50000)

for i := 0; i < 50000; i++ {
  a = append(a, i)
  b = append(b, i)
}

intersection(a, b) // -> [0,1,2,3,..., 49999]
```