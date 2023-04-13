// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitosis/multisig/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	exported "github.com/many-things/mitosis/x/multisig/exported"
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

// GenesisKeygen defines the multisig keygen genesis state.
type GenesisKeygen struct {
	ChainSet []*GenesisKeygen_ChainSet `protobuf:"bytes,1,rep,name=chain_set,json=chainSet,proto3" json:"chain_set,omitempty"`
}

func (m *GenesisKeygen) Reset()         { *m = GenesisKeygen{} }
func (m *GenesisKeygen) String() string { return proto.CompactTextString(m) }
func (*GenesisKeygen) ProtoMessage()    {}
func (*GenesisKeygen) Descriptor() ([]byte, []int) {
	return fileDescriptor_455e9f637adacbd8, []int{0}
}
func (m *GenesisKeygen) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisKeygen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisKeygen.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisKeygen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisKeygen.Merge(m, src)
}
func (m *GenesisKeygen) XXX_Size() int {
	return m.Size()
}
func (m *GenesisKeygen) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisKeygen.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisKeygen proto.InternalMessageInfo

func (m *GenesisKeygen) GetChainSet() []*GenesisKeygen_ChainSet {
	if m != nil {
		return m.ChainSet
	}
	return nil
}

// Key-Value set of chain
type GenesisKeygen_ChainSet struct {
	Chain     string          `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	LastId    uint64          `protobuf:"varint,2,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	ItemSet   []*Keygen       `protobuf:"bytes,3,rep,name=item_set,json=itemSet,proto3" json:"item_set,omitempty"`
	ResultSet []*KeygenResult `protobuf:"bytes,4,rep,name=result_set,json=resultSet,proto3" json:"result_set,omitempty"`
}

func (m *GenesisKeygen_ChainSet) Reset()         { *m = GenesisKeygen_ChainSet{} }
func (m *GenesisKeygen_ChainSet) String() string { return proto.CompactTextString(m) }
func (*GenesisKeygen_ChainSet) ProtoMessage()    {}
func (*GenesisKeygen_ChainSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_455e9f637adacbd8, []int{0, 0}
}
func (m *GenesisKeygen_ChainSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisKeygen_ChainSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisKeygen_ChainSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisKeygen_ChainSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisKeygen_ChainSet.Merge(m, src)
}
func (m *GenesisKeygen_ChainSet) XXX_Size() int {
	return m.Size()
}
func (m *GenesisKeygen_ChainSet) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisKeygen_ChainSet.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisKeygen_ChainSet proto.InternalMessageInfo

func (m *GenesisKeygen_ChainSet) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *GenesisKeygen_ChainSet) GetLastId() uint64 {
	if m != nil {
		return m.LastId
	}
	return 0
}

func (m *GenesisKeygen_ChainSet) GetItemSet() []*Keygen {
	if m != nil {
		return m.ItemSet
	}
	return nil
}

func (m *GenesisKeygen_ChainSet) GetResultSet() []*KeygenResult {
	if m != nil {
		return m.ResultSet
	}
	return nil
}

// GenesisSign define the multisig sign genesis state.
type GenesisSign struct {
	ChainSet []*GenesisSign_ChainSet `protobuf:"bytes,1,rep,name=chain_set,json=chainSet,proto3" json:"chain_set,omitempty"`
}

func (m *GenesisSign) Reset()         { *m = GenesisSign{} }
func (m *GenesisSign) String() string { return proto.CompactTextString(m) }
func (*GenesisSign) ProtoMessage()    {}
func (*GenesisSign) Descriptor() ([]byte, []int) {
	return fileDescriptor_455e9f637adacbd8, []int{1}
}
func (m *GenesisSign) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisSign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisSign.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisSign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisSign.Merge(m, src)
}
func (m *GenesisSign) XXX_Size() int {
	return m.Size()
}
func (m *GenesisSign) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisSign.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisSign proto.InternalMessageInfo

func (m *GenesisSign) GetChainSet() []*GenesisSign_ChainSet {
	if m != nil {
		return m.ChainSet
	}
	return nil
}

