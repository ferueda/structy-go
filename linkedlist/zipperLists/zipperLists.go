package zipperLists

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func zipperLists[T any](node1, node2 *linkedlist.Node[T]) *linkedlist.Node[T] {
// 	tail, curr1, curr2 := node1, node1.Next, node2
// 	isEven := true

// 	for curr1 != nil && curr2 != nil {
// 		if isEven {
// 			tail.Next = curr2
// 			curr2 = curr2.Next
// 		} else {
// 			tail.Next = curr1
// 			curr1 = curr1.Next
// 		}

// 		tail = tail.Next
// 		isEven = !isEven
// 	}

// 	if curr1 != nil {
// 		tail.Next = curr1
// 	}
// 	if curr2 != nil {
// 		tail.Next = curr2
// 	}

// 	return node1
// }

// recursive

func zipperLists[T any](node1, node2 *linkedlist.Node[T]) *linkedlist.Node[T] {
	if node1 == nil && node2 == nil {
		return nil
	}
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	next := node1.Next
	node1.Next = node2
	node2.Next = zipperLists(next, node2.Next)

	return node1
}
