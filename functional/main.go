package main

import (
	"bufio"
	"fmt"
	"github.com/wangmengyu/go-learning/functional/fib"
	"io"
	"strings"
)

func main() {
	a := fib.Fibonacci()
	printFileContents(a)

}

/**
  打印文件的每一行内容
*/
func printFileContents(reader io.Reader) {
	scaner := bufio.NewScanner(reader)
	for scaner.Scan() { // 只保留退出条件，连段分号都可以省略，等效于while
		fmt.Println(scaner.Text())
	}
}

//定义一个类，实现reader的接口
type intGen func() int

//函数也能实现接口。因为函数本质上是一个type, 任何的type 是可以实现接口的
func (g intGen) Read(p []byte) (n int, err error) {
	next := g() // 取得下一个元素
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p) //让S成为下一个序列数。
}
