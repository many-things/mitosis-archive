// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitosis/multisig/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgStartKeygen struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgStartKeygen) Reset()         { *m = MsgStartKeygen{} }
func (m *MsgStartKeygen) String() string { return proto.CompactTextString(m) }
func (*MsgStartKeygen) ProtoMessage()    {}
func (*MsgStartKeygen) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5960799c442c59, []int{0}
}
func (m *MsgStartKeygen) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStartKeygen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStartKeygen.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStartKeygen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStartKeygen.Merge(m, src)
}
func (m *MsgStartKeygen) XXX_Size() int {
	return m.Size()
}
func (m *MsgStartKeygen) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStartKeygen.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStartKeygen proto.InternalMessageInfo

func (m *MsgStartKeygen) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgStartKeygenResponse struct {
}

func (m *MsgStartKeygenResponse) Reset()         { *m = MsgStartKeygenResponse{} }
func (m *MsgStartKeygenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStartKeygenResponse) ProtoMessage()    {}
func (*MsgStartKeygenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5960799c442c59, []int{1}
}
func (m *MsgStartKeygenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStartKeygenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStartKeygenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStartKeygenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStartKeygenResponse.Merge(m, src)
}
func (m *MsgStartKeygenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStartKeygenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStartKeygenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStartKeygenResponse proto.InternalMessageInfo

type MsgSubmitPubkey struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgSubmitPubkey) Reset()         { *m = MsgSubmitPubkey{} }
func (m *MsgSubmitPubkey) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitPubkey) ProtoMessage()    {}
func (*MsgSubmitPubkey) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5960799c442c59, []int{2}
}
func (m *MsgSubmitPubkey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitPubkey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitPubkey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitPubkey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitPubkey.Merge(m, src)
}
func (m *MsgSubmitPubkey) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitPubkey) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitPubkey.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitPubkey proto.InternalMessageInfo

func (m *MsgSubmitPubkey) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgSubmitPubkeyResponse struct {
}

func (m *MsgSubmitPubkeyResponse) Reset()         { *m = MsgSubmitPubkeyResponse{} }
func (m *MsgSubmitPubkeyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitPubkeyResponse) ProtoMessage()    {}
func (*MsgSubmitPubkeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5960799c442c59, []int{3}
}
func (m *MsgSubmitPubkeyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitPubkeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitPubkeyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitPubkeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitPubkeyResponse.Merge(m, src)
}
func (m *MsgSubmitPubkeyResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitPubkeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitPubkeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitPubkeyResponse proto.InternalMessageInfo

type MsgSubmitSignature struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgSubmitSignature) Reset()         { *m = MsgSubmitSignature{} }
func (m *MsgSubmitSignature) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitSignature) ProtoMessage()    {}
func (*MsgSubmitSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5960799c442c59, []int{4}
}
func (m *MsgSubmitSignature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitSignature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitSignature.Merge(m, src)
}
func (m *MsgSubmitSignature) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitSignature.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitSignature proto.InternalMessageInfo

func (m *MsgSubmitSignature) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgSubmitSignatureResponse struct {
}

func (m *MsgSubmitSignatureResponse) Reset()         { *m = MsgSubmitSignatureResponse{} }
func (m *MsgSubmitSignatureResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitSignatureResponse) ProtoMessage()    {}
func (*MsgSubmitSignatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5960799c442c59, []int{5}
}
func (m *MsgSubmitSignatureResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitSignatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitSignatureResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitSignatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitSignatureResponse.Merge(m, src)
}
func (m *MsgSubmitSignatureResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitSignatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitSignatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitSignatureResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgStartKeygen)(nil), "manythings.mitosis.multisig.MsgStartKeygen")
	proto.RegisterType((*MsgStartKeygenResponse)(nil), "manythings.mitosis.multisig.MsgStartKeygenResponse")
	proto.RegisterType((*MsgSubmitPubkey)(nil), "manythings.mitosis.multisig.MsgSubmitPubkey")
	proto.RegisterType((*MsgSubmitPubkeyResponse)(nil), "manythings.mitosis.multisig.MsgSubmitPubkeyResponse")
	proto.RegisterType((*MsgSubmitSignature)(nil), "manythings.mitosis.multisig.MsgSubmitSignature")
	proto.RegisterType((*MsgSubmitSignatureResponse)(nil), "manythings.mitosis.multisig.MsgSubmitSignatureResponse")
}

func init() { proto.RegisterFile("mitosis/multisig/tx.proto", fileDescriptor_6c5960799c442c59) }