// Key-Value set of item
type GenesisSign_ChainSet struct {
	Chain     string                 `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	LastId    uint64                 `protobuf:"varint,2,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	ItemSet   []*exported.Sign       `protobuf:"bytes,3,rep,name=item_set,json=itemSet,proto3" json:"item_set,omitempty"`
	ResultSet []*exported.SignResult `protobuf:"bytes,4,rep,name=result_set,json=resultSet,proto3" json:"result_set,omitempty"`
}

func (m *GenesisSign_ChainSet) Reset()         { *m = GenesisSign_ChainSet{} }
func (m *GenesisSign_ChainSet) String() string { return proto.CompactTextString(m) }
func (*GenesisSign_ChainSet) ProtoMessage()    {}
func (*GenesisSign_ChainSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_455e9f637adacbd8, []int{1, 0}
}
func (m *GenesisSign_ChainSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisSign_ChainSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisSign_ChainSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisSign_ChainSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisSign_ChainSet.Merge(m, src)
}
func (m *GenesisSign_ChainSet) XXX_Size() int {
	return m.Size()
}
func (m *GenesisSign_ChainSet) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisSign_ChainSet.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisSign_ChainSet proto.InternalMessageInfo

func (m *GenesisSign_ChainSet) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *GenesisSign_ChainSet) GetLastId() uint64 {
	if m != nil {
		return m.LastId
	}
	return 0
}

func (m *GenesisSign_ChainSet) GetItemSet() []*exported.Sign {
	if m != nil {
		return m.ItemSet
	}
	return nil
}

func (m *GenesisSign_ChainSet) GetResultSet() []*exported.SignResult {
	if m != nil {
		return m.ResultSet
	}
	return nil
}

