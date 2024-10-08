// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: 06.ws_grpc/proto/rpc/chat.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求数据 Request格式定义
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file__06_ws_grpc_proto_rpc_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file__06_ws_grpc_proto_rpc_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file__06_ws_grpc_proto_rpc_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

// 响应数据Response格式定义
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file__06_ws_grpc_proto_rpc_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file__06_ws_grpc_proto_rpc_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file__06_ws_grpc_proto_rpc_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File__06_ws_grpc_proto_rpc_chat_proto protoreflect.FileDescriptor

var file__06_ws_grpc_proto_rpc_chat_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x30, 0x36, 0x2e, 0x77, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x1f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x36, 0x0a, 0x04, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x42, 0x69, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x0c, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x7a, 0x6c, 0x74, 0x6f, 0x6d, 0x6d, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x30, 0x36, 0x2e, 0x77, 0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file__06_ws_grpc_proto_rpc_chat_proto_rawDescOnce sync.Once
	file__06_ws_grpc_proto_rpc_chat_proto_rawDescData = file__06_ws_grpc_proto_rpc_chat_proto_rawDesc
)

func file__06_ws_grpc_proto_rpc_chat_proto_rawDescGZIP() []byte {
	file__06_ws_grpc_proto_rpc_chat_proto_rawDescOnce.Do(func() {
		file__06_ws_grpc_proto_rpc_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file__06_ws_grpc_proto_rpc_chat_proto_rawDescData)
	})
	return file__06_ws_grpc_proto_rpc_chat_proto_rawDescData
}

var file__06_ws_grpc_proto_rpc_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file__06_ws_grpc_proto_rpc_chat_proto_goTypes = []any{
	(*Request)(nil),  // 0: rpc.Request
	(*Response)(nil), // 1: rpc.Response
}
var file__06_ws_grpc_proto_rpc_chat_proto_depIdxs = []int32{
	0, // 0: rpc.Chat.BidStream:input_type -> rpc.Request
	1, // 1: rpc.Chat.BidStream:output_type -> rpc.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file__06_ws_grpc_proto_rpc_chat_proto_init() }
func file__06_ws_grpc_proto_rpc_chat_proto_init() {
	if File__06_ws_grpc_proto_rpc_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__06_ws_grpc_proto_rpc_chat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file__06_ws_grpc_proto_rpc_chat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file__06_ws_grpc_proto_rpc_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file__06_ws_grpc_proto_rpc_chat_proto_goTypes,
		DependencyIndexes: file__06_ws_grpc_proto_rpc_chat_proto_depIdxs,
		MessageInfos:      file__06_ws_grpc_proto_rpc_chat_proto_msgTypes,
	}.Build()
	File__06_ws_grpc_proto_rpc_chat_proto = out.File
	file__06_ws_grpc_proto_rpc_chat_proto_rawDesc = nil
	file__06_ws_grpc_proto_rpc_chat_proto_goTypes = nil
	file__06_ws_grpc_proto_rpc_chat_proto_depIdxs = nil
}
