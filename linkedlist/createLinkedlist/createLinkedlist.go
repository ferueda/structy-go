package createLinkedlist

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func createLinkedlist[T any](values []T) *linkedlist.Node[T] {
// 	dummyHead := &linkedlist.Node[T]{}
// 	tail := dummyHead

// 	for _, v := range values {
// 		tail.Next = &linkedlist.Node[T]{Val: v}
// 		tail = tail.Next
// 	}

// 	return dummyHead.Next
// }

// recursive

func createLinkedlist[T any](values []T) *linkedlist.Node[T] {
	return _createLinkedlist(values, 0)
}

func _createLinkedlist[T any](values []T, i int) *linkedlist.Node[T] {
	if i == len(values) {
		return nil
	}
	return &linkedlist.Node[T]{Val: values[i], Next: _createLinkedlist(values, i+1)}
}
