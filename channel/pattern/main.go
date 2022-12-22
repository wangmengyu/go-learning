package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 消息生成器， 返回一个传输string的channel
func msgGen() <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("message %d", i)
			i++
		}
	}()
	return c
}
func main() {
	m := msgGen() // 不断生成数据
	for {
		fmt.Println(<-m) //打印生成的字符串.
		// m <- "abc"
	}

}
