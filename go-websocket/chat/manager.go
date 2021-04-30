package chat

import (
	"encoding/json"
	"log"
)

type ClientManager struct { // 客户端管理器
	Clients    map[*Client]bool //持有的客户端连接
	Broadcast  chan []byte      // 广播通道
	Register   chan *Client     // 注册管道
	Unregister chan *Client     // 注销管道
}

func NewManager() *ClientManager {
	return &ClientManager{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (manager *ClientManager) Start() { // 启动客户端管理器
	for {
		select {
		case client := <-manager.Register: // 注册客户
			manager.Clients[client] = true
			client.Manager = manager
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.Send(jsonMessage, client) //
		case client := <-manager.Unregister: // 注销客户
			if _, ok := manager.Clients[client]; ok {
				close(client.Send)
				delete(manager.Clients, client)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.Send(jsonMessage, client)
			}
		case message := <-manager.Broadcast: // 从广播读取消息发给每个用户.
			for client := range manager.Clients {
				select {
				case client.Send <- message: // 监听消息,发送到每个客户
					log.Printf("send msg[%s] to client [%s]\n", string(message), client.Id)
				default: //没有消息时关闭连接.
					close(client.Send)
					delete(manager.Clients, client)
				}
			}
		}
	}
}

func (manager *ClientManager) Send(message []byte, ignore *Client) { // 群发消息, 将消息发送到每个客户的send通道里, 除了指定的用户.
	for conn := range manager.Clients {
		if conn != ignore {
			conn.Send <- message
		}
	}
}
