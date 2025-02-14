// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: profile/v1/profile.proto

package profilev1

import (
	key "github.com/MemeLabs/strims/pkg/apis/type/key"
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

type KDFType int32

const (
	KDFType_KDF_TYPE_UNDEFINED     KDFType = 0
	KDFType_KDF_TYPE_PBKDF2_SHA256 KDFType = 1
)

// Enum value maps for KDFType.
var (
	KDFType_name = map[int32]string{
		0: "KDF_TYPE_UNDEFINED",
		1: "KDF_TYPE_PBKDF2_SHA256",
	}
	KDFType_value = map[string]int32{
		"KDF_TYPE_UNDEFINED":     0,
		"KDF_TYPE_PBKDF2_SHA256": 1,
	}
)

func (x KDFType) Enum() *KDFType {
	p := new(KDFType)
	*p = x
	return p
}

func (x KDFType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KDFType) Descriptor() protoreflect.EnumDescriptor {
	return file_profile_v1_profile_proto_enumTypes[0].Descriptor()
}

func (KDFType) Type() protoreflect.EnumType {
	return &file_profile_v1_profile_proto_enumTypes[0]
}

func (x KDFType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KDFType.Descriptor instead.
func (KDFType) EnumDescriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{0}
}

type UpdateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UpdateProfileRequest) Reset() {
	*x = UpdateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileRequest) ProtoMessage() {}

func (x *UpdateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProfileRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *UpdateProfileResponse) Reset() {
	*x = UpdateProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileResponse) ProtoMessage() {}

func (x *UpdateProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileResponse.ProtoReflect.Descriptor instead.
func (*UpdateProfileResponse) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateProfileResponse) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type GetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{2}
}

