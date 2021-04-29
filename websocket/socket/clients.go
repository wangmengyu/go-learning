package socket

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
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

	//设置最大消息数量
	c.Conn.SetReadLimit(maxMessageSize)
	//设置最长等待时间
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))

	//设置收到信息之后的操作. 延长等待市场
	c.Conn.SetPongHandler(func(appData string) error {
		_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	//不断接收数据
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			//遇到错误, 关闭响应通道
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		//去掉空格后把消息推送到广播通道里
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.Hub.Broadcast <- message
	}
}

// WritePump 从中心区域抽取数据放到连接的websocket通道内
// 一个 gouroutine 运行 WritePump 为每一个连接, 这个应用 通过执行所有的写操作从gouroutine, 每个连接至少一个write
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	//定时请求器
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		//关闭掉ticker, 关闭掉连接
		ticker.Stop()
		_ = c.Conn.Close()
	}()

}
