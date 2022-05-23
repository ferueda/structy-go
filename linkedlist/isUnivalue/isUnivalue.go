package isUnivalue

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func isUnivalue[T comparable](node *linkedlist.Node[T]) bool {
// 	for n := node; n != nil; n = n.Next {
// 		if n.Val != node.Val {
// 			return false
// 		}
// 	}

// 	return true
// }

// recursive

func isUnivalue[T comparable](node *linkedlist.Node[T]) bool {
	return _isUniValue(node, node.Val)
}

func _isUniValue[T comparable](node *linkedlist.Node[T], val T) bool {
	if node == nil {
		return true
	}
	if node.Val != val {
		return false
	}

	return _isUniValue(node.Next, val)
}
