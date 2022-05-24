package addLists

import (
	"github.com/ferueda/structy-go/linkedlist"
)

// iterative

// func addLists(node1, node2 *linkedlist.Node[int]) *linkedlist.Node[int] {
// 	dummyHead := &linkedlist.Node[int]{}
// 	tail := dummyHead
// 	carry, curr1, curr2 := 0, node1, node2

// 	for curr1 != nil || curr2 != nil || carry != 0 {
// 		var val1, val2 int
// 		if curr1 != nil {
// 			val1 = curr1.Val
// 		}
// 		if curr2 != nil {
// 			val2 = curr2.Val
// 		}

// 		sum := val1 + val2 + carry
// 		if sum > 9 {
// 			carry = 1
// 		} else {
// 			carry = 0
// 		}

// 		tail.Next = &linkedlist.Node[int]{Val: sum % 10}
// 		tail = tail.Next

// 		if curr1 != nil {
// 			curr1 = curr1.Next
// 		}
// 		if curr2 != nil {
// 			curr2 = curr2.Next
// 		}
// 	}

// 	return dummyHead.Next
// }

// recursive

func addLists(node1, node2 *linkedlist.Node[int]) *linkedlist.Node[int] {
	return _addLists(node1, node2, 0)
}

func _addLists(node1, node2 *linkedlist.Node[int], carry int) *linkedlist.Node[int] {
	if node1 == nil && node2 == nil && carry == 0 {
		return nil
	}

	var val1, val2 int
	if node1 != nil {
		val1 = node1.Val
	}
	if node2 != nil {
		val2 = node2.Val
	}

	sum := val1 + val2 + carry
	if sum > 9 {
		carry = 1
	} else {
		carry = 0
	}

	result := &linkedlist.Node[int]{Val: sum % 10}

	var next1, next2 *linkedlist.Node[int]
	if node1 != nil {
		next1 = node1.Next
	}
	if node2 != nil {
		next2 = node2.Next
	}

	result.Next = _addLists(next1, next2, carry)
	return result
}
