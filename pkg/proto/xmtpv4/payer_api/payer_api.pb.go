// Payer API

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: xmtpv4/payer_api/payer_api.proto

package payer_api

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

type PublishClientEnvelopesRequest struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Envelopes     []*envelopes.ClientEnvelope `protobuf:"bytes,1,rep,name=envelopes,proto3" json:"envelopes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishClientEnvelopesRequest) Reset() {
	*x = PublishClientEnvelopesRequest{}
	mi := &file_xmtpv4_payer_api_payer_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishClientEnvelopesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishClientEnvelopesRequest) ProtoMessage() {}

func (x *PublishClientEnvelopesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_xmtpv4_payer_api_payer_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishClientEnvelopesRequest.ProtoReflect.Descriptor instead.
func (*PublishClientEnvelopesRequest) Descriptor() ([]byte, []int) {
	return file_xmtpv4_payer_api_payer_api_proto_rawDescGZIP(), []int{0}
}

func (x *PublishClientEnvelopesRequest) GetEnvelopes() []*envelopes.ClientEnvelope {
	if x != nil {
		return x.Envelopes
	}
	return nil
}

type PublishClientEnvelopesResponse struct {
	state               protoimpl.MessageState          `protogen:"open.v1"`
	OriginatorEnvelopes []*envelopes.OriginatorEnvelope `protobuf:"bytes,1,rep,name=originator_envelopes,json=originatorEnvelopes,proto3" json:"originator_envelopes,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *PublishClientEnvelopesResponse) Reset() {
	*x = PublishClientEnvelopesResponse{}
	mi := &file_xmtpv4_payer_api_payer_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishClientEnvelopesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishClientEnvelopesResponse) ProtoMessage() {}

func (x *PublishClientEnvelopesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_xmtpv4_payer_api_payer_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishClientEnvelopesResponse.ProtoReflect.Descriptor instead.
func (*PublishClientEnvelopesResponse) Descriptor() ([]byte, []int) {
	return file_xmtpv4_payer_api_payer_api_proto_rawDescGZIP(), []int{1}
}

func (x *PublishClientEnvelopesResponse) GetOriginatorEnvelopes() []*envelopes.OriginatorEnvelope {
	if x != nil {
		return x.OriginatorEnvelopes
	}
	return nil
}

var File_xmtpv4_payer_api_payer_api_proto protoreflect.FileDescriptor

var file_xmtpv4_payer_api_payer_api_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2f, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e,
	0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2f,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x1d, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x09, 0x65, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x78, 0x6d, 0x74, 0x70, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e, 0x65, 0x6e, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x22,
	0x7e, 0x0a, 0x1e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5c, 0x0a, 0x14, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e, 0x65, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x6f, 0x72, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x13, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x32,
	0xc5, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x69, 0x12, 0xb8, 0x01, 0x0a,
	0x16, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x34, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x78,
	0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e,
	0x78, 0x6d, 0x74, 0x70, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e, 0x70, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22,
	0x26, 0x2f, 0x6d, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x65, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x42, 0x81, 0x02, 0x0a, 0x33, 0x6f, 0x72, 0x67, 0x2e,
	0x78, 0x6d, 0x74, 0x70, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x78,
	0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x42,
	0x0d, 0x50, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x74,
	0x70, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x76,
	0x34, 0x2f, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x03, 0x58, 0x58,
	0x50, 0xaa, 0x02, 0x14, 0x58, 0x6d, 0x74, 0x70, 0x2e, 0x58, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x2e,
	0x50, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x69, 0xca, 0x02, 0x14, 0x58, 0x6d, 0x74, 0x70, 0x5c,
	0x58, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x5c, 0x50, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x69, 0xe2,
	0x02, 0x20, 0x58, 0x6d, 0x74, 0x70, 0x5c, 0x58, 0x6d, 0x74, 0x70, 0x76, 0x34, 0x5c, 0x50, 0x61,
	0x79, 0x65, 0x72, 0x41, 0x70, 0x69, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x16, 0x58, 0x6d, 0x74, 0x70, 0x3a, 0x3a, 0x58, 0x6d, 0x74, 0x70, 0x76,
	0x34, 0x3a, 0x3a, 0x50, 0x61, 0x79, 0x65, 0x72, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_xmtpv4_payer_api_payer_api_proto_rawDescOnce sync.Once
	file_xmtpv4_payer_api_payer_api_proto_rawDescData []byte
)

func file_xmtpv4_payer_api_payer_api_proto_rawDescGZIP() []byte {
	file_xmtpv4_payer_api_payer_api_proto_rawDescOnce.Do(func() {
		file_xmtpv4_payer_api_payer_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_xmtpv4_payer_api_payer_api_proto_rawDesc), len(file_xmtpv4_payer_api_payer_api_proto_rawDesc)))
	})
	return file_xmtpv4_payer_api_payer_api_proto_rawDescData
}

var file_xmtpv4_payer_api_payer_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_xmtpv4_payer_api_payer_api_proto_goTypes = []any{
	(*PublishClientEnvelopesRequest)(nil),  // 0: xmtp.xmtpv4.payer_api.PublishClientEnvelopesRequest
	(*PublishClientEnvelopesResponse)(nil), // 1: xmtp.xmtpv4.payer_api.PublishClientEnvelopesResponse
	(*envelopes.ClientEnvelope)(nil),       // 2: xmtp.xmtpv4.envelopes.ClientEnvelope
	(*envelopes.OriginatorEnvelope)(nil),   // 3: xmtp.xmtpv4.envelopes.OriginatorEnvelope
}
var file_xmtpv4_payer_api_payer_api_proto_depIdxs = []int32{
	2, // 0: xmtp.xmtpv4.payer_api.PublishClientEnvelopesRequest.envelopes:type_name -> xmtp.xmtpv4.envelopes.ClientEnvelope
	3, // 1: xmtp.xmtpv4.payer_api.PublishClientEnvelopesResponse.originator_envelopes:type_name -> xmtp.xmtpv4.envelopes.OriginatorEnvelope
	0, // 2: xmtp.xmtpv4.payer_api.PayerApi.PublishClientEnvelopes:input_type -> xmtp.xmtpv4.payer_api.PublishClientEnvelopesRequest
	1, // 3: xmtp.xmtpv4.payer_api.PayerApi.PublishClientEnvelopes:output_type -> xmtp.xmtpv4.payer_api.PublishClientEnvelopesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_xmtpv4_payer_api_payer_api_proto_init() }
func file_xmtpv4_payer_api_payer_api_proto_init() {
	if File_xmtpv4_payer_api_payer_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_xmtpv4_payer_api_payer_api_proto_rawDesc), len(file_xmtpv4_payer_api_payer_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xmtpv4_payer_api_payer_api_proto_goTypes,
		DependencyIndexes: file_xmtpv4_payer_api_payer_api_proto_depIdxs,
		MessageInfos:      file_xmtpv4_payer_api_payer_api_proto_msgTypes,
	}.Build()
	File_xmtpv4_payer_api_payer_api_proto = out.File
	file_xmtpv4_payer_api_payer_api_proto_goTypes = nil
	file_xmtpv4_payer_api_payer_api_proto_depIdxs = nil
}
