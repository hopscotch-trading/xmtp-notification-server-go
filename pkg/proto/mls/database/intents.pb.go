// V3 invite message structure

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: mls/database/intents.proto

package database

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

// The data required to publish a message
type SendMessageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Version:
	//
	//	*SendMessageData_V1_
	Version isSendMessageData_Version `protobuf_oneof:"version"`
}

func (x *SendMessageData) Reset() {
	*x = SendMessageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageData) ProtoMessage() {}

func (x *SendMessageData) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageData.ProtoReflect.Descriptor instead.
func (*SendMessageData) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{0}
}

func (m *SendMessageData) GetVersion() isSendMessageData_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (x *SendMessageData) GetV1() *SendMessageData_V1 {
	if x, ok := x.GetVersion().(*SendMessageData_V1_); ok {
		return x.V1
	}
	return nil
}

type isSendMessageData_Version interface {
	isSendMessageData_Version()
}

type SendMessageData_V1_ struct {
	V1 *SendMessageData_V1 `protobuf:"bytes,1,opt,name=v1,proto3,oneof"`
}

func (*SendMessageData_V1_) isSendMessageData_Version() {}

// Wrapper around a list af repeated EVM Account Addresses
type AccountAddresses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountAddresses []string `protobuf:"bytes,1,rep,name=account_addresses,json=accountAddresses,proto3" json:"account_addresses,omitempty"`
}

func (x *AccountAddresses) Reset() {
	*x = AccountAddresses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountAddresses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountAddresses) ProtoMessage() {}

func (x *AccountAddresses) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountAddresses.ProtoReflect.Descriptor instead.
func (*AccountAddresses) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{1}
}

func (x *AccountAddresses) GetAccountAddresses() []string {
	if x != nil {
		return x.AccountAddresses
	}
	return nil
}

// Wrapper around a list of repeated Installation IDs
type InstallationIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationIds [][]byte `protobuf:"bytes,1,rep,name=installation_ids,json=installationIds,proto3" json:"installation_ids,omitempty"`
}

func (x *InstallationIds) Reset() {
	*x = InstallationIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstallationIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallationIds) ProtoMessage() {}

func (x *InstallationIds) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallationIds.ProtoReflect.Descriptor instead.
func (*InstallationIds) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{2}
}

func (x *InstallationIds) GetInstallationIds() [][]byte {
	if x != nil {
		return x.InstallationIds
	}
	return nil
}

// One of an EVM account address or Installation ID
type AddressesOrInstallationIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AddressesOrInstallationIds:
	//
	//	*AddressesOrInstallationIds_AccountAddresses
	//	*AddressesOrInstallationIds_InstallationIds
	AddressesOrInstallationIds isAddressesOrInstallationIds_AddressesOrInstallationIds `protobuf_oneof:"addresses_or_installation_ids"`
}

func (x *AddressesOrInstallationIds) Reset() {
	*x = AddressesOrInstallationIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressesOrInstallationIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressesOrInstallationIds) ProtoMessage() {}

func (x *AddressesOrInstallationIds) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressesOrInstallationIds.ProtoReflect.Descriptor instead.
func (*AddressesOrInstallationIds) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{3}
}

func (m *AddressesOrInstallationIds) GetAddressesOrInstallationIds() isAddressesOrInstallationIds_AddressesOrInstallationIds {
	if m != nil {
		return m.AddressesOrInstallationIds
	}
	return nil
}

func (x *AddressesOrInstallationIds) GetAccountAddresses() *AccountAddresses {
	if x, ok := x.GetAddressesOrInstallationIds().(*AddressesOrInstallationIds_AccountAddresses); ok {
		return x.AccountAddresses
	}
	return nil
}

func (x *AddressesOrInstallationIds) GetInstallationIds() *InstallationIds {
	if x, ok := x.GetAddressesOrInstallationIds().(*AddressesOrInstallationIds_InstallationIds); ok {
		return x.InstallationIds
	}
	return nil
}

type isAddressesOrInstallationIds_AddressesOrInstallationIds interface {
	isAddressesOrInstallationIds_AddressesOrInstallationIds()
}

