package main

import (
	"fmt"
	"github.com/wangmengyu/go-learning/infra"
)

/**
获取接受者对象，需要独立建立方法，
原因： 可以快速修改方法供应的类
*/
func getRetriever() retrieverI {
	return infra.Retriever{}
}

// 需要一个东西，可以获取URL网页的内容，
// 定义个retriever接口,接口类名和各个厂家写的struct类名要一致,方法
type retrieverI interface {
	Get(string) string
}

func main() {
	var r = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com/"))
}
