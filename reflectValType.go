package main

import (
	"fmt"
	"reflect"
)

func registerModel(model interface{}) {
	val := reflect.ValueOf(model)
	typ := reflect.TypeOf(model)
	fmt.Println("val:", val) //hello
	fmt.Println("typ:", typ) //type

}

type Student struct {
	Name string
	Age  int
}

func main() {
	s := Student{Name: "fanpf", Age: 32}
	registerModel(s)

}
