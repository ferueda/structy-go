package treeLevels

import (
	"github.com/ferueda/structy-go/binarytree"
	"github.com/ferueda/structy-go/ds"
)

type queueItem[T comparable] struct {
	node  *binarytree.Node[T]
	level int
}

// breadth first (iterative)

func treeLevels[T comparable](root *binarytree.Node[T]) [][]T {
	if root == nil {
		return [][]T{}
	}

	var levels [][]T

	queue := ds.NewQueue[queueItem[T]]()
	queue.Enqueue(queueItem[T]{node: root, level: 0})

	for queue.Size() > 0 {
		item, _ := queue.Dequeue()
		if item.level == len(levels) {
			levels = append(levels, []T{item.node.Val})
		} else {
			levels[item.level] = append(levels[item.level], item.node.Val)
		}

		if item.node.Left != nil {
			queue.Enqueue(queueItem[T]{node: item.node.Left, level: item.level + 1})
		}
		if item.node.Right != nil {
			queue.Enqueue(queueItem[T]{node: item.node.Right, level: item.level + 1})
		}
	}

	return levels
}
