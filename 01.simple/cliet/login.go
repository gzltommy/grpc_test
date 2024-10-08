package main

import (
	"context"
	"github.com/gzltommy/grpc_test/01.simple/proto/rpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Login(grpcConn *grpc.ClientConn) {
	// 通过 grpc 连接创建一个客户端实例对象
	client := rpc.NewLoginClient(grpcConn)

	//设置ctx超时（根据情况设定）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//通过client客户端对象，调用Login函数
	res, err := client.Login(ctx, &rpc.LoginReq{
		Username: "root",
		Password: "123456",
	})
	if err != nil {
		log.Fatalln(err)
	}

	//输出登陆结果
	log.Println("the login answer is", res.Msg)

	// 返回结果后还是可以继续发送消息的，说明面向的是长链接
	{
		res, err := client.Login(ctx, &rpc.LoginReq{
			Username: "root",
			Password: "123456",
		})
		if err != nil {
			log.Fatalln(err)
		}

		//输出登陆结果
		log.Println("the login answer is", res.Msg)
	}
}

func LoginOut(grpcConn *grpc.ClientConn) {
	// 通过 grpc 连接创建一个客户端实例对象
	client := rpc.NewLoginClient(grpcConn)

	//设置ctx超时（根据情况设定）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//通过client客户端对象，调用Login函数
	res, err := client.Logout(ctx, &rpc.LoginReq{
		Username: "root",
		Password: "123456",
	})
	if err != nil {
		log.Fatalln(err)
	}

	//输出登出结果
	log.Println("the login answer is", res.Msg)
}
