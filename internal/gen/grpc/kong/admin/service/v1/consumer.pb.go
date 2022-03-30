// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: kong/admin/service/v1/consumer.proto

package v1

import (
	v1 "github.com/kong/koko/internal/gen/grpc/kong/admin/model/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *GetConsumerRequest) Reset() {
	*x = GetConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumerRequest) ProtoMessage() {}

func (x *GetConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumerRequest.ProtoReflect.Descriptor instead.
func (*GetConsumerRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_consumer_proto_rawDescGZIP(), []int{0}
}

func (x *GetConsumerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetConsumerRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type GetConsumerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.Consumer `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetConsumerResponse) Reset() {
	*x = GetConsumerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumerResponse) ProtoMessage() {}

func (x *GetConsumerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumerResponse.ProtoReflect.Descriptor instead.
func (*GetConsumerResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_consumer_proto_rawDescGZIP(), []int{1}
}

func (x *GetConsumerResponse) GetItem() *v1.Consumer {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreateConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item    *v1.Consumer       `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *CreateConsumerRequest) Reset() {
	*x = CreateConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConsumerRequest) ProtoMessage() {}

func (x *CreateConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConsumerRequest.ProtoReflect.Descriptor instead.
func (*CreateConsumerRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_consumer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateConsumerRequest) GetItem() *v1.Consumer {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *CreateConsumerRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type CreateConsumerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.Consumer `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateConsumerResponse) Reset() {
	*x = CreateConsumerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConsumerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConsumerResponse) ProtoMessage() {}

func (x *CreateConsumerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConsumerResponse.ProtoReflect.Descriptor instead.
func (*CreateConsumerResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_consumer_proto_rawDescGZIP(), []int{3}
}

func (x *CreateConsumerResponse) GetItem() *v1.Consumer {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpsertConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item    *v1.Consumer       `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *UpsertConsumerRequest) Reset() {
	*x = UpsertConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertConsumerRequest) ProtoMessage() {}

func (x *UpsertConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertConsumerRequest.ProtoReflect.Descriptor instead.
func (*UpsertConsumerRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_consumer_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertConsumerRequest) GetItem() *v1.Consumer {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *UpsertConsumerRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type UpsertConsumerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.Consumer `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *UpsertConsumerResponse) Reset() {
	*x = UpsertConsumerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertConsumerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertConsumerResponse) ProtoMessage() {}

func (x *UpsertConsumerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertConsumerResponse.ProtoReflect.Descriptor instead.
func (*UpsertConsumerResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_consumer_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertConsumerResponse) GetItem() *v1.Consumer {
	if x != nil {
		return x.Item
	}
	return nil
}

type DeleteConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *DeleteConsumerRequest) Reset() {
	*x = DeleteConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConsumerRequest) ProtoMessage() {}

func (x *DeleteConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConsumerRequest.ProtoReflect.Descriptor instead.
func (*DeleteConsumerRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_consumer_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteConsumerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteConsumerRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type DeleteConsumerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteConsumerResponse) Reset() {
	*x = DeleteConsumerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConsumerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConsumerResponse) ProtoMessage() {}

func (x *DeleteConsumerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConsumerResponse.ProtoReflect.Descriptor instead.
func (*DeleteConsumerResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_consumer_proto_rawDescGZIP(), []int{7}
}

type ListConsumersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *v1.RequestCluster    `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Page    *v1.PaginationRequest `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListConsumersRequest) Reset() {
	*x = ListConsumersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConsumersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConsumersRequest) ProtoMessage() {}

func (x *ListConsumersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConsumersRequest.ProtoReflect.Descriptor instead.
func (*ListConsumersRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_consumer_proto_rawDescGZIP(), []int{8}
}

func (x *ListConsumersRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *ListConsumersRequest) GetPage() *v1.PaginationRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListConsumersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*v1.Consumer         `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Page  *v1.PaginationResponse `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListConsumersResponse) Reset() {
	*x = ListConsumersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConsumersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConsumersResponse) ProtoMessage() {}

func (x *ListConsumersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_consumer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConsumersResponse.ProtoReflect.Descriptor instead.
func (*ListConsumersResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_consumer_proto_rawDescGZIP(), []int{9}
}

func (x *ListConsumersResponse) GetItems() []*v1.Consumer {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListConsumersResponse) GetPage() *v1.PaginationResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_kong_admin_service_v1_consumer_proto protoreflect.FileDescriptor

var file_kong_admin_service_v1_consumer_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6b, 0x6f, 0x6e,
	0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d,
	0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x48, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x89, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x6f,
	0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x22, 0x89, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x3d, 0x0a,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x16,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x66, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x89, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3b,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6b,
	0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0xc8, 0x05, 0x0a, 0x0f,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x80, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12,
	0x29, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6b, 0x6f, 0x6e,
	0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x8a, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12,
	0x94, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x12, 0x2c, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x1a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x64, 0x7d,
	0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x89, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x6b, 0x6f, 0x6e, 0x67,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x81, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x6b, 0x6f, 0x6e, 0x67,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kong_admin_service_v1_consumer_proto_rawDescOnce sync.Once
	file_kong_admin_service_v1_consumer_proto_rawDescData = file_kong_admin_service_v1_consumer_proto_rawDesc
)

func file_kong_admin_service_v1_consumer_proto_rawDescGZIP() []byte {
	file_kong_admin_service_v1_consumer_proto_rawDescOnce.Do(func() {
		file_kong_admin_service_v1_consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_kong_admin_service_v1_consumer_proto_rawDescData)
	})
	return file_kong_admin_service_v1_consumer_proto_rawDescData
}

var file_kong_admin_service_v1_consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_kong_admin_service_v1_consumer_proto_goTypes = []interface{}{
	(*GetConsumerRequest)(nil),     // 0: kong.admin.service.v1.GetConsumerRequest
	(*GetConsumerResponse)(nil),    // 1: kong.admin.service.v1.GetConsumerResponse
	(*CreateConsumerRequest)(nil),  // 2: kong.admin.service.v1.CreateConsumerRequest
	(*CreateConsumerResponse)(nil), // 3: kong.admin.service.v1.CreateConsumerResponse
	(*UpsertConsumerRequest)(nil),  // 4: kong.admin.service.v1.UpsertConsumerRequest
	(*UpsertConsumerResponse)(nil), // 5: kong.admin.service.v1.UpsertConsumerResponse
	(*DeleteConsumerRequest)(nil),  // 6: kong.admin.service.v1.DeleteConsumerRequest
	(*DeleteConsumerResponse)(nil), // 7: kong.admin.service.v1.DeleteConsumerResponse
	(*ListConsumersRequest)(nil),   // 8: kong.admin.service.v1.ListConsumersRequest
	(*ListConsumersResponse)(nil),  // 9: kong.admin.service.v1.ListConsumersResponse
	(*v1.RequestCluster)(nil),      // 10: kong.admin.model.v1.RequestCluster
	(*v1.Consumer)(nil),            // 11: kong.admin.model.v1.Consumer
	(*v1.PaginationRequest)(nil),   // 12: kong.admin.model.v1.PaginationRequest
	(*v1.PaginationResponse)(nil),  // 13: kong.admin.model.v1.PaginationResponse
}
var file_kong_admin_service_v1_consumer_proto_depIdxs = []int32{
	10, // 0: kong.admin.service.v1.GetConsumerRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	11, // 1: kong.admin.service.v1.GetConsumerResponse.item:type_name -> kong.admin.model.v1.Consumer
	11, // 2: kong.admin.service.v1.CreateConsumerRequest.item:type_name -> kong.admin.model.v1.Consumer
	10, // 3: kong.admin.service.v1.CreateConsumerRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	11, // 4: kong.admin.service.v1.CreateConsumerResponse.item:type_name -> kong.admin.model.v1.Consumer
	11, // 5: kong.admin.service.v1.UpsertConsumerRequest.item:type_name -> kong.admin.model.v1.Consumer
	10, // 6: kong.admin.service.v1.UpsertConsumerRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	11, // 7: kong.admin.service.v1.UpsertConsumerResponse.item:type_name -> kong.admin.model.v1.Consumer
	10, // 8: kong.admin.service.v1.DeleteConsumerRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	10, // 9: kong.admin.service.v1.ListConsumersRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	12, // 10: kong.admin.service.v1.ListConsumersRequest.page:type_name -> kong.admin.model.v1.PaginationRequest
	11, // 11: kong.admin.service.v1.ListConsumersResponse.items:type_name -> kong.admin.model.v1.Consumer
	13, // 12: kong.admin.service.v1.ListConsumersResponse.page:type_name -> kong.admin.model.v1.PaginationResponse
	0,  // 13: kong.admin.service.v1.ConsumerService.GetConsumer:input_type -> kong.admin.service.v1.GetConsumerRequest
	2,  // 14: kong.admin.service.v1.ConsumerService.CreateConsumer:input_type -> kong.admin.service.v1.CreateConsumerRequest
	4,  // 15: kong.admin.service.v1.ConsumerService.UpsertConsumer:input_type -> kong.admin.service.v1.UpsertConsumerRequest
	6,  // 16: kong.admin.service.v1.ConsumerService.DeleteConsumer:input_type -> kong.admin.service.v1.DeleteConsumerRequest
	8,  // 17: kong.admin.service.v1.ConsumerService.ListConsumers:input_type -> kong.admin.service.v1.ListConsumersRequest
	1,  // 18: kong.admin.service.v1.ConsumerService.GetConsumer:output_type -> kong.admin.service.v1.GetConsumerResponse
	3,  // 19: kong.admin.service.v1.ConsumerService.CreateConsumer:output_type -> kong.admin.service.v1.CreateConsumerResponse
	5,  // 20: kong.admin.service.v1.ConsumerService.UpsertConsumer:output_type -> kong.admin.service.v1.UpsertConsumerResponse
	7,  // 21: kong.admin.service.v1.ConsumerService.DeleteConsumer:output_type -> kong.admin.service.v1.DeleteConsumerResponse
	9,  // 22: kong.admin.service.v1.ConsumerService.ListConsumers:output_type -> kong.admin.service.v1.ListConsumersResponse
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_kong_admin_service_v1_consumer_proto_init() }
func file_kong_admin_service_v1_consumer_proto_init() {
	if File_kong_admin_service_v1_consumer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kong_admin_service_v1_consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumerRequest); i {
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
		file_kong_admin_service_v1_consumer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumerResponse); i {
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
		file_kong_admin_service_v1_consumer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConsumerRequest); i {
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
		file_kong_admin_service_v1_consumer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConsumerResponse); i {
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
		file_kong_admin_service_v1_consumer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertConsumerRequest); i {
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
		file_kong_admin_service_v1_consumer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertConsumerResponse); i {
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
		file_kong_admin_service_v1_consumer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConsumerRequest); i {
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
		file_kong_admin_service_v1_consumer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConsumerResponse); i {
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
		file_kong_admin_service_v1_consumer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConsumersRequest); i {
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
		file_kong_admin_service_v1_consumer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConsumersResponse); i {
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
			RawDescriptor: file_kong_admin_service_v1_consumer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kong_admin_service_v1_consumer_proto_goTypes,
		DependencyIndexes: file_kong_admin_service_v1_consumer_proto_depIdxs,
		MessageInfos:      file_kong_admin_service_v1_consumer_proto_msgTypes,
	}.Build()
	File_kong_admin_service_v1_consumer_proto = out.File
	file_kong_admin_service_v1_consumer_proto_rawDesc = nil
	file_kong_admin_service_v1_consumer_proto_goTypes = nil
	file_kong_admin_service_v1_consumer_proto_depIdxs = nil
}
