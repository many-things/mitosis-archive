// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitosis/multisig/v1beta1/server/server.proto

package server

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/many-things/mitosis/x/multisig/exported"
	_ "github.com/many-things/mitosis/x/multisig/types"
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
	proto.RegisterFile("mitosis/multisig/v1beta1/server/server.proto", fileDescriptor_39308092bd06c3d7)
}

var fileDescriptor_39308092bd06c3d7 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xcb, 0x6b, 0x13, 0x41,
	0x1c, 0xc7, 0x3b, 0x56, 0x8b, 0x8c, 0xf5, 0xc1, 0xe0, 0x69, 0x91, 0x3d, 0x88, 0x0f, 0xfa, 0xc8,
	0x8e, 0x69, 0x2d, 0x45, 0x4b, 0x2d, 0x6d, 0x51, 0xf1, 0x51, 0xa8, 0xed, 0xcd, 0x4b, 0x99, 0x6d,
	0x86, 0xe9, 0x90, 0xec, 0xce, 0x76, 0x67, 0x36, 0x74, 0x09, 0xb9, 0x78, 0x17, 0x0a, 0x9e, 0xfc,
	0x33, 0x3c, 0x49, 0x8f, 0x42, 0x0f, 0x1e, 0x0b, 0x01, 0xf1, 0x28, 0x89, 0xff, 0x86, 0x20, 0x3b,
	0xfb, 0x48, 0xa8, 0xa6, 0xc9, 0xee, 0x9e, 0xb2, 0x64, 0xe7, 0xf3, 0xdd, 0xef, 0xf7, 0xf7, 0xc8,
	0x06, 0xce, 0x3b, 0x5c, 0x09, 0xc9, 0x25, 0x76, 0x82, 0x86, 0xe2, 0x92, 0x33, 0xdc, 0xac, 0xda,
	0x54, 0x91, 0x2a, 0x96, 0xd4, 0x6f, 0x52, 0x3f, 0xf9, 0xb0, 0x3c, 0x5f, 0x28, 0x81, 0x66, 0x1d,
	0xe2, 0x86, 0xea, 0x80, 0xbb, 0x4c, 0x5a, 0x09, 0x68, 0x25, 0xe7, 0xad, 0x54, 0xc0, 0x8a, 0x09,
	0x63, 0x76, 0x5f, 0x48, 0x47, 0x48, 0x6c, 0x13, 0x49, 0xf1, 0x61, 0x40, 0xfd, 0x30, 0x93, 0xf6,
	0x08, 0xe3, 0x2e, 0x51, 0x5c, 0xb8, 0xb1, 0xae, 0x71, 0x9b, 0x09, 0x26, 0xf4, 0x25, 0x8e, 0xae,
	0x92, 0x6f, 0xef, 0x30, 0x21, 0x58, 0x83, 0x62, 0xe2, 0x71, 0x4c, 0x5c, 0x57, 0x28, 0x8d, 0xc8,
	0xe4, 0x6e, 0x65, 0xa8, 0x73, 0x7a, 0xe4, 0x09, 0x5f, 0xd1, 0x1a, 0x56, 0xa1, 0x47, 0xd3, 0xe3,
	0xf7, 0x87, 0x1e, 0xf7, 0x88, 0x4f, 0x9c, 0xf4, 0x18, 0x1e, 0x55, 0x0f, 0xad, 0xb9, 0xe7, 0x48,
	0x96, 0x00, 0xd5, 0xf1, 0x00, 0x5d, 0x80, 0x18, 0x59, 0x38, 0x9d, 0x84, 0x93, 0x5b, 0x92, 0xa1,
	0x8f, 0x00, 0x5e, 0xdb, 0x55, 0xc4, 0x57, 0x6f, 0x68, 0xc8, 0xa8, 0x8b, 0x9e, 0x5a, 0xe3, 0x97,
	0xd7, 0xda, 0x92, 0x6c, 0x80, 0x35, 0x36, 0x8a, 0xb3, 0x3b, 0x54, 0x7a, 0xc2, 0x95, 0x14, 0x1d,
	0x03, 0x38, 0xbd, 0x1b, 0xd8, 0x0e, 0x57, 0xdb, 0x81, 0x5d, 0xa7, 0x21, 0x5a, 0xc9, 0x2b, 0x3a,
	0x00, 0x1b, 0x9b, 0x25, 0xe0, 0xcc, 0xd2, 0x67, 0x00, 0x6f, 0xc6, 0x37, 0x76, 0x39, 0x73, 0x89,
	0x0a, 0x7c, 0x8a, 0x9e, 0x15, 0x12, 0xce, 0x78, 0xe3, 0x45, 0x39, 0x3e, 0xf5, 0xb6, 0xf0, 0xe7,
	0x3a, 0xbc, 0xf2, 0x2e, 0x6a, 0x2b, 0xfa, 0x0a, 0xe0, 0xd4, 0xb6, 0x9e, 0xa2, 0x7c, 0xe6, 0x34,
	0x1d, 0x83, 0x3b, 0xf4, 0x30, 0xa0, 0x52, 0x19, 0x6b, 0x85, 0xf9, 0xd8, 0xd5, 0xdd, 0xf9, 0x0f,
	0x9d, 0xdf, 0x9f, 0x2e, 0x3d, 0x40, 0xf7, 0x70, 0x24, 0x54, 0x89, 0x95, 0xfe, 0x9d, 0xea, 0x78,
	0xe8, 0xd1, 0x09, 0x80, 0x53, 0xc9, 0xf4, 0x2d, 0xe7, 0x7e, 0x72, 0x32, 0x7a, 0x6b, 0x05, 0xc1,
	0xcc, 0xf2, 0x92, 0xb6, 0x8c, 0x51, 0x65, 0xf8, 0xf2, 0xd5, 0x35, 0x81, 0x5b, 0xfb, 0x07, 0x84,
	0xbb, 0x6d, 0xdc, 0xe2, 0xb5, 0x36, 0xfa, 0x06, 0x20, 0x8c, 0x95, 0xde, 0x72, 0xa9, 0xf2, 0x0d,
	0xeb, 0x80, 0x8d, 0x08, 0xce, 0x37, 0xac, 0xe7, 0xe0, 0x2c, 0x47, 0x55, 0xe7, 0x98, 0x43, 0x33,
	0xa3, 0x72, 0xc8, 0x34, 0x08, 0xea, 0x00, 0x38, 0x9d, 0x55, 0x23, 0x68, 0x28, 0xb4, 0x5a, 0xbc,
	0x98, 0x41, 0x43, 0x19, 0xcf, 0x4b, 0xe1, 0x59, 0x92, 0x55, 0x9d, 0x64, 0x19, 0x2d, 0x0d, 0x4f,
	0xe2, 0xe9, 0x45, 0xc5, 0xad, 0x3a, 0x0d, 0xf7, 0x78, 0xad, 0x8d, 0x5b, 0x4d, 0xd2, 0xe0, 0x35,
	0xa2, 0x84, 0xdf, 0x46, 0x3f, 0x00, 0xbc, 0x35, 0xa8, 0xab, 0xfb, 0xb3, 0x5e, 0xca, 0x9a, 0xee,
	0xd2, 0xab, 0xd2, 0x12, 0x59, 0xc2, 0xc7, 0x3a, 0xa1, 0x85, 0xe6, 0x47, 0x25, 0x4c, 0x7b, 0x15,
	0x85, 0x44, 0x5f, 0x00, 0xbc, 0x1c, 0xfd, 0x10, 0xa0, 0xa5, 0xdc, 0x4e, 0x22, 0xcc, 0x58, 0x2d,
	0x84, 0x65, 0xa6, 0x17, 0xb5, 0xe9, 0x0a, 0x9a, 0xbb, 0xe0, 0x2d, 0xc5, 0xd9, 0xb9, 0x35, 0x39,
	0x01, 0xf0, 0x6a, 0xa4, 0xa2, 0x9b, 0xf0, 0xa4, 0x90, 0x01, 0x5d, 0xfc, 0xf5, 0xc2, 0x68, 0xe6,
	0x1f, 0x6b, 0xff, 0x33, 0xe8, 0xe1, 0xc5, 0xfe, 0xfb, 0xeb, 0x71, 0x0a, 0x20, 0x4c, 0x2a, 0x10,
	0x2d, 0xc7, 0x4a, 0xd1, 0xf2, 0x45, 0xab, 0xb1, 0x59, 0x02, 0xce, 0x33, 0x36, 0x32, 0x7d, 0x51,
	0xe0, 0x96, 0xe4, 0x4c, 0x8f, 0x4d, 0x07, 0xc0, 0x1b, 0x7d, 0x31, 0xdd, 0x88, 0xb5, 0x12, 0x6e,
	0x74, 0x3b, 0x5e, 0x96, 0x14, 0xc8, 0x22, 0x2d, 0xeb, 0x48, 0x55, 0x84, 0xc7, 0x88, 0x34, 0xb8,
	0x0c, 0x1b, 0xaf, 0xbf, 0x77, 0x4d, 0x70, 0xd6, 0x35, 0xc1, 0xaf, 0xae, 0x09, 0x8e, 0x7b, 0xe6,
	0xc4, 0x59, 0xcf, 0x9c, 0xf8, 0xd9, 0x33, 0x27, 0xde, 0x3f, 0x62, 0x5c, 0x1d, 0x04, 0xb6, 0xb5,
	0x2f, 0x9c, 0xff, 0xbe, 0x85, 0x8e, 0xfa, 0x8f, 0x88, 0x0d, 0xda, 0x53, 0xfa, 0x9f, 0xd1, 0xe2,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x54, 0x64, 0xcb, 0x8f, 0x0a, 0x00, 0x00,
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
	StartKeygen(ctx context.Context, in *MsgStartKeygen, opts ...grpc.CallOption) (*MsgStartKeygenResponse, error)
	SubmitPubkey(ctx context.Context, in *MsgSubmitPubkey, opts ...grpc.CallOption) (*MsgSubmitPubkeyResponse, error)
	SubmitSignature(ctx context.Context, in *MsgSubmitSignature, opts ...grpc.CallOption) (*MsgSubmitSignatureResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) StartKeygen(ctx context.Context, in *MsgStartKeygen, opts ...grpc.CallOption) (*MsgStartKeygenResponse, error) {
	out := new(MsgStartKeygenResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Msg/StartKeygen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitPubkey(ctx context.Context, in *MsgSubmitPubkey, opts ...grpc.CallOption) (*MsgSubmitPubkeyResponse, error) {
	out := new(MsgSubmitPubkeyResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Msg/SubmitPubkey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitSignature(ctx context.Context, in *MsgSubmitSignature, opts ...grpc.CallOption) (*MsgSubmitSignatureResponse, error) {
	out := new(MsgSubmitSignatureResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Msg/SubmitSignature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	StartKeygen(context.Context, *MsgStartKeygen) (*MsgStartKeygenResponse, error)
	SubmitPubkey(context.Context, *MsgSubmitPubkey) (*MsgSubmitPubkeyResponse, error)
	SubmitSignature(context.Context, *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) StartKeygen(ctx context.Context, req *MsgStartKeygen) (*MsgStartKeygenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartKeygen not implemented")
}
func (*UnimplementedMsgServer) SubmitPubkey(ctx context.Context, req *MsgSubmitPubkey) (*MsgSubmitPubkeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPubkey not implemented")
}
func (*UnimplementedMsgServer) SubmitSignature(ctx context.Context, req *MsgSubmitSignature) (*MsgSubmitSignatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitSignature not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_StartKeygen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStartKeygen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StartKeygen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Msg/StartKeygen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StartKeygen(ctx, req.(*MsgStartKeygen))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitPubkey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitPubkey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitPubkey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Msg/SubmitPubkey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitPubkey(ctx, req.(*MsgSubmitPubkey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitSignature)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Msg/SubmitSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitSignature(ctx, req.(*MsgSubmitSignature))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manythings.mitosis.v1beta1.multisig.server.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartKeygen",
			Handler:    _Msg_StartKeygen_Handler,
		},
		{
			MethodName: "SubmitPubkey",
			Handler:    _Msg_SubmitPubkey_Handler,
		},
		{
			MethodName: "SubmitSignature",
			Handler:    _Msg_SubmitSignature_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitosis/multisig/v1beta1/server/server.proto",
}

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Keygen query
	Keygen(ctx context.Context, in *QueryKeygen, opts ...grpc.CallOption) (*QueryKeygenResponse, error)
	// Keygen list query
	KeygenList(ctx context.Context, in *QueryKeygenList, opts ...grpc.CallOption) (*QueryKeygenListResponse, error)
	// specific pubkey query
	KeygenResult(ctx context.Context, in *QueryKeygenResult, opts ...grpc.CallOption) (*QueryKeygenResultResponse, error)
	// list pubkey query
	KeygenResultList(ctx context.Context, in *QueryKeygenResultList, opts ...grpc.CallOption) (*QueryKeygenResultListResponse, error)
	// specific sign query
	Sign(ctx context.Context, in *QuerySign, opts ...grpc.CallOption) (*QuerySignResponse, error)
	// list sign query
	SignList(ctx context.Context, in *QuerySignList, opts ...grpc.CallOption) (*QuerySignListResponse, error)
	// specific signature query
	SignResult(ctx context.Context, in *QuerySignResult, opts ...grpc.CallOption) (*QuerySignResultResponse, error)
	// list signature query
	SignResultList(ctx context.Context, in *QuerySignResultList, opts ...grpc.CallOption) (*QuerySignResultListResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Keygen(ctx context.Context, in *QueryKeygen, opts ...grpc.CallOption) (*QueryKeygenResponse, error) {
	out := new(QueryKeygenResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Query/Keygen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) KeygenList(ctx context.Context, in *QueryKeygenList, opts ...grpc.CallOption) (*QueryKeygenListResponse, error) {
	out := new(QueryKeygenListResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Query/KeygenList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) KeygenResult(ctx context.Context, in *QueryKeygenResult, opts ...grpc.CallOption) (*QueryKeygenResultResponse, error) {
	out := new(QueryKeygenResultResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Query/KeygenResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) KeygenResultList(ctx context.Context, in *QueryKeygenResultList, opts ...grpc.CallOption) (*QueryKeygenResultListResponse, error) {
	out := new(QueryKeygenResultListResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Query/KeygenResultList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Sign(ctx context.Context, in *QuerySign, opts ...grpc.CallOption) (*QuerySignResponse, error) {
	out := new(QuerySignResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Query/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SignList(ctx context.Context, in *QuerySignList, opts ...grpc.CallOption) (*QuerySignListResponse, error) {
	out := new(QuerySignListResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Query/SignList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SignResult(ctx context.Context, in *QuerySignResult, opts ...grpc.CallOption) (*QuerySignResultResponse, error) {
	out := new(QuerySignResultResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Query/SignResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SignResultList(ctx context.Context, in *QuerySignResultList, opts ...grpc.CallOption) (*QuerySignResultListResponse, error) {
	out := new(QuerySignResultListResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.v1beta1.multisig.server.Query/SignResultList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Keygen query
	Keygen(context.Context, *QueryKeygen) (*QueryKeygenResponse, error)
	// Keygen list query
	KeygenList(context.Context, *QueryKeygenList) (*QueryKeygenListResponse, error)
	// specific pubkey query
	KeygenResult(context.Context, *QueryKeygenResult) (*QueryKeygenResultResponse, error)
	// list pubkey query
	KeygenResultList(context.Context, *QueryKeygenResultList) (*QueryKeygenResultListResponse, error)
	// specific sign query
	Sign(context.Context, *QuerySign) (*QuerySignResponse, error)
	// list sign query
	SignList(context.Context, *QuerySignList) (*QuerySignListResponse, error)
	// specific signature query
	SignResult(context.Context, *QuerySignResult) (*QuerySignResultResponse, error)
	// list signature query
	SignResultList(context.Context, *QuerySignResultList) (*QuerySignResultListResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Keygen(ctx context.Context, req *QueryKeygen) (*QueryKeygenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Keygen not implemented")
}
func (*UnimplementedQueryServer) KeygenList(ctx context.Context, req *QueryKeygenList) (*QueryKeygenListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeygenList not implemented")
}
func (*UnimplementedQueryServer) KeygenResult(ctx context.Context, req *QueryKeygenResult) (*QueryKeygenResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeygenResult not implemented")
}
func (*UnimplementedQueryServer) KeygenResultList(ctx context.Context, req *QueryKeygenResultList) (*QueryKeygenResultListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeygenResultList not implemented")
}
func (*UnimplementedQueryServer) Sign(ctx context.Context, req *QuerySign) (*QuerySignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (*UnimplementedQueryServer) SignList(ctx context.Context, req *QuerySignList) (*QuerySignListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignList not implemented")
}
func (*UnimplementedQueryServer) SignResult(ctx context.Context, req *QuerySignResult) (*QuerySignResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignResult not implemented")
}
func (*UnimplementedQueryServer) SignResultList(ctx context.Context, req *QuerySignResultList) (*QuerySignResultListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignResultList not implemented")
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
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Keygen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryKeygen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Keygen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Query/Keygen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Keygen(ctx, req.(*QueryKeygen))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_KeygenList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryKeygenList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).KeygenList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Query/KeygenList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).KeygenList(ctx, req.(*QueryKeygenList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_KeygenResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryKeygenResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).KeygenResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Query/KeygenResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).KeygenResult(ctx, req.(*QueryKeygenResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_KeygenResultList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryKeygenResultList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).KeygenResultList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Query/KeygenResultList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).KeygenResultList(ctx, req.(*QueryKeygenResultList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySign)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Query/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Sign(ctx, req.(*QuerySign))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SignList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySignList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SignList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Query/SignList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SignList(ctx, req.(*QuerySignList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SignResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySignResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SignResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Query/SignResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SignResult(ctx, req.(*QuerySignResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SignResultList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySignResultList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SignResultList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.v1beta1.multisig.server.Query/SignResultList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SignResultList(ctx, req.(*QuerySignResultList))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manythings.mitosis.v1beta1.multisig.server.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Keygen",
			Handler:    _Query_Keygen_Handler,
		},
		{
			MethodName: "KeygenList",
			Handler:    _Query_KeygenList_Handler,
		},
		{
			MethodName: "KeygenResult",
			Handler:    _Query_KeygenResult_Handler,
		},
		{
			MethodName: "KeygenResultList",
			Handler:    _Query_KeygenResultList_Handler,
		},
		{
			MethodName: "Sign",
			Handler:    _Query_Sign_Handler,
		},
		{
			MethodName: "SignList",
			Handler:    _Query_SignList_Handler,
		},
		{
			MethodName: "SignResult",
			Handler:    _Query_SignResult_Handler,
		},
		{
			MethodName: "SignResultList",
			Handler:    _Query_SignResultList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitosis/multisig/v1beta1/server/server.proto",
}
