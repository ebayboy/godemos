package main

import (
	"context"

	pb "github.com/godemos/grpc-gateway/proto/hello_http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// 实现rpc hello方法服务接口

const (
	Address = "127.0.0.1:8000"
)

func main() {

	//create grpc conn
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	//使用conn创建客户端
	c := pb.NewHelloHTTPClient(conn)

	//调用方法
	req := &pb.HelloHTTPRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	grpclog.Println(res.Message)
}
