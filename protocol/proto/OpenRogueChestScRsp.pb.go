// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: OpenRogueChestScRsp.proto

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

type OpenRogueChestScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DropData *ItemList `protobuf:"bytes,7,opt,name=drop_data,json=dropData,proto3" json:"drop_data,omitempty"`
	Retcode  uint32    `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Reward   *ItemList `protobuf:"bytes,14,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (x *OpenRogueChestScRsp) Reset() {
	*x = OpenRogueChestScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpenRogueChestScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenRogueChestScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenRogueChestScRsp) ProtoMessage() {}

func (x *OpenRogueChestScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_OpenRogueChestScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenRogueChestScRsp.ProtoReflect.Descriptor instead.
func (*OpenRogueChestScRsp) Descriptor() ([]byte, []int) {
	return file_OpenRogueChestScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *OpenRogueChestScRsp) GetDropData() *ItemList {
	if x != nil {
		return x.DropData
	}
	return nil
}

func (x *OpenRogueChestScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *OpenRogueChestScRsp) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

var File_OpenRogueChestScRsp_proto protoreflect.FileDescriptor

var file_OpenRogueChestScRsp_proto_rawDesc = []byte{
	0x0a, 0x19, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x68, 0x65, 0x73, 0x74,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x13, 0x4f,
	0x70, 0x65, 0x6e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x68, 0x65, 0x73, 0x74, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x12, 0x26, 0x0a, 0x09, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x08, 0x64, 0x72, 0x6f, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OpenRogueChestScRsp_proto_rawDescOnce sync.Once
	file_OpenRogueChestScRsp_proto_rawDescData = file_OpenRogueChestScRsp_proto_rawDesc
)

func file_OpenRogueChestScRsp_proto_rawDescGZIP() []byte {
	file_OpenRogueChestScRsp_proto_rawDescOnce.Do(func() {
		file_OpenRogueChestScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_OpenRogueChestScRsp_proto_rawDescData)
	})
	return file_OpenRogueChestScRsp_proto_rawDescData
}

var file_OpenRogueChestScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OpenRogueChestScRsp_proto_goTypes = []interface{}{
	(*OpenRogueChestScRsp)(nil), // 0: OpenRogueChestScRsp
	(*ItemList)(nil),            // 1: ItemList
}
var file_OpenRogueChestScRsp_proto_depIdxs = []int32{
	1, // 0: OpenRogueChestScRsp.drop_data:type_name -> ItemList
	1, // 1: OpenRogueChestScRsp.reward:type_name -> ItemList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_OpenRogueChestScRsp_proto_init() }
func file_OpenRogueChestScRsp_proto_init() {
	if File_OpenRogueChestScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_OpenRogueChestScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenRogueChestScRsp); i {
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
			RawDescriptor: file_OpenRogueChestScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OpenRogueChestScRsp_proto_goTypes,
		DependencyIndexes: file_OpenRogueChestScRsp_proto_depIdxs,
		MessageInfos:      file_OpenRogueChestScRsp_proto_msgTypes,
	}.Build()
	File_OpenRogueChestScRsp_proto = out.File
	file_OpenRogueChestScRsp_proto_rawDesc = nil
	file_OpenRogueChestScRsp_proto_goTypes = nil
	file_OpenRogueChestScRsp_proto_depIdxs = nil
}
