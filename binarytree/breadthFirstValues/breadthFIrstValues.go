package breadthFirstValues

import (
	"github.com/ferueda/structy-go/binarytree"
	"github.com/ferueda/structy-go/ds"
)

func breadthFirstValues[T any](root *binarytree.Node[T]) []T {
	values := []T{}

	if root == nil {
		return values
	}

	queue := ds.NewQueue[*binarytree.Node[T]]()
	queue.Enqueue(root)

	for queue.Size() > 0 {
		node, _ := queue.Dequeue()
		values = append(values, node.Val)

		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
	}

	return values
}
