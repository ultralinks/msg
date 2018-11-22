// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgLogic.proto

package msgLogic

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
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
	return fileDescriptor_1463db17ff44f5fc, []int{0}
}

func (m *ParseMsgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseMsgRequest.Unmarshal(m, b)
}
func (m *ParseMsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseMsgRequest.Marshal(b, m, deterministic)
}
func (m *ParseMsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseMsgRequest.Merge(m, src)
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

type GetLinkByTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLinkByTokenRequest) Reset()         { *m = GetLinkByTokenRequest{} }
func (m *GetLinkByTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetLinkByTokenRequest) ProtoMessage()    {}
func (*GetLinkByTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1463db17ff44f5fc, []int{1}
}

func (m *GetLinkByTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLinkByTokenRequest.Unmarshal(m, b)
}
func (m *GetLinkByTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLinkByTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetLinkByTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLinkByTokenRequest.Merge(m, src)
}
func (m *GetLinkByTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetLinkByTokenRequest.Size(m)
}
func (m *GetLinkByTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLinkByTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLinkByTokenRequest proto.InternalMessageInfo

func (m *GetLinkByTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
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
	return fileDescriptor_1463db17ff44f5fc, []int{2}
}

func (m *ParseMsgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseMsgResponse.Unmarshal(m, b)
}
func (m *ParseMsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseMsgResponse.Marshal(b, m, deterministic)
}
func (m *ParseMsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseMsgResponse.Merge(m, src)
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

type GetLinkByTokenResponse struct {
	LinkId               string   `protobuf:"bytes,1,opt,name=linkId,proto3" json:"linkId,omitempty"`
	LinkKey              string   `protobuf:"bytes,2,opt,name=linkKey,proto3" json:"linkKey,omitempty"`
	Nick                 string   `protobuf:"bytes,3,opt,name=nick,proto3" json:"nick,omitempty"`
	Avt                  string   `protobuf:"bytes,4,opt,name=avt,proto3" json:"avt,omitempty"`
	AppId                string   `protobuf:"bytes,5,opt,name=appId,proto3" json:"appId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLinkByTokenResponse) Reset()         { *m = GetLinkByTokenResponse{} }
func (m *GetLinkByTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetLinkByTokenResponse) ProtoMessage()    {}
func (*GetLinkByTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1463db17ff44f5fc, []int{3}
}

func (m *GetLinkByTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLinkByTokenResponse.Unmarshal(m, b)
}
func (m *GetLinkByTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLinkByTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetLinkByTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLinkByTokenResponse.Merge(m, src)
}
func (m *GetLinkByTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetLinkByTokenResponse.Size(m)
}
func (m *GetLinkByTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLinkByTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLinkByTokenResponse proto.InternalMessageInfo

func (m *GetLinkByTokenResponse) GetLinkId() string {
	if m != nil {
		return m.LinkId
	}
	return ""
}

func (m *GetLinkByTokenResponse) GetLinkKey() string {
	if m != nil {
		return m.LinkKey
	}
	return ""
}

func (m *GetLinkByTokenResponse) GetNick() string {
	if m != nil {
		return m.Nick
	}
	return ""
}

func (m *GetLinkByTokenResponse) GetAvt() string {
	if m != nil {
		return m.Avt
	}
	return ""
}

func (m *GetLinkByTokenResponse) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func init() {
	proto.RegisterType((*ParseMsgRequest)(nil), "msgLogic.ParseMsgRequest")
	proto.RegisterType((*GetLinkByTokenRequest)(nil), "msgLogic.GetLinkByTokenRequest")
	proto.RegisterType((*ParseMsgResponse)(nil), "msgLogic.ParseMsgResponse")
	proto.RegisterType((*GetLinkByTokenResponse)(nil), "msgLogic.GetLinkByTokenResponse")
}

func init() { proto.RegisterFile("msgLogic.proto", fileDescriptor_1463db17ff44f5fc) }

var fileDescriptor_1463db17ff44f5fc = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xef, 0x4a, 0xc3, 0x30,
	0x1c, 0x34, 0xee, 0x8f, 0xdd, 0x0f, 0xd9, 0x46, 0xd0, 0x11, 0xeb, 0x07, 0x47, 0x41, 0xd8, 0x17,
	0x8b, 0xe8, 0x1b, 0xd4, 0x0f, 0x63, 0xb8, 0xc1, 0x08, 0xfa, 0x00, 0xb1, 0x0d, 0x25, 0x74, 0x6b,
	0x62, 0x93, 0x8a, 0x7b, 0x01, 0x1f, 0xc5, 0xe7, 0x94, 0xa4, 0x7f, 0xa6, 0x63, 0x7e, 0xbb, 0x4b,
	0xaf, 0x77, 0xb9, 0x23, 0x30, 0xdc, 0xea, 0x74, 0x29, 0x53, 0x11, 0x87, 0xaa, 0x90, 0x46, 0x62,
	0xaf, 0xe1, 0xc1, 0x2d, 0x8c, 0xd6, 0xac, 0xd0, 0x7c, 0xa5, 0x53, 0xca, 0xdf, 0x4b, 0xae, 0x0d,
	0xc6, 0xd0, 0x4d, 0x98, 0x61, 0x04, 0x4d, 0xd1, 0xec, 0x9c, 0x3a, 0x1c, 0xdc, 0xc1, 0xe5, 0x9c,
	0x9b, 0xa5, 0xc8, 0xb3, 0x68, 0xf7, 0x22, 0x33, 0x9e, 0x37, 0xe2, 0x0b, 0xe8, 0x19, 0xcb, 0x9d,
	0x7a, 0x40, 0x2b, 0x12, 0x44, 0x30, 0xde, 0xbb, 0x6a, 0x25, 0x73, 0xcd, 0xb1, 0x0f, 0xde, 0x46,
	0xe4, 0xd9, 0x33, 0xdf, 0x69, 0x82, 0xa6, 0x9d, 0xd9, 0x80, 0xb6, 0xbc, 0x8d, 0x3c, 0xfd, 0x15,
	0xf9, 0x85, 0x60, 0x72, 0x98, 0x59, 0x5b, 0x4d, 0xa0, 0x6f, 0x7f, 0x5d, 0x24, 0x75, 0x6a, 0xcd,
	0x30, 0x81, 0xb3, 0xda, 0xd2, 0x39, 0x0d, 0x68, 0x43, 0x6d, 0x40, 0x2e, 0xe2, 0x8c, 0x74, 0xdc,
	0xb1, 0xc3, 0x78, 0x0c, 0x1d, 0xf6, 0x61, 0x48, 0xd7, 0x1d, 0x59, 0x68, 0xcb, 0x30, 0xa5, 0x16,
	0x09, 0xe9, 0x55, 0x65, 0x1c, 0x79, 0xf8, 0x46, 0xe0, 0xad, 0xea, 0xbd, 0xf0, 0x13, 0x78, 0x4d,
	0x33, 0x7c, 0x15, 0xb6, 0xb3, 0x1e, 0x6c, 0xe8, 0xfb, 0xc7, 0x3e, 0x55, 0xb7, 0x0f, 0x4e, 0xf0,
	0x2b, 0x0c, 0xff, 0x36, 0xc3, 0x37, 0x7b, 0xfd, 0xd1, 0x9d, 0xfd, 0xe9, 0xff, 0x82, 0xc6, 0x36,
	0xba, 0x87, 0x6b, 0x21, 0xc3, 0xb4, 0x50, 0x71, 0xc8, 0x3f, 0xd9, 0x56, 0x6d, 0xb8, 0x0e, 0x0b,
	0x59, 0x1a, 0x9e, 0x96, 0x22, 0xe1, 0xd1, 0x88, 0x5a, 0x3c, 0xb7, 0x78, 0x6d, 0x5f, 0xc1, 0x1a,
	0xbd, 0xf5, 0xdd, 0x73, 0x78, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x04, 0x98, 0xfd, 0x20,
	0x02, 0x00, 0x00,
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
	GetLinkByToken(ctx context.Context, in *GetLinkByTokenRequest, opts ...grpc.CallOption) (*GetLinkByTokenResponse, error)
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

func (c *msgLogicClient) GetLinkByToken(ctx context.Context, in *GetLinkByTokenRequest, opts ...grpc.CallOption) (*GetLinkByTokenResponse, error) {
	out := new(GetLinkByTokenResponse)
	err := c.cc.Invoke(ctx, "/msgLogic.MsgLogic/GetLinkByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgLogicServer is the server API for MsgLogic service.
type MsgLogicServer interface {
	ParseMsg(context.Context, *ParseMsgRequest) (*ParseMsgResponse, error)
	GetLinkByToken(context.Context, *GetLinkByTokenRequest) (*GetLinkByTokenResponse, error)
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

func _MsgLogic_GetLinkByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkByTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgLogicServer).GetLinkByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msgLogic.MsgLogic/GetLinkByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgLogicServer).GetLinkByToken(ctx, req.(*GetLinkByTokenRequest))
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
		{
			MethodName: "GetLinkByToken",
			Handler:    _MsgLogic_GetLinkByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msgLogic.proto",
}
