package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	cnt := 0

	//创建有缓冲通道, 异步通道
	ch := make(chan int, 1)

	wg.Add(1)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("ch closed by main!")
				break
			}
			cnt += 1 //闭包函数cnt是引用
			fmt.Println("Recv: ", v, " ok:", ok, " cnt:", cnt)
			time.Sleep(time.Second * 1)
		}

		fmt.Println("goroutine exited!")
		wg.Done()
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Send [%d] ok!\n", i)
	}
	close(ch)
	fmt.Println("chan closed!")

	wg.Wait()

	fmt.Println("3 second exit main!")
	time.Sleep(time.Second * 3)
	fmt.Println("main exited! cnt:", cnt)
}
