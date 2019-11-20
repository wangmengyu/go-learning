package main

import (
	"fmt"
	"github.com/wangmengyu/go-learning/tree"
)

/**
对于tree.Node的扩展类 - 内嵌做法
*/
type myTreeNode struct {
	*tree.Node //Embedding
}

/**
扩展了之前遍历方法
左，右，中 顺序遍历树, 后续遍历
*/
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()

}

func (myNode *myTreeNode) Traverse() { //重载
	fmt.Println("This is shadowed")
}

func main() {
	//定义treeMode类型值，全部用点访问下去元素，
	fmt.Println(1)
	var root = myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = &tree.Node{}
	root.Right.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	fmt.Println("In-order traversal")
	root.Traverse()      //调用自己的（重载后）的遍历方法
	root.Node.Traverse() // 调用原有的遍历方法
	cnt := 0
	root.TraverseFunc(func(node *tree.Node) {
		cnt++
		fmt.Println("TR-FUNC:", node.Value)
	})
	fmt.Println("cnt=", cnt)
	fmt.Println()
	fmt.Println("post-order traversal")
	root.postOrder()
	fmt.Println()

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}

	fmt.Println("max node value :", maxNode)

}
