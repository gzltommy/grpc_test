package main

import (
	"github.com/gzltommy/grpc_test/01.simple/proto/rpc"
	"io"
	"log"
	"strconv"
)

// ChatServer 服务端
type ChatServer struct {
	rpc.UnimplementedChatServer
}

// BidStream 实现了 ChatServer 接口中定义的 BidStream 方法
func (s *ChatServer) BidStream(stream rpc.Chat_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("收到客户端通过 context 发出的终止信号")
			return ctx.Err()
		default:
			// 接收从客户端发来的消息
			request, err := stream.Recv()
			if err == io.EOF {
				log.Println("客户端发送的数据流结束")
				return nil
			}
			if err != nil {
				log.Println("接收数据出错:", err)
				return err
			}

			// 如果接收正常，则根据接收到的 字符串 执行相应的指令
			switch request.Input {
			case "结束对话\n":
				log.Println("收到'结束对话'指令")
				if err := stream.Send(&rpc.Response{Output: "收到结束指令"}); err != nil {
					return err
				}
				// 收到结束指令时，通过 return nil 终止双向数据流
				return nil

			case "返回数据流\n":
				log.Println("收到'返回数据流'指令")
				// 收到 收到'返回数据流'指令， 连续返回 10 条数据
				for i := 0; i < 10; i++ {
					if err := stream.Send(&rpc.Response{Output: "数据流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}

			default:
				// 缺省情况下， 返回 '服务端返回: ' + 输入信息
				log.Printf("[收到消息]: %s", request.Input)
				if err := stream.Send(&rpc.Response{Output: "服务端返回: " + request.Input}); err != nil {
					return err
				}
			}
		}
	}
}
