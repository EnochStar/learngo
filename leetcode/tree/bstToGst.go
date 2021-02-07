package tree

func bstToGst(root *TreeNode) *TreeNode {
	var traverse func(root *TreeNode)
	sum := 0
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Right)
		sum += root.Val
		root.Val = sum
		traverse(root.Left)
	}
	traverse(root)
	return root
}
