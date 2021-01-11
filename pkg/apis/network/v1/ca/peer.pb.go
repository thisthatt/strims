// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.1
// source: network/v1/ca/peer.proto

package ca

import (
	certificate "github.com/MemeLabs/go-ppspp/pkg/apis/type/certificate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CAPeerRenewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate        *certificate.Certificate        `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	CertificateRequest *certificate.CertificateRequest `protobuf:"bytes,2,opt,name=certificate_request,json=certificateRequest,proto3" json:"certificate_request,omitempty"`
}

func (x *CAPeerRenewRequest) Reset() {
	*x = CAPeerRenewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_ca_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CAPeerRenewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CAPeerRenewRequest) ProtoMessage() {}

func (x *CAPeerRenewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_ca_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CAPeerRenewRequest.ProtoReflect.Descriptor instead.
func (*CAPeerRenewRequest) Descriptor() ([]byte, []int) {
	return file_network_v1_ca_peer_proto_rawDescGZIP(), []int{0}
}

func (x *CAPeerRenewRequest) GetCertificate() *certificate.Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *CAPeerRenewRequest) GetCertificateRequest() *certificate.CertificateRequest {
	if x != nil {
		return x.CertificateRequest
	}
	return nil
}

type CAPeerRenewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate *certificate.Certificate `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *CAPeerRenewResponse) Reset() {
	*x = CAPeerRenewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_v1_ca_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CAPeerRenewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CAPeerRenewResponse) ProtoMessage() {}

func (x *CAPeerRenewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_v1_ca_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CAPeerRenewResponse.ProtoReflect.Descriptor instead.
func (*CAPeerRenewResponse) Descriptor() ([]byte, []int) {
	return file_network_v1_ca_peer_proto_rawDescGZIP(), []int{1}
}

func (x *CAPeerRenewResponse) GetCertificate() *certificate.Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

var File_network_v1_ca_peer_proto protoreflect.FileDescriptor

var file_network_v1_ca_peer_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x2f,
	0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x74, 0x72, 0x69,
	0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61,
	0x1a, 0x16, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x12, 0x43, 0x41, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0b,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x13, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d,
	0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x12, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a,
	0x13, 0x43, 0x41, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x69,
	0x6d, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x32, 0x66, 0x0a, 0x06, 0x43, 0x41, 0x50, 0x65, 0x65, 0x72, 0x12, 0x5c, 0x0a, 0x05, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x12, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x61, 0x2e, 0x43, 0x41, 0x50, 0x65, 0x65,
	0x72, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x73, 0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x61, 0x2e, 0x43, 0x41, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x57, 0x0a, 0x17, 0x67, 0x67, 0x2e, 0x73,
	0x74, 0x72, 0x69, 0x6d, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x61, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4d, 0x65, 0x6d, 0x65, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x70, 0x73, 0x70,
	0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x3b, 0x63, 0x61, 0xba, 0x02, 0x03, 0x53, 0x4e,
	0x43, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_v1_ca_peer_proto_rawDescOnce sync.Once
	file_network_v1_ca_peer_proto_rawDescData = file_network_v1_ca_peer_proto_rawDesc
)

func file_network_v1_ca_peer_proto_rawDescGZIP() []byte {
	file_network_v1_ca_peer_proto_rawDescOnce.Do(func() {
		file_network_v1_ca_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_v1_ca_peer_proto_rawDescData)
	})
	return file_network_v1_ca_peer_proto_rawDescData
}

var file_network_v1_ca_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_network_v1_ca_peer_proto_goTypes = []interface{}{
	(*CAPeerRenewRequest)(nil),             // 0: strims.network.v1.ca.CAPeerRenewRequest
	(*CAPeerRenewResponse)(nil),            // 1: strims.network.v1.ca.CAPeerRenewResponse
	(*certificate.Certificate)(nil),        // 2: strims.type.Certificate
	(*certificate.CertificateRequest)(nil), // 3: strims.type.CertificateRequest
}
var file_network_v1_ca_peer_proto_depIdxs = []int32{
	2, // 0: strims.network.v1.ca.CAPeerRenewRequest.certificate:type_name -> strims.type.Certificate
	3, // 1: strims.network.v1.ca.CAPeerRenewRequest.certificate_request:type_name -> strims.type.CertificateRequest
	2, // 2: strims.network.v1.ca.CAPeerRenewResponse.certificate:type_name -> strims.type.Certificate
	0, // 3: strims.network.v1.ca.CAPeer.Renew:input_type -> strims.network.v1.ca.CAPeerRenewRequest
	1, // 4: strims.network.v1.ca.CAPeer.Renew:output_type -> strims.network.v1.ca.CAPeerRenewResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_network_v1_ca_peer_proto_init() }
func file_network_v1_ca_peer_proto_init() {
	if File_network_v1_ca_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_v1_ca_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CAPeerRenewRequest); i {
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
		file_network_v1_ca_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CAPeerRenewResponse); i {
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
			RawDescriptor: file_network_v1_ca_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_v1_ca_peer_proto_goTypes,
		DependencyIndexes: file_network_v1_ca_peer_proto_depIdxs,
		MessageInfos:      file_network_v1_ca_peer_proto_msgTypes,
	}.Build()
	File_network_v1_ca_peer_proto = out.File
	file_network_v1_ca_peer_proto_rawDesc = nil
	file_network_v1_ca_peer_proto_goTypes = nil
	file_network_v1_ca_peer_proto_depIdxs = nil
}
