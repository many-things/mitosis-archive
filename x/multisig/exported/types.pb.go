// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitosis/multisig/v1beta1/exported/types.proto

package exported

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

// Sign Status
type Sign_Status int32

const (
	// Unspecified Status
	Sign_StatusUnspeicified Sign_Status = 0
	// Sign event Assigned ( Not executed )
	Sign_StatusAssign Sign_Status = 1
	// Sign event executed
	Sign_StatusExecute Sign_Status = 2
	// Sign event Completed
	Sign_StatusComplete Sign_Status = 3
	// Sign event Success
	Sign_StatusSuccess Sign_Status = 4
	// Sign event Failed
	Sign_StatusFailed Sign_Status = 5
)

var Sign_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ASSIGN",
	2: "STATUS_EXECUTE",
	3: "STATUS_COMPLETE",
	4: "STATUS_SUCCESS",
	5: "STATUS_FAILED",
}

var Sign_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_ASSIGN":      1,
	"STATUS_EXECUTE":     2,
	"STATUS_COMPLETE":    3,
	"STATUS_SUCCESS":     4,
	"STATUS_FAILED":      5,
}

func (x Sign_Status) String() string {
	return proto.EnumName(Sign_Status_name, int32(x))
}

func (Sign_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_678ce0504116f33b, []int{0, 0}
}

// Sign Message
type Sign struct {
	// sign target chain id
	Chain string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	// signature id
	SigID uint64 `protobuf:"varint,2,opt,name=sig_id,json=sigId,proto3" json:"sig_id,omitempty"`
	// try-to-sign key id
	KeyID string `protobuf:"bytes,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// operation id
	OpId uint64 `protobuf:"varint,4,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	// participant_ids
	Participants []github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,5,rep,name=participants,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"participants,omitempty"`
	// Message to sign
	MessageToSign Hash        `protobuf:"bytes,6,opt,name=message_to_sign,json=messageToSign,proto3,casttype=Hash" json:"message_to_sign,omitempty"`
	Status        Sign_Status `protobuf:"varint,7,opt,name=status,proto3,enum=manythings.mitosis.v1beta1.multisig.exported.Sign_Status" json:"status,omitempty"`
}

func (m *Sign) Reset()         { *m = Sign{} }
func (m *Sign) String() string { return proto.CompactTextString(m) }
func (*Sign) ProtoMessage()    {}
func (*Sign) Descriptor() ([]byte, []int) {
	return fileDescriptor_678ce0504116f33b, []int{0}
}
func (m *Sign) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sign.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sign.Merge(m, src)
}
func (m *Sign) XXX_Size() int {
	return m.Size()
}
func (m *Sign) XXX_DiscardUnknown() {
	xxx_messageInfo_Sign.DiscardUnknown(m)
}

var xxx_messageInfo_Sign proto.InternalMessageInfo

func (m *Sign) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *Sign) GetSigID() uint64 {
	if m != nil {
		return m.SigID
	}
	return 0
}

func (m *Sign) GetKeyID() string {
	if m != nil {
		return m.KeyID
	}
	return ""
}

func (m *Sign) GetOpId() uint64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *Sign) GetParticipants() []github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *Sign) GetMessageToSign() Hash {
	if m != nil {
		return m.MessageToSign
	}
	return nil
}

func (m *Sign) GetStatus() Sign_Status {
	if m != nil {
		return m.Status
	}
	return Sign_StatusUnspeicified
}

// SignSignature Message
type SignResult struct {
	Chain           string             `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	SigID           uint64             `protobuf:"varint,2,opt,name=sig_id,json=sigId,proto3" json:"sig_id,omitempty"`
	Items           []*SignResult_Item `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	ResultSignature Signature          `protobuf:"bytes,4,opt,name=result_signature,json=resultSignature,proto3,casttype=Signature" json:"result_signature,omitempty"`
}

func (m *SignResult) Reset()         { *m = SignResult{} }
func (m *SignResult) String() string { return proto.CompactTextString(m) }
func (*SignResult) ProtoMessage()    {}
func (*SignResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_678ce0504116f33b, []int{1}
}
func (m *SignResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignResult.Merge(m, src)
}
func (m *SignResult) XXX_Size() int {
	return m.Size()
}
func (m *SignResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SignResult.DiscardUnknown(m)
}

var xxx_messageInfo_SignResult proto.InternalMessageInfo

func (m *SignResult) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *SignResult) GetSigID() uint64 {
	if m != nil {
		return m.SigID
	}
	return 0
}

func (m *SignResult) GetItems() []*SignResult_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SignResult) GetResultSignature() Signature {
	if m != nil {
		return m.ResultSignature
	}
	return nil
}

// SignSignature result message
type SignResult_Item struct {
	Participant github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=participant,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"participant,omitempty"`
	Signature   Signature                                     `protobuf:"bytes,2,opt,name=signature,proto3,casttype=Signature" json:"signature,omitempty"`
}

