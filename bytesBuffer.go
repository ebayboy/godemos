/**
 * @Author: fanpengfei
 * @Description:
 * @File:  bytesBuffer
 * @Version: 1.0.0
 * @Date: 2020/5/28 15:36
 */

package main

import (
	"bytes"
	"fmt"
)

//缓冲区 bytes.buffer
func main() {
	//声明一个变量
	var b bytes.Buffer

	//写入字符串
	b.WriteString("hello world")

	//将buffer内容字符串化
	fmt.Println(b.String())
}
