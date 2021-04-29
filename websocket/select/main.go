package main

import (
	"fmt"
	"time"
)

func main() {

	// create two channel
	chA := make(chan int, 10)
	chB := make(chan int, 50)

	// 起一个监听协程, 当两个管道进入数据的时候, 要做什么
	go func() {
		for { // 一直监听
			select {
			case va := <-chA:
				fmt.Printf("channel A receive %d", va)
			case vb := <-chB:
				fmt.Printf("channel B receive %d", vb)
			default:
				time.Sleep(2 * time.Second) // 睡两秒
				fmt.Printf("nothing happen")
			}
		}
	}()
	//又是一个协程, 放入数据到通道A
	go func() {
		chA <- 1
		chA <- 2
	}()

	//这里是主进程的, 放入数据到通道B
	chB <- 1
	chB <- 2

}
