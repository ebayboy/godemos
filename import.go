package main

import (
	"fmt"

	"./Sumx"
)

func main() {
	s := Sumx.NewSumx(1, 2)
	fmt.Println("s:", s.Add())
}
