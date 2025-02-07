// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneBattleInfo.proto

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

type SceneBattleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogicRandomSeed  uint32                       `protobuf:"varint,12,opt,name=logic_random_seed,json=logicRandomSeed,proto3" json:"logic_random_seed,omitempty"`
	BattleAvatarList []*BattleAvatar              `protobuf:"bytes,15,rep,name=battle_avatar_list,json=battleAvatarList,proto3" json:"battle_avatar_list,omitempty"`
	BattleEvent      []*BattleEventBattleInfo     `protobuf:"bytes,1965,rep,name=battle_event,json=battleEvent,proto3" json:"battle_event,omitempty"`
	BuffList         []*BattleBuff                `protobuf:"bytes,3,rep,name=buff_list,json=buffList,proto3" json:"buff_list,omitempty"`
	HEAMIJGFDMO      *EvolveBuildBattleInfo       `protobuf:"bytes,763,opt,name=HEAMIJGFDMO,proto3" json:"HEAMIJGFDMO,omitempty"`
	StageId          uint32                       `protobuf:"varint,4,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	FNLHAHFIGNC      *OFHNBLOGEME                 `protobuf:"bytes,854,opt,name=FNLHAHFIGNC,proto3" json:"FNLHAHFIGNC,omitempty"`
	HKOOBMMLGME      *AELAFNKGADP                 `protobuf:"bytes,63,opt,name=HKOOBMMLGME,proto3" json:"HKOOBMMLGME,omitempty"`
	BattleTargetInfo map[uint32]*BattleTargetList `protobuf:"bytes,149,rep,name=battle_target_info,json=battleTargetInfo,proto3" json:"battle_target_info,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MOJLNCNDIOB      bool                         `protobuf:"varint,8,opt,name=MOJLNCNDIOB,proto3" json:"MOJLNCNDIOB,omitempty"`
	RoundsLimit      uint32                       `protobuf:"varint,2,opt,name=rounds_limit,json=roundsLimit,proto3" json:"rounds_limit,omitempty"`
	MonsterWaveList  []*SceneMonsterWave          `protobuf:"bytes,1,rep,name=monster_wave_list,json=monsterWaveList,proto3" json:"monster_wave_list,omitempty"`
	WorldLevel       uint32                       `protobuf:"varint,6,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	BOJHPNCAKOP      uint32                       `protobuf:"varint,11,opt,name=BOJHPNCAKOP,proto3" json:"BOJHPNCAKOP,omitempty"`
	BattleId         uint32                       `protobuf:"varint,14,opt,name=battle_id,json=battleId,proto3" json:"battle_id,omitempty"`
}

func (x *SceneBattleInfo) Reset() {
	*x = SceneBattleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneBattleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneBattleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneBattleInfo) ProtoMessage() {}

func (x *SceneBattleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneBattleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneBattleInfo.ProtoReflect.Descriptor instead.
func (*SceneBattleInfo) Descriptor() ([]byte, []int) {
	return file_SceneBattleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneBattleInfo) GetLogicRandomSeed() uint32 {
	if x != nil {
		return x.LogicRandomSeed
	}
	return 0
}

func (x *SceneBattleInfo) GetBattleAvatarList() []*BattleAvatar {
	if x != nil {
		return x.BattleAvatarList
	}
	return nil
}

func (x *SceneBattleInfo) GetBattleEvent() []*BattleEventBattleInfo {
	if x != nil {
		return x.BattleEvent
	}
	return nil
}

func (x *SceneBattleInfo) GetBuffList() []*BattleBuff {
	if x != nil {
		return x.BuffList
	}
	return nil
}

func (x *SceneBattleInfo) GetHEAMIJGFDMO() *EvolveBuildBattleInfo {
	if x != nil {
		return x.HEAMIJGFDMO
	}
	return nil
}

func (x *SceneBattleInfo) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *SceneBattleInfo) GetFNLHAHFIGNC() *OFHNBLOGEME {
	if x != nil {
		return x.FNLHAHFIGNC
	}
	return nil
}

func (x *SceneBattleInfo) GetHKOOBMMLGME() *AELAFNKGADP {
	if x != nil {
		return x.HKOOBMMLGME
	}
	return nil
}

func (x *SceneBattleInfo) GetBattleTargetInfo() map[uint32]*BattleTargetList {
	if x != nil {
		return x.BattleTargetInfo
	}
	return nil
}

func (x *SceneBattleInfo) GetMOJLNCNDIOB() bool {
	if x != nil {
		return x.MOJLNCNDIOB
	}
	return false
}

func (x *SceneBattleInfo) GetRoundsLimit() uint32 {
	if x != nil {
		return x.RoundsLimit
	}
	return 0
}

func (x *SceneBattleInfo) GetMonsterWaveList() []*SceneMonsterWave {
	if x != nil {
		return x.MonsterWaveList
	}
	return nil
}

func (x *SceneBattleInfo) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *SceneBattleInfo) GetBOJHPNCAKOP() uint32 {
	if x != nil {
		return x.BOJHPNCAKOP
	}
	return 0
}

func (x *SceneBattleInfo) GetBattleId() uint32 {
	if x != nil {
		return x.BattleId
	}
	return 0
}

var File_SceneBattleInfo_proto protoreflect.FileDescriptor

var file_SceneBattleInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x57, 0x61, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41,
	0x45, 0x4c, 0x41, 0x46, 0x4e, 0x4b, 0x47, 0x41, 0x44, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x46, 0x48, 0x4e, 0x42, 0x4c, 0x4f, 0x47, 0x45, 0x4d, 0x45, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x06,
	0x0a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x65, 0x65, 0x64, 0x12, 0x3b, 0x0a,
	0x12, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x10, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x62, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0xad, 0x0f, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x42, 0x75, 0x66, 0x66, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x39, 0x0a, 0x0b, 0x48, 0x45, 0x41, 0x4d, 0x49, 0x4a, 0x47, 0x46, 0x44, 0x4d, 0x4f, 0x18,
	0xfb, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b,
	0x48, 0x45, 0x41, 0x4d, 0x49, 0x4a, 0x47, 0x46, 0x44, 0x4d, 0x4f, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0b, 0x46, 0x4e, 0x4c, 0x48, 0x41, 0x48,
	0x46, 0x49, 0x47, 0x4e, 0x43, 0x18, 0xd6, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f,
	0x46, 0x48, 0x4e, 0x42, 0x4c, 0x4f, 0x47, 0x45, 0x4d, 0x45, 0x52, 0x0b, 0x46, 0x4e, 0x4c, 0x48,
	0x41, 0x48, 0x46, 0x49, 0x47, 0x4e, 0x43, 0x12, 0x2e, 0x0a, 0x0b, 0x48, 0x4b, 0x4f, 0x4f, 0x42,
	0x4d, 0x4d, 0x4c, 0x47, 0x4d, 0x45, 0x18, 0x3f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41,
	0x45, 0x4c, 0x41, 0x46, 0x4e, 0x4b, 0x47, 0x41, 0x44, 0x50, 0x52, 0x0b, 0x48, 0x4b, 0x4f, 0x4f,
	0x42, 0x4d, 0x4d, 0x4c, 0x47, 0x4d, 0x45, 0x12, 0x55, 0x0a, 0x12, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x95, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x62, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x4d, 0x4f, 0x4a, 0x4c, 0x4e, 0x43, 0x4e, 0x44, 0x49, 0x4f, 0x42, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x4d, 0x4f, 0x4a, 0x4c, 0x4e, 0x43, 0x4e, 0x44, 0x49, 0x4f, 0x42,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x3d, 0x0a, 0x11, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x77,
	0x61, 0x76, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x57, 0x61, 0x76,
	0x65, 0x52, 0x0f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x57, 0x61, 0x76, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x4f, 0x4a, 0x48, 0x50, 0x4e, 0x43, 0x41, 0x4b,
	0x4f, 0x50, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x4f, 0x4a, 0x48, 0x50, 0x4e,
	0x43, 0x41, 0x4b, 0x4f, 0x50, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x64, 0x1a, 0x56, 0x0a, 0x15, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_SceneBattleInfo_proto_rawDescOnce sync.Once
	file_SceneBattleInfo_proto_rawDescData = file_SceneBattleInfo_proto_rawDesc
)

func file_SceneBattleInfo_proto_rawDescGZIP() []byte {
	file_SceneBattleInfo_proto_rawDescOnce.Do(func() {
		file_SceneBattleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneBattleInfo_proto_rawDescData)
	})
	return file_SceneBattleInfo_proto_rawDescData
}

var file_SceneBattleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SceneBattleInfo_proto_goTypes = []interface{}{
	(*SceneBattleInfo)(nil),       // 0: SceneBattleInfo
	nil,                           // 1: SceneBattleInfo.BattleTargetInfoEntry
	(*BattleAvatar)(nil),          // 2: BattleAvatar
	(*BattleEventBattleInfo)(nil), // 3: BattleEventBattleInfo
	(*BattleBuff)(nil),            // 4: BattleBuff
	(*EvolveBuildBattleInfo)(nil), // 5: EvolveBuildBattleInfo
	(*OFHNBLOGEME)(nil),           // 6: OFHNBLOGEME
	(*AELAFNKGADP)(nil),           // 7: AELAFNKGADP
	(*SceneMonsterWave)(nil),      // 8: SceneMonsterWave
	(*BattleTargetList)(nil),      // 9: BattleTargetList
}
var file_SceneBattleInfo_proto_depIdxs = []int32{
	2, // 0: SceneBattleInfo.battle_avatar_list:type_name -> BattleAvatar
	3, // 1: SceneBattleInfo.battle_event:type_name -> BattleEventBattleInfo
	4, // 2: SceneBattleInfo.buff_list:type_name -> BattleBuff
	5, // 3: SceneBattleInfo.HEAMIJGFDMO:type_name -> EvolveBuildBattleInfo
	6, // 4: SceneBattleInfo.FNLHAHFIGNC:type_name -> OFHNBLOGEME
	7, // 5: SceneBattleInfo.HKOOBMMLGME:type_name -> AELAFNKGADP
	1, // 6: SceneBattleInfo.battle_target_info:type_name -> SceneBattleInfo.BattleTargetInfoEntry
	8, // 7: SceneBattleInfo.monster_wave_list:type_name -> SceneMonsterWave
	9, // 8: SceneBattleInfo.BattleTargetInfoEntry.value:type_name -> BattleTargetList
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_SceneBattleInfo_proto_init() }
func file_SceneBattleInfo_proto_init() {
	if File_SceneBattleInfo_proto != nil {
		return
	}
	file_BattleEventBattleInfo_proto_init()
	file_BattleAvatar_proto_init()
	file_SceneMonsterWave_proto_init()
	file_EvolveBuildBattleInfo_proto_init()
	file_AELAFNKGADP_proto_init()
	file_BattleBuff_proto_init()
	file_OFHNBLOGEME_proto_init()
	file_BattleTargetList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneBattleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneBattleInfo); i {
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
			RawDescriptor: file_SceneBattleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneBattleInfo_proto_goTypes,
		DependencyIndexes: file_SceneBattleInfo_proto_depIdxs,
		MessageInfos:      file_SceneBattleInfo_proto_msgTypes,
	}.Build()
	File_SceneBattleInfo_proto = out.File
	file_SceneBattleInfo_proto_rawDesc = nil
	file_SceneBattleInfo_proto_goTypes = nil
	file_SceneBattleInfo_proto_depIdxs = nil
}
