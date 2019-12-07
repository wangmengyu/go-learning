package main

import (
	"fmt"
	"sync"
)

func main() {
	chanDemo()
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func CreateWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{in: make(chan int), wg: wg}

	//create goroutine for the worker
	go DoWorker(i, w)

	return w

}

/**
receive data from channel, and print them
*/
func DoWorker(i int, w worker) {
	for {
		fmt.Printf("worker %d receive %c\n", i, <-w.in)
		w.wg.Done()
	}
}

func chanDemo() {

	//create a waitGroup variable
	var wg sync.WaitGroup

	//create ten workers
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(i, &wg)
	}
	//add work number for wg
	wg.Add(20)
	//concurrently print a to j
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	//concurrently print A to J
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	//wait for all worker done
	wg.Wait()

}
