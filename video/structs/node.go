package structs

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// 工厂函数，返回了局部变量的地址
// 那么结构是在堆还是在栈？     不需要知道?????????视情况而定
// nil指针也可以调用方法
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

//传值
func (node TreeNode) Print() {
	fmt.Println(node.Value)
}
func (node TreeNode) SetByNum(value int) {
	node.Value = value
}
func (node *TreeNode) SetByAddress(value int) {
	if node == nil {
		fmt.Println("Setting value To nil")
		return
	}
	node.Value = value
}

//调用函数的对象可以为nil
func (node *TreeNode) Traverse() {
	node.TraverseFunc(func(n *TreeNode) {
		n.Print()
	})
	fmt.Println()
}

func (node *TreeNode) TraverseFunc(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
