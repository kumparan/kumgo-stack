// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/server/echosvcserver/buff/echo.proto

/*
Package buff is a generated protocol buffer package.

It is generated from these files:
	example/server/echosvcserver/buff/echo.proto

It has these top-level messages:
	EchoRequest
	EchoResponse
*/
package buff

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type EchoRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EchoResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "buff.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "buff.EchoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoService service

type EchoServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := grpc.Invoke(ctx, "/buff.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buff.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buff.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/server/echosvcserver/buff/echo.proto",
}

func init() { proto.RegisterFile("example/server/echosvcserver/buff/echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x4f, 0x4d, 0xce, 0xc8, 0x2f,
	0x2e, 0x4b, 0x86, 0xf2, 0x92, 0x4a, 0xd3, 0xd2, 0xc0, 0x42, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x2c, 0x20, 0x01, 0x25, 0x75, 0x2e, 0x6e, 0xd7, 0xe4, 0x8c, 0xfc, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x49, 0x83, 0x8b, 0x07, 0xa2, 0xb0, 0xb8, 0x20, 0x3f,
	0xaf, 0x38, 0x15, 0xb7, 0x4a, 0x23, 0x3b, 0x88, 0x91, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9,
	0x42, 0xfa, 0x5c, 0x2c, 0x20, 0xae, 0x90, 0xa0, 0x1e, 0xc8, 0x42, 0x3d, 0x24, 0xdb, 0xa4, 0x84,
	0x90, 0x85, 0x20, 0xe6, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0xdd, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x94, 0x6a, 0xe6, 0x0c, 0xcf, 0x00, 0x00, 0x00,
}