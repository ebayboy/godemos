/**
 * @Author: fanpengfei
 * @Description:
 * @File:  test
 * @Version: 1.0.0
 * @Date: 2020/5/28 11:18
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	ts := "1585944382"
	t, err := strconv.ParseFloat(ts, 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(int64(t) - int64(t)%60)
	fmt.Println("int64(t)", int64(t))
	fmt.Println("int64(t)%60", int64(t)%60)
	//int64(t) - int64(t)%self.conf.Wind
	//int64(t)%self.conf.Wind
	//日志的时间和60求余=剩余的秒数
	//int64(t) - 剩余的描述 = 能与60整除的秒数

}
