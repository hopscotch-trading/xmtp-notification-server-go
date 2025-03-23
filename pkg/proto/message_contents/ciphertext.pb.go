// Ciphertext is a generic structure for encrypted payload.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: message_contents/ciphertext.proto

package message_contents

import (
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

// Ciphertext represents encrypted payload.
// It is definited as a union to support cryptographic algorithm agility.
// The payload is accompanied by the cryptographic parameters
// required by the chosen encryption scheme.
type Ciphertext struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Union:
	//
	//	*Ciphertext_Aes256GcmHkdfSha256
	Union         isCiphertext_Union `protobuf_oneof:"union"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ciphertext) Reset() {
	*x = Ciphertext{}
	mi := &file_message_contents_ciphertext_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ciphertext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ciphertext) ProtoMessage() {}

func (x *Ciphertext) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_ciphertext_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ciphertext.ProtoReflect.Descriptor instead.
func (*Ciphertext) Descriptor() ([]byte, []int) {
	return file_message_contents_ciphertext_proto_rawDescGZIP(), []int{0}
}

func (x *Ciphertext) GetUnion() isCiphertext_Union {
	if x != nil {
		return x.Union
	}
	return nil
}

func (x *Ciphertext) GetAes256GcmHkdfSha256() *Ciphertext_Aes256GcmHkdfsha256 {
	if x != nil {
		if x, ok := x.Union.(*Ciphertext_Aes256GcmHkdfSha256); ok {
			return x.Aes256GcmHkdfSha256
		}
	}
	return nil
}

type isCiphertext_Union interface {
	isCiphertext_Union()
}

type Ciphertext_Aes256GcmHkdfSha256 struct {
	Aes256GcmHkdfSha256 *Ciphertext_Aes256GcmHkdfsha256 `protobuf:"bytes,1,opt,name=aes256_gcm_hkdf_sha256,json=aes256GcmHkdfSha256,proto3,oneof"`
}

func (*Ciphertext_Aes256GcmHkdfSha256) isCiphertext_Union() {}

// SignedEciesCiphertext represents an ECIES encrypted payload and a signature
type SignedEciesCiphertext struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// serialized Ecies message
	EciesBytes []byte `protobuf:"bytes,1,opt,name=ecies_bytes,json=eciesBytes,proto3" json:"ecies_bytes,omitempty"`
	// signature of sha256(ecies_bytes) signed with the IdentityKey
	Signature     *Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignedEciesCiphertext) Reset() {
	*x = SignedEciesCiphertext{}
	mi := &file_message_contents_ciphertext_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignedEciesCiphertext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedEciesCiphertext) ProtoMessage() {}

func (x *SignedEciesCiphertext) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_ciphertext_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedEciesCiphertext.ProtoReflect.Descriptor instead.
func (*SignedEciesCiphertext) Descriptor() ([]byte, []int) {
	return file_message_contents_ciphertext_proto_rawDescGZIP(), []int{1}
}

func (x *SignedEciesCiphertext) GetEciesBytes() []byte {
	if x != nil {
		return x.EciesBytes
	}
	return nil
}

func (x *SignedEciesCiphertext) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// Encryption: AES256-GCM
// Key derivation function: HKDF-SHA256
type Ciphertext_Aes256GcmHkdfsha256 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HkdfSalt      []byte                 `protobuf:"bytes,1,opt,name=hkdf_salt,json=hkdfSalt,proto3" json:"hkdf_salt,omitempty"` // 32 bytes
	GcmNonce      []byte                 `protobuf:"bytes,2,opt,name=gcm_nonce,json=gcmNonce,proto3" json:"gcm_nonce,omitempty"` // 12 bytes
	Payload       []byte                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`                   // encrypted payload
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ciphertext_Aes256GcmHkdfsha256) Reset() {
	*x = Ciphertext_Aes256GcmHkdfsha256{}
	mi := &file_message_contents_ciphertext_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ciphertext_Aes256GcmHkdfsha256) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ciphertext_Aes256GcmHkdfsha256) ProtoMessage() {}

func (x *Ciphertext_Aes256GcmHkdfsha256) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_ciphertext_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ciphertext_Aes256GcmHkdfsha256.ProtoReflect.Descriptor instead.
func (*Ciphertext_Aes256GcmHkdfsha256) Descriptor() ([]byte, []int) {
	return file_message_contents_ciphertext_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Ciphertext_Aes256GcmHkdfsha256) GetHkdfSalt() []byte {
	if x != nil {
		return x.HkdfSalt
	}
	return nil
}

func (x *Ciphertext_Aes256GcmHkdfsha256) GetGcmNonce() []byte {
	if x != nil {
		return x.GcmNonce
	}
	return nil
}

func (x *Ciphertext_Aes256GcmHkdfsha256) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Ecies is ciphertext encrypted using ECIES with a MAC
type SignedEciesCiphertext_Ecies struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	EphemeralPublicKey []byte                 `protobuf:"bytes,1,opt,name=ephemeral_public_key,json=ephemeralPublicKey,proto3" json:"ephemeral_public_key,omitempty"` // 65 bytes
	Iv                 []byte                 `protobuf:"bytes,2,opt,name=iv,proto3" json:"iv,omitempty"`                                                             // 16 bytes
	Mac                []byte                 `protobuf:"bytes,3,opt,name=mac,proto3" json:"mac,omitempty"`                                                           // 32 bytes
	Ciphertext         []byte                 `protobuf:"bytes,4,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`                                             // encrypted payload with block size of 16
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *SignedEciesCiphertext_Ecies) Reset() {
	*x = SignedEciesCiphertext_Ecies{}
	mi := &file_message_contents_ciphertext_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignedEciesCiphertext_Ecies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedEciesCiphertext_Ecies) ProtoMessage() {}

