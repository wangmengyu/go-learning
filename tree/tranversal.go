package tree

import "fmt"

//遍历结构体，需要传入指针类型
//先遍历左节点，在打印自己，再遍历右节点
func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
	fmt.Println()

}

/**
  闭包函数，遍历节点
  参数：带一个Node节点指针的函数。
  意义： f函数定义了要对节点操作的具体内容
*/
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)

}
