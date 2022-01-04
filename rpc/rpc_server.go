package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

//例题：golang实现RPC程序，实现求矩形面积和周长
/*
golang写RPC程序，必须符合4个基本条件，不然RPC用不了
+ 结构体字段首字母要大写，可以别人调用 [done]
+ 函数名必须首字母大写 [done]
+ 函数第一参数是接收参数，第二个参数是返回给客户端的参数，必须是指针类型 [注意]
+ 函数还必须有一个返回值error [注意]
*/

type Params struct {
	Length int
	Width  int
}

type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Length * p.Width
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Length + p.Width) * 2
	return nil
}

func Test() {
	r := Rect{}
	p := Params{Length: 5, Width: 4}
	var area, per int
	if err := r.Area(p, &area); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if err := r.Perimeter(p, &per); err != nil {
		log.Fatalln("Error:", err.Error())
	}
	fmt.Printf("Area:%v Perimeter:%v\n", area, per)
}

func main() {
	//1. 注册rpc服务
	r1 := new(Rect) //r1必须是指针，
	//否则会报错；  rpc.Register: type Rect has no exported methods of suitable type (hint: pass a pointer to value of that type)
	//new的， 或者是用&Rect{}
	err := rpc.Register(r1)
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	//2. bind rpc to http
	rpc.HandleHTTP()

	//3. listen http
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln("Error:", err.Error())
	}
}
