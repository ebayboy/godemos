package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

//使用普通的map：
//因为是非线程安全的map： 多协程写会报错error: concurrent map writes错误。
func testMap() {
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(idx int) {
			key := strconv.Itoa(idx)
			set(key, idx)
			fmt.Printf("key=%v, v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}

	wg.Wait()
}

var sm sync.Map

//使用sync.Map, 并发读写协程安全
func testSyncMap() {
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(idx int) {
			key := strconv.Itoa(idx)
			sm.Store(key, idx)
			val, ok := sm.Load(key)
			if ok {
				fmt.Printf("key=%v, v:=%v\n", key, val)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
func main() {
	//testMap()

	testSyncMap()
}
