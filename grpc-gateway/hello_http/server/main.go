package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	pb "github.com/godemos/grpc-gateway/proto/hello"
	"golang.org/x/net/trace"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var Address string = "127.0.0.1:8000"

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func init() {
	//##  必须开启后， 访问debug/[events|requests]才会有信息
	grpc.EnableTracing = true
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	// 开启trace
	go startTrace()

	grpclog.Println("Listen on " + Address)
	s.Serve(listen)
}

func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	go http.ListenAndServe(":50051", nil)
	grpclog.Println("Trace listen on 50051")
}
