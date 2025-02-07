// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CurChallenge.proto

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

type CurChallenge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScoreId         uint32              `protobuf:"varint,3,opt,name=score_id,json=scoreId,proto3" json:"score_id,omitempty"`
	Status          ChallengeStatus     `protobuf:"varint,13,opt,name=status,proto3,enum=ChallengeStatus" json:"status,omitempty"`
	PlayerInfo      *ChallengeStoryInfo `protobuf:"bytes,9,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	ScoreTwo        uint32              `protobuf:"varint,12,opt,name=score_two,json=scoreTwo,proto3" json:"score_two,omitempty"`
	ExtraLineupType ExtraLineupType     `protobuf:"varint,7,opt,name=extra_lineup_type,json=extraLineupType,proto3,enum=ExtraLineupType" json:"extra_lineup_type,omitempty"`
	ChallengeId     uint32              `protobuf:"varint,5,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	RoundCount      uint32              `protobuf:"varint,14,opt,name=round_count,json=roundCount,proto3" json:"round_count,omitempty"`
	KillMonsterList []*KillMonster      `protobuf:"bytes,15,rep,name=kill_monster_list,json=killMonsterList,proto3" json:"kill_monster_list,omitempty"`
	DeadAvatarNum   uint32              `protobuf:"varint,11,opt,name=dead_avatar_num,json=deadAvatarNum,proto3" json:"dead_avatar_num,omitempty"`
}

func (x *CurChallenge) Reset() {
	*x = CurChallenge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CurChallenge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurChallenge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurChallenge) ProtoMessage() {}

func (x *CurChallenge) ProtoReflect() protoreflect.Message {
	mi := &file_CurChallenge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurChallenge.ProtoReflect.Descriptor instead.
func (*CurChallenge) Descriptor() ([]byte, []int) {
	return file_CurChallenge_proto_rawDescGZIP(), []int{0}
}

func (x *CurChallenge) GetScoreId() uint32 {
	if x != nil {
		return x.ScoreId
	}
	return 0
}

func (x *CurChallenge) GetStatus() ChallengeStatus {
	if x != nil {
		return x.Status
	}
	return ChallengeStatus_CHALLENGE_UNKNOWN
}

func (x *CurChallenge) GetPlayerInfo() *ChallengeStoryInfo {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

func (x *CurChallenge) GetScoreTwo() uint32 {
	if x != nil {
		return x.ScoreTwo
	}
	return 0
}

func (x *CurChallenge) GetExtraLineupType() ExtraLineupType {
	if x != nil {
		return x.ExtraLineupType
	}
	return ExtraLineupType_LINEUP_NONE
}

func (x *CurChallenge) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *CurChallenge) GetRoundCount() uint32 {
	if x != nil {
		return x.RoundCount
	}
	return 0
}

func (x *CurChallenge) GetKillMonsterList() []*KillMonster {
	if x != nil {
		return x.KillMonsterList
	}
	return nil
}

func (x *CurChallenge) GetDeadAvatarNum() uint32 {
	if x != nil {
		return x.DeadAvatarNum
	}
	return 0
}

var File_CurChallenge_proto protoreflect.FileDescriptor

var file_CurChallenge_proto_rawDesc = []byte{
	0x0a, 0x12, 0x43, 0x75, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x69, 0x6c,
	0x6c, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8a, 0x03, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x77, 0x6f, 0x12, 0x3c, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x11, 0x6b, 0x69,
	0x6c, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x4d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x0f, 0x6b, 0x69, 0x6c, 0x6c, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x64,
	0x65, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4e, 0x75, 0x6d, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CurChallenge_proto_rawDescOnce sync.Once
	file_CurChallenge_proto_rawDescData = file_CurChallenge_proto_rawDesc
)

func file_CurChallenge_proto_rawDescGZIP() []byte {
	file_CurChallenge_proto_rawDescOnce.Do(func() {
		file_CurChallenge_proto_rawDescData = protoimpl.X.CompressGZIP(file_CurChallenge_proto_rawDescData)
	})
	return file_CurChallenge_proto_rawDescData
}

var file_CurChallenge_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CurChallenge_proto_goTypes = []interface{}{
	(*CurChallenge)(nil),       // 0: CurChallenge
	(ChallengeStatus)(0),       // 1: ChallengeStatus
	(*ChallengeStoryInfo)(nil), // 2: ChallengeStoryInfo
	(ExtraLineupType)(0),       // 3: ExtraLineupType
	(*KillMonster)(nil),        // 4: KillMonster
}
var file_CurChallenge_proto_depIdxs = []int32{
	1, // 0: CurChallenge.status:type_name -> ChallengeStatus
	2, // 1: CurChallenge.player_info:type_name -> ChallengeStoryInfo
	3, // 2: CurChallenge.extra_lineup_type:type_name -> ExtraLineupType
	4, // 3: CurChallenge.kill_monster_list:type_name -> KillMonster
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_CurChallenge_proto_init() }
func file_CurChallenge_proto_init() {
	if File_CurChallenge_proto != nil {
		return
	}
	file_ChallengeStatus_proto_init()
	file_KillMonster_proto_init()
	file_ExtraLineupType_proto_init()
	file_ChallengeStoryInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CurChallenge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurChallenge); i {
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
			RawDescriptor: file_CurChallenge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CurChallenge_proto_goTypes,
		DependencyIndexes: file_CurChallenge_proto_depIdxs,
		MessageInfos:      file_CurChallenge_proto_msgTypes,
	}.Build()
	File_CurChallenge_proto = out.File
	file_CurChallenge_proto_rawDesc = nil
	file_CurChallenge_proto_goTypes = nil
	file_CurChallenge_proto_depIdxs = nil
}
