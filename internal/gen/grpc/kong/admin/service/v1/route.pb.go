// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: kong/admin/service/v1/route.proto

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

type GetRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *GetRouteRequest) Reset() {
	*x = GetRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteRequest) ProtoMessage() {}

func (x *GetRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteRequest.ProtoReflect.Descriptor instead.
func (*GetRouteRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_route_proto_rawDescGZIP(), []int{0}
}

func (x *GetRouteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetRouteRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type GetRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.Route `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetRouteResponse) Reset() {
	*x = GetRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteResponse) ProtoMessage() {}

func (x *GetRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteResponse.ProtoReflect.Descriptor instead.
func (*GetRouteResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_route_proto_rawDescGZIP(), []int{1}
}

func (x *GetRouteResponse) GetItem() *v1.Route {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreateRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item    *v1.Route          `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *CreateRouteRequest) Reset() {
	*x = CreateRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRouteRequest) ProtoMessage() {}

func (x *CreateRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRouteRequest.ProtoReflect.Descriptor instead.
func (*CreateRouteRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_route_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRouteRequest) GetItem() *v1.Route {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *CreateRouteRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type CreateRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.Route `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateRouteResponse) Reset() {
	*x = CreateRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRouteResponse) ProtoMessage() {}

func (x *CreateRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRouteResponse.ProtoReflect.Descriptor instead.
func (*CreateRouteResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_route_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRouteResponse) GetItem() *v1.Route {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpsertRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item    *v1.Route          `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *UpsertRouteRequest) Reset() {
	*x = UpsertRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertRouteRequest) ProtoMessage() {}

func (x *UpsertRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertRouteRequest.ProtoReflect.Descriptor instead.
func (*UpsertRouteRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_route_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertRouteRequest) GetItem() *v1.Route {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *UpsertRouteRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type UpsertRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *v1.Route `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *UpsertRouteResponse) Reset() {
	*x = UpsertRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertRouteResponse) ProtoMessage() {}

func (x *UpsertRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertRouteResponse.ProtoReflect.Descriptor instead.
func (*UpsertRouteResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_route_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertRouteResponse) GetItem() *v1.Route {
	if x != nil {
		return x.Item
	}
	return nil
}

type DeleteRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cluster *v1.RequestCluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *DeleteRouteRequest) Reset() {
	*x = DeleteRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_route_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRouteRequest) ProtoMessage() {}

func (x *DeleteRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_route_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRouteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRouteRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_route_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRouteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteRouteRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type DeleteRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRouteResponse) Reset() {
	*x = DeleteRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_route_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRouteResponse) ProtoMessage() {}

func (x *DeleteRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_route_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRouteResponse.ProtoReflect.Descriptor instead.
func (*DeleteRouteResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_route_proto_rawDescGZIP(), []int{7}
}

type ListRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster   *v1.RequestCluster    `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	ServiceId string                `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Page      *v1.PaginationRequest `protobuf:"bytes,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListRoutesRequest) Reset() {
	*x = ListRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_route_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutesRequest) ProtoMessage() {}

func (x *ListRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_route_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutesRequest.ProtoReflect.Descriptor instead.
func (*ListRoutesRequest) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_route_proto_rawDescGZIP(), []int{8}
}

func (x *ListRoutesRequest) GetCluster() *v1.RequestCluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *ListRoutesRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *ListRoutesRequest) GetPage() *v1.PaginationRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListRoutesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*v1.Route            `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Page  *v1.PaginationResponse `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListRoutesResponse) Reset() {
	*x = ListRoutesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_service_v1_route_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoutesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoutesResponse) ProtoMessage() {}

func (x *ListRoutesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_service_v1_route_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoutesResponse.ProtoReflect.Descriptor instead.
func (*ListRoutesResponse) Descriptor() ([]byte, []int) {
	return file_kong_admin_service_v1_route_proto_rawDescGZIP(), []int{9}
}

func (x *ListRoutesResponse) GetItems() []*v1.Route {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListRoutesResponse) GetPage() *v1.PaginationResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_kong_admin_service_v1_route_proto protoreflect.FileDescriptor

var file_kong_admin_service_v1_route_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6b, 0x6f, 0x6e,
	0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x60, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12,
	0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x45,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x6f, 0x6e,
	0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x3d, 0x0a, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x13, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x22, 0x63, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xad,
	0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x3a, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x83,
	0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x32, 0x85, 0x05, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x12, 0x26, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6b, 0x6f, 0x6e, 0x67,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6b, 0x6f, 0x6e,
	0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x88, 0x01, 0x0a, 0x0b,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6b, 0x6f,
	0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x1a, 0x14, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x69, 0x64, 0x7d,
	0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x7d, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x75, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c,
	0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f,
	0x6b, 0x6f, 0x6b, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_kong_admin_service_v1_route_proto_rawDescOnce sync.Once
	file_kong_admin_service_v1_route_proto_rawDescData = file_kong_admin_service_v1_route_proto_rawDesc
)

func file_kong_admin_service_v1_route_proto_rawDescGZIP() []byte {
	file_kong_admin_service_v1_route_proto_rawDescOnce.Do(func() {
		file_kong_admin_service_v1_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_kong_admin_service_v1_route_proto_rawDescData)
	})
	return file_kong_admin_service_v1_route_proto_rawDescData
}

var file_kong_admin_service_v1_route_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_kong_admin_service_v1_route_proto_goTypes = []interface{}{
	(*GetRouteRequest)(nil),       // 0: kong.admin.service.v1.GetRouteRequest
	(*GetRouteResponse)(nil),      // 1: kong.admin.service.v1.GetRouteResponse
	(*CreateRouteRequest)(nil),    // 2: kong.admin.service.v1.CreateRouteRequest
	(*CreateRouteResponse)(nil),   // 3: kong.admin.service.v1.CreateRouteResponse
	(*UpsertRouteRequest)(nil),    // 4: kong.admin.service.v1.UpsertRouteRequest
	(*UpsertRouteResponse)(nil),   // 5: kong.admin.service.v1.UpsertRouteResponse
	(*DeleteRouteRequest)(nil),    // 6: kong.admin.service.v1.DeleteRouteRequest
	(*DeleteRouteResponse)(nil),   // 7: kong.admin.service.v1.DeleteRouteResponse
	(*ListRoutesRequest)(nil),     // 8: kong.admin.service.v1.ListRoutesRequest
	(*ListRoutesResponse)(nil),    // 9: kong.admin.service.v1.ListRoutesResponse
	(*v1.RequestCluster)(nil),     // 10: kong.admin.model.v1.RequestCluster
	(*v1.Route)(nil),              // 11: kong.admin.model.v1.Route
	(*v1.PaginationRequest)(nil),  // 12: kong.admin.model.v1.PaginationRequest
	(*v1.PaginationResponse)(nil), // 13: kong.admin.model.v1.PaginationResponse
}
var file_kong_admin_service_v1_route_proto_depIdxs = []int32{
	10, // 0: kong.admin.service.v1.GetRouteRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	11, // 1: kong.admin.service.v1.GetRouteResponse.item:type_name -> kong.admin.model.v1.Route
	11, // 2: kong.admin.service.v1.CreateRouteRequest.item:type_name -> kong.admin.model.v1.Route
	10, // 3: kong.admin.service.v1.CreateRouteRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	11, // 4: kong.admin.service.v1.CreateRouteResponse.item:type_name -> kong.admin.model.v1.Route
	11, // 5: kong.admin.service.v1.UpsertRouteRequest.item:type_name -> kong.admin.model.v1.Route
	10, // 6: kong.admin.service.v1.UpsertRouteRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	11, // 7: kong.admin.service.v1.UpsertRouteResponse.item:type_name -> kong.admin.model.v1.Route
	10, // 8: kong.admin.service.v1.DeleteRouteRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	10, // 9: kong.admin.service.v1.ListRoutesRequest.cluster:type_name -> kong.admin.model.v1.RequestCluster
	12, // 10: kong.admin.service.v1.ListRoutesRequest.page:type_name -> kong.admin.model.v1.PaginationRequest
	11, // 11: kong.admin.service.v1.ListRoutesResponse.items:type_name -> kong.admin.model.v1.Route
	13, // 12: kong.admin.service.v1.ListRoutesResponse.page:type_name -> kong.admin.model.v1.PaginationResponse
	0,  // 13: kong.admin.service.v1.RouteService.GetRoute:input_type -> kong.admin.service.v1.GetRouteRequest
	2,  // 14: kong.admin.service.v1.RouteService.CreateRoute:input_type -> kong.admin.service.v1.CreateRouteRequest
	4,  // 15: kong.admin.service.v1.RouteService.UpsertRoute:input_type -> kong.admin.service.v1.UpsertRouteRequest
	6,  // 16: kong.admin.service.v1.RouteService.DeleteRoute:input_type -> kong.admin.service.v1.DeleteRouteRequest
	8,  // 17: kong.admin.service.v1.RouteService.ListRoutes:input_type -> kong.admin.service.v1.ListRoutesRequest
	1,  // 18: kong.admin.service.v1.RouteService.GetRoute:output_type -> kong.admin.service.v1.GetRouteResponse
	3,  // 19: kong.admin.service.v1.RouteService.CreateRoute:output_type -> kong.admin.service.v1.CreateRouteResponse
	5,  // 20: kong.admin.service.v1.RouteService.UpsertRoute:output_type -> kong.admin.service.v1.UpsertRouteResponse
	7,  // 21: kong.admin.service.v1.RouteService.DeleteRoute:output_type -> kong.admin.service.v1.DeleteRouteResponse
	9,  // 22: kong.admin.service.v1.RouteService.ListRoutes:output_type -> kong.admin.service.v1.ListRoutesResponse
	18, // [18:23] is the sub-list for method output_type
	13, // [13:18] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_kong_admin_service_v1_route_proto_init() }
func file_kong_admin_service_v1_route_proto_init() {
	if File_kong_admin_service_v1_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kong_admin_service_v1_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRouteRequest); i {
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
		file_kong_admin_service_v1_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRouteResponse); i {
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
		file_kong_admin_service_v1_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRouteRequest); i {
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
		file_kong_admin_service_v1_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRouteResponse); i {
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
		file_kong_admin_service_v1_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertRouteRequest); i {
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
		file_kong_admin_service_v1_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertRouteResponse); i {
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
		file_kong_admin_service_v1_route_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRouteRequest); i {
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
		file_kong_admin_service_v1_route_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRouteResponse); i {
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
		file_kong_admin_service_v1_route_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoutesRequest); i {
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
		file_kong_admin_service_v1_route_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoutesResponse); i {
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
			RawDescriptor: file_kong_admin_service_v1_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kong_admin_service_v1_route_proto_goTypes,
		DependencyIndexes: file_kong_admin_service_v1_route_proto_depIdxs,
		MessageInfos:      file_kong_admin_service_v1_route_proto_msgTypes,
	}.Build()
	File_kong_admin_service_v1_route_proto = out.File
	file_kong_admin_service_v1_route_proto_rawDesc = nil
	file_kong_admin_service_v1_route_proto_goTypes = nil
	file_kong_admin_service_v1_route_proto_depIdxs = nil
}
