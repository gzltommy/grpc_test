// upload.go

package main

import (
	"github.com/gzltommy/grpc_test/01.simple/proto/rpc"
	"io"
	"log"
)

type UploadServer struct {
	rpc.UnimplementedUploadServer
}

func (*UploadServer) Upload(uploadServer rpc.Upload_UploadServer) error {
	for {
		//循环接受客户端传的流数据
		recv, err := uploadServer.Recv()

		//检测到EOF（客户端调用close）
		if err == io.EOF {
			//发送res
			err := uploadServer.SendAndClose(&rpc.UploadRes{Msg: "finish"})
			if err != nil {
				return err
			}
			return nil
		} else if err != nil {
			return err
		}

		// 输出接收的流数据
		log.Printf("get a upload data package~ offset:%v, size:%v\n", recv.Offset, recv.Size)
	}
}
