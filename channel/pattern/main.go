package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 消息生成器， 返回一个传输string的channel
func msgGen(name string) <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service : %s message %d", name, i)
			i++
		}
	}()
	return c
}

// 在不明确c1,c2 具体数量的时候用这种方式.
func fanIn(chs ...<-chan string) chan string {
	c := make(chan string)

	// 谁的数据做完了就先处理谁
	for _, ch := range chs {
		// 以下goroutine中永远是第二个ch.如果不传参数到goroutine 全局会变成第二个ch.
		go func(in <-chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}

	return c
}

// fanInBySelect 在明确知道c1,c2 数量的情况下用 select
func fanInBySelect(c1, c2 <-chan string) chan string {

	c := make(chan string)
	go func() {
		for {
			select {
			case v := <-c1:
				c <- v
			case v := <-c2:
				c <- v

			}
		}
	}()
	return c
}
func main() {
	m1 := msgGen("service 1") // 不断生成数据
	m2 := msgGen("service 2")
	m3 := msgGen("service 3")
	m := fanIn(m1, m2, m3)
	for {
		fmt.Println(<-m) //打印生成的字符串.
	}

}
