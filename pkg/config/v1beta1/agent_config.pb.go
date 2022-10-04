// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v1.0.0
// source: github.com/rancher/opni/pkg/config/v1beta1/agent_config.proto

package v1beta1

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

type RulesSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Discovery *DiscoverySpec `protobuf:"bytes,1,opt,name=discovery,proto3" json:"discovery,omitempty"`
}

func (x *RulesSpec) Reset() {
	*x = RulesSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RulesSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RulesSpec) ProtoMessage() {}

func (x *RulesSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RulesSpec.ProtoReflect.Descriptor instead.
func (*RulesSpec) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescGZIP(), []int{0}
}

func (x *RulesSpec) GetDiscovery() *DiscoverySpec {
	if x != nil {
		return x.Discovery
	}
	return nil
}

type DiscoverySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrometheusRules *PrometheusRulesSpec `protobuf:"bytes,1,opt,name=prometheusRules,proto3" json:"prometheusRules,omitempty"`
	Filesystem      *FilesystemRulesSpec `protobuf:"bytes,2,opt,name=filesystem,proto3" json:"filesystem,omitempty"`
	Interval        string               `protobuf:"bytes,3,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *DiscoverySpec) Reset() {
	*x = DiscoverySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverySpec) ProtoMessage() {}

func (x *DiscoverySpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverySpec.ProtoReflect.Descriptor instead.
func (*DiscoverySpec) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescGZIP(), []int{1}
}

func (x *DiscoverySpec) GetPrometheusRules() *PrometheusRulesSpec {
	if x != nil {
		return x.PrometheusRules
	}
	return nil
}

func (x *DiscoverySpec) GetFilesystem() *FilesystemRulesSpec {
	if x != nil {
		return x.Filesystem
	}
	return nil
}

func (x *DiscoverySpec) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

type PrometheusRulesSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Namespaces to search for rules in. If empty, will search all accessible
	// namespaces.
	SearchNamespaces []string `protobuf:"bytes,1,rep,name=searchNamespaces,proto3" json:"searchNamespaces,omitempty"`
	// Kubeconfig to use for rule discovery. If nil, will use the in-cluster
	// kubeconfig.
	Kubeconfig *string `protobuf:"bytes,2,opt,name=kubeconfig,proto3,oneof" json:"kubeconfig,omitempty"`
}

func (x *PrometheusRulesSpec) Reset() {
	*x = PrometheusRulesSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrometheusRulesSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrometheusRulesSpec) ProtoMessage() {}

func (x *PrometheusRulesSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrometheusRulesSpec.ProtoReflect.Descriptor instead.
func (*PrometheusRulesSpec) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescGZIP(), []int{2}
}

func (x *PrometheusRulesSpec) GetSearchNamespaces() []string {
	if x != nil {
		return x.SearchNamespaces
	}
	return nil
}

func (x *PrometheusRulesSpec) GetKubeconfig() string {
	if x != nil && x.Kubeconfig != nil {
		return *x.Kubeconfig
	}
	return ""
}

type FilesystemRulesSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PathExpressions []string `protobuf:"bytes,1,rep,name=pathExpressions,proto3" json:"pathExpressions,omitempty"`
}

func (x *FilesystemRulesSpec) Reset() {
	*x = FilesystemRulesSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesystemRulesSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesystemRulesSpec) ProtoMessage() {}

func (x *FilesystemRulesSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesystemRulesSpec.ProtoReflect.Descriptor instead.
func (*FilesystemRulesSpec) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescGZIP(), []int{3}
}

func (x *FilesystemRulesSpec) GetPathExpressions() []string {
	if x != nil {
		return x.PathExpressions
	}
	return nil
}

var File_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22,
	0x48, 0x0a, 0x09, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3b, 0x0a, 0x09,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x22, 0xbf, 0x01, 0x0a, 0x0d, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4d, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x65, 0x75, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x0a, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x75, 0x0a, 0x13, 0x50,
	0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x23,
	0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x3f, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x74,
	0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x74, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescData = file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDesc
)

func file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescData)
	})
	return file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDescData
}

var file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_goTypes = []interface{}{
	(*RulesSpec)(nil),           // 0: config.v1beta1.RulesSpec
	(*DiscoverySpec)(nil),       // 1: config.v1beta1.DiscoverySpec
	(*PrometheusRulesSpec)(nil), // 2: config.v1beta1.PrometheusRulesSpec
	(*FilesystemRulesSpec)(nil), // 3: config.v1beta1.FilesystemRulesSpec
}
var file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_depIdxs = []int32{
	1, // 0: config.v1beta1.RulesSpec.discovery:type_name -> config.v1beta1.DiscoverySpec
	2, // 1: config.v1beta1.DiscoverySpec.prometheusRules:type_name -> config.v1beta1.PrometheusRulesSpec
	3, // 2: config.v1beta1.DiscoverySpec.filesystem:type_name -> config.v1beta1.FilesystemRulesSpec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_init() }
func file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_init() {
	if File_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RulesSpec); i {
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
		file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverySpec); i {
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
		file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrometheusRulesSpec); i {
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
		file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesystemRulesSpec); i {
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
	file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto = out.File
	file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_rawDesc = nil
	file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_goTypes = nil
	file_github_com_rancher_opni_pkg_config_v1beta1_agent_config_proto_depIdxs = nil
}