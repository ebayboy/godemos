package main

import (
	"fmt"
	"io"
	"net/http"
)

//重点内容:
//http.Get
//resp.Body.Close
//resp.Body.Reader

func main() {

	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Println("status:", resp.Status)
	fmt.Println("header:", resp.Header)

	buf := make([]byte, 4096)

	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("error:", err.Error())
			return
		}
		fmt.Printf("buf:[%v]\n", string(buf[:n]))
		break
	}
}
