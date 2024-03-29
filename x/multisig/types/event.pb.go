// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitosis/multisig/v1beta1/event.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/many-things/mitosis/pkg/types"
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

// Event for Emit Sign
type EventSigningStart struct {
	Chain         string                                          `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	SigId         uint64                                          `protobuf:"varint,2,opt,name=sig_id,json=sigId,proto3" json:"sig_id,omitempty"`
	KeyId         string                                          `protobuf:"bytes,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	OpId          uint64                                          `protobuf:"varint,4,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	Participants  []github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,5,rep,name=participants,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"participants,omitempty"`
	MessageToSign []byte                                          `protobuf:"bytes,6,opt,name=message_to_sign,json=messageToSign,proto3" json:"message_to_sign,omitempty"`
}

func (m *EventSigningStart) Reset()         { *m = EventSigningStart{} }
func (m *EventSigningStart) String() string { return proto.CompactTextString(m) }
func (*EventSigningStart) ProtoMessage()    {}
func (*EventSigningStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6868c47c227711f, []int{0}
}
func (m *EventSigningStart) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSigningStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSigningStart.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSigningStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSigningStart.Merge(m, src)
}
func (m *EventSigningStart) XXX_Size() int {
	return m.Size()
}
func (m *EventSigningStart) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSigningStart.DiscardUnknown(m)
}

var xxx_messageInfo_EventSigningStart proto.InternalMessageInfo

func (m *EventSigningStart) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *EventSigningStart) GetSigId() uint64 {
	if m != nil {
		return m.SigId
	}
	return 0
}

func (m *EventSigningStart) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *EventSigningStart) GetOpId() uint64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *EventSigningStart) GetParticipants() []github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *EventSigningStart) GetMessageToSign() []byte {
	if m != nil {
		return m.MessageToSign
	}
	return nil
}

func init() {
	proto.RegisterType((*EventSigningStart)(nil), "manythings.mitosis.v1beta1.multisig.EventSigningStart")
}

func init() {
	proto.RegisterFile("mitosis/multisig/v1beta1/event.proto", fileDescriptor_d6868c47c227711f)
}

var fileDescriptor_d6868c47c227711f = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0xc6, 0x9d, 0xab, 0x11, 0x0c, 0x5e, 0x2e, 0x37, 0xb5, 0x10, 0x5c, 0xa4, 0xc1, 0x96, 0x92,
	0x8d, 0x09, 0xd2, 0x27, 0xa8, 0xd0, 0x45, 0xb6, 0xb1, 0xed, 0xa2, 0x1b, 0x19, 0x93, 0x61, 0x3c,
	0xe8, 0xcc, 0x84, 0x9c, 0x51, 0x9a, 0xb7, 0xe8, 0x63, 0x75, 0xe9, 0xb2, 0xab, 0x52, 0x74, 0xd9,
	0x37, 0xe8, 0xaa, 0x4c, 0x12, 0xfb, 0x67, 0x95, 0xcc, 0xc7, 0xef, 0x3b, 0xdf, 0x99, 0x6f, 0xec,
	0x0b, 0x01, 0x5a, 0x21, 0x60, 0x24, 0x36, 0x6b, 0x0d, 0x08, 0x3c, 0xda, 0x4e, 0x16, 0x4c, 0xd3,
	0x49, 0xc4, 0xb6, 0x4c, 0xea, 0x30, 0x2f, 0x94, 0x56, 0xce, 0xb9, 0xa0, 0xb2, 0xd4, 0x4b, 0x90,
	0x1c, 0xc3, 0xc6, 0x10, 0x36, 0x5c, 0x78, 0x34, 0x0e, 0x07, 0x5c, 0x71, 0x55, 0xf1, 0x91, 0xf9,
	0xab, 0xad, 0xc3, 0xd1, 0x31, 0x20, 0x55, 0x42, 0x28, 0xf9, 0x35, 0x5e, 0x97, 0x39, 0xc3, 0x9a,
	0x19, 0xbd, 0x13, 0xfb, 0xff, 0x8d, 0x89, 0x9b, 0x01, 0x97, 0x20, 0xf9, 0x4c, 0xd3, 0x42, 0x3b,
	0x03, 0xdb, 0x4a, 0x97, 0x14, 0xa4, 0x4b, 0x7c, 0x12, 0xf4, 0x92, 0xfa, 0xe0, 0x9c, 0xda, 0x5d,
	0x04, 0x3e, 0x87, 0xcc, 0xfd, 0xe3, 0x93, 0xa0, 0x93, 0x58, 0x08, 0x3c, 0xce, 0x8c, 0xbc, 0x62,
	0xa5, 0x91, 0xdb, 0x35, 0xbd, 0x62, 0x65, 0x9c, 0x39, 0x27, 0xb6, 0xa5, 0x72, 0xa3, 0x76, 0x2a,
	0xb8, 0xa3, 0xf2, 0x38, 0x73, 0xee, 0xec, 0x7e, 0x4e, 0x0b, 0x0d, 0x29, 0xe4, 0x54, 0x6a, 0x74,
	0x2d, 0xbf, 0x1d, 0xf4, 0xa6, 0x93, 0x8f, 0xd7, 0xb3, 0x31, 0x07, 0xbd, 0xdc, 0x2c, 0xc2, 0x54,
	0x89, 0x28, 0x55, 0x28, 0x14, 0x36, 0x9f, 0x31, 0x66, 0xab, 0x66, 0xe5, 0x7b, 0xba, 0xbe, 0xce,
	0xb2, 0x82, 0x21, 0x26, 0xbf, 0xc6, 0x38, 0x97, 0xf6, 0x3f, 0xc1, 0x10, 0x29, 0x67, 0x73, 0xad,
	0xe6, 0x08, 0x5c, 0xba, 0x5d, 0x9f, 0x04, 0xfd, 0xe4, 0x6f, 0x23, 0xdf, 0x2a, 0x73, 0xbf, 0x69,
	0xfc, 0xbc, 0xf7, 0xc8, 0x6e, 0xef, 0x91, 0xb7, 0xbd, 0x47, 0x9e, 0x0e, 0x5e, 0x6b, 0x77, 0xf0,
	0x5a, 0x2f, 0x07, 0xaf, 0xf5, 0x10, 0xfd, 0x88, 0x37, 0x8d, 0x8f, 0xeb, 0xca, 0xa3, 0x63, 0x85,
	0x8f, 0xdf, 0xaf, 0x54, 0xed, 0xb2, 0xe8, 0x56, 0xfd, 0x5d, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x52, 0xdc, 0xf2, 0xd8, 0xc6, 0x01, 0x00, 0x00,
}

func (m *EventSigningStart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSigningStart) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSigningStart) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageToSign) > 0 {
		i -= len(m.MessageToSign)
		copy(dAtA[i:], m.MessageToSign)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.MessageToSign)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Participants[iNdEx])
			copy(dAtA[i:], m.Participants[iNdEx])
			i = encodeVarintEvent(dAtA, i, uint64(len(m.Participants[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.OpId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.OpId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.KeyId) > 0 {
		i -= len(m.KeyId)
		copy(dAtA[i:], m.KeyId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.KeyId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SigId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.SigId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventSigningStart) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.SigId != 0 {
		n += 1 + sovEvent(uint64(m.SigId))
	}
	l = len(m.KeyId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.OpId != 0 {
		n += 1 + sovEvent(uint64(m.OpId))
	}
	if len(m.Participants) > 0 {
		for _, s := range m.Participants {
			l = len(s)
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	l = len(m.MessageToSign)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventSigningStart) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventSigningStart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSigningStart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigId", wireType)
			}
			m.SigId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpId", wireType)
			}
			m.OpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OpId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participants", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participants = append(m.Participants, github_com_cosmos_cosmos_sdk_types.ValAddress(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageToSign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageToSign = append(m.MessageToSign[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageToSign == nil {
				m.MessageToSign = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
