package main

import (
	"fmt"
	"time"
)

func Sum(s []int, c chan int) {
	sum := 0
	for _, i2 := range s {
		sum += i2
	}
	fmt.Println("sub sleep start")
	time.Sleep(100 * time.Second)
	fmt.Println("sub sleep over")

	//写入结果到通道变量
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go Sum(s[:len(s)/2], c)
	go Sum(s[len(s)/2:], c)

	time.Sleep(3 * time.Second)
	for {
		fmt.Println("main print!")
		time.Sleep(time.Second * 1)
	}

	//读取通道变量
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
