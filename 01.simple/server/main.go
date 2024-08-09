// main_server.go
package main

import (
	"github.com/gzltommy/grpc_test/01.simple/proto/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 构建一个新的服务端对象
	s := grpc.NewServer()

	// 向这个服务端对象注册服务
	//rpc.RegisterLoginServer(s, &LoginServer{})
	//rpc.RegisterUploadServer(s, &UploadServer{})
	//rpc.RegisterDownloadServer(s, &DownloadServer{})
	rpc.RegisterChatServer(s, &ChatServer{})

	// 注册服务端反射服务
	reflection.Register(s)

	// 启动服务,阻塞，这里一般单独启动一个 goroutine
	err = s.Serve(lis)

	//可配合ctx实现服务端的动态终止
	//s.Stop()
}
