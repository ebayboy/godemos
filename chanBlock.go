package main

import (
	"fmt"
	"time"
)

/*
- 往一个已经关闭了的通道里面发送值时会引发“panic”。比如上面注释报错3处，前面已执行“close(channel3)”关闭通道操作，再往里面发送值就会引发panic。
- 关闭一个已经关闭的通道时，会引发“panic”。比如上面注释“报错4”处。
- 示例5和示例6的区别仅仅在于关闭通道后，里面是否还有值剩余？假设有剩余，我们就可以从通道取值同时赋给两个变量，第二个变量是bool类型值，
其为true表示取到了值，其为false表示没有取到值，这样仅仅可以避免引发“panic”,如果通道已经关闭且无元素值，
则取出的第二个bool值为false;若从已关闭的通道里面（里面无剩余元素值）再次读取元素值，则第二个值为true。
：第二个bool值为false,则通道肯定关闭了，值为true,可能关闭也可能没有关闭
*/

func main() {

	// 示例2
	channel := make(chan int, 3)
	//写通道
	channel <- 1
	channel <- 2
	channel <- 3

	// 报错1：fatal error: all goroutines are asleep - deadlock!
	//只有三个通道变量，已经写满， 在读取之前继续写会阻塞
	//channel <- 4

	fmt.Printf("Outsize write the first channel value is %v\n", <-channel)
	//读通道
	v := <-channel
	fmt.Printf("Outsize write the first channel value is %v\n", v)
	fmt.Printf("Outsize write the first channel value is %v\n", <-channel)

	//3个通道变量已经全部被读取， 继续读会阻塞
	// 报错2：fatal error: all goroutines are asleep - deadlock!
	//fmt.Printf("the first channel value is %v\n",<-channel)

	// 示例3 goroutine sleep, 等待主进程写入， sleep时间过后读取写入的数据
	channel2 := make(chan int, 0)
	go func() {
		time.Sleep(time.Second * 5)
		v := <-channel2
		fmt.Printf("In goroutine read the value is %v\n", v)
	}()
	channel2 <- 1
	fmt.Print("the time is over\n")
}
