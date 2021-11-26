#!/bin/bash


# 生成test文件
# gotests -all fibonacci.go  > fibonacci_test.go
go test fibonacci_test.go  fibonacci.go
