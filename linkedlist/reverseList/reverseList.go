package reverseList

import "github.com/ferueda/structy-go/linkedlist"

// iterative

// func reverseList[T any](node *linkedlist.Node[T]) *linkedlist.Node[T] {
// 	curr := node
// 	var prev *linkedlist.Node[T]

// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}

// 	return prev
// }

// recursive

func reverseList[T any](node *linkedlist.Node[T]) *linkedlist.Node[T] {
	return _reverseList(node, nil)
}

func _reverseList[T any](curr, prev *linkedlist.Node[T]) *linkedlist.Node[T] {
	if curr == nil {
		return prev
	}

	next := curr.Next
	curr.Next = prev

	return _reverseList(next, curr)
}
