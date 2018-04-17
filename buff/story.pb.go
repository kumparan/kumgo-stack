// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buff/story.proto

/*
Package buff is a generated protocol buffer package.

It is generated from these files:
	buff/story.proto

It has these top-level messages:
	Empty
	GetStoriesResponse
	Story
	Author
	Editor
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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetStoriesResponse struct {
	Stories []*Story `protobuf:"bytes,1,rep,name=stories" json:"stories,omitempty"`
}

func (m *GetStoriesResponse) Reset()                    { *m = GetStoriesResponse{} }
func (m *GetStoriesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetStoriesResponse) ProtoMessage()               {}
func (*GetStoriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetStoriesResponse) GetStories() []*Story {
	if m != nil {
		return m.Stories
	}
	return nil
}

type Story struct {
	Author *Author `protobuf:"bytes,1,opt,name=author" json:"author,omitempty"`
	Editor *Editor `protobuf:"bytes,2,opt,name=editor" json:"editor,omitempty"`
}

func (m *Story) Reset()                    { *m = Story{} }
func (m *Story) String() string            { return proto.CompactTextString(m) }
func (*Story) ProtoMessage()               {}
func (*Story) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Story) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Story) GetEditor() *Editor {
	if m != nil {
		return m.Editor
	}
	return nil
}

type Author struct {
	Group    string `protobuf:"bytes,1,opt,name=group" json:"group,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Username string `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
}

func (m *Author) Reset()                    { *m = Author{} }
func (m *Author) String() string            { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()               {}
func (*Author) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Author) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Author) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Author) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Editor struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username" json:"username,omitempty"`
}

func (m *Editor) Reset()                    { *m = Editor{} }
func (m *Editor) String() string            { return proto.CompactTextString(m) }
func (*Editor) ProtoMessage()               {}
func (*Editor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Editor) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Editor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Editor) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "buff.Empty")
	proto.RegisterType((*GetStoriesResponse)(nil), "buff.GetStoriesResponse")
	proto.RegisterType((*Story)(nil), "buff.Story")
	proto.RegisterType((*Author)(nil), "buff.Author")
	proto.RegisterType((*Editor)(nil), "buff.Editor")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StoryService service

type StoryServiceClient interface {
	GetStories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetStoriesResponse, error)
}

type storyServiceClient struct {
	cc *grpc.ClientConn
}

func NewStoryServiceClient(cc *grpc.ClientConn) StoryServiceClient {
	return &storyServiceClient{cc}
}

func (c *storyServiceClient) GetStories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetStoriesResponse, error) {
	out := new(GetStoriesResponse)
	err := grpc.Invoke(ctx, "/buff.StoryService/GetStories", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StoryService service

type StoryServiceServer interface {
	GetStories(context.Context, *Empty) (*GetStoriesResponse, error)
}

func RegisterStoryServiceServer(s *grpc.Server, srv StoryServiceServer) {
	s.RegisterService(&_StoryService_serviceDesc, srv)
}

func _StoryService_GetStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).GetStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buff.StoryService/GetStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).GetStories(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buff.StoryService",
	HandlerType: (*StoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStories",
			Handler:    _StoryService_GetStories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buff/story.proto",
}

func init() { proto.RegisterFile("buff/story.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xed, 0xe7, 0xba, 0xd3, 0x45, 0x64, 0xf0, 0x10, 0xf6, 0xb4, 0x04, 0x85, 0x3d, 0x55,
	0xa8, 0x78, 0xf2, 0xe4, 0xa1, 0xe8, 0x39, 0xbd, 0x0b, 0xbb, 0x76, 0x56, 0x73, 0xd8, 0x4d, 0x49,
	0x52, 0xa1, 0xff, 0xbd, 0x74, 0x52, 0xbf, 0x28, 0xde, 0x32, 0xbf, 0xf7, 0xe6, 0x85, 0xbc, 0xc0,
	0xe5, 0xbe, 0x3f, 0x1c, 0x6e, 0x9d, 0x37, 0x76, 0x28, 0x3b, 0x6b, 0xbc, 0xc1, 0x74, 0x24, 0x72,
	0x01, 0x59, 0x7d, 0xec, 0xfc, 0x20, 0x1f, 0x00, 0x9f, 0xc8, 0x37, 0xde, 0x58, 0x4d, 0x4e, 0x91,
	0xeb, 0xcc, 0xc9, 0x11, 0xde, 0xc0, 0xc2, 0x05, 0x24, 0xa2, 0x4d, 0xb2, 0x2d, 0xaa, 0xa2, 0x1c,
	0xd7, 0xca, 0xd1, 0x37, 0xa8, 0x2f, 0x4d, 0x36, 0x90, 0x31, 0xc1, 0x6b, 0xc8, 0x77, 0xbd, 0x7f,
	0x37, 0x56, 0x44, 0x9b, 0x68, 0x5b, 0x54, 0xab, 0x60, 0x7f, 0x64, 0xa6, 0x26, 0x6d, 0x74, 0x51,
	0xab, 0xbd, 0xb1, 0x22, 0xfe, 0xed, 0xaa, 0x99, 0xa9, 0x49, 0x93, 0x2f, 0x90, 0x87, 0x3d, 0xbc,
	0x82, 0xec, 0xcd, 0x9a, 0xbe, 0xe3, 0xd0, 0xa5, 0x0a, 0x03, 0x5e, 0x40, 0xac, 0x5b, 0x4e, 0x58,
	0xaa, 0x58, 0xb7, 0x88, 0x90, 0x9e, 0x76, 0x47, 0x12, 0x09, 0x13, 0x3e, 0xe3, 0x1a, 0xce, 0x7b,
	0x47, 0x96, 0x79, 0xca, 0xfc, 0x7b, 0x96, 0xcf, 0x90, 0x87, 0x1b, 0xa7, 0xa4, 0x68, 0x96, 0x14,
	0xff, 0x93, 0x94, 0xfc, 0x4d, 0xaa, 0x6a, 0x58, 0xf1, 0xf3, 0x1b, 0xb2, 0x1f, 0xfa, 0x95, 0xf0,
	0x1e, 0xe0, 0xa7, 0x4b, 0x9c, 0x2a, 0xe3, 0x9a, 0xd7, 0x22, 0x0c, 0xf3, 0xaa, 0xe5, 0xd9, 0x3e,
	0xe7, 0x8f, 0xb9, 0xfb, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x85, 0x1b, 0x0d, 0x53, 0xac, 0x01, 0x00,
	0x00,
}