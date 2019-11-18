package mock

type Retriever struct {
	// 这里类的名字最好是Retriever, 不要叫MockRetriever ,因为已经在mock下了，不要重复前缀
	// 鼠标移动到IDE中 Retriever 类名上 出现小黄灯可以创建实现某个接口
	Contents string
}

/**
要修改结构体内容 需要指针传入
*/
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

/**
最好都用指针接受者，因为表示操作同一个对象
*/
func (r *Retriever) Get(url string) string {
	return r.Contents
}
