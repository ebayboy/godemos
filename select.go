package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func test1() {
	wg := sync.WaitGroup{}

	c1 := make(chan int, 1)
	c2 := make(chan int, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//will block1
		d1, ok := <-c1
		fmt.Println("read d1:", d1, " ok:", ok)

		//will block2
		d2, ok := <-c2
		fmt.Println("read d2:", d2, " ok:", ok)
	}()

	c1 <- 10

	time.Sleep(time.Second * 3)
	c2 <- 11

	wg.Wait()
	fmt.Println("main exit!")
}

func selectTest1() {

	ctx, cancleFunc := context.WithCancel(context.Background())

	c1 := make(chan int, 1)
	c2 := make(chan int, 1)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx.Done! goroutine exit!")
				return
			case d1, ok := <-c1:
				fmt.Println("read d1:", d1, " ok:", ok)
			case d2, ok := <-c2:
				fmt.Println("read d2:", d2, " ok:", ok)
			default:
				fmt.Println("default:", time.Now())
				time.Sleep(time.Millisecond * 500)
			}
		}
	}(ctx)

	c1 <- 10

	time.Sleep(time.Second * 3)
	c2 <- 11

	cancleFunc()

	time.Sleep(time.Second * 1)

	fmt.Println("main exit!")
}

func main() {
	//test1()
	selectTest1()
}
