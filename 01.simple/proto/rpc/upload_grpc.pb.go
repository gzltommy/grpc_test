// upload.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: 01.simple/proto/rpc/upload.proto

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
	Upload_Upload_FullMethodName = "/rpc.Upload/Upload"
)

// UploadClient is the client API for Upload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 上传服务
// 请求接受一个UploadReq
// 响应回发多条数据（"true" or "false")
type UploadClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadReq, UploadRes], error)
}

type uploadClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadClient(cc grpc.ClientConnInterface) UploadClient {
	return &uploadClient{cc}
}

func (c *uploadClient) Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadReq, UploadRes], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Upload_ServiceDesc.Streams[0], Upload_Upload_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadReq, UploadRes]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Upload_UploadClient = grpc.ClientStreamingClient[UploadReq, UploadRes]

// UploadServer is the server API for Upload service.
// All implementations must embed UnimplementedUploadServer
// for forward compatibility.
//
// 上传服务
// 请求接受一个UploadReq
// 响应回发多条数据（"true" or "false")
type UploadServer interface {
	Upload(grpc.ClientStreamingServer[UploadReq, UploadRes]) error
	mustEmbedUnimplementedUploadServer()
}

// UnimplementedUploadServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUploadServer struct{}

func (UnimplementedUploadServer) Upload(grpc.ClientStreamingServer[UploadReq, UploadRes]) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedUploadServer) mustEmbedUnimplementedUploadServer() {}
func (UnimplementedUploadServer) testEmbeddedByValue()                {}

// UnsafeUploadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServer will
// result in compilation errors.
type UnsafeUploadServer interface {
	mustEmbedUnimplementedUploadServer()
}

func RegisterUploadServer(s grpc.ServiceRegistrar, srv UploadServer) {
	// If the following call pancis, it indicates UnimplementedUploadServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Upload_ServiceDesc, srv)
}

func _Upload_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServer).Upload(&grpc.GenericServerStream[UploadReq, UploadRes]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Upload_UploadServer = grpc.ClientStreamingServer[UploadReq, UploadRes]

// Upload_ServiceDesc is the grpc.ServiceDesc for Upload service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Upload_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Upload",
	HandlerType: (*UploadServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Upload_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "01.simple/proto/rpc/upload.proto",
}
