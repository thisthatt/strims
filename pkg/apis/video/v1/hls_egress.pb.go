// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: video/v1/hls_egress.proto

package video

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

type HLSEgressConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled          bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	PublicServerAddr string `protobuf:"bytes,2,opt,name=public_server_addr,json=publicServerAddr,proto3" json:"public_server_addr,omitempty"`
}

func (x *HLSEgressConfig) Reset() {
	*x = HLSEgressConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressConfig) ProtoMessage() {}

func (x *HLSEgressConfig) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressConfig.ProtoReflect.Descriptor instead.
func (*HLSEgressConfig) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{0}
}

func (x *HLSEgressConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *HLSEgressConfig) GetPublicServerAddr() string {
	if x != nil {
		return x.PublicServerAddr
	}
	return ""
}

type HLSEgressIsSupportedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HLSEgressIsSupportedRequest) Reset() {
	*x = HLSEgressIsSupportedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressIsSupportedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressIsSupportedRequest) ProtoMessage() {}

func (x *HLSEgressIsSupportedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressIsSupportedRequest.ProtoReflect.Descriptor instead.
func (*HLSEgressIsSupportedRequest) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{1}
}

type HLSEgressIsSupportedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Supported bool `protobuf:"varint,1,opt,name=supported,proto3" json:"supported,omitempty"`
}

func (x *HLSEgressIsSupportedResponse) Reset() {
	*x = HLSEgressIsSupportedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressIsSupportedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressIsSupportedResponse) ProtoMessage() {}

func (x *HLSEgressIsSupportedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressIsSupportedResponse.ProtoReflect.Descriptor instead.
func (*HLSEgressIsSupportedResponse) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{2}
}

func (x *HLSEgressIsSupportedResponse) GetSupported() bool {
	if x != nil {
		return x.Supported
	}
	return false
}

type HLSEgressGetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HLSEgressGetConfigRequest) Reset() {
	*x = HLSEgressGetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressGetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressGetConfigRequest) ProtoMessage() {}

func (x *HLSEgressGetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressGetConfigRequest.ProtoReflect.Descriptor instead.
func (*HLSEgressGetConfigRequest) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{3}
}

type HLSEgressGetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *HLSEgressConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *HLSEgressGetConfigResponse) Reset() {
	*x = HLSEgressGetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressGetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressGetConfigResponse) ProtoMessage() {}

func (x *HLSEgressGetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressGetConfigResponse.ProtoReflect.Descriptor instead.
func (*HLSEgressGetConfigResponse) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{4}
}

func (x *HLSEgressGetConfigResponse) GetConfig() *HLSEgressConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type HLSEgressSetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *HLSEgressConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *HLSEgressSetConfigRequest) Reset() {
	*x = HLSEgressSetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressSetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressSetConfigRequest) ProtoMessage() {}

func (x *HLSEgressSetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressSetConfigRequest.ProtoReflect.Descriptor instead.
func (*HLSEgressSetConfigRequest) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{5}
}

func (x *HLSEgressSetConfigRequest) GetConfig() *HLSEgressConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type HLSEgressSetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *HLSEgressConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *HLSEgressSetConfigResponse) Reset() {
	*x = HLSEgressSetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressSetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressSetConfigResponse) ProtoMessage() {}

func (x *HLSEgressSetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressSetConfigResponse.ProtoReflect.Descriptor instead.
func (*HLSEgressSetConfigResponse) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{6}
}

func (x *HLSEgressSetConfigResponse) GetConfig() *HLSEgressConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type HLSEgressOpenStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SwarmUri    string   `protobuf:"bytes,1,opt,name=swarm_uri,json=swarmUri,proto3" json:"swarm_uri,omitempty"`
	NetworkKeys [][]byte `protobuf:"bytes,2,rep,name=network_keys,json=networkKeys,proto3" json:"network_keys,omitempty"`
}

func (x *HLSEgressOpenStreamRequest) Reset() {
	*x = HLSEgressOpenStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressOpenStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressOpenStreamRequest) ProtoMessage() {}

func (x *HLSEgressOpenStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressOpenStreamRequest.ProtoReflect.Descriptor instead.
func (*HLSEgressOpenStreamRequest) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{7}
}

func (x *HLSEgressOpenStreamRequest) GetSwarmUri() string {
	if x != nil {
		return x.SwarmUri
	}
	return ""
}

func (x *HLSEgressOpenStreamRequest) GetNetworkKeys() [][]byte {
	if x != nil {
		return x.NetworkKeys
	}
	return nil
}

type HLSEgressOpenStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaylistUrl string `protobuf:"bytes,1,opt,name=playlist_url,json=playlistUrl,proto3" json:"playlist_url,omitempty"`
}

func (x *HLSEgressOpenStreamResponse) Reset() {
	*x = HLSEgressOpenStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressOpenStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressOpenStreamResponse) ProtoMessage() {}

func (x *HLSEgressOpenStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressOpenStreamResponse.ProtoReflect.Descriptor instead.
func (*HLSEgressOpenStreamResponse) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{8}
}

func (x *HLSEgressOpenStreamResponse) GetPlaylistUrl() string {
	if x != nil {
		return x.PlaylistUrl
	}
	return ""
}

type HLSEgressCloseStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferId []byte `protobuf:"bytes,1,opt,name=transfer_id,json=transferId,proto3" json:"transfer_id,omitempty"`
}

func (x *HLSEgressCloseStreamRequest) Reset() {
	*x = HLSEgressCloseStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressCloseStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressCloseStreamRequest) ProtoMessage() {}

func (x *HLSEgressCloseStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressCloseStreamRequest.ProtoReflect.Descriptor instead.
func (*HLSEgressCloseStreamRequest) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{9}
}

func (x *HLSEgressCloseStreamRequest) GetTransferId() []byte {
	if x != nil {
		return x.TransferId
	}
	return nil
}

type HLSEgressCloseStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HLSEgressCloseStreamResponse) Reset() {
	*x = HLSEgressCloseStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_v1_hls_egress_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HLSEgressCloseStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HLSEgressCloseStreamResponse) ProtoMessage() {}

func (x *HLSEgressCloseStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_v1_hls_egress_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HLSEgressCloseStreamResponse.ProtoReflect.Descriptor instead.
func (*HLSEgressCloseStreamResponse) Descriptor() ([]byte, []int) {
	return file_video_v1_hls_egress_proto_rawDescGZIP(), []int{10}
}

var File_video_v1_hls_egress_proto protoreflect.FileDescriptor

var file_video_v1_hls_egress_proto_rawDesc = []byte{
	0x0a, 0x19, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6c, 0x73, 0x5f, 0x65,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x74, 0x72,
	0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x59, 0x0a, 0x0f,
	0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x22, 0x1d, 0x0a, 0x1b, 0x48, 0x4c, 0x53, 0x45, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x1c, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x56, 0x0a, 0x1a, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x55, 0x0a, 0x19, 0x48, 0x4c, 0x53,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x56, 0x0a, 0x1a, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x5c, 0x0a, 0x1a, 0x48, 0x4c, 0x53, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x77, 0x61, 0x72, 0x6d,
	0x55, 0x72, 0x69, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x40, 0x0a, 0x1b, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x3e, 0x0a, 0x1b, 0x48, 0x4c, 0x53, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x48, 0x4c, 0x53, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x98, 0x04, 0x0a, 0x09, 0x48, 0x4c, 0x53,
	0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x6a, 0x0a, 0x0b, 0x49, 0x73, 0x53, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x49, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x64, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x2a, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x4c,
	0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67,
	0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2b, 0x2e, 0x73,
	0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x74, 0x72, 0x69,
	0x6d, 0x73, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x4c, 0x53, 0x45,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2c, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x4c, 0x53, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x4e, 0x0a, 0x12, 0x67, 0x67, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x6d, 0x65, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x73,
	0x74, 0x72, 0x69, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0xba, 0x02, 0x03,
	0x53, 0x56, 0x4f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_v1_hls_egress_proto_rawDescOnce sync.Once
	file_video_v1_hls_egress_proto_rawDescData = file_video_v1_hls_egress_proto_rawDesc
)

func file_video_v1_hls_egress_proto_rawDescGZIP() []byte {
	file_video_v1_hls_egress_proto_rawDescOnce.Do(func() {
		file_video_v1_hls_egress_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_v1_hls_egress_proto_rawDescData)
	})
	return file_video_v1_hls_egress_proto_rawDescData
}

