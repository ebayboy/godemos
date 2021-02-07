package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	c1 := make(chan int)
	fmt.Println("cap(c1):", cap(c1))

	sg := sync.WaitGroup{}

	sg.Add(1)
	go func() {
		fmt.Println("arr:", arr)
		for _, v := range arr {
			fmt.Println("send v:", v)
			c1 <- v
		}
		sg.Done()
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		default: //省略会出现问题： fatal error: all goroutines are asleep - deadlock!
		}
	}

	sg.Wait()
	fmt.Println("main over")
}
