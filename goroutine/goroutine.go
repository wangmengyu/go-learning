package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) { //go 关键词后定义的匿名函数可以被并发执行。
			for {
				//a[i]++ // 此处不执行IO操作或者runtime.Gosched()就没有机会交出控制权。会导致不断执行
				//runtime.Gosched() // 手动交出控制权
				fmt.Printf("Hello from goroutine %d\n", i)
			}

		}(i)
	}
	time.Sleep(time.Minute)
}
