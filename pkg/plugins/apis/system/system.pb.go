// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v1.0.0
// source: github.com/rancher/opni/pkg/plugins/apis/system/system.proto

package system

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BrokerID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BrokerID) Reset() {
	*x = BrokerID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerID) ProtoMessage() {}

func (x *BrokerID) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerID.ProtoReflect.Descriptor instead.
func (*BrokerID) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescGZIP(), []int{0}
}

func (x *BrokerID) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescGZIP(), []int{1}
}

func (x *Key) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescGZIP(), []int{2}
}

func (x *Value) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescGZIP(), []int{3}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type KeyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *KeyList) Reset() {
	*x = KeyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyList) ProtoMessage() {}

func (x *KeyList) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyList.ProtoReflect.Descriptor instead.
func (*KeyList) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescGZIP(), []int{4}
}

func (x *KeyList) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

type DialAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DialAddress) Reset() {
	*x = DialAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DialAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DialAddress) ProtoMessage() {}

func (x *DialAddress) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DialAddress.ProtoReflect.Descriptor instead.
func (*DialAddress) Descriptor() ([]byte, []int) {
	return file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescGZIP(), []int{5}
}

func (x *DialAddress) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_github_com_rancher_opni_pkg_plugins_apis_system_system_proto protoreflect.FileDescriptor

var file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x17, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x1d, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1f, 0x0a, 0x07, 0x4b,
	0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x23, 0x0a, 0x0b,
	0x44, 0x69, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x32, 0xcd, 0x02, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x3c, 0x0a, 0x10,
	0x55, 0x73, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x50, 0x49,
	0x12, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x14, 0x55, 0x73,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x10,
	0x55, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x10, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x10, 0x55, 0x73,
	0x65, 0x41, 0x50, 0x49, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x13,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x12, 0x55,
	0x73, 0x65, 0x43, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x32, 0xbc, 0x01, 0x0a, 0x0d, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x0d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x0b, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65,
	0x79, 0x73, 0x12, 0x0b, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x1a,
	0x0f, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescOnce sync.Once
	file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescData = file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDesc
)

func file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescGZIP() []byte {
	file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescOnce.Do(func() {
		file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescData)
	})
	return file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDescData
}

var file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_goTypes = []interface{}{
	(*BrokerID)(nil),      // 0: system.BrokerID
	(*Key)(nil),           // 1: system.Key
	(*Value)(nil),         // 2: system.Value
	(*KeyValue)(nil),      // 3: system.KeyValue
	(*KeyList)(nil),       // 4: system.KeyList
	(*DialAddress)(nil),   // 5: system.DialAddress
	(*emptypb.Empty)(nil), // 6: google.protobuf.Empty
}
var file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_depIdxs = []int32{
	0, // 0: system.System.UseManagementAPI:input_type -> system.BrokerID
	0, // 1: system.System.UseNodeManagerClient:input_type -> system.BrokerID
	0, // 2: system.System.UseKeyValueStore:input_type -> system.BrokerID
	5, // 3: system.System.UseAPIExtensions:input_type -> system.DialAddress
	6, // 4: system.System.UseCachingProvider:input_type -> google.protobuf.Empty
	3, // 5: system.KeyValueStore.Put:input_type -> system.KeyValue
	1, // 6: system.KeyValueStore.Get:input_type -> system.Key
	1, // 7: system.KeyValueStore.Delete:input_type -> system.Key
	1, // 8: system.KeyValueStore.ListKeys:input_type -> system.Key
	6, // 9: system.System.UseManagementAPI:output_type -> google.protobuf.Empty
	6, // 10: system.System.UseNodeManagerClient:output_type -> google.protobuf.Empty
	6, // 11: system.System.UseKeyValueStore:output_type -> google.protobuf.Empty
	6, // 12: system.System.UseAPIExtensions:output_type -> google.protobuf.Empty
	6, // 13: system.System.UseCachingProvider:output_type -> google.protobuf.Empty
	6, // 14: system.KeyValueStore.Put:output_type -> google.protobuf.Empty
	2, // 15: system.KeyValueStore.Get:output_type -> system.Value
	6, // 16: system.KeyValueStore.Delete:output_type -> google.protobuf.Empty
	4, // 17: system.KeyValueStore.ListKeys:output_type -> system.KeyList
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_init() }
func file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_init() {
	if File_github_com_rancher_opni_pkg_plugins_apis_system_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerID); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyList); i {
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
		file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DialAddress); i {
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
			RawDescriptor: file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_goTypes,
		DependencyIndexes: file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_depIdxs,
		MessageInfos:      file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_msgTypes,
	}.Build()
	File_github_com_rancher_opni_pkg_plugins_apis_system_system_proto = out.File
	file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_rawDesc = nil
	file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_goTypes = nil
	file_github_com_rancher_opni_pkg_plugins_apis_system_system_proto_depIdxs = nil
}