var fileDescriptor_6c5960799c442c59 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x57, 0x05, 0xc5, 0xa7, 0x38, 0xe8, 0x41, 0xbb, 0x2a, 0x41, 0x7a, 0x12, 0xa7, 0x09,
	0x38, 0xc1, 0xbb, 0x37, 0x91, 0x81, 0x6c, 0x37, 0x6f, 0xed, 0x08, 0x59, 0xd0, 0x36, 0x25, 0x79,
	0x85, 0x15, 0xbf, 0x84, 0x1f, 0xcb, 0xe3, 0x8e, 0x1e, 0xa5, 0x3d, 0xf8, 0x35, 0x64, 0x6a, 0x4a,
	0x3b, 0xa1, 0xda, 0x63, 0xc8, 0xef, 0xfd, 0x1e, 0xef, 0xcf, 0x1f, 0x06, 0xb1, 0x44, 0x65, 0xa4,
	0x61, 0x71, 0xf6, 0x84, 0xd2, 0x48, 0xc1, 0x70, 0x41, 0x53, 0xad, 0x50, 0xb9, 0x47, 0x71, 0x98,
	0xe4, 0x38, 0x97, 0x89, 0x30, 0xf4, 0x87, 0xa2, 0x96, 0x0a, 0xce, 0x60, 0x7f, 0x6c, 0xc4, 0x14,
	0x43, 0x8d, 0x77, 0x3c, 0x17, 0x3c, 0x71, 0x3d, 0xd8, 0x9e, 0x69, 0x1e, 0xa2, 0xd2, 0x9e, 0x73,
	0xe2, 0x9c, 0xee, 0x4c, 0xec, 0x33, 0xf0, 0xe0, 0xa0, 0xc9, 0x4e, 0xb8, 0x49, 0x55, 0x62, 0x78,
	0x30, 0x84, 0xfe, 0xea, 0x27, 0x8b, 0x62, 0x89, 0xf7, 0x59, 0xf4, 0xc8, 0xf3, 0x16, 0xcd, 0x00,
	0x0e, 0xd7, 0xe0, 0xca, 0x43, 0xc1, 0xad, 0xbe, 0xa6, 0x52, 0x24, 0x21, 0x66, 0x9a, 0xb7, 0xa8,
	0x8e, 0xc1, 0xff, 0xcd, 0x5b, 0xdb, 0xe5, 0xc7, 0x06, 0x6c, 0x8e, 0x8d, 0x70, 0x15, 0xec, 0xd6,
	0x0f, 0x1c, 0xd2, 0x96, 0x40, 0x68, 0xf3, 0x42, 0x7f, 0xd4, 0x01, 0xb6, 0x8b, 0x5d, 0x0d, 0x7b,
	0x8d, 0x2c, 0xce, 0xff, 0x94, 0xd4, 0x68, 0xff, 0xaa, 0x0b, 0x5d, 0xed, 0x7c, 0x86, 0xfe, 0x7a,
	0x6e, 0xec, 0x7f, 0xa2, 0x6a, 0xc0, 0xbf, 0xee, 0x38, 0x60, 0x97, 0xdf, 0xdc, 0xbe, 0x16, 0xc4,
	0x59, 0x16, 0xc4, 0x79, 0x2f, 0x88, 0xf3, 0x52, 0x92, 0xde, 0xb2, 0x24, 0xbd, 0xb7, 0x92, 0xf4,
	0x1e, 0x98, 0x90, 0x38, 0xcf, 0x22, 0x3a, 0x53, 0x31, 0x5b, 0xc9, 0x2f, 0xbe, 0xed, 0xcc, 0xd6,
	0x75, 0x51, 0x2b, 0x6c, 0x9e, 0x72, 0x13, 0x6d, 0x7d, 0x95, 0x76, 0xf4, 0x19, 0x00, 0x00, 0xff,
	0xff, 0x8b, 0x22, 0x45, 0xd6, 0xd1, 0x02, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/manythings.mitosis.multisig.Msg/StartKeygen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitPubkey(ctx context.Context, in *MsgSubmitPubkey, opts ...grpc.CallOption) (*MsgSubmitPubkeyResponse, error) {
	out := new(MsgSubmitPubkeyResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.multisig.Msg/SubmitPubkey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitSignature(ctx context.Context, in *MsgSubmitSignature, opts ...grpc.CallOption) (*MsgSubmitSignatureResponse, error) {
	out := new(MsgSubmitSignatureResponse)
	err := c.cc.Invoke(ctx, "/manythings.mitosis.multisig.Msg/SubmitSignature", in, out, opts...)
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
		FullMethod: "/manythings.mitosis.multisig.Msg/StartKeygen",
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
		FullMethod: "/manythings.mitosis.multisig.Msg/SubmitPubkey",
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
		FullMethod: "/manythings.mitosis.multisig.Msg/SubmitSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitSignature(ctx, req.(*MsgSubmitSignature))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "manythings.mitosis.multisig.Msg",
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
	Metadata: "mitosis/multisig/tx.proto",
}

func (m *MsgStartKeygen) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStartKeygen) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStartKeygen) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStartKeygenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStartKeygenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStartKeygenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitPubkey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitPubkey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitPubkey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitPubkeyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitPubkeyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitPubkeyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitSignature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitSignature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitSignature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitSignatureResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitSignatureResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitSignatureResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgStartKeygen) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgStartKeygenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitPubkey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitPubkeyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitSignature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitSignatureResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgStartKeygen) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStartKeygen: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStartKeygen: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStartKeygenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStartKeygenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStartKeygenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitPubkey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitPubkey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitPubkey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitPubkeyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitPubkeyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitPubkeyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitSignature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitSignature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitSignature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitSignatureResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitSignatureResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitSignatureResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)