package tree

import "fmt"

//定义struct类型
type Node struct {
	Value       int   //自身的值
	Left, Right *Node //左右treeNode形结构的指针
}

//工厂函数,返回局部变量地址，外部可以调用
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//打印节点的value
func (node Node) Print() {
	fmt.Println(node.Value)
}

//设置值，一定要用指针传递，只有传递指针才可以改变结构内容
func (node *Node) SetValue(val int) {
	if node == nil {
		//当节点为空时无法设置值，需要忽略处理
		fmt.Println("setting value to nil node, Ignored")
		return
	}
	fmt.Println(node.Value)
	node.Value = val
}

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
