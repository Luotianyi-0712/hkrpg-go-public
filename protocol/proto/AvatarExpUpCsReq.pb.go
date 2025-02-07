// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AvatarExpUpCsReq.proto

package proto

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

type AvatarExpUpCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAvatarId uint32        `protobuf:"varint,12,opt,name=base_avatar_id,json=baseAvatarId,proto3" json:"base_avatar_id,omitempty"`
	ItemCost     *ItemCostData `protobuf:"bytes,14,opt,name=item_cost,json=itemCost,proto3" json:"item_cost,omitempty"`
}

func (x *AvatarExpUpCsReq) Reset() {
	*x = AvatarExpUpCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarExpUpCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarExpUpCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarExpUpCsReq) ProtoMessage() {}

func (x *AvatarExpUpCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarExpUpCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarExpUpCsReq.ProtoReflect.Descriptor instead.
func (*AvatarExpUpCsReq) Descriptor() ([]byte, []int) {
	return file_AvatarExpUpCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarExpUpCsReq) GetBaseAvatarId() uint32 {
	if x != nil {
		return x.BaseAvatarId
	}
	return 0
}

func (x *AvatarExpUpCsReq) GetItemCost() *ItemCostData {
	if x != nil {
		return x.ItemCost
	}
	return nil
}

var File_AvatarExpUpCsReq_proto protoreflect.FileDescriptor

var file_AvatarExpUpCsReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x55, 0x70, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x10,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x70, 0x55, 0x70, 0x43, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x6f,
	0x73, 0x74, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarExpUpCsReq_proto_rawDescOnce sync.Once
	file_AvatarExpUpCsReq_proto_rawDescData = file_AvatarExpUpCsReq_proto_rawDesc
)

func file_AvatarExpUpCsReq_proto_rawDescGZIP() []byte {
	file_AvatarExpUpCsReq_proto_rawDescOnce.Do(func() {
		file_AvatarExpUpCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarExpUpCsReq_proto_rawDescData)
	})
	return file_AvatarExpUpCsReq_proto_rawDescData
}

var file_AvatarExpUpCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarExpUpCsReq_proto_goTypes = []interface{}{
	(*AvatarExpUpCsReq)(nil), // 0: AvatarExpUpCsReq
	(*ItemCostData)(nil),     // 1: ItemCostData
}
var file_AvatarExpUpCsReq_proto_depIdxs = []int32{
	1, // 0: AvatarExpUpCsReq.item_cost:type_name -> ItemCostData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AvatarExpUpCsReq_proto_init() }
func file_AvatarExpUpCsReq_proto_init() {
	if File_AvatarExpUpCsReq_proto != nil {
		return
	}
	file_ItemCostData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarExpUpCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarExpUpCsReq); i {
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
			RawDescriptor: file_AvatarExpUpCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarExpUpCsReq_proto_goTypes,
		DependencyIndexes: file_AvatarExpUpCsReq_proto_depIdxs,
		MessageInfos:      file_AvatarExpUpCsReq_proto_msgTypes,
	}.Build()
	File_AvatarExpUpCsReq_proto = out.File
	file_AvatarExpUpCsReq_proto_rawDesc = nil
	file_AvatarExpUpCsReq_proto_goTypes = nil
	file_AvatarExpUpCsReq_proto_depIdxs = nil
}