func (x *SignedEciesCiphertext_Ecies) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_ciphertext_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedEciesCiphertext_Ecies.ProtoReflect.Descriptor instead.
func (*SignedEciesCiphertext_Ecies) Descriptor() ([]byte, []int) {
	return file_message_contents_ciphertext_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SignedEciesCiphertext_Ecies) GetEphemeralPublicKey() []byte {
	if x != nil {
		return x.EphemeralPublicKey
	}
	return nil
}

func (x *SignedEciesCiphertext_Ecies) GetIv() []byte {
	if x != nil {
		return x.Iv
	}
	return nil
}

func (x *SignedEciesCiphertext_Ecies) GetMac() []byte {
	if x != nil {
		return x.Mac
	}
	return nil
}

func (x *SignedEciesCiphertext_Ecies) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

var File_message_contents_ciphertext_proto protoreflect.FileDescriptor

var file_message_contents_ciphertext_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x20, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a,
	0x0a, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x6c, 0x0a, 0x16, 0x61,
	0x65, 0x73, 0x32, 0x35, 0x36, 0x5f, 0x67, 0x63, 0x6d, 0x5f, 0x68, 0x6b, 0x64, 0x66, 0x5f, 0x73,
	0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x78, 0x6d,
	0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x41,
	0x65, 0x73, 0x32, 0x35, 0x36, 0x67, 0x63, 0x6d, 0x48, 0x6b, 0x64, 0x66, 0x73, 0x68, 0x61, 0x32,
	0x35, 0x36, 0x48, 0x00, 0x52, 0x13, 0x61, 0x65, 0x73, 0x32, 0x35, 0x36, 0x47, 0x63, 0x6d, 0x48,
	0x6b, 0x64, 0x66, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x1a, 0x69, 0x0a, 0x13, 0x41, 0x65, 0x73,
	0x32, 0x35, 0x36, 0x67, 0x63, 0x6d, 0x48, 0x6b, 0x64, 0x66, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36,
	0x12, 0x1b, 0x0a, 0x09, 0x68, 0x6b, 0x64, 0x66, 0x5f, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x68, 0x6b, 0x64, 0x66, 0x53, 0x61, 0x6c, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x67, 0x63, 0x6d, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x67, 0x63, 0x6d, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x22, 0xf5, 0x01,
	0x0a, 0x15, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x45, 0x63, 0x69, 0x65, 0x73, 0x43, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x63, 0x69, 0x65, 0x73,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x65, 0x63,
	0x69, 0x65, 0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x6d,
	0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x7b, 0x0a, 0x05, 0x45, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x12, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x69, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x74, 0x65, 0x78, 0x74, 0x42, 0x82, 0x02, 0x0a, 0x33, 0x6f, 0x72, 0x67, 0x2e, 0x78, 0x6d,
	0x74, 0x70, 0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x0f, 0x43,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x74,
	0x70, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xa2, 0x02, 0x03, 0x58, 0x4d,
	0x58, 0xaa, 0x02, 0x14, 0x58, 0x6d, 0x74, 0x70, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xca, 0x02, 0x14, 0x58, 0x6d, 0x74, 0x70, 0x5c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0xe2,
	0x02, 0x20, 0x58, 0x6d, 0x74, 0x70, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x15, 0x58, 0x6d, 0x74, 0x70, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_message_contents_ciphertext_proto_rawDescOnce sync.Once
	file_message_contents_ciphertext_proto_rawDescData []byte
)

func file_message_contents_ciphertext_proto_rawDescGZIP() []byte {
	file_message_contents_ciphertext_proto_rawDescOnce.Do(func() {
		file_message_contents_ciphertext_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_message_contents_ciphertext_proto_rawDesc), len(file_message_contents_ciphertext_proto_rawDesc)))
	})
	return file_message_contents_ciphertext_proto_rawDescData
}

var file_message_contents_ciphertext_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message_contents_ciphertext_proto_goTypes = []any{
	(*Ciphertext)(nil),                     // 0: xmtp.message_contents.Ciphertext
	(*SignedEciesCiphertext)(nil),          // 1: xmtp.message_contents.SignedEciesCiphertext
	(*Ciphertext_Aes256GcmHkdfsha256)(nil), // 2: xmtp.message_contents.Ciphertext.Aes256gcmHkdfsha256
	(*SignedEciesCiphertext_Ecies)(nil),    // 3: xmtp.message_contents.SignedEciesCiphertext.Ecies
	(*Signature)(nil),                      // 4: xmtp.message_contents.Signature
}
var file_message_contents_ciphertext_proto_depIdxs = []int32{
	2, // 0: xmtp.message_contents.Ciphertext.aes256_gcm_hkdf_sha256:type_name -> xmtp.message_contents.Ciphertext.Aes256gcmHkdfsha256
	4, // 1: xmtp.message_contents.SignedEciesCiphertext.signature:type_name -> xmtp.message_contents.Signature
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_contents_ciphertext_proto_init() }
func file_message_contents_ciphertext_proto_init() {
	if File_message_contents_ciphertext_proto != nil {
		return
	}
	file_message_contents_signature_proto_init()
	file_message_contents_ciphertext_proto_msgTypes[0].OneofWrappers = []any{
		(*Ciphertext_Aes256GcmHkdfSha256)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_message_contents_ciphertext_proto_rawDesc), len(file_message_contents_ciphertext_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_contents_ciphertext_proto_goTypes,
		DependencyIndexes: file_message_contents_ciphertext_proto_depIdxs,
		MessageInfos:      file_message_contents_ciphertext_proto_msgTypes,
	}.Build()
	File_message_contents_ciphertext_proto = out.File
	file_message_contents_ciphertext_proto_goTypes = nil
	file_message_contents_ciphertext_proto_depIdxs = nil
}
