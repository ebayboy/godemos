package main

import (
	"fmt"
	"sync"
	"time"
)

// 10个协程处理1000个任务

func main() {

	ch := make(chan int)

	TCount := 10
	var wg sync.WaitGroup

	wg.Add(TCount)
	for i := 0; i < TCount; i++ {
		go func() {
			for {
				select {
				case data, ok := <-ch:
					if !ok {
						fmt.Println("exit! ok:", ok)
						wg.Done()
						return

					}
					fmt.Println("read data:", data, " ok:", ok)
				default:
					time.Sleep(time.Second)
					fmt.Println("default")

				}

			}

		}()

	}

	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i

		}
		close(ch)

	}()

	wg.Wait()

	fmt.Println("hello")

}
