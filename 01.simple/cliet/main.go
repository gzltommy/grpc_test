// main_client.go
package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//创立 grpc 连接
	//grpcConn, err := grpc.Dial("127.0.0.1:6012", grpc.WithInsecure())
	grpcConn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	grpcConn.GetState()

	// 登录
	Login(grpcConn)
	LoginOut(grpcConn)

	// 上传
	//Upload(grpcConn)

	// 下载
	//Download(grpcConn)

	// 聊天
	//BidStream(grpcConn)
}
