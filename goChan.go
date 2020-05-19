package main

import "fmt"

func main() {

	// 示例4
	channel3 := make(chan int, 2)
	channel3 <- 1
	channel3 <- 2
	close(channel3)

	// 报错3: panic: send on closed channel
	// channel3 <- 3

	// 报错4：panic: close of closed channel
	//close(channel3)

	// 示例5
	channel5 := make(chan int, 2)
	channel5 <- 1
	channel5 <- 2

	v1, b1 := <-channel5
	fmt.Printf("v1:%v  b1:%v\n", v1, b1)

	v2, b2 := <-channel5
	fmt.Printf("v2:%v  b2:%v\n", v2, b2)

	close(channel5)
	v3, b3 := <-channel5
	fmt.Printf("v3:%v  b3:%v\n", v3, b3)
	/*输出：
	  v1:1  b1:true
	  v2:2  b2:true
	  v3:0  b3:false
	*/

	// 示例6
	channel6 := make(chan int, 2)
	channel6 <- 1
	channel6 <- 2

	v4, b4 := <-channel6
	fmt.Printf("v4:%v  b4:%v\n", v4, b4)

	close(channel6)
	v5, b5 := <-channel6
	fmt.Printf("v5:%v  b5:%v\n", v5, b5)

	/*输出：
	  v4:1  b4:true
	  v5:2  b5:true
	*/
}
