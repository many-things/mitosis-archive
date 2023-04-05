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
	Participant github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,3,opt,name=participant,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"participant,omitempty"`
	// pub key
	PubKey github_com_many_things_mitosis_x_multisig_exported.PublicKey `protobuf:"bytes,4,opt,name=pub_key,json=pubKey,proto3,casttype=github.com/many-things/mitosis/x/multisig/exported.PublicKey" json:"pub_key,omitempty"`
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

func (m *PubKey) GetParticipant() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Participant
	}
	return nil
}

func (m *PubKey) GetPubKey() github_com_many_things_mitosis_x_multisig_exported.PublicKey {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func init() {
	proto.RegisterEnum("manythings.mitosis.v1beta1.multisig.Keygen_Status", Keygen_Status_name, Keygen_Status_value)
	proto.RegisterType((*Keygen)(nil), "manythings.mitosis.v1beta1.multisig.Keygen")
	proto.RegisterType((*PubKey)(nil), "manythings.mitosis.v1beta1.multisig.PubKey")
}

func init() {
	proto.RegisterFile("mitosis/multisig/v1beta1/types.proto", fileDescriptor_1f5869ea852ab5d9)
}

var fileDescriptor_1f5869ea852ab5d9 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcb, 0x6e, 0xda, 0x40,
	0x18, 0x85, 0x99, 0x10, 0x5c, 0x65, 0x4a, 0x90, 0x33, 0x4a, 0x25, 0xc4, 0xc2, 0x58, 0xa4, 0x95,
	0xd8, 0x60, 0x8b, 0x74, 0xdb, 0x45, 0xb9, 0x98, 0xca, 0x25, 0x4d, 0x11, 0x86, 0xaa, 0xed, 0x06,
	0xf9, 0xf2, 0xd7, 0x8c, 0xc0, 0x17, 0x31, 0xe3, 0x0a, 0xbf, 0x02, 0xab, 0xbe, 0x00, 0xef, 0xd1,
	0x47, 0xe8, 0x32, 0xea, 0xaa, 0x2b, 0x54, 0xc1, 0x0b, 0x74, 0x9d, 0x55, 0x15, 0xdb, 0x51, 0xdd,
	0x5d, 0x94, 0x95, 0x2f, 0xfa, 0xce, 0xf9, 0xe7, 0xcc, 0xf9, 0xf1, 0x73, 0x8f, 0xf2, 0x80, 0x51,
	0xa6, 0x7a, 0xd1, 0x92, 0x53, 0x46, 0x5d, 0xf5, 0x6b, 0xdb, 0x02, 0x6e, 0xb6, 0x55, 0x1e, 0x87,
	0xc0, 0x94, 0x70, 0x15, 0xf0, 0x80, 0x5c, 0x78, 0xa6, 0x1f, 0xf3, 0x39, 0xf5, 0x5d, 0xa6, 0x64,
	0x02, 0x25, 0xe3, 0x94, 0x7b, 0x61, 0xed, 0xdc, 0x0d, 0xdc, 0x20, 0xe1, 0xd5, 0xbb, 0xb7, 0x54,
	0xda, 0xf8, 0x5e, 0xc4, 0xc2, 0x10, 0x62, 0x17, 0x7c, 0x72, 0x8e, 0x4b, 0xf6, 0xdc, 0xa4, 0x7e,
	0x15, 0xc9, 0xa8, 0x79, 0x32, 0x4e, 0x3f, 0x88, 0x8c, 0x85, 0x05, 0xc4, 0x33, 0xea, 0x54, 0x8f,
	0x64, 0xd4, 0x3c, 0xee, 0x9e, 0xec, 0x77, 0xf5, 0xd2, 0x10, 0x62, 0xbd, 0x3f, 0x2e, 0x2d, 0x20,
	0xd6, 0x1d, 0x32, 0xc5, 0xe5, 0xd0, 0x5c, 0x71, 0x6a, 0xd3, 0xd0, 0xf4, 0x39, 0xab, 0x16, 0xe5,
	0x62, 0xb3, 0xdc, 0x6d, 0xdf, 0xee, 0xea, 0x2d, 0x97, 0xf2, 0x79, 0x64, 0x29, 0x76, 0xe0, 0xa9,
	0x76, 0xc0, 0xbc, 0x80, 0x65, 0x8f, 0x16, 0x73, 0x16, 0x59, 0x82, 0x0f, 0xe6, 0xb2, 0xe3, 0x38,
	0x2b, 0x60, 0x6c, 0xfc, 0x9f, 0x0d, 0x79, 0x8b, 0x05, 0xc6, 0x4d, 0x1e, 0xb1, 0xea, 0xb1, 0x8c,
	0x9a, 0x95, 0xcb, 0x4b, 0xe5, 0x01, 0x29, 0x95, 0x34, 0x8b, 0x62, 0x24, 0xca, 0x71, 0xe6, 0xd0,
	0xf8, 0x89, 0xb0, 0x90, 0xfe, 0x22, 0x2d, 0x4c, 0x8c, 0x49, 0x67, 0x32, 0x35, 0x66, 0xd3, 0x6b,
	0x63, 0xa4, 0xf5, 0xf4, 0x81, 0xae, 0xf5, 0xc5, 0x42, 0xed, 0xd9, 0x66, 0x2b, 0x9f, 0xa5, 0xcc,
	0xd4, 0x67, 0x21, 0xd8, 0xf4, 0x0b, 0x05, 0x87, 0x5c, 0xe0, 0xd3, 0x0c, 0xef, 0x18, 0x86, 0xfe,
	0xe6, 0x5a, 0x44, 0x35, 0x71, 0xb3, 0x95, 0xcb, 0x29, 0xd9, 0x61, 0x8c, 0xba, 0x3e, 0x79, 0x81,
	0x2b, 0x19, 0xa4, 0x7d, 0xd4, 0x7a, 0xd3, 0x89, 0x26, 0x1e, 0xd5, 0xce, 0x36, 0x5b, 0xf9, 0x34,
	0xa5, 0xb4, 0x35, 0xd8, 0x11, 0x07, 0xd2, 0xc4, 0x62, 0x86, 0xf5, 0xde, 0xbf, 0x1b, 0x5d, 0x69,
	0x13, 0xad, 0x2f, 0x16, 0x6b, 0x64, 0xb3, 0x95, 0x2b, 0x29, 0xd8, 0x0b, 0xbc, 0x70, 0x09, 0x1c,
	0x72, 0x53, 0x07, 0x1d, 0xfd, 0x4a, 0xeb, 0x8b, 0xc7, 0xf9, 0xa9, 0x03, 0x93, 0x2e, 0xc1, 0x69,
	0xfc, 0x41, 0x58, 0x18, 0x45, 0xd6, 0x10, 0xe2, 0x47, 0x57, 0x67, 0xe0, 0xa7, 0xb9, 0x3b, 0xaf,
	0x16, 0x65, 0xf4, 0xb8, 0xe6, 0xf2, 0x2e, 0xe4, 0x13, 0x7e, 0x12, 0x46, 0xd6, 0x6c, 0x01, 0x71,
	0xd2, 0x5c, 0xb9, 0xfb, 0xfa, 0x76, 0x57, 0x7f, 0x95, 0x33, 0xbc, 0xeb, 0xb1, 0x95, 0x16, 0xa9,
	0xde, 0xef, 0xf7, 0xfa, 0xdf, 0x86, 0xc3, 0x3a, 0x0c, 0x56, 0x1c, 0x1c, 0x65, 0x14, 0x59, 0x4b,
	0x6a, 0x0f, 0x21, 0x1e, 0x0b, 0x61, 0x92, 0xb3, 0xab, 0xff, 0xd8, 0x4b, 0xe8, 0x66, 0x2f, 0xa1,
	0xdf, 0x7b, 0x09, 0x7d, 0x3b, 0x48, 0x85, 0x9b, 0x83, 0x54, 0xf8, 0x75, 0x90, 0x0a, 0x9f, 0xd5,
	0x87, 0xfb, 0x27, 0xa7, 0xb7, 0x84, 0x64, 0xff, 0x5f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x5e,
	0x9c, 0x57, 0x10, 0x62, 0x03, 0x00, 0x00,
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
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0x1a
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
		case 4:
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