func (m *SignResult_Item) Reset()         { *m = SignResult_Item{} }
func (m *SignResult_Item) String() string { return proto.CompactTextString(m) }
func (*SignResult_Item) ProtoMessage()    {}
func (*SignResult_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_678ce0504116f33b, []int{1, 0}
}
func (m *SignResult_Item) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignResult_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignResult_Item.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignResult_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignResult_Item.Merge(m, src)
}
func (m *SignResult_Item) XXX_Size() int {
	return m.Size()
}
func (m *SignResult_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_SignResult_Item.DiscardUnknown(m)
}

var xxx_messageInfo_SignResult_Item proto.InternalMessageInfo

func (m *SignResult_Item) GetParticipant() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Participant
	}
	return nil
}

func (m *SignResult_Item) GetSignature() Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("manythings.mitosis.v1beta1.multisig.exported.Sign_Status", Sign_Status_name, Sign_Status_value)
	proto.RegisterType((*Sign)(nil), "manythings.mitosis.v1beta1.multisig.exported.Sign")
	proto.RegisterType((*SignResult)(nil), "manythings.mitosis.v1beta1.multisig.exported.SignResult")
	proto.RegisterType((*SignResult_Item)(nil), "manythings.mitosis.v1beta1.multisig.exported.SignResult.Item")
}

func init() {
	proto.RegisterFile("mitosis/multisig/v1beta1/exported/types.proto", fileDescriptor_678ce0504116f33b)
}

