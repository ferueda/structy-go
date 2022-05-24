package insertNode

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func insertNode[T any](node *linkedlist.Node[T], val T, i int) *linkedlist.Node[T] {
// 	if i == 0 {
// 		return &linkedlist.Node[T]{Val: val, Next: node}
// 	}

// 	c := 0
// 	for n := node; n != nil; n = n.Next {
// 		if c == i-1 {
// 			next := n.Next
// 			n.Next = &linkedlist.Node[T]{Val: val, Next: next}
// 			break
// 		}

// 		c += 1
// 	}

// 	return node
// }

// recursive

func insertNode[T any](node *linkedlist.Node[T], val T, i int) *linkedlist.Node[T] {
	if i == 0 {
		return &linkedlist.Node[T]{Val: val, Next: node}
	}

	return _insertNode(node, val, i, 0)
}

func _insertNode[T any](node *linkedlist.Node[T], val T, i, c int) *linkedlist.Node[T] {
	if c == i-1 {
		next := node.Next
		node.Next = &linkedlist.Node[T]{Val: val, Next: next}
		return node
	}

	_insertNode(node.Next, val, i, c+1)
	return node
}
