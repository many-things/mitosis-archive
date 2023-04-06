// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitosis/multisig/v1beta1/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_many_things_mitosis_x_multisig_exported "github.com/many-things/mitosis/x/multisig/exported"
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

// Keygen Status
type Keygen_Status int32

const (
	// Unspecified Status
	Keygen_StatusUnspecified Keygen_Status = 0
	// Keygen Assigned Status ( Not executed )
	Keygen_StatusAssign Keygen_Status = 1
	// Keygen Started
	Keygen_StatusExecute Keygen_Status = 2
	// Keygen Completed
	Keygen_StatusComplete Keygen_Status = 3
	// Keygen Failed
	Keygen_StatusFailed Keygen_Status = 4
)

var Keygen_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ASSIGN",
	2: "STATUS_EXECUTE",
	3: "STATUS_COMPLETED",
	4: "STATUS_FAILED",
}

var Keygen_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_ASSIGN":      1,
	"STATUS_EXECUTE":     2,
	"STATUS_COMPLETED":   3,
	"STATUS_FAILED":      4,
}

func (x Keygen_Status) String() string {
	return proto.EnumName(Keygen_Status_name, int32(x))
}

func (Keygen_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1f5869ea852ab5d9, []int{0, 0}
}

// Keygen Message
type Keygen struct {
	// key target chain id
	Chain string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	// key id
	KeyID uint64 `protobuf:"varint,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// keygen participants
	Participants []github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,3,rep,name=participants,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"participants,omitempty"`
	// keygen event status
	Status Keygen_Status `protobuf:"varint,4,opt,name=status,proto3,enum=manythings.mitosis.v1beta1.multisig.Keygen_Status" json:"status,omitempty"`
}

func (m *Keygen) Reset()         { *m = Keygen{} }
func (m *Keygen) String() string { return proto.CompactTextString(m) }
func (*Keygen) ProtoMessage()    {}
func (*Keygen) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5869ea852ab5d9, []int{0}
}
func (m *Keygen) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Keygen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Keygen.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Keygen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keygen.Merge(m, src)
}
func (m *Keygen) XXX_Size() int {
	return m.Size()
}
func (m *Keygen) XXX_DiscardUnknown() {
	xxx_messageInfo_Keygen.DiscardUnknown(m)
}

var xxx_messageInfo_Keygen proto.InternalMessageInfo

func (m *Keygen) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *Keygen) GetKeyID() uint64 {
	if m != nil {
		return m.KeyID
	}
	return 0
}

func (m *Keygen) GetParticipants() []github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *Keygen) GetStatus() Keygen_Status {
	if m != nil {
		return m.Status
	}
	return Keygen_StatusUnspecified
}

// Pubkey Message
type PubKey struct {
	// key target chain id
	Chain string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	// key id
	KeyID uint64 `protobuf:"varint,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// participant id
	Items []*PubKey_Item `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (m *PubKey) Reset()         { *m = PubKey{} }
func (m *PubKey) String() string { return proto.CompactTextString(m) }
func (*PubKey) ProtoMessage()    {}
func (*PubKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5869ea852ab5d9, []int{1}
}
func (m *PubKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKey.Merge(m, src)
}
func (m *PubKey) XXX_Size() int {
	return m.Size()
}
func (m *PubKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKey.DiscardUnknown(m)
}

var xxx_messageInfo_PubKey proto.InternalMessageInfo

func (m *PubKey) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *PubKey) GetKeyID() uint64 {
	if m != nil {
		return m.KeyID
	}
	return 0
}

func (m *PubKey) GetItems() []*PubKey_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

// Item returns
type PubKey_Item struct {
	Participant github_com_cosmos_cosmos_sdk_types.ValAddress                `protobuf:"bytes,1,opt,name=participant,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"participant,omitempty"`
	PubKey      github_com_many_things_mitosis_x_multisig_exported.PublicKey `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3,casttype=github.com/many-things/mitosis/x/multisig/exported.PublicKey" json:"pub_key,omitempty"`
}

func (m *PubKey_Item) Reset()         { *m = PubKey_Item{} }
func (m *PubKey_Item) String() string { return proto.CompactTextString(m) }
func (*PubKey_Item) ProtoMessage()    {}
func (*PubKey_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5869ea852ab5d9, []int{1, 0}
}
func (m *PubKey_Item) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKey_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKey_Item.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKey_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKey_Item.Merge(m, src)
}
func (m *PubKey_Item) XXX_Size() int {
	return m.Size()
}
func (m *PubKey_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKey_Item.DiscardUnknown(m)
}

var xxx_messageInfo_PubKey_Item proto.InternalMessageInfo

func (m *PubKey_Item) GetParticipant() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Participant
	}
	return nil
}

