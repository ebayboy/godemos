package main

import (
	"encoding/json"
	"fmt"
)

type addr struct {
	Province string `json:"province"` //转化后对应字段的json名为province
	City     string //未转化， 输出为City
}
type stu struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr addr   `json:"addr"`
}

func main() {
	var xm = stu{Name: "xiaoming", Age: 18, Addr: addr{Province: "Hunan", City: "ChangSha"}}

	//struct -> json
	fmt.Println("struct -> json")
	js, err := json.Marshal(xm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(js))

	//json -> struct
	fmt.Println("json -> struct")
	var xxm stu
	err = json.Unmarshal(js, &xxm)
	fmt.Println(xxm)
}
