package main

import (
	"fmt"
	"github.com/wangmengyu/go-learning/tree"
)

/**
对于tree.Node的扩展类 - 组合式的
*/
type myTreeNode struct {
	node *tree.Node
}

/**
扩展了之前遍历方法
左，右，中 顺序遍历树, 后续遍历
*/
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}

func main() {
	//定义treeMode类型值，全部用点访问下去元素，
	fmt.Println(1)
	var root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = &tree.Node{}
	root.Right.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

}
