package tree

import "fmt"

//定义struct类型
type TreeNode struct {
	value       int       //自身的值
	left, right *TreeNode //左右treeNode形结构的指针
}

//工厂函数,返回局部变量地址，外部可以调用
func CreateNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

//打印节点的value
func (node TreeNode) Print() {
	fmt.Println(node.value)
}

//设置值，一定要用指针传递，只有传递指针才可以改变结构内容
func (node *TreeNode) SetValue(val int) {
	if node == nil {
		//当节点为空时无法设置值，需要忽略处理
		fmt.Println("setting value to nil node, Ignored")
		return
	}
	node.value = val
}

//遍历结构体，需要传入指针类型
//先遍历左节点，在打印自己，再遍历右节点
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.left.Traverse()
	node.Print()
	node.right.Traverse()
}
