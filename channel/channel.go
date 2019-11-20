package main

import (
	"fmt"
	"time"
)

func work(id int, c chan int) {
	for { // 不断的在接收数据,如果不写成死循环，只获取一次，下面如果发送了不止一次就会 deadlock

		fmt.Printf("Work %d Received  %c\n", id, <-c)
	}
}

func chanDemo() {

	var channels [10]chan int
	// 发送数据没人收是会 deadlock 的
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int) // 定义变量c, 类型是个channel, 里面内容是个int, 只有用make方法定义的channel才可以进行使用
		go work(i, channels[i])      // 将10个channel分发给10个work
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)

}
func main() {
	chanDemo()

}
