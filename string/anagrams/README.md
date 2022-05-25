# Anagrams

Write a function, anagrams, that takes in two strings as arguments. The function should return a boolean indicating whether or not the strings are anagrams. Anagrams are strings that contain the same characters, but in any order.

**test_00:**
```go
anagrams("restful", "fluster") // -> true
```
**test_01:**
```go
anagrams("cats", "tocs") // -> false
```
**test_02:**
```go
anagrams("monkeyswrite", "newyorktimes") // -> true
```
**test_03:**
```go
anagrams("paper", "reapa") // -> false
```
**test_04:**
```go
anagrams("elbow", "below") // -> true
```
**test_05:**
```go
anagrams("tax", "taxi") // -> false
```
**test_06:**
```go
anagrams("taxi", "tax") // -> false
```
**test_07:**
```go
anagrams("night", "thing") // -> true
```
**test_08:**
```go
anagrams("abbc", "aabc") // -> false
```