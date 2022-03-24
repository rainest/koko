// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: kong/services/event_hooks/v1/event_hooks.proto

package v1

import (
	events "github.com/kong/koko/internal/gen/wrpc/kong/model/events"
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

type SyncHooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hooks []*events.EventHook `protobuf:"bytes,1,rep,name=hooks,proto3" json:"hooks,omitempty"`
	// On every configuration change, CP MUST increment the version field
	// in the request.
	// Version field has no significance outside the context of a single ephemeral
	// connection between a DP node and a CP node.
	Version uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SyncHooksRequest) Reset() {
	*x = SyncHooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_services_event_hooks_v1_event_hooks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncHooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncHooksRequest) ProtoMessage() {}

func (x *SyncHooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kong_services_event_hooks_v1_event_hooks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncHooksRequest.ProtoReflect.Descriptor instead.
func (*SyncHooksRequest) Descriptor() ([]byte, []int) {
	return file_kong_services_event_hooks_v1_event_hooks_proto_rawDescGZIP(), []int{0}
}

func (x *SyncHooksRequest) GetHooks() []*events.EventHook {
	if x != nil {
		return x.Hooks
	}
	return nil
}

func (x *SyncHooksRequest) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

type SyncHooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// accepted is set to true when the DP has accepted the configuration.
	// Acceptance of configuration indicates that the configuration is successfully
	// processed by the DP.
	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	// If accepted is set to false, errors denote the errors with the configuration.
	// CP MAY analyze the errors and send back a correct configuration.
	// If accepted is true, this field must be empty
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *SyncHooksResponse) Reset() {
	*x = SyncHooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kong_services_event_hooks_v1_event_hooks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncHooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncHooksResponse) ProtoMessage() {}

func (x *SyncHooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kong_services_event_hooks_v1_event_hooks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncHooksResponse.ProtoReflect.Descriptor instead.
func (*SyncHooksResponse) Descriptor() ([]byte, []int) {
	return file_kong_services_event_hooks_v1_event_hooks_proto_rawDescGZIP(), []int{1}
}

func (x *SyncHooksResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

func (x *SyncHooksResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_kong_services_event_hooks_v1_event_hooks_proto protoreflect.FileDescriptor

var file_kong_services_event_hooks_v1_event_hooks_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1c, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x23,
	0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x10, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x6f, 0x6f, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x68, 0x6f, 0x6f, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x6f, 0x6f,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x32, 0x81,
	0x01, 0x0a, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x6f, 0x6f, 0x6b,
	0x73, 0x12, 0x2e, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2f, 0x2e, 0x6b, 0x6f, 0x6e, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x6f, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x6b, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x77, 0x72, 0x70, 0x63, 0x2f, 0x6b, 0x6f, 0x6e,
	0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68,
	0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_kong_services_event_hooks_v1_event_hooks_proto_rawDescOnce sync.Once
	file_kong_services_event_hooks_v1_event_hooks_proto_rawDescData = file_kong_services_event_hooks_v1_event_hooks_proto_rawDesc
)

func file_kong_services_event_hooks_v1_event_hooks_proto_rawDescGZIP() []byte {
	file_kong_services_event_hooks_v1_event_hooks_proto_rawDescOnce.Do(func() {
		file_kong_services_event_hooks_v1_event_hooks_proto_rawDescData = protoimpl.X.CompressGZIP(file_kong_services_event_hooks_v1_event_hooks_proto_rawDescData)
	})
	return file_kong_services_event_hooks_v1_event_hooks_proto_rawDescData
}

var file_kong_services_event_hooks_v1_event_hooks_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kong_services_event_hooks_v1_event_hooks_proto_goTypes = []interface{}{
	(*SyncHooksRequest)(nil),  // 0: kong.services.event_hooks.v1.SyncHooksRequest
	(*SyncHooksResponse)(nil), // 1: kong.services.event_hooks.v1.SyncHooksResponse
	(*events.EventHook)(nil),  // 2: kong.model.events.EventHook
}
var file_kong_services_event_hooks_v1_event_hooks_proto_depIdxs = []int32{
	2, // 0: kong.services.event_hooks.v1.SyncHooksRequest.hooks:type_name -> kong.model.events.EventHook
	0, // 1: kong.services.event_hooks.v1.EventHooksService.SyncHooks:input_type -> kong.services.event_hooks.v1.SyncHooksRequest
	1, // 2: kong.services.event_hooks.v1.EventHooksService.SyncHooks:output_type -> kong.services.event_hooks.v1.SyncHooksResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kong_services_event_hooks_v1_event_hooks_proto_init() }
func file_kong_services_event_hooks_v1_event_hooks_proto_init() {
	if File_kong_services_event_hooks_v1_event_hooks_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kong_services_event_hooks_v1_event_hooks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncHooksRequest); i {
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
		file_kong_services_event_hooks_v1_event_hooks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncHooksResponse); i {
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
			RawDescriptor: file_kong_services_event_hooks_v1_event_hooks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kong_services_event_hooks_v1_event_hooks_proto_goTypes,
		DependencyIndexes: file_kong_services_event_hooks_v1_event_hooks_proto_depIdxs,
		MessageInfos:      file_kong_services_event_hooks_v1_event_hooks_proto_msgTypes,
	}.Build()
	File_kong_services_event_hooks_v1_event_hooks_proto = out.File
	file_kong_services_event_hooks_v1_event_hooks_proto_rawDesc = nil
	file_kong_services_event_hooks_v1_event_hooks_proto_goTypes = nil
	file_kong_services_event_hooks_v1_event_hooks_proto_depIdxs = nil
}
