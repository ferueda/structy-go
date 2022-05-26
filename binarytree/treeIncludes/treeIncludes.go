package treeIncludes

import (
	"github.com/ferueda/structy-go/binarytree"
	"github.com/ferueda/structy-go/ds"
)

// depth first

// func treeIncludes[T comparable](root *binarytree.Node[T], target T) bool {
// 	if root == nil {
// 		return false
// 	}

// 	if root.Val == target {
// 		return true
// 	}

// 	return treeIncludes(root.Left, target) || treeIncludes(root.Right, target)
// }

// breadth first

func treeIncludes[T comparable](root *binarytree.Node[T], target T) bool {
	if root == nil {
		return false
	}

	queue := ds.NewQueue[*binarytree.Node[T]]()
	queue.Enqueue(root)

	for queue.Size() > 0 {
		node, _ := queue.Dequeue()

		if node.Val == target {
			return true
		}

		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}

	return false
}