func (m *PubKey_Item) GetPubKey() github_com_many_things_mitosis_x_multisig_exported.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func init() {
	proto.RegisterEnum("manythings.mitosis.v1beta1.multisig.Keygen_Status", Keygen_Status_name, Keygen_Status_value)
	proto.RegisterType((*Keygen)(nil), "manythings.mitosis.v1beta1.multisig.Keygen")
	proto.RegisterType((*PubKey)(nil), "manythings.mitosis.v1beta1.multisig.PubKey")
	proto.RegisterType((*PubKey_Item)(nil), "manythings.mitosis.v1beta1.multisig.PubKey.Item")
}

func init() {
	proto.RegisterFile("mitosis/multisig/v1beta1/types.proto", fileDescriptor_1f5869ea852ab5d9)
}

var fileDescriptor_1f5869ea852ab5d9 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xbb, 0x8e, 0xda, 0x4e,
	0x14, 0xc6, 0x31, 0x17, 0xff, 0xb5, 0xb3, 0x2c, 0xf2, 0x8e, 0xf6, 0x2f, 0x21, 0x17, 0xc6, 0x62,
	0x13, 0x89, 0x06, 0x3b, 0x90, 0x36, 0x45, 0xb8, 0x98, 0xc8, 0x61, 0xb3, 0x41, 0x36, 0x44, 0x49,
	0x1a, 0x64, 0xec, 0x89, 0x19, 0x81, 0x2f, 0x62, 0xc6, 0x11, 0x7e, 0x05, 0xaa, 0xbc, 0x00, 0xaf,
	0x11, 0xa5, 0x4d, 0x97, 0x72, 0x95, 0x2a, 0x15, 0x8a, 0xe0, 0x2d, 0xb6, 0x8a, 0xf0, 0x78, 0x15,
	0xa7, 0x43, 0x5b, 0xd9, 0x63, 0xfd, 0xce, 0x77, 0xbe, 0xcf, 0x67, 0x0e, 0x78, 0xe2, 0x61, 0x1a,
	0x10, 0x4c, 0x54, 0x2f, 0x5a, 0x52, 0x4c, 0xb0, 0xab, 0x7e, 0x6e, 0xcd, 0x10, 0xb5, 0x5a, 0x2a,
	0x8d, 0x43, 0x44, 0x94, 0x70, 0x15, 0xd0, 0x00, 0x5e, 0x7b, 0x96, 0x1f, 0xd3, 0x39, 0xf6, 0x5d,
	0xa2, 0xa4, 0x05, 0x4a, 0xca, 0x29, 0x0f, 0x85, 0xe2, 0x95, 0x1b, 0xb8, 0x41, 0xc2, 0xab, 0xc7,
	0x37, 0x56, 0x5a, 0xff, 0x56, 0x00, 0xfc, 0x10, 0xc5, 0x2e, 0xf2, 0xe1, 0x15, 0x28, 0xd9, 0x73,
	0x0b, 0xfb, 0x55, 0x4e, 0xe6, 0x1a, 0x67, 0x06, 0x3b, 0x40, 0x19, 0xf0, 0x0b, 0x14, 0x4f, 0xb1,
	0x53, 0xcd, 0xcb, 0x5c, 0xa3, 0xd8, 0x3d, 0xdb, 0xef, 0x6a, 0xa5, 0x21, 0x8a, 0xf5, 0xbe, 0x51,
	0x5a, 0xa0, 0x58, 0x77, 0xe0, 0x04, 0x94, 0x43, 0x6b, 0x45, 0xb1, 0x8d, 0x43, 0xcb, 0xa7, 0xa4,
	0x5a, 0x90, 0x0b, 0x8d, 0x72, 0xb7, 0x75, 0xbf, 0xab, 0x35, 0x5d, 0x4c, 0xe7, 0xd1, 0x4c, 0xb1,
	0x03, 0x4f, 0xb5, 0x03, 0xe2, 0x05, 0x24, 0x7d, 0x34, 0x89, 0xb3, 0x48, 0x13, 0xbc, 0xb3, 0x96,
	0x1d, 0xc7, 0x59, 0x21, 0x42, 0x8c, 0x7f, 0x64, 0xe0, 0x6b, 0xc0, 0x13, 0x6a, 0xd1, 0x88, 0x54,
	0x8b, 0x32, 0xd7, 0xa8, 0xb4, 0xdb, 0xca, 0x09, 0x29, 0x15, 0x96, 0x45, 0x31, 0x93, 0x4a, 0x23,
	0x55, 0xa8, 0xff, 0xe4, 0x00, 0xcf, 0x3e, 0xc1, 0x26, 0x80, 0xe6, 0xb8, 0x33, 0x9e, 0x98, 0xd3,
	0xc9, 0xad, 0x39, 0xd2, 0x7a, 0xfa, 0x40, 0xd7, 0xfa, 0x42, 0x4e, 0xfc, 0x7f, 0xb3, 0x95, 0x2f,
	0x19, 0x33, 0xf1, 0x49, 0x88, 0x6c, 0xfc, 0x09, 0x23, 0x07, 0x5e, 0x83, 0x8b, 0x14, 0xef, 0x98,
	0xa6, 0xfe, 0xea, 0x56, 0xe0, 0x44, 0x61, 0xb3, 0x95, 0xcb, 0x8c, 0xec, 0x10, 0x82, 0x5d, 0x1f,
	0x3e, 0x05, 0x95, 0x14, 0xd2, 0xde, 0x6b, 0xbd, 0xc9, 0x58, 0x13, 0xf2, 0xe2, 0xe5, 0x66, 0x2b,
	0x5f, 0x30, 0x4a, 0x5b, 0x23, 0x3b, 0xa2, 0x08, 0x36, 0x80, 0x90, 0x62, 0xbd, 0xb7, 0x6f, 0x46,
	0x37, 0xda, 0x58, 0xeb, 0x0b, 0x05, 0x11, 0x6e, 0xb6, 0x72, 0x85, 0x81, 0xbd, 0xc0, 0x0b, 0x97,
	0x88, 0xa2, 0x4c, 0xd7, 0x41, 0x47, 0xbf, 0xd1, 0xfa, 0x42, 0x31, 0xdb, 0x75, 0x60, 0xe1, 0x25,
	0x72, 0xea, 0xdf, 0xf3, 0x80, 0x1f, 0x45, 0xb3, 0x21, 0x8a, 0x1f, 0x3d, 0xba, 0x01, 0x28, 0x61,
	0x8a, 0x3c, 0x36, 0xb3, 0xf3, 0xf6, 0xb3, 0x93, 0x7e, 0x31, 0xeb, 0xa9, 0xe8, 0x14, 0x79, 0x06,
	0x2b, 0x17, 0xbf, 0x72, 0xa0, 0x78, 0x3c, 0x43, 0x13, 0x9c, 0x67, 0x86, 0x98, 0xd8, 0x79, 0xd4,
	0x55, 0xc8, 0xaa, 0xc0, 0x0f, 0xe0, 0xbf, 0x30, 0x9a, 0x4d, 0x17, 0x28, 0x4e, 0x82, 0x94, 0xbb,
	0x2f, 0xef, 0x77, 0xb5, 0x17, 0x19, 0xc1, 0xa3, 0xeb, 0x26, 0xb3, 0xad, 0x3e, 0x2c, 0xcc, 0xfa,
	0xef, 0xca, 0xa0, 0x75, 0x18, 0xac, 0x28, 0x72, 0x8e, 0xce, 0x97, 0xd8, 0x1e, 0xa2, 0xd8, 0xe0,
	0xc3, 0x24, 0x44, 0x57, 0xff, 0xb1, 0x97, 0xb8, 0xbb, 0xbd, 0xc4, 0xfd, 0xde, 0x4b, 0xdc, 0x97,
	0x83, 0x94, 0xbb, 0x3b, 0x48, 0xb9, 0x5f, 0x07, 0x29, 0xf7, 0x51, 0x3d, 0x5d, 0x3f, 0x71, 0x3f,
	0xe3, 0x93, 0x85, 0x7a, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xb5, 0xda, 0x71, 0xb3, 0x03,
	0x00, 0x00,
}

func (m *Keygen) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Keygen) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Keygen) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Participants[iNdEx])
			copy(dAtA[i:], m.Participants[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Participants[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.KeyID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.KeyID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PubKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.KeyID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.KeyID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PubKey_Item) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKey_Item) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKey_Item) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Keygen) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.KeyID != 0 {
		n += 1 + sovTypes(uint64(m.KeyID))
	}
	if len(m.Participants) > 0 {
		for _, b := range m.Participants {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovTypes(uint64(m.Status))
	}
	return n
}

func (m *PubKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.KeyID != 0 {
		n += 1 + sovTypes(uint64(m.KeyID))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *PubKey_Item) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Keygen) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Keygen: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Keygen: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			m.KeyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participants = append(m.Participants, make([]byte, postIndex-iNdEx))
			copy(m.Participants[len(m.Participants)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Keygen_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *PubKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PubKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			m.KeyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &PubKey_Item{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *PubKey_Item) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Item: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Item: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participant", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participant = append(m.Participant[:0], dAtA[iNdEx:postIndex]...)
			if m.Participant == nil {
				m.Participant = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
