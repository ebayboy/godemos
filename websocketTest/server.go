package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var srvAddr = flag.String("srvAddr", "0.0.0.0:8000", "http service srvAddress")

var upgrader = websocket.Upgrader{} // use default options

// 重点关注函数
func echo(w http.ResponseWriter, r *http.Request) {

	fmt.Println("====Enter echo ...")
	fmt.Printf("request:[%v]\n", *r)

	//update http to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s mt:%v", message, mt)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func echo_once(w http.ResponseWriter, r *http.Request) {

	fmt.Println("====Enter echo_once ...")
	fmt.Printf("request:[%v]\n", *r)

	//update http to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte("hello"))
	if err != nil {
		log.Println("write:", err)
	}
	conn.Close()
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	//建立路由， 也key用mux开源模块

	//作用：通过client.go 访问这个路由
	//客户端连接的是ws://127.0.0.1:8000/echo , 这里的路由是http://127.0.0.1:8000/echo, 怎么联通的？
	//使用ws://实际上也是走的http协议，只是请求头有变化
	//websocket 也是走的http 协议，只是一些头部字段有区别， 例如：
	//request:[{GET /echo HTTP/1.1 1 1 map[Connection:[Upgrade] Sec-Websocket-Key:[xJuxUVPvz97A1tt+8Fz8Eg==] Sec-Websocket-Version:[13] Upgrade:[websocket] User-Agent:[Go-http-client/1.1]] {} <nil> 0 [] false 0.0.0.0:8000 map[] map[] <nil> map[] 127.0.0.1:36824 /echo <nil> <nil> <nil> 0xc0000b43c0}]
	//所以到echo的是http,进入echo函数内部会将http请求升级为websocket请求

	http.HandleFunc("/echo", echo)
	http.HandleFunc("/echo_once", echo_once)

	//开启http监听
	fmt.Println("http.ListenAndServe:", *srvAddr)
	err := http.ListenAndServe(*srvAddr, nil)
	if err != nil {
		log.Fatalf("Error:%v\n", err.Error())
	}
}
