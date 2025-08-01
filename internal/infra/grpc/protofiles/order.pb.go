// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: internal/infra/grpc/protofiles/order.proto

package orderpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Product       string                 `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Price         float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *Order) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ListOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersRequest) Reset() {
	*x = ListOrdersRequest{}
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersRequest) ProtoMessage() {}

func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP(), []int{1}
}

type ListOrdersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOrdersResponse) Reset() {
	*x = ListOrdersResponse{}
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersResponse) ProtoMessage() {}

func (x *ListOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP(), []int{2}
}

func (x *ListOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Product       string                 `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Price         float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrderRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *CreateOrderRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderRequest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Product       string                 `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Price         float32                `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	FinalPrice    float32                `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infra_grpc_protofiles_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateOrderResponse) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *CreateOrderResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderResponse) GetFinalPrice() float32 {
	if x != nil {
		return x.FinalPrice
	}
	return 0
}

var File_internal_infra_grpc_protofiles_order_proto protoreflect.FileDescriptor

const file_internal_infra_grpc_protofiles_order_proto_rawDesc = "" +
	"\n" +
	"*internal/infra/grpc/protofiles/order.proto\x12\aorderpb\"f\n" +
	"\x05Order\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x18\n" +
	"\aproduct\x18\x02 \x01(\tR\aproduct\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x02R\x05price\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\tR\tcreatedAt\"\x13\n" +
	"\x11ListOrdersRequest\"<\n" +
	"\x12ListOrdersResponse\x12&\n" +
	"\x06orders\x18\x01 \x03(\v2\x0e.orderpb.OrderR\x06orders\"s\n" +
	"\x12CreateOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x18\n" +
	"\aproduct\x18\x02 \x01(\tR\aproduct\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x02R\x05price\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\tR\tcreatedAt\"v\n" +
	"\x13CreateOrderResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x18\n" +
	"\aproduct\x18\x02 \x01(\tR\aproduct\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x02R\x05price\x12\x1f\n" +
	"\vfinal_price\x18\x04 \x01(\x02R\n" +
	"finalPrice2\x9f\x01\n" +
	"\fOrderService\x12H\n" +
	"\vCreateOrder\x12\x1b.orderpb.CreateOrderRequest\x1a\x1c.orderpb.CreateOrderResponse\x12E\n" +
	"\n" +
	"ListOrders\x12\x1a.orderpb.ListOrdersRequest\x1a\x1b.orderpb.ListOrdersResponseB;Z9desafio_clean_architecture/internal/infra/grpc/pb/orderpbb\x06proto3"

var (
	file_internal_infra_grpc_protofiles_order_proto_rawDescOnce sync.Once
	file_internal_infra_grpc_protofiles_order_proto_rawDescData []byte
)

func file_internal_infra_grpc_protofiles_order_proto_rawDescGZIP() []byte {
	file_internal_infra_grpc_protofiles_order_proto_rawDescOnce.Do(func() {
		file_internal_infra_grpc_protofiles_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_infra_grpc_protofiles_order_proto_rawDesc), len(file_internal_infra_grpc_protofiles_order_proto_rawDesc)))
	})
	return file_internal_infra_grpc_protofiles_order_proto_rawDescData
}

var file_internal_infra_grpc_protofiles_order_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_infra_grpc_protofiles_order_proto_goTypes = []any{
	(*Order)(nil),               // 0: orderpb.Order
	(*ListOrdersRequest)(nil),   // 1: orderpb.ListOrdersRequest
	(*ListOrdersResponse)(nil),  // 2: orderpb.ListOrdersResponse
	(*CreateOrderRequest)(nil),  // 3: orderpb.CreateOrderRequest
	(*CreateOrderResponse)(nil), // 4: orderpb.CreateOrderResponse
}
var file_internal_infra_grpc_protofiles_order_proto_depIdxs = []int32{
	0, // 0: orderpb.ListOrdersResponse.orders:type_name -> orderpb.Order
	3, // 1: orderpb.OrderService.CreateOrder:input_type -> orderpb.CreateOrderRequest
	1, // 2: orderpb.OrderService.ListOrders:input_type -> orderpb.ListOrdersRequest
	4, // 3: orderpb.OrderService.CreateOrder:output_type -> orderpb.CreateOrderResponse
	2, // 4: orderpb.OrderService.ListOrders:output_type -> orderpb.ListOrdersResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_infra_grpc_protofiles_order_proto_init() }
func file_internal_infra_grpc_protofiles_order_proto_init() {
	if File_internal_infra_grpc_protofiles_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_infra_grpc_protofiles_order_proto_rawDesc), len(file_internal_infra_grpc_protofiles_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_infra_grpc_protofiles_order_proto_goTypes,
		DependencyIndexes: file_internal_infra_grpc_protofiles_order_proto_depIdxs,
		MessageInfos:      file_internal_infra_grpc_protofiles_order_proto_msgTypes,
	}.Build()
	File_internal_infra_grpc_protofiles_order_proto = out.File
	file_internal_infra_grpc_protofiles_order_proto_goTypes = nil
	file_internal_infra_grpc_protofiles_order_proto_depIdxs = nil
}
