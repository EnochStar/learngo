package recursion

import "math"

func maxDepth(root *TreeNode) (res int) {
	if root == nil {
		return 0
	}
	var getDepth func(root *TreeNode, res int) int
	getDepth = func(root *TreeNode, res int) int {
		if root == nil {
			return res
		}
		return int(math.Max(float64(getDepth(root.Left, res+1)), float64(getDepth(root.Right, res+1))))
	}
	return getDepth(root, 0)
}
