/**
 * @Author: fanpengfei
 * @Description:
 * @File:  test
 * @Version: 1.0.0
 * @Date: 2020/5/28 11:18
 */

package main

import "log"

//值传递
func test1(data interface{}) {
	m := data.(string)
	log.Println("test1:", &m)
}

//引用传递
func test3(data interface{}) {
	m := data.(*string)
	log.Println("test3:", m)
	log.Println("test3 value:", *m)
}

/*
2020/06/15 18:01:05 main: 0xc0000301f0
2020/06/15 18:01:05 test1: 0xc000030210
2020/06/15 18:01:05 test3: 0xc0000301f0
*/
func main() {
	m := "hello world"
	log.Println("main:", &m)
	test1(m)
	test3(&m)
}
