package recursion

import "math"

var pre = math.MinInt64

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBST(root.Left) {
		return false
	}
	if root.Val <= pre {
		return false
	}
	return isValidBST(root.Right)
}
