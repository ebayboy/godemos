package main

import (
	"fmt"
	"sync/atomic"
)

//原子变量的操作
func main() {
	var flag int64 = 123

	//atomic.StoreInt64
	atomic.StoreInt64(&flag, 456)
	fmt.Println("atomic.StoreInt64 flag:", atomic.LoadInt64(&flag))

	//atomic.AddInt64
	flag = 123
	atomic.AddInt64(&flag, 2)
	fmt.Println("atomic.AddInt64:", flag)

	//atomic.SwapInt64
	flag = 123
	old := atomic.SwapInt64(&flag, 2)
	fmt.Println("atomic.SwapInt64:", flag, " old:", old)

	//atomic.CompareAndSwapInt64
	flag = 123
	swapped := atomic.CompareAndSwapInt64(&flag, 123, 456)
	fmt.Println("atomic.CompareAndSwapInt64:", flag, " swapped:", swapped)

	//atomic.CompareAndSwapInt64
	flag = 123
	swapped = atomic.CompareAndSwapInt64(&flag, 789, 456)
	fmt.Println("atomic.CompareAndSwapInt64:", flag, " swapped:", swapped)
}
