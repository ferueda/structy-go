package treeMinValue

import (
	"math"

	"github.com/ferueda/structy-go/binarytree"
	"github.com/ferueda/structy-go/ds"
	// "github.com/ferueda/structy-go/utils"
)

// depth first

// func treeMinValue(root *binarytree.Node[int]) int {
// 	if root == nil {
// 		return math.MaxInt
// 	}
// 	return utils.Min(root.Val, utils.Min(treeMinValue(root.Left), treeMinValue(root.Right)))
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
