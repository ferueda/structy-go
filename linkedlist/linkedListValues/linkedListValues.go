package linkedListValues

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func linkedListValues[T any](node *linkedlist.Node[T]) []T {
// 	r := []T{}

// 	for n := node; n != nil; n = n.Next {
// 		r = append(r, n.Val)
// 	}

// 	return r
// }

// recursive

func linkedListValues[T any](node *linkedlist.Node[T]) []T {
	values := []T{}
	return _linkedListValues(node, values)
}

func _linkedListValues[T any](node *linkedlist.Node[T], values []T) []T {
	if node == nil {
		return values
	}
	values = append(values, node.Val)
	return _linkedListValues(node.Next, values)
}
