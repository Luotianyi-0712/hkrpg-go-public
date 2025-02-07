// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueStartCsReq.proto

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

type ChessRogueStartCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAvatarIdList  []uint32 `protobuf:"varint,8,rep,packed,name=base_avatar_id_list,json=baseAvatarIdList,proto3" json:"base_avatar_id_list,omitempty"`
	DiceBranchId      uint32   `protobuf:"varint,3,opt,name=dice_branch_id,json=diceBranchId,proto3" json:"dice_branch_id,omitempty"`
	AeonId            uint32   `protobuf:"varint,13,opt,name=aeon_id,json=aeonId,proto3" json:"aeon_id,omitempty"`
	TrialAvatarIdList []uint32 `protobuf:"varint,1,rep,packed,name=trial_avatar_id_list,json=trialAvatarIdList,proto3" json:"trial_avatar_id_list,omitempty"`
	DisableAeonIdList []uint32 `protobuf:"varint,15,rep,packed,name=disable_aeon_id_list,json=disableAeonIdList,proto3" json:"disable_aeon_id_list,omitempty"`
	Id                uint32   `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
	DifficultyIdList  []uint32 `protobuf:"varint,7,rep,packed,name=difficulty_id_list,json=difficultyIdList,proto3" json:"difficulty_id_list,omitempty"`
}

func (x *ChessRogueStartCsReq) Reset() {
	*x = ChessRogueStartCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueStartCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueStartCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueStartCsReq) ProtoMessage() {}

func (x *ChessRogueStartCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueStartCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueStartCsReq.ProtoReflect.Descriptor instead.
func (*ChessRogueStartCsReq) Descriptor() ([]byte, []int) {
	return file_ChessRogueStartCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueStartCsReq) GetBaseAvatarIdList() []uint32 {
	if x != nil {
		return x.BaseAvatarIdList
	}
	return nil
}

func (x *ChessRogueStartCsReq) GetDiceBranchId() uint32 {
	if x != nil {
		return x.DiceBranchId
	}
	return 0
}

func (x *ChessRogueStartCsReq) GetAeonId() uint32 {
	if x != nil {
		return x.AeonId
	}
	return 0
}

func (x *ChessRogueStartCsReq) GetTrialAvatarIdList() []uint32 {
	if x != nil {
		return x.TrialAvatarIdList
	}
	return nil
}

func (x *ChessRogueStartCsReq) GetDisableAeonIdList() []uint32 {
	if x != nil {
		return x.DisableAeonIdList
	}
	return nil
}

func (x *ChessRogueStartCsReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChessRogueStartCsReq) GetDifficultyIdList() []uint32 {
	if x != nil {
		return x.DifficultyIdList
	}
	return nil
}

var File_ChessRogueStartCsReq_proto protoreflect.FileDescriptor

var file_ChessRogueStartCsReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a,
	0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x13, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x10, 0x62, 0x61, 0x73, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x69, 0x63, 0x65, 0x5f, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x69,
	0x63, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x65,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x65, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x11, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x65, 0x6f, 0x6e, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x10, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueStartCsReq_proto_rawDescOnce sync.Once
	file_ChessRogueStartCsReq_proto_rawDescData = file_ChessRogueStartCsReq_proto_rawDesc
)

func file_ChessRogueStartCsReq_proto_rawDescGZIP() []byte {
	file_ChessRogueStartCsReq_proto_rawDescOnce.Do(func() {
		file_ChessRogueStartCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueStartCsReq_proto_rawDescData)
	})
	return file_ChessRogueStartCsReq_proto_rawDescData
}

var file_ChessRogueStartCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueStartCsReq_proto_goTypes = []interface{}{
	(*ChessRogueStartCsReq)(nil), // 0: ChessRogueStartCsReq
}
var file_ChessRogueStartCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChessRogueStartCsReq_proto_init() }
func file_ChessRogueStartCsReq_proto_init() {
	if File_ChessRogueStartCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueStartCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueStartCsReq); i {
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
			RawDescriptor: file_ChessRogueStartCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueStartCsReq_proto_goTypes,
		DependencyIndexes: file_ChessRogueStartCsReq_proto_depIdxs,
		MessageInfos:      file_ChessRogueStartCsReq_proto_msgTypes,
	}.Build()
	File_ChessRogueStartCsReq_proto = out.File
	file_ChessRogueStartCsReq_proto_rawDesc = nil
	file_ChessRogueStartCsReq_proto_goTypes = nil
	file_ChessRogueStartCsReq_proto_depIdxs = nil
}
