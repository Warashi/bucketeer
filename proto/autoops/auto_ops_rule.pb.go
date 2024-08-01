// Copyright 2024 The Bucketeer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: proto/autoops/auto_ops_rule.proto

package autoops

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

type OpsType int32

const (
	// Refactor as Unknown instead of deleting
	//
	// Deprecated: Do not use.
	OpsType_ENABLE_FEATURE OpsType = 0
	// Deprecated: Do not use.
	OpsType_DISABLE_FEATURE OpsType = 1
	OpsType_SCHEDULE        OpsType = 2
	OpsType_EVENT_RATE      OpsType = 3
)

// Enum value maps for OpsType.
var (
	OpsType_name = map[int32]string{
		0: "ENABLE_FEATURE",
		1: "DISABLE_FEATURE",
		2: "SCHEDULE",
		3: "EVENT_RATE",
	}
	OpsType_value = map[string]int32{
		"ENABLE_FEATURE":  0,
		"DISABLE_FEATURE": 1,
		"SCHEDULE":        2,
		"EVENT_RATE":      3,
	}
)

func (x OpsType) Enum() *OpsType {
	p := new(OpsType)
	*p = x
	return p
}

func (x OpsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpsType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_autoops_auto_ops_rule_proto_enumTypes[0].Descriptor()
}

func (OpsType) Type() protoreflect.EnumType {
	return &file_proto_autoops_auto_ops_rule_proto_enumTypes[0]
}

func (x OpsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpsType.Descriptor instead.
func (OpsType) EnumDescriptor() ([]byte, []int) {
	return file_proto_autoops_auto_ops_rule_proto_rawDescGZIP(), []int{0}
}

type AutoOpsStatus int32

const (
	AutoOpsStatus_WAITING  AutoOpsStatus = 0
	AutoOpsStatus_RUNNING  AutoOpsStatus = 1
	AutoOpsStatus_FINISHED AutoOpsStatus = 2
	AutoOpsStatus_STOPPED  AutoOpsStatus = 3
)

// Enum value maps for AutoOpsStatus.
var (
	AutoOpsStatus_name = map[int32]string{
		0: "WAITING",
		1: "RUNNING",
		2: "FINISHED",
		3: "STOPPED",
	}
	AutoOpsStatus_value = map[string]int32{
		"WAITING":  0,
		"RUNNING":  1,
		"FINISHED": 2,
		"STOPPED":  3,
	}
)

func (x AutoOpsStatus) Enum() *AutoOpsStatus {
	p := new(AutoOpsStatus)
	*p = x
	return p
}

func (x AutoOpsStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AutoOpsStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_autoops_auto_ops_rule_proto_enumTypes[1].Descriptor()
}

func (AutoOpsStatus) Type() protoreflect.EnumType {
	return &file_proto_autoops_auto_ops_rule_proto_enumTypes[1]
}

func (x AutoOpsStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AutoOpsStatus.Descriptor instead.
func (AutoOpsStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_autoops_auto_ops_rule_proto_rawDescGZIP(), []int{1}
}

type AutoOpsRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	FeatureId     string        `protobuf:"bytes,2,opt,name=feature_id,json=featureId,proto3" json:"feature_id"`
	OpsType       OpsType       `protobuf:"varint,3,opt,name=ops_type,json=opsType,proto3,enum=bucketeer.autoops.OpsType" json:"ops_type"`
	Clauses       []*Clause     `protobuf:"bytes,4,rep,name=clauses,proto3" json:"clauses"`
	CreatedAt     int64         `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     int64         `protobuf:"varint,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Deleted       bool          `protobuf:"varint,9,opt,name=deleted,proto3" json:"deleted"`
	AutoOpsStatus AutoOpsStatus `protobuf:"varint,10,opt,name=auto_ops_status,json=autoOpsStatus,proto3,enum=bucketeer.autoops.AutoOpsStatus" json:"auto_ops_status"`
}

func (x *AutoOpsRule) Reset() {
	*x = AutoOpsRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_autoops_auto_ops_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoOpsRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoOpsRule) ProtoMessage() {}

func (x *AutoOpsRule) ProtoReflect() protoreflect.Message {
	mi := &file_proto_autoops_auto_ops_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoOpsRule.ProtoReflect.Descriptor instead.
func (*AutoOpsRule) Descriptor() ([]byte, []int) {
	return file_proto_autoops_auto_ops_rule_proto_rawDescGZIP(), []int{0}
}

func (x *AutoOpsRule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AutoOpsRule) GetFeatureId() string {
	if x != nil {
		return x.FeatureId
	}
	return ""
}

func (x *AutoOpsRule) GetOpsType() OpsType {
	if x != nil {
		return x.OpsType
	}
	return OpsType_ENABLE_FEATURE
}

func (x *AutoOpsRule) GetClauses() []*Clause {
	if x != nil {
		return x.Clauses
	}
	return nil
}

func (x *AutoOpsRule) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AutoOpsRule) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *AutoOpsRule) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *AutoOpsRule) GetAutoOpsStatus() AutoOpsStatus {
	if x != nil {
		return x.AutoOpsStatus
	}
	return AutoOpsStatus_WAITING
}

type AutoOpsRules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutoOpsRules []*AutoOpsRule `protobuf:"bytes,1,rep,name=auto_ops_rules,json=autoOpsRules,proto3" json:"auto_ops_rules"`
}

