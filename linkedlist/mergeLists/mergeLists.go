package mergeLists

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func mergeLists(node1, node2 *linkedlist.Node[int]) *linkedlist.Node[int] {
// 	dummyHead := &linkedlist.Node[int]{}
// 	tail, curr1, curr2 := dummyHead, node1, node2

// 	for curr1 != nil && curr2 != nil {
// 		if curr1.Val < curr2.Val {
// 			tail.Next = curr1
// 			curr1 = curr1.Next
// 		} else {
// 			tail.Next = curr2
// 			curr2 = curr2.Next
// 		}

// 		tail = tail.Next
// 	}

// 	if curr1 != nil {
// 		tail.Next = curr1
// 	}
// 	if curr2 != nil {
// 		tail.Next = curr2
// 	}

// 	return dummyHead.Next
// }

// recursive

func mergeLists(node1, node2 *linkedlist.Node[int]) *linkedlist.Node[int] {
	if node1 == nil && node2 == nil {
		return nil
	}
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	if node1.Val > node2.Val {
		next := node2.Next
		node2.Next = mergeLists(node1, next)
		return node2
	} else {
		next := node1.Next
		node1.Next = mergeLists(node2, next)
		return node1
	}
}
