package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "0.0.0.0:8000", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	//websocket的schema是ws, 例如:  ws://127.0.0.1:8000/
	//url: ws://127.0.0.1:8000/echo
	urlStr := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", urlStr.String())

	//Dial
	conn, _, err := websocket.DefaultDialer.Dial(urlStr.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	//启动协程读取websocket消息
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			//服务器主动关闭处理
			return
		case timeNow := <-ticker.C:
			//每秒向websocket连接写入消息
			err := conn.WriteMessage(websocket.TextMessage, []byte(timeNow.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}

			//收到中断信号, 写入websocket关闭消息
		case <-interrupt:
			//客户端主动关闭处理
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			//收到服务器关闭确认消息后, 1秒后退出程序
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}