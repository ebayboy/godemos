package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type ControllerMapsType map[string]reflect.Value

var ControllerMaps ControllerMapsType

type Routers struct {
}

func (this *Routers) AND(data ...interface{}) {
	rules := data[0].([]bool)
	ret := data[1].(*bool)
	for _, rule := range rules {
		*ret = *ret && rule
	}
	log.Println("AND data:", data, " ret:", ret)
}

func (this *Routers) OR(data ...interface{}) {
	rules := data[0].([]bool)
	ret := data[1].(*bool)

	for _, rule := range rules {
		*ret = *ret || rule
	}
	log.Println("OR data:", data, " ret:", ret)
}

func (this *Routers) NOT(data ...interface{}) {
	rules := data[0].([]bool)
	ret := data[1].(*bool)

	*ret = !rules[0]
	log.Println("NOT data:", data, " ret:", ret)
}

var LogicComputer ControllerMapsType

func init() {
	var ruTest Routers

	LogicComputer = make(ControllerMapsType, 0)
	//创建反射变量，注意这里需要传入ruTest变量的地址；
	//不传入地址就只能反射Routers静态定义的方法
	vf := reflect.ValueOf(&ruTest)
	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()
	fmt.Println("NumMethod:", mNum)
	//遍历路由器的方法，并将其存入控制器映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		fmt.Println("index:", i, " MethodName:", mName)
		LogicComputer[mName] = vf.Method(i) //<<<
	}
}

func main() {
	var ret bool = false
	//演示
	rules := []bool{true, false, true, false}
	//创建带调用方法时需要传入的参数列表
	parms := []reflect.Value{reflect.ValueOf(rules), reflect.ValueOf(&ret)}

	//使用方法名字符串调用指定方法
	f, ok := LogicComputer[strings.ToUpper("and")]
	if ok {
		f.Call(parms)
		log.Println(ret)
	} else {
		log.Println("not exist and func!!!")
	}

	f, ok = LogicComputer[strings.ToUpper("or")]
	if ok {
		f.Call(parms)
		log.Println(ret)
	} else {
		log.Println("or exist and func!!!")
	}

	f, ok = LogicComputer[strings.ToUpper("not")]
	if ok {
		f.Call(parms)
		log.Println(ret)
	} else {
		log.Println("not exist and func!!!")
	}
}
