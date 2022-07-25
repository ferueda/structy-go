package treePathFinder

import (
	"github.com/ferueda/structy-go/binarytree"
)

func pathFinder(root *binarytree.Node[string], target string) []string {
	result := _pathFinder(root, target)
	if result == nil {
		return nil
	}

	// reverse result slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func _pathFinder(root *binarytree.Node[string], target string) []string {
	if root == nil {
		return nil
	}
	if root.Val == target {
		return []string{root.Val}
	}

	leftPath := _pathFinder(root.Left, target)
	if leftPath != nil {
		return append(leftPath, root.Val)
	}

	rightPath := _pathFinder(root.Right, target)
	if rightPath != nil {
		return append(rightPath, root.Val)
	}

	return nil
}
