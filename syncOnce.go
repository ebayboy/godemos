package main

import (
	"fmt"
	"sync"
)

// sync.Once: 协程安全的只初始化一次

var cfg map[string]string
var loadCfgOnce sync.Once // goroutine safe

func loadCfg() {
	cfg = make(map[string]string)

	cfg["input"] = "input1"
	cfg["stat"] = "stat1"
	cfg["output"] = "output1"
}

func getNode(name string) string {
	/*
		loadCfg() // is not goroutine safe
	*/

	loadCfgOnce.Do(loadCfg) // goroutine safe

	res, exist := cfg[name]
	if !exist {
		return ""
	}

	return res
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fmt.Printf("node[%d]:%s\n", idx, getNode("input"))
		}(i)
	}

	wg.Wait()
}
