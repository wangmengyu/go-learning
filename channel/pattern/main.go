package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 消息生成器， 返回一个传输string的channel, done 代表任务结束的信号. 接到之后关闭掉
func msgGen(name string, done chan struct{}) <-chan string {
	//
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-done:
				// 收到结束命令. 退出程序
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second) // 需要2秒才能清理完
				fmt.Println("cleaning up done")
				done <- struct{}{} // 完成之后向done再发一个空消息体
				return
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				// 等待指定秒数之后哦iu输出字符串到c中.
				c <- fmt.Sprintf("service : %s message %d", name, i)
			}
			i++
		}
	}()
	return c
}

// 在不明确c1,c2 具体数量的时候用这种方式.
func fanIn(chs ...<-chan string) chan string {
	c := make(chan string)

	// 谁的数据做完了就先处理谁
	for _, ch := range chs {
		// 以下goroutine中永远是第二个ch.如果不传参数到goroutine 全局会变成第二个ch.
		go func(in <-chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}

	return c
}

// nonBlockingWait 非阻塞式锁 , 返回等到/未等到, 和等到的字符串
func nonBlockingWait(c <-chan string) (string, bool) {

	select {
	case m := <-c: // 等到直接返回
		return m, true
	default: // 非阻塞等待. 没有等到
		return "", false
	}

}

// timeoutWait 超时等待 . 在指定时间内等到了就返回, 没有等到就返回已经 没有等到消息
func timeoutWait(c <-chan string, duration time.Duration) (string, bool) {
	select {
	case m := <-c: // 等到了直接返回得到的数据
		return m, true
	case <-time.After(duration):
		return "", false // 指定时间内没有等到就直接返回没有等到

	}
}

// fanInBySelect 在明确知道c1,c2 数量的情况下用 select
func fanInBySelect(c1, c2 <-chan string) chan string {

	c := make(chan string)
	go func() {
		for {
			select {
			case v := <-c1:
				c <- v
			case v := <-c2:
				c <- v

			}
		}
	}()
	return c
}
func main() {
	done := make(chan struct{})
	m1 := msgGen("service 1", done) // 不断生成数据
	//m2 := msgGen("service 2")

	for i := 1; i <= 5; i++ {
		//fmt.Println(<-m1)                                      // 打m1中收到的数值
		if v2, ok := timeoutWait(m1, 1*time.Second); ok { // 持续尝试从M2中获得数据.
			fmt.Println("receive from service 1 : ", v2)
		} else {
			fmt.Println(" timeout ")
		}
	}
	done <- struct{}{}

	<-done // 收到以后再推出main函数, 通过done channel的通知, 来进行优雅退出
	fmt.Println("done")

}
