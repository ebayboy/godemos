package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/godemos/grpc_hello/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
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
	creds, err := credentials.NewServerTLSFromFile("../../keys/example.com.cert", "../../keys/example.com.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}

	var opts []grpc.ServerOption

	opts = append(opts, grpc.Creds(creds))

	//append interceptor
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	//create rpc server with TLS
	srv := grpc.NewServer(opts...)

	//registe HelloService
	pb.RegisterHelloServer(srv, HelloService)

	grpclog.Println("Listen on " + Address)

	srv.Serve(listen)

}

func auth(ctx context.Context) error {
	//从ctx中取出appid和key验证

	//type MD map[string][]string
	mdata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "获取鉴权信息失败!")
	}

	var appid, appkey string

	if val, ok := mdata["appid"]; !ok {
		return grpc.Errorf(codes.Unauthenticated, "鉴权获取appid失败!")
	} else {
		appid = val[0]
	}

	if val, ok := mdata["appkey"]; !ok {
		return grpc.Errorf(codes.Unauthenticated, "鉴权获取appkey失败！")
	} else {
		appkey = val[0]
	}

	if appid != "101010" && appkey != "i am key" {
		return grpc.Errorf(codes.Unauthenticated, "鉴权失败!")
	}

	return nil
}

func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := auth(ctx); err != nil {
		grpclog.Errorf(err.Error())
		return nil, err
	}

	grpclog.Printf("auth success!")

	return handler(ctx, req)
}
