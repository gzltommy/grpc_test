package main

import (
	"context"
	"fmt"
	"github.com/gzltommy/grpc_test/load_balance/proto"
	"github.com/hashicorp/consul/api"
	"github.com/prometheus/common/log"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"strings"
)

const (
	serverHost = "172.30.134.7"
	consulAddr = "172.30.134.7:8500"
)

type Hello struct {
	proto.UnimplementedHelloServer
}

func (s Hello) Hi(ctx context.Context, empty *emptypb.Empty) (*proto.HiResponse, error) {
	log.Info("收到一个请求")
	return &proto.HiResponse{Msg: "test"}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterHelloServer(server, &Hello{})

	port := GenFreePort()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("监听端口:%d失败: %s", port, err.Error())
	}

	config := api.DefaultConfig()
	config.Address = consulAddr

	consulClient, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("连接consul失败: %s", err.Error())
	}

	// grpc 注册服务的健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 每个服务的 ID 必须不同;这里使用 uuid;
	// Name 相同 ID 不同 consul 会认为是两个实例;
	// Name 和 ID 都相同 consul 会认为是一个实例会出现覆盖
	registration := &api.AgentServiceRegistration{
		Address: serverHost,
		Port:    port,
		ID:      fmt.Sprintf("%s", strings.ReplaceAll(uuid.NewV4().String(), "-", "")),
		Name:    "hello-server",
		Tags:    []string{"manual"},
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",                                   // 指定运行此检查的频率
			Timeout:                        "5s",                                   // 超时时间
			GRPC:                           fmt.Sprintf("%s:%d", serverHost, port), // 健康检查HTTP请求
			DeregisterCriticalServiceAfter: "30s",                                  // 触发注销的时间
		},
	}
	err = consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("注册服务失败: %s", err.Error())
	}

	fmt.Printf("服务启动成功;PORT:%d\n", port)
	err = server.Serve(listen)
}

// GenFreePort 获取一个空闲的端口;端口避免写死,因为要启动多个实例,测试负载均衡
func GenFreePort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port
}
