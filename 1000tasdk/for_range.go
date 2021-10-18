package main

import "fmt"

func main() {

	go func(in <-chan int) {
		// Using for-range to exit goroutine
		// range has the ability to detect the close/end of a channel
		for x := range in {
			fmt.Printf("Process %d\n", x)
		}
	}(inCh)
}
