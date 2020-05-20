package main

import "fmt"

func main() {

	//生命切片
	s := make([]string, 5)

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	//初始化
	for n := 0; n < len(s); n++ {
		s[n] = fmt.Sprintf("name_%d", n)
	}

	//ffr 打印 key, value
	for i, i2 := range s {
		fmt.Printf("i:%d i2:%s\n", i, i2)
	}

}
