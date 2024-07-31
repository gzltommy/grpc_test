// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: 06.ws_grpc/proto/rpc/chat.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Chat_BidStream_FullMethodName = "/rpc.Chat/BidStream"
)

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	BidStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Request, Response], error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) BidStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Request, Response], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], Chat_BidStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Request, Response]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Chat_BidStreamClient = grpc.BidiStreamingClient[Request, Response]

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility.
type ChatServer interface {
	BidStream(grpc.BidiStreamingServer[Request, Response]) error
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatServer struct{}

func (UnimplementedChatServer) BidStream(grpc.BidiStreamingServer[Request, Response]) error {
	return status.Errorf(codes.Unimplemented, "method BidStream not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}
func (UnimplementedChatServer) testEmbeddedByValue()              {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	// If the following call pancis, it indicates UnimplementedChatServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_BidStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).BidStream(&grpc.GenericServerStream[Request, Response]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Chat_BidStreamServer = grpc.BidiStreamingServer[Request, Response]

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BidStream",
			Handler:       _Chat_BidStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "06.ws_grpc/proto/rpc/chat.proto",
}
