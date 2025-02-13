// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: StartStarFightLevelScRsp.proto

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

type StartStarFightLevelScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32           `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	IHKPKPJKKBI uint32           `protobuf:"varint,7,opt,name=IHKPKPJKKBI,proto3" json:"IHKPKPJKKBI,omitempty"`
	GroupId     uint32           `protobuf:"varint,14,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	BattleInfo  *SceneBattleInfo `protobuf:"bytes,6,opt,name=battle_info,json=battleInfo,proto3" json:"battle_info,omitempty"`
}

func (x *StartStarFightLevelScRsp) Reset() {
	*x = StartStarFightLevelScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StartStarFightLevelScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartStarFightLevelScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartStarFightLevelScRsp) ProtoMessage() {}

func (x *StartStarFightLevelScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_StartStarFightLevelScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartStarFightLevelScRsp.ProtoReflect.Descriptor instead.
func (*StartStarFightLevelScRsp) Descriptor() ([]byte, []int) {
	return file_StartStarFightLevelScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *StartStarFightLevelScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *StartStarFightLevelScRsp) GetIHKPKPJKKBI() uint32 {
	if x != nil {
		return x.IHKPKPJKKBI
	}
	return 0
}

func (x *StartStarFightLevelScRsp) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *StartStarFightLevelScRsp) GetBattleInfo() *SceneBattleInfo {
	if x != nil {
		return x.BattleInfo
	}
	return nil
}

var File_StartStarFightLevelScRsp_proto protoreflect.FileDescriptor

var file_StartStarFightLevelScRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x72, 0x46, 0x69, 0x67, 0x68, 0x74,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x53, 0x74, 0x61, 0x72, 0x46, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x49, 0x48, 0x4b, 0x50, 0x4b, 0x50, 0x4a, 0x4b, 0x4b, 0x42, 0x49, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x48, 0x4b, 0x50, 0x4b, 0x50, 0x4a, 0x4b, 0x4b, 0x42, 0x49,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x62,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StartStarFightLevelScRsp_proto_rawDescOnce sync.Once
	file_StartStarFightLevelScRsp_proto_rawDescData = file_StartStarFightLevelScRsp_proto_rawDesc
)

func file_StartStarFightLevelScRsp_proto_rawDescGZIP() []byte {
	file_StartStarFightLevelScRsp_proto_rawDescOnce.Do(func() {
		file_StartStarFightLevelScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_StartStarFightLevelScRsp_proto_rawDescData)
	})
	return file_StartStarFightLevelScRsp_proto_rawDescData
}

var file_StartStarFightLevelScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StartStarFightLevelScRsp_proto_goTypes = []interface{}{
	(*StartStarFightLevelScRsp)(nil), // 0: StartStarFightLevelScRsp
	(*SceneBattleInfo)(nil),          // 1: SceneBattleInfo
}
var file_StartStarFightLevelScRsp_proto_depIdxs = []int32{
	1, // 0: StartStarFightLevelScRsp.battle_info:type_name -> SceneBattleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_StartStarFightLevelScRsp_proto_init() }
func file_StartStarFightLevelScRsp_proto_init() {
	if File_StartStarFightLevelScRsp_proto != nil {
		return
	}
	file_SceneBattleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_StartStarFightLevelScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartStarFightLevelScRsp); i {
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
			RawDescriptor: file_StartStarFightLevelScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StartStarFightLevelScRsp_proto_goTypes,
		DependencyIndexes: file_StartStarFightLevelScRsp_proto_depIdxs,
		MessageInfos:      file_StartStarFightLevelScRsp_proto_msgTypes,
	}.Build()
	File_StartStarFightLevelScRsp_proto = out.File
	file_StartStarFightLevelScRsp_proto_rawDesc = nil
	file_StartStarFightLevelScRsp_proto_goTypes = nil
	file_StartStarFightLevelScRsp_proto_depIdxs = nil
}
