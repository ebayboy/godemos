package main

import "fmt"

/* 接口开始 ============================================ */
type Phone interface {
	call()
}

/* 接口结束 ===========================================  */

/* ========================== 接口实现1  ==============  */
type NokiaPhone struct {
}

// func
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, i cal call you!")
}

/* =========================接口实现2 =================== */
type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println(("I am IPhone, I can call you!"))
}

func main() {
	//定义一个接口
	var phone Phone

	//接口实例化1
	phone = new(NokiaPhone)
	phone.call()

	//接口实例化2
	phone = new(IPhone)
	phone.call()

	//空接口
	var nullFace interface{} = 1
	fmt.Println("nullFace=", nullFace)
}
