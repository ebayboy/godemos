#!/bin/bash

go test -v -bench="BenchmarkSum$"  --run=none main_test.go  main.go

