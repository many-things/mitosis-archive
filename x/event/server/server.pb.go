// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitosis/event/v1beta1/server/server.proto

package server

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("mitosis/event/v1beta1/server/server.proto", fileDescriptor_8493cefcc4fba2c7)
}

var fileDescriptor_8493cefcc4fba2c7 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcc, 0xcd, 0x2c, 0xc9,
	0x2f, 0xce, 0x2c, 0xd6, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x82, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0xea, 0xb9, 0x89, 0x79, 0x95, 0x25, 0x19, 0x99, 0x79, 0xe9, 0xc5, 0x7a, 0x50, 0x5d, 0x7a,
	0x50, 0xf5, 0x7a, 0x60, 0xdd, 0x7a, 0x10, 0xe5, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9,
	0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5,
	0x10, 0x63, 0xa4, 0xf4, 0xf0, 0xda, 0x58, 0x52, 0x59, 0x90, 0x5a, 0x1c, 0x5f, 0x58, 0x9a, 0x5a,
	0x54, 0x09, 0x51, 0x6f, 0xc4, 0xca, 0xc5, 0xec, 0x5b, 0x9c, 0x6e, 0xb4, 0x83, 0x91, 0x8b, 0x35,
	0x10, 0x24, 0x2c, 0xb4, 0x8e, 0x91, 0x8b, 0x2d, 0x20, 0xb1, 0x28, 0x31, 0xb7, 0x58, 0xc8, 0x5a,
	0x8f, 0x48, 0x37, 0xe9, 0x81, 0xb5, 0x42, 0x74, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48,
	0xd9, 0x90, 0xa7, 0xb9, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x49, 0xb3, 0xe9, 0xf2, 0x93, 0xc9,
	0x4c, 0xca, 0x42, 0x8a, 0xfa, 0x20, 0x53, 0x74, 0x21, 0xc6, 0xe8, 0xa3, 0xfa, 0xad, 0x00, 0xac,
	0xc5, 0xc9, 0xfd, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd3, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xb1, 0x1a, 0x53, 0x01, 0x35, 0x08, 0xe2, 0x8e,
	0x24, 0x36, 0x70, 0x88, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x7b, 0x46, 0x43, 0xb5,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manythings.mitosis.v1beta1.event.server.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "mitosis/event/v1beta1/server/server.proto",
}

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.event.server.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.event.server.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manythings.mitosis.v1beta1.event.server.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitosis/event/v1beta1/server/server.proto",
}