type AddressesOrInstallationIds_AccountAddresses struct {
	AccountAddresses *AccountAddresses `protobuf:"bytes,1,opt,name=account_addresses,json=accountAddresses,proto3,oneof"`
}

type AddressesOrInstallationIds_InstallationIds struct {
	InstallationIds *InstallationIds `protobuf:"bytes,2,opt,name=installation_ids,json=installationIds,proto3,oneof"`
}

func (*AddressesOrInstallationIds_AccountAddresses) isAddressesOrInstallationIds_AddressesOrInstallationIds() {
}

func (*AddressesOrInstallationIds_InstallationIds) isAddressesOrInstallationIds_AddressesOrInstallationIds() {
}

// The data required to add members to a group
type AddMembersData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Version:
	//
	//	*AddMembersData_V1_
	Version isAddMembersData_Version `protobuf_oneof:"version"`
}

func (x *AddMembersData) Reset() {
	*x = AddMembersData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMembersData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMembersData) ProtoMessage() {}

func (x *AddMembersData) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMembersData.ProtoReflect.Descriptor instead.
func (*AddMembersData) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{4}
}

func (m *AddMembersData) GetVersion() isAddMembersData_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (x *AddMembersData) GetV1() *AddMembersData_V1 {
	if x, ok := x.GetVersion().(*AddMembersData_V1_); ok {
		return x.V1
	}
	return nil
}

type isAddMembersData_Version interface {
	isAddMembersData_Version()
}

type AddMembersData_V1_ struct {
	V1 *AddMembersData_V1 `protobuf:"bytes,1,opt,name=v1,proto3,oneof"`
}

func (*AddMembersData_V1_) isAddMembersData_Version() {}

// The data required to remove members from a group
type RemoveMembersData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Version:
	//
	//	*RemoveMembersData_V1_
	Version isRemoveMembersData_Version `protobuf_oneof:"version"`
}

func (x *RemoveMembersData) Reset() {
	*x = RemoveMembersData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMembersData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMembersData) ProtoMessage() {}

func (x *RemoveMembersData) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMembersData.ProtoReflect.Descriptor instead.
func (*RemoveMembersData) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{5}
}

func (m *RemoveMembersData) GetVersion() isRemoveMembersData_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (x *RemoveMembersData) GetV1() *RemoveMembersData_V1 {
	if x, ok := x.GetVersion().(*RemoveMembersData_V1_); ok {
		return x.V1
	}
	return nil
}

type isRemoveMembersData_Version interface {
	isRemoveMembersData_Version()
}

type RemoveMembersData_V1_ struct {
	V1 *RemoveMembersData_V1 `protobuf:"bytes,1,opt,name=v1,proto3,oneof"`
}

func (*RemoveMembersData_V1_) isRemoveMembersData_Version() {}

// Generic data-type for all post-commit actions
type PostCommitAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*PostCommitAction_SendWelcomes_
	Kind isPostCommitAction_Kind `protobuf_oneof:"kind"`
}

func (x *PostCommitAction) Reset() {
	*x = PostCommitAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCommitAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCommitAction) ProtoMessage() {}

func (x *PostCommitAction) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCommitAction.ProtoReflect.Descriptor instead.
func (*PostCommitAction) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{6}
}

func (m *PostCommitAction) GetKind() isPostCommitAction_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *PostCommitAction) GetSendWelcomes() *PostCommitAction_SendWelcomes {
	if x, ok := x.GetKind().(*PostCommitAction_SendWelcomes_); ok {
		return x.SendWelcomes
	}
	return nil
}

type isPostCommitAction_Kind interface {
	isPostCommitAction_Kind()
}

type PostCommitAction_SendWelcomes_ struct {
	SendWelcomes *PostCommitAction_SendWelcomes `protobuf:"bytes,1,opt,name=send_welcomes,json=sendWelcomes,proto3,oneof"`
}

func (*PostCommitAction_SendWelcomes_) isPostCommitAction_Kind() {}

// V1 of SendMessagePublishData
type SendMessageData_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayloadBytes []byte `protobuf:"bytes,1,opt,name=payload_bytes,json=payloadBytes,proto3" json:"payload_bytes,omitempty"`
}

