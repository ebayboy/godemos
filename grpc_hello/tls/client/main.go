package main

import (
	"context"

	pb "github.com/godemos/grpc_hello/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// 实现rpc hello方法服务接口

const (
	Address = "127.0.0.1:8000"
)

func main() {

	//CN = go-grpc-example
	creds, err := credentials.NewClientTLSFromFile("../../keys/example.com.cert", "www.example.com")
	if err != nil {
		grpclog.Fatalln(err)
	}

	//create grpc conn
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	//使用conn创建客户端
	c := pb.NewHelloClient(conn)

	//调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	grpclog.Println(res.Message)
}
