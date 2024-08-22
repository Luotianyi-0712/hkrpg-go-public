// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerDetailInfo.proto

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

type PlayerDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsBanned          bool                       `protobuf:"varint,11,opt,name=is_banned,json=isBanned,proto3" json:"is_banned,omitempty"`
	WorldLevel        uint32                     `protobuf:"varint,6,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	MLDNNOMLHJD       *DNCLPGJGHHK               `protobuf:"bytes,1265,opt,name=MLDNNOMLHJD,proto3" json:"MLDNNOMLHJD,omitempty"`
	LFFJMADBHNN       uint32                     `protobuf:"varint,10,opt,name=LFFJMADBHNN,proto3" json:"LFFJMADBHNN,omitempty"`
	HeadIcon          uint32                     `protobuf:"varint,12,opt,name=head_icon,json=headIcon,proto3" json:"head_icon,omitempty"`
	OPIACEKOANJ       string                     `protobuf:"bytes,3,opt,name=OPIACEKOANJ,proto3" json:"OPIACEKOANJ,omitempty"`
	Signature         string                     `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	DisplayAvatarList []*DisplayAvatarDetailInfo `protobuf:"bytes,8,rep,name=display_avatar_list,json=displayAvatarList,proto3" json:"display_avatar_list,omitempty"`
	AssistAvatarList  []*DisplayAvatarDetailInfo `protobuf:"bytes,1854,rep,name=assist_avatar_list,json=assistAvatarList,proto3" json:"assist_avatar_list,omitempty"`
	Platform          PlatformType               `protobuf:"varint,5,opt,name=platform,proto3,enum=PlatformType" json:"platform,omitempty"`
	RecordInfo        *PlayerRecordInfo          `protobuf:"bytes,13,opt,name=record_info,json=recordInfo,proto3" json:"record_info,omitempty"`
	DisplaySettings   *PlayerDisplaySettings     `protobuf:"bytes,1079,opt,name=display_settings,json=displaySettings,proto3" json:"display_settings,omitempty"`
	Uid               uint32                     `protobuf:"varint,9,opt,name=uid,proto3" json:"uid,omitempty"`
	COAELHIGPNG       uint32                     `protobuf:"varint,15,opt,name=COAELHIGPNG,proto3" json:"COAELHIGPNG,omitempty"`
	Level             uint32                     `protobuf:"varint,14,opt,name=level,proto3" json:"level,omitempty"`
	Nickname          string                     `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	HNADOAKOMEE       bool                       `protobuf:"varint,1,opt,name=HNADOAKOMEE,proto3" json:"HNADOAKOMEE,omitempty"`
	MCKKBKPOMLI       string                     `protobuf:"bytes,7,opt,name=MCKKBKPOMLI,proto3" json:"MCKKBKPOMLI,omitempty"`
}

func (x *PlayerDetailInfo) Reset() {
	*x = PlayerDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerDetailInfo) ProtoMessage() {}

func (x *PlayerDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerDetailInfo.ProtoReflect.Descriptor instead.
func (*PlayerDetailInfo) Descriptor() ([]byte, []int) {
	return file_PlayerDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerDetailInfo) GetIsBanned() bool {
	if x != nil {
		return x.IsBanned
	}
	return false
}

func (x *PlayerDetailInfo) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *PlayerDetailInfo) GetMLDNNOMLHJD() *DNCLPGJGHHK {
	if x != nil {
		return x.MLDNNOMLHJD
	}
	return nil
}

func (x *PlayerDetailInfo) GetLFFJMADBHNN() uint32 {
	if x != nil {
		return x.LFFJMADBHNN
	}
	return 0
}

func (x *PlayerDetailInfo) GetHeadIcon() uint32 {
	if x != nil {
		return x.HeadIcon
	}
	return 0
}

func (x *PlayerDetailInfo) GetOPIACEKOANJ() string {
	if x != nil {
		return x.OPIACEKOANJ
	}
	return ""
}

func (x *PlayerDetailInfo) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *PlayerDetailInfo) GetDisplayAvatarList() []*DisplayAvatarDetailInfo {
	if x != nil {
		return x.DisplayAvatarList
	}
	return nil
}

func (x *PlayerDetailInfo) GetAssistAvatarList() []*DisplayAvatarDetailInfo {
	if x != nil {
		return x.AssistAvatarList
	}
	return nil
}

func (x *PlayerDetailInfo) GetPlatform() PlatformType {
	if x != nil {
		return x.Platform
	}
	return PlatformType_EDITOR
}

func (x *PlayerDetailInfo) GetRecordInfo() *PlayerRecordInfo {
	if x != nil {
		return x.RecordInfo
	}
	return nil
}

func (x *PlayerDetailInfo) GetDisplaySettings() *PlayerDisplaySettings {
	if x != nil {
		return x.DisplaySettings
	}
	return nil
}

func (x *PlayerDetailInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PlayerDetailInfo) GetCOAELHIGPNG() uint32 {
	if x != nil {
		return x.COAELHIGPNG
	}
	return 0
}

func (x *PlayerDetailInfo) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *PlayerDetailInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *PlayerDetailInfo) GetHNADOAKOMEE() bool {
	if x != nil {
		return x.HNADOAKOMEE
	}
	return false
}

func (x *PlayerDetailInfo) GetMCKKBKPOMLI() string {
	if x != nil {
		return x.MCKKBKPOMLI
	}
	return ""
}

var File_PlayerDetailInfo_proto protoreflect.FileDescriptor

var file_PlayerDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x44, 0x4e, 0x43, 0x4c, 0x50, 0x47, 0x4a, 0x47, 0x48, 0x48, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x05, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x42,
	0x61, 0x6e, 0x6e, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2f, 0x0a, 0x0b, 0x4d, 0x4c, 0x44, 0x4e, 0x4e, 0x4f,
	0x4d, 0x4c, 0x48, 0x4a, 0x44, 0x18, 0xf1, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44,
	0x4e, 0x43, 0x4c, 0x50, 0x47, 0x4a, 0x47, 0x48, 0x48, 0x4b, 0x52, 0x0b, 0x4d, 0x4c, 0x44, 0x4e,
	0x4e, 0x4f, 0x4d, 0x4c, 0x48, 0x4a, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x46, 0x46, 0x4a, 0x4d,
	0x41, 0x44, 0x42, 0x48, 0x4e, 0x4e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x46,
	0x46, 0x4a, 0x4d, 0x41, 0x44, 0x42, 0x48, 0x4e, 0x4e, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x61,
	0x64, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x68, 0x65,
	0x61, 0x64, 0x49, 0x63, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x50, 0x49, 0x41, 0x43, 0x45,
	0x4b, 0x4f, 0x41, 0x4e, 0x4a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x50, 0x49,
	0x41, 0x43, 0x45, 0x4b, 0x4f, 0x41, 0x4e, 0x4a, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x48, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x47, 0x0a, 0x12, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xbe, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x32, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0xb7, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0f, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x43, 0x4f, 0x41, 0x45, 0x4c, 0x48, 0x49, 0x47, 0x50, 0x4e, 0x47, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x4f, 0x41, 0x45, 0x4c, 0x48, 0x49, 0x47, 0x50, 0x4e, 0x47,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4e, 0x41, 0x44, 0x4f, 0x41, 0x4b, 0x4f, 0x4d, 0x45,
	0x45, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x48, 0x4e, 0x41, 0x44, 0x4f, 0x41, 0x4b,
	0x4f, 0x4d, 0x45, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x43, 0x4b, 0x4b, 0x42, 0x4b, 0x50, 0x4f,
	0x4d, 0x4c, 0x49, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x43, 0x4b, 0x4b, 0x42,
	0x4b, 0x50, 0x4f, 0x4d, 0x4c, 0x49, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerDetailInfo_proto_rawDescOnce sync.Once
	file_PlayerDetailInfo_proto_rawDescData = file_PlayerDetailInfo_proto_rawDesc
)

func file_PlayerDetailInfo_proto_rawDescGZIP() []byte {
	file_PlayerDetailInfo_proto_rawDescOnce.Do(func() {
		file_PlayerDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerDetailInfo_proto_rawDescData)
	})
	return file_PlayerDetailInfo_proto_rawDescData
}

var file_PlayerDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerDetailInfo_proto_goTypes = []interface{}{
	(*PlayerDetailInfo)(nil),        // 0: PlayerDetailInfo
	(*DNCLPGJGHHK)(nil),             // 1: DNCLPGJGHHK
	(*DisplayAvatarDetailInfo)(nil), // 2: DisplayAvatarDetailInfo
	(PlatformType)(0),               // 3: PlatformType
	(*PlayerRecordInfo)(nil),        // 4: PlayerRecordInfo
	(*PlayerDisplaySettings)(nil),   // 5: PlayerDisplaySettings
}
var file_PlayerDetailInfo_proto_depIdxs = []int32{
	1, // 0: PlayerDetailInfo.MLDNNOMLHJD:type_name -> DNCLPGJGHHK
	2, // 1: PlayerDetailInfo.display_avatar_list:type_name -> DisplayAvatarDetailInfo
	2, // 2: PlayerDetailInfo.assist_avatar_list:type_name -> DisplayAvatarDetailInfo
	3, // 3: PlayerDetailInfo.platform:type_name -> PlatformType
	4, // 4: PlayerDetailInfo.record_info:type_name -> PlayerRecordInfo
	5, // 5: PlayerDetailInfo.display_settings:type_name -> PlayerDisplaySettings
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_PlayerDetailInfo_proto_init() }
func file_PlayerDetailInfo_proto_init() {
	if File_PlayerDetailInfo_proto != nil {
		return
	}
	file_PlayerRecordInfo_proto_init()
	file_PlayerDisplaySettings_proto_init()
	file_PlatformType_proto_init()
	file_DisplayAvatarDetailInfo_proto_init()
	file_DNCLPGJGHHK_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerDetailInfo); i {
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
			RawDescriptor: file_PlayerDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerDetailInfo_proto_goTypes,
		DependencyIndexes: file_PlayerDetailInfo_proto_depIdxs,
		MessageInfos:      file_PlayerDetailInfo_proto_msgTypes,
	}.Build()
	File_PlayerDetailInfo_proto = out.File
	file_PlayerDetailInfo_proto_rawDesc = nil
	file_PlayerDetailInfo_proto_goTypes = nil
	file_PlayerDetailInfo_proto_depIdxs = nil
}
