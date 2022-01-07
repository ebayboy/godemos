#!/bin/bash


# 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

# 编译hello_http.proto
protoc -I . --go_out=plugins=grpc,Mannotations.proto=./google/api:. hello_http/*.proto

# 编译hello_http.proto gateway
protoc --grpc-gateway_out=logtostderr=true:. hello_http/hello_http.proto
