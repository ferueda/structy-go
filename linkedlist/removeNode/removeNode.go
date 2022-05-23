package removeNode

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func removeNode[T comparable](node *linkedlist.Node[T], val T) *linkedlist.Node[T] {
// 	if node != nil && node.Val == val {
// 		return node.Next
// 	}

// 	var prev *linkedlist.Node[T]
// 	for n := node; n != nil; n = n.Next {
// 		if n.Val == val {
// 			prev.Next = n.Next
// 			break
// 		}

// 		prev = n
// 	}

// 	return node
// }

// recursive

func removeNode[T comparable](node *linkedlist.Node[T], val T) *linkedlist.Node[T] {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node.Next
	}

	node.Next = removeNode(node.Next, val)

	return node
}
