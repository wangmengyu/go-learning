package main

import (
	"fmt"
	"sync"
	"time"
)

/**
原子化的int
*/
type atomicInt struct {
	value int
	lock  sync.Mutex // 互斥锁
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	//下面一块函数会被一起执行，或者一起不执行。
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
	}()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())

}
