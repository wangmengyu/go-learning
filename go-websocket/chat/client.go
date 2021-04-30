package chat

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct { //客户, 每个客户持有socket连接
	Manager   *ClientManager
	Id        string
	Websocket *websocket.Conn
	Send      chan []byte // 发送消息通道
}

type Message struct { // 消息结构体.
	Sender    string `json:"sender,omitempty"`    // 发送人
	Recipient string `json:"recipient,omitempty"` //接收人
	Content   string `json:"content,omitempty"`   // 内容/
}

// Read 不断从socket连接 读取消息.
func (c *Client) Read() {

	defer func() {
		//注销并关闭连接
		c.Manager.Unregister <- c
		_ = c.Websocket.Close()

	}()
	for {
		_, message, err := c.Websocket.ReadMessage()
		if err != nil {
			//注销, 关闭通道
			c.Manager.Unregister <- c
			_ = c.Websocket.Close()
			break
		}

		//将信息写入到广播通道
		jsonMsg, _ := json.Marshal(&Message{
			Sender:  c.Id,
			Content: string(message),
		})
		log.Println("[send msg to broadcast]", string(jsonMsg))
		c.Manager.Broadcast <- jsonMsg

	}
}

// Write 从 send 通道读取数据 并且写入到通道里.
func (c *Client) Write() {
	defer func() {
		//关闭掉连接
		_ = c.Websocket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				//写一个关闭消息, 返回
				_ = c.Websocket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//将消息写入到socket 中
			log.Printf("write msg [%s] from send channel in client[%s] \b", string(message), c.Id)
			_ = c.Websocket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
