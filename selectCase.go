/*
通道select && case 用法
*/
package main

import (
	"fmt"
)

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			//将x写入通道c
			x, y = y, x+y
		case <-quit:
			//读quit通道， 成功读quit通道后函数返回
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		//routine runing
		for i := 0; i < 10; i++ {
			//读c通道， 等待主协程写入
			fmt.Println(<-c)
		}
		//读取10次c通道后，将0写入quit
		quit <- 0
	}()
	//主协程运行fibonacci函数
	fibonacci(c, quit)

	//函数返回后主程序退出
}
