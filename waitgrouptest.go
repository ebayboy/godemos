/**
 * @Author: fanpengfei
 * @Description:
 * @File:  waitgrouptest
 * @Version: 1.0.0
 * @Date: 2020/5/28 14:55
 */

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func test(i int) {
	fmt.Println("test num:", i)
	wg.Done()
}

func main() {

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go test(i)
	}

	wg.Wait()

	fmt.Println("Done")
}
