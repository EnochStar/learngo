package main

import (
	"fmt"
	"learngo/video/structs"
)

//type myTreeNode struct {
//	node *structs.TreeNode
//}

//内嵌 Embedding
type myTreeNode struct {
	// 相当于将内嵌结构的方法和属性作为自己的
	*structs.TreeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.TreeNode == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

//值接收者：指针接收者
// 1、改变内容必须使用指针接收者
// 2、结构过大也考虑使用指针接收者
// 3、一致性：如果有指针接收者，最好都是指针接收者

// 根据名字设定可见性
// 首字母大写表示public 小写表示private
// 可见性是相对于包而言的
// 与java不同 包名不用与路径名相同
// 一个目录下的同级文件归属一个包
func main() {

	root := myTreeNode{&structs.TreeNode{Value: 3}}
	root.Left = &structs.TreeNode{}
	root.Right = &structs.TreeNode{5, nil, nil}
	root.Right.Left = new(structs.TreeNode)
	root.Left.Right = structs.CreateNode(2)
	nodes := []structs.TreeNode{
		{Value: 3},
		{},
		{6, nil, root.TreeNode},
	}
	fmt.Println(nodes)

	root.SetByNum(2)
	fmt.Println("root Value by num", root.Value)
	root.SetByAddress(5)
	fmt.Println("root Value by address", root.Value)
	var zero *structs.TreeNode
	zero.SetByAddress(10)
	zero = root.TreeNode
	zero.SetByAddress(10)
	zero.Print()
	fmt.Println("遍历树")
	root.Traverse()

	node := root
	node.postOrder()
}
