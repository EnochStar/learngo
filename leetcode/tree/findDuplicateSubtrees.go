package tree

import "strconv"

func findDuplicateSubtrees(root *TreeNode) (res []*TreeNode) {
	var traverse func(root *TreeNode) string
	stringMap := make(map[string]int)
	traverse = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		left := traverse(root.Left)
		right := traverse(root.Right)
		var s string
		s += left + "," + right + "," + strconv.Itoa(root.Val)
		if v, ok := stringMap[s]; !ok {
			stringMap[s] = 1
		} else {
			if v == 1 {
				res = append(res, root)
			}
			stringMap[s]++
		}
		return s
	}
	traverse(root)
	return res
}
