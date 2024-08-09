package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gzltommy/grpc_test/01.simple/proto/rpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
	"time"
)

// BidStream 双向流
func BidStream(grpcConn *grpc.ClientConn) {
	//通过grpc连接创建一个客户端实例对象
	client := rpc.NewChatClient(grpcConn)

	//设置ctx超时（根据情况设定）
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	//ctx := context.Background()

	// 创建双向数据流
	stream, err := client.BidStream(ctx)
	if err != nil {
		log.Printf("创建数据流失败: [%v]\n", err)
	}

	// 启动一个 goroutine 接收命令行输入的指令
	go func() {
		log.Println("请输入消息...")
		reader := bufio.NewReader(os.Stdin)
		for {
			// 获取 命令行输入的字符串， 以回车 \n 作为结束标志
			stdInput, _ := reader.ReadString('\n')

			// 向服务端发送 指令
			if err := stream.Send(&rpc.Request{Input: stdInput}); err != nil {
				return
			}
		}
	}()

	for {
		// 接收从 服务端返回的数据流
		response, err := stream.Recv()
		if err == io.EOF {
			log.Println("⚠️ 收到服务端的结束信号")
			break //如果收到结束信号，则退出“接收循环”，结束客户端程序
		}

		if errors.As(err, status.Error) {
			fmt.Println("=======超时拉=======")
		}

		if err != nil {
			// TODO: 处理接收错误
			log.Printf("接收数据出错:%T,%v", err, err)
		}

		// 没有错误的情况下，打印来自服务端的消息
		log.Printf("[客户端收到]: %s", response.Output)
	}
}
