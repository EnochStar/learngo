package recursion

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	var ans int
	if left == 0 || right == 0 {
		ans = left + right + 1
	} else {
		ans = int(math.Min(float64(left), float64(right)) + 1)
	}
	return ans
}
