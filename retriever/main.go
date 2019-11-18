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

/**
  POST数据的接口定义
  可以提交数据到指定Url
*/
type Poster interface {
	Post(url string, form map[string]string) string
}

//下载方法，需要 调用 接口类， 用这个类的实例，来调用方法。
func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func post(post Poster) {
	post.Post("https://www.imooc.com",
		map[string]string{
			"name":   "wmy",
			"course": "golang",
		})
}

/**
  组合的接口，里面包含一个Retriever 和一个Poster接口
*/
type RetrieverPoster interface {
	Retriever
	Poster
}

/**
这个方法你要用到Retriever接口的方法。又要用到Poster的方法
*/
const url = "https://www.imooc.com"

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another faked imooc.come"})
	return s.Get(url)

}

func inspect(r Retriever) {
	//fmt.Printf("%T %v",r,r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent, "TimeOut:", v.TimeOut)
	default:
		fmt.Println(reflect.TypeOf(r))
	}

}

func main() {

	var r Retriever

	retriever := mock.Retriever{Contents: "This is mock"}

	inspect(r)

	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	//fmt.Println(download(r))
	inspect(r)

	//类型判断: Type assertion
	if realRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(realRetriever.Contents)
	} else {
		fmt.Println("not Mock retriever")
	}
	//fmt.Println(realRetriever.Contents)
	//mockRetriever := r.(mock.Retriever)
	//fmt.Println(mockRetriever.Contents)

	fmt.Println(session(&retriever))

	//

}
