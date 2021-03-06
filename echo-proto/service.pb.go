// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package echoproto is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	EchoMessage
*/
package echoproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoMessage struct {
	Body string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *EchoMessage) Reset()                    { *m = EchoMessage{} }
func (m *EchoMessage) String() string            { return proto.CompactTextString(m) }
func (*EchoMessage) ProtoMessage()               {}
func (*EchoMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoMessage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoMessage)(nil), "echoproto.EchoMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoService service

type EchoServiceClient interface {
	Hello(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
	Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Hello(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := grpc.Invoke(ctx, "/echoproto.EchoService/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := grpc.Invoke(ctx, "/echoproto.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	Hello(context.Context, *EchoMessage) (*EchoMessage, error)
	Echo(context.Context, *EchoMessage) (*EchoMessage, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echoproto.EchoService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Hello(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echoproto.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echoproto.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _EchoService_Hello_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x4d, 0xce, 0xc8, 0x07,
	0x33, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3,
	0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0x0a, 0x95, 0x14, 0xb9, 0xb8, 0x5d,
	0x93, 0x33, 0xf2, 0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x92, 0xf2,
	0x53, 0x2a, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0xa3, 0x65, 0x8c, 0x10, 0x35,
	0xc1, 0x10, 0x1b, 0x84, 0xbc, 0xb8, 0x58, 0x3d, 0x52, 0x73, 0x72, 0xf2, 0x85, 0xc4, 0xf4, 0xe0,
	0xb6, 0xe8, 0x21, 0x19, 0x22, 0x85, 0x43, 0x5c, 0x49, 0xb0, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xdc,
	0x42, 0x9c, 0xfa, 0x65, 0x86, 0xfa, 0x19, 0x60, 0x23, 0xbc, 0xb9, 0x58, 0x40, 0x2a, 0x48, 0x36,
	0x4a, 0x18, 0x6c, 0x14, 0xaf, 0x12, 0x07, 0xc8, 0x28, 0x90, 0x12, 0x2b, 0x46, 0xad, 0x24, 0x36,
	0xb0, 0x3a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x13, 0xed, 0x00, 0x0c, 0x01, 0x00,
	0x00,
}
