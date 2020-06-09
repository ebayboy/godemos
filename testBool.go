/**
 * @Author: fanpengfei
 * @Description:
 * @File:  testBool
 * @Version: 1.0.0
 * @Date: 2020/5/26 15:13
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func splitSep(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/' || r == '+' || r == '(' || r == ')' || r == '%' || r == '[' || r == ']' || r == '^' || r == ' ' || r == '\t' || r == '<' || r == '>' || r == '=' || r == ','
}

//字符切分
func Split_string(s string) []string {

	a := strings.FieldsFunc(s, splitSep)
	return a
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

/*
i: 0 s: vGID_vSNM_vIP_count_404_5m
i: 1 s: vGID_vSNM_vIP_count_5m
i: 3 s: 0.5
*/
func main() {
	formula := " GD(vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m, 0.5)"
	arr := Split_string(formula)

	for _, s := range arr {
		if IsNum(s) {
			fmt.Println("IsNumber:", s)
		}
	}

	for i, s := range arr {
		fmt.Printf("i:[%v] s:[%v]\n", i, s)
	}
}
