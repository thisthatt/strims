// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: dao/v1/dao.proto

package dao

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

type SecondaryIndexKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Id  uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SecondaryIndexKey) Reset() {
	*x = SecondaryIndexKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dao_v1_dao_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecondaryIndexKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecondaryIndexKey) ProtoMessage() {}

func (x *SecondaryIndexKey) ProtoReflect() protoreflect.Message {
	mi := &file_dao_v1_dao_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecondaryIndexKey.ProtoReflect.Descriptor instead.
func (*SecondaryIndexKey) Descriptor() ([]byte, []int) {
	return file_dao_v1_dao_proto_rawDescGZIP(), []int{0}
}

func (x *SecondaryIndexKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SecondaryIndexKey) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Mutex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eol   int64  `protobuf:"varint,1,opt,name=eol,proto3" json:"eol,omitempty"`
	Token []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Mutex) Reset() {
	*x = Mutex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dao_v1_dao_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mutex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mutex) ProtoMessage() {}

func (x *Mutex) ProtoReflect() protoreflect.Message {
	mi := &file_dao_v1_dao_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mutex.ProtoReflect.Descriptor instead.
func (*Mutex) Descriptor() ([]byte, []int) {
	return file_dao_v1_dao_proto_rawDescGZIP(), []int{1}
}

func (x *Mutex) GetEol() int64 {
	if x != nil {
		return x.Eol
	}
	return 0
}

func (x *Mutex) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_dao_v1_dao_proto protoreflect.FileDescriptor

var file_dao_v1_dao_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x6f, 0x2e, 0x76,
	0x31, 0x22, 0x35, 0x0a, 0x11, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x05, 0x4d, 0x75, 0x74, 0x65,
	0x78, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x65, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x48, 0x0a, 0x10, 0x67, 0x67, 0x2e,
	0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x61, 0x6f, 0x2e, 0x76, 0x31, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x6d, 0x65, 0x4c, 0x61,
	0x62, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x64, 0x61, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x6f, 0xba, 0x02, 0x03,
	0x53, 0x44, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dao_v1_dao_proto_rawDescOnce sync.Once
	file_dao_v1_dao_proto_rawDescData = file_dao_v1_dao_proto_rawDesc
)

func file_dao_v1_dao_proto_rawDescGZIP() []byte {
	file_dao_v1_dao_proto_rawDescOnce.Do(func() {
		file_dao_v1_dao_proto_rawDescData = protoimpl.X.CompressGZIP(file_dao_v1_dao_proto_rawDescData)
	})
	return file_dao_v1_dao_proto_rawDescData
}

var file_dao_v1_dao_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dao_v1_dao_proto_goTypes = []interface{}{
	(*SecondaryIndexKey)(nil), // 0: strims.dao.v1.SecondaryIndexKey
	(*Mutex)(nil),             // 1: strims.dao.v1.Mutex
}
var file_dao_v1_dao_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dao_v1_dao_proto_init() }
func file_dao_v1_dao_proto_init() {
	if File_dao_v1_dao_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dao_v1_dao_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecondaryIndexKey); i {
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
		file_dao_v1_dao_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mutex); i {
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
			RawDescriptor: file_dao_v1_dao_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dao_v1_dao_proto_goTypes,
		DependencyIndexes: file_dao_v1_dao_proto_depIdxs,
		MessageInfos:      file_dao_v1_dao_proto_msgTypes,
	}.Build()
	File_dao_v1_dao_proto = out.File
	file_dao_v1_dao_proto_rawDesc = nil
	file_dao_v1_dao_proto_goTypes = nil
	file_dao_v1_dao_proto_depIdxs = nil
}