func (x *SendMessageData_V1) Reset() {
	*x = SendMessageData_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageData_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageData_V1) ProtoMessage() {}

func (x *SendMessageData_V1) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageData_V1.ProtoReflect.Descriptor instead.
func (*SendMessageData_V1) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SendMessageData_V1) GetPayloadBytes() []byte {
	if x != nil {
		return x.PayloadBytes
	}
	return nil
}

// V1 of AddMembersPublishData
type AddMembersData_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressesOrInstallationIds *AddressesOrInstallationIds `protobuf:"bytes,1,opt,name=addresses_or_installation_ids,json=addressesOrInstallationIds,proto3" json:"addresses_or_installation_ids,omitempty"`
}

func (x *AddMembersData_V1) Reset() {
	*x = AddMembersData_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMembersData_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMembersData_V1) ProtoMessage() {}

func (x *AddMembersData_V1) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMembersData_V1.ProtoReflect.Descriptor instead.
func (*AddMembersData_V1) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{4, 0}
}

func (x *AddMembersData_V1) GetAddressesOrInstallationIds() *AddressesOrInstallationIds {
	if x != nil {
		return x.AddressesOrInstallationIds
	}
	return nil
}

// V1 of RemoveMembersPublishData
type RemoveMembersData_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddressesOrInstallationIds *AddressesOrInstallationIds `protobuf:"bytes,1,opt,name=addresses_or_installation_ids,json=addressesOrInstallationIds,proto3" json:"addresses_or_installation_ids,omitempty"`
}

func (x *RemoveMembersData_V1) Reset() {
	*x = RemoveMembersData_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMembersData_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMembersData_V1) ProtoMessage() {}

func (x *RemoveMembersData_V1) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMembersData_V1.ProtoReflect.Descriptor instead.
func (*RemoveMembersData_V1) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{5, 0}
}

func (x *RemoveMembersData_V1) GetAddressesOrInstallationIds() *AddressesOrInstallationIds {
	if x != nil {
		return x.AddressesOrInstallationIds
	}
	return nil
}

// An installation
type PostCommitAction_Installation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationKey []byte `protobuf:"bytes,1,opt,name=installation_key,json=installationKey,proto3" json:"installation_key,omitempty"`
	HpkePublicKey   []byte `protobuf:"bytes,2,opt,name=hpke_public_key,json=hpkePublicKey,proto3" json:"hpke_public_key,omitempty"`
}

func (x *PostCommitAction_Installation) Reset() {
	*x = PostCommitAction_Installation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCommitAction_Installation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCommitAction_Installation) ProtoMessage() {}

func (x *PostCommitAction_Installation) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCommitAction_Installation.ProtoReflect.Descriptor instead.
func (*PostCommitAction_Installation) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{6, 0}
}

func (x *PostCommitAction_Installation) GetInstallationKey() []byte {
	if x != nil {
		return x.InstallationKey
	}
	return nil
}

func (x *PostCommitAction_Installation) GetHpkePublicKey() []byte {
	if x != nil {
		return x.HpkePublicKey
	}
	return nil
}

// SendWelcome message
type PostCommitAction_SendWelcomes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Installations  []*PostCommitAction_Installation `protobuf:"bytes,1,rep,name=installations,proto3" json:"installations,omitempty"`
	WelcomeMessage []byte                           `protobuf:"bytes,2,opt,name=welcome_message,json=welcomeMessage,proto3" json:"welcome_message,omitempty"`
}

func (x *PostCommitAction_SendWelcomes) Reset() {
	*x = PostCommitAction_SendWelcomes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mls_database_intents_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostCommitAction_SendWelcomes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostCommitAction_SendWelcomes) ProtoMessage() {}

