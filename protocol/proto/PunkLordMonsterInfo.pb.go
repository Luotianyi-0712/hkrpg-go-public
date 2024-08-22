// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PunkLordMonsterInfo.proto

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

type PunkLordMonsterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasicInfo      *PunkLordMonsterBasicInfo `protobuf:"bytes,11,opt,name=basic_info,json=basicInfo,proto3" json:"basic_info,omitempty"`
	JGBGIKPBLFP    uint32                    `protobuf:"varint,14,opt,name=JGBGIKPBLFP,proto3" json:"JGBGIKPBLFP,omitempty"`
	BattleRecord   *PunkLordBattleRecordList `protobuf:"bytes,4,opt,name=battle_record,json=battleRecord,proto3" json:"battle_record,omitempty"`
	AttackerStatus PunkLordAttackerStatus    `protobuf:"varint,2,opt,name=attacker_status,json=attackerStatus,proto3,enum=PunkLordAttackerStatus" json:"attacker_status,omitempty"`
}

func (x *PunkLordMonsterInfo) Reset() {
	*x = PunkLordMonsterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PunkLordMonsterInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PunkLordMonsterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PunkLordMonsterInfo) ProtoMessage() {}

func (x *PunkLordMonsterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PunkLordMonsterInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PunkLordMonsterInfo.ProtoReflect.Descriptor instead.
func (*PunkLordMonsterInfo) Descriptor() ([]byte, []int) {
	return file_PunkLordMonsterInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PunkLordMonsterInfo) GetBasicInfo() *PunkLordMonsterBasicInfo {
	if x != nil {
		return x.BasicInfo
	}
	return nil
}

func (x *PunkLordMonsterInfo) GetJGBGIKPBLFP() uint32 {
	if x != nil {
		return x.JGBGIKPBLFP
	}
	return 0
}

func (x *PunkLordMonsterInfo) GetBattleRecord() *PunkLordBattleRecordList {
	if x != nil {
		return x.BattleRecord
	}
	return nil
}

func (x *PunkLordMonsterInfo) GetAttackerStatus() PunkLordAttackerStatus {
	if x != nil {
		return x.AttackerStatus
	}
	return PunkLordAttackerStatus_PUNK_LORD_ATTACKER_STATUS_NONE
}

var File_PunkLordMonsterInfo_proto protoreflect.FileDescriptor

var file_PunkLordMonsterInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x50, 0x75, 0x6e,
	0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x50, 0x75, 0x6e,
	0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x50, 0x75, 0x6e,
	0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x13, 0x50, 0x75,
	0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x38, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64,
	0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x62, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4a,
	0x47, 0x42, 0x47, 0x49, 0x4b, 0x50, 0x42, 0x4c, 0x46, 0x50, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4a, 0x47, 0x42, 0x47, 0x49, 0x4b, 0x50, 0x42, 0x4c, 0x46, 0x50, 0x12, 0x3e, 0x0a,
	0x0d, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x0c, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x40, 0x0a,
	0x0f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72,
	0x64, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x0e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_PunkLordMonsterInfo_proto_rawDescOnce sync.Once
	file_PunkLordMonsterInfo_proto_rawDescData = file_PunkLordMonsterInfo_proto_rawDesc
)

func file_PunkLordMonsterInfo_proto_rawDescGZIP() []byte {
	file_PunkLordMonsterInfo_proto_rawDescOnce.Do(func() {
		file_PunkLordMonsterInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PunkLordMonsterInfo_proto_rawDescData)
	})
	return file_PunkLordMonsterInfo_proto_rawDescData
}

var file_PunkLordMonsterInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PunkLordMonsterInfo_proto_goTypes = []interface{}{
	(*PunkLordMonsterInfo)(nil),      // 0: PunkLordMonsterInfo
	(*PunkLordMonsterBasicInfo)(nil), // 1: PunkLordMonsterBasicInfo
	(*PunkLordBattleRecordList)(nil), // 2: PunkLordBattleRecordList
	(PunkLordAttackerStatus)(0),      // 3: PunkLordAttackerStatus
}
var file_PunkLordMonsterInfo_proto_depIdxs = []int32{
	1, // 0: PunkLordMonsterInfo.basic_info:type_name -> PunkLordMonsterBasicInfo
	2, // 1: PunkLordMonsterInfo.battle_record:type_name -> PunkLordBattleRecordList
	3, // 2: PunkLordMonsterInfo.attacker_status:type_name -> PunkLordAttackerStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_PunkLordMonsterInfo_proto_init() }
func file_PunkLordMonsterInfo_proto_init() {
	if File_PunkLordMonsterInfo_proto != nil {
		return
	}
	file_PunkLordMonsterBasicInfo_proto_init()
	file_PunkLordBattleRecordList_proto_init()
	file_PunkLordAttackerStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PunkLordMonsterInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PunkLordMonsterInfo); i {
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
			RawDescriptor: file_PunkLordMonsterInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PunkLordMonsterInfo_proto_goTypes,
		DependencyIndexes: file_PunkLordMonsterInfo_proto_depIdxs,
		MessageInfos:      file_PunkLordMonsterInfo_proto_msgTypes,
	}.Build()
	File_PunkLordMonsterInfo_proto = out.File
	file_PunkLordMonsterInfo_proto_rawDesc = nil
	file_PunkLordMonsterInfo_proto_goTypes = nil
	file_PunkLordMonsterInfo_proto_depIdxs = nil
}
