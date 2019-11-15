package main

import (
	"fmt"
	"github.com/wangmengyu/go-learning/retriever/mock"
	real2 "github.com/wangmengyu/go-learning/retriever/real"
	"reflect"
	"time"
)

/**
  接口定义
*/
type Retriever interface {
	Get(url string) string //获取指定URL页面内容的接口
}

//下载方法，需要 调用 接口类， 用这个类的实例，来调用方法。
func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func inspect(r Retriever) {
	//fmt.Printf("%T %v",r,r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent, "TimeOut:", v.TimeOut)
	default:
		fmt.Println(reflect.TypeOf(r))
	}

}

func main() {

	var r Retriever
	r = mock.Retriever{Contents: "This is mock"}
	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	//fmt.Println(download(r))
	inspect(r)

}
