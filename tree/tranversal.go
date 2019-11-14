package tree

//遍历结构体，需要传入指针类型
//先遍历左节点，在打印自己，再遍历右节点
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
