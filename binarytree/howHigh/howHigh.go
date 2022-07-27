package howHigh

import "github.com/ferueda/structy-go/binarytree"

// depth first (recursive)

func howHigh(root *binarytree.Node[string]) int {
	if root == nil {
		return -1
	}

	leftHight := howHigh(root.Left)
	rightHight := howHigh(root.Right)

	if leftHight > rightHight {
		return 1 + leftHight
	} else {
		return 1 + rightHight
	}
}
