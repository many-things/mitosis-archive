// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.4
// source: pkg/msgconv/osmo/swap.proto

package osmo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom  string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_msgconv_osmo_swap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_msgconv_osmo_swap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file_pkg_msgconv_osmo_swap_proto_rawDescGZIP(), []int{0}
}

func (x *Coin) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Coin) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type SwapAmountInRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolId        uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TokenOutDenom string `protobuf:"bytes,2,opt,name=token_out_denom,json=tokenOutDenom,proto3" json:"token_out_denom,omitempty"`
}

func (x *SwapAmountInRoute) Reset() {
	*x = SwapAmountInRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_msgconv_osmo_swap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapAmountInRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapAmountInRoute) ProtoMessage() {}

func (x *SwapAmountInRoute) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_msgconv_osmo_swap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapAmountInRoute.ProtoReflect.Descriptor instead.
func (*SwapAmountInRoute) Descriptor() ([]byte, []int) {
	return file_pkg_msgconv_osmo_swap_proto_rawDescGZIP(), []int{1}
}

func (x *SwapAmountInRoute) GetPoolId() uint64 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *SwapAmountInRoute) GetTokenOutDenom() string {
	if x != nil {
		return x.TokenOutDenom
	}
	return ""
}

type SwapAmountOutRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolId       uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TokenInDenom string `protobuf:"bytes,2,opt,name=token_in_denom,json=tokenInDenom,proto3" json:"token_in_denom,omitempty"`
}

func (x *SwapAmountOutRoute) Reset() {
	*x = SwapAmountOutRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_msgconv_osmo_swap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapAmountOutRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapAmountOutRoute) ProtoMessage() {}

func (x *SwapAmountOutRoute) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_msgconv_osmo_swap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapAmountOutRoute.ProtoReflect.Descriptor instead.
func (*SwapAmountOutRoute) Descriptor() ([]byte, []int) {
	return file_pkg_msgconv_osmo_swap_proto_rawDescGZIP(), []int{2}
}

func (x *SwapAmountOutRoute) GetPoolId() uint64 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *SwapAmountOutRoute) GetTokenInDenom() string {
	if x != nil {
		return x.TokenInDenom
	}
	return ""
}

type MsgSwapExactAmountIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender            string               `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Routes            []*SwapAmountInRoute `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	TokenIn           *Coin                `protobuf:"bytes,3,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	TokenOutMinAmount string               `protobuf:"bytes,4,opt,name=token_out_min_amount,json=tokenOutMinAmount,proto3" json:"token_out_min_amount,omitempty"`
}

func (x *MsgSwapExactAmountIn) Reset() {
	*x = MsgSwapExactAmountIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_msgconv_osmo_swap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSwapExactAmountIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSwapExactAmountIn) ProtoMessage() {}

func (x *MsgSwapExactAmountIn) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_msgconv_osmo_swap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSwapExactAmountIn.ProtoReflect.Descriptor instead.
func (*MsgSwapExactAmountIn) Descriptor() ([]byte, []int) {
	return file_pkg_msgconv_osmo_swap_proto_rawDescGZIP(), []int{3}
}

func (x *MsgSwapExactAmountIn) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MsgSwapExactAmountIn) GetRoutes() []*SwapAmountInRoute {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *MsgSwapExactAmountIn) GetTokenIn() *Coin {
	if x != nil {
		return x.TokenIn
	}
	return nil
}

func (x *MsgSwapExactAmountIn) GetTokenOutMinAmount() string {
	if x != nil {
		return x.TokenOutMinAmount
	}
	return ""
}

type MsgSwapExactAmountOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender           string                `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Routes           []*SwapAmountOutRoute `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	TokenInMaxAmount string                `protobuf:"bytes,3,opt,name=token_in_max_amount,json=tokenInMaxAmount,proto3" json:"token_in_max_amount,omitempty"`
	TokenOut         *Coin                 `protobuf:"bytes,4,opt,name=token_out,json=tokenOut,proto3" json:"token_out,omitempty"`
}

func (x *MsgSwapExactAmountOut) Reset() {
	*x = MsgSwapExactAmountOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_msgconv_osmo_swap_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSwapExactAmountOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSwapExactAmountOut) ProtoMessage() {}

func (x *MsgSwapExactAmountOut) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_msgconv_osmo_swap_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSwapExactAmountOut.ProtoReflect.Descriptor instead.
func (*MsgSwapExactAmountOut) Descriptor() ([]byte, []int) {
	return file_pkg_msgconv_osmo_swap_proto_rawDescGZIP(), []int{4}
}

func (x *MsgSwapExactAmountOut) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MsgSwapExactAmountOut) GetRoutes() []*SwapAmountOutRoute {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *MsgSwapExactAmountOut) GetTokenInMaxAmount() string {
	if x != nil {
		return x.TokenInMaxAmount
	}
	return ""
}

func (x *MsgSwapExactAmountOut) GetTokenOut() *Coin {
	if x != nil {
		return x.TokenOut
	}
	return nil
}

var File_pkg_msgconv_osmo_swap_proto protoreflect.FileDescriptor

var file_pkg_msgconv_osmo_swap_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x63, 0x6f, 0x6e, 0x76, 0x2f, 0x6f, 0x73,
	0x6d, 0x6f, 0x2f, 0x73, 0x77, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a,
	0x04, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x54, 0x0a, 0x11, 0x53, 0x77, 0x61, 0x70, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49,
	0x64, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x4f, 0x75, 0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x22, 0x53, 0x0a, 0x12, 0x53, 0x77, 0x61,
	0x70, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x22, 0xad,
	0x01, 0x0a, 0x14, 0x4d, 0x73, 0x67, 0x53, 0x77, 0x61, 0x70, 0x45, 0x78, 0x61, 0x63, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x08, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e,
	0x43, 0x6f, 0x69, 0x6e, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x12, 0x2f, 0x0a,
	0x14, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x4f, 0x75, 0x74, 0x4d, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xaf,
	0x01, 0x0a, 0x15, 0x4d, 0x73, 0x67, 0x53, 0x77, 0x61, 0x70, 0x45, 0x78, 0x61, 0x63, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x2b, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x53, 0x77, 0x61, 0x70, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a,
	0x13, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x6e, 0x4d, 0x61, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x09,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x61, 0x6e, 0x79, 0x2d, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x6d, 0x69, 0x74, 0x6f, 0x73,
	0x69, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x73, 0x67, 0x63, 0x6f, 0x6e, 0x76, 0x2f, 0x6f,
	0x73, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_msgconv_osmo_swap_proto_rawDescOnce sync.Once
	file_pkg_msgconv_osmo_swap_proto_rawDescData = file_pkg_msgconv_osmo_swap_proto_rawDesc
)

func file_pkg_msgconv_osmo_swap_proto_rawDescGZIP() []byte {
	file_pkg_msgconv_osmo_swap_proto_rawDescOnce.Do(func() {
		file_pkg_msgconv_osmo_swap_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_msgconv_osmo_swap_proto_rawDescData)
	})
	return file_pkg_msgconv_osmo_swap_proto_rawDescData
}

var file_pkg_msgconv_osmo_swap_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_msgconv_osmo_swap_proto_goTypes = []interface{}{
	(*Coin)(nil),                  // 0: Coin
	(*SwapAmountInRoute)(nil),     // 1: SwapAmountInRoute
	(*SwapAmountOutRoute)(nil),    // 2: SwapAmountOutRoute
	(*MsgSwapExactAmountIn)(nil),  // 3: MsgSwapExactAmountIn
	(*MsgSwapExactAmountOut)(nil), // 4: MsgSwapExactAmountOut
}
var file_pkg_msgconv_osmo_swap_proto_depIdxs = []int32{
	1, // 0: MsgSwapExactAmountIn.routes:type_name -> SwapAmountInRoute
	0, // 1: MsgSwapExactAmountIn.token_in:type_name -> Coin
	2, // 2: MsgSwapExactAmountOut.routes:type_name -> SwapAmountOutRoute
	0, // 3: MsgSwapExactAmountOut.token_out:type_name -> Coin
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_msgconv_osmo_swap_proto_init() }
func file_pkg_msgconv_osmo_swap_proto_init() {
	if File_pkg_msgconv_osmo_swap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_msgconv_osmo_swap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_msgconv_osmo_swap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapAmountInRoute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_msgconv_osmo_swap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapAmountOutRoute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_msgconv_osmo_swap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSwapExactAmountIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_msgconv_osmo_swap_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSwapExactAmountOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_msgconv_osmo_swap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_msgconv_osmo_swap_proto_goTypes,
		DependencyIndexes: file_pkg_msgconv_osmo_swap_proto_depIdxs,
		MessageInfos:      file_pkg_msgconv_osmo_swap_proto_msgTypes,
	}.Build()
	File_pkg_msgconv_osmo_swap_proto = out.File
	file_pkg_msgconv_osmo_swap_proto_rawDesc = nil
	file_pkg_msgconv_osmo_swap_proto_goTypes = nil
	file_pkg_msgconv_osmo_swap_proto_depIdxs = nil
}