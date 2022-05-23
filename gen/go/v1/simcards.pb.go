// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: simcards.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SimCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *SimCard) Reset() {
	*x = SimCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simcards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimCard) ProtoMessage() {}

func (x *SimCard) ProtoReflect() protoreflect.Message {
	mi := &file_simcards_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimCard.ProtoReflect.Descriptor instead.
func (*SimCard) Descriptor() ([]byte, []int) {
	return file_simcards_proto_rawDescGZIP(), []int{0}
}

func (x *SimCard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SimCard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SimCard) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetSimCardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SimCards []*SimCard `protobuf:"bytes,1,rep,name=sim_cards,json=simCards,proto3" json:"sim_cards,omitempty"`
}

func (x *GetSimCardsResponse) Reset() {
	*x = GetSimCardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simcards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSimCardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSimCardsResponse) ProtoMessage() {}

func (x *GetSimCardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simcards_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSimCardsResponse.ProtoReflect.Descriptor instead.
func (*GetSimCardsResponse) Descriptor() ([]byte, []int) {
	return file_simcards_proto_rawDescGZIP(), []int{1}
}

func (x *GetSimCardsResponse) GetSimCards() []*SimCard {
	if x != nil {
		return x.SimCards
	}
	return nil
}

type GetSimCardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnterpriseId string `protobuf:"bytes,1,opt,name=enterprise_id,json=enterpriseId,proto3" json:"enterprise_id,omitempty"`
	SiteId       string `protobuf:"bytes,2,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
}

func (x *GetSimCardsRequest) Reset() {
	*x = GetSimCardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simcards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSimCardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSimCardsRequest) ProtoMessage() {}

func (x *GetSimCardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simcards_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSimCardsRequest.ProtoReflect.Descriptor instead.
func (*GetSimCardsRequest) Descriptor() ([]byte, []int) {
	return file_simcards_proto_rawDescGZIP(), []int{2}
}

func (x *GetSimCardsRequest) GetEnterpriseId() string {
	if x != nil {
		return x.EnterpriseId
	}
	return ""
}

func (x *GetSimCardsRequest) GetSiteId() string {
	if x != nil {
		return x.SiteId
	}
	return ""
}

var File_simcards_proto protoreflect.FileDescriptor

var file_simcards_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x69, 0x6d, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x07, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x5f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x73, 0x69,
	0x6d, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x42, 0x1e, 0x92, 0x41, 0x1b, 0x32,
	0x19, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x53, 0x69, 0x6d, 0x43, 0x61,
	0x72, 0x64, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x08, 0x73, 0x69, 0x6d, 0x43,
	0x61, 0x72, 0x64, 0x73, 0x22, 0xf5, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x43,
	0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x77, 0x0a, 0x0d, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x52, 0x92, 0x41, 0x4f, 0x2a, 0x0c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x73, 0x65, 0x49, 0x64, 0x32, 0x23, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6e, 0x20,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x4a, 0x06, 0x22, 0x61, 0x63, 0x6d,
	0x65, 0x22, 0x8a, 0x01, 0x08, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0xa2, 0x02, 0x06,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69,
	0x73, 0x65, 0x49, 0x64, 0x12, 0x66, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4d, 0x92, 0x41, 0x4a, 0x2a, 0x06, 0x53, 0x69, 0x74, 0x65,
	0x49, 0x64, 0x32, 0x1c, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x73, 0x69, 0x74, 0x65,
	0x4a, 0x0e, 0x22, 0x61, 0x63, 0x6d, 0x65, 0x2d, 0x63, 0x68, 0x69, 0x63, 0x61, 0x67, 0x6f, 0x22,
	0x8a, 0x01, 0x08, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0xa2, 0x02, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x32, 0xad, 0x01, 0x0a,
	0x0e, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x9a, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x54, 0x5a, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x6d, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x12, 0x3d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x73, 0x2f, 0x7b, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x73, 0x69, 0x6d, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x6f, 0x73, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2d, 0x75,
	0x6d, 0x62, 0x72, 0x65, 0x6c, 0x6c, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simcards_proto_rawDescOnce sync.Once
	file_simcards_proto_rawDescData = file_simcards_proto_rawDesc
)

func file_simcards_proto_rawDescGZIP() []byte {
	file_simcards_proto_rawDescOnce.Do(func() {
		file_simcards_proto_rawDescData = protoimpl.X.CompressGZIP(file_simcards_proto_rawDescData)
	})
	return file_simcards_proto_rawDescData
}

var file_simcards_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_simcards_proto_goTypes = []interface{}{
	(*SimCard)(nil),             // 0: v1.SimCard
	(*GetSimCardsResponse)(nil), // 1: v1.GetSimCardsResponse
	(*GetSimCardsRequest)(nil),  // 2: v1.GetSimCardsRequest
}
var file_simcards_proto_depIdxs = []int32{
	0, // 0: v1.GetSimCardsResponse.sim_cards:type_name -> v1.SimCard
	2, // 1: v1.SimCardService.GetSimCards:input_type -> v1.GetSimCardsRequest
	1, // 2: v1.SimCardService.GetSimCards:output_type -> v1.GetSimCardsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_simcards_proto_init() }
func file_simcards_proto_init() {
	if File_simcards_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simcards_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimCard); i {
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
		file_simcards_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSimCardsResponse); i {
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
		file_simcards_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSimCardsRequest); i {
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
			RawDescriptor: file_simcards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_simcards_proto_goTypes,
		DependencyIndexes: file_simcards_proto_depIdxs,
		MessageInfos:      file_simcards_proto_msgTypes,
	}.Build()
	File_simcards_proto = out.File
	file_simcards_proto_rawDesc = nil
	file_simcards_proto_goTypes = nil
	file_simcards_proto_depIdxs = nil
}
