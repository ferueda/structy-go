package depthFirstValues

import (
	// "github.com/ferueda/structy-go/ds"
	"github.com/ferueda/structy-go/binarytree"
)

// iterative

// func depthFirstValues[T any](root *binarytree.Node[T]) []T {
// 	values := []T{}
// 	if root == nil {
// 		return values
// 	}

// 	queue := ds.Stack[*binarytree.Node[T]]{}
// 	queue.Push(root)

// 	for queue.Size() > 0 {
// 		curr, _ := queue.Pop()
// 		values = append(values, curr.Val)

// 		if curr.Right != nil {
// 			queue.Push(curr.Right)
// 		}
// 		if curr.Left != nil {
// 			queue.Push(curr.Left)
// 		}
// 	}

// 	return values
// }

// recursive

func depthFirstValues[T any](root *binarytree.Node[T]) []T {
	if root == nil {
		return []T{}
	}

	left := depthFirstValues(root.Left)
	right := depthFirstValues(root.Right)

	values := make([]T, 0, len(left)+len(right)+1)
	values = append(values, root.Val)
	values = append(values, left...)
	values = append(values, right...)

	return values
}
