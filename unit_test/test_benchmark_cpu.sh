#!/bin/bash

go test -v -bench="BenchmarkSum$" --run=none -cpuprofile cpu.out main_test.go main.go
go tool pprof cpu.out

#start web port for check
#go tool pprof -http=192.168.56.101:9090 cpu.out
