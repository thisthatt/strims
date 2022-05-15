// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: network/v1/bootstrap/peer.proto

package bootstrap

import (
	certificate "github.com/MemeLabs/strims/pkg/apis/type/certificate"
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

type BootstrapPeerGetPublishEnabledRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BootstrapPeerGetPublishEnabledRequest) Reset() {
	*x = BootstrapPeerGetPublishEnabledRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_bootstrap_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapPeerGetPublishEnabledRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapPeerGetPublishEnabledRequest) ProtoMessage() {}

func (x *BootstrapPeerGetPublishEnabledRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_bootstrap_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapPeerGetPublishEnabledRequest.ProtoReflect.Descriptor instead.
func (*BootstrapPeerGetPublishEnabledRequest) Descriptor() ([]byte, []int) {
	return file_network_v1_bootstrap_peer_proto_rawDescGZIP(), []int{0}
}

type BootstrapPeerGetPublishEnabledResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *BootstrapPeerGetPublishEnabledResponse) Reset() {
	*x = BootstrapPeerGetPublishEnabledResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_bootstrap_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapPeerGetPublishEnabledResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapPeerGetPublishEnabledResponse) ProtoMessage() {}

func (x *BootstrapPeerGetPublishEnabledResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_bootstrap_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapPeerGetPublishEnabledResponse.ProtoReflect.Descriptor instead.
func (*BootstrapPeerGetPublishEnabledResponse) Descriptor() ([]byte, []int) {
	return file_network_v1_bootstrap_peer_proto_rawDescGZIP(), []int{1}
}

func (x *BootstrapPeerGetPublishEnabledResponse) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type BootstrapPeerListNetworksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BootstrapPeerListNetworksRequest) Reset() {
	*x = BootstrapPeerListNetworksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_bootstrap_peer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapPeerListNetworksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapPeerListNetworksRequest) ProtoMessage() {}

func (x *BootstrapPeerListNetworksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_bootstrap_peer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapPeerListNetworksRequest.ProtoReflect.Descriptor instead.
func (*BootstrapPeerListNetworksRequest) Descriptor() ([]byte, []int) {
	return file_network_v1_bootstrap_peer_proto_rawDescGZIP(), []int{2}
}

type BootstrapPeerListNetworksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BootstrapPeerListNetworksResponse) Reset() {
	*x = BootstrapPeerListNetworksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_bootstrap_peer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapPeerListNetworksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapPeerListNetworksResponse) ProtoMessage() {}

func (x *BootstrapPeerListNetworksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_bootstrap_peer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapPeerListNetworksResponse.ProtoReflect.Descriptor instead.
func (*BootstrapPeerListNetworksResponse) Descriptor() ([]byte, []int) {
	return file_network_v1_bootstrap_peer_proto_rawDescGZIP(), []int{3}
}

type BootstrapPeerPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate *certificate.Certificate `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *BootstrapPeerPublishRequest) Reset() {
	*x = BootstrapPeerPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_bootstrap_peer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapPeerPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapPeerPublishRequest) ProtoMessage() {}

func (x *BootstrapPeerPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_bootstrap_peer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapPeerPublishRequest.ProtoReflect.Descriptor instead.
func (*BootstrapPeerPublishRequest) Descriptor() ([]byte, []int) {
	return file_network_v1_bootstrap_peer_proto_rawDescGZIP(), []int{4}
}

func (x *BootstrapPeerPublishRequest) GetCertificate() *certificate.Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

type BootstrapPeerPublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BootstrapPeerPublishResponse) Reset() {
	*x = BootstrapPeerPublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_bootstrap_peer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapPeerPublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapPeerPublishResponse) ProtoMessage() {}

func (x *BootstrapPeerPublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_bootstrap_peer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapPeerPublishResponse.ProtoReflect.Descriptor instead.
func (*BootstrapPeerPublishResponse) Descriptor() ([]byte, []int) {
	return file_network_v1_bootstrap_peer_proto_rawDescGZIP(), []int{5}
}

var File_network_v1_bootstrap_peer_proto protoreflect.FileDescriptor

var file_network_v1_bootstrap_peer_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x1a, 0x16,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x25, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x50, 0x65, 0x65, 0x72, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x42, 0x0a, 0x26, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x50, 0x65, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x22, 0x22, 0x0a, 0x20, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x50, 0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x21, 0x42, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x0a, 0x1b,
	0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x50, 0x65, 0x65, 0x72, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x1e, 0x0a, 0x1c, 0x42, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x50, 0x65, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbc, 0x03, 0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x42, 0x2e,
	0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x50, 0x65, 0x65, 0x72, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x43, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e,
	0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x50, 0x65, 0x65, 0x72, 0x47, 0x65, 0x74,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x3d, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x50,
	0x65, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x50, 0x65,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7e, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x12, 0x38, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e,
	0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x50, 0x65, 0x65, 0x72, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x50, 0x65, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6a, 0x0a, 0x1e, 0x67, 0x67, 0x2e, 0x73, 0x74, 0x72,
	0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x62,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x6d, 0x65, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72,
	0x61, 0x70, 0x3b, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0xba, 0x02, 0x03, 0x53,
	0x4e, 0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_v1_bootstrap_peer_proto_rawDescOnce sync.Once
	file_network_v1_bootstrap_peer_proto_rawDescData = file_network_v1_bootstrap_peer_proto_rawDesc
)

func file_network_v1_bootstrap_peer_proto_rawDescGZIP() []byte {
	file_network_v1_bootstrap_peer_proto_rawDescOnce.Do(func() {
		file_network_v1_bootstrap_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_v1_bootstrap_peer_proto_rawDescData)
	})
	return file_network_v1_bootstrap_peer_proto_rawDescData
}

var file_network_v1_bootstrap_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_network_v1_bootstrap_peer_proto_goTypes = []interface{}{
	(*BootstrapPeerGetPublishEnabledRequest)(nil),  // 0: strims.network.v1.bootstrap.BootstrapPeerGetPublishEnabledRequest
	(*BootstrapPeerGetPublishEnabledResponse)(nil), // 1: strims.network.v1.bootstrap.BootstrapPeerGetPublishEnabledResponse
	(*BootstrapPeerListNetworksRequest)(nil),       // 2: strims.network.v1.bootstrap.BootstrapPeerListNetworksRequest
	(*BootstrapPeerListNetworksResponse)(nil),      // 3: strims.network.v1.bootstrap.BootstrapPeerListNetworksResponse
	(*BootstrapPeerPublishRequest)(nil),            // 4: strims.network.v1.bootstrap.BootstrapPeerPublishRequest
	(*BootstrapPeerPublishResponse)(nil),           // 5: strims.network.v1.bootstrap.BootstrapPeerPublishResponse
	(*certificate.Certificate)(nil),                // 6: strims.type.Certificate
}
var file_network_v1_bootstrap_peer_proto_depIdxs = []int32{
	6, // 0: strims.network.v1.bootstrap.BootstrapPeerPublishRequest.certificate:type_name -> strims.type.Certificate
	0, // 1: strims.network.v1.bootstrap.PeerService.GetPublishEnabled:input_type -> strims.network.v1.bootstrap.BootstrapPeerGetPublishEnabledRequest
	2, // 2: strims.network.v1.bootstrap.PeerService.ListNetworks:input_type -> strims.network.v1.bootstrap.BootstrapPeerListNetworksRequest
	4, // 3: strims.network.v1.bootstrap.PeerService.Publish:input_type -> strims.network.v1.bootstrap.BootstrapPeerPublishRequest
	1, // 4: strims.network.v1.bootstrap.PeerService.GetPublishEnabled:output_type -> strims.network.v1.bootstrap.BootstrapPeerGetPublishEnabledResponse
	3, // 5: strims.network.v1.bootstrap.PeerService.ListNetworks:output_type -> strims.network.v1.bootstrap.BootstrapPeerListNetworksResponse
	5, // 6: strims.network.v1.bootstrap.PeerService.Publish:output_type -> strims.network.v1.bootstrap.BootstrapPeerPublishResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_network_v1_bootstrap_peer_proto_init() }
func file_network_v1_bootstrap_peer_proto_init() {
	if File_network_v1_bootstrap_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_v1_bootstrap_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapPeerGetPublishEnabledRequest); i {
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
		file_network_v1_bootstrap_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapPeerGetPublishEnabledResponse); i {
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
		file_network_v1_bootstrap_peer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapPeerListNetworksRequest); i {
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
		file_network_v1_bootstrap_peer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapPeerListNetworksResponse); i {
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
		file_network_v1_bootstrap_peer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapPeerPublishRequest); i {
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
		file_network_v1_bootstrap_peer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapPeerPublishResponse); i {
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
			RawDescriptor: file_network_v1_bootstrap_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_v1_bootstrap_peer_proto_goTypes,
		DependencyIndexes: file_network_v1_bootstrap_peer_proto_depIdxs,
		MessageInfos:      file_network_v1_bootstrap_peer_proto_msgTypes,
	}.Build()
	File_network_v1_bootstrap_peer_proto = out.File
	file_network_v1_bootstrap_peer_proto_rawDesc = nil
	file_network_v1_bootstrap_peer_proto_goTypes = nil
	file_network_v1_bootstrap_peer_proto_depIdxs = nil
}
