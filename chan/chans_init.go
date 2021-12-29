package main

import "fmt"

func main() {
	//1.1 初始化
	ch := make(chan int, 5)
	fmt.Println("ch:", ch)

	//2.1 未初始化
	chs := make([]chan int, 5)
	fmt.Println("chs:", chs)
	for j := 0; j < len(chs); j++ {
		chs[j] = make(chan int, 2)
	}
	//2.2 初始化
	fmt.Println("after init chs:", chs)
}
