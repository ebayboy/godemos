package main

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	pb "github.com/godemos/grpc-gateway/proto/hello_http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// 定义helloHTTPService并实现约定的接口
type helloHTTPService struct{}

// HelloHTTPService Hello HTTP服务
var HelloHTTPService = helloHTTPService{}

// 实现proto的SayHello服务接口
func (h helloHTTPService) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	resp := new(pb.HelloHTTPResponse)
	resp.Message = "Hello " + in.Name + "."

	return resp, nil
}

func main() {

	endpoint := "127.0.0.1:8000"
	conn, err := net.Listen("tcp", endpoint)
	if err != nil {
		grpclog.Fatalf("TCP Listen err:%v\n", err)
	}

	// grpc tls server
	creds, err := credentials.NewServerTLSFromFile("../../keys/example.com.cert", "../../keys/example.com.key")
	if err != nil {
		grpclog.Fatalf("Failed to create server TLS credentials %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	//注册protobuf 服务
	pb.RegisterHelloHTTPServer(grpcServer, HelloHTTPService)

	// gw server
	// ./proto/hello_http/hello_http.pb.gw.go
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile("../../keys/example.com.cert", "www.example.com")
	if err != nil {
		grpclog.Fatalf("Failed to create client TLS credentials %v", err)
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)} //Dial opt with TLS
	gwmux := runtime.NewServeMux()
	if err = pb.RegisterHelloHTTPHandlerFromEndpoint(ctx, gwmux, endpoint, dopts); err != nil {
		grpclog.Fatalf("Failed to register gw server: %v\n", err)
	}

	// 根服务路由
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	grpclog.Infof("gRPC and https listen on: %s\n", endpoint)
	srv := &http.Server{
		Addr:      endpoint,
		Handler:   grpcHandlerFunc(grpcServer, mux),
		TLSConfig: getTLSConfig(),
	}
	if err = srv.Serve(tls.NewListener(conn, srv.TLSConfig)); err != nil {
		grpclog.Fatal("ListenAndServe: ", err)
	}

	return
}

func getTLSConfig() *tls.Config {
	cert, _ := ioutil.ReadFile("../../keys/example.com.cert")
	key, _ := ioutil.ReadFile("../../keys/example.com.key")
	var demoKeyPair *tls.Certificate
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		grpclog.Fatalf("TLS KeyPair err: %v\n", err)
	}
	demoKeyPair = &pair
	return &tls.Config{
		Certificates: []tls.Certificate{*demoKeyPair},
		NextProtos:   []string{http2.NextProtoTLS}, // HTTP2 TLS支持
	}
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	if otherHandler == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			grpcServer.ServeHTTP(w, r)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
