package main

import (
	"context"

	pb "github.com/godemos/grpc_hello/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// TLS 认证 + 自定义token认证

// token认证模块
type customCredential struct{}

//实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

// 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}

const (
	Address = "127.0.0.1:8000"
	OpenTLS = true
)

func main() {

	var err error
	var opts []grpc.DialOption

	// 1. 使用TLS证书或者不使用
	if OpenTLS {
		creds, err := credentials.NewClientTLSFromFile("../../keys/example.com.cert", "www.example.com")
		if err != nil {
			grpclog.Fatalln(err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	//2. 必须使用自定义token认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	//create grpc conn with tls/no + token
	conn, err := grpc.Dial(Address, opts...)
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
