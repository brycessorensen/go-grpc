// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	auth.proto
	currency.proto

It has these top-level messages:
	AuthRequest
	AuthResponse
	Currency
	CurrencyList
	CurrencyRequest
*/
package protobuf

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

type AuthRequest struct {
	Uname string `protobuf:"bytes,1,opt,name=uname" json:"uname,omitempty"`
	Pwd   string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *AuthRequest) Reset()                    { *m = AuthRequest{} }
func (m *AuthRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()               {}
func (*AuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthRequest) GetUname() string {
	if m != nil {
		return m.Uname
	}
	return ""
}

func (m *AuthRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type AuthResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *AuthResponse) Reset()                    { *m = AuthResponse{} }
func (m *AuthResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()               {}
func (*AuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "protobuf.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "protobuf.AuthResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	Login(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/protobuf.Auth/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Login(context.Context, *AuthRequest) (*AuthResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Auth/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0xa6, 0x5c,
	0xdc, 0x8e, 0xa5, 0x25, 0x19, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x22, 0x5c, 0xac,
	0xa5, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x00,
	0x17, 0x73, 0x41, 0x79, 0x8a, 0x04, 0x13, 0x58, 0x0c, 0xc4, 0x54, 0x52, 0xe1, 0xe2, 0x81, 0x68,
	0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x05, 0xe9, 0x2b, 0xc9, 0xcf, 0x4e, 0xcd, 0x83, 0xe9, 0x03,
	0x73, 0x8c, 0x1c, 0xb8, 0x58, 0x40, 0xaa, 0x84, 0x2c, 0xb8, 0x58, 0x7d, 0xf2, 0xd3, 0x33, 0xf3,
	0x84, 0x44, 0xf5, 0x60, 0x16, 0xeb, 0x21, 0xd9, 0x2a, 0x25, 0x86, 0x2e, 0x0c, 0x31, 0x55, 0x89,
	0x21, 0x89, 0x0d, 0x2c, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x05, 0x4b, 0x23, 0xbd,
	0x00, 0x00, 0x00,
}