// GenesisState defines the multisig module's genesis state.
type GenesisState struct {
	Params Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Keygen *GenesisKeygen `protobuf:"bytes,2,opt,name=keygen,proto3" json:"keygen,omitempty"`
	Sign   *GenesisSign   `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_455e9f637adacbd8, []int{2}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetKeygen() *GenesisKeygen {
	if m != nil {
		return m.Keygen
	}
	return nil
}

func (m *GenesisState) GetSign() *GenesisSign {
	if m != nil {
		return m.Sign
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisKeygen)(nil), "manythings.mitosis.v1beta1.multisig.GenesisKeygen")
	proto.RegisterType((*GenesisKeygen_ChainSet)(nil), "manythings.mitosis.v1beta1.multisig.GenesisKeygen.ChainSet")
	proto.RegisterType((*GenesisSign)(nil), "manythings.mitosis.v1beta1.multisig.GenesisSign")
	proto.RegisterType((*GenesisSign_ChainSet)(nil), "manythings.mitosis.v1beta1.multisig.GenesisSign.ChainSet")
	proto.RegisterType((*GenesisState)(nil), "manythings.mitosis.v1beta1.multisig.GenesisState")
}

func init() {
	proto.RegisterFile("mitosis/multisig/v1beta1/genesis.proto", fileDescriptor_455e9f637adacbd8)
}

var fileDescriptor_455e9f637adacbd8 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4f, 0x8b, 0x13, 0x31,
	0x18, 0xc6, 0x27, 0xed, 0x38, 0xdb, 0x66, 0xf5, 0x12, 0x16, 0x1c, 0xe6, 0x30, 0x96, 0xf5, 0x0f,
	0x05, 0xd9, 0xc4, 0xae, 0x17, 0xc5, 0xdb, 0x2a, 0x4a, 0x15, 0x61, 0x49, 0x41, 0xc5, 0xcb, 0x32,
	0x6d, 0x43, 0x1a, 0xec, 0xfc, 0x61, 0xf2, 0x56, 0xb6, 0xdf, 0xc2, 0x6f, 0xe4, 0x75, 0xbd, 0x2d,
	0x78, 0xf1, 0x24, 0xda, 0x1e, 0xfd, 0x12, 0x92, 0xcc, 0x0c, 0x6b, 0x69, 0x0b, 0x53, 0xf7, 0x96,
	0xcc, 0xbc, 0xcf, 0x2f, 0xcf, 0xfb, 0x24, 0x2f, 0x7e, 0x10, 0x2b, 0x48, 0xb5, 0xd2, 0x2c, 0x9e,
	0x4d, 0x41, 0x69, 0x25, 0xd9, 0xe7, 0xde, 0x50, 0x40, 0xd4, 0x63, 0x52, 0x24, 0x42, 0x2b, 0x4d,
	0xb3, 0x3c, 0x85, 0x94, 0xdc, 0x8d, 0xa3, 0x64, 0x0e, 0x13, 0x95, 0x48, 0x4d, 0x4b, 0x09, 0x2d,
	0x2b, 0x69, 0x25, 0x0d, 0x0e, 0x64, 0x2a, 0x53, 0x5b, 0xcf, 0xcc, 0xaa, 0x90, 0x06, 0x47, 0x5b,
	0x8f, 0x10, 0xe7, 0x59, 0x9a, 0x83, 0x18, 0x33, 0x98, 0x67, 0xa2, 0x3c, 0x29, 0xb8, 0xbf, 0xb5,
	0x3c, 0x8b, 0xf2, 0x28, 0xae, 0xca, 0xee, 0x6d, 0x2d, 0xfb, 0x07, 0x76, 0xf8, 0xb5, 0x81, 0x6f,
	0xbd, 0x2a, 0x1a, 0x79, 0x23, 0xe6, 0x52, 0x24, 0xe4, 0x03, 0x6e, 0x8f, 0x26, 0x91, 0x4a, 0xce,
	0xb4, 0x00, 0x1f, 0x75, 0x9a, 0xdd, 0xfd, 0xe3, 0x67, 0xb4, 0x46, 0x73, 0x74, 0x05, 0x43, 0x9f,
	0x1b, 0xc6, 0x40, 0x00, 0x6f, 0x8d, 0xca, 0x55, 0xf0, 0x1d, 0xe1, 0x56, 0xf5, 0x99, 0x1c, 0xe0,
	0x1b, 0xf6, 0x87, 0x8f, 0x3a, 0xa8, 0xdb, 0xe6, 0xc5, 0x86, 0xdc, 0xc6, 0x7b, 0xd3, 0x48, 0xc3,
	0x99, 0x1a, 0xfb, 0x8d, 0x0e, 0xea, 0xba, 0xdc, 0x33, 0xdb, 0xfe, 0x98, 0xbc, 0xc4, 0x2d, 0x05,
	0x22, 0xb6, 0xa6, 0x9a, 0xd6, 0xd4, 0xc3, 0x5a, 0xa6, 0x0a, 0x37, 0x7c, 0xcf, 0x88, 0xcd, 0xb1,
	0xa7, 0x18, 0xe7, 0x42, 0xcf, 0xa6, 0x60, 0x49, 0xae, 0x25, 0xf5, 0x76, 0x21, 0x59, 0x31, 0x6f,
	0x17, 0x90, 0x81, 0x80, 0xc3, 0x6f, 0x0d, 0xbc, 0x5f, 0xb6, 0x3e, 0x50, 0x32, 0x21, 0xef, 0xd6,
	0xf3, 0x7b, 0xba, 0x4b, 0x7e, 0x06, 0xb2, 0x29, 0xbd, 0xdf, 0xd7, 0x48, 0xef, 0xed, 0x5a, 0x7a,
	0xc7, 0xb5, 0x2c, 0x55, 0xef, 0x8f, 0x1a, 0x53, 0x57, 0x21, 0xbe, 0xdf, 0x10, 0xe2, 0x93, 0xff,
	0x00, 0xae, 0x65, 0xf9, 0x07, 0xe1, 0x9b, 0x55, 0x0c, 0x10, 0x81, 0x20, 0x7d, 0xec, 0x15, 0x8f,
	0xda, 0x36, 0x5a, 0xf7, 0xd2, 0x4f, 0xad, 0xe4, 0xc4, 0xbd, 0xf8, 0x79, 0xc7, 0xe1, 0x25, 0x80,
	0xbc, 0xc6, 0xde, 0x27, 0x7b, 0x85, 0x36, 0x9b, 0xba, 0x09, 0xac, 0x3c, 0x6a, 0x5e, 0x12, 0xc8,
	0x0b, 0xec, 0x6a, 0x25, 0x13, 0xbf, 0x69, 0x49, 0x8f, 0x76, 0xbd, 0x5e, 0x6e, 0xd5, 0x27, 0xfd,
	0x8b, 0x45, 0x88, 0x2e, 0x17, 0x21, 0xfa, 0xb5, 0x08, 0xd1, 0x97, 0x65, 0xe8, 0x5c, 0x2e, 0x43,
	0xe7, 0xc7, 0x32, 0x74, 0x3e, 0x32, 0xa9, 0x60, 0x32, 0x1b, 0xd2, 0x51, 0x1a, 0x33, 0xc3, 0x3e,
	0x2a, 0xe0, 0xac, 0x1a, 0xe9, 0xf3, 0xab, 0xa1, 0xb6, 0xc3, 0x3c, 0xf4, 0xec, 0x34, 0x3f, 0xfe,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x90, 0x02, 0x6f, 0xae, 0x04, 0x00, 0x00,
}

func (m *GenesisKeygen) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisKeygen) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisKeygen) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainSet) > 0 {
		for iNdEx := len(m.ChainSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenesisKeygen_ChainSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisKeygen_ChainSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisKeygen_ChainSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResultSet) > 0 {
		for iNdEx := len(m.ResultSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResultSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ItemSet) > 0 {
		for iNdEx := len(m.ItemSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ItemSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.LastId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisSign) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisSign) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisSign) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainSet) > 0 {
		for iNdEx := len(m.ChainSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenesisSign_ChainSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisSign_ChainSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisSign_ChainSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResultSet) > 0 {
		for iNdEx := len(m.ResultSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResultSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ItemSet) > 0 {
		for iNdEx := len(m.ItemSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ItemSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.LastId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sign != nil {
		{
			size, err := m.Sign.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Keygen != nil {
		{
			size, err := m.Keygen.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisKeygen) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChainSet) > 0 {
		for _, e := range m.ChainSet {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisKeygen_ChainSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LastId != 0 {
		n += 1 + sovGenesis(uint64(m.LastId))
	}
	if len(m.ItemSet) > 0 {
		for _, e := range m.ItemSet {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ResultSet) > 0 {
		for _, e := range m.ResultSet {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisSign) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChainSet) > 0 {
		for _, e := range m.ChainSet {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisSign_ChainSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.LastId != 0 {
		n += 1 + sovGenesis(uint64(m.LastId))
	}
	if len(m.ItemSet) > 0 {
		for _, e := range m.ItemSet {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ResultSet) > 0 {
		for _, e := range m.ResultSet {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Keygen != nil {
		l = m.Keygen.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Sign != nil {
		l = m.Sign.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisKeygen) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisKeygen: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisKeygen: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainSet = append(m.ChainSet, &GenesisKeygen_ChainSet{})
			if err := m.ChainSet[len(m.ChainSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisKeygen_ChainSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ChainSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastId", wireType)
			}
			m.LastId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemSet = append(m.ItemSet, &Keygen{})
			if err := m.ItemSet[len(m.ItemSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultSet = append(m.ResultSet, &KeygenResult{})
			if err := m.ResultSet[len(m.ResultSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisSign) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisSign: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisSign: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainSet = append(m.ChainSet, &GenesisSign_ChainSet{})
			if err := m.ChainSet[len(m.ChainSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisSign_ChainSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ChainSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastId", wireType)
			}
			m.LastId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ItemSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ItemSet = append(m.ItemSet, &exported.Sign{})
			if err := m.ItemSet[len(m.ItemSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultSet = append(m.ResultSet, &exported.SignResult{})
			if err := m.ResultSet[len(m.ResultSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keygen", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Keygen == nil {
				m.Keygen = &GenesisKeygen{}
			}
			if err := m.Keygen.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sign == nil {
				m.Sign = &GenesisSign{}
			}
			if err := m.Sign.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
