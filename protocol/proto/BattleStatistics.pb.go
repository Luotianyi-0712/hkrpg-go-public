// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BattleStatistics.proto

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

type BattleStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalBattleTurns   uint32                       `protobuf:"varint,1,opt,name=total_battle_turns,json=totalBattleTurns,proto3" json:"total_battle_turns,omitempty"`
	TotalAutoTurns     uint32                       `protobuf:"varint,2,opt,name=total_auto_turns,json=totalAutoTurns,proto3" json:"total_auto_turns,omitempty"`
	AvatarIdList       []uint32                     `protobuf:"varint,3,rep,packed,name=avatar_id_list,json=avatarIdList,proto3" json:"avatar_id_list,omitempty"`
	UltraCnt           uint32                       `protobuf:"varint,4,opt,name=ultra_cnt,json=ultraCnt,proto3" json:"ultra_cnt,omitempty"`
	TotalDelayCumulate float64                      `protobuf:"fixed64,5,opt,name=total_delay_cumulate,json=totalDelayCumulate,proto3" json:"total_delay_cumulate,omitempty"`
	CostTime           float64                      `protobuf:"fixed64,6,opt,name=cost_time,json=costTime,proto3" json:"cost_time,omitempty"`
	BattleAvatarList   []*AvatarBattleInfo          `protobuf:"bytes,7,rep,name=battle_avatar_list,json=battleAvatarList,proto3" json:"battle_avatar_list,omitempty"`
	MonsterList        []*MonsterBattleInfo         `protobuf:"bytes,8,rep,name=monster_list,json=monsterList,proto3" json:"monster_list,omitempty"`
	RoundCnt           uint32                       `protobuf:"varint,9,opt,name=round_cnt,json=roundCnt,proto3" json:"round_cnt,omitempty"`
	CocoonDeadWave     uint32                       `protobuf:"varint,10,opt,name=cocoon_dead_wave,json=cocoonDeadWave,proto3" json:"cocoon_dead_wave,omitempty"`
	AvatarBattleTurns  uint32                       `protobuf:"varint,11,opt,name=avatar_battle_turns,json=avatarBattleTurns,proto3" json:"avatar_battle_turns,omitempty"`
	MonsterBattleTurns uint32                       `protobuf:"varint,12,opt,name=monster_battle_turns,json=monsterBattleTurns,proto3" json:"monster_battle_turns,omitempty"`
	CustomValues       map[string]float32           `protobuf:"bytes,13,rep,name=custom_values,json=customValues,proto3" json:"custom_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	ChallengeScore     uint32                       `protobuf:"varint,14,opt,name=challenge_score,json=challengeScore,proto3" json:"challenge_score,omitempty"`
	PJOECEPBPOJ        []*BattleEventBattleInfo     `protobuf:"bytes,16,rep,name=PJOECEPBPOJ,proto3" json:"PJOECEPBPOJ,omitempty"`
	EndReason          BattleEndReason              `protobuf:"varint,19,opt,name=end_reason,json=endReason,proto3,enum=BattleEndReason" json:"end_reason,omitempty"`
	OIPOMFPKFIE        []*LGFELGLKHJB               `protobuf:"bytes,21,rep,name=OIPOMFPKFIE,proto3" json:"OIPOMFPKFIE,omitempty"`
	DENNDAGNJNN        []int32                      `protobuf:"varint,22,rep,packed,name=DENNDAGNJNN,proto3" json:"DENNDAGNJNN,omitempty"`
	CGOHELIDBBC        []*EDMLINEHPPE               `protobuf:"bytes,23,rep,name=CGOHELIDBBC,proto3" json:"CGOHELIDBBC,omitempty"`
	ENCPJPEHIDD        []*IGOJKPCDFNN               `protobuf:"bytes,26,rep,name=ENCPJPEHIDD,proto3" json:"ENCPJPEHIDD,omitempty"`
	OGEGBEKLPAB        []*CKLPMMEKCDF               `protobuf:"bytes,27,rep,name=OGEGBEKLPAB,proto3" json:"OGEGBEKLPAB,omitempty"`
	BattleTargetInfo   map[uint32]*BattleTargetList `protobuf:"bytes,28,rep,name=battle_target_info,json=battleTargetInfo,proto3" json:"battle_target_info,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ENKHGICLCFO        []*MFBFMKMABAO               `protobuf:"bytes,29,rep,name=ENKHGICLCFO,proto3" json:"ENKHGICLCFO,omitempty"`
	HEAMIJGFDMO        *EvolveBuildBattleInfo       `protobuf:"bytes,30,opt,name=HEAMIJGFDMO,proto3" json:"HEAMIJGFDMO,omitempty"`
	FDBABLGMGKN        *HHJEAHMHGFK                 `protobuf:"bytes,31,opt,name=FDBABLGMGKN,proto3" json:"FDBABLGMGKN,omitempty"`
	DBHGJCODLBK        bool                         `protobuf:"varint,32,opt,name=DBHGJCODLBK,proto3" json:"DBHGJCODLBK,omitempty"`
	BJHPKLNLGLG        []*DKODPCDGOOL               `protobuf:"bytes,33,rep,name=BJHPKLNLGLG,proto3" json:"BJHPKLNLGLG,omitempty"`
	LCBLKBIMDHL        []*CLJLLOFIBML               `protobuf:"bytes,34,rep,name=LCBLKBIMDHL,proto3" json:"LCBLKBIMDHL,omitempty"`
	DAHDDICCOGD        uint32                       `protobuf:"varint,35,opt,name=DAHDDICCOGD,proto3" json:"DAHDDICCOGD,omitempty"`
}

func (x *BattleStatistics) Reset() {
	*x = BattleStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BattleStatistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleStatistics) ProtoMessage() {}

func (x *BattleStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_BattleStatistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleStatistics.ProtoReflect.Descriptor instead.
func (*BattleStatistics) Descriptor() ([]byte, []int) {
	return file_BattleStatistics_proto_rawDescGZIP(), []int{0}
}

func (x *BattleStatistics) GetTotalBattleTurns() uint32 {
	if x != nil {
		return x.TotalBattleTurns
	}
	return 0
}

func (x *BattleStatistics) GetTotalAutoTurns() uint32 {
	if x != nil {
		return x.TotalAutoTurns
	}
	return 0
}

func (x *BattleStatistics) GetAvatarIdList() []uint32 {
	if x != nil {
		return x.AvatarIdList
	}
	return nil
}

func (x *BattleStatistics) GetUltraCnt() uint32 {
	if x != nil {
		return x.UltraCnt
	}
	return 0
}

func (x *BattleStatistics) GetTotalDelayCumulate() float64 {
	if x != nil {
		return x.TotalDelayCumulate
	}
	return 0
}

func (x *BattleStatistics) GetCostTime() float64 {
	if x != nil {
		return x.CostTime
	}
	return 0
}

func (x *BattleStatistics) GetBattleAvatarList() []*AvatarBattleInfo {
	if x != nil {
		return x.BattleAvatarList
	}
	return nil
}

func (x *BattleStatistics) GetMonsterList() []*MonsterBattleInfo {
	if x != nil {
		return x.MonsterList
	}
	return nil
}

func (x *BattleStatistics) GetRoundCnt() uint32 {
	if x != nil {
		return x.RoundCnt
	}
	return 0
}

func (x *BattleStatistics) GetCocoonDeadWave() uint32 {
	if x != nil {
		return x.CocoonDeadWave
	}
	return 0
}

func (x *BattleStatistics) GetAvatarBattleTurns() uint32 {
	if x != nil {
		return x.AvatarBattleTurns
	}
	return 0
}

func (x *BattleStatistics) GetMonsterBattleTurns() uint32 {
	if x != nil {
		return x.MonsterBattleTurns
	}
	return 0
}

func (x *BattleStatistics) GetCustomValues() map[string]float32 {
	if x != nil {
		return x.CustomValues
	}
	return nil
}

func (x *BattleStatistics) GetChallengeScore() uint32 {
	if x != nil {
		return x.ChallengeScore
	}
	return 0
}

func (x *BattleStatistics) GetPJOECEPBPOJ() []*BattleEventBattleInfo {
	if x != nil {
		return x.PJOECEPBPOJ
	}
	return nil
}

func (x *BattleStatistics) GetEndReason() BattleEndReason {
	if x != nil {
		return x.EndReason
	}
	return BattleEndReason_BATTLE_END_REASON_NONE
}

func (x *BattleStatistics) GetOIPOMFPKFIE() []*LGFELGLKHJB {
	if x != nil {
		return x.OIPOMFPKFIE
	}
	return nil
}

func (x *BattleStatistics) GetDENNDAGNJNN() []int32 {
	if x != nil {
		return x.DENNDAGNJNN
	}
	return nil
}

func (x *BattleStatistics) GetCGOHELIDBBC() []*EDMLINEHPPE {
	if x != nil {
		return x.CGOHELIDBBC
	}
	return nil
}

func (x *BattleStatistics) GetENCPJPEHIDD() []*IGOJKPCDFNN {
	if x != nil {
		return x.ENCPJPEHIDD
	}
	return nil
}

func (x *BattleStatistics) GetOGEGBEKLPAB() []*CKLPMMEKCDF {
	if x != nil {
		return x.OGEGBEKLPAB
	}
	return nil
}

func (x *BattleStatistics) GetBattleTargetInfo() map[uint32]*BattleTargetList {
	if x != nil {
		return x.BattleTargetInfo
	}
	return nil
}

func (x *BattleStatistics) GetENKHGICLCFO() []*MFBFMKMABAO {
	if x != nil {
		return x.ENKHGICLCFO
	}
	return nil
}

func (x *BattleStatistics) GetHEAMIJGFDMO() *EvolveBuildBattleInfo {
	if x != nil {
		return x.HEAMIJGFDMO
	}
	return nil
}

func (x *BattleStatistics) GetFDBABLGMGKN() *HHJEAHMHGFK {
	if x != nil {
		return x.FDBABLGMGKN
	}
	return nil
}

func (x *BattleStatistics) GetDBHGJCODLBK() bool {
	if x != nil {
		return x.DBHGJCODLBK
	}
	return false
}

func (x *BattleStatistics) GetBJHPKLNLGLG() []*DKODPCDGOOL {
	if x != nil {
		return x.BJHPKLNLGLG
	}
	return nil
}

func (x *BattleStatistics) GetLCBLKBIMDHL() []*CLJLLOFIBML {
	if x != nil {
		return x.LCBLKBIMDHL
	}
	return nil
}

func (x *BattleStatistics) GetDAHDDICCOGD() uint32 {
	if x != nil {
		return x.DAHDDICCOGD
	}
	return 0
}

var File_BattleStatistics_proto protoreflect.FileDescriptor

var file_BattleStatistics_proto_rawDesc = []byte{
	0x0a, 0x16, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x47, 0x46, 0x45, 0x4c, 0x47, 0x4c, 0x4b, 0x48,
	0x4a, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x45, 0x76, 0x6f, 0x6c, 0x76,
	0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x48, 0x4a, 0x45, 0x41, 0x48, 0x4d, 0x48,
	0x47, 0x46, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x4b, 0x4f, 0x44, 0x50,
	0x43, 0x44, 0x47, 0x4f, 0x4f, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x44, 0x4d, 0x4c, 0x49, 0x4e, 0x45, 0x48, 0x50, 0x50, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x46, 0x42, 0x46, 0x4d, 0x4b, 0x4d, 0x41,
	0x42, 0x41, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x47, 0x4f, 0x4a, 0x4b,
	0x50, 0x43, 0x44, 0x46, 0x4e, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4c,
	0x4a, 0x4c, 0x4c, 0x4f, 0x46, 0x49, 0x42, 0x4d, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4b, 0x4c, 0x50, 0x4d, 0x4d, 0x45,
	0x4b, 0x43, 0x44, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x0c, 0x0a, 0x10, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12,
	0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x74, 0x75, 0x72, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x73, 0x12, 0x28, 0x0a,
	0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x74, 0x75, 0x72, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x75,
	0x74, 0x6f, 0x54, 0x75, 0x72, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x6c, 0x74, 0x72, 0x61, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x75, 0x6c, 0x74, 0x72, 0x61, 0x43, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x12, 0x62, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0c, 0x6d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6e, 0x74, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x6f, 0x63, 0x6f, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x5f, 0x77, 0x61,
	0x76, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f, 0x63, 0x6f, 0x6f, 0x6e,
	0x44, 0x65, 0x61, 0x64, 0x57, 0x61, 0x76, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x74, 0x75, 0x72, 0x6e, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x73, 0x12, 0x48, 0x0a, 0x0d, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x38, 0x0a,
	0x0b, 0x50, 0x4a, 0x4f, 0x45, 0x43, 0x45, 0x50, 0x42, 0x50, 0x4f, 0x4a, 0x18, 0x10, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x50, 0x4a, 0x4f, 0x45,
	0x43, 0x45, 0x50, 0x42, 0x50, 0x4f, 0x4a, 0x12, 0x2f, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x09, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x49, 0x50, 0x4f,
	0x4d, 0x46, 0x50, 0x4b, 0x46, 0x49, 0x45, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4c, 0x47, 0x46, 0x45, 0x4c, 0x47, 0x4c, 0x4b, 0x48, 0x4a, 0x42, 0x52, 0x0b, 0x4f, 0x49, 0x50,
	0x4f, 0x4d, 0x46, 0x50, 0x4b, 0x46, 0x49, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x45, 0x4e, 0x4e,
	0x44, 0x41, 0x47, 0x4e, 0x4a, 0x4e, 0x4e, 0x18, 0x16, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x44,
	0x45, 0x4e, 0x4e, 0x44, 0x41, 0x47, 0x4e, 0x4a, 0x4e, 0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x43, 0x47,
	0x4f, 0x48, 0x45, 0x4c, 0x49, 0x44, 0x42, 0x42, 0x43, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x45, 0x44, 0x4d, 0x4c, 0x49, 0x4e, 0x45, 0x48, 0x50, 0x50, 0x45, 0x52, 0x0b, 0x43,
	0x47, 0x4f, 0x48, 0x45, 0x4c, 0x49, 0x44, 0x42, 0x42, 0x43, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x4e,
	0x43, 0x50, 0x4a, 0x50, 0x45, 0x48, 0x49, 0x44, 0x44, 0x18, 0x1a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x49, 0x47, 0x4f, 0x4a, 0x4b, 0x50, 0x43, 0x44, 0x46, 0x4e, 0x4e, 0x52, 0x0b, 0x45,
	0x4e, 0x43, 0x50, 0x4a, 0x50, 0x45, 0x48, 0x49, 0x44, 0x44, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x47,
	0x45, 0x47, 0x42, 0x45, 0x4b, 0x4c, 0x50, 0x41, 0x42, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x43, 0x4b, 0x4c, 0x50, 0x4d, 0x4d, 0x45, 0x4b, 0x43, 0x44, 0x46, 0x52, 0x0b, 0x4f,
	0x47, 0x45, 0x47, 0x42, 0x45, 0x4b, 0x4c, 0x50, 0x41, 0x42, 0x12, 0x55, 0x0a, 0x12, 0x62, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x1c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x10, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x4e, 0x4b, 0x48, 0x47, 0x49, 0x43, 0x4c, 0x43, 0x46, 0x4f,
	0x18, 0x1d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x46, 0x42, 0x46, 0x4d, 0x4b, 0x4d,
	0x41, 0x42, 0x41, 0x4f, 0x52, 0x0b, 0x45, 0x4e, 0x4b, 0x48, 0x47, 0x49, 0x43, 0x4c, 0x43, 0x46,
	0x4f, 0x12, 0x38, 0x0a, 0x0b, 0x48, 0x45, 0x41, 0x4d, 0x49, 0x4a, 0x47, 0x46, 0x44, 0x4d, 0x4f,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b,
	0x48, 0x45, 0x41, 0x4d, 0x49, 0x4a, 0x47, 0x46, 0x44, 0x4d, 0x4f, 0x12, 0x2e, 0x0a, 0x0b, 0x46,
	0x44, 0x42, 0x41, 0x42, 0x4c, 0x47, 0x4d, 0x47, 0x4b, 0x4e, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x48, 0x48, 0x4a, 0x45, 0x41, 0x48, 0x4d, 0x48, 0x47, 0x46, 0x4b, 0x52, 0x0b,
	0x46, 0x44, 0x42, 0x41, 0x42, 0x4c, 0x47, 0x4d, 0x47, 0x4b, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x42, 0x48, 0x47, 0x4a, 0x43, 0x4f, 0x44, 0x4c, 0x42, 0x4b, 0x18, 0x20, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x44, 0x42, 0x48, 0x47, 0x4a, 0x43, 0x4f, 0x44, 0x4c, 0x42, 0x4b, 0x12, 0x2e, 0x0a,
	0x0b, 0x42, 0x4a, 0x48, 0x50, 0x4b, 0x4c, 0x4e, 0x4c, 0x47, 0x4c, 0x47, 0x18, 0x21, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4b, 0x4f, 0x44, 0x50, 0x43, 0x44, 0x47, 0x4f, 0x4f, 0x4c,
	0x52, 0x0b, 0x42, 0x4a, 0x48, 0x50, 0x4b, 0x4c, 0x4e, 0x4c, 0x47, 0x4c, 0x47, 0x12, 0x2e, 0x0a,
	0x0b, 0x4c, 0x43, 0x42, 0x4c, 0x4b, 0x42, 0x49, 0x4d, 0x44, 0x48, 0x4c, 0x18, 0x22, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x4c, 0x4a, 0x4c, 0x4c, 0x4f, 0x46, 0x49, 0x42, 0x4d, 0x4c,
	0x52, 0x0b, 0x4c, 0x43, 0x42, 0x4c, 0x4b, 0x42, 0x49, 0x4d, 0x44, 0x48, 0x4c, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x41, 0x48, 0x44, 0x44, 0x49, 0x43, 0x43, 0x4f, 0x47, 0x44, 0x18, 0x23, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x44, 0x41, 0x48, 0x44, 0x44, 0x49, 0x43, 0x43, 0x4f, 0x47, 0x44, 0x1a,
	0x3f, 0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x56, 0x0a, 0x15, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattleStatistics_proto_rawDescOnce sync.Once
	file_BattleStatistics_proto_rawDescData = file_BattleStatistics_proto_rawDesc
)

func file_BattleStatistics_proto_rawDescGZIP() []byte {
	file_BattleStatistics_proto_rawDescOnce.Do(func() {
		file_BattleStatistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattleStatistics_proto_rawDescData)
	})
	return file_BattleStatistics_proto_rawDescData
}

var file_BattleStatistics_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_BattleStatistics_proto_goTypes = []interface{}{
	(*BattleStatistics)(nil),      // 0: BattleStatistics
	nil,                           // 1: BattleStatistics.CustomValuesEntry
	nil,                           // 2: BattleStatistics.BattleTargetInfoEntry
	(*AvatarBattleInfo)(nil),      // 3: AvatarBattleInfo
	(*MonsterBattleInfo)(nil),     // 4: MonsterBattleInfo
	(*BattleEventBattleInfo)(nil), // 5: BattleEventBattleInfo
	(BattleEndReason)(0),          // 6: BattleEndReason
	(*LGFELGLKHJB)(nil),           // 7: LGFELGLKHJB
	(*EDMLINEHPPE)(nil),           // 8: EDMLINEHPPE
	(*IGOJKPCDFNN)(nil),           // 9: IGOJKPCDFNN
	(*CKLPMMEKCDF)(nil),           // 10: CKLPMMEKCDF
	(*MFBFMKMABAO)(nil),           // 11: MFBFMKMABAO
	(*EvolveBuildBattleInfo)(nil), // 12: EvolveBuildBattleInfo
	(*HHJEAHMHGFK)(nil),           // 13: HHJEAHMHGFK
	(*DKODPCDGOOL)(nil),           // 14: DKODPCDGOOL
	(*CLJLLOFIBML)(nil),           // 15: CLJLLOFIBML
	(*BattleTargetList)(nil),      // 16: BattleTargetList
}
var file_BattleStatistics_proto_depIdxs = []int32{
	3,  // 0: BattleStatistics.battle_avatar_list:type_name -> AvatarBattleInfo
	4,  // 1: BattleStatistics.monster_list:type_name -> MonsterBattleInfo
	1,  // 2: BattleStatistics.custom_values:type_name -> BattleStatistics.CustomValuesEntry
	5,  // 3: BattleStatistics.PJOECEPBPOJ:type_name -> BattleEventBattleInfo
	6,  // 4: BattleStatistics.end_reason:type_name -> BattleEndReason
	7,  // 5: BattleStatistics.OIPOMFPKFIE:type_name -> LGFELGLKHJB
	8,  // 6: BattleStatistics.CGOHELIDBBC:type_name -> EDMLINEHPPE
	9,  // 7: BattleStatistics.ENCPJPEHIDD:type_name -> IGOJKPCDFNN
	10, // 8: BattleStatistics.OGEGBEKLPAB:type_name -> CKLPMMEKCDF
	2,  // 9: BattleStatistics.battle_target_info:type_name -> BattleStatistics.BattleTargetInfoEntry
	11, // 10: BattleStatistics.ENKHGICLCFO:type_name -> MFBFMKMABAO
	12, // 11: BattleStatistics.HEAMIJGFDMO:type_name -> EvolveBuildBattleInfo
	13, // 12: BattleStatistics.FDBABLGMGKN:type_name -> HHJEAHMHGFK
	14, // 13: BattleStatistics.BJHPKLNLGLG:type_name -> DKODPCDGOOL
	15, // 14: BattleStatistics.LCBLKBIMDHL:type_name -> CLJLLOFIBML
	16, // 15: BattleStatistics.BattleTargetInfoEntry.value:type_name -> BattleTargetList
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_BattleStatistics_proto_init() }
func file_BattleStatistics_proto_init() {
	if File_BattleStatistics_proto != nil {
		return
	}
	file_BattleEventBattleInfo_proto_init()
	file_LGFELGLKHJB_proto_init()
	file_AvatarBattleInfo_proto_init()
	file_MonsterBattleInfo_proto_init()
	file_EvolveBuildBattleInfo_proto_init()
	file_HHJEAHMHGFK_proto_init()
	file_DKODPCDGOOL_proto_init()
	file_BattleEndReason_proto_init()
	file_EDMLINEHPPE_proto_init()
	file_MFBFMKMABAO_proto_init()
	file_IGOJKPCDFNN_proto_init()
	file_CLJLLOFIBML_proto_init()
	file_BattleTargetList_proto_init()
	file_CKLPMMEKCDF_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BattleStatistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleStatistics); i {
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
			RawDescriptor: file_BattleStatistics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattleStatistics_proto_goTypes,
		DependencyIndexes: file_BattleStatistics_proto_depIdxs,
		MessageInfos:      file_BattleStatistics_proto_msgTypes,
	}.Build()
	File_BattleStatistics_proto = out.File
	file_BattleStatistics_proto_rawDesc = nil
	file_BattleStatistics_proto_goTypes = nil
	file_BattleStatistics_proto_depIdxs = nil
}
