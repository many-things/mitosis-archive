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
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x3f, 0x4b, 0xf3, 0x40,
	0x18, 0xc0, 0x7b, 0x7d, 0xdf, 0x16, 0xde, 0x83, 0x77, 0xb9, 0x31, 0xbc, 0x84, 0xd7, 0x0a, 0xda,
	0x0e, 0xbd, 0xd0, 0x8a, 0x83, 0x56, 0xf1, 0x0f, 0x88, 0x53, 0xa1, 0x56, 0x70, 0x70, 0x91, 0xab,
	0x3c, 0xa4, 0x01, 0x93, 0x8b, 0xb9, 0x27, 0xa1, 0x19, 0x75, 0x17, 0x04, 0x27, 0x3f, 0x81, 0xe0,
	0xe0, 0x57, 0x70, 0x75, 0x2c, 0xb8, 0x38, 0x4a, 0xeb, 0x07, 0x91, 0xe4, 0x52, 0xa5, 0x20, 0xa5,
	0xcd, 0x94, 0xe5, 0xf9, 0xfd, 0xee, 0xc7, 0x73, 0x39, 0x5a, 0x73, 0x1d, 0x94, 0xca, 0x51, 0x16,
	0x44, 0xe0, 0xa1, 0x15, 0x35, 0x7a, 0x80, 0xa2, 0x61, 0x29, 0x08, 0x22, 0x08, 0xb2, 0x0f, 0xf7,
	0x03, 0x89, 0x92, 0xad, 0xba, 0xc2, 0x8b, 0xb1, 0xef, 0x78, 0xb6, 0xe2, 0x19, 0xc5, 0xb3, 0x79,
	0x9e, 0xd2, 0x5c, 0x8f, 0x1b, 0xff, 0x6c, 0x29, 0xed, 0x0b, 0xb0, 0x84, 0xef, 0x58, 0xc2, 0xf3,
	0x24, 0x0a, 0x74, 0xa4, 0xa7, 0xb4, 0xc6, 0xa8, 0xce, 0x3c, 0x11, 0x63, 0x1f, 0xb2, 0xc9, 0xe6,
	0x7d, 0x91, 0xfe, 0x6a, 0x2b, 0x9b, 0x5d, 0x11, 0xfa, 0xe7, 0x44, 0x22, 0x1c, 0x24, 0x00, 0x5b,
	0xe7, 0x73, 0x76, 0xf0, 0xb6, 0xb2, 0xbf, 0x30, 0x63, 0x3b, 0x17, 0xd6, 0x05, 0xe5, 0x4b, 0x4f,
	0x01, 0xbb, 0x21, 0xf4, 0x6f, 0x17, 0x6c, 0x47, 0x21, 0x04, 0x9d, 0x40, 0x0e, 0x62, 0xb6, 0xb1,
	0x88, 0x70, 0x0a, 0x35, 0xf6, 0x72, 0xa3, 0x93, 0x9e, 0xe6, 0xc3, 0x6f, 0x5a, 0x3a, 0x0a, 0x21,
	0x88, 0xd9, 0x13, 0xa1, 0xe5, 0x8e, 0x08, 0x84, 0xab, 0x58, 0x6b, 0x6e, 0x6f, 0x8a, 0x6a, 0xaa,
	0x0b, 0x97, 0x21, 0x28, 0x34, 0xb6, 0xf2, 0xc1, 0xba, 0xa7, 0x52, 0xbb, 0x7e, 0xfd, 0xb8, 0x2b,
	0x2e, 0xb3, 0x25, 0x2b, 0xb1, 0xd4, 0xb5, 0xc6, 0x9a, 0xbe, 0x6a, 0x5f, 0x57, 0x3e, 0x13, 0x4a,
	0x93, 0x05, 0x1f, 0xa3, 0xc0, 0x50, 0xb1, 0x9d, 0xc5, 0xce, 0xfd, 0x26, 0x27, 0xe1, 0xbb, 0xf9,
	0x05, 0x59, 0x3c, 0x4f, 0xe3, 0xab, 0x6c, 0x65, 0x46, 0x7c, 0x24, 0x11, 0xce, 0x94, 0x4e, 0x7e,
	0x24, 0xb4, 0xa4, 0x7f, 0x82, 0xcd, 0x05, 0x97, 0xa6, 0xef, 0x50, 0x77, 0xb7, 0x72, 0xb1, 0x59,
	0x72, 0x35, 0x4d, 0xae, 0xb0, 0xff, 0xb3, 0xf6, 0x9d, 0x10, 0xfb, 0x87, 0x2f, 0x23, 0x93, 0x0c,
	0x47, 0x26, 0x79, 0x1f, 0x99, 0xe4, 0x76, 0x6c, 0x16, 0x86, 0x63, 0xb3, 0xf0, 0x36, 0x36, 0x0b,
	0xa7, 0x75, 0xdb, 0xc1, 0x7e, 0xd8, 0xe3, 0xe7, 0xd2, 0xfd, 0xd1, 0x32, 0xc8, 0x3c, 0xba, 0xa2,
	0x57, 0x4e, 0x5f, 0xe5, 0xda, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x9e, 0x65, 0xc1, 0x33,
	0x04, 0x00, 0x00,
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
	// VoteEvent
	VoteEvent(ctx context.Context, in *MsgVoteEvent, opts ...grpc.CallOption) (*MsgVoteEventResponse, error)
	// RegisterProxy
	RegisterProxy(ctx context.Context, in *MsgRegisterProxy, opts ...grpc.CallOption) (*MsgRegisterProxyResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) VoteEvent(ctx context.Context, in *MsgVoteEvent, opts ...grpc.CallOption) (*MsgVoteEventResponse, error) {
	out := new(MsgVoteEventResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.event.server.Msg/VoteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterProxy(ctx context.Context, in *MsgRegisterProxy, opts ...grpc.CallOption) (*MsgRegisterProxyResponse, error) {
	out := new(MsgRegisterProxyResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.event.server.Msg/RegisterProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// VoteEvent
	VoteEvent(context.Context, *MsgVoteEvent) (*MsgVoteEventResponse, error)
	// RegisterProxy
	RegisterProxy(context.Context, *MsgRegisterProxy) (*MsgRegisterProxyResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) VoteEvent(ctx context.Context, req *MsgVoteEvent) (*MsgVoteEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteEvent not implemented")
}
func (*UnimplementedMsgServer) RegisterProxy(ctx context.Context, req *MsgRegisterProxy) (*MsgRegisterProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterProxy not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_VoteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVoteEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VoteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.event.server.Msg/VoteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VoteEvent(ctx, req.(*MsgVoteEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterProxy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.event.server.Msg/RegisterProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterProxy(ctx, req.(*MsgRegisterProxy))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manythings.mitosis.v1beta1.event.server.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VoteEvent",
			Handler:    _Msg_VoteEvent_Handler,
		},
		{
			MethodName: "RegisterProxy",
			Handler:    _Msg_RegisterProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitosis/event/v1beta1/server/server.proto",
}

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of VoteStatus items.
	VoteStatus(ctx context.Context, in *QueryVoteStatusRequest, opts ...grpc.CallOption) (*QueryVoteStatusResponse, error)
	// Queries a proxy account's address of validator
	Proxy(ctx context.Context, in *QueryProxyRequest, opts ...grpc.CallOption) (*QueryProxyResponse, error)
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

func (c *queryClient) VoteStatus(ctx context.Context, in *QueryVoteStatusRequest, opts ...grpc.CallOption) (*QueryVoteStatusResponse, error) {
	out := new(QueryVoteStatusResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.event.server.Query/VoteStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Proxy(ctx context.Context, in *QueryProxyRequest, opts ...grpc.CallOption) (*QueryProxyResponse, error) {
	out := new(QueryProxyResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.event.server.Query/Proxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of VoteStatus items.
	VoteStatus(context.Context, *QueryVoteStatusRequest) (*QueryVoteStatusResponse, error)
	// Queries a proxy account's address of validator
	Proxy(context.Context, *QueryProxyRequest) (*QueryProxyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) VoteStatus(ctx context.Context, req *QueryVoteStatusRequest) (*QueryVoteStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteStatus not implemented")
}
func (*UnimplementedQueryServer) Proxy(ctx context.Context, req *QueryProxyRequest) (*QueryProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proxy not implemented")
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

func _Query_VoteStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVoteStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VoteStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.event.server.Query/VoteStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VoteStatus(ctx, req.(*QueryVoteStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Proxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.event.server.Query/Proxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proxy(ctx, req.(*QueryProxyRequest))
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
		{
			MethodName: "VoteStatus",
			Handler:    _Query_VoteStatus_Handler,
		},
		{
			MethodName: "Proxy",
			Handler:    _Query_Proxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitosis/event/v1beta1/server/server.proto",
}
