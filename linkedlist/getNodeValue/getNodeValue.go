package getNodeValue

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func getNodeValue[T any](node *linkedlist.Node[T], index int) T {
// 	var r T
// 	for n, c := node, 0; n != nil; n = n.Next {
// 		if c == index {
// 			r = n.Val
// 			break
// 		}
// 		c += 1
// 	}
// 	return r
// }

// recursive

func getNodeValue[T any](node *linkedlist.Node[T], index int) T {
	if node == nil {
		var r T
		return r
	}

	if index == 0 {
		return node.Val
	}

	return getNodeValue(node.Next, index-1)
}