func (x *AutoOpsRules) Reset() {
	*x = AutoOpsRules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_autoops_auto_ops_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoOpsRules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoOpsRules) ProtoMessage() {}

func (x *AutoOpsRules) ProtoReflect() protoreflect.Message {
	mi := &file_proto_autoops_auto_ops_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoOpsRules.ProtoReflect.Descriptor instead.
func (*AutoOpsRules) Descriptor() ([]byte, []int) {
	return file_proto_autoops_auto_ops_rule_proto_rawDescGZIP(), []int{1}
}

func (x *AutoOpsRules) GetAutoOpsRules() []*AutoOpsRule {
	if x != nil {
		return x.AutoOpsRules
	}
	return nil
}

var File_proto_autoops_auto_ops_rule_proto protoreflect.FileDescriptor

var file_proto_autoops_auto_ops_rule_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6f, 0x70, 0x73, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x6f, 0x70, 0x73, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6f, 0x70, 0x73, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd0, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x6f, 0x4f, 0x70, 0x73, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49,
	0x64, 0x12, 0x35, 0x0a, 0x08, 0x6f, 0x70, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6f, 0x70, 0x73, 0x2e, 0x4f, 0x70, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x6f, 0x70, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x75,
	0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6f, 0x70, 0x73, 0x2e, 0x43, 0x6c,
	0x61, 0x75, 0x73, 0x65, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x48, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x6f, 0x70,
	0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6f,
	0x70, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4f, 0x70, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x4f, 0x70, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a,
	0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0x54, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x6f, 0x4f, 0x70, 0x73,
	0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x6f, 0x70,
	0x73, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6f, 0x70,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4f, 0x70, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0c, 0x61,
	0x75, 0x74, 0x6f, 0x4f, 0x70, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x2a, 0x58, 0x0a, 0x07, 0x4f,
	0x70, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x0e, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x00, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x17,
	0x0a, 0x0f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52,
	0x45, 0x10, 0x01, 0x1a, 0x02, 0x08, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x43, 0x48, 0x45, 0x44,
	0x55, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x52,
	0x41, 0x54, 0x45, 0x10, 0x03, 0x2a, 0x44, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x6f, 0x4f, 0x70, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x03, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x65, 0x65, 0x72, 0x2d, 0x69, 0x6f, 0x2f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6f, 0x70, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_autoops_auto_ops_rule_proto_rawDescOnce sync.Once
	file_proto_autoops_auto_ops_rule_proto_rawDescData = file_proto_autoops_auto_ops_rule_proto_rawDesc
)

func file_proto_autoops_auto_ops_rule_proto_rawDescGZIP() []byte {
	file_proto_autoops_auto_ops_rule_proto_rawDescOnce.Do(func() {
		file_proto_autoops_auto_ops_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_autoops_auto_ops_rule_proto_rawDescData)
	})
	return file_proto_autoops_auto_ops_rule_proto_rawDescData
}

var file_proto_autoops_auto_ops_rule_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_autoops_auto_ops_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_autoops_auto_ops_rule_proto_goTypes = []interface{}{
	(OpsType)(0),         // 0: bucketeer.autoops.OpsType
	(AutoOpsStatus)(0),   // 1: bucketeer.autoops.AutoOpsStatus
	(*AutoOpsRule)(nil),  // 2: bucketeer.autoops.AutoOpsRule
	(*AutoOpsRules)(nil), // 3: bucketeer.autoops.AutoOpsRules
	(*Clause)(nil),       // 4: bucketeer.autoops.Clause
}
var file_proto_autoops_auto_ops_rule_proto_depIdxs = []int32{
	0, // 0: bucketeer.autoops.AutoOpsRule.ops_type:type_name -> bucketeer.autoops.OpsType
	4, // 1: bucketeer.autoops.AutoOpsRule.clauses:type_name -> bucketeer.autoops.Clause
	1, // 2: bucketeer.autoops.AutoOpsRule.auto_ops_status:type_name -> bucketeer.autoops.AutoOpsStatus
	2, // 3: bucketeer.autoops.AutoOpsRules.auto_ops_rules:type_name -> bucketeer.autoops.AutoOpsRule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_autoops_auto_ops_rule_proto_init() }
func file_proto_autoops_auto_ops_rule_proto_init() {
	if File_proto_autoops_auto_ops_rule_proto != nil {
		return
	}
	file_proto_autoops_clause_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_autoops_auto_ops_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoOpsRule); i {
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
		file_proto_autoops_auto_ops_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoOpsRules); i {
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
			RawDescriptor: file_proto_autoops_auto_ops_rule_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_autoops_auto_ops_rule_proto_goTypes,
		DependencyIndexes: file_proto_autoops_auto_ops_rule_proto_depIdxs,
		EnumInfos:         file_proto_autoops_auto_ops_rule_proto_enumTypes,
		MessageInfos:      file_proto_autoops_auto_ops_rule_proto_msgTypes,
	}.Build()
	File_proto_autoops_auto_ops_rule_proto = out.File
	file_proto_autoops_auto_ops_rule_proto_rawDesc = nil
	file_proto_autoops_auto_ops_rule_proto_goTypes = nil
	file_proto_autoops_auto_ops_rule_proto_depIdxs = nil
}
