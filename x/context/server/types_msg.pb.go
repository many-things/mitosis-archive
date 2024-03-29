// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitosis/context/v1beta1/server/types_msg.proto

package server

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type MsgRegisterVault struct {
	Sender    github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	Chain     string                                        `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	VaultAddr string                                        `protobuf:"bytes,3,opt,name=vault_addr,json=vaultAddr,proto3" json:"vault_addr,omitempty"`
}

func (m *MsgRegisterVault) Reset()         { *m = MsgRegisterVault{} }
func (m *MsgRegisterVault) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterVault) ProtoMessage()    {}
func (*MsgRegisterVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc66b8b1d363cb, []int{0}
}
func (m *MsgRegisterVault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterVault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterVault.Merge(m, src)
}
func (m *MsgRegisterVault) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterVault) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterVault.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterVault proto.InternalMessageInfo

func (m *MsgRegisterVault) GetSender() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *MsgRegisterVault) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *MsgRegisterVault) GetVaultAddr() string {
	if m != nil {
		return m.VaultAddr
	}
	return ""
}

// Response of [MsgRegisterVault]
type MsgRegisterVaultResponse struct {
}

func (m *MsgRegisterVaultResponse) Reset()         { *m = MsgRegisterVaultResponse{} }
func (m *MsgRegisterVaultResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterVaultResponse) ProtoMessage()    {}
func (*MsgRegisterVaultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc66b8b1d363cb, []int{1}
}
func (m *MsgRegisterVaultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterVaultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterVaultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterVaultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterVaultResponse.Merge(m, src)
}
func (m *MsgRegisterVaultResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterVaultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterVaultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterVaultResponse proto.InternalMessageInfo

type MsgClearVault struct {
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	Chain  string                                        `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (m *MsgClearVault) Reset()         { *m = MsgClearVault{} }
func (m *MsgClearVault) String() string { return proto.CompactTextString(m) }
func (*MsgClearVault) ProtoMessage()    {}
func (*MsgClearVault) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc66b8b1d363cb, []int{2}
}
func (m *MsgClearVault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClearVault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClearVault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClearVault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClearVault.Merge(m, src)
}
func (m *MsgClearVault) XXX_Size() int {
	return m.Size()
}
func (m *MsgClearVault) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClearVault.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClearVault proto.InternalMessageInfo

func (m *MsgClearVault) GetSender() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *MsgClearVault) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

// Response of [MsgClearVault]
type MsgClearVaultResponse struct {
}

func (m *MsgClearVaultResponse) Reset()         { *m = MsgClearVaultResponse{} }
func (m *MsgClearVaultResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClearVaultResponse) ProtoMessage()    {}
func (*MsgClearVaultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3dc66b8b1d363cb, []int{3}
}
func (m *MsgClearVaultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClearVaultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClearVaultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClearVaultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClearVaultResponse.Merge(m, src)
}
func (m *MsgClearVaultResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClearVaultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClearVaultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClearVaultResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterVault)(nil), "manythings.mitosis.v1beta1.context.server.MsgRegisterVault")
	proto.RegisterType((*MsgRegisterVaultResponse)(nil), "manythings.mitosis.v1beta1.context.server.MsgRegisterVaultResponse")
	proto.RegisterType((*MsgClearVault)(nil), "manythings.mitosis.v1beta1.context.server.MsgClearVault")
	proto.RegisterType((*MsgClearVaultResponse)(nil), "manythings.mitosis.v1beta1.context.server.MsgClearVaultResponse")
}

func init() {
	proto.RegisterFile("mitosis/context/v1beta1/server/types_msg.proto", fileDescriptor_d3dc66b8b1d363cb)
}