var fileDescriptor_678ce0504116f33b = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x80, 0xe3, 0xc4, 0xce, 0xff, 0x67, 0x9b, 0xb4, 0x66, 0xa9, 0x50, 0x94, 0x43, 0x62, 0x15,
	0x21, 0x22, 0x41, 0x6c, 0x5a, 0x2e, 0x70, 0xe0, 0x90, 0xa4, 0x2e, 0x58, 0x94, 0x52, 0xbc, 0x09,
	0x42, 0x5c, 0x22, 0xd7, 0x5e, 0x9c, 0x55, 0x63, 0xaf, 0x95, 0x5d, 0xa3, 0xe6, 0x0d, 0x50, 0x4e,
	0xbc, 0x40, 0xde, 0x87, 0x0b, 0x52, 0x8f, 0x9c, 0x22, 0x94, 0x3e, 0x05, 0x3d, 0x21, 0x7b, 0x1d,
	0x62, 0x24, 0x2e, 0xed, 0x29, 0xf1, 0xf8, 0x9b, 0x4f, 0x33, 0xe3, 0xd1, 0x80, 0x4e, 0x40, 0x38,
	0x65, 0x84, 0x19, 0x41, 0x3c, 0xe1, 0x84, 0x11, 0xdf, 0xf8, 0xbc, 0x7f, 0x86, 0xb9, 0xb3, 0x6f,
	0xe0, 0x8b, 0x88, 0x4e, 0x39, 0xf6, 0x0c, 0x3e, 0x8b, 0x30, 0xd3, 0xa3, 0x29, 0xe5, 0x14, 0x3e,
	0x0e, 0x9c, 0x70, 0xc6, 0xc7, 0x24, 0xf4, 0x99, 0x9e, 0x65, 0xea, 0x59, 0x82, 0xbe, 0x36, 0xe8,
	0xeb, 0xcc, 0xc6, 0xae, 0x4f, 0x7d, 0x9a, 0x26, 0x1a, 0xc9, 0x3f, 0xe1, 0xd8, 0xfb, 0x2e, 0x03,
	0x19, 0x11, 0x3f, 0x84, 0xbb, 0x40, 0x71, 0xc7, 0x0e, 0x09, 0xeb, 0x92, 0x26, 0xb5, 0x2b, 0xb6,
	0x78, 0x80, 0x1a, 0x28, 0x33, 0xe2, 0x8f, 0x88, 0x57, 0x2f, 0x6a, 0x52, 0x5b, 0xee, 0x55, 0x56,
	0xcb, 0x96, 0x82, 0x88, 0x6f, 0x1d, 0xda, 0x0a, 0x23, 0xbe, 0xe5, 0x25, 0xc4, 0x39, 0x9e, 0x25,
	0x44, 0x29, 0x49, 0x14, 0xc4, 0x6b, 0x3c, 0x4b, 0x88, 0x73, 0x3c, 0xb3, 0x3c, 0x78, 0x17, 0x28,
	0x34, 0x4a, 0x00, 0x39, 0x51, 0xd8, 0x32, 0x8d, 0x2c, 0x0f, 0x0e, 0x41, 0x35, 0x72, 0xa6, 0x9c,
	0xb8, 0x24, 0x72, 0x42, 0xce, 0xea, 0x8a, 0x56, 0x6a, 0x57, 0x7a, 0xfb, 0xd7, 0xcb, 0x56, 0xc7,
	0x27, 0x7c, 0x1c, 0x9f, 0xe9, 0x2e, 0x0d, 0x0c, 0x97, 0xb2, 0x80, 0xb2, 0xec, 0xa7, 0xc3, 0xbc,
	0xf3, 0xac, 0xff, 0xf7, 0xce, 0xa4, 0xeb, 0x79, 0x53, 0xcc, 0x98, 0xfd, 0x97, 0x06, 0x3e, 0x01,
	0x3b, 0x01, 0x66, 0xcc, 0xf1, 0xf1, 0x88, 0xd3, 0x11, 0x23, 0x7e, 0x58, 0x2f, 0x6b, 0x52, 0xbb,
	0xda, 0xfb, 0xff, 0x7a, 0xd9, 0x92, 0x5f, 0x39, 0x6c, 0x6c, 0xd7, 0x32, 0x60, 0x40, 0xd3, 0xbe,
	0xdf, 0x81, 0x32, 0xe3, 0x0e, 0x8f, 0x59, 0xfd, 0x3f, 0x4d, 0x6a, 0x6f, 0x1f, 0x3c, 0xd7, 0x6f,
	0x32, 0x55, 0x3d, 0x71, 0xe8, 0x28, 0x15, 0xd8, 0x99, 0x68, 0xef, 0x97, 0x04, 0xca, 0x22, 0x04,
	0x75, 0x00, 0xd1, 0xa0, 0x3b, 0x18, 0xa2, 0xd1, 0xf0, 0x04, 0x9d, 0x9a, 0x7d, 0xeb, 0xc8, 0x32,
	0x0f, 0xd5, 0x42, 0xe3, 0xde, 0x7c, 0xa1, 0x41, 0xc1, 0x0c, 0x43, 0x16, 0x61, 0xe2, 0x92, 0x4f,
	0x04, 0x7b, 0xf0, 0x3e, 0xa8, 0x65, 0x7c, 0x17, 0x21, 0xeb, 0xe5, 0x89, 0x2a, 0x35, 0xd4, 0xf9,
	0x42, 0xab, 0x0a, 0xb4, 0xcb, 0x92, 0x8e, 0xe0, 0x03, 0xb0, 0x9d, 0x41, 0xe6, 0x07, 0xb3, 0x3f,
	0x1c, 0x98, 0x6a, 0xb1, 0x71, 0x67, 0xbe, 0xd0, 0x6a, 0x82, 0x32, 0x2f, 0xb0, 0x1b, 0x73, 0x0c,
	0x1f, 0x82, 0x9d, 0x0c, 0xeb, 0xbf, 0x7d, 0x73, 0x7a, 0x6c, 0x0e, 0x4c, 0xb5, 0xd4, 0x80, 0xf3,
	0x85, 0xb6, 0x2d, 0xb8, 0x3e, 0x0d, 0xa2, 0x09, 0xe6, 0x38, 0xe7, 0x43, 0xc3, 0x7e, 0xdf, 0x44,
	0x48, 0x95, 0xf3, 0x3e, 0x14, 0xbb, 0x2e, 0x66, 0x2c, 0x57, 0xdb, 0x51, 0xd7, 0x3a, 0x36, 0x0f,
	0x55, 0x25, 0x5f, 0xdb, 0x91, 0x43, 0x26, 0xd8, 0xdb, 0x5b, 0x16, 0x01, 0x48, 0x66, 0x62, 0x63,
	0x16, 0x4f, 0xf8, 0xad, 0xb7, 0x0a, 0x01, 0x85, 0x70, 0x1c, 0xb0, 0x7a, 0x49, 0x2b, 0xb5, 0xb7,
	0x0e, 0x5e, 0xdc, 0xfc, 0xa3, 0x88, 0x02, 0x74, 0x8b, 0xe3, 0xc0, 0x16, 0x2e, 0xf8, 0x0c, 0xa8,
	0xd3, 0x34, 0x9a, 0x2e, 0x86, 0xc3, 0xe3, 0x29, 0x4e, 0x77, 0xb2, 0xda, 0xab, 0x5d, 0x2f, 0x5b,
	0x15, 0xb4, 0x0e, 0xda, 0x3b, 0x02, 0xfb, 0x13, 0x68, 0x7c, 0x91, 0x80, 0x9c, 0x98, 0x20, 0x02,
	0x5b, 0xb9, 0x7d, 0x4b, 0xbb, 0xaa, 0xde, 0x66, 0x6b, 0xf3, 0x16, 0xf8, 0x08, 0x54, 0x36, 0x05,
	0x15, 0xff, 0x55, 0xd0, 0xe6, 0x7d, 0xef, 0xf8, 0xdb, 0xaa, 0x29, 0x5d, 0xae, 0x9a, 0xd2, 0xcf,
	0x55, 0x53, 0xfa, 0x7a, 0xd5, 0x2c, 0x5c, 0x5e, 0x35, 0x0b, 0x3f, 0xae, 0x9a, 0x85, 0x8f, 0x07,
	0xb9, 0x12, 0x92, 0x71, 0x75, 0xc4, 0xbc, 0x8c, 0xf5, 0x51, 0xb9, 0xd8, 0x9c, 0x95, 0xf5, 0xa4,
	0xce, 0xca, 0xe9, 0x15, 0x78, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x63, 0x7c, 0x7f, 0x14, 0x7a,
	0x04, 0x00, 0x00,
}

