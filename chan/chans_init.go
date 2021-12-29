package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	wg := sync.WaitGroup{}

	//1.1 初始化
	ch := make(chan int, 5)
	fmt.Println("ch:", ch)

	//2.1 创建大小为10的数组chan
	chs := make([]chan int, 3)
	fmt.Println("chs:", chs)
	for j := 0; j < len(chs); j++ {
		chs[j] = make(chan int, 2)
	}

	//2.2 初始化chan
	fmt.Println("after init chs:", chs)

	sum := 0
	//启动输出累加协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			num, ok := <-ch
			if !ok {
				break
			}
			fmt.Printf("out chan recv %d\n", num)
			sum += num
		}
	}()

	//启动计算协程组
	for i := 0; i < len(chs); i++ {
		wg.Add(1)
		go func(in chan int) {
			defer wg.Done()
			for {
				data, ok := <-in
				if !ok {
					break
				}
				fmt.Printf("in chan recv %d\n", data)
				data += 1
				ch <- data //write out chan
			}
		}(chs[i])
	}

	for j := 0; j < 5; j++ {
		idx := j % len(chs)
		fmt.Printf("write %d to chan[%d]\n ", j, idx)
		chs[idx] <- j
	}

	//必须有协程退出机制， 否则会deadlock
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		defer wg.Done()

	SIG_LOOP:
		for {
			sig := <-sigCh
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				close(ch)
				for i := 0; i < len(chs); i++ {
					close(chs[i])
				}
				fmt.Println("close all chans done!")
				break SIG_LOOP //break lable 跳出指定代码块
			}
		}
	}()

	//前面的步骤不能少
	wg.Wait()
	close(sigCh)

	fmt.Println("sum:", sum)
}
