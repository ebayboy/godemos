package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	inCh := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(in <-chan int) {
			for x := range in {
				fmt.Printf("Process %d\n", x)
			}
			wg.Done()
		}(inCh)
	}

	go func() {
		for i := 0; i < 1000; i++ {
			inCh <- i
		}
		close(inCh)
	}()

	wg.Wait()
}
