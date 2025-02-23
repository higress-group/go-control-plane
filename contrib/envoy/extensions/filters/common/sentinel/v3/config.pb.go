// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: contrib/envoy/extensions/filters/common/sentinel/v3/config.proto

package sentinelv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// [#protodoc-title: Sentinel]
// [#comment:next free field: 2]
// [#next-free-field: 9]
type CommonConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName                string `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	AppType                uint32 `protobuf:"varint,2,opt,name=app_type,json=appType,proto3" json:"app_type,omitempty"`
	SystemMetricIntervalMs uint32 `protobuf:"varint,3,opt,name=system_metric_interval_ms,json=systemMetricIntervalMs,proto3" json:"system_metric_interval_ms,omitempty"`
	UseCommandCenter       bool   `protobuf:"varint,4,opt,name=use_command_center,json=useCommandCenter,proto3" json:"use_command_center,omitempty"`
	AhasRegionId           string `protobuf:"bytes,5,opt,name=ahas_region_id,json=ahasRegionId,proto3" json:"ahas_region_id,omitempty"`
	License                string `protobuf:"bytes,6,opt,name=license,proto3" json:"license,omitempty"`
	AppNamespace           string `protobuf:"bytes,7,opt,name=app_namespace,json=appNamespace,proto3" json:"app_namespace,omitempty"`
	LogDir                 string `protobuf:"bytes,8,opt,name=log_dir,json=logDir,proto3" json:"log_dir,omitempty"`
}

func (x *CommonConfig) Reset() {
	*x = CommonConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonConfig) ProtoMessage() {}

func (x *CommonConfig) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonConfig.ProtoReflect.Descriptor instead.
func (*CommonConfig) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDescGZIP(), []int{0}
}

func (x *CommonConfig) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *CommonConfig) GetAppType() uint32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *CommonConfig) GetSystemMetricIntervalMs() uint32 {
	if x != nil {
		return x.SystemMetricIntervalMs
	}
	return 0
}

func (x *CommonConfig) GetUseCommandCenter() bool {
	if x != nil {
		return x.UseCommandCenter
	}
	return false
}

func (x *CommonConfig) GetAhasRegionId() string {
	if x != nil {
		return x.AhasRegionId
	}
	return ""
}

func (x *CommonConfig) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *CommonConfig) GetAppNamespace() string {
	if x != nil {
		return x.AppNamespace
	}
	return ""
}

func (x *CommonConfig) GetLogDir() string {
	if x != nil {
		return x.LogDir
	}
	return ""
}

var File_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDesc = []byte{
	0x0a, 0x40, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6e,
	0x65, 0x6c, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x33, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x70, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x4d, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x75, 0x73, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x0e, 0x61, 0x68, 0x61, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x68, 0x61, 0x73, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x69, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x44, 0x69, 0x72, 0x42, 0xb1,
	0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x39, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65,
	0x6c, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x6e, 0x74,
	0x69, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6c,
	0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDescData = file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDesc
)

func file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDescData
}

var file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_goTypes = []interface{}{
	(*CommonConfig)(nil), // 0: envoy.extensions.filters.common.sentinel.v3.CommonConfig
}
var file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_init() }
func file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_init() {
	if File_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonConfig); i {
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
			RawDescriptor: file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_depIdxs,
		MessageInfos:      file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto = out.File
	file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_rawDesc = nil
	file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_goTypes = nil
	file_contrib_envoy_extensions_filters_common_sentinel_v3_config_proto_depIdxs = nil
}
