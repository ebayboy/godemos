package main

import (
	"fmt"
	"log"
	"reflect"
)

type LogicCompute struct {
}

//定义控制器函数Map类型，便于后续快捷使用
type LogicMapsType map[string]reflect.Value

var logicMap LogicMapsType

func (this *LogicCompute) And(data []int) bool {
	log.Println("LogicCompute and...")
	return false
}

func (this *LogicCompute) Or(data []int) bool {
	log.Println("LogicCompute or...")
	return false
}

func (this *LogicCompute) Not(data []int) bool {
	log.Println("LogicCompute not...")
	return false
}

func init() {
	var router LogicCompute
	logicMap := make(LogicMapsType, 0)

	vf := reflect.ValueOf(&router)
	vft := vf.Type()
	//读取方法数量
	mNum := vf.NumMethod()

	log.Println("INFO: LogicCompute init NumMethod:", mNum)
	//遍历路由器的方法，并将其存入控制器映射变量中
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		log.Println("INFO: LogicCompute index:", i, " MethodName:", mName)
		logicMap[mName] = vf.Method(i)
	}
}

func main() {
	//演示
	testStr := "Hello Go"
	//创建带调用方法时需要传入的参数列表
	parms := []reflect.Value{reflect.ValueOf(testStr)}
	//使用方法名字符串调用指定方法
	logicMap["Login"].Call(parms)

	//创建带调用方法时需要传入的参数列表
	parms = []reflect.Value{reflect.ValueOf(&testStr)}
	//使用方法名字符串调用指定方法
	logicMap["ChangeName"].Call(parms)
	//可见，testStr的值已经被修改了
	fmt.Println("testStr:", testStr)
}
