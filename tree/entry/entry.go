package main

import (
	"fmt"
	"github.com/wangmengyu/go-learning/tree"
)

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

}
