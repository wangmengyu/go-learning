package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
channel 生成器
*/
func generator() chan int {
	out := make(chan int)
	go func() {
		//开启一个并发，每1500毫秒以内推送一个i给当前创建的通道
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

/**
创建work
*/
func createWork(i int) chan int {
	//定义一个channel, 和go func 不断读取channel中的数据
	c := make(chan int)
	go worker(i, c)
	return c
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("work %d, received %d\n", id, n)
	}
	fmt.Println("no data")
}

func main() {
	//定义两个存储整数的channel
	c1, c2 := generator(), generator()
	var worker = createWork(0) //创建一个worker , 它会并发的打印收到的整数
	n := 0
	var values []int                   //定义个切片，用于收集所有输出的数字
	tm := time.After(10 * time.Second) //time.After返回的是个通道， 时间一到会里面里会有当前的时间字符串进去
	tick := time.Tick(time.Second)     // 每一秒像tick发出信号

	for {
		//无限循环，从通道接收数据，有可能是C1有数据有可能是C2有出数据，用select做出不同的操作。
		var activeWorker chan int // 定义激活中的workder
		var activeValue int       //定义当前输出的数字
		//检查当前是否有值，有值的话，将激活通道设置为之前定义的workder,要是没有值就不处理
		if len(values) > 0 {
			//如果切片长度>0, 设置当前激活的work, 并将最早的切片元素进行存储
			activeWorker = worker
			activeValue = values[0]

		}

		select {
		case n = <-c1: //如果是1通道来的。存储数字到序列
			values = append(values, n)
		case n = <-c2: //如果是2通道来的，存储数字到序列
			values = append(values, n)
		case activeWorker <- activeValue: //两个通道都没有，则将当前激活的数字给到当前激活的work
			values = values[1:] // 并且去掉已处理的数字
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out") // 每次循环超过1500毫秒会报错time out
		case <-tick:
			fmt.Println("len of queue=", len(values))
		case n2 := <-tm:
			fmt.Println("buy, now:", n2)
			return //10秒到了后退出运行

		}
	}
}
