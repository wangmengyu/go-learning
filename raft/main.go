package main

import (
	"fmt"
	"time"
)

type Node struct {
	State      string
	Term       int      // Raft协议关键概念，每个term内都会产生一个新的leader
	Val        int      // 存储的值
	MsgChannel chan int // 接收leader消息的channel
}

func CreateWorker() {

}

func DoWorker() {
	//等待3秒后, 期间没有收到领导者发来的信息, 自己变成Candidate身份.

}
func main() {
	nodes := make([]Node, 0)
	for i := 0; i < 3; i++ {
		nodes = append(nodes, Node{State: "Follower", Term: 0})
	}

	//每个节点给一个监听的worker.
	tick := time.Tick(3 * time.Second)

	//假设先判断第一个节点, 3
	for {
		//远远不断的从第一个节点的msg通道获得消息.
		select {
		case n := <-nodes[0].MsgChannel:
			//收到了通知
			fmt.Println("receive n=", n)
			return
		case <-tick:
			//3秒内判断有没有收到通知, 且当前是follower身份的, 提升为candidate
			if nodes[0].State == "Follower" {
				nodes[0].State = "Candidate"
			}
		}
	}

}
