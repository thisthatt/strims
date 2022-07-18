// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.1
// source: network/v1/invite/invite.proto

package networkv1invite

import (
	v1 "github.com/MemeLabs/strims/pkg/apis/network/v1"
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

type GetInvitationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetInvitationRequest) Reset() {
	*x = GetInvitationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_invite_invite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvitationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvitationRequest) ProtoMessage() {}

func (x *GetInvitationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_invite_invite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvitationRequest.ProtoReflect.Descriptor instead.
func (*GetInvitationRequest) Descriptor() ([]byte, []int) {
	return file_network_v1_invite_invite_proto_rawDescGZIP(), []int{0}
}

func (x *GetInvitationRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetInvitationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invitation *v1.Invitation `protobuf:"bytes,1,opt,name=invitation,proto3" json:"invitation,omitempty"`
}

func (x *GetInvitationResponse) Reset() {
	*x = GetInvitationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_invite_invite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvitationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvitationResponse) ProtoMessage() {}

func (x *GetInvitationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_invite_invite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvitationResponse.ProtoReflect.Descriptor instead.
func (*GetInvitationResponse) Descriptor() ([]byte, []int) {
	return file_network_v1_invite_invite_proto_rawDescGZIP(), []int{1}
}

func (x *GetInvitationResponse) GetInvitation() *v1.Invitation {
	if x != nil {
		return x.Invitation
	}
	return nil
}

var File_network_v1_invite_invite_proto protoreflect.FileDescriptor

var file_network_v1_invite_invite_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x1a, 0x18, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x56, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x7e, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x70, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6a, 0x0a, 0x1b, 0x67, 0x67, 0x2e, 0x73,
	0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x6d, 0x65, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x73, 0x74, 0x72,
	0x69, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x3b, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x76, 0x31, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0xba, 0x02,
	0x03, 0x53, 0x4e, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_v1_invite_invite_proto_rawDescOnce sync.Once
	file_network_v1_invite_invite_proto_rawDescData = file_network_v1_invite_invite_proto_rawDesc
)

func file_network_v1_invite_invite_proto_rawDescGZIP() []byte {
	file_network_v1_invite_invite_proto_rawDescOnce.Do(func() {
		file_network_v1_invite_invite_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_v1_invite_invite_proto_rawDescData)
	})
	return file_network_v1_invite_invite_proto_rawDescData
}

var file_network_v1_invite_invite_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_network_v1_invite_invite_proto_goTypes = []interface{}{
	(*GetInvitationRequest)(nil),  // 0: strims.network.v1.invite.GetInvitationRequest
	(*GetInvitationResponse)(nil), // 1: strims.network.v1.invite.GetInvitationResponse
	(*v1.Invitation)(nil),         // 2: strims.network.v1.Invitation
}
var file_network_v1_invite_invite_proto_depIdxs = []int32{
	2, // 0: strims.network.v1.invite.GetInvitationResponse.invitation:type_name -> strims.network.v1.Invitation
	0, // 1: strims.network.v1.invite.InviteLink.GetInvitation:input_type -> strims.network.v1.invite.GetInvitationRequest
	1, // 2: strims.network.v1.invite.InviteLink.GetInvitation:output_type -> strims.network.v1.invite.GetInvitationResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_network_v1_invite_invite_proto_init() }
func file_network_v1_invite_invite_proto_init() {
	if File_network_v1_invite_invite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_v1_invite_invite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvitationRequest); i {
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
		file_network_v1_invite_invite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvitationResponse); i {
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
			RawDescriptor: file_network_v1_invite_invite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_v1_invite_invite_proto_goTypes,
		DependencyIndexes: file_network_v1_invite_invite_proto_depIdxs,
		MessageInfos:      file_network_v1_invite_invite_proto_msgTypes,
	}.Build()
	File_network_v1_invite_invite_proto = out.File
	file_network_v1_invite_invite_proto_rawDesc = nil
	file_network_v1_invite_invite_proto_goTypes = nil
	file_network_v1_invite_invite_proto_depIdxs = nil
}
