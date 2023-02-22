// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitosis/context/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae61d901f635151, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae61d901f635151, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryContextByAddressRequest struct {
}

func (m *QueryContextByAddressRequest) Reset()         { *m = QueryContextByAddressRequest{} }
func (m *QueryContextByAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContextByAddressRequest) ProtoMessage()    {}
func (*QueryContextByAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae61d901f635151, []int{2}
}
func (m *QueryContextByAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContextByAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContextByAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContextByAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContextByAddressRequest.Merge(m, src)
}
func (m *QueryContextByAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContextByAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContextByAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContextByAddressRequest proto.InternalMessageInfo

type QueryContextByAddressResponse struct {
}

func (m *QueryContextByAddressResponse) Reset()         { *m = QueryContextByAddressResponse{} }
func (m *QueryContextByAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContextByAddressResponse) ProtoMessage()    {}
func (*QueryContextByAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae61d901f635151, []int{3}
}
func (m *QueryContextByAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContextByAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContextByAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContextByAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContextByAddressResponse.Merge(m, src)
}
func (m *QueryContextByAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContextByAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContextByAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContextByAddressResponse proto.InternalMessageInfo

type QueryContextsByAddressRequest struct {
}

func (m *QueryContextsByAddressRequest) Reset()         { *m = QueryContextsByAddressRequest{} }
func (m *QueryContextsByAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContextsByAddressRequest) ProtoMessage()    {}
func (*QueryContextsByAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae61d901f635151, []int{4}
}
func (m *QueryContextsByAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContextsByAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContextsByAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContextsByAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContextsByAddressRequest.Merge(m, src)
}
func (m *QueryContextsByAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContextsByAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContextsByAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContextsByAddressRequest proto.InternalMessageInfo

type QueryContextsByAddressResponse struct {
}

func (m *QueryContextsByAddressResponse) Reset()         { *m = QueryContextsByAddressResponse{} }
func (m *QueryContextsByAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContextsByAddressResponse) ProtoMessage()    {}
func (*QueryContextsByAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae61d901f635151, []int{5}
}
func (m *QueryContextsByAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContextsByAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContextsByAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContextsByAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContextsByAddressResponse.Merge(m, src)
}
func (m *QueryContextsByAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContextsByAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContextsByAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContextsByAddressResponse proto.InternalMessageInfo

type QueryContextByTxHashRequest struct {
}

func (m *QueryContextByTxHashRequest) Reset()         { *m = QueryContextByTxHashRequest{} }
func (m *QueryContextByTxHashRequest) String() string { return proto.CompactTextString(m) }
func (*QueryContextByTxHashRequest) ProtoMessage()    {}
func (*QueryContextByTxHashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae61d901f635151, []int{6}
}
func (m *QueryContextByTxHashRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContextByTxHashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContextByTxHashRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContextByTxHashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContextByTxHashRequest.Merge(m, src)
}
func (m *QueryContextByTxHashRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryContextByTxHashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContextByTxHashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContextByTxHashRequest proto.InternalMessageInfo

type QueryContextByTxHashResponse struct {
}

func (m *QueryContextByTxHashResponse) Reset()         { *m = QueryContextByTxHashResponse{} }
func (m *QueryContextByTxHashResponse) String() string { return proto.CompactTextString(m) }
func (*QueryContextByTxHashResponse) ProtoMessage()    {}
func (*QueryContextByTxHashResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ae61d901f635151, []int{7}
}
func (m *QueryContextByTxHashResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryContextByTxHashResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryContextByTxHashResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryContextByTxHashResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryContextByTxHashResponse.Merge(m, src)
}
func (m *QueryContextByTxHashResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryContextByTxHashResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryContextByTxHashResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryContextByTxHashResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "manythings.mitosis.context.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "manythings.mitosis.context.QueryParamsResponse")
	proto.RegisterType((*QueryContextByAddressRequest)(nil), "manythings.mitosis.context.QueryContextByAddressRequest")
	proto.RegisterType((*QueryContextByAddressResponse)(nil), "manythings.mitosis.context.QueryContextByAddressResponse")
	proto.RegisterType((*QueryContextsByAddressRequest)(nil), "manythings.mitosis.context.QueryContextsByAddressRequest")
	proto.RegisterType((*QueryContextsByAddressResponse)(nil), "manythings.mitosis.context.QueryContextsByAddressResponse")
	proto.RegisterType((*QueryContextByTxHashRequest)(nil), "manythings.mitosis.context.QueryContextByTxHashRequest")
	proto.RegisterType((*QueryContextByTxHashResponse)(nil), "manythings.mitosis.context.QueryContextByTxHashResponse")
}

func init() { proto.RegisterFile("mitosis/context/query.proto", fileDescriptor_6ae61d901f635151) }

var fileDescriptor_6ae61d901f635151 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x41, 0x6b, 0x14, 0x31,
	0x14, 0xde, 0x88, 0xdd, 0x43, 0x3c, 0xa8, 0xb1, 0x07, 0x99, 0xb6, 0x69, 0x59, 0x11, 0x44, 0x31,
	0x71, 0xd7, 0x43, 0x57, 0x4f, 0xba, 0x5e, 0x7a, 0xd4, 0x22, 0x08, 0x5e, 0x96, 0xcc, 0x36, 0xcc,
	0x0c, 0x38, 0xc9, 0x74, 0x5e, 0x56, 0x66, 0xae, 0xfe, 0x02, 0x41, 0xf0, 0xdf, 0x08, 0x1e, 0x3c,
	0xf4, 0x58, 0xf0, 0xe2, 0x49, 0x64, 0xd7, 0x93, 0xbf, 0x42, 0x36, 0xc9, 0x16, 0x77, 0xa6, 0x53,
	0xa7, 0x7b, 0x9a, 0x90, 0xf7, 0x7d, 0xdf, 0xfb, 0xbe, 0xf7, 0xc2, 0xe0, 0xad, 0x34, 0x31, 0x1a,
	0x12, 0xe0, 0x13, 0xad, 0x8c, 0x2c, 0x0c, 0x3f, 0x9e, 0xca, 0xbc, 0x64, 0x59, 0xae, 0x8d, 0x26,
	0x41, 0x2a, 0x54, 0x69, 0xe2, 0x44, 0x45, 0xc0, 0x3c, 0x8e, 0x79, 0x5c, 0x70, 0x7f, 0xa2, 0x21,
	0xd5, 0xc0, 0x43, 0x01, 0xd2, 0x91, 0xf8, 0xfb, 0x7e, 0x28, 0x8d, 0xe8, 0xf3, 0x4c, 0x44, 0x89,
	0x12, 0x26, 0xd1, 0xca, 0xe9, 0x04, 0x9b, 0x91, 0x8e, 0xb4, 0x3d, 0xf2, 0xc5, 0xc9, 0xdf, 0x6e,
	0x47, 0x5a, 0x47, 0xef, 0x24, 0x17, 0x59, 0xc2, 0x85, 0x52, 0xda, 0x58, 0x0a, 0x2c, 0xab, 0x55,
	0x63, 0x99, 0xc8, 0x45, 0xea, 0xab, 0xbd, 0x4d, 0x4c, 0x5e, 0x2d, 0x7a, 0xbe, 0xb4, 0x97, 0x87,
	0xf2, 0x78, 0x2a, 0xc1, 0xf4, 0xde, 0xe0, 0x5b, 0x2b, 0xb7, 0x90, 0x69, 0x05, 0x92, 0x3c, 0xc3,
	0x5d, 0x47, 0xbe, 0x8d, 0xf6, 0xd0, 0xbd, 0x6b, 0x83, 0x1e, 0x6b, 0xce, 0xc5, 0x1c, 0x77, 0x74,
	0xf5, 0xe4, 0xe7, 0x6e, 0xe7, 0xd0, 0xf3, 0x7a, 0x14, 0x6f, 0x5b, 0xe1, 0x17, 0x0e, 0x34, 0x2a,
	0x9f, 0x1f, 0x1d, 0xe5, 0x12, 0xce, 0x1a, 0xef, 0xe2, 0x9d, 0x86, 0xba, 0xb3, 0x50, 0x05, 0x40,
	0x4d, 0x61, 0x0f, 0xd3, 0x26, 0x80, 0x97, 0xd8, 0xc1, 0x5b, 0xab, 0x3d, 0x5e, 0x17, 0x07, 0x02,
	0xe2, 0xa5, 0x40, 0xcd, 0xe2, 0xb2, 0xec, 0xe8, 0x83, 0x3f, 0x1b, 0x78, 0xc3, 0x02, 0xc8, 0x67,
	0x84, 0xbb, 0x2e, 0x25, 0x61, 0x17, 0x4d, 0xa2, 0x3e, 0xe0, 0x80, 0xb7, 0xc6, 0x7b, 0xd3, 0x0f,
	0x3e, 0x7c, 0xff, 0xfd, 0xe9, 0xca, 0x5d, 0x72, 0x87, 0x2f, 0x88, 0x0f, 0x1d, 0x93, 0x9f, 0xbf,
	0x5a, 0xf2, 0x15, 0xe1, 0x1b, 0xd5, 0x09, 0x92, 0xe1, 0x7f, 0x5b, 0x36, 0x2c, 0x25, 0x78, 0xb2,
	0x06, 0xd3, 0xdb, 0xde, 0xb7, 0xb6, 0xfb, 0x84, 0x5f, 0x68, 0xdb, 0x7f, 0xc7, 0x61, 0x39, 0x16,
	0xde, 0xed, 0x37, 0x84, 0x6f, 0xd6, 0x56, 0x48, 0x5a, 0x3b, 0xa9, 0xbd, 0x8b, 0xe0, 0xe9, 0x3a,
	0x54, 0x9f, 0x62, 0x68, 0x53, 0x0c, 0xc8, 0xa3, 0x36, 0x29, 0xe0, 0xdf, 0x18, 0x5f, 0x10, 0xbe,
	0x5e, 0x79, 0x48, 0x64, 0xbf, 0xfd, 0x38, 0x57, 0x5e, 0x66, 0x30, 0xbc, 0x3c, 0x71, 0xdd, 0x35,
	0x98, 0x62, 0x1c, 0x0b, 0x88, 0x47, 0x07, 0x27, 0x33, 0x8a, 0x4e, 0x67, 0x14, 0xfd, 0x9a, 0x51,
	0xf4, 0x71, 0x4e, 0x3b, 0xa7, 0x73, 0xda, 0xf9, 0x31, 0xa7, 0x9d, 0xb7, 0x2c, 0x4a, 0x4c, 0x3c,
	0x0d, 0xd9, 0x44, 0xa7, 0xe7, 0x8a, 0x16, 0x67, 0xb2, 0xa6, 0xcc, 0x24, 0x84, 0x5d, 0xfb, 0xbf,
	0x79, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x3c, 0x27, 0xb1, 0x28, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of ContextByAddress items.
	ContextByAddress(ctx context.Context, in *QueryContextByAddressRequest, opts ...grpc.CallOption) (*QueryContextByAddressResponse, error)
	// Queries a list of ContextsByAddress items.
	ContextsByAddress(ctx context.Context, in *QueryContextsByAddressRequest, opts ...grpc.CallOption) (*QueryContextsByAddressResponse, error)
	// Queries a list of ContextByTxHash items.
	ContextByTxHash(ctx context.Context, in *QueryContextByTxHashRequest, opts ...grpc.CallOption) (*QueryContextByTxHashResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.context.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContextByAddress(ctx context.Context, in *QueryContextByAddressRequest, opts ...grpc.CallOption) (*QueryContextByAddressResponse, error) {
	out := new(QueryContextByAddressResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.context.Query/ContextByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContextsByAddress(ctx context.Context, in *QueryContextsByAddressRequest, opts ...grpc.CallOption) (*QueryContextsByAddressResponse, error) {
	out := new(QueryContextsByAddressResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.context.Query/ContextsByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContextByTxHash(ctx context.Context, in *QueryContextByTxHashRequest, opts ...grpc.CallOption) (*QueryContextByTxHashResponse, error) {
	out := new(QueryContextByTxHashResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.context.Query/ContextByTxHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ContextByAddress items.
	ContextByAddress(context.Context, *QueryContextByAddressRequest) (*QueryContextByAddressResponse, error)
	// Queries a list of ContextsByAddress items.
	ContextsByAddress(context.Context, *QueryContextsByAddressRequest) (*QueryContextsByAddressResponse, error)
	// Queries a list of ContextByTxHash items.
	ContextByTxHash(context.Context, *QueryContextByTxHashRequest) (*QueryContextByTxHashResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ContextByAddress(ctx context.Context, req *QueryContextByAddressRequest) (*QueryContextByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContextByAddress not implemented")
}
func (*UnimplementedQueryServer) ContextsByAddress(ctx context.Context, req *QueryContextsByAddressRequest) (*QueryContextsByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContextsByAddress not implemented")
}
func (*UnimplementedQueryServer) ContextByTxHash(ctx context.Context, req *QueryContextByTxHashRequest) (*QueryContextByTxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContextByTxHash not implemented")
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
		FullMethod: "/manythings.mitosis.context.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContextByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContextByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContextByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.context.Query/ContextByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContextByAddress(ctx, req.(*QueryContextByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContextsByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContextsByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContextsByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.context.Query/ContextsByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContextsByAddress(ctx, req.(*QueryContextsByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContextByTxHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContextByTxHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContextByTxHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manythings.mitosis.context.Query/ContextByTxHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContextByTxHash(ctx, req.(*QueryContextByTxHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manythings.mitosis.context.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ContextByAddress",
			Handler:    _Query_ContextByAddress_Handler,
		},
		{
			MethodName: "ContextsByAddress",
			Handler:    _Query_ContextsByAddress_Handler,
		},
		{
			MethodName: "ContextByTxHash",
			Handler:    _Query_ContextByTxHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitosis/context/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryContextByAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContextByAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContextByAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryContextByAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContextByAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContextByAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryContextsByAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContextsByAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContextsByAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryContextsByAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContextsByAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContextsByAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryContextByTxHashRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContextByTxHashRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContextByTxHashRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryContextByTxHashResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryContextByTxHashResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryContextByTxHashResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryContextByAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryContextByAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryContextsByAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryContextsByAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryContextByTxHashRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryContextByTxHashResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryContextByAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContextByAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContextByAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryContextByAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContextByAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContextByAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryContextsByAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContextsByAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContextsByAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryContextsByAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContextsByAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContextsByAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryContextByTxHashRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContextByTxHashRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContextByTxHashRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryContextByTxHashResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryContextByTxHashResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryContextByTxHashResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
