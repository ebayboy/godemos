package main

import (
	"fmt"
	"time"
)

func main() {
	//创建有缓冲通道, 异步通道
	ch := make(chan int, 1)
	go func() {
		for v := range ch {
			fmt.Println("recv v:", v)
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("Send [%d] ok!\n", i)
		time.Sleep(time.Second * 1)
	}
}
