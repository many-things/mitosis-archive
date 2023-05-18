// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sidecar/common.proto

package types

import (
	fmt "fmt"
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

type KeyPresenceResponse_Response int32

const (
	KeyPresenceResponse_RESPONSE_UNSPECIFIED KeyPresenceResponse_Response = 0
	KeyPresenceResponse_RESPONSE_PRESENT     KeyPresenceResponse_Response = 1
	KeyPresenceResponse_RESPONSE_ABSENT      KeyPresenceResponse_Response = 2
	KeyPresenceResponse_RESPONSE_FAIL        KeyPresenceResponse_Response = 3
)

var KeyPresenceResponse_Response_name = map[int32]string{
	0: "RESPONSE_UNSPECIFIED",
	1: "RESPONSE_PRESENT",
	2: "RESPONSE_ABSENT",
	3: "RESPONSE_FAIL",
}

var KeyPresenceResponse_Response_value = map[string]int32{
	"RESPONSE_UNSPECIFIED": 0,
	"RESPONSE_PRESENT":     1,
	"RESPONSE_ABSENT":      2,
	"RESPONSE_FAIL":        3,
}

func (x KeyPresenceResponse_Response) String() string {
	return proto.EnumName(KeyPresenceResponse_Response_name, int32(x))
}

func (KeyPresenceResponse_Response) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c9bf3db3c598a19f, []int{1, 0}
}

