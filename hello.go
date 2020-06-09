package main

import (
	"fmt"
	"strings"
)

func hello(name string) {
	fmt.Println("hello", name)
}

func main() {
	str := "GT(vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m, 0.5)"
	str = strings.ReplaceAll(str, "vGID_vSNM_vIP_count_404_5m", "%v")
	str = strings.ReplaceAll(str, "vGID_vSNM_vIP_count_5m", "%v")

	fmt.Println("str:", str)
	newStr := fmt.Sprintf(str, 100, 200)
	fmt.Println("newStr", newStr)

}
