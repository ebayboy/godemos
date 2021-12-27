package main

import (
	"fmt"
	"math"
)

//import "math"

func main() {
	fmt.Println("vim-go")
	a := math.Inf(1)
	b := math.Inf(-1)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	fmt.Println(math.IsInf(a, 0), math.IsInf(b, 0))
}
