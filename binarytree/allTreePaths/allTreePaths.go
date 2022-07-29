package allTreePaths

import (
	"github.com/ferueda/structy-go/binarytree"
)

func allTreePaths(root *binarytree.Node[string]) [][]string {
	if root == nil {
		return [][]string{}
	}
	if root.Left == nil && root.Right == nil {
		return [][]string{{root.Val}}
	}

	var paths [][]string

	leftSubPath := allTreePaths(root.Left)
	for _, sp := range leftSubPath {
		arr := []string{root.Val}
		for _, p := range sp {
			arr = append(arr, p)
		}
		paths = append(paths, arr)
	}

	rightSubPath := allTreePaths(root.Right)
	for _, sp := range rightSubPath {
		arr := []string{root.Val}
		for _, p := range sp {
			arr = append(arr, p)
		}
		paths = append(paths, arr)
	}

	return paths
}
