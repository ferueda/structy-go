package findValue

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func findValue[T comparable](node *linkedlist.Node[T], target T) bool {
// 	for n := node; n != nil; n = n.Next {
// 		if n.Val == target {
// 			return true
// 		}
// 	}
// 	return false
// }

// recursive

func findValue[T comparable](node *linkedlist.Node[T], target T) bool {
	if node == nil {
		return false
	}

	if node.Val == target {
		return true
	}

	return findValue(node.Next, target)
}
