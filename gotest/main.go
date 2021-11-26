package main

import (
	"fmt"

	"github.com/godemos/gotest/lib"
)

func main() {
	ret := lib.Fibonacci(32)
	fmt.Println("ret:", ret)
}
