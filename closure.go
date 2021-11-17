//测试闭包函数

package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2

	Add := func() int {
		a++
		return a + b
	}

	c := Add()

	fmt.Println("c:", c, "a:", a)
}
