package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Length int
	Width  int
}

func main() {
	// 1. rpc dailhttp
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatalln("Error:", err.Error())
	}

	// 2. call rpc
	p := Params{Width: 3, Length: 5}
	var ret int = 0
	if err := conn.Call("Rect.Area", &p, &ret); err != nil {
		log.Fatalln("Error:", err.Error())
	}
	fmt.Printf("Area:%v\n", ret)
}
