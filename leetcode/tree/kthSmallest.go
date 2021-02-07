package tree

func kthSmallest(root *TreeNode, k int) int {
	var midTraverse func(root *TreeNode)
	var res []int
	midTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		midTraverse(root.Left)
		res = append(res, root.Val)
		midTraverse(root.Right)
	}
	midTraverse(root)
	return res[k-1]
}
