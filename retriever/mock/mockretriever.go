package mock

type Retriever struct {
	// 这里类的名字最好是Retriever, 不要叫MockRetriever ,因为已经在mock下了，不要重复前缀
	// 鼠标移动到IDE中 Retriever 类名上 出现小黄灯可以创建实现某个接口
	Contents string
}

func (r Retriever) Get(url string) string {
	return r.Contents
}
