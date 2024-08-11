package service

import (
	"context"
	"github.com/gzltommy/grpc_test/07.load_balance/proto"
	"github.com/prometheus/common/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Hello struct {
	proto.UnimplementedHelloServer
}

func (s Hello) Hi(ctx context.Context, empty *emptypb.Empty) (*proto.HiResponse, error) {
	log.Info("收到一个请求")
	return &proto.HiResponse{Msg: "test"}, nil
}
