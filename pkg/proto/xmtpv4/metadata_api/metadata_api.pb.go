// Metadata API

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: xmtpv4/metadata_api/metadata_api.proto

package metadata_api

import (
	envelopes "github.com/xmtp/example-notification-server-go/pkg/proto/xmtpv4/envelopes"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetSyncCursorRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSyncCursorRequest) Reset() {
	*x = GetSyncCursorRequest{}
	mi := &file_xmtpv4_metadata_api_metadata_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSyncCursorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSyncCursorRequest) ProtoMessage() {}

func (x *GetSyncCursorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmtpv4_metadata_api_metadata_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSyncCursorRequest.ProtoReflect.Descriptor instead.
func (*GetSyncCursorRequest) Descriptor() ([]byte, []int) {
	return file_xmtpv4_metadata_api_metadata_api_proto_rawDescGZIP(), []int{0}
}

type GetSyncCursorResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LatestSync    *envelopes.Cursor      `protobuf:"bytes,1,opt,name=latest_sync,json=latestSync,proto3" json:"latest_sync,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSyncCursorResponse) Reset() {
	*x = GetSyncCursorResponse{}
	mi := &file_xmtpv4_metadata_api_metadata_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSyncCursorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSyncCursorResponse) ProtoMessage() {}

func (x *GetSyncCursorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xmtpv4_metadata_api_metadata_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSyncCursorResponse.ProtoReflect.Descriptor instead.
func (*GetSyncCursorResponse) Descriptor() ([]byte, []int) {
	return file_xmtpv4_metadata_api_metadata_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetSyncCursorResponse) GetLatestSync() *envelopes.Cursor {
	if x != nil {
		return x.LatestSync
	}
	return nil
}

var File_xmtpv4_metadata_api_metadata_api_proto protoreflect.FileDescriptor

var file_xmtpv4_metadata_api_metadata_api_proto_rawDesc = string([]byte{
	0x0a, 0x26, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x78,
	0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61,
	0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x57, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e,
	0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73,
	0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x32, 0xdb, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x41, 0x70, 0x69, 0x12, 0x9d, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x43,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x2e, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x78, 0x6d, 0x74,
	0x70, 0x76, 0x34, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x78, 0x6d, 0x74,
	0x70, 0x76, 0x34, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01,
	0x2a, 0x22, 0x20, 0x2f, 0x6d, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x63, 0x75, 0x72,
	0x73, 0x6f, 0x72, 0x12, 0xab, 0x01, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x2e, 0x2e, 0x78, 0x6d,
	0x74, 0x70, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x78, 0x6d,
	0x74, 0x70, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x6d, 0x6c, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x2d, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x30,
	0x01, 0x42, 0x96, 0x02, 0x0a, 0x36, 0x6f, 0x72, 0x67, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x61,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2e, 0x70,
	0x75, 0x73, 0x68, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x42, 0x10, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x74,
	0x70, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x76,
	0x34, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x70, 0x69, 0xa2, 0x02,
	0x03, 0x58, 0x58, 0x4d, 0xaa, 0x02, 0x17, 0x58, 0x6d, 0x74, 0x70, 0x2e, 0x58, 0x6d, 0x74, 0x70,
	0x76, 0x34, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x41, 0x70, 0x69, 0xca, 0x02,
	0x17, 0x58, 0x6d, 0x74, 0x70, 0x5c, 0x58, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x5c, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x41, 0x70, 0x69, 0xe2, 0x02, 0x23, 0x58, 0x6d, 0x74, 0x70, 0x5c,
	0x58, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x41,
	0x70, 0x69, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x19, 0x58, 0x6d, 0x74, 0x70, 0x3a, 0x3a, 0x58, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x3a, 0x3a, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_xmtpv4_metadata_api_metadata_api_proto_rawDescOnce sync.Once
	file_xmtpv4_metadata_api_metadata_api_proto_rawDescData []byte
)

func file_xmtpv4_metadata_api_metadata_api_proto_rawDescGZIP() []byte {
	file_xmtpv4_metadata_api_metadata_api_proto_rawDescOnce.Do(func() {
		file_xmtpv4_metadata_api_metadata_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_xmtpv4_metadata_api_metadata_api_proto_rawDesc), len(file_xmtpv4_metadata_api_metadata_api_proto_rawDesc)))
	})
	return file_xmtpv4_metadata_api_metadata_api_proto_rawDescData
}

var file_xmtpv4_metadata_api_metadata_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_xmtpv4_metadata_api_metadata_api_proto_goTypes = []any{
	(*GetSyncCursorRequest)(nil),  // 0: xmtp.xmtpv4.metadata_api.GetSyncCursorRequest
	(*GetSyncCursorResponse)(nil), // 1: xmtp.xmtpv4.metadata_api.GetSyncCursorResponse
	(*envelopes.Cursor)(nil),      // 2: xmtp.xmtpv4.envelopes.Cursor
}
var file_xmtpv4_metadata_api_metadata_api_proto_depIdxs = []int32{
	2, // 0: xmtp.xmtpv4.metadata_api.GetSyncCursorResponse.latest_sync:type_name -> xmtp.xmtpv4.envelopes.Cursor
	0, // 1: xmtp.xmtpv4.metadata_api.MetadataApi.GetSyncCursor:input_type -> xmtp.xmtpv4.metadata_api.GetSyncCursorRequest
	0, // 2: xmtp.xmtpv4.metadata_api.MetadataApi.SubscribeSyncCursor:input_type -> xmtp.xmtpv4.metadata_api.GetSyncCursorRequest
	1, // 3: xmtp.xmtpv4.metadata_api.MetadataApi.GetSyncCursor:output_type -> xmtp.xmtpv4.metadata_api.GetSyncCursorResponse
	1, // 4: xmtp.xmtpv4.metadata_api.MetadataApi.SubscribeSyncCursor:output_type -> xmtp.xmtpv4.metadata_api.GetSyncCursorResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_xmtpv4_metadata_api_metadata_api_proto_init() }
func file_xmtpv4_metadata_api_metadata_api_proto_init() {
	if File_xmtpv4_metadata_api_metadata_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_xmtpv4_metadata_api_metadata_api_proto_rawDesc), len(file_xmtpv4_metadata_api_metadata_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xmtpv4_metadata_api_metadata_api_proto_goTypes,
		DependencyIndexes: file_xmtpv4_metadata_api_metadata_api_proto_depIdxs,
		MessageInfos:      file_xmtpv4_metadata_api_metadata_api_proto_msgTypes,
	}.Build()
	File_xmtpv4_metadata_api_metadata_api_proto = out.File
	file_xmtpv4_metadata_api_metadata_api_proto_goTypes = nil
	file_xmtpv4_metadata_api_metadata_api_proto_depIdxs = nil
}
