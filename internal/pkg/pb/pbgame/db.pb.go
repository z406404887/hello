// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db.proto

package pbgame

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

type LoadRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadRequest) Reset()         { *m = LoadRequest{} }
func (m *LoadRequest) String() string { return proto.CompactTextString(m) }
func (*LoadRequest) ProtoMessage()    {}
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_440a6f2790ff5b24, []int{0}
}
func (m *LoadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadRequest.Unmarshal(m, b)
}
func (m *LoadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadRequest.Marshal(b, m, deterministic)
}
func (dst *LoadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadRequest.Merge(dst, src)
}
func (m *LoadRequest) XXX_Size() int {
	return xxx_messageInfo_LoadRequest.Size(m)
}
func (m *LoadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadRequest proto.InternalMessageInfo

func (m *LoadRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type LoadResponse struct {
	Result               ErrorCode `protobuf:"varint,1,opt,name=result,enum=pbgame.ErrorCode" json:"result,omitempty"`
	Uid                  uint32    `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	Name                 string    `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Money                int32     `protobuf:"varint,4,opt,name=money" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LoadResponse) Reset()         { *m = LoadResponse{} }
func (m *LoadResponse) String() string { return proto.CompactTextString(m) }
func (*LoadResponse) ProtoMessage()    {}
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_440a6f2790ff5b24, []int{1}
}
func (m *LoadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadResponse.Unmarshal(m, b)
}
func (m *LoadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadResponse.Marshal(b, m, deterministic)
}
func (dst *LoadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadResponse.Merge(dst, src)
}
func (m *LoadResponse) XXX_Size() int {
	return xxx_messageInfo_LoadResponse.Size(m)
}
func (m *LoadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadResponse proto.InternalMessageInfo

func (m *LoadResponse) GetResult() ErrorCode {
	if m != nil {
		return m.Result
	}
	return ErrorCode_SUCCESS
}

func (m *LoadResponse) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *LoadResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoadResponse) GetMoney() int32 {
	if m != nil {
		return m.Money
	}
	return 0
}

type SaveRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Uid                  uint32   `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Money                int32    `protobuf:"varint,4,opt,name=money" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveRequest) Reset()         { *m = SaveRequest{} }
func (m *SaveRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()    {}
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_440a6f2790ff5b24, []int{2}
}
func (m *SaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest.Unmarshal(m, b)
}
func (m *SaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest.Marshal(b, m, deterministic)
}
func (dst *SaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest.Merge(dst, src)
}
func (m *SaveRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRequest.Size(m)
}
func (m *SaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest proto.InternalMessageInfo

func (m *SaveRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *SaveRequest) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SaveRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SaveRequest) GetMoney() int32 {
	if m != nil {
		return m.Money
	}
	return 0
}

type SaveResponse struct {
	Result               ErrorCode `protobuf:"varint,1,opt,name=result,enum=pbgame.ErrorCode" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SaveResponse) Reset()         { *m = SaveResponse{} }
func (m *SaveResponse) String() string { return proto.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()    {}
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_440a6f2790ff5b24, []int{3}
}
func (m *SaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResponse.Unmarshal(m, b)
}
func (m *SaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResponse.Marshal(b, m, deterministic)
}
func (dst *SaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResponse.Merge(dst, src)
}
func (m *SaveResponse) XXX_Size() int {
	return xxx_messageInfo_SaveResponse.Size(m)
}
func (m *SaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResponse proto.InternalMessageInfo

func (m *SaveResponse) GetResult() ErrorCode {
	if m != nil {
		return m.Result
	}
	return ErrorCode_SUCCESS
}

type CreatePlayerRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Money                int32    `protobuf:"varint,3,opt,name=money" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePlayerRequest) Reset()         { *m = CreatePlayerRequest{} }
func (m *CreatePlayerRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePlayerRequest) ProtoMessage()    {}
func (*CreatePlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_440a6f2790ff5b24, []int{4}
}
func (m *CreatePlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlayerRequest.Unmarshal(m, b)
}
func (m *CreatePlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlayerRequest.Marshal(b, m, deterministic)
}
func (dst *CreatePlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayerRequest.Merge(dst, src)
}
func (m *CreatePlayerRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePlayerRequest.Size(m)
}
func (m *CreatePlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayerRequest proto.InternalMessageInfo

func (m *CreatePlayerRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CreatePlayerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreatePlayerRequest) GetMoney() int32 {
	if m != nil {
		return m.Money
	}
	return 0
}

type CreatePlayerResponse struct {
	Result               ErrorCode `protobuf:"varint,1,opt,name=result,enum=pbgame.ErrorCode" json:"result,omitempty"`
	Uid                  uint32    `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreatePlayerResponse) Reset()         { *m = CreatePlayerResponse{} }
func (m *CreatePlayerResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePlayerResponse) ProtoMessage()    {}
func (*CreatePlayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db_440a6f2790ff5b24, []int{5}
}
func (m *CreatePlayerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlayerResponse.Unmarshal(m, b)
}
func (m *CreatePlayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlayerResponse.Marshal(b, m, deterministic)
}
func (dst *CreatePlayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayerResponse.Merge(dst, src)
}
func (m *CreatePlayerResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePlayerResponse.Size(m)
}
func (m *CreatePlayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayerResponse proto.InternalMessageInfo

func (m *CreatePlayerResponse) GetResult() ErrorCode {
	if m != nil {
		return m.Result
	}
	return ErrorCode_SUCCESS
}

func (m *CreatePlayerResponse) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterType((*LoadRequest)(nil), "pbgame.LoadRequest")
	proto.RegisterType((*LoadResponse)(nil), "pbgame.LoadResponse")
	proto.RegisterType((*SaveRequest)(nil), "pbgame.SaveRequest")
	proto.RegisterType((*SaveResponse)(nil), "pbgame.SaveResponse")
	proto.RegisterType((*CreatePlayerRequest)(nil), "pbgame.CreatePlayerRequest")
	proto.RegisterType((*CreatePlayerResponse)(nil), "pbgame.CreatePlayerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DBClient is the client API for DB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DBClient interface {
	LoadPlayer(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error)
	SavePlayer(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
	CreatePlayer(ctx context.Context, in *CreatePlayerRequest, opts ...grpc.CallOption) (*CreatePlayerResponse, error)
}

type dBClient struct {
	cc *grpc.ClientConn
}

func NewDBClient(cc *grpc.ClientConn) DBClient {
	return &dBClient{cc}
}

func (c *dBClient) LoadPlayer(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error) {
	out := new(LoadResponse)
	err := c.cc.Invoke(ctx, "/pbgame.DB/LoadPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) SavePlayer(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/pbgame.DB/SavePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) CreatePlayer(ctx context.Context, in *CreatePlayerRequest, opts ...grpc.CallOption) (*CreatePlayerResponse, error) {
	out := new(CreatePlayerResponse)
	err := c.cc.Invoke(ctx, "/pbgame.DB/CreatePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DB service

type DBServer interface {
	LoadPlayer(context.Context, *LoadRequest) (*LoadResponse, error)
	SavePlayer(context.Context, *SaveRequest) (*SaveResponse, error)
	CreatePlayer(context.Context, *CreatePlayerRequest) (*CreatePlayerResponse, error)
}

func RegisterDBServer(s *grpc.Server, srv DBServer) {
	s.RegisterService(&_DB_serviceDesc, srv)
}

func _DB_LoadPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).LoadPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgame.DB/LoadPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).LoadPlayer(ctx, req.(*LoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_SavePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).SavePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgame.DB/SavePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).SavePlayer(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DB_CreatePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServer).CreatePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbgame.DB/CreatePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServer).CreatePlayer(ctx, req.(*CreatePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbgame.DB",
	HandlerType: (*DBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadPlayer",
			Handler:    _DB_LoadPlayer_Handler,
		},
		{
			MethodName: "SavePlayer",
			Handler:    _DB_SavePlayer_Handler,
		},
		{
			MethodName: "CreatePlayer",
			Handler:    _DB_CreatePlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db.proto",
}

func init() { proto.RegisterFile("db.proto", fileDescriptor_db_440a6f2790ff5b24) }

var fileDescriptor_db_440a6f2790ff5b24 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x2d, 0xd0, 0xa2, 0x0e, 0x68, 0x74, 0xca, 0x81, 0xa0, 0x07, 0xb2, 0x17, 0xf1, 0xc2, 0xa1,
	0x9e, 0x7a, 0xb5, 0x7a, 0xd2, 0x83, 0xa1, 0x27, 0x8f, 0x0b, 0x4c, 0x8c, 0x49, 0x61, 0x71, 0x01,
	0x93, 0xfe, 0x42, 0xff, 0x96, 0x81, 0x95, 0xb8, 0x35, 0x24, 0x8d, 0xb1, 0xb7, 0xfd, 0x78, 0xf3,
	0xde, 0x9b, 0x37, 0x03, 0xc7, 0x79, 0x1a, 0x57, 0x52, 0x34, 0x02, 0xed, 0x2a, 0x7d, 0xe5, 0x05,
	0x05, 0x0e, 0x49, 0x29, 0xa4, 0x7a, 0x64, 0xd7, 0xe0, 0x3c, 0x09, 0x9e, 0x27, 0xf4, 0xde, 0x52,
	0xdd, 0xa0, 0x0f, 0x47, 0x3c, 0xcb, 0x44, 0x5b, 0x36, 0xbe, 0x11, 0x1a, 0xd1, 0x49, 0x32, 0x5c,
	0x59, 0x0b, 0xae, 0x02, 0xd6, 0x95, 0x28, 0x6b, 0xc2, 0x1b, 0xb0, 0x25, 0xd5, 0xed, 0x46, 0x01,
	0xcf, 0x16, 0x17, 0xb1, 0xa2, 0x8f, 0x1f, 0x3a, 0xf6, 0x95, 0xc8, 0x29, 0xf9, 0x06, 0xe0, 0x39,
	0x58, 0xed, 0x5b, 0xee, 0x9b, 0xa1, 0x11, 0x9d, 0x26, 0xdd, 0x11, 0x11, 0xa6, 0x25, 0x2f, 0xc8,
	0xb7, 0x7a, 0x8d, 0xfe, 0x8c, 0x1e, 0xcc, 0x0a, 0x51, 0xd2, 0xd6, 0x9f, 0x86, 0x46, 0x34, 0x4b,
	0xd4, 0x85, 0x65, 0xe0, 0xac, 0xf9, 0x07, 0xed, 0xf5, 0xf7, 0x2f, 0x91, 0x25, 0xb8, 0x4a, 0xe4,
	0xcf, 0xbd, 0xb1, 0x17, 0x98, 0xaf, 0x24, 0xf1, 0x86, 0x9e, 0x37, 0x7c, 0x4b, 0x72, 0xbf, 0xcf,
	0xc1, 0x95, 0x39, 0xe6, 0xca, 0xd2, 0x5d, 0xad, 0xc1, 0xdb, 0xa5, 0x3e, 0x40, 0xf2, 0x8b, 0x4f,
	0x03, 0xcc, 0xfb, 0x3b, 0x5c, 0x02, 0x74, 0xd3, 0x54, 0xcc, 0x38, 0x1f, 0x18, 0xb4, 0x55, 0x08,
	0xbc, 0xdd, 0x47, 0x25, 0xce, 0x26, 0x5d, 0x69, 0x17, 0xd6, 0xef, 0x52, 0x6d, 0x4a, 0x3f, 0xa5,
	0x7a, 0xaa, 0x6c, 0x82, 0x8f, 0xe0, 0xea, 0x1d, 0xe1, 0xe5, 0x80, 0x1b, 0x89, 0x30, 0xb8, 0x1a,
	0xff, 0x1c, 0xc8, 0x52, 0xbb, 0x5f, 0xe0, 0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0x16,
	0x73, 0xa4, 0xe1, 0x02, 0x00, 0x00,
}
