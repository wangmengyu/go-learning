package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wangmengyu/go-learning/websocket/socket"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const keyRequestId = "requestId"

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

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.GET("/ws", socket.ServeWs)

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
