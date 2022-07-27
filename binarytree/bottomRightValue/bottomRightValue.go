package bottomRightValue

import (
	"github.com/ferueda/structy-go/binarytree"
	"github.com/ferueda/structy-go/ds"
)

// breadth first

func bottomRightValue(root *binarytree.Node[string]) string {
	queue := ds.NewQueue[*binarytree.Node[string]]()
	queue.Enqueue(root)
	var current *binarytree.Node[string]

	for queue.Size() > 0 {
		current, _ = queue.Dequeue()
		if current.Left != nil {
			queue.Enqueue(current.Left)
		}
		if current.Right != nil {
			queue.Enqueue(current.Right)
		}
	}

	return current.Val
}
