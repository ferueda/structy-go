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

// 	stack := ds.NewStack[*binarytree.Node[T]]()
// 	stack.Push(root)

// 	for stack.Size() > 0 {
// 		node, _ := stack.Pop()
// 		values = append(values, node.Val)

// 		if node.Right != nil {
// 			stack.Push(node.Right)
// 		}
// 		if node.Left != nil {
// 			stack.Push(node.Left)
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