type KeyPresenceRequest struct {
	KeyUid string `protobuf:"bytes,1,opt,name=key_uid,json=keyUid,proto3" json:"key_uid,omitempty"`
	PubKey []byte `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *KeyPresenceRequest) Reset()         { *m = KeyPresenceRequest{} }
func (m *KeyPresenceRequest) String() string { return proto.CompactTextString(m) }
func (*KeyPresenceRequest) ProtoMessage()    {}
func (*KeyPresenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9bf3db3c598a19f, []int{0}
}
func (m *KeyPresenceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyPresenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyPresenceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyPresenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPresenceRequest.Merge(m, src)
}
func (m *KeyPresenceRequest) XXX_Size() int {
	return m.Size()
}
func (m *KeyPresenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPresenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPresenceRequest proto.InternalMessageInfo

func (m *KeyPresenceRequest) GetKeyUid() string {
	if m != nil {
		return m.KeyUid
	}
	return ""
}

func (m *KeyPresenceRequest) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type KeyPresenceResponse struct {
	Response KeyPresenceResponse_Response `protobuf:"varint,1,opt,name=response,proto3,enum=tofnd.KeyPresenceResponse_Response" json:"response,omitempty"`
}

func (m *KeyPresenceResponse) Reset()         { *m = KeyPresenceResponse{} }
func (m *KeyPresenceResponse) String() string { return proto.CompactTextString(m) }
func (*KeyPresenceResponse) ProtoMessage()    {}
func (*KeyPresenceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9bf3db3c598a19f, []int{1}
}
func (m *KeyPresenceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyPresenceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyPresenceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyPresenceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPresenceResponse.Merge(m, src)
}
func (m *KeyPresenceResponse) XXX_Size() int {
	return m.Size()
}
func (m *KeyPresenceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPresenceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPresenceResponse proto.InternalMessageInfo

func (m *KeyPresenceResponse) GetResponse() KeyPresenceResponse_Response {
	if m != nil {
		return m.Response
	}
	return KeyPresenceResponse_RESPONSE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("tofnd.KeyPresenceResponse_Response", KeyPresenceResponse_Response_name, KeyPresenceResponse_Response_value)
	proto.RegisterType((*KeyPresenceRequest)(nil), "tofnd.KeyPresenceRequest")
	proto.RegisterType((*KeyPresenceResponse)(nil), "tofnd.KeyPresenceResponse")
}

func init() { proto.RegisterFile("sidecar/common.proto", fileDescriptor_c9bf3db3c598a19f) }

var fileDescriptor_c9bf3db3c598a19f = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xc7, 0x9b, 0x89, 0x73, 0x06, 0xff, 0xd4, 0x6c, 0xe0, 0x4e, 0x65, 0xcc, 0xcb, 0x0e, 0xda,
	0x82, 0x3e, 0x80, 0x6c, 0x9a, 0xc2, 0x98, 0xcc, 0x92, 0xba, 0x8b, 0x97, 0xb1, 0xb6, 0x3f, 0xb7,
	0x30, 0x9a, 0xd4, 0x26, 0x3d, 0xe4, 0x2d, 0x7c, 0x1e, 0x9f, 0xc0, 0xe3, 0x8e, 0x1e, 0x65, 0x7d,
	0x11, 0xb1, 0xb8, 0x82, 0xe0, 0xed, 0xcb, 0xe7, 0x9b, 0x7c, 0x7e, 0xf0, 0xc5, 0x1d, 0xc5, 0x13,
	0x88, 0x17, 0xb9, 0x17, 0xcb, 0x34, 0x95, 0xc2, 0xcd, 0x72, 0xa9, 0x25, 0xd9, 0xd7, 0xf2, 0x45,
	0x24, 0x7d, 0x1f, 0x93, 0x09, 0x98, 0x20, 0x07, 0x05, 0x22, 0x06, 0x06, 0xaf, 0x05, 0x28, 0x4d,
	0xce, 0xf1, 0xc1, 0x1a, 0xcc, 0xbc, 0xe0, 0x49, 0x17, 0xf5, 0xd0, 0xe0, 0x90, 0x35, 0xd7, 0x60,
	0x66, 0x3c, 0xf9, 0x29, 0xb2, 0x22, 0x9a, 0xaf, 0xc1, 0x74, 0x1b, 0x3d, 0x34, 0x38, 0x62, 0xcd,
	0xac, 0x88, 0x26, 0x60, 0xfa, 0xef, 0x08, 0xb7, 0xff, 0x88, 0x54, 0x26, 0x85, 0x02, 0x72, 0x8b,
	0x5b, 0xf9, 0x6f, 0xae, 0x54, 0x27, 0xd7, 0x17, 0x6e, 0x75, 0xd9, 0xfd, 0xe7, 0xb5, 0xbb, 0x0b,
	0xac, 0xfe, 0xd4, 0x8f, 0x70, 0xab, 0x96, 0x75, 0x71, 0x87, 0xd1, 0x30, 0x78, 0x9c, 0x86, 0x74,
	0x3e, 0x9b, 0x86, 0x01, 0xbd, 0x1b, 0xfb, 0x63, 0x7a, 0x6f, 0x5b, 0xa4, 0x83, 0xed, 0xba, 0x09,
	0x18, 0x0d, 0xe9, 0xf4, 0xc9, 0x46, 0xa4, 0x8d, 0x4f, 0x6b, 0x3a, 0x1c, 0x55, 0xb0, 0x41, 0xce,
	0xf0, 0x71, 0x0d, 0xfd, 0xe1, 0xf8, 0xc1, 0xde, 0x1b, 0xf9, 0x1f, 0x5b, 0x07, 0x6d, 0xb6, 0x0e,
	0xfa, 0xda, 0x3a, 0xe8, 0xad, 0x74, 0xac, 0x4d, 0xe9, 0x58, 0x9f, 0xa5, 0x63, 0x3d, 0x5f, 0x2e,
	0xb9, 0x5e, 0x15, 0x91, 0x1b, 0xcb, 0xd4, 0x4b, 0x17, 0xc2, 0x5c, 0xe9, 0x15, 0x17, 0x4b, 0xe5,
	0xa5, 0x5c, 0x4b, 0xc5, 0x95, 0xb7, 0x9b, 0x56, 0x9b, 0x0c, 0x54, 0xd4, 0xac, 0xa6, 0xbd, 0xf9,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x30, 0xcc, 0xb4, 0x72, 0x01, 0x00, 0x00,
}

func (m *KeyPresenceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyPresenceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyPresenceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.KeyUid) > 0 {
		i -= len(m.KeyUid)
		copy(dAtA[i:], m.KeyUid)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.KeyUid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KeyPresenceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyPresenceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyPresenceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Response != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Response))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyPresenceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyUid)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *KeyPresenceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Response != 0 {
		n += 1 + sovCommon(uint64(m.Response))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyPresenceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: KeyPresenceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyPresenceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyUid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyUid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
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
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *KeyPresenceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: KeyPresenceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyPresenceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			m.Response = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Response |= KeyPresenceResponse_Response(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
