package main

import (
	"fmt"
	pb "github.com/gzltommy/grpc_test/02.auth/proto/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials" // 引入 grpc 认证包
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata" // 引入 grpc meta 包
	"google.golang.org/grpc/status"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8080"
)

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}

	// 实例化 grpc Server, 并开启 TLS 认证
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册 HelloService
	pb.RegisterHelloServer(s, helloService{})

	fmt.Println("Listen on " + Address + " with TLS + Token")

	// 阻塞
	err = s.Serve(listen)
	_ = err
}

// 定义helloService并实现约定的接口
type helloService struct {
	pb.UnimplementedHelloServer
}

// SayHello 实现 Hello 服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 解析 metada 中的信息并验证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}

	var (
		appid  string
		appkey string
	)

	if val, ok := md["appid"]; ok {
		appid = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "101010" || appkey != "i am key" {
		return nil, status.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}

	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", in.Name, appid, appkey)

	return resp, nil
}
