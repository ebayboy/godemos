package main

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	gw "github.com/godemos/grpc-gateway/proto/hello_http"
)

func main() {
	ctx := context.Background()
	ctx, cancle := context.WithCancel(ctx)
	defer cancle()

	//grpc服务地址
	endpoint := "127.0.0.1:8000"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	//http转grpc
	err := gw.RegisterHelloHTTPHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		grpclog.Fatalf("RegisterHelloHTTPHandlerFromEndpoint error:%v", err.Error())
	}

	grpclog.Println("HTTP Listen on 8080")
	http.ListenAndServe(":8080", mux)
}
