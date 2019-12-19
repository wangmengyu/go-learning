package main

import (
	rpcdemo "github.com/wangmengyu/go-learning/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
  简单的rpc服务器
*/
func main() {
	//注册rpc服务
	_ = rpc.Register(rpcdemo.DemoService{})

	//监听tcp 1234端口
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	//一直从tcp接收信息，
	for {
		conn, err := listener.Accept()
		if err != nil {
			//发生错误，接收下一个请求
			log.Printf("accept err:%v", err)
			continue
		}
		//接收成功，异步的完成服务
		go jsonrpc.ServeConn(conn)
	}
}
