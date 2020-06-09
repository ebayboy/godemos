/*
获取MD5以及hashcode
*/

package main

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
)

func main() {
	data := []byte("test")
	fmt.Printf("md5:[%x]\n", md5.Sum(data))
	fmt.Printf("hashcode:%x\n", crc32.ChecksumIEEE(data))
}
