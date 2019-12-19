package main

import (
	"fmt"
	rpcdemo "github.com/wangmengyu/go-learning/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	//telnet localhost 1234
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div",
		rpcdemo.Args{A: 10, B: 4}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	err = client.Call("DemoService.Div",
		rpcdemo.Args{A: 3, B: 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
