package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//比较互斥锁和原子变量
// 1. 原子变量性能大于互斥锁，
// 2. 互斥锁比原子变量应用场景广（锁定业务块）

//互斥锁方案
var x int64 = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 50000000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func testLock() {
	wg.Add(2)
	start := time.Now()
	go add()
	go add()
	wg.Wait()
	fmt.Println("time since:", time.Since(start), " x=", x)
}

//原子变量方案
func testAtomic() {
	wg.Add(2)

	go func() {
		for i := 0; i < 50000000; i++ {
			atomic.AddInt64(&x, 1)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 50000000; i++ {
			atomic.AddInt64(&x, 1)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("x = ", x)
}

func main() {

	s := time.Now()
	testLock() //2 sec
	fmt.Println("mutex cost:", time.Since(s))

	s2 := time.Now()
	x = 0
	testAtomic()
	fmt.Println("atomic cost:", time.Since(s2))
}
