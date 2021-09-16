package main

import "fmt"

//接口的多态性
//实现uniqFilter && multiFileter
//uniqFilter: remove duplicated number.(去重)
//multiFilter: only keep numbers which are multiples(保留重复数)

type Filter interface {
	/* TODO: add methods */
	About() string
	Process([]int) []int
}

//UniqueFilter
type UniqueFilter struct {
}

func (u *UniqueFilter) About() string {
	return "remove duplicated number"
}

func (u *UniqueFilter) Process(in []int) []int {
	var outs []int

	for _, v := range in {
		found := false
		for _, o := range outs {
			if o == v {
				found = true
				break
			}
		}
		if !found {
			outs = append(outs, v)
		}
	}

	return outs
}

//MultiFilter
type MultiFilter struct {
}

func (m *MultiFilter) About() string {
	return "only keep numbers which are multiples"
}

func (m *MultiFilter) Process(in []int) []int {

	var outs []int

	tmp := make(map[int]int)

	for _, v := range in {

		if _, exist := tmp[v]; !exist {
			tmp[v] = 1
		} else {
			tmp[v] = tmp[v] + 1
		}
	}

	for k, v := range tmp {
		if v > 1 {
			outs = append(outs, k)
		}
	}

	return outs
}

//接口作为函数参数, 可以传递实现该接口的任何结构体
func testFilter(f Filter, in []int) (out []int) {
	fmt.Println("About:", f.About())
	out = f.Process(in)
	return out
}

//main
func main() {
	fmt.Println("vim-go")

	in := []int{1, 2, 2, 3, 5, 7, 7, 9}

	//test uniqFilter
	var uq UniqueFilter
	out := testFilter(&uq, in)
	fmt.Println("UniqueFilter out:", out)

	//testMultiFilter
	var mul MultiFilter
	out = testFilter(&mul, in)
	fmt.Println("mul:", out)
}
