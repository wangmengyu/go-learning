package main

import (
	"fmt"
	"sync"
)

/**
worker 结构
包含两个属性,
一个用于接收数据的channel
一个用于完成done通知动作的function
*/
type worker struct {
	in   chan int
	done func()
}

/**
  初始化worker对象，
  包含一个用于数据接收的channel,
  还有一个完成通知的waitGroup指针
*/
func createWork(i int, wg *sync.WaitGroup) worker {
	//定义一个worker 对象， 里面包含一个channel, 和go func 不断读取channel中的数据
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	//开启并发动作
	go doWork(i, w)
	return w
}

/**
创建work , 接收来自缓冲类channel的数据
done 参数， 当执行完成后，使用channel通知完成
*/
func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("work %d, received %c\n", id, n)
		//需要在开个goroutine， 不会卡住
		w.done()

	}
	fmt.Println("no data")
}

func chanDemo() {

	//定义10个元素的 worker类的数组
	var workers [10]worker

	//一个waitGroup对象，预计等待20个完成
	var wg sync.WaitGroup
	wg.Add(20) // 等待20个任务被完成

	//为10个worker元素初始化
	for i := 0; i < 10; i++ {
		workers[i] = createWork(i, &wg)
	}

	//内置等待工具
	for i, work := range workers {
		work.in <- 'a' + i
	}

	for i, work := range workers {
		work.in <- 'A' + i
	}

	//执行等待
	wg.Wait()

}

func main() {

	chanDemo()

}
