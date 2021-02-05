package main

import "fmt"

/**
...的用法
**/

//1. 变长参数
func Sum(nums ...int) int {
	res := 0

	for _, n := range nums {
		res += n
	}

	return res
}

func main() {
	fmt.Println("sum:", Sum(1, 2, 3))

	//2. 展开切片, 注意不能展开数组
	nums := []int{1, 2, 3} //此处是切片
	fmt.Println("sum2:", Sum(nums...))

	//3. 代表数组元素的个数
	array := [...]int{4, 5, 6, 7}
	fmt.Println("array:", array)
}
