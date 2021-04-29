package socket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second //等待输入的时间 10秒

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second //读取回应的等待时间 60秒

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10 // 等到序响应过程中,提前发送的时间

	// Maximum message size allowed from peer.
	maxMessageSize = 512 //最大消息数量限制.
)

type Client struct {
	Hub  *Hub            // 所在中心
	Conn *websocket.Conn // 持有连接
	Send chan []byte     // 发送通道
}

var (
	newline = []byte{'\n'} //回车
	space   = []byte{' '}  // 空格
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // 检查origin
	},
}

// ReadPump 从管道中抽取消息, 放入到Hub中
func (c *Client) ReadPump() {

	//预定注销客户端, 关闭连接
	defer func() {
		c.Hub.Unregister <- c
		_ = c.Conn.Close()
	}()

}
