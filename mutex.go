package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁 + 变量 控制协程退出
func main() {

	rwlock := sync.RWMutex{}
	stopFlag := false
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(idx int) {
			for {
				rwlock.RLock()
				if stopFlag {
					wg.Done()
					rwlock.RUnlock()
					fmt.Println("goroutine exit!", idx)
					return
				}
				rwlock.RUnlock()
				fmt.Println("goroutine alive...", idx)
				time.Sleep(time.Second * 1)
			}
		}(i)
	}

	//fatal error: sync: Unlock of unlocked RWMutex
	wg.Add(1)
	go func() {
		fmt.Println("after 3 seconds cancelFunc!")
		time.Sleep(time.Second * 3)
		rwlock.Lock()
		stopFlag = true
		rwlock.Unlock()
		fmt.Println("Start cancelFunc...")
		wg.Done()
	}()

	fmt.Println("main start to wait exit...")
	wg.Wait()
	fmt.Println("main exit!")
}
