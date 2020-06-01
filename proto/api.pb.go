// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.12.0
// source: api.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x33, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xef, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x37,
	0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x40, 0x0a, 0x18, 0x42, 0x69, 0x44, 0x72, 0x69,
	0x63, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: proto.Message
}
var file_api_proto_depIdxs = []int32{
	0, // 0: proto.SendMessage.SayHello:input_type -> proto.Message
	0, // 1: proto.SendMessage.ServerStreamHello:input_type -> proto.Message
	0, // 2: proto.SendMessage.ClientStreamHello:input_type -> proto.Message
	0, // 3: proto.SendMessage.BiDricetionalStreamHello:input_type -> proto.Message
	0, // 4: proto.SendMessage.SayHello:output_type -> proto.Message
	0, // 5: proto.SendMessage.ServerStreamHello:output_type -> proto.Message
	0, // 6: proto.SendMessage.ClientStreamHello:output_type -> proto.Message
	0, // 7: proto.SendMessage.BiDricetionalStreamHello:output_type -> proto.Message
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SendMessageClient is the client API for SendMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SendMessageClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	ServerStreamHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (SendMessage_ServerStreamHelloClient, error)
	ClientStreamHello(ctx context.Context, opts ...grpc.CallOption) (SendMessage_ClientStreamHelloClient, error)
	BiDricetionalStreamHello(ctx context.Context, opts ...grpc.CallOption) (SendMessage_BiDricetionalStreamHelloClient, error)
}

type sendMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewSendMessageClient(cc grpc.ClientConnInterface) SendMessageClient {
	return &sendMessageClient{cc}
}

func (c *sendMessageClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/proto.SendMessage/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendMessageClient) ServerStreamHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (SendMessage_ServerStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SendMessage_serviceDesc.Streams[0], "/proto.SendMessage/ServerStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &sendMessageServerStreamHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SendMessage_ServerStreamHelloClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type sendMessageServerStreamHelloClient struct {
	grpc.ClientStream
}

func (x *sendMessageServerStreamHelloClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sendMessageClient) ClientStreamHello(ctx context.Context, opts ...grpc.CallOption) (SendMessage_ClientStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SendMessage_serviceDesc.Streams[1], "/proto.SendMessage/ClientStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &sendMessageClientStreamHelloClient{stream}
	return x, nil
}

type SendMessage_ClientStreamHelloClient interface {
	Send(*Message) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type sendMessageClientStreamHelloClient struct {
	grpc.ClientStream
}

func (x *sendMessageClientStreamHelloClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sendMessageClientStreamHelloClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sendMessageClient) BiDricetionalStreamHello(ctx context.Context, opts ...grpc.CallOption) (SendMessage_BiDricetionalStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SendMessage_serviceDesc.Streams[2], "/proto.SendMessage/BiDricetionalStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &sendMessageBiDricetionalStreamHelloClient{stream}
	return x, nil
}

type SendMessage_BiDricetionalStreamHelloClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type sendMessageBiDricetionalStreamHelloClient struct {
	grpc.ClientStream
}

func (x *sendMessageBiDricetionalStreamHelloClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sendMessageBiDricetionalStreamHelloClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SendMessageServer is the server API for SendMessage service.
type SendMessageServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	ServerStreamHello(*Message, SendMessage_ServerStreamHelloServer) error
	ClientStreamHello(SendMessage_ClientStreamHelloServer) error
	BiDricetionalStreamHello(SendMessage_BiDricetionalStreamHelloServer) error
}

// UnimplementedSendMessageServer can be embedded to have forward compatible implementations.
type UnimplementedSendMessageServer struct {
}

func (*UnimplementedSendMessageServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedSendMessageServer) ServerStreamHello(*Message, SendMessage_ServerStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamHello not implemented")
}
func (*UnimplementedSendMessageServer) ClientStreamHello(SendMessage_ClientStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamHello not implemented")
}
func (*UnimplementedSendMessageServer) BiDricetionalStreamHello(SendMessage_BiDricetionalStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDricetionalStreamHello not implemented")
}

func RegisterSendMessageServer(s *grpc.Server, srv SendMessageServer) {
	s.RegisterService(&_SendMessage_serviceDesc, srv)
}

func _SendMessage_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendMessageServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SendMessage/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendMessageServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendMessage_ServerStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SendMessageServer).ServerStreamHello(m, &sendMessageServerStreamHelloServer{stream})
}

type SendMessage_ServerStreamHelloServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type sendMessageServerStreamHelloServer struct {
	grpc.ServerStream
}

func (x *sendMessageServerStreamHelloServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _SendMessage_ClientStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SendMessageServer).ClientStreamHello(&sendMessageClientStreamHelloServer{stream})
}

type SendMessage_ClientStreamHelloServer interface {
	SendAndClose(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type sendMessageClientStreamHelloServer struct {
	grpc.ServerStream
}

func (x *sendMessageClientStreamHelloServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sendMessageClientStreamHelloServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SendMessage_BiDricetionalStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SendMessageServer).BiDricetionalStreamHello(&sendMessageBiDricetionalStreamHelloServer{stream})
}

type SendMessage_BiDricetionalStreamHelloServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type sendMessageBiDricetionalStreamHelloServer struct {
	grpc.ServerStream
}

func (x *sendMessageBiDricetionalStreamHelloServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sendMessageBiDricetionalStreamHelloServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SendMessage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SendMessage",
	HandlerType: (*SendMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SendMessage_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamHello",
			Handler:       _SendMessage_ServerStreamHello_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamHello",
			Handler:       _SendMessage_ClientStreamHello_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BiDricetionalStreamHello",
			Handler:       _SendMessage_BiDricetionalStreamHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}
