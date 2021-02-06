package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	length := len(nums)

	maxValue, minIndex := 0, -1
	for i := 0; i < length; i++ {
		if maxValue <= nums[i] {
			maxValue = nums[i]
			minIndex = i
		}
	}
	if minIndex == -1 {
		return nil
	}
	tree := TreeNode{
		Val: maxValue,
	}
	tree.Left = constructMaximumBinaryTree(nums[:minIndex])
	tree.Right = constructMaximumBinaryTree(nums[minIndex+1:])
	return &tree
}
