package channel

import (
	"fmt"
	"time"
)

/**
创建work
*/
func createWork(i int) chan int {
	//定义一个channel, 和go func 不断读取channel中的数据
	c := make(chan int)
	go worker(i, c)
	return c
}

func chanDemo() {

	//定义10个元素的协程数组
	var channels [10]chan int

	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)

}

/**
创建work , 接收来自缓冲类channel的数据
*/
func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("work %d, received %c\n", id, n)
	}
	fmt.Println("no data")
}

/**
缓冲类的channel,
初始化时，需要给出缓冲长度。
如果发送数据时，超过缓冲区长度，有没有人接收数据，会报deadlock
定义好 go work 协程一定要写在数据发送之前，才能防止deadlock, 缓存式的写法性能会比不缓存的好
*/
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)

}

/**
channel的关闭，需要在使用前判断收到的字符是否为空, 语法 for n:= range c，为空的跳出，不为空的执行
close()调用在发送数据之后。
*/
func channelClose() {

	c := make(chan int, 3)
	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c)

	time.Sleep(time.Millisecond)

}
func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	fmt.Println("Buffered Channel")
	bufferedChannel()
	fmt.Println("Channel close")
	channelClose()

}
