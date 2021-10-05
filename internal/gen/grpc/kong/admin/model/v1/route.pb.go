// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: kong/admin/model/v1/route.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                    string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Headers                 map[string]*HeaderValues `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Hosts                   []string                 `protobuf:"bytes,4,rep,name=hosts,proto3" json:"hosts,omitempty"`
	CreatedAt               int32                    `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Methods                 []string                 `protobuf:"bytes,6,rep,name=methods,proto3" json:"methods,omitempty"`
	Paths                   []string                 `protobuf:"bytes,7,rep,name=paths,proto3" json:"paths,omitempty"`
	PathHandling            string                   `protobuf:"bytes,8,opt,name=path_handling,json=pathHandling,proto3" json:"path_handling,omitempty"`
	PreserveHost            *wrapperspb.BoolValue    `protobuf:"bytes,9,opt,name=preserve_host,json=preserveHost,proto3" json:"preserve_host,omitempty"`
	Protocols               []string                 `protobuf:"bytes,10,rep,name=protocols,proto3" json:"protocols,omitempty"`
	RegexPriority           *wrapperspb.Int32Value   `protobuf:"bytes,11,opt,name=regex_priority,json=regexPriority,proto3" json:"regex_priority,omitempty"`
	StripPath               *wrapperspb.BoolValue    `protobuf:"bytes,12,opt,name=strip_path,json=stripPath,proto3" json:"strip_path,omitempty"`
	UpdatedAt               int32                    `protobuf:"varint,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Snis                    []string                 `protobuf:"bytes,14,rep,name=snis,proto3" json:"snis,omitempty"`
	Sources                 []*CIDRPort              `protobuf:"bytes,15,rep,name=sources,proto3" json:"sources,omitempty"`
	Destinations            []*CIDRPort              `protobuf:"bytes,16,rep,name=destinations,proto3" json:"destinations,omitempty"`
	Tags                    []string                 `protobuf:"bytes,17,rep,name=tags,proto3" json:"tags,omitempty"`
	HttpsRedirectStatusCode int32                    `protobuf:"varint,18,opt,name=https_redirect_status_code,json=httpsRedirectStatusCode,proto3" json:"https_redirect_status_code,omitempty"`
	RequestBuffering        *wrapperspb.BoolValue    `protobuf:"bytes,19,opt,name=request_buffering,json=requestBuffering,proto3" json:"request_buffering,omitempty"`
	ResponseBuffering       *wrapperspb.BoolValue    `protobuf:"bytes,20,opt,name=response_buffering,json=responseBuffering,proto3" json:"response_buffering,omitempty"`
	Service                 *Service                 `protobuf:"bytes,21,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_model_v1_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_model_v1_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_kong_admin_model_v1_route_proto_rawDescGZIP(), []int{0}
}

func (x *Route) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Route) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Route) GetHeaders() map[string]*HeaderValues {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Route) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *Route) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Route) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Route) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *Route) GetPathHandling() string {
	if x != nil {
		return x.PathHandling
	}
	return ""
}

func (x *Route) GetPreserveHost() *wrapperspb.BoolValue {
	if x != nil {
		return x.PreserveHost
	}
	return nil
}

func (x *Route) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *Route) GetRegexPriority() *wrapperspb.Int32Value {
	if x != nil {
		return x.RegexPriority
	}
	return nil
}

func (x *Route) GetStripPath() *wrapperspb.BoolValue {
	if x != nil {
		return x.StripPath
	}
	return nil
}

func (x *Route) GetUpdatedAt() int32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Route) GetSnis() []string {
	if x != nil {
		return x.Snis
	}
	return nil
}

func (x *Route) GetSources() []*CIDRPort {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *Route) GetDestinations() []*CIDRPort {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *Route) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Route) GetHttpsRedirectStatusCode() int32 {
	if x != nil {
		return x.HttpsRedirectStatusCode
	}
	return 0
}

func (x *Route) GetRequestBuffering() *wrapperspb.BoolValue {
	if x != nil {
		return x.RequestBuffering
	}
	return nil
}

func (x *Route) GetResponseBuffering() *wrapperspb.BoolValue {
	if x != nil {
		return x.ResponseBuffering
	}
	return nil
}

func (x *Route) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

type HeaderValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *HeaderValues) Reset() {
	*x = HeaderValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_model_v1_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValues) ProtoMessage() {}

func (x *HeaderValues) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_model_v1_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValues.ProtoReflect.Descriptor instead.
func (*HeaderValues) Descriptor() ([]byte, []int) {
	return file_kong_admin_model_v1_route_proto_rawDescGZIP(), []int{1}
}

func (x *HeaderValues) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type CIDRPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *CIDRPort) Reset() {
	*x = CIDRPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_admin_model_v1_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIDRPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIDRPort) ProtoMessage() {}

func (x *CIDRPort) ProtoReflect() protoreflect.Message {
	mi := &file_kong_admin_model_v1_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIDRPort.ProtoReflect.Descriptor instead.
func (*CIDRPort) Descriptor() ([]byte, []int) {
	return file_kong_admin_model_v1_route_proto_rawDescGZIP(), []int{2}
}

func (x *CIDRPort) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CIDRPort) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_kong_admin_model_v1_route_proto protoreflect.FileDescriptor

var file_kong_admin_model_v1_route_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x08, 0x0a, 0x05, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f,
	0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74,
	0x68, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x74, 0x68, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x69, 0x6e, 0x67, 0x12, 0x3f, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x65, 0x78, 0x5f, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x67, 0x65, 0x78, 0x50,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x70,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x70, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6e, 0x69, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x6e, 0x69, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x49, 0x44,
	0x52, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x41,
	0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x10,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x49, 0x44, 0x52, 0x50,
	0x6f, 0x72, 0x74, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x5f, 0x72,
	0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x47, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x49, 0x0a, 0x12, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x5d,
	0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a,
	0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x43, 0x49, 0x44, 0x52, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kong_admin_model_v1_route_proto_rawDescOnce sync.Once
	file_kong_admin_model_v1_route_proto_rawDescData = file_kong_admin_model_v1_route_proto_rawDesc
)

func file_kong_admin_model_v1_route_proto_rawDescGZIP() []byte {
	file_kong_admin_model_v1_route_proto_rawDescOnce.Do(func() {
		file_kong_admin_model_v1_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_kong_admin_model_v1_route_proto_rawDescData)
	})
	return file_kong_admin_model_v1_route_proto_rawDescData
}

var file_kong_admin_model_v1_route_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kong_admin_model_v1_route_proto_goTypes = []interface{}{
	(*Route)(nil),                 // 0: kong.admin.model.v1.Route
	(*HeaderValues)(nil),          // 1: kong.admin.model.v1.HeaderValues
	(*CIDRPort)(nil),              // 2: kong.admin.model.v1.CIDRPort
	nil,                           // 3: kong.admin.model.v1.Route.HeadersEntry
	(*wrapperspb.BoolValue)(nil),  // 4: google.protobuf.BoolValue
	(*wrapperspb.Int32Value)(nil), // 5: google.protobuf.Int32Value
	(*Service)(nil),               // 6: kong.admin.model.v1.Service
}
var file_kong_admin_model_v1_route_proto_depIdxs = []int32{
	3,  // 0: kong.admin.model.v1.Route.headers:type_name -> kong.admin.model.v1.Route.HeadersEntry
	4,  // 1: kong.admin.model.v1.Route.preserve_host:type_name -> google.protobuf.BoolValue
	5,  // 2: kong.admin.model.v1.Route.regex_priority:type_name -> google.protobuf.Int32Value
	4,  // 3: kong.admin.model.v1.Route.strip_path:type_name -> google.protobuf.BoolValue
	2,  // 4: kong.admin.model.v1.Route.sources:type_name -> kong.admin.model.v1.CIDRPort
	2,  // 5: kong.admin.model.v1.Route.destinations:type_name -> kong.admin.model.v1.CIDRPort
	4,  // 6: kong.admin.model.v1.Route.request_buffering:type_name -> google.protobuf.BoolValue
	4,  // 7: kong.admin.model.v1.Route.response_buffering:type_name -> google.protobuf.BoolValue
	6,  // 8: kong.admin.model.v1.Route.service:type_name -> kong.admin.model.v1.Service
	1,  // 9: kong.admin.model.v1.Route.HeadersEntry.value:type_name -> kong.admin.model.v1.HeaderValues
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_kong_admin_model_v1_route_proto_init() }
func file_kong_admin_model_v1_route_proto_init() {
	if File_kong_admin_model_v1_route_proto != nil {
		return
	}
	file_kong_admin_model_v1_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kong_admin_model_v1_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
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
		file_kong_admin_model_v1_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValues); i {
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
		file_kong_admin_model_v1_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIDRPort); i {
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
			RawDescriptor: file_kong_admin_model_v1_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kong_admin_model_v1_route_proto_goTypes,
		DependencyIndexes: file_kong_admin_model_v1_route_proto_depIdxs,
		MessageInfos:      file_kong_admin_model_v1_route_proto_msgTypes,
	}.Build()
	File_kong_admin_model_v1_route_proto = out.File
	file_kong_admin_model_v1_route_proto_rawDesc = nil
	file_kong_admin_model_v1_route_proto_goTypes = nil
	file_kong_admin_model_v1_route_proto_depIdxs = nil
}
