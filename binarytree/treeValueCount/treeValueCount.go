package treeValueCount

import "github.com/ferueda/structy-go/binarytree"

// breadth first search

// func treeValueCount(root *binarytree.Node[int], target int) int {
// 	if root == nil {
// 		return 0
// 	}

// 	count := 0
// 	queue := ds.NewQueue[*binarytree.Node[int]]()
// 	queue.Enqueue(root)

// 	for queue.Size() > 0 {
// 		node, _ := queue.Dequeue()
// 		if node.Val == target {
// 			count += 1
// 		}
// 		if node.Left != nil {
// 			queue.Enqueue(node.Left)
// 		}
// 		if node.Right != nil {
// 			queue.Enqueue(node.Right)
// 		}
// 	}

// 	return count
// }

// depth first (recursive)

func treeValueCount(root *binarytree.Node[int], target int) int {
	if root == nil {
		return 0
	}

	match := 0
	if root.Val == target {
		match = 1
	}

	return match + treeValueCount(root.Left, target) + treeValueCount(root.Right, target)
}
