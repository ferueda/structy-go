package treesum

import (
	"github.com/ferueda/structy-go/binarytree"
	"github.com/ferueda/structy-go/ds"
)

// depth first

// func treeSum(root *binarytree.Node[int]) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return root.Val + treeSum(root.Left) + treeSum(root.Right)
// }

// breadth first

func treeSum(root *binarytree.Node[int]) int {
	var sum int

	if root == nil {
		return sum
	}

	queue := ds.NewQueue[*binarytree.Node[int]]()
	queue.Enqueue(root)

	for queue.Size() > 0 {
		node, _ := queue.Dequeue()
		sum += node.Val

		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}

	return sum
}
