// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EvolveBuildFinishScNotify.proto

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

type EvolveBuildFinishScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wave             uint32                `protobuf:"varint,9,opt,name=wave,proto3" json:"wave,omitempty"`
	IsLose           bool                  `protobuf:"varint,3,opt,name=is_lose,json=isLose,proto3" json:"is_lose,omitempty"`
	ScoreId          uint32                `protobuf:"varint,11,opt,name=score_id,json=scoreId,proto3" json:"score_id,omitempty"`
	LevelInfo        *EvolveBuildLevelInfo `protobuf:"bytes,7,opt,name=level_info,json=levelInfo,proto3" json:"level_info,omitempty"`
	LevelId          uint32                `protobuf:"varint,15,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
	Coin             uint32                `protobuf:"varint,8,opt,name=coin,proto3" json:"coin,omitempty"`
	BattleResultType GADPGIMLECD           `protobuf:"varint,10,opt,name=battle_result_type,json=battleResultType,proto3,enum=GADPGIMLECD" json:"battle_result_type,omitempty"`
	Exp              uint32                `protobuf:"varint,12,opt,name=exp,proto3" json:"exp,omitempty"`
	CurPeriodType    uint32                `protobuf:"varint,5,opt,name=cur_period_type,json=curPeriodType,proto3" json:"cur_period_type,omitempty"`
}

func (x *EvolveBuildFinishScNotify) Reset() {
	*x = EvolveBuildFinishScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvolveBuildFinishScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvolveBuildFinishScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvolveBuildFinishScNotify) ProtoMessage() {}

func (x *EvolveBuildFinishScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvolveBuildFinishScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvolveBuildFinishScNotify.ProtoReflect.Descriptor instead.
func (*EvolveBuildFinishScNotify) Descriptor() ([]byte, []int) {
	return file_EvolveBuildFinishScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvolveBuildFinishScNotify) GetWave() uint32 {
	if x != nil {
		return x.Wave
	}
	return 0
}

func (x *EvolveBuildFinishScNotify) GetIsLose() bool {
	if x != nil {
		return x.IsLose
	}
	return false
}

func (x *EvolveBuildFinishScNotify) GetScoreId() uint32 {
	if x != nil {
		return x.ScoreId
	}
	return 0
}

func (x *EvolveBuildFinishScNotify) GetLevelInfo() *EvolveBuildLevelInfo {
	if x != nil {
		return x.LevelInfo
	}
	return nil
}

func (x *EvolveBuildFinishScNotify) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *EvolveBuildFinishScNotify) GetCoin() uint32 {
	if x != nil {
		return x.Coin
	}
	return 0
}

func (x *EvolveBuildFinishScNotify) GetBattleResultType() GADPGIMLECD {
	if x != nil {
		return x.BattleResultType
	}
	return GADPGIMLECD_EVOLVE_BATTLE_RESULT_NONE
}

func (x *EvolveBuildFinishScNotify) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *EvolveBuildFinishScNotify) GetCurPeriodType() uint32 {
	if x != nil {
		return x.CurPeriodType
	}
	return 0
}

var File_EvolveBuildFinishScNotify_proto protoreflect.FileDescriptor

var file_EvolveBuildFinishScNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47,
	0x41, 0x44, 0x50, 0x47, 0x49, 0x4d, 0x4c, 0x45, 0x43, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbe, 0x02, 0x0a, 0x19, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x77, 0x61, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x77, 0x61,
	0x76, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4c, 0x6f, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x45, 0x76, 0x6f,
	0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x3a, 0x0a, 0x12, 0x62,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x47, 0x41, 0x44, 0x50, 0x47, 0x49,
	0x4d, 0x4c, 0x45, 0x43, 0x44, 0x52, 0x10, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x75, 0x72,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_EvolveBuildFinishScNotify_proto_rawDescOnce sync.Once
	file_EvolveBuildFinishScNotify_proto_rawDescData = file_EvolveBuildFinishScNotify_proto_rawDesc
)

func file_EvolveBuildFinishScNotify_proto_rawDescGZIP() []byte {
	file_EvolveBuildFinishScNotify_proto_rawDescOnce.Do(func() {
		file_EvolveBuildFinishScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvolveBuildFinishScNotify_proto_rawDescData)
	})
	return file_EvolveBuildFinishScNotify_proto_rawDescData
}

var file_EvolveBuildFinishScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvolveBuildFinishScNotify_proto_goTypes = []interface{}{
	(*EvolveBuildFinishScNotify)(nil), // 0: EvolveBuildFinishScNotify
	(*EvolveBuildLevelInfo)(nil),      // 1: EvolveBuildLevelInfo
	(GADPGIMLECD)(0),                  // 2: GADPGIMLECD
}
var file_EvolveBuildFinishScNotify_proto_depIdxs = []int32{
	1, // 0: EvolveBuildFinishScNotify.level_info:type_name -> EvolveBuildLevelInfo
	2, // 1: EvolveBuildFinishScNotify.battle_result_type:type_name -> GADPGIMLECD
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvolveBuildFinishScNotify_proto_init() }
func file_EvolveBuildFinishScNotify_proto_init() {
	if File_EvolveBuildFinishScNotify_proto != nil {
		return
	}
	file_EvolveBuildLevelInfo_proto_init()
	file_GADPGIMLECD_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvolveBuildFinishScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvolveBuildFinishScNotify); i {
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
			RawDescriptor: file_EvolveBuildFinishScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvolveBuildFinishScNotify_proto_goTypes,
		DependencyIndexes: file_EvolveBuildFinishScNotify_proto_depIdxs,
		MessageInfos:      file_EvolveBuildFinishScNotify_proto_msgTypes,
	}.Build()
	File_EvolveBuildFinishScNotify_proto = out.File
	file_EvolveBuildFinishScNotify_proto_rawDesc = nil
	file_EvolveBuildFinishScNotify_proto_goTypes = nil
	file_EvolveBuildFinishScNotify_proto_depIdxs = nil
}
