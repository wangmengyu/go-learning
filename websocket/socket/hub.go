package socket

type Hub struct {
	Clients    map[*Client]bool // 当前所含客户端Map
	Broadcast  chan []byte      // 广播通道
	Register   chan *Client     //注册通道
	Unregister chan *Client     //注销通道
}
