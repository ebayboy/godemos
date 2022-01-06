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
	//解析medata中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "缺少token认证信息")
	}

	var appid string
	var appkey string

	//type MD map[string][]string
	//val type []string
	if val, ok := md["appid"]; ok {
		appid = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "101010" || appkey != "i am key" {
		return nil, grpc.Errorf(codes.Unauthenticated, "Token认证失败!: appid:[%s] != [%s], appkey=[%s] != [%s]",
			appid, "101010", appkey, "i am key")
	}

	resp := new(pb.HelloResponse)
	msg := fmt.Sprintf("hello %s. appkd:%s appkey:%s", in.Name, appid, appkey)
	resp.Message = msg
	grpclog.Println("Send to client:", msg)

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

	//create rpc server with TLS
	srv := grpc.NewServer(grpc.Creds(creds))

	//registe HelloService
	pb.RegisterHelloServer(srv, HelloService)

	grpclog.Println("Listen on " + Address + " With TLS + Token")

	srv.Serve(listen)

}
