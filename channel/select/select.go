package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
  generate a channel
  always put increment number into channel - goroutine, (random time for each number)
*/
func generator() chan int {
	c := make(chan int)
	i := 0
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- i
			i++
		}
	}()
	return c
}

/**
create a channel, set operation of the work
*/
func createWorker(i int) chan int {
	c := make(chan int)
	go doWork(i, c)
	return c

}

func doWork(i int, work chan int) {
	for n := range work {
		time.Sleep(time.Second)
		fmt.Printf("worder %d receice %d \n", i, n)
	}
}

func main() {

	//create two channel to receive numbers
	var c1, c2 = generator(), generator()
	n := 0

	//create one work to receive data from two channel and print data
	worker := createWorker(0)

	//int slice for pending numbers
	values := make([]int, 0)
	actVal := 0
	actWorker := make(chan int)

	tm := time.After(100 * time.Second) // after 10 second, tm channel will receive a time value
	tick := time.Tick(time.Second)      // every second to put time value  into a channel:(a timer)

	for {
		if len(values) > 0 {
			actVal = values[0]
			actWorker = worker
		}
		timeOutTm := time.After(800 * time.Millisecond)
		select {
		case n = <-c1: //if c1 has val, put number to pending slice
			values = append(values, n)
		case n = <-c2: //if c1 has val, put number to pending slice
			values = append(values, n)
		case actWorker <- actVal: //if both c1 and c2 has no data ,get first number from slice and put into worker
			values = values[1:]
		case <-timeOutTm: //if wait time over 800 Millisecond, print time out
			fmt.Println("time out")
		case <-tick:
			fmt.Println("len of queue:", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}

}
