// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgLogic.proto

package msgLogic

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

type ParseMsgRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseMsgRequest) Reset()         { *m = ParseMsgRequest{} }
func (m *ParseMsgRequest) String() string { return proto.CompactTextString(m) }
func (*ParseMsgRequest) ProtoMessage()    {}
func (*ParseMsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_msgLogic_0ea004e4935b65cc, []int{0}
}
func (m *ParseMsgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseMsgRequest.Unmarshal(m, b)
}
func (m *ParseMsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseMsgRequest.Marshal(b, m, deterministic)
}
func (dst *ParseMsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseMsgRequest.Merge(dst, src)
}
func (m *ParseMsgRequest) XXX_Size() int {
	return xxx_messageInfo_ParseMsgRequest.Size(m)
}
func (m *ParseMsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseMsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParseMsgRequest proto.InternalMessageInfo

func (m *ParseMsgRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ParseMsgResponse struct {
	LinkKeys             []string `protobuf:"bytes,1,rep,name=linkKeys,proto3" json:"linkKeys,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseMsgResponse) Reset()         { *m = ParseMsgResponse{} }
func (m *ParseMsgResponse) String() string { return proto.CompactTextString(m) }
func (*ParseMsgResponse) ProtoMessage()    {}
func (*ParseMsgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_msgLogic_0ea004e4935b65cc, []int{1}
}
func (m *ParseMsgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseMsgResponse.Unmarshal(m, b)
}
func (m *ParseMsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseMsgResponse.Marshal(b, m, deterministic)
}
func (dst *ParseMsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseMsgResponse.Merge(dst, src)
}
func (m *ParseMsgResponse) XXX_Size() int {
	return xxx_messageInfo_ParseMsgResponse.Size(m)
}
func (m *ParseMsgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseMsgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParseMsgResponse proto.InternalMessageInfo

func (m *ParseMsgResponse) GetLinkKeys() []string {
	if m != nil {
		return m.LinkKeys
	}
	return nil
}

func (m *ParseMsgResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ParseMsgRequest)(nil), "msgLogic.ParseMsgRequest")
	proto.RegisterType((*ParseMsgResponse)(nil), "msgLogic.ParseMsgResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgLogicClient is the client API for MsgLogic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgLogicClient interface {
	ParseMsg(ctx context.Context, in *ParseMsgRequest, opts ...grpc.CallOption) (*ParseMsgResponse, error)
}

type msgLogicClient struct {
	cc *grpc.ClientConn
}

func NewMsgLogicClient(cc *grpc.ClientConn) MsgLogicClient {
	return &msgLogicClient{cc}
}

func (c *msgLogicClient) ParseMsg(ctx context.Context, in *ParseMsgRequest, opts ...grpc.CallOption) (*ParseMsgResponse, error) {
	out := new(ParseMsgResponse)
	err := c.cc.Invoke(ctx, "/msgLogic.MsgLogic/ParseMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgLogicServer is the server API for MsgLogic service.
type MsgLogicServer interface {
	ParseMsg(context.Context, *ParseMsgRequest) (*ParseMsgResponse, error)
}

func RegisterMsgLogicServer(s *grpc.Server, srv MsgLogicServer) {
	s.RegisterService(&_MsgLogic_serviceDesc, srv)
}

func _MsgLogic_ParseMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgLogicServer).ParseMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgLogic.MsgLogic/ParseMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgLogicServer).ParseMsg(ctx, req.(*ParseMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgLogic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "msgLogic.MsgLogic",
	HandlerType: (*MsgLogicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseMsg",
			Handler:    _MsgLogic_ParseMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msgLogic.proto",
}

func init() { proto.RegisterFile("msgLogic.proto", fileDescriptor_msgLogic_0ea004e4935b65cc) }

var fileDescriptor_msgLogic_0ea004e4935b65cc = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2d, 0x4e, 0xf7,
	0xc9, 0x4f, 0xcf, 0x4c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x54,
	0xb9, 0xf8, 0x03, 0x12, 0x8b, 0x8a, 0x53, 0x7d, 0x8b, 0xd3, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82,
	0xc0, 0x6c, 0x25, 0x27, 0x2e, 0x01, 0x84, 0xb2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x29,
	0x2e, 0x8e, 0x9c, 0xcc, 0xbc, 0x6c, 0xef, 0xd4, 0xca, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xce,
	0x20, 0x38, 0x1f, 0x6e, 0x06, 0x13, 0xc2, 0x0c, 0x23, 0x7f, 0x2e, 0x0e, 0x5f, 0xa8, 0xb5, 0x42,
	0xce, 0x5c, 0x1c, 0x30, 0xf3, 0x84, 0x24, 0xf5, 0xe0, 0xae, 0x43, 0x73, 0x8a, 0x94, 0x14, 0x36,
	0x29, 0x88, 0xf5, 0x4a, 0x0c, 0x4e, 0x06, 0x5c, 0xd2, 0x99, 0xf9, 0x7a, 0xe9, 0x45, 0x05, 0xc9,
	0x7a, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0xc5, 0x7a, 0x45, 0xf9, 0xa5, 0x25, 0xa9, 0xe9,
	0xa5, 0x99, 0x29, 0xa9, 0x4e, 0xfc, 0x41, 0x20, 0xb6, 0x3b, 0x88, 0x1d, 0x00, 0xf2, 0x75, 0x00,
	0x63, 0x12, 0x1b, 0xd8, 0xfb, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xe9, 0x53, 0x7a,
	0x10, 0x01, 0x00, 0x00,
}
