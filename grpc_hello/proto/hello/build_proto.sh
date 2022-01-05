protoc -I . --go_out=plugins=grpc:. ./hello.proto
cp hello/hello.pb.go  . 
