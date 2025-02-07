// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AvatarBattleInfo.proto

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

type AvatarBattleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarType            AvatarType              `protobuf:"varint,1,opt,name=avatar_type,json=avatarType,proto3,enum=AvatarType" json:"avatar_type,omitempty"`
	Id                    uint32                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	AvatarLevel           uint32                  `protobuf:"varint,3,opt,name=avatar_level,json=avatarLevel,proto3" json:"avatar_level,omitempty"`
	AvatarRank            uint32                  `protobuf:"varint,4,opt,name=avatar_rank,json=avatarRank,proto3" json:"avatar_rank,omitempty"`
	AvatarPromotion       uint32                  `protobuf:"varint,5,opt,name=avatar_promotion,json=avatarPromotion,proto3" json:"avatar_promotion,omitempty"`
	AvatarStatus          *AvatarProperty         `protobuf:"bytes,6,opt,name=avatar_status,json=avatarStatus,proto3" json:"avatar_status,omitempty"`
	AvatarSkill           []*AvatarSkillTree      `protobuf:"bytes,7,rep,name=avatar_skill,json=avatarSkill,proto3" json:"avatar_skill,omitempty"`
	AvatarEquipment       []*EquipmentProperty    `protobuf:"bytes,8,rep,name=avatar_equipment,json=avatarEquipment,proto3" json:"avatar_equipment,omitempty"`
	TotalTurns            uint32                  `protobuf:"varint,9,opt,name=total_turns,json=totalTurns,proto3" json:"total_turns,omitempty"`
	TotalDamage           float64                 `protobuf:"fixed64,10,opt,name=total_damage,json=totalDamage,proto3" json:"total_damage,omitempty"`
	TotalHeal             float64                 `protobuf:"fixed64,11,opt,name=total_heal,json=totalHeal,proto3" json:"total_heal,omitempty"`
	TotalDamageTaken      float64                 `protobuf:"fixed64,12,opt,name=total_damage_taken,json=totalDamageTaken,proto3" json:"total_damage_taken,omitempty"`
	TotalHpRecover        float64                 `protobuf:"fixed64,13,opt,name=total_hp_recover,json=totalHpRecover,proto3" json:"total_hp_recover,omitempty"`
	TotalSpCost           float64                 `protobuf:"fixed64,14,opt,name=total_sp_cost,json=totalSpCost,proto3" json:"total_sp_cost,omitempty"`
	StageId               uint32                  `protobuf:"varint,15,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	StageType             uint32                  `protobuf:"varint,16,opt,name=stage_type,json=stageType,proto3" json:"stage_type,omitempty"`
	TotalBreakDamage      float64                 `protobuf:"fixed64,17,opt,name=total_break_damage,json=totalBreakDamage,proto3" json:"total_break_damage,omitempty"`
	AttackTypeDamage      []*AttackDamageProperty `protobuf:"bytes,18,rep,name=attack_type_damage,json=attackTypeDamage,proto3" json:"attack_type_damage,omitempty"`
	AttackTypeBreakDamage []*AttackDamageProperty `protobuf:"bytes,19,rep,name=attack_type_break_damage,json=attackTypeBreakDamage,proto3" json:"attack_type_break_damage,omitempty"`
	AttackTypeMaxDamage   []*AttackDamageProperty `protobuf:"bytes,20,rep,name=attack_type_max_damage,json=attackTypeMaxDamage,proto3" json:"attack_type_max_damage,omitempty"`
	SkillTimes            []*SkillUseProperty     `protobuf:"bytes,21,rep,name=skill_times,json=skillTimes,proto3" json:"skill_times,omitempty"`
	DelayCumulate         float64                 `protobuf:"fixed64,22,opt,name=delay_cumulate,json=delayCumulate,proto3" json:"delay_cumulate,omitempty"`
	TotalSpAdd            uint32                  `protobuf:"varint,23,opt,name=total_sp_add,json=totalSpAdd,proto3" json:"total_sp_add,omitempty"`
	SpAddSource           []*SpAddSource          `protobuf:"bytes,24,rep,name=sp_add_source,json=spAddSource,proto3" json:"sp_add_source,omitempty"`
	TotalBpCost           uint32                  `protobuf:"varint,25,opt,name=total_bp_cost,json=totalBpCost,proto3" json:"total_bp_cost,omitempty"`
	DieTimes              uint32                  `protobuf:"varint,26,opt,name=die_times,json=dieTimes,proto3" json:"die_times,omitempty"`
	ReviveTimes           uint32                  `protobuf:"varint,27,opt,name=revive_times,json=reviveTimes,proto3" json:"revive_times,omitempty"`
	BreakTimes            uint32                  `protobuf:"varint,28,opt,name=break_times,json=breakTimes,proto3" json:"break_times,omitempty"`
	ExtraTurns            uint32                  `protobuf:"varint,29,opt,name=extra_turns,json=extraTurns,proto3" json:"extra_turns,omitempty"`
	TotalShield           float64                 `protobuf:"fixed64,30,opt,name=total_shield,json=totalShield,proto3" json:"total_shield,omitempty"`
	TotalShieldTaken      float64                 `protobuf:"fixed64,31,opt,name=total_shield_taken,json=totalShieldTaken,proto3" json:"total_shield_taken,omitempty"`
	TotalShieldDamage     float64                 `protobuf:"fixed64,32,opt,name=total_shield_damage,json=totalShieldDamage,proto3" json:"total_shield_damage,omitempty"`
	InitialStatus         *AvatarProperty         `protobuf:"bytes,33,opt,name=initial_status,json=initialStatus,proto3" json:"initial_status,omitempty"`
	Relics                []*BattleRelic          `protobuf:"bytes,34,rep,name=relics,proto3" json:"relics,omitempty"`
	AssistUid             uint32                  `protobuf:"varint,35,opt,name=assist_uid,json=assistUid,proto3" json:"assist_uid,omitempty"`
	NEIMBIEPHJM           []*AttackDamageProperty `protobuf:"bytes,36,rep,name=NEIMBIEPHJM,proto3" json:"NEIMBIEPHJM,omitempty"`
	DHIAIMNOAHJ           float64                 `protobuf:"fixed64,37,opt,name=DHIAIMNOAHJ,proto3" json:"DHIAIMNOAHJ,omitempty"`
	EHELGCFKAFI           float64                 `protobuf:"fixed64,38,opt,name=EHELGCFKAFI,proto3" json:"EHELGCFKAFI,omitempty"`
	AGBFFAPEGNE           float64                 `protobuf:"fixed64,39,opt,name=AGBFFAPEGNE,proto3" json:"AGBFFAPEGNE,omitempty"`
	GIPBPADKIEL           float64                 `protobuf:"fixed64,40,opt,name=GIPBPADKIEL,proto3" json:"GIPBPADKIEL,omitempty"`
	IAINOBNIKIK           []*AbilityUseStt        `protobuf:"bytes,41,rep,name=IAINOBNIKIK,proto3" json:"IAINOBNIKIK,omitempty"`
	ACBLIHHMDBE           uint32                  `protobuf:"varint,42,opt,name=ACBLIHHMDBE,proto3" json:"ACBLIHHMDBE,omitempty"`
	PKOCPLEPGMB           uint32                  `protobuf:"varint,43,opt,name=PKOCPLEPGMB,proto3" json:"PKOCPLEPGMB,omitempty"`
}

func (x *AvatarBattleInfo) Reset() {
	*x = AvatarBattleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarBattleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarBattleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarBattleInfo) ProtoMessage() {}

func (x *AvatarBattleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarBattleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarBattleInfo.ProtoReflect.Descriptor instead.
func (*AvatarBattleInfo) Descriptor() ([]byte, []int) {
	return file_AvatarBattleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarBattleInfo) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

func (x *AvatarBattleInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AvatarBattleInfo) GetAvatarLevel() uint32 {
	if x != nil {
		return x.AvatarLevel
	}
	return 0
}

func (x *AvatarBattleInfo) GetAvatarRank() uint32 {
	if x != nil {
		return x.AvatarRank
	}
	return 0
}

func (x *AvatarBattleInfo) GetAvatarPromotion() uint32 {
	if x != nil {
		return x.AvatarPromotion
	}
	return 0
}

func (x *AvatarBattleInfo) GetAvatarStatus() *AvatarProperty {
	if x != nil {
		return x.AvatarStatus
	}
	return nil
}

func (x *AvatarBattleInfo) GetAvatarSkill() []*AvatarSkillTree {
	if x != nil {
		return x.AvatarSkill
	}
	return nil
}

func (x *AvatarBattleInfo) GetAvatarEquipment() []*EquipmentProperty {
	if x != nil {
		return x.AvatarEquipment
	}
	return nil
}

func (x *AvatarBattleInfo) GetTotalTurns() uint32 {
	if x != nil {
		return x.TotalTurns
	}
	return 0
}

func (x *AvatarBattleInfo) GetTotalDamage() float64 {
	if x != nil {
		return x.TotalDamage
	}
	return 0
}

func (x *AvatarBattleInfo) GetTotalHeal() float64 {
	if x != nil {
		return x.TotalHeal
	}
	return 0
}

func (x *AvatarBattleInfo) GetTotalDamageTaken() float64 {
	if x != nil {
		return x.TotalDamageTaken
	}
	return 0
}

func (x *AvatarBattleInfo) GetTotalHpRecover() float64 {
	if x != nil {
		return x.TotalHpRecover
	}
	return 0
}

func (x *AvatarBattleInfo) GetTotalSpCost() float64 {
	if x != nil {
		return x.TotalSpCost
	}
	return 0
}

func (x *AvatarBattleInfo) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *AvatarBattleInfo) GetStageType() uint32 {
	if x != nil {
		return x.StageType
	}
	return 0
}

func (x *AvatarBattleInfo) GetTotalBreakDamage() float64 {
	if x != nil {
		return x.TotalBreakDamage
	}
	return 0
}

func (x *AvatarBattleInfo) GetAttackTypeDamage() []*AttackDamageProperty {
	if x != nil {
		return x.AttackTypeDamage
	}
	return nil
}

func (x *AvatarBattleInfo) GetAttackTypeBreakDamage() []*AttackDamageProperty {
	if x != nil {
		return x.AttackTypeBreakDamage
	}
	return nil
}

func (x *AvatarBattleInfo) GetAttackTypeMaxDamage() []*AttackDamageProperty {
	if x != nil {
		return x.AttackTypeMaxDamage
	}
	return nil
}

func (x *AvatarBattleInfo) GetSkillTimes() []*SkillUseProperty {
	if x != nil {
		return x.SkillTimes
	}
	return nil
}

func (x *AvatarBattleInfo) GetDelayCumulate() float64 {
	if x != nil {
		return x.DelayCumulate
	}
	return 0
}

func (x *AvatarBattleInfo) GetTotalSpAdd() uint32 {
	if x != nil {
		return x.TotalSpAdd
	}
	return 0
}

func (x *AvatarBattleInfo) GetSpAddSource() []*SpAddSource {
	if x != nil {
		return x.SpAddSource
	}
	return nil
}

func (x *AvatarBattleInfo) GetTotalBpCost() uint32 {
	if x != nil {
		return x.TotalBpCost
	}
	return 0
}

func (x *AvatarBattleInfo) GetDieTimes() uint32 {
	if x != nil {
		return x.DieTimes
	}
	return 0
}

func (x *AvatarBattleInfo) GetReviveTimes() uint32 {
	if x != nil {
		return x.ReviveTimes
	}
	return 0
}

func (x *AvatarBattleInfo) GetBreakTimes() uint32 {
	if x != nil {
		return x.BreakTimes
	}
	return 0
}

func (x *AvatarBattleInfo) GetExtraTurns() uint32 {
	if x != nil {
		return x.ExtraTurns
	}
	return 0
}

func (x *AvatarBattleInfo) GetTotalShield() float64 {
	if x != nil {
		return x.TotalShield
	}
	return 0
}

func (x *AvatarBattleInfo) GetTotalShieldTaken() float64 {
	if x != nil {
		return x.TotalShieldTaken
	}
	return 0
}

func (x *AvatarBattleInfo) GetTotalShieldDamage() float64 {
	if x != nil {
		return x.TotalShieldDamage
	}
	return 0
}

func (x *AvatarBattleInfo) GetInitialStatus() *AvatarProperty {
	if x != nil {
		return x.InitialStatus
	}
	return nil
}

func (x *AvatarBattleInfo) GetRelics() []*BattleRelic {
	if x != nil {
		return x.Relics
	}
	return nil
}

func (x *AvatarBattleInfo) GetAssistUid() uint32 {
	if x != nil {
		return x.AssistUid
	}
	return 0
}

func (x *AvatarBattleInfo) GetNEIMBIEPHJM() []*AttackDamageProperty {
	if x != nil {
		return x.NEIMBIEPHJM
	}
	return nil
}

func (x *AvatarBattleInfo) GetDHIAIMNOAHJ() float64 {
	if x != nil {
		return x.DHIAIMNOAHJ
	}
	return 0
}

func (x *AvatarBattleInfo) GetEHELGCFKAFI() float64 {
	if x != nil {
		return x.EHELGCFKAFI
	}
	return 0
}

func (x *AvatarBattleInfo) GetAGBFFAPEGNE() float64 {
	if x != nil {
		return x.AGBFFAPEGNE
	}
	return 0
}

func (x *AvatarBattleInfo) GetGIPBPADKIEL() float64 {
	if x != nil {
		return x.GIPBPADKIEL
	}
	return 0
}

func (x *AvatarBattleInfo) GetIAINOBNIKIK() []*AbilityUseStt {
	if x != nil {
		return x.IAINOBNIKIK
	}
	return nil
}

func (x *AvatarBattleInfo) GetACBLIHHMDBE() uint32 {
	if x != nil {
		return x.ACBLIHHMDBE
	}
	return 0
}

func (x *AvatarBattleInfo) GetPKOCPLEPGMB() uint32 {
	if x != nil {
		return x.PKOCPLEPGMB
	}
	return 0
}

var File_AvatarBattleInfo_proto protoreflect.FileDescriptor

var file_AvatarBattleInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x73, 0x65, 0x53, 0x74, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x53, 0x70, 0x41, 0x64, 0x64, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x54, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9b, 0x0e, 0x0a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x0d, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x54, 0x72, 0x65, 0x65, 0x52,
	0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x3d, 0x0a, 0x10,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0f, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x75, 0x72, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x12, 0x2c,
	0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x61, 0x6b, 0x65, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x10,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x68, 0x70, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x48, 0x70, 0x52,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x70, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x72,
	0x65, 0x61, 0x6b, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x44, 0x61, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x43, 0x0a, 0x12, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x10, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x18, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f, 0x64, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x6b, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x52, 0x15, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x4a, 0x0a, 0x16, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x13,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x78, 0x44, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0a, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x5f, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x43, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x41, 0x64, 0x64,
	0x12, 0x30, 0x0a, 0x0d, 0x73, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x70, 0x41, 0x64, 0x64, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x73, 0x70, 0x41, 0x64, 0x64, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x70, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x42, 0x70, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x69, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x76,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x62, 0x72, 0x65,
	0x61, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x5f, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x54, 0x75, 0x72, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x61, 0x6b, 0x65,
	0x6e, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68,
	0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x68, 0x69,
	0x65, 0x6c, 0x64, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x0e, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x21, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x73, 0x18, 0x22, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x52,
	0x06, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x73, 0x73,
	0x69, 0x73, 0x74, 0x55, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x4e, 0x45, 0x49, 0x4d, 0x42, 0x49,
	0x45, 0x50, 0x48, 0x4a, 0x4d, 0x18, 0x24, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x52, 0x0b, 0x4e, 0x45, 0x49, 0x4d, 0x42, 0x49, 0x45, 0x50, 0x48, 0x4a, 0x4d, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x48, 0x49, 0x41, 0x49, 0x4d, 0x4e, 0x4f, 0x41, 0x48, 0x4a, 0x18, 0x25,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x44, 0x48, 0x49, 0x41, 0x49, 0x4d, 0x4e, 0x4f, 0x41, 0x48,
	0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x48, 0x45, 0x4c, 0x47, 0x43, 0x46, 0x4b, 0x41, 0x46, 0x49,
	0x18, 0x26, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x45, 0x48, 0x45, 0x4c, 0x47, 0x43, 0x46, 0x4b,
	0x41, 0x46, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x47, 0x42, 0x46, 0x46, 0x41, 0x50, 0x45, 0x47,
	0x4e, 0x45, 0x18, 0x27, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x41, 0x47, 0x42, 0x46, 0x46, 0x41,
	0x50, 0x45, 0x47, 0x4e, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x49, 0x50, 0x42, 0x50, 0x41, 0x44,
	0x4b, 0x49, 0x45, 0x4c, 0x18, 0x28, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x47, 0x49, 0x50, 0x42,
	0x50, 0x41, 0x44, 0x4b, 0x49, 0x45, 0x4c, 0x12, 0x30, 0x0a, 0x0b, 0x49, 0x41, 0x49, 0x4e, 0x4f,
	0x42, 0x4e, 0x49, 0x4b, 0x49, 0x4b, 0x18, 0x29, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x55, 0x73, 0x65, 0x53, 0x74, 0x74, 0x52, 0x0b, 0x49, 0x41,
	0x49, 0x4e, 0x4f, 0x42, 0x4e, 0x49, 0x4b, 0x49, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x43, 0x42,
	0x4c, 0x49, 0x48, 0x48, 0x4d, 0x44, 0x42, 0x45, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x41, 0x43, 0x42, 0x4c, 0x49, 0x48, 0x48, 0x4d, 0x44, 0x42, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x4b, 0x4f, 0x43, 0x50, 0x4c, 0x45, 0x50, 0x47, 0x4d, 0x42, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x50, 0x4b, 0x4f, 0x43, 0x50, 0x4c, 0x45, 0x50, 0x47, 0x4d, 0x42, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarBattleInfo_proto_rawDescOnce sync.Once
	file_AvatarBattleInfo_proto_rawDescData = file_AvatarBattleInfo_proto_rawDesc
)

func file_AvatarBattleInfo_proto_rawDescGZIP() []byte {
	file_AvatarBattleInfo_proto_rawDescOnce.Do(func() {
		file_AvatarBattleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarBattleInfo_proto_rawDescData)
	})
	return file_AvatarBattleInfo_proto_rawDescData
}

var file_AvatarBattleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarBattleInfo_proto_goTypes = []interface{}{
	(*AvatarBattleInfo)(nil),     // 0: AvatarBattleInfo
	(AvatarType)(0),              // 1: AvatarType
	(*AvatarProperty)(nil),       // 2: AvatarProperty
	(*AvatarSkillTree)(nil),      // 3: AvatarSkillTree
	(*EquipmentProperty)(nil),    // 4: EquipmentProperty
	(*AttackDamageProperty)(nil), // 5: AttackDamageProperty
	(*SkillUseProperty)(nil),     // 6: SkillUseProperty
	(*SpAddSource)(nil),          // 7: SpAddSource
	(*BattleRelic)(nil),          // 8: BattleRelic
	(*AbilityUseStt)(nil),        // 9: AbilityUseStt
}
var file_AvatarBattleInfo_proto_depIdxs = []int32{
	1,  // 0: AvatarBattleInfo.avatar_type:type_name -> AvatarType
	2,  // 1: AvatarBattleInfo.avatar_status:type_name -> AvatarProperty
	3,  // 2: AvatarBattleInfo.avatar_skill:type_name -> AvatarSkillTree
	4,  // 3: AvatarBattleInfo.avatar_equipment:type_name -> EquipmentProperty
	5,  // 4: AvatarBattleInfo.attack_type_damage:type_name -> AttackDamageProperty
	5,  // 5: AvatarBattleInfo.attack_type_break_damage:type_name -> AttackDamageProperty
	5,  // 6: AvatarBattleInfo.attack_type_max_damage:type_name -> AttackDamageProperty
	6,  // 7: AvatarBattleInfo.skill_times:type_name -> SkillUseProperty
	7,  // 8: AvatarBattleInfo.sp_add_source:type_name -> SpAddSource
	2,  // 9: AvatarBattleInfo.initial_status:type_name -> AvatarProperty
	8,  // 10: AvatarBattleInfo.relics:type_name -> BattleRelic
	5,  // 11: AvatarBattleInfo.NEIMBIEPHJM:type_name -> AttackDamageProperty
	9,  // 12: AvatarBattleInfo.IAINOBNIKIK:type_name -> AbilityUseStt
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_AvatarBattleInfo_proto_init() }
func file_AvatarBattleInfo_proto_init() {
	if File_AvatarBattleInfo_proto != nil {
		return
	}
	file_EquipmentProperty_proto_init()
	file_AbilityUseStt_proto_init()
	file_AvatarProperty_proto_init()
	file_SkillUseProperty_proto_init()
	file_BattleRelic_proto_init()
	file_SpAddSource_proto_init()
	file_AttackDamageProperty_proto_init()
	file_AvatarSkillTree_proto_init()
	file_AvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarBattleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarBattleInfo); i {
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
			RawDescriptor: file_AvatarBattleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarBattleInfo_proto_goTypes,
		DependencyIndexes: file_AvatarBattleInfo_proto_depIdxs,
		MessageInfos:      file_AvatarBattleInfo_proto_msgTypes,
	}.Build()
	File_AvatarBattleInfo_proto = out.File
	file_AvatarBattleInfo_proto_rawDesc = nil
	file_AvatarBattleInfo_proto_goTypes = nil
	file_AvatarBattleInfo_proto_depIdxs = nil
}
