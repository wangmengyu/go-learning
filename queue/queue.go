package queue

// 整数切片类型的类，名字是Queue。目的，定义别名. 元素可以是任意类型的
type Queue []interface{}

//push元素到队列中
func (q *Queue) Push(v interface{}) {
	//指针接收者，是可以改变又有的对象的
	*q = append(*q, v)
}

//从队列中获取第一个元素并从原有的队列里删掉第一个元素
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
