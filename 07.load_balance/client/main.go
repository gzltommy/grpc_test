package main

import (
	"context"
	"fmt"
	"github.com/gzltommy/grpc_test/07.load_balance/proto"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"

	// 这里一定要 import。只有 import 进去 grpc-consul-resolver 中 consul 包下的 init 函数才会执行,
	// init 是为了把自己的解析器注册进去;如果不 import 进去就获取不到对应的服务;
	_ "github.com/mbobakov/grpc-consul-resolver"

	"google.golang.org/grpc"
)

const (
	consulAddr = "172.30.134.7:8500"
)

func main() {
	conn, err := grpc.NewClient(
		// consul://192.168.193.128:8500 consul 地址
		// hello-server 拉取的服务名
		// wait=14s 等待时间
		// tag=manual 筛选条件
		// 底层就是利用 grpc-consul-resolver 将参数解析成 HTTP 请求获取对应的服务
		fmt.Sprintf("consul://%s/hello-server?wait=14s&tag=manual", consulAddr),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`), // 轮询
		/*
			gRPC-GO 内置了pick_first，round_robin 两种负载均衡策略。
			pick_first(默认策略):尝试逐个连接客户端地址，如果某一地址连接成功，则将其用于所有 RPC ，如果所有的失败，则报告错误
			round_robin:连接所有地址，并依次向每个可用的后端发送 RPC 请求

		*/
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewHelloClient(conn)

	for i := 0; i < 12; i++ {
		resp, err := client.Hi(context.Background(), &emptypb.Empty{})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Msg)
	}
}
