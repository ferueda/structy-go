package maxpathsum

import (
	"math"

	"github.com/ferueda/structy-go/binarytree"
	"github.com/ferueda/structy-go/utils"
)

// depth first (recursive)

func maxPathSum(root *binarytree.Node[int]) int {
	if root == nil {
		return math.MinInt
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	return root.Val + utils.Max(maxPathSum(root.Left), maxPathSum(root.Right))
}
