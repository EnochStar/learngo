package tree

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (res []int) {
	if root == nil {
		return res
	}
	for _, node := range root.Children {
		res = append(res, postorder(node)...)
	}
	res = append(res, root.Val)
	return res
}