var fileDescriptor_d3dc66b8b1d363cb = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0xc1, 0x4a, 0xfb, 0x30,
	0x1c, 0xc7, 0x97, 0xff, 0x1f, 0x07, 0x0b, 0x0a, 0x52, 0x26, 0x96, 0x81, 0x75, 0xf4, 0x34, 0x0f,
	0x4b, 0x18, 0x3e, 0xc1, 0xe6, 0x69, 0x87, 0x5d, 0x7a, 0xf0, 0xe0, 0x65, 0x74, 0xcd, 0x8f, 0x2c,
	0xb8, 0x26, 0x25, 0xbf, 0x6c, 0x6c, 0x6f, 0x21, 0x3e, 0x95, 0xc7, 0x1d, 0x3d, 0x89, 0x6c, 0x6f,
	0xe1, 0x49, 0xda, 0xb4, 0xa2, 0x3e, 0x80, 0xa7, 0x36, 0x09, 0x9f, 0xef, 0xf7, 0x0b, 0x1f, 0xca,
	0x72, 0xe5, 0x0c, 0x2a, 0xe4, 0x99, 0xd1, 0x0e, 0xb6, 0x8e, 0x6f, 0x46, 0x0b, 0x70, 0xe9, 0x88,
	0x23, 0xd8, 0x0d, 0x58, 0xee, 0x76, 0x05, 0xe0, 0x3c, 0x47, 0xc9, 0x0a, 0x6b, 0x9c, 0x09, 0x6e,
	0xf2, 0x54, 0xef, 0xdc, 0x52, 0x69, 0x89, 0x0d, 0xca, 0x6a, 0x84, 0xd5, 0x11, 0xcc, 0xa3, 0xbd,
	0xae, 0x34, 0xd2, 0x54, 0x14, 0x2f, 0xff, 0x7c, 0x40, 0xfc, 0x4c, 0xe8, 0xf9, 0x0c, 0x65, 0x02,
	0x52, 0xa1, 0x03, 0x7b, 0x9f, 0xae, 0x57, 0x2e, 0x98, 0xd2, 0x36, 0x82, 0x16, 0x60, 0x43, 0xd2,
	0x27, 0x83, 0xd3, 0xc9, 0xe8, 0xe3, 0xed, 0x7a, 0x28, 0x95, 0x5b, 0xae, 0x17, 0x2c, 0x33, 0x39,
	0xcf, 0x0c, 0xe6, 0x06, 0xeb, 0xcf, 0x10, 0xc5, 0xa3, 0xdf, 0xc5, 0xc6, 0x59, 0x36, 0x16, 0xc2,
	0x02, 0x62, 0x52, 0x07, 0x04, 0x5d, 0x7a, 0x92, 0x2d, 0x53, 0xa5, 0xc3, 0x7f, 0x7d, 0x32, 0xe8,
	0x24, 0xfe, 0x10, 0x5c, 0x51, 0xba, 0x29, 0x9b, 0xe6, 0xa9, 0x10, 0x36, 0xfc, 0x5f, 0x3d, 0x75,
	0xaa, 0x9b, 0x92, 0x8f, 0x7b, 0x34, 0xfc, 0xbd, 0x29, 0x01, 0x2c, 0x8c, 0x46, 0x88, 0x0b, 0x7a,
	0x36, 0x43, 0x79, 0xb7, 0x82, 0xf4, 0x8f, 0xc6, 0xc6, 0x97, 0xf4, 0xe2, 0x47, 0x63, 0x33, 0x65,
	0x32, 0x7d, 0x39, 0x44, 0x64, 0x7f, 0x88, 0xc8, 0xfb, 0x21, 0x22, 0x4f, 0xc7, 0xa8, 0xb5, 0x3f,
	0x46, 0xad, 0xd7, 0x63, 0xd4, 0x7a, 0xe0, 0xdf, 0xfa, 0x4b, 0x43, 0x43, 0xaf, 0x88, 0x37, 0x76,
	0xb7, 0x5f, 0x7e, 0xbd, 0x9c, 0x45, 0xbb, 0xb2, 0x71, 0xfb, 0x19, 0x00, 0x00, 0xff, 0xff, 0x1d,
	0x21, 0x7e, 0x7b, 0x00, 0x02, 0x00, 0x00,
}

func (m *MsgRegisterVault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterVault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterVault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VaultAddr) > 0 {
		i -= len(m.VaultAddr)
		copy(dAtA[i:], m.VaultAddr)
		i = encodeVarintTypesMsg(dAtA, i, uint64(len(m.VaultAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTypesMsg(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTypesMsg(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterVaultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterVaultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterVaultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgClearVault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClearVault) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClearVault) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTypesMsg(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTypesMsg(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClearVaultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClearVaultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClearVaultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTypesMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypesMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterVault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTypesMsg(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypesMsg(uint64(l))
	}
	l = len(m.VaultAddr)
	if l > 0 {
		n += 1 + l + sovTypesMsg(uint64(l))
	}
	return n
}

func (m *MsgRegisterVaultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgClearVault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTypesMsg(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypesMsg(uint64(l))
	}
	return n
}

func (m *MsgClearVaultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTypesMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypesMsg(x uint64) (n int) {
	return sovTypesMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterVault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesMsg
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
			return fmt.Errorf("proto: MsgRegisterVault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterVault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesMsg
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
				return ErrInvalidLengthTypesMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesMsg
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
				return ErrInvalidLengthTypesMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesMsg
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
				return ErrInvalidLengthTypesMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VaultAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypesMsg
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
func (m *MsgRegisterVaultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesMsg
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
			return fmt.Errorf("proto: MsgRegisterVaultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterVaultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypesMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypesMsg
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
func (m *MsgClearVault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesMsg
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
			return fmt.Errorf("proto: MsgClearVault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClearVault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesMsg
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
				return ErrInvalidLengthTypesMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesMsg
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
				return ErrInvalidLengthTypesMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypesMsg
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
func (m *MsgClearVaultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesMsg
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
			return fmt.Errorf("proto: MsgClearVaultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClearVaultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypesMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypesMsg
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
func skipTypesMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypesMsg
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
					return 0, ErrIntOverflowTypesMsg
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
					return 0, ErrIntOverflowTypesMsg
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
				return 0, ErrInvalidLengthTypesMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypesMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypesMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypesMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypesMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypesMsg = fmt.Errorf("proto: unexpected end of group")
)
