package leafList

import (
	"github.com/ferueda/structy-go/binarytree"
)

// depth first (recursive)

func leafList[T comparable](root *binarytree.Node[T]) []T {
	leaves := []T{}
	fillLeaves(root, &leaves)
	return leaves
}

func fillLeaves[T comparable](root *binarytree.Node[T], leaves *[]T) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
	}

	fillLeaves(root.Left, leaves)
	fillLeaves(root.Right, leaves)
}
