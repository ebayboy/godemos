package main

import (
	"fmt"
	"net/http"
)

func myHandler(resp http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "连接成功")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL)
	fmt.Println("body:", r.Body)
	fmt.Printf("header:%v\n", r.Header)
	fmt.Println("user-agent:", r.UserAgent())

	//write response
	resp.Write(([]byte)("www.5lhm.com"))
}

func main() {
	//注册路由
	http.HandleFunc("/go", myHandler)

	//启动监听
	fmt.Println("http.ListenAndServe 127.0.0.1:8000")
	http.ListenAndServe("127.0.0.1:8000", nil)
}