type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *Profile `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{3}
}

func (x *GetProfileResponse) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type StorageKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KdfType KDFType `protobuf:"varint,1,opt,name=kdf_type,json=kdfType,proto3,enum=strims.profile.v1.KDFType" json:"kdf_type,omitempty"`
	// Types that are assignable to KdfOptions:
	//	*StorageKey_Pbkdf2Options
	KdfOptions isStorageKey_KdfOptions `protobuf_oneof:"kdf_options"`
}

func (x *StorageKey) Reset() {
	*x = StorageKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageKey) ProtoMessage() {}

func (x *StorageKey) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageKey.ProtoReflect.Descriptor instead.
func (*StorageKey) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{4}
}

func (x *StorageKey) GetKdfType() KDFType {
	if x != nil {
		return x.KdfType
	}
	return KDFType_KDF_TYPE_UNDEFINED
}

func (m *StorageKey) GetKdfOptions() isStorageKey_KdfOptions {
	if m != nil {
		return m.KdfOptions
	}
	return nil
}

func (x *StorageKey) GetPbkdf2Options() *StorageKey_PBKDF2Options {
	if x, ok := x.GetKdfOptions().(*StorageKey_Pbkdf2Options); ok {
		return x.Pbkdf2Options
	}
	return nil
}

type isStorageKey_KdfOptions interface {
	isStorageKey_KdfOptions()
}

type StorageKey_Pbkdf2Options struct {
	Pbkdf2Options *StorageKey_PBKDF2Options `protobuf:"bytes,2,opt,name=pbkdf2_options,json=pbkdf2Options,proto3,oneof"`
}

func (*StorageKey_Pbkdf2Options) isStorageKey_KdfOptions() {}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Secret []byte   `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	Key    *key.Key `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{5}
}

func (x *Profile) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *Profile) GetKey() *key.Key {
	if x != nil {
		return x.Key
	}
	return nil
}

type ProfileID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextId uint64 `protobuf:"varint,1,opt,name=next_id,json=nextId,proto3" json:"next_id,omitempty"`
}

func (x *ProfileID) Reset() {
	*x = ProfileID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileID) ProtoMessage() {}

func (x *ProfileID) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileID.ProtoReflect.Descriptor instead.
func (*ProfileID) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{6}
}

func (x *ProfileID) GetNextId() uint64 {
	if x != nil {
		return x.NextId
	}
	return 0
}

type StorageKey_PBKDF2Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iterations uint32 `protobuf:"varint,1,opt,name=iterations,proto3" json:"iterations,omitempty"`
	KeySize    uint32 `protobuf:"varint,2,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	Salt       []byte `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty"`
}

func (x *StorageKey_PBKDF2Options) Reset() {
	*x = StorageKey_PBKDF2Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_v1_profile_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageKey_PBKDF2Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageKey_PBKDF2Options) ProtoMessage() {}

func (x *StorageKey_PBKDF2Options) ProtoReflect() protoreflect.Message {
	mi := &file_profile_v1_profile_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageKey_PBKDF2Options.ProtoReflect.Descriptor instead.
func (*StorageKey_PBKDF2Options) Descriptor() ([]byte, []int) {
	return file_profile_v1_profile_proto_rawDescGZIP(), []int{4, 0}
}

func (x *StorageKey_PBKDF2Options) GetIterations() uint32 {
	if x != nil {
		return x.Iterations
	}
	return 0
}

func (x *StorageKey_PBKDF2Options) GetKeySize() uint32 {
	if x != nil {
		return x.KeySize
	}
	return 0
}

func (x *StorageKey_PBKDF2Options) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

var File_profile_v1_profile_proto protoreflect.FileDescriptor

var file_profile_v1_profile_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x74, 0x72, 0x69,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x0e, 0x74,
	0x79, 0x70, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4d, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x88, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x08, 0x6b, 0x64, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x44, 0x46, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x07, 0x6b, 0x64, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x54, 0x0a, 0x0e, 0x70,
	0x62, 0x6b, 0x64, 0x66, 0x32, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4b,
	0x65, 0x79, 0x2e, 0x50, 0x42, 0x4b, 0x44, 0x46, 0x32, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x48, 0x00, 0x52, 0x0d, 0x70, 0x62, 0x6b, 0x64, 0x66, 0x32, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x5e, 0x0a, 0x0d, 0x50, 0x42, 0x4b, 0x44, 0x46, 0x32, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x61, 0x6c,
	0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x6b, 0x64, 0x66, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x69, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x24, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x65, 0x78, 0x74, 0x49,
	0x64, 0x2a, 0x3d, 0x0a, 0x07, 0x4b, 0x44, 0x46, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x4b, 0x44, 0x46, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x4b, 0x44, 0x46, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x42, 0x4b, 0x44, 0x46, 0x32, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x01,
	0x32, 0xc2, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x12, 0x52, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x56, 0x0a, 0x14, 0x67, 0x67, 0x2e, 0x73, 0x74, 0x72, 0x69,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x6d, 0x65, 0x4c, 0x61,
	0x62, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x76, 0x31, 0xba, 0x02, 0x03, 0x53, 0x50, 0x46, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_v1_profile_proto_rawDescOnce sync.Once
	file_profile_v1_profile_proto_rawDescData = file_profile_v1_profile_proto_rawDesc
)

func file_profile_v1_profile_proto_rawDescGZIP() []byte {
	file_profile_v1_profile_proto_rawDescOnce.Do(func() {
		file_profile_v1_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_v1_profile_proto_rawDescData)
	})
	return file_profile_v1_profile_proto_rawDescData
}

var file_profile_v1_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_profile_v1_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_profile_v1_profile_proto_goTypes = []interface{}{
	(KDFType)(0),                     // 0: strims.profile.v1.KDFType
	(*UpdateProfileRequest)(nil),     // 1: strims.profile.v1.UpdateProfileRequest
	(*UpdateProfileResponse)(nil),    // 2: strims.profile.v1.UpdateProfileResponse
	(*GetProfileRequest)(nil),        // 3: strims.profile.v1.GetProfileRequest
	(*GetProfileResponse)(nil),       // 4: strims.profile.v1.GetProfileResponse
	(*StorageKey)(nil),               // 5: strims.profile.v1.StorageKey
	(*Profile)(nil),                  // 6: strims.profile.v1.Profile
	(*ProfileID)(nil),                // 7: strims.profile.v1.ProfileID
	(*StorageKey_PBKDF2Options)(nil), // 8: strims.profile.v1.StorageKey.PBKDF2Options
	(*key.Key)(nil),                  // 9: strims.type.Key
}
var file_profile_v1_profile_proto_depIdxs = []int32{
	6, // 0: strims.profile.v1.UpdateProfileResponse.profile:type_name -> strims.profile.v1.Profile
	6, // 1: strims.profile.v1.GetProfileResponse.profile:type_name -> strims.profile.v1.Profile
	0, // 2: strims.profile.v1.StorageKey.kdf_type:type_name -> strims.profile.v1.KDFType
	8, // 3: strims.profile.v1.StorageKey.pbkdf2_options:type_name -> strims.profile.v1.StorageKey.PBKDF2Options
	9, // 4: strims.profile.v1.Profile.key:type_name -> strims.type.Key
	3, // 5: strims.profile.v1.ProfileFrontend.Get:input_type -> strims.profile.v1.GetProfileRequest
	1, // 6: strims.profile.v1.ProfileFrontend.Update:input_type -> strims.profile.v1.UpdateProfileRequest
	4, // 7: strims.profile.v1.ProfileFrontend.Get:output_type -> strims.profile.v1.GetProfileResponse
	2, // 8: strims.profile.v1.ProfileFrontend.Update:output_type -> strims.profile.v1.UpdateProfileResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_profile_v1_profile_proto_init() }
func file_profile_v1_profile_proto_init() {
	if File_profile_v1_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_v1_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileRequest); i {
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
		file_profile_v1_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileResponse); i {
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
		file_profile_v1_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileRequest); i {
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
		file_profile_v1_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponse); i {
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
		file_profile_v1_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageKey); i {
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
		file_profile_v1_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_profile_v1_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileID); i {
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
		file_profile_v1_profile_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageKey_PBKDF2Options); i {
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
	file_profile_v1_profile_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*StorageKey_Pbkdf2Options)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_profile_v1_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_v1_profile_proto_goTypes,
		DependencyIndexes: file_profile_v1_profile_proto_depIdxs,
		EnumInfos:         file_profile_v1_profile_proto_enumTypes,
		MessageInfos:      file_profile_v1_profile_proto_msgTypes,
	}.Build()
	File_profile_v1_profile_proto = out.File
	file_profile_v1_profile_proto_rawDesc = nil
	file_profile_v1_profile_proto_goTypes = nil
	file_profile_v1_profile_proto_depIdxs = nil
}
