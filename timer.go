package main

import (
	"fmt"
	"time"
)

//Timer：时间到了，执行只执行1次
//Ticker：时间到了，多次执行

func main() {
	// Timer：时间到了，执行只执行1次

	// 1. timer基本使用
	/*
		timer1 := time.NewTimer(time.Second * 2)
		t1 := time.Now()
		fmt.Println("t1:", t1)
		t2 := <-timer1.C
		fmt.Println("t2:", t2)
	*/

	// 2. 验证timer只能响应一次
	timer2 := time.NewTimer(time.Second * 1)
	fmt.Println("时间到:", <-timer2.C)

	// 3. 延迟功能
	timer3 := time.NewTimer(time.Second * 2)
	fmt.Println("2秒到：", <-timer3.C)

	fmt.Println("After 2秒到:", <-time.After(time.Second*2))

	//4. 停止定时器
	timer4 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer4.C
		fmt.Println("timer4 定时器执行了") //因为main关闭了timer4, 此处没有执行
	}()

	b := timer4.Stop()
	if b {
		fmt.Println("timer4已经关闭")
	}

	//5. 重置定时器
	timer5 := time.NewTimer(time.Second * 5)
	timer5.Reset(time.Second * 1)
	fmt.Println("now:", time.Now())
	fmt.Println("now:", <-timer5.C)
}
