package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"github.com/wangmengyu/go-learning/go-websocket/chat"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var manager = chat.NewManager()

func Hello(c *gin.Context) {
	c.String(200, "hello")
}

func WsEndpoint(c *gin.Context) {
	w := c.Writer
	r := c.Request
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.String(200, err.Error())
		return
	}
	log.Println("clients successfully connected")

	client := &chat.Client{
		Id:        uuid.NewV4().String(),
		Websocket: ws,
		Send:      make(chan []byte),
	}
	//注册客户端
	manager.Register <- client
	go client.Read()
	go client.Write()

}

const keyRequestId = "requestId"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	r.Use(func(c *gin.Context) {
		//使用中间件,无论用户访问/ping, 还是/hello,一定先进入到use中
		s := time.Now()
		c.Next() // 这是执行后续请求 例如下方的/ping, 还有/hello
		logger.Info(
			"incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status code", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)
	}, func(c *gin.Context) {
		c.Set(keyRequestId, rand.Int())
		c.Next()
	})
	//创建manager

	go manager.Start()

	r.GET("/hello", Hello)
	r.GET("/ws", WsEndpoint)

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
