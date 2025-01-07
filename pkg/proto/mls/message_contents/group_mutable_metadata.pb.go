// Group mutable metadata

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: mls/message_contents/group_mutable_metadata.proto

package message_contents

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

// Message for group mutable metadata
type GroupMutableMetadataV1 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Map to store various metadata attributes (Group name, etc.)
	Attributes map[string]string `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AdminList  *Inboxes          `protobuf:"bytes,2,opt,name=admin_list,json=adminList,proto3" json:"admin_list,omitempty"`
	// Creator starts as only super_admin
	// Only super_admin can add/remove other super_admin
	SuperAdminList *Inboxes `protobuf:"bytes,3,opt,name=super_admin_list,json=superAdminList,proto3" json:"super_admin_list,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GroupMutableMetadataV1) Reset() {
	*x = GroupMutableMetadataV1{}
	mi := &file_mls_message_contents_group_mutable_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupMutableMetadataV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupMutableMetadataV1) ProtoMessage() {}

func (x *GroupMutableMetadataV1) ProtoReflect() protoreflect.Message {
	mi := &file_mls_message_contents_group_mutable_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupMutableMetadataV1.ProtoReflect.Descriptor instead.
func (*GroupMutableMetadataV1) Descriptor() ([]byte, []int) {
	return file_mls_message_contents_group_mutable_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *GroupMutableMetadataV1) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *GroupMutableMetadataV1) GetAdminList() *Inboxes {
	if x != nil {
		return x.AdminList
	}
	return nil
}

func (x *GroupMutableMetadataV1) GetSuperAdminList() *Inboxes {
	if x != nil {
		return x.SuperAdminList
	}
	return nil
}

// Wrapper around a list of repeated Inbox Ids
type Inboxes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InboxIds      []string               `protobuf:"bytes,1,rep,name=inbox_ids,json=inboxIds,proto3" json:"inbox_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Inboxes) Reset() {
	*x = Inboxes{}
	mi := &file_mls_message_contents_group_mutable_metadata_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Inboxes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inboxes) ProtoMessage() {}

func (x *Inboxes) ProtoReflect() protoreflect.Message {
	mi := &file_mls_message_contents_group_mutable_metadata_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inboxes.ProtoReflect.Descriptor instead.
func (*Inboxes) Descriptor() ([]byte, []int) {
	return file_mls_message_contents_group_mutable_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *Inboxes) GetInboxIds() []string {
	if x != nil {
		return x.InboxIds
	}
	return nil
}

var File_mls_message_contents_group_mutable_metadata_proto protoreflect.FileDescriptor

var file_mls_message_contents_group_mutable_metadata_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6d, 0x6c, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x75, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xcb,
	0x02, 0x0a, 0x16, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x31, 0x12, 0x61, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41, 0x2e,
	0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x31,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0a,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x62,
	0x6f, 0x78, 0x65, 0x73, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x4c, 0x0a, 0x10, 0x73, 0x75, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x78, 0x6d, 0x74, 0x70,
	0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x52, 0x0e, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x3d, 0x0a,
	0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a, 0x07,
	0x49, 0x6e, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x62, 0x6f, 0x78,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x62, 0x6f,
	0x78, 0x49, 0x64, 0x73, 0x42, 0xa5, 0x02, 0x0a, 0x37, 0x6f, 0x72, 0x67, 0x2e, 0x78, 0x6d, 0x74,
	0x70, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x42, 0x19, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6c, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xa2, 0x02, 0x03, 0x58,
	0x4d, 0x4d, 0xaa, 0x02, 0x18, 0x58, 0x6d, 0x74, 0x70, 0x2e, 0x4d, 0x6c, 0x73, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xca, 0x02, 0x18,
	0x58, 0x6d, 0x74, 0x70, 0x5c, 0x4d, 0x6c, 0x73, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xe2, 0x02, 0x24, 0x58, 0x6d, 0x74, 0x70, 0x5c,
	0x4d, 0x6c, 0x73, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1a, 0x58, 0x6d, 0x74, 0x70, 0x3a, 0x3a, 0x4d, 0x6c, 0x73, 0x3a, 0x3a, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mls_message_contents_group_mutable_metadata_proto_rawDescOnce sync.Once
	file_mls_message_contents_group_mutable_metadata_proto_rawDescData = file_mls_message_contents_group_mutable_metadata_proto_rawDesc
)

func file_mls_message_contents_group_mutable_metadata_proto_rawDescGZIP() []byte {
	file_mls_message_contents_group_mutable_metadata_proto_rawDescOnce.Do(func() {
		file_mls_message_contents_group_mutable_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_mls_message_contents_group_mutable_metadata_proto_rawDescData)
	})
	return file_mls_message_contents_group_mutable_metadata_proto_rawDescData
}

var file_mls_message_contents_group_mutable_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mls_message_contents_group_mutable_metadata_proto_goTypes = []any{
	(*GroupMutableMetadataV1)(nil), // 0: xmtp.mls.message_contents.GroupMutableMetadataV1
	(*Inboxes)(nil),                // 1: xmtp.mls.message_contents.Inboxes
	nil,                            // 2: xmtp.mls.message_contents.GroupMutableMetadataV1.AttributesEntry
}
var file_mls_message_contents_group_mutable_metadata_proto_depIdxs = []int32{
	2, // 0: xmtp.mls.message_contents.GroupMutableMetadataV1.attributes:type_name -> xmtp.mls.message_contents.GroupMutableMetadataV1.AttributesEntry
	1, // 1: xmtp.mls.message_contents.GroupMutableMetadataV1.admin_list:type_name -> xmtp.mls.message_contents.Inboxes
	1, // 2: xmtp.mls.message_contents.GroupMutableMetadataV1.super_admin_list:type_name -> xmtp.mls.message_contents.Inboxes
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mls_message_contents_group_mutable_metadata_proto_init() }
func file_mls_message_contents_group_mutable_metadata_proto_init() {
	if File_mls_message_contents_group_mutable_metadata_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mls_message_contents_group_mutable_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mls_message_contents_group_mutable_metadata_proto_goTypes,
		DependencyIndexes: file_mls_message_contents_group_mutable_metadata_proto_depIdxs,
		MessageInfos:      file_mls_message_contents_group_mutable_metadata_proto_msgTypes,
	}.Build()
	File_mls_message_contents_group_mutable_metadata_proto = out.File
	file_mls_message_contents_group_mutable_metadata_proto_rawDesc = nil
	file_mls_message_contents_group_mutable_metadata_proto_goTypes = nil
	file_mls_message_contents_group_mutable_metadata_proto_depIdxs = nil
}
