// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: devtools/v1/devtools.proto

package devtools

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

type DevToolsTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DevToolsTestRequest) Reset() {
	*x = DevToolsTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devtools_v1_devtools_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevToolsTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevToolsTestRequest) ProtoMessage() {}

func (x *DevToolsTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_devtools_v1_devtools_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevToolsTestRequest.ProtoReflect.Descriptor instead.
func (*DevToolsTestRequest) Descriptor() ([]byte, []int) {
	return file_devtools_v1_devtools_proto_rawDescGZIP(), []int{0}
}

func (x *DevToolsTestRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DevToolsTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DevToolsTestResponse) Reset() {
	*x = DevToolsTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devtools_v1_devtools_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevToolsTestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevToolsTestResponse) ProtoMessage() {}

func (x *DevToolsTestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_devtools_v1_devtools_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevToolsTestResponse.ProtoReflect.Descriptor instead.
func (*DevToolsTestResponse) Descriptor() ([]byte, []int) {
	return file_devtools_v1_devtools_proto_rawDescGZIP(), []int{1}
}

func (x *DevToolsTestResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_devtools_v1_devtools_proto protoreflect.FileDescriptor

var file_devtools_v1_devtools_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x22, 0x29, 0x0a, 0x13, 0x44, 0x65, 0x76, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x44,
	0x65, 0x76, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x65, 0x0a,
	0x08, 0x44, 0x65, 0x76, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x59, 0x0a, 0x04, 0x54, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x74, 0x72,
	0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x76, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x57, 0x0a, 0x15, 0x67, 0x67, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d,
	0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x6d, 0x65, 0x4c, 0x61,
	0x62, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0xba, 0x02, 0x03, 0x53, 0x44, 0x54, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_devtools_v1_devtools_proto_rawDescOnce sync.Once
	file_devtools_v1_devtools_proto_rawDescData = file_devtools_v1_devtools_proto_rawDesc
)

func file_devtools_v1_devtools_proto_rawDescGZIP() []byte {
	file_devtools_v1_devtools_proto_rawDescOnce.Do(func() {
		file_devtools_v1_devtools_proto_rawDescData = protoimpl.X.CompressGZIP(file_devtools_v1_devtools_proto_rawDescData)
	})
	return file_devtools_v1_devtools_proto_rawDescData
}

var file_devtools_v1_devtools_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_devtools_v1_devtools_proto_goTypes = []interface{}{
	(*DevToolsTestRequest)(nil),  // 0: strims.devtools.v1.DevToolsTestRequest
	(*DevToolsTestResponse)(nil), // 1: strims.devtools.v1.DevToolsTestResponse
}
var file_devtools_v1_devtools_proto_depIdxs = []int32{
	0, // 0: strims.devtools.v1.DevTools.Test:input_type -> strims.devtools.v1.DevToolsTestRequest
	1, // 1: strims.devtools.v1.DevTools.Test:output_type -> strims.devtools.v1.DevToolsTestResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_devtools_v1_devtools_proto_init() }
func file_devtools_v1_devtools_proto_init() {
	if File_devtools_v1_devtools_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_devtools_v1_devtools_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevToolsTestRequest); i {
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
		file_devtools_v1_devtools_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevToolsTestResponse); i {
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
			RawDescriptor: file_devtools_v1_devtools_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_devtools_v1_devtools_proto_goTypes,
		DependencyIndexes: file_devtools_v1_devtools_proto_depIdxs,
		MessageInfos:      file_devtools_v1_devtools_proto_msgTypes,
	}.Build()
	File_devtools_v1_devtools_proto = out.File
	file_devtools_v1_devtools_proto_rawDesc = nil
	file_devtools_v1_devtools_proto_goTypes = nil
	file_devtools_v1_devtools_proto_depIdxs = nil
}
