package main

import "go-learning/tree"

func main() {
	//定义treeMode类型值，全部用点访问下去元素，
	var root tree.TreeNode
	root.Value = 3
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, &tree.TreeNode{}, &tree.TreeNode{}}
	root.Left.Right = tree.CreateNode(2)
	//打印值
	root.Print()
	//修改值
	root.Right.Left.SetValue(8) //定义时候是指针，但是实际调用不用做特殊标注，会传地址过去的
	root.Right.Left.Print()

	//给空指针设置值
	var pRoot *tree.TreeNode
	pRoot.SetValue(200)
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()

	root.Traverse()

}
