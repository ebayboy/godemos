/**
 * @Author: fanpengfei
 * @Description:
 * @File:  mapFuncTest
 * @Version: 1.0.0
 * @Date: 2020/6/17 10:35
 */

package LogicCompute

import (
	"fmt"
	"reflect"
)

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
