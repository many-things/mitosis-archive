// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitosis/event/v1beta1/types_event.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/many-things/mitosis/pkg/types"
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

// TxReqEvent Message
type TxReqEvent struct {
	DestChain string        `protobuf:"bytes,1,opt,name=dest_chain,json=destChain,proto3" json:"dest_chain,omitempty"`
	DestAddr  []byte        `protobuf:"bytes,2,opt,name=dest_addr,json=destAddr,proto3" json:"dest_addr,omitempty"`
	OpId      uint64        `protobuf:"varint,3,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	OpArgs    [][]byte      `protobuf:"bytes,4,rep,name=op_args,json=opArgs,proto3" json:"op_args,omitempty"`
	Funds     []*types.Coin `protobuf:"bytes,5,rep,name=funds,proto3" json:"funds,omitempty"`
}

func (m *TxReqEvent) Reset()         { *m = TxReqEvent{} }
func (m *TxReqEvent) String() string { return proto.CompactTextString(m) }
func (*TxReqEvent) ProtoMessage()    {}
func (*TxReqEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3829ccfefb0f93fa, []int{0}
}
func (m *TxReqEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxReqEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxReqEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxReqEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxReqEvent.Merge(m, src)
}
func (m *TxReqEvent) XXX_Size() int {
	return m.Size()
}
func (m *TxReqEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TxReqEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TxReqEvent proto.InternalMessageInfo

func (m *TxReqEvent) GetDestChain() string {
	if m != nil {
		return m.DestChain
	}
	return ""
}

func (m *TxReqEvent) GetDestAddr() []byte {
	if m != nil {
		return m.DestAddr
	}
	return nil
}

func (m *TxReqEvent) GetOpId() uint64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *TxReqEvent) GetOpArgs() [][]byte {
	if m != nil {
		return m.OpArgs
	}
	return nil
}

func (m *TxReqEvent) GetFunds() []*types.Coin {
	if m != nil {
		return m.Funds
	}
	return nil
}

// TxResEvent Message
type TxResEvent struct {
	// Request Event id correspond to given response
	ReqEvtId uint64 `protobuf:"varint,1,opt,name=req_evt_id,json=reqEvtId,proto3" json:"req_evt_id,omitempty"`
	Ok       bool   `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	Result   []byte `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *TxResEvent) Reset()         { *m = TxResEvent{} }
func (m *TxResEvent) String() string { return proto.CompactTextString(m) }
func (*TxResEvent) ProtoMessage()    {}
func (*TxResEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3829ccfefb0f93fa, []int{1}
}
func (m *TxResEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxResEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxResEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxResEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResEvent.Merge(m, src)
}
func (m *TxResEvent) XXX_Size() int {
	return m.Size()
}
func (m *TxResEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TxResEvent proto.InternalMessageInfo

func (m *TxResEvent) GetReqEvtId() uint64 {
	if m != nil {
		return m.ReqEvtId
	}
	return 0
}

func (m *TxResEvent) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *TxResEvent) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

// Event Message
type Event struct {
	// Block height
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// Transaction hash
	TxHash []byte `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	// Event Index
	EvtIdx uint32 `protobuf:"varint,3,opt,name=evt_idx,json=evtIdx,proto3" json:"evt_idx,omitempty"`
	// Actual event payload
	//
	// Types that are valid to be assigned to Event:
	//
	//	*Event_Req
	//	*Event_Res
	Event isEvent_Event `protobuf_oneof:"event"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_3829ccfefb0f93fa, []int{2}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_Event interface {
	isEvent_Event()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Event_Req struct {
	Req *TxReqEvent `protobuf:"bytes,4,opt,name=req,proto3,oneof" json:"req,omitempty"`
}
type Event_Res struct {
	Res *TxResEvent `protobuf:"bytes,5,opt,name=res,proto3,oneof" json:"res,omitempty"`
}

func (*Event_Req) isEvent_Event() {}
func (*Event_Res) isEvent_Event() {}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Event) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Event) GetEvtIdx() uint32 {
	if m != nil {
		return m.EvtIdx
	}
	return 0
}

func (m *Event) GetReq() *TxReqEvent {
	if x, ok := m.GetEvent().(*Event_Req); ok {
		return x.Req
	}
	return nil
}

func (m *Event) GetRes() *TxResEvent {
	if x, ok := m.GetEvent().(*Event_Res); ok {
		return x.Res
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_Req)(nil),
		(*Event_Res)(nil),
	}
}

func init() {
	proto.RegisterType((*TxReqEvent)(nil), "manythings.mitosis.v1beta1.event.TxReqEvent")
	proto.RegisterType((*TxResEvent)(nil), "manythings.mitosis.v1beta1.event.TxResEvent")
	proto.RegisterType((*Event)(nil), "manythings.mitosis.v1beta1.event.Event")
}

func init() {
	proto.RegisterFile("mitosis/event/v1beta1/types_event.proto", fileDescriptor_3829ccfefb0f93fa)
}

var fileDescriptor_3829ccfefb0f93fa = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0xb6, 0xc9, 0xba, 0xb7, 0xc2, 0xc1, 0xa0, 0x11, 0x0d, 0x88, 0xa2, 0x5e, 0x96,
	0xc3, 0x70, 0xb4, 0x71, 0x46, 0x62, 0x9b, 0x40, 0xdb, 0xd5, 0xe2, 0xc4, 0x25, 0x4a, 0x67, 0x13,
	0x5b, 0xa3, 0x71, 0x6a, 0xbb, 0x55, 0xf6, 0x2d, 0xf8, 0x32, 0x7c, 0x07, 0x8e, 0x3b, 0x72, 0x42,
	0xa8, 0xfd, 0x22, 0xc8, 0x76, 0x2a, 0xe0, 0x02, 0xda, 0xcd, 0xef, 0xf9, 0xbd, 0xff, 0xef, 0xff,
	0xec, 0x07, 0xc7, 0x0b, 0x69, 0x95, 0x91, 0xa6, 0xe0, 0x6b, 0xde, 0xd8, 0x62, 0x7d, 0x3a, 0xe7,
	0xb6, 0x3a, 0x2d, 0xec, 0x5d, 0xcb, 0x4d, 0xe9, 0x73, 0xa4, 0xd5, 0xca, 0x2a, 0x9c, 0x2d, 0xaa,
	0xe6, 0xce, 0x0a, 0xd9, 0xd4, 0x86, 0xf4, 0x3d, 0xa4, 0xaf, 0x26, 0xbe, 0xee, 0xe8, 0x69, 0xad,
	0x6a, 0xe5, 0x8b, 0x0b, 0x77, 0x0a, 0x7d, 0x47, 0xb3, 0x1d, 0xe0, 0x46, 0x2d, 0x16, 0xaa, 0xf9,
	0x9b, 0x10, 0x6a, 0x66, 0x5f, 0x11, 0xc0, 0x87, 0x8e, 0xf2, 0xe5, 0x3b, 0x27, 0x84, 0x5f, 0x02,
	0x30, 0x6e, 0x6c, 0x79, 0x23, 0x2a, 0xd9, 0x24, 0x28, 0x43, 0xf9, 0x3e, 0xdd, 0x77, 0x99, 0x4b,
	0x97, 0xc0, 0xcf, 0xc1, 0x07, 0x65, 0xc5, 0x98, 0x4e, 0x86, 0x19, 0xca, 0xa7, 0x74, 0xe2, 0x12,
	0xe7, 0x8c, 0x69, 0xfc, 0x04, 0x22, 0xd5, 0x96, 0x92, 0x25, 0xa3, 0x0c, 0xe5, 0x63, 0x3a, 0x56,
	0xed, 0x35, 0xc3, 0xcf, 0x60, 0x4f, 0xb5, 0x65, 0xa5, 0x6b, 0x93, 0x8c, 0xb3, 0x51, 0x3e, 0xa5,
	0xb1, 0x6a, 0xcf, 0x75, 0x6d, 0xf0, 0x1b, 0x88, 0x3e, 0xad, 0x1a, 0x66, 0x92, 0x28, 0x1b, 0xe5,
	0x07, 0x67, 0xc7, 0xe4, 0x1f, 0x43, 0x06, 0xff, 0xe4, 0x52, 0xc9, 0x86, 0x86, 0xae, 0x19, 0x0d,
	0xb6, 0x4d, 0xb0, 0xfd, 0x02, 0x40, 0xf3, 0x65, 0xc9, 0xd7, 0xd6, 0xf1, 0x91, 0xe7, 0x4f, 0xb4,
	0x1b, 0xca, 0x5e, 0x33, 0xfc, 0x18, 0x86, 0xea, 0xd6, 0xdb, 0x9d, 0xd0, 0xa1, 0xba, 0xc5, 0x87,
	0x10, 0x6b, 0x6e, 0x56, 0x9f, 0xad, 0x77, 0x3a, 0xa5, 0x7d, 0x34, 0xfb, 0x81, 0x20, 0x0a, 0x7a,
	0x87, 0x10, 0x0b, 0x2e, 0x6b, 0x61, 0x7b, 0xad, 0x3e, 0x72, 0xd3, 0xd8, 0xae, 0x14, 0x95, 0x11,
	0xfd, 0xf4, 0xb1, 0xed, 0xae, 0x2a, 0x23, 0xdc, 0x45, 0x80, 0x77, 0x5e, 0xf3, 0x11, 0x8d, 0xb9,
	0x43, 0x77, 0xf8, 0x2d, 0x8c, 0x34, 0x5f, 0x26, 0xe3, 0x0c, 0xe5, 0x07, 0x67, 0x27, 0xe4, 0x7f,
	0x3f, 0x49, 0x7e, 0xff, 0xc5, 0xd5, 0x80, 0xba, 0xd6, 0xa0, 0xe0, 0x9e, 0xe9, 0x01, 0x0a, 0xe6,
	0x0f, 0x05, 0x73, 0xb1, 0x07, 0x91, 0xbf, 0xba, 0x78, 0xff, 0x6d, 0x93, 0xa2, 0xfb, 0x4d, 0x8a,
	0x7e, 0x6e, 0x52, 0xf4, 0x65, 0x9b, 0x0e, 0xee, 0xb7, 0xe9, 0xe0, 0xfb, 0x36, 0x1d, 0x7c, 0x3c,
	0xa9, 0xa5, 0x15, 0xab, 0xb9, 0x7b, 0xe9, 0xc2, 0x11, 0x5e, 0x05, 0x44, 0xb1, 0xdb, 0xa0, 0xae,
	0x5f, 0x52, 0xbf, 0x3a, 0xf3, 0xd8, 0xef, 0xce, 0xeb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4,
	0xf3, 0xe1, 0x7b, 0xc2, 0x02, 0x00, 0x00,
}

func (m *TxReqEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxReqEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxReqEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Funds) > 0 {
		for iNdEx := len(m.Funds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Funds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypesEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.OpArgs) > 0 {
		for iNdEx := len(m.OpArgs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.OpArgs[iNdEx])
			copy(dAtA[i:], m.OpArgs[iNdEx])
			i = encodeVarintTypesEvent(dAtA, i, uint64(len(m.OpArgs[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.OpId != 0 {
		i = encodeVarintTypesEvent(dAtA, i, uint64(m.OpId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.DestAddr) > 0 {
		i -= len(m.DestAddr)
		copy(dAtA[i:], m.DestAddr)
		i = encodeVarintTypesEvent(dAtA, i, uint64(len(m.DestAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DestChain) > 0 {
		i -= len(m.DestChain)
		copy(dAtA[i:], m.DestChain)
		i = encodeVarintTypesEvent(dAtA, i, uint64(len(m.DestChain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxResEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxResEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxResEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintTypesEvent(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Ok {
		i--
		if m.Ok {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.ReqEvtId != 0 {
		i = encodeVarintTypesEvent(dAtA, i, uint64(m.ReqEvtId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Event != nil {
		{
			size := m.Event.Size()
			i -= size
			if _, err := m.Event.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.EvtIdx != 0 {
		i = encodeVarintTypesEvent(dAtA, i, uint64(m.EvtIdx))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintTypesEvent(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintTypesEvent(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Event_Req) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event_Req) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Req != nil {
		{
			size, err := m.Req.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypesEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *Event_Res) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event_Res) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Res != nil {
		{
			size, err := m.Res.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypesEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func encodeVarintTypesEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypesEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxReqEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DestChain)
	if l > 0 {
		n += 1 + l + sovTypesEvent(uint64(l))
	}
	l = len(m.DestAddr)
	if l > 0 {
		n += 1 + l + sovTypesEvent(uint64(l))
	}
	if m.OpId != 0 {
		n += 1 + sovTypesEvent(uint64(m.OpId))
	}
	if len(m.OpArgs) > 0 {
		for _, b := range m.OpArgs {
			l = len(b)
			n += 1 + l + sovTypesEvent(uint64(l))
		}
	}
	if len(m.Funds) > 0 {
		for _, e := range m.Funds {
			l = e.Size()
			n += 1 + l + sovTypesEvent(uint64(l))
		}
	}
	return n
}

func (m *TxResEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReqEvtId != 0 {
		n += 1 + sovTypesEvent(uint64(m.ReqEvtId))
	}
	if m.Ok {
		n += 2
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovTypesEvent(uint64(l))
	}
	return n
}

func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypesEvent(uint64(m.Height))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovTypesEvent(uint64(l))
	}
	if m.EvtIdx != 0 {
		n += 1 + sovTypesEvent(uint64(m.EvtIdx))
	}
	if m.Event != nil {
		n += m.Event.Size()
	}
	return n
}

func (m *Event_Req) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Req != nil {
		l = m.Req.Size()
		n += 1 + l + sovTypesEvent(uint64(l))
	}
	return n
}
func (m *Event_Res) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Res != nil {
		l = m.Res.Size()
		n += 1 + l + sovTypesEvent(uint64(l))
	}
	return n
}

func sovTypesEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypesEvent(x uint64) (n int) {
	return sovTypesEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxReqEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesEvent
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
			return fmt.Errorf("proto: TxReqEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxReqEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
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
				return ErrInvalidLengthTypesEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
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
				return ErrInvalidLengthTypesEvent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestAddr = append(m.DestAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.DestAddr == nil {
				m.DestAddr = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpId", wireType)
			}
			m.OpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OpArgs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
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
				return ErrInvalidLengthTypesEvent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OpArgs = append(m.OpArgs, make([]byte, postIndex-iNdEx))
			copy(m.OpArgs[len(m.OpArgs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
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
				return ErrInvalidLengthTypesEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Funds = append(m.Funds, &types.Coin{})
			if err := m.Funds[len(m.Funds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypesEvent
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
func (m *TxResEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesEvent
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
			return fmt.Errorf("proto: TxResEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxResEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqEvtId", wireType)
			}
			m.ReqEvtId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqEvtId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ok = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
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
				return ErrInvalidLengthTypesEvent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypesEvent
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
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypesEvent
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
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
				return ErrInvalidLengthTypesEvent
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypesEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvtIdx", wireType)
			}
			m.EvtIdx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EvtIdx |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Req", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
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
				return ErrInvalidLengthTypesEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TxReqEvent{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &Event_Req{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Res", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypesEvent
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
				return ErrInvalidLengthTypesEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypesEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TxResEvent{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &Event_Res{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypesEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypesEvent
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
func skipTypesEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypesEvent
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
					return 0, ErrIntOverflowTypesEvent
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
					return 0, ErrIntOverflowTypesEvent
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
				return 0, ErrInvalidLengthTypesEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypesEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypesEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypesEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypesEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypesEvent = fmt.Errorf("proto: unexpected end of group")
)
