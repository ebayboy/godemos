package main

import (
	"fmt"
	"time"
)

type Element struct {
	Value   float64
	Expired uint64
}

type MinHeap struct {
	Elements []Element
}

// MinHeap构造方法
func NewMinHeap() *MinHeap {
	// TODO: 第一个元素仅用于结束insert中的 for 循环
	return &MinHeap{}
}

// 插入数字,插入数字需要保证堆的性质
func (H *MinHeap) Insert(v Element) {
	H.Elements = append(H.Elements, v)
	i := len(H.Elements) - 1

	// 上浮
	for ; H.Elements[i/2].Value > v.Value; i /= 2 {
		H.Elements[i] = H.Elements[i/2]
	}

	H.Elements[i] = v
}

// 删除并返回最小值
func (H *MinHeap) DeleteMin() (Element, error) {
	if len(H.Elements) < 1 {
		return Element{}, fmt.Errorf("MinHeap is empty")
	}

	minElements := H.Elements[0]
	lastElements := H.Elements[len(H.Elements)-1]

	var i, child int
	for i = 0; i*2 < len(H.Elements); i = child {
		child = i * 2
		if child < len(H.Elements)-1 && H.Elements[child+1].Value < H.Elements[child].Value {
			child++
		}

		// 下滤一层
		if lastElements.Value > H.Elements[child].Value {
			H.Elements[i] = H.Elements[child]
		} else {
			break
		}
	}
	H.Elements[i] = lastElements
	H.Elements = H.Elements[:len(H.Elements)-1]

	return minElements, nil
}

// 堆的大小
func (H *MinHeap) Length() int {
	return len(H.Elements) - 1
}

// 获取最小堆的最小值
func (H *MinHeap) Min() (int, error) {
	if len(H.Elements) > 0 {
		return int(H.Elements[0].Value), nil
	}
	return 0, fmt.Errorf("heap is empty")
}

// MinHeap格式化输出
func (H *MinHeap) String() string {
	return fmt.Sprintf("%v", H.Elements[1:])
}

func main() {
	nm := NewMinHeap()

	tmw := time.Now().Unix()
	vals := []float64{1.11, 50.11, 2.11, 100.11, 10.11}
	for _, v := range vals {
		nm.Insert(Element{Expired: uint64(tmw), Value: v})
	}

	v, _ := nm.Min()
	fmt.Println("min:", v)

	nm.DeleteMin()
	v, _ = nm.Min()
	fmt.Println("min:", v)
}
