package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, i2 := range s {
		sum += i2
	}

	//写入结果到通道变量
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	//读取通道变量
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