func (x *PostCommitAction_SendWelcomes) ProtoReflect() protoreflect.Message {
	mi := &file_mls_database_intents_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostCommitAction_SendWelcomes.ProtoReflect.Descriptor instead.
func (*PostCommitAction_SendWelcomes) Descriptor() ([]byte, []int) {
	return file_mls_database_intents_proto_rawDescGZIP(), []int{6, 1}
}

func (x *PostCommitAction_SendWelcomes) GetInstallations() []*PostCommitAction_Installation {
	if x != nil {
		return x.Installations
	}
	return nil
}

func (x *PostCommitAction_SendWelcomes) GetWelcomeMessage() []byte {
	if x != nil {
		return x.WelcomeMessage
	}
	return nil
}

var File_mls_database_intents_proto protoreflect.FileDescriptor

var file_mls_database_intents_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6c, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x78, 0x6d,
	0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22,
	0x80, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x56, 0x31, 0x48, 0x00, 0x52, 0x02, 0x76, 0x31, 0x1a, 0x29, 0x0a, 0x02,
	0x56, 0x31, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x73, 0x22, 0xe2, 0x01, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x4f,
	0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73,
	0x12, 0x52, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x78, 0x6d,
	0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x48, 0x00, 0x52, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x73, 0x42, 0x1f, 0x0a, 0x1d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x5f, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x02, 0x76, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x31, 0x48, 0x00, 0x52, 0x02, 0x76,
	0x31, 0x1a, 0x76, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x70, 0x0a, 0x1d, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x5f, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x4f, 0x72, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x1a, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x4f, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0xd1, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x02, 0x76, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c,
	0x73, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x31, 0x48,
	0x00, 0x52, 0x02, 0x76, 0x31, 0x1a, 0x76, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x70, 0x0a, 0x1d, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x5f, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x4f, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x73, 0x52, 0x1a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x4f, 0x72, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x42, 0x09, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe8, 0x02, 0x0a, 0x10, 0x50, 0x6f, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a,
	0x0d, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x65,
	0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x57, 0x65,
	0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x1a, 0x61, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65,
	0x79, 0x12, 0x26, 0x0a, 0x0f, 0x68, 0x70, 0x6b, 0x65, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x68, 0x70, 0x6b, 0x65,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x1a, 0x8f, 0x01, 0x0a, 0x0c, 0x53, 0x65,
	0x6e, 0x64, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x0d, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x77, 0x65, 0x6c,
	0x63, 0x6f, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x42, 0xec, 0x01, 0x0a, 0x2f, 0x6f, 0x72, 0x67, 0x2e, 0x78, 0x6d, 0x74, 0x70,
	0x2e, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x6c, 0x73, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x42, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x6c, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0xa2, 0x02,
	0x03, 0x58, 0x4d, 0x44, 0xaa, 0x02, 0x11, 0x58, 0x6d, 0x74, 0x70, 0x2e, 0x4d, 0x6c, 0x73, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0xca, 0x02, 0x11, 0x58, 0x6d, 0x74, 0x70, 0x5c,
	0x4d, 0x6c, 0x73, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0xe2, 0x02, 0x1d, 0x58,
	0x6d, 0x74, 0x70, 0x5c, 0x4d, 0x6c, 0x73, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x58,
	0x6d, 0x74, 0x70, 0x3a, 0x3a, 0x4d, 0x6c, 0x73, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mls_database_intents_proto_rawDescOnce sync.Once
	file_mls_database_intents_proto_rawDescData = file_mls_database_intents_proto_rawDesc
)

func file_mls_database_intents_proto_rawDescGZIP() []byte {
	file_mls_database_intents_proto_rawDescOnce.Do(func() {
		file_mls_database_intents_proto_rawDescData = protoimpl.X.CompressGZIP(file_mls_database_intents_proto_rawDescData)
	})
	return file_mls_database_intents_proto_rawDescData
}

