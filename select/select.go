package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成 不断接收数据的 channel
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	//定义两个channel , 同时接收数据, 谁快先处理谁,.
	var c1, c2 = generator(), generator()
	// 一直不断的监听c1, c2 通道里有没有数据并且打印
	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1: ", n)
		case n := <-c2:
			fmt.Println("Received from c2: ", n)

		}
	}

}
