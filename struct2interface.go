package main

import (
	"fmt"
)

type ST struct {
	Name string
	Age  int
}

func Hello(v interface{}) {
	fmt.Println("hello", v)
	//空接口转为结构体
	s := v.(ST)
	fmt.Println(s)
}

func main() {
	s := ST{Name: "fanpf", Age: 30}
	Hello(s)
	fmt.Println(s)
}