var file_mls_database_intents_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_mls_database_intents_proto_goTypes = []interface{}{
	(*SendMessageData)(nil),               // 0: xmtp.mls.database.SendMessageData
	(*AccountAddresses)(nil),              // 1: xmtp.mls.database.AccountAddresses
	(*InstallationIds)(nil),               // 2: xmtp.mls.database.InstallationIds
	(*AddressesOrInstallationIds)(nil),    // 3: xmtp.mls.database.AddressesOrInstallationIds
	(*AddMembersData)(nil),                // 4: xmtp.mls.database.AddMembersData
	(*RemoveMembersData)(nil),             // 5: xmtp.mls.database.RemoveMembersData
	(*PostCommitAction)(nil),              // 6: xmtp.mls.database.PostCommitAction
	(*SendMessageData_V1)(nil),            // 7: xmtp.mls.database.SendMessageData.V1
	(*AddMembersData_V1)(nil),             // 8: xmtp.mls.database.AddMembersData.V1
	(*RemoveMembersData_V1)(nil),          // 9: xmtp.mls.database.RemoveMembersData.V1
	(*PostCommitAction_Installation)(nil), // 10: xmtp.mls.database.PostCommitAction.Installation
	(*PostCommitAction_SendWelcomes)(nil), // 11: xmtp.mls.database.PostCommitAction.SendWelcomes
}
var file_mls_database_intents_proto_depIdxs = []int32{
	7,  // 0: xmtp.mls.database.SendMessageData.v1:type_name -> xmtp.mls.database.SendMessageData.V1
	1,  // 1: xmtp.mls.database.AddressesOrInstallationIds.account_addresses:type_name -> xmtp.mls.database.AccountAddresses
	2,  // 2: xmtp.mls.database.AddressesOrInstallationIds.installation_ids:type_name -> xmtp.mls.database.InstallationIds
	8,  // 3: xmtp.mls.database.AddMembersData.v1:type_name -> xmtp.mls.database.AddMembersData.V1
	9,  // 4: xmtp.mls.database.RemoveMembersData.v1:type_name -> xmtp.mls.database.RemoveMembersData.V1
	11, // 5: xmtp.mls.database.PostCommitAction.send_welcomes:type_name -> xmtp.mls.database.PostCommitAction.SendWelcomes
	3,  // 6: xmtp.mls.database.AddMembersData.V1.addresses_or_installation_ids:type_name -> xmtp.mls.database.AddressesOrInstallationIds
	3,  // 7: xmtp.mls.database.RemoveMembersData.V1.addresses_or_installation_ids:type_name -> xmtp.mls.database.AddressesOrInstallationIds
	10, // 8: xmtp.mls.database.PostCommitAction.SendWelcomes.installations:type_name -> xmtp.mls.database.PostCommitAction.Installation
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_mls_database_intents_proto_init() }
func file_mls_database_intents_proto_init() {
	if File_mls_database_intents_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mls_database_intents_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageData); i {
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
		file_mls_database_intents_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountAddresses); i {
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
		file_mls_database_intents_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstallationIds); i {
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
		file_mls_database_intents_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressesOrInstallationIds); i {
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
		file_mls_database_intents_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMembersData); i {
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
		file_mls_database_intents_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMembersData); i {
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
		file_mls_database_intents_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCommitAction); i {
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
		file_mls_database_intents_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageData_V1); i {
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
		file_mls_database_intents_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMembersData_V1); i {
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
		file_mls_database_intents_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMembersData_V1); i {
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
		file_mls_database_intents_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCommitAction_Installation); i {
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
		file_mls_database_intents_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostCommitAction_SendWelcomes); i {
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
	file_mls_database_intents_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SendMessageData_V1_)(nil),
	}
	file_mls_database_intents_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*AddressesOrInstallationIds_AccountAddresses)(nil),
		(*AddressesOrInstallationIds_InstallationIds)(nil),
	}
	file_mls_database_intents_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*AddMembersData_V1_)(nil),
	}
	file_mls_database_intents_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*RemoveMembersData_V1_)(nil),
	}
	file_mls_database_intents_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*PostCommitAction_SendWelcomes_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mls_database_intents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mls_database_intents_proto_goTypes,
		DependencyIndexes: file_mls_database_intents_proto_depIdxs,
		MessageInfos:      file_mls_database_intents_proto_msgTypes,
	}.Build()
	File_mls_database_intents_proto = out.File
	file_mls_database_intents_proto_rawDesc = nil
	file_mls_database_intents_proto_goTypes = nil
	file_mls_database_intents_proto_depIdxs = nil
}
