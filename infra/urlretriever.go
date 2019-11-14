package infra

import (
	"io/ioutil"
	"net/http"
)

//定义一个接受者结构体
type Retriever struct {
}

/**
获取某个网页的内容
*/
func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
