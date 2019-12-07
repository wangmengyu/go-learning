package main

import (
	"fmt"
	"time"
)

func main() {
	chanDemo()
	bufferedChannel()
	ChannelClose()
}

func ChannelClose() {
	c := make(chan int)
	go Worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

/**
create three buffer size of channel
 it can receive  3 data at most with no deadlock
*/
func bufferedChannel() {
	c := make(chan int, 3)
	go Worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)

}

func Worker(i int, c chan int) {
	for num := range c {
		fmt.Printf("worker %d receive %c\n", i, num)
	}
}

func CreateWorker(i int) chan<- int {
	c := make(chan int)
	go Worker(i, c)
	return c
}

/**
  create ten channels of int in an array
  put char into theses channels
  read them  from channels by their goroutine worker
*/
func chanDemo() {

	//create an array  of ten chan of int element
	var channels [10]chan<- int

	//create 10 ten channels with their goroutine
	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i)
	}

	//put ten char into ten channels
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)

}
