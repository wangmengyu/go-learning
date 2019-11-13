package main

import ""

func main() {
	//定义treeMode类型值，全部用点访问下去元素，
	var root tree.TreeNode
	root.value = 3
	root.left = &treeNode{}
	root.right = &treeNode{5, &treeNode{}, &treeNode{}}
	root.left.right = createNode(2)
	//打印值
	root.print()
	//修改值
	root.right.left.setValue(8) //定义时候是指针，但是实际调用不用做特殊标注，会传地址过去的
	root.right.left.print()

	//给空指针设置值
	var pRoot *treeNode
	pRoot.setValue(200)
	pRoot = &root
	pRoot.setValue(300)
	pRoot.print()

	root.traverse()

}
