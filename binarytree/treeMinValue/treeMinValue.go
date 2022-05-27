package treeMinValue

import (
	"math"

	"github.com/ferueda/structy-go/binarytree"
	"github.com/ferueda/structy-go/ds"
)

// depth first

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func treeMinValue(root *binarytree.Node[int]) int {
// 	if root == nil {
// 		return math.MaxInt
// 	}
// 	return min(root.Val, min(treeMinValue(root.Left), treeMinValue(root.Right)))
// }

// breadth first

func treeMinValue(root *binarytree.Node[int]) int {
	queue := ds.NewQueue[*binarytree.Node[int]]()
	queue.Enqueue(root)
	min := math.MaxInt

	for queue.Size() > 0 {
		node, _ := queue.Dequeue()

		if node.Val < min {
			min = node.Val
		}

		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}

	return min
}
