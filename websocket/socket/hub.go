package socket

type Hub struct {
	Clients    map[*Client]bool // 当前所含客户端Map
	Broadcast  chan []byte      // 广播通道
	Register   chan *Client     //注册通道
	Unregister chan *Client     //注销通道
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

//Run 运行hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true // 接收客户
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				//注销客户
				delete(h.Clients, client)
				close(client.Send) //关闭发送通道
			}
		case message := <-h.Broadcast: // 从广播中获得消息
			//循环每个客户端, 把消息发送给每个客户, 没有消息了就关闭掉通道
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					//如果没有消息.删掉客户端
					close(client.Send)
					delete(h.Clients, client) //
				}
			}

		}
	}
}
