package main

import (
	"fmt"
	"net"
)

/*
  1.建立与服务端的链接
  2.进行数据收发
  3.关闭链接
*/

func main() {

	addr := "127.0.0.1:20000"

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	defer conn.Close()

	msg := ([]byte)("hello world")
	for {
		n, err := conn.Write(msg)
		if err != nil {
			fmt.Println("Error:", err.Error())
			break
		}
		fmt.Println("write size:", n)
		if n == len(msg) {
			break
		}
	}
}
