package queue

// 整数切片类型的类，名字是Queue
type Queue []int

//push元素到队列中
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

//从队列中获取第一个元素并从原有的队列里删掉第一个元素
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
