package main

import (
	"fmt"
	"github.com/wangmengyu/go-learning/queue"
	//"github.com/wangmengyu/go-learning/tree"
)

func main() {

	fmt.Println(1)
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	q.Pop()
	fmt.Println(q.IsEmpty())
	q.Pop()
	q.Pop()
	fmt.Println(q.IsEmpty())

	q.Push("abc")
	fmt.Println(q.Pop())

}