func (m *Sign) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sign) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sign) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	if len(m.MessageToSign) > 0 {
		i -= len(m.MessageToSign)
		copy(dAtA[i:], m.MessageToSign)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.MessageToSign)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Participants) > 0 {
		for iNdEx := len(m.Participants) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Participants[iNdEx])
			copy(dAtA[i:], m.Participants[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Participants[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.OpId != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.OpId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SigID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SigID))
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

func (m *SignResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResultSignature) > 0 {
		i -= len(m.ResultSignature)
		copy(dAtA[i:], m.ResultSignature)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ResultSignature)))
		i--
		dAtA[i] = 0x22
	}
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
	if m.SigID != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SigID))
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

func (m *SignResult_Item) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignResult_Item) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignResult_Item) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Signature)))
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
func (m *Sign) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.SigID != 0 {
		n += 1 + sovTypes(uint64(m.SigID))
	}
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.OpId != 0 {
		n += 1 + sovTypes(uint64(m.OpId))
	}
	if len(m.Participants) > 0 {
		for _, s := range m.Participants {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.MessageToSign)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovTypes(uint64(m.Status))
	}
	return n
}

func (m *SignResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.SigID != 0 {
		n += 1 + sovTypes(uint64(m.SigID))
	}
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.ResultSignature)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *SignResult_Item) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Signature)
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
func (m *Sign) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Sign: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sign: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field SigID", wireType)
			}
			m.SigID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
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
			m.KeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpId", wireType)
			}
			m.OpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.Participants = append(m.Participants, github_com_cosmos_cosmos_sdk_types.ValAddress(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageToSign", wireType)
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
			m.MessageToSign = append(m.MessageToSign[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageToSign == nil {
				m.MessageToSign = []byte{}
			}
			iNdEx = postIndex
		case 7:
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
				m.Status |= Sign_Status(b&0x7F) << shift
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
func (m *SignResult) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SignResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignResult: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field SigID", wireType)
			}
			m.SigID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigID |= uint64(b&0x7F) << shift
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
			m.Items = append(m.Items, &SignResult_Item{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultSignature", wireType)
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
			m.ResultSignature = append(m.ResultSignature[:0], dAtA[iNdEx:postIndex]...)
			if m.ResultSignature == nil {
				m.ResultSignature = []byte{}
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
func (m *SignResult_Item) Unmarshal(dAtA []byte) error {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
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
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
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