var file_video_v1_hls_egress_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_video_v1_hls_egress_proto_goTypes = []interface{}{
	(*HLSEgressConfig)(nil),              // 0: strims.video.v1.HLSEgressConfig
	(*HLSEgressIsSupportedRequest)(nil),  // 1: strims.video.v1.HLSEgressIsSupportedRequest
	(*HLSEgressIsSupportedResponse)(nil), // 2: strims.video.v1.HLSEgressIsSupportedResponse
	(*HLSEgressGetConfigRequest)(nil),    // 3: strims.video.v1.HLSEgressGetConfigRequest
	(*HLSEgressGetConfigResponse)(nil),   // 4: strims.video.v1.HLSEgressGetConfigResponse
	(*HLSEgressSetConfigRequest)(nil),    // 5: strims.video.v1.HLSEgressSetConfigRequest
	(*HLSEgressSetConfigResponse)(nil),   // 6: strims.video.v1.HLSEgressSetConfigResponse
	(*HLSEgressOpenStreamRequest)(nil),   // 7: strims.video.v1.HLSEgressOpenStreamRequest
	(*HLSEgressOpenStreamResponse)(nil),  // 8: strims.video.v1.HLSEgressOpenStreamResponse
	(*HLSEgressCloseStreamRequest)(nil),  // 9: strims.video.v1.HLSEgressCloseStreamRequest
	(*HLSEgressCloseStreamResponse)(nil), // 10: strims.video.v1.HLSEgressCloseStreamResponse
}
var file_video_v1_hls_egress_proto_depIdxs = []int32{
	0,  // 0: strims.video.v1.HLSEgressGetConfigResponse.config:type_name -> strims.video.v1.HLSEgressConfig
	0,  // 1: strims.video.v1.HLSEgressSetConfigRequest.config:type_name -> strims.video.v1.HLSEgressConfig
	0,  // 2: strims.video.v1.HLSEgressSetConfigResponse.config:type_name -> strims.video.v1.HLSEgressConfig
	1,  // 3: strims.video.v1.HLSEgress.IsSupported:input_type -> strims.video.v1.HLSEgressIsSupportedRequest
	3,  // 4: strims.video.v1.HLSEgress.GetConfig:input_type -> strims.video.v1.HLSEgressGetConfigRequest
	5,  // 5: strims.video.v1.HLSEgress.SetConfig:input_type -> strims.video.v1.HLSEgressSetConfigRequest
	7,  // 6: strims.video.v1.HLSEgress.OpenStream:input_type -> strims.video.v1.HLSEgressOpenStreamRequest
	9,  // 7: strims.video.v1.HLSEgress.CloseStream:input_type -> strims.video.v1.HLSEgressCloseStreamRequest
	2,  // 8: strims.video.v1.HLSEgress.IsSupported:output_type -> strims.video.v1.HLSEgressIsSupportedResponse
	4,  // 9: strims.video.v1.HLSEgress.GetConfig:output_type -> strims.video.v1.HLSEgressGetConfigResponse
	6,  // 10: strims.video.v1.HLSEgress.SetConfig:output_type -> strims.video.v1.HLSEgressSetConfigResponse
	8,  // 11: strims.video.v1.HLSEgress.OpenStream:output_type -> strims.video.v1.HLSEgressOpenStreamResponse
	10, // 12: strims.video.v1.HLSEgress.CloseStream:output_type -> strims.video.v1.HLSEgressCloseStreamResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_video_v1_hls_egress_proto_init() }
func file_video_v1_hls_egress_proto_init() {
	if File_video_v1_hls_egress_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_v1_hls_egress_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressConfig); i {
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
		file_video_v1_hls_egress_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressIsSupportedRequest); i {
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
		file_video_v1_hls_egress_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressIsSupportedResponse); i {
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
		file_video_v1_hls_egress_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressGetConfigRequest); i {
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
		file_video_v1_hls_egress_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressGetConfigResponse); i {
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
		file_video_v1_hls_egress_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressSetConfigRequest); i {
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
		file_video_v1_hls_egress_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressSetConfigResponse); i {
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
		file_video_v1_hls_egress_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressOpenStreamRequest); i {
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
		file_video_v1_hls_egress_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressOpenStreamResponse); i {
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
		file_video_v1_hls_egress_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressCloseStreamRequest); i {
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
		file_video_v1_hls_egress_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HLSEgressCloseStreamResponse); i {
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
			RawDescriptor: file_video_v1_hls_egress_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_v1_hls_egress_proto_goTypes,
		DependencyIndexes: file_video_v1_hls_egress_proto_depIdxs,
		MessageInfos:      file_video_v1_hls_egress_proto_msgTypes,
	}.Build()
	File_video_v1_hls_egress_proto = out.File
	file_video_v1_hls_egress_proto_rawDesc = nil
	file_video_v1_hls_egress_proto_goTypes = nil
	file_video_v1_hls_egress_proto_depIdxs = nil
}
