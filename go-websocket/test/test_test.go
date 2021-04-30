package testclient

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/wangmengyu/go-learning/go-websocket/chat"
	"log"
	"net/url"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestHello(t *testing.T) {

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	var addr = flag.String("addr", "localhost:8080", "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	socket, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer socket.Close()

	//申请一个ClientManager
	content := chat.Message{
		Sender:    "system",
		Recipient: "1",
		Content:   "hello",
	}
	log.Println("[content]", content)
	jsonData, err := json.Marshal(content)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	done := make(chan struct{})
	for {
		select {
		case v := <-ticker.C:
			fmt.Println(v)
			err = socket.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				log.Println("[SEND ERR]", err.Error())
			} else {
				log.Println("[send success]", content)
			}
		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := socket.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("write err:", err.Error())
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}

	}

}
