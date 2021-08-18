package main

import (
	"testing"
)

//单元测试
func TestSum(t *testing.T) {
	//准备参数
	param := 10
	//执行函数
	ret := Sum(param)
	//判断结果是否符合预期
	if ret != 9 {
		t.Error("Sum result failed")
	}
}

//性能测试
func BenchmarkSum(b *testing.B) {
	//准备参数
	param := 10
	//执行函数
	for i := 0; i < b.N; i++ {
		Sum(param)
	}
}
