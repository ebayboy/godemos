package main

import (
	"context"
	"time"

	pb "github.com/godemos/grpc_hello/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// 拦截器实现 inteprator: TLS 认证 + 自定义token认证

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

//客户端使用的都是With开头的函数： 例如WithTransportCredentials, WithInsecure, WithPerRPCCredentials,WithUnaryInterceptor
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

	//3. 添加客户端拦截器inteprator
	opts = append(opts, grpc.WithUnaryInterceptor(inteceptor))

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

//客户端拦截器
func inteceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()

	// ## invoker: 先调用拦截器unaryInt，
	// 然后在1里面调用invoke处理真正的一次请求过程
	// ## var invoker grpc.UnaryInvoker :
	grpclog.Printf("invoker start...")

	err := invoker(ctx, method, req, reply, cc, opts...)

	grpclog.Printf("invokker end: method:%s req:%v rep:%v duration=%s error:%v\n", method, req, reply, time.Since(start), err)

	return err
}
