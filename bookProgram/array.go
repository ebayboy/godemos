package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	for i, _ := range a {
		fmt.Printf("%d %d\n", i, a[i])
	}
}
