package main

import (
	"context"
	"github.com/gzltommy/grpc_test/01.simple/proto/rpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Download(grpcConn *grpc.ClientConn) {
	//通过grpc连接创建一个客户端实例对象
	client := rpc.NewDownloadClient(grpcConn)

	//设置ctx超时（根据情况设定）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//和简单rpc不同，此时获得的不是res，而是一个client的对象，通过这个连接对象去读取数据
	downloadClient, err := client.Download(ctx, &rpc.DownloadReq{
		Path:   "../test_grpc.txt",
		Offset: 0,
		Size:   64 * 1024,
	})
	if err != nil {
		log.Fatalln(err)
	}

	//循环处理数据，当监测到读取完成后退出
	for {
		res, err := downloadClient.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("get a date package~ offset:%v, size:%v\n", res.Offset, res.Size)
		if res.Size+res.Offset >= 64*1024 {
			break
		}
	}

	log.Println("download over~")
}
