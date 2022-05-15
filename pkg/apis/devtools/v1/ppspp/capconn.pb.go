// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: devtools/v1/ppspp/capconn.proto

package ppspp

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

type CapConnLog_PeerLog_Event_Code int32

const (
	CapConnLog_PeerLog_Event_EVENT_CODE_INIT      CapConnLog_PeerLog_Event_Code = 0
	CapConnLog_PeerLog_Event_EVENT_CODE_WRITE     CapConnLog_PeerLog_Event_Code = 1
	CapConnLog_PeerLog_Event_EVENT_CODE_WRITE_ERR CapConnLog_PeerLog_Event_Code = 2
	CapConnLog_PeerLog_Event_EVENT_CODE_FLUSH     CapConnLog_PeerLog_Event_Code = 3
	CapConnLog_PeerLog_Event_EVENT_CODE_FLUSH_ERR CapConnLog_PeerLog_Event_Code = 4
	CapConnLog_PeerLog_Event_EVENT_CODE_READ      CapConnLog_PeerLog_Event_Code = 5
	CapConnLog_PeerLog_Event_EVENT_CODE_READ_ERR  CapConnLog_PeerLog_Event_Code = 6
)

// Enum value maps for CapConnLog_PeerLog_Event_Code.
var (
	CapConnLog_PeerLog_Event_Code_name = map[int32]string{
		0: "EVENT_CODE_INIT",
		1: "EVENT_CODE_WRITE",
		2: "EVENT_CODE_WRITE_ERR",
		3: "EVENT_CODE_FLUSH",
		4: "EVENT_CODE_FLUSH_ERR",
		5: "EVENT_CODE_READ",
		6: "EVENT_CODE_READ_ERR",
	}
	CapConnLog_PeerLog_Event_Code_value = map[string]int32{
		"EVENT_CODE_INIT":      0,
		"EVENT_CODE_WRITE":     1,
		"EVENT_CODE_WRITE_ERR": 2,
		"EVENT_CODE_FLUSH":     3,
		"EVENT_CODE_FLUSH_ERR": 4,
		"EVENT_CODE_READ":      5,
		"EVENT_CODE_READ_ERR":  6,
	}
)

func (x CapConnLog_PeerLog_Event_Code) Enum() *CapConnLog_PeerLog_Event_Code {
	p := new(CapConnLog_PeerLog_Event_Code)
	*p = x
	return p
}

func (x CapConnLog_PeerLog_Event_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CapConnLog_PeerLog_Event_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_devtools_v1_ppspp_capconn_proto_enumTypes[0].Descriptor()
}

func (CapConnLog_PeerLog_Event_Code) Type() protoreflect.EnumType {
	return &file_devtools_v1_ppspp_capconn_proto_enumTypes[0]
}

