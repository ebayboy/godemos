package main

import (
	"context"
	"fmt"
	"time"
)

func DoWork(ctx context.Context, idx int) {
	fmt.Println("Start goroutne: ", idx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutne done! idx:", idx)
			return
		default:
			fmt.Println("goroutne alive... idx:", idx)
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {

	//1. init ctx
	ctx, cancelFunc := context.WithCancel(context.Background())

	//2. start goroutine
	for i := 0; i < 5; i++ {
		go DoWork(ctx, i)
	}

	fmt.Println("Main will sleep 5...")
	time.Sleep(time.Second * 5)

	cancelFunc()
	//3. execute cancelFunc: goroutine  case <-ctx.Done() -> cancelFunc -> main exit
}
