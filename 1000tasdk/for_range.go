package main

import (
	"fmt"
	"sync"
	"time"
)

//问题： 使用10个协程对1-1000进行累加， 并输出累加后的总和.
//1. 使用goroutine + channel方式?
//2. 使用互斥锁实现?

func method2() int {
	var wg sync.WaitGroup
	sum := 0
	var l sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		l.Lock()
		go func() {
			for j := 0; j < 100; j++ {
				sum += j
			}
			l.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("method2 sum:", sum)
	return sum
}

func method1() int {
	var wg sync.WaitGroup

	inCh := make(chan int)
	outCh := make(chan int)
	sum := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// ch <- chan sometype 只读通道
		// ch chan <- sometype 写入通道
		go func(iCh <-chan int, oCh chan<- int) {
			for x := range iCh {
				fmt.Printf("Process %d\n", x)
				oCh <- x
			}
			wg.Done()
		}(inCh, outCh)
	}

	go func() {
		for {
			select {
			case data, ok := <-outCh:
				if !ok {
					//检查到通道关闭后退出协程
					fmt.Println("Error: channel is closed! exit goroutine!")
					return
				}
				sum += data
				fmt.Println("sum:", sum)
			default:
				time.Sleep(time.Millisecond * 10)
			}
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			inCh <- i
		}
		close(inCh)
	}()

	wg.Wait()
	close(outCh)

	fmt.Println("Sum:", sum)

	return sum
}

func main() {
	fmt.Println("methiod1:", method1())
	fmt.Println("methiod2:", method2())
}
