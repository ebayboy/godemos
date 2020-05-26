/**
 * @Author: fanpengfei
 * @Description:
 * @File:  hello_test.go
 * @Version: 1.0.0
 * @Date: 2020/5/25 13:25
 */

package main

import "testing"

func Test_hello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"test1", args{"fanpf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
