// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: devicegroups/v1/devicegroups.proto

package v1

import (
	_ "github.com/danielvladco/go-proto-gql/pb"
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

type DeviceGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *DeviceGroup) Reset() {
	*x = DeviceGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devicegroups_v1_devicegroups_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceGroup) ProtoMessage() {}

func (x *DeviceGroup) ProtoReflect() protoreflect.Message {
	mi := &file_devicegroups_v1_devicegroups_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceGroup.ProtoReflect.Descriptor instead.
func (*DeviceGroup) Descriptor() ([]byte, []int) {
	return file_devicegroups_v1_devicegroups_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetDeviceGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceGroups []*DeviceGroup `protobuf:"bytes,1,rep,name=device_groups,json=deviceGroups,proto3" json:"device_groups,omitempty"`
}

func (x *GetDeviceGroupsResponse) Reset() {
	*x = GetDeviceGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devicegroups_v1_devicegroups_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceGroupsResponse) ProtoMessage() {}

func (x *GetDeviceGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_devicegroups_v1_devicegroups_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceGroupsResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceGroupsResponse) Descriptor() ([]byte, []int) {
	return file_devicegroups_v1_devicegroups_proto_rawDescGZIP(), []int{1}
}

func (x *GetDeviceGroupsResponse) GetDeviceGroups() []*DeviceGroup {
	if x != nil {
		return x.DeviceGroups
	}
	return nil
}

type GetDeviceGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnterpriseId string `protobuf:"bytes,1,opt,name=enterprise_id,json=enterpriseId,proto3" json:"enterprise_id,omitempty"`
	SiteId       string `protobuf:"bytes,2,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
}

func (x *GetDeviceGroupsRequest) Reset() {
	*x = GetDeviceGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devicegroups_v1_devicegroups_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceGroupsRequest) ProtoMessage() {}

func (x *GetDeviceGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_devicegroups_v1_devicegroups_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceGroupsRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceGroupsRequest) Descriptor() ([]byte, []int) {
	return file_devicegroups_v1_devicegroups_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeviceGroupsRequest) GetEnterpriseId() string {
	if x != nil {
		return x.EnterpriseId
	}
	return ""
}

func (x *GetDeviceGroupsRequest) GetSiteId() string {
	if x != nil {
		return x.SiteId
	}
	return ""
}

var File_devicegroups_v1_devicegroups_proto protoreflect.FileDescriptor

var file_devicegroups_v1_devicegroups_proto_rawDesc = []byte{
	0x0a, 0x22, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x22, 0x92, 0x41,
	0x1f, 0x32, 0x1d, 0x41, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0xf9,
	0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x77, 0x0a, 0x0d, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x52, 0x92, 0x41, 0x4f, 0x2a, 0x0c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73,
	0x65, 0x49, 0x64, 0x32, 0x23, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x4a, 0x06, 0x22, 0x61, 0x63, 0x6d, 0x65, 0x22,
	0x8a, 0x01, 0x08, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0xa2, 0x02, 0x06, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x66, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x4d, 0x92, 0x41, 0x4a, 0x2a, 0x06, 0x53, 0x69, 0x74, 0x65, 0x49, 0x64,
	0x32, 0x1c, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x73, 0x69, 0x74, 0x65, 0x4a, 0x0e,
	0x22, 0x61, 0x63, 0x6d, 0x65, 0x2d, 0x63, 0x68, 0x69, 0x63, 0x61, 0x67, 0x6f, 0x22, 0x8a, 0x01,
	0x08, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x5d, 0xa2, 0x02, 0x06, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x06, 0x73, 0x69, 0x74, 0x65, 0x49, 0x64, 0x32, 0xe5, 0x01, 0x0a, 0x12, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xce, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x27, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x68, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x5c,
	0x5a, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x41, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x73, 0x2f, 0x7b,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73,
	0x69, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0xb2, 0xe0, 0x1f, 0x02,
	0x08, 0x02, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x6e, 0x6f, 0x73, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x63, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x2d, 0x75, 0x6d, 0x62, 0x72, 0x65, 0x6c, 0x6c, 0x61, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_devicegroups_v1_devicegroups_proto_rawDescOnce sync.Once
	file_devicegroups_v1_devicegroups_proto_rawDescData = file_devicegroups_v1_devicegroups_proto_rawDesc
)

func file_devicegroups_v1_devicegroups_proto_rawDescGZIP() []byte {
	file_devicegroups_v1_devicegroups_proto_rawDescOnce.Do(func() {
		file_devicegroups_v1_devicegroups_proto_rawDescData = protoimpl.X.CompressGZIP(file_devicegroups_v1_devicegroups_proto_rawDescData)
	})
	return file_devicegroups_v1_devicegroups_proto_rawDescData
}

var file_devicegroups_v1_devicegroups_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_devicegroups_v1_devicegroups_proto_goTypes = []interface{}{
	(*DeviceGroup)(nil),             // 0: devicegroups.v1.DeviceGroup
	(*GetDeviceGroupsResponse)(nil), // 1: devicegroups.v1.GetDeviceGroupsResponse
	(*GetDeviceGroupsRequest)(nil),  // 2: devicegroups.v1.GetDeviceGroupsRequest
}
var file_devicegroups_v1_devicegroups_proto_depIdxs = []int32{
	0, // 0: devicegroups.v1.GetDeviceGroupsResponse.device_groups:type_name -> devicegroups.v1.DeviceGroup
	2, // 1: devicegroups.v1.DeviceGroupService.GetDeviceGroups:input_type -> devicegroups.v1.GetDeviceGroupsRequest
	1, // 2: devicegroups.v1.DeviceGroupService.GetDeviceGroups:output_type -> devicegroups.v1.GetDeviceGroupsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_devicegroups_v1_devicegroups_proto_init() }
func file_devicegroups_v1_devicegroups_proto_init() {
	if File_devicegroups_v1_devicegroups_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_devicegroups_v1_devicegroups_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceGroup); i {
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
		file_devicegroups_v1_devicegroups_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceGroupsResponse); i {
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
		file_devicegroups_v1_devicegroups_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceGroupsRequest); i {
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
			RawDescriptor: file_devicegroups_v1_devicegroups_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_devicegroups_v1_devicegroups_proto_goTypes,
		DependencyIndexes: file_devicegroups_v1_devicegroups_proto_depIdxs,
		MessageInfos:      file_devicegroups_v1_devicegroups_proto_msgTypes,
	}.Build()
	File_devicegroups_v1_devicegroups_proto = out.File
	file_devicegroups_v1_devicegroups_proto_rawDesc = nil
	file_devicegroups_v1_devicegroups_proto_goTypes = nil
	file_devicegroups_v1_devicegroups_proto_depIdxs = nil
}
