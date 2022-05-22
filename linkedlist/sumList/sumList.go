package sumList

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func sumList(node *linkedlist.Node[int]) int {
// 	s := 0

// 	for n := node; n != nil; n = n.Next {
// 		s += n.Val
// 	}

// 	return s
// }

// recursive

func sumList(node *linkedlist.Node[int]) int {
	return _sumList(node, 0)
}

func _sumList(node *linkedlist.Node[int], sum int) int {
	if node == nil {
		return sum
	}

	return _sumList(node.Next, sum+node.Val)
}
