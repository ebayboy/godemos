package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/godemos/grpc_hello/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// 实现rpc hello方法服务接口

const (
	Address = "127.0.0.1:8000"
)

type helloService struct{}

var HelloService = helloService{}

//模块helloService实现SayHello方法
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("hello %s.", in.Name)
	return resp, nil
}

func main() {
	//with tcp protocol
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// ## TLS认证 ##
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}

	//create rpc server with TLS
	srv := grpc.NewServer(grpc.Creds(creds))

	//registe HelloService
	pb.RegisterHelloServer(srv, HelloService)

	grpclog.Println("Listen on " + Address)

	srv.Serve(listen)

}
