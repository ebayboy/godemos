package main

import "fmt"

//无缓冲通道(同步通道)
//必须有接收方先接收， 否则发送出错
func recv(ch chan int) {
	ret := <-ch
	fmt.Println("recv ok! ret:", ret)
}

func main() {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("send ok!")
}
