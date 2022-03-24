package main

import (
	"fmt"
	"time"
)

func longOperation(ch chan<- string) {
	// Simulate long operation.
	// Change it to more than 10 seconds to get server timeout.
	time.Sleep(time.Second * 3)
	ch <- "Successful result."
}

func Pooling() string {
	ch := make(chan string)
	defer close(ch)
	tick := time.Tick(time.Second)

	go longOperation(ch)
	cnt := 0
	result := ""

	for {
		select {
		case result = <-ch:
			return result
		case <-tick:
			cnt++
			fmt.Println("polling...", cnt)
		case <-time.After(time.Second * 10):
			fmt.Println("Server is busy.")
			<-ch
			return ""
		}
	}

}

func main() {
	res := Pooling()
	fmt.Println(res)
}
