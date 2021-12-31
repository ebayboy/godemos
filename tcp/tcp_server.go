package main

import (
	"fmt"
	"net"
)

//tcp server

/*
   1.监听端口
   2.接收客户端请求建立链接
   3.创建goroutine处理链接。
*/

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("Error:%v\n", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Listen:", "127.0.0.1:20000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error:%v\n", err.Error())
			continue
		}

		fmt.Println("Accept conn:", conn)
		go func(conn net.Conn) {
			defer conn.Close()

			buf := make([]byte, 4096)
			var recvStr string
			for {
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Printf("Error : read:%s\n", err.Error())
					break
				}
				recvStr = string(buf[:n])
				fmt.Printf("Read:[%d][%v]\n", n, recvStr)
				conn.Write([]byte(recvStr))
			}
		}(conn)
	}
}