func (x CapConnLog_PeerLog_Event_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CapConnLog_PeerLog_Event_Code.Descriptor instead.
func (CapConnLog_PeerLog_Event_Code) EnumDescriptor() ([]byte, []int) {
	return file_devtools_v1_ppspp_capconn_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

type CapConnLog_PeerLog_Event_MessageType int32

const (
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_HANDSHAKE        CapConnLog_PeerLog_Event_MessageType = 0
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_DATA             CapConnLog_PeerLog_Event_MessageType = 1
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_ACK              CapConnLog_PeerLog_Event_MessageType = 2
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_HAVE             CapConnLog_PeerLog_Event_MessageType = 3
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_INTEGRITY        CapConnLog_PeerLog_Event_MessageType = 4
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_SIGNED_INTEGRITY CapConnLog_PeerLog_Event_MessageType = 5
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_REQUEST          CapConnLog_PeerLog_Event_MessageType = 6
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_CANCEL           CapConnLog_PeerLog_Event_MessageType = 7
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_CHOKE            CapConnLog_PeerLog_Event_MessageType = 8
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_UNCHOKE          CapConnLog_PeerLog_Event_MessageType = 9
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_PING             CapConnLog_PeerLog_Event_MessageType = 10
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_PONG             CapConnLog_PeerLog_Event_MessageType = 11
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_STREAM_REQUEST   CapConnLog_PeerLog_Event_MessageType = 12
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_STREAM_CANCEL    CapConnLog_PeerLog_Event_MessageType = 13
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_STREAM_OPEN      CapConnLog_PeerLog_Event_MessageType = 14
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_STREAM_CLOSE     CapConnLog_PeerLog_Event_MessageType = 15
	CapConnLog_PeerLog_Event_MESSAGE_TYPE_END              CapConnLog_PeerLog_Event_MessageType = 255
)

// Enum value maps for CapConnLog_PeerLog_Event_MessageType.
var (
	CapConnLog_PeerLog_Event_MessageType_name = map[int32]string{
		0:   "MESSAGE_TYPE_HANDSHAKE",
		1:   "MESSAGE_TYPE_DATA",
		2:   "MESSAGE_TYPE_ACK",
		3:   "MESSAGE_TYPE_HAVE",
		4:   "MESSAGE_TYPE_INTEGRITY",
		5:   "MESSAGE_TYPE_SIGNED_INTEGRITY",
		6:   "MESSAGE_TYPE_REQUEST",
		7:   "MESSAGE_TYPE_CANCEL",
		8:   "MESSAGE_TYPE_CHOKE",
		9:   "MESSAGE_TYPE_UNCHOKE",
		10:  "MESSAGE_TYPE_PING",
		11:  "MESSAGE_TYPE_PONG",
		12:  "MESSAGE_TYPE_STREAM_REQUEST",
		13:  "MESSAGE_TYPE_STREAM_CANCEL",
		14:  "MESSAGE_TYPE_STREAM_OPEN",
		15:  "MESSAGE_TYPE_STREAM_CLOSE",
		255: "MESSAGE_TYPE_END",
	}
	CapConnLog_PeerLog_Event_MessageType_value = map[string]int32{
		"MESSAGE_TYPE_HANDSHAKE":        0,
		"MESSAGE_TYPE_DATA":             1,
		"MESSAGE_TYPE_ACK":              2,
		"MESSAGE_TYPE_HAVE":             3,
		"MESSAGE_TYPE_INTEGRITY":        4,
		"MESSAGE_TYPE_SIGNED_INTEGRITY": 5,
		"MESSAGE_TYPE_REQUEST":          6,
		"MESSAGE_TYPE_CANCEL":           7,
		"MESSAGE_TYPE_CHOKE":            8,
		"MESSAGE_TYPE_UNCHOKE":          9,
		"MESSAGE_TYPE_PING":             10,
		"MESSAGE_TYPE_PONG":             11,
		"MESSAGE_TYPE_STREAM_REQUEST":   12,
		"MESSAGE_TYPE_STREAM_CANCEL":    13,
		"MESSAGE_TYPE_STREAM_OPEN":      14,
		"MESSAGE_TYPE_STREAM_CLOSE":     15,
		"MESSAGE_TYPE_END":              255,
	}
)

func (x CapConnLog_PeerLog_Event_MessageType) Enum() *CapConnLog_PeerLog_Event_MessageType {
	p := new(CapConnLog_PeerLog_Event_MessageType)
	*p = x
	return p
}

func (x CapConnLog_PeerLog_Event_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CapConnLog_PeerLog_Event_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_devtools_v1_ppspp_capconn_proto_enumTypes[1].Descriptor()
}

func (CapConnLog_PeerLog_Event_MessageType) Type() protoreflect.EnumType {
	return &file_devtools_v1_ppspp_capconn_proto_enumTypes[1]
}

func (x CapConnLog_PeerLog_Event_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CapConnLog_PeerLog_Event_MessageType.Descriptor instead.
func (CapConnLog_PeerLog_Event_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_devtools_v1_ppspp_capconn_proto_rawDescGZIP(), []int{0, 0, 0, 1}
}

type CapConnWatchLogsResponse_Op int32

const (
	CapConnWatchLogsResponse_CREATE CapConnWatchLogsResponse_Op = 0
	CapConnWatchLogsResponse_REMOVE CapConnWatchLogsResponse_Op = 1
)

// Enum value maps for CapConnWatchLogsResponse_Op.
var (
	CapConnWatchLogsResponse_Op_name = map[int32]string{
		0: "CREATE",
		1: "REMOVE",
	}
	CapConnWatchLogsResponse_Op_value = map[string]int32{
		"CREATE": 0,
		"REMOVE": 1,
	}
)

func (x CapConnWatchLogsResponse_Op) Enum() *CapConnWatchLogsResponse_Op {
	p := new(CapConnWatchLogsResponse_Op)
	*p = x
	return p
}

func (x CapConnWatchLogsResponse_Op) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CapConnWatchLogsResponse_Op) Descriptor() protoreflect.EnumDescriptor {
	return file_devtools_v1_ppspp_capconn_proto_enumTypes[2].Descriptor()
}

func (CapConnWatchLogsResponse_Op) Type() protoreflect.EnumType {
	return &file_devtools_v1_ppspp_capconn_proto_enumTypes[2]
}

func (x CapConnWatchLogsResponse_Op) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CapConnWatchLogsResponse_Op.Descriptor instead.
func (CapConnWatchLogsResponse_Op) EnumDescriptor() ([]byte, []int) {
	return file_devtools_v1_ppspp_capconn_proto_rawDescGZIP(), []int{2, 0}
}

type CapConnLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerLogs []*CapConnLog_PeerLog `protobuf:"bytes,1,rep,name=peer_logs,json=peerLogs,proto3" json:"peer_logs,omitempty"`
}

func (x *CapConnLog) Reset() {
	*x = CapConnLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapConnLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapConnLog) ProtoMessage() {}

func (x *CapConnLog) ProtoReflect() protoreflect.Message {
	mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapConnLog.ProtoReflect.Descriptor instead.
func (*CapConnLog) Descriptor() ([]byte, []int) {
	return file_devtools_v1_ppspp_capconn_proto_rawDescGZIP(), []int{0}
}

func (x *CapConnLog) GetPeerLogs() []*CapConnLog_PeerLog {
	if x != nil {
		return x.PeerLogs
	}
	return nil
}

type CapConnWatchLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CapConnWatchLogsRequest) Reset() {
	*x = CapConnWatchLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapConnWatchLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapConnWatchLogsRequest) ProtoMessage() {}

func (x *CapConnWatchLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapConnWatchLogsRequest.ProtoReflect.Descriptor instead.
func (*CapConnWatchLogsRequest) Descriptor() ([]byte, []int) {
	return file_devtools_v1_ppspp_capconn_proto_rawDescGZIP(), []int{1}
}

type CapConnWatchLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op   CapConnWatchLogsResponse_Op `protobuf:"varint,1,opt,name=op,proto3,enum=strims.devtools.v1.ppspp.CapConnWatchLogsResponse_Op" json:"op,omitempty"`
	Name string                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CapConnWatchLogsResponse) Reset() {
	*x = CapConnWatchLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapConnWatchLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapConnWatchLogsResponse) ProtoMessage() {}

func (x *CapConnWatchLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapConnWatchLogsResponse.ProtoReflect.Descriptor instead.
func (*CapConnWatchLogsResponse) Descriptor() ([]byte, []int) {
	return file_devtools_v1_ppspp_capconn_proto_rawDescGZIP(), []int{2}
}

func (x *CapConnWatchLogsResponse) GetOp() CapConnWatchLogsResponse_Op {
	if x != nil {
		return x.Op
	}
	return CapConnWatchLogsResponse_CREATE
}

func (x *CapConnWatchLogsResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CapConnLoadLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CapConnLoadLogRequest) Reset() {
	*x = CapConnLoadLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapConnLoadLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapConnLoadLogRequest) ProtoMessage() {}

func (x *CapConnLoadLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapConnLoadLogRequest.ProtoReflect.Descriptor instead.
func (*CapConnLoadLogRequest) Descriptor() ([]byte, []int) {
	return file_devtools_v1_ppspp_capconn_proto_rawDescGZIP(), []int{3}
}

func (x *CapConnLoadLogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CapConnLoadLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log *CapConnLog `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *CapConnLoadLogResponse) Reset() {
	*x = CapConnLoadLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapConnLoadLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapConnLoadLogResponse) ProtoMessage() {}

func (x *CapConnLoadLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapConnLoadLogResponse.ProtoReflect.Descriptor instead.
func (*CapConnLoadLogResponse) Descriptor() ([]byte, []int) {
	return file_devtools_v1_ppspp_capconn_proto_rawDescGZIP(), []int{4}
}

func (x *CapConnLoadLogResponse) GetLog() *CapConnLog {
	if x != nil {
		return x.Log
	}
	return nil
}

type CapConnLog_PeerLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label  string                      `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Events []*CapConnLog_PeerLog_Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *CapConnLog_PeerLog) Reset() {
	*x = CapConnLog_PeerLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapConnLog_PeerLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapConnLog_PeerLog) ProtoMessage() {}

func (x *CapConnLog_PeerLog) ProtoReflect() protoreflect.Message {
	mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapConnLog_PeerLog.ProtoReflect.Descriptor instead.
func (*CapConnLog_PeerLog) Descriptor() ([]byte, []int) {
	return file_devtools_v1_ppspp_capconn_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CapConnLog_PeerLog) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *CapConnLog_PeerLog) GetEvents() []*CapConnLog_PeerLog_Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type CapConnLog_PeerLog_Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             CapConnLog_PeerLog_Event_Code          `protobuf:"varint,1,opt,name=code,proto3,enum=strims.devtools.v1.ppspp.CapConnLog_PeerLog_Event_Code" json:"code,omitempty"`
	Timestamp        int64                                  `protobuf:"fixed64,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MessageTypes     []CapConnLog_PeerLog_Event_MessageType `protobuf:"varint,3,rep,packed,name=message_types,json=messageTypes,proto3,enum=strims.devtools.v1.ppspp.CapConnLog_PeerLog_Event_MessageType" json:"message_types,omitempty"`
	MessageAddresses []uint64                               `protobuf:"varint,4,rep,packed,name=message_addresses,json=messageAddresses,proto3" json:"message_addresses,omitempty"`
}

func (x *CapConnLog_PeerLog_Event) Reset() {
	*x = CapConnLog_PeerLog_Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapConnLog_PeerLog_Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapConnLog_PeerLog_Event) ProtoMessage() {}

func (x *CapConnLog_PeerLog_Event) ProtoReflect() protoreflect.Message {
	mi := &file_devtools_v1_ppspp_capconn_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapConnLog_PeerLog_Event.ProtoReflect.Descriptor instead.
func (*CapConnLog_PeerLog_Event) Descriptor() ([]byte, []int) {
	return file_devtools_v1_ppspp_capconn_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *CapConnLog_PeerLog_Event) GetCode() CapConnLog_PeerLog_Event_Code {
	if x != nil {
		return x.Code
	}
	return CapConnLog_PeerLog_Event_EVENT_CODE_INIT
}

func (x *CapConnLog_PeerLog_Event) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *CapConnLog_PeerLog_Event) GetMessageTypes() []CapConnLog_PeerLog_Event_MessageType {
	if x != nil {
		return x.MessageTypes
	}
	return nil
}

func (x *CapConnLog_PeerLog_Event) GetMessageAddresses() []uint64 {
	if x != nil {
		return x.MessageAddresses
	}
	return nil
}

var File_devtools_v1_ppspp_capconn_proto protoreflect.FileDescriptor

var file_devtools_v1_ppspp_capconn_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x70,
	0x73, 0x70, 0x70, 0x2f, 0x63, 0x61, 0x70, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x70, 0x73, 0x70, 0x70, 0x22, 0xcf, 0x08, 0x0a, 0x0a,
	0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x49, 0x0a, 0x09, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x70, 0x70, 0x73, 0x70, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e,
	0x4c, 0x6f, 0x67, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x08, 0x70, 0x65, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0xf5, 0x07, 0x0a, 0x07, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x4a, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x70, 0x73,
	0x70, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x1a, 0x87, 0x07, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x70, 0x73, 0x70, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x4c, 0x6f,
	0x67, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x10, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x63, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x3e, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x70, 0x73, 0x70, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x43, 0x6f,
	0x6e, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x2b, 0x0a,
	0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x01, 0x12, 0x18,
	0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x4c, 0x55, 0x53, 0x48, 0x10, 0x03, 0x12, 0x18,
	0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x4c, 0x55,
	0x53, 0x48, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x05, 0x12, 0x17, 0x0a,
	0x13, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44,
	0x5f, 0x45, 0x52, 0x52, 0x10, 0x06, 0x22, 0xd4, 0x03, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x53, 0x48, 0x41, 0x4b, 0x45,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x12,
	0x15, 0x0a, 0x11, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x48, 0x41, 0x56, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x52, 0x49, 0x54, 0x59,
	0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x52,
	0x49, 0x54, 0x59, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x06, 0x12,
	0x17, 0x0a, 0x13, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x4f, 0x4b, 0x45, 0x10, 0x08,
	0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x43, 0x48, 0x4f, 0x4b, 0x45, 0x10, 0x09, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x49, 0x4e, 0x47, 0x10,
	0x0a, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x50, 0x4f, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x0c, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x0d, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d,
	0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x0e, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x53, 0x53, 0x41,
	0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x43,
	0x4c, 0x4f, 0x53, 0x45, 0x10, 0x0f, 0x12, 0x15, 0x0a, 0x10, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0xff, 0x01, 0x22, 0x19, 0x0a,
	0x17, 0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x18, 0x43, 0x61, 0x70,
	0x43, 0x6f, 0x6e, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x35, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x70, 0x73, 0x70, 0x70, 0x2e, 0x43, 0x61, 0x70,
	0x43, 0x6f, 0x6e, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x1c, 0x0a, 0x02, 0x4f, 0x70, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x22, 0x2b,
	0x0a, 0x15, 0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x16, 0x43,
	0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x70, 0x73, 0x70, 0x70, 0x2e, 0x43, 0x61,
	0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x32, 0xed, 0x01,
	0x0a, 0x07, 0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x74, 0x0a, 0x09, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x70, 0x73, 0x70,
	0x70, 0x2e, 0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x73, 0x74, 0x72, 0x69,
	0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70,
	0x70, 0x73, 0x70, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x6c, 0x0a, 0x07, 0x4c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x2f, 0x2e, 0x73, 0x74, 0x72,
	0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x70, 0x70, 0x73, 0x70, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x4c, 0x6f, 0x61,
	0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x74,
	0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x70, 0x73, 0x70, 0x70, 0x2e, 0x43, 0x61, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x4c, 0x6f,
	0x61, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x60, 0x0a,
	0x1b, 0x67, 0x67, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x70, 0x73, 0x70, 0x70, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x65, 0x6d, 0x65, 0x4c, 0x61, 0x62,
	0x73, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x70,
	0x73, 0x70, 0x70, 0x3b, 0x70, 0x70, 0x73, 0x70, 0x70, 0xba, 0x02, 0x03, 0x53, 0x44, 0x54, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_devtools_v1_ppspp_capconn_proto_rawDescOnce sync.Once
	file_devtools_v1_ppspp_capconn_proto_rawDescData = file_devtools_v1_ppspp_capconn_proto_rawDesc
)

func file_devtools_v1_ppspp_capconn_proto_rawDescGZIP() []byte {
	file_devtools_v1_ppspp_capconn_proto_rawDescOnce.Do(func() {
		file_devtools_v1_ppspp_capconn_proto_rawDescData = protoimpl.X.CompressGZIP(file_devtools_v1_ppspp_capconn_proto_rawDescData)
	})
	return file_devtools_v1_ppspp_capconn_proto_rawDescData
}

var file_devtools_v1_ppspp_capconn_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_devtools_v1_ppspp_capconn_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_devtools_v1_ppspp_capconn_proto_goTypes = []interface{}{
	(CapConnLog_PeerLog_Event_Code)(0),        // 0: strims.devtools.v1.ppspp.CapConnLog.PeerLog.Event.Code
	(CapConnLog_PeerLog_Event_MessageType)(0), // 1: strims.devtools.v1.ppspp.CapConnLog.PeerLog.Event.MessageType
	(CapConnWatchLogsResponse_Op)(0),          // 2: strims.devtools.v1.ppspp.CapConnWatchLogsResponse.Op
	(*CapConnLog)(nil),                        // 3: strims.devtools.v1.ppspp.CapConnLog
	(*CapConnWatchLogsRequest)(nil),           // 4: strims.devtools.v1.ppspp.CapConnWatchLogsRequest
	(*CapConnWatchLogsResponse)(nil),          // 5: strims.devtools.v1.ppspp.CapConnWatchLogsResponse
	(*CapConnLoadLogRequest)(nil),             // 6: strims.devtools.v1.ppspp.CapConnLoadLogRequest
	(*CapConnLoadLogResponse)(nil),            // 7: strims.devtools.v1.ppspp.CapConnLoadLogResponse
	(*CapConnLog_PeerLog)(nil),                // 8: strims.devtools.v1.ppspp.CapConnLog.PeerLog
	(*CapConnLog_PeerLog_Event)(nil),          // 9: strims.devtools.v1.ppspp.CapConnLog.PeerLog.Event
}
var file_devtools_v1_ppspp_capconn_proto_depIdxs = []int32{
	8, // 0: strims.devtools.v1.ppspp.CapConnLog.peer_logs:type_name -> strims.devtools.v1.ppspp.CapConnLog.PeerLog
	2, // 1: strims.devtools.v1.ppspp.CapConnWatchLogsResponse.op:type_name -> strims.devtools.v1.ppspp.CapConnWatchLogsResponse.Op
	3, // 2: strims.devtools.v1.ppspp.CapConnLoadLogResponse.log:type_name -> strims.devtools.v1.ppspp.CapConnLog
	9, // 3: strims.devtools.v1.ppspp.CapConnLog.PeerLog.events:type_name -> strims.devtools.v1.ppspp.CapConnLog.PeerLog.Event
	0, // 4: strims.devtools.v1.ppspp.CapConnLog.PeerLog.Event.code:type_name -> strims.devtools.v1.ppspp.CapConnLog.PeerLog.Event.Code
	1, // 5: strims.devtools.v1.ppspp.CapConnLog.PeerLog.Event.message_types:type_name -> strims.devtools.v1.ppspp.CapConnLog.PeerLog.Event.MessageType
	4, // 6: strims.devtools.v1.ppspp.CapConn.WatchLogs:input_type -> strims.devtools.v1.ppspp.CapConnWatchLogsRequest
	6, // 7: strims.devtools.v1.ppspp.CapConn.LoadLog:input_type -> strims.devtools.v1.ppspp.CapConnLoadLogRequest
	5, // 8: strims.devtools.v1.ppspp.CapConn.WatchLogs:output_type -> strims.devtools.v1.ppspp.CapConnWatchLogsResponse
	7, // 9: strims.devtools.v1.ppspp.CapConn.LoadLog:output_type -> strims.devtools.v1.ppspp.CapConnLoadLogResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_devtools_v1_ppspp_capconn_proto_init() }
func file_devtools_v1_ppspp_capconn_proto_init() {
	if File_devtools_v1_ppspp_capconn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_devtools_v1_ppspp_capconn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapConnLog); i {
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
		file_devtools_v1_ppspp_capconn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapConnWatchLogsRequest); i {
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
		file_devtools_v1_ppspp_capconn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapConnWatchLogsResponse); i {
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
		file_devtools_v1_ppspp_capconn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapConnLoadLogRequest); i {
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
		file_devtools_v1_ppspp_capconn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapConnLoadLogResponse); i {
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
		file_devtools_v1_ppspp_capconn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapConnLog_PeerLog); i {
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
		file_devtools_v1_ppspp_capconn_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapConnLog_PeerLog_Event); i {
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
			RawDescriptor: file_devtools_v1_ppspp_capconn_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_devtools_v1_ppspp_capconn_proto_goTypes,
		DependencyIndexes: file_devtools_v1_ppspp_capconn_proto_depIdxs,
		EnumInfos:         file_devtools_v1_ppspp_capconn_proto_enumTypes,
		MessageInfos:      file_devtools_v1_ppspp_capconn_proto_msgTypes,
	}.Build()
	File_devtools_v1_ppspp_capconn_proto = out.File
	file_devtools_v1_ppspp_capconn_proto_rawDesc = nil
	file_devtools_v1_ppspp_capconn_proto_goTypes = nil
	file_devtools_v1_ppspp_capconn_proto_depIdxs = nil
}
