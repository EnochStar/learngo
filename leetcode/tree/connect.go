package tree

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) (res *Node) {
	if root == nil {
		return root
	}
	var connectTwo func(node1 *Node, node2 *Node)
	connectTwo = func(node1 *Node, node2 *Node) {
		if node1 == nil || node2 == nil {
			return
		}
		node1.Next = node2
		connectTwo(node1.Left, node1.Right)
		connectTwo(node2.Left, node2.Right)
		connectTwo(node1.Right, node2.Left)
	}
	connectTwo(root.Left, root.Right)
	return root
}
