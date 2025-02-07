// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: StartChallengeScRsp.proto

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

type StartChallengeScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scene        *SceneInfo    `protobuf:"bytes,11,opt,name=scene,proto3" json:"scene,omitempty"`
	PlayerInfo   *LMJLNMPCJJA  `protobuf:"bytes,12,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	LineupList   []*LineupInfo `protobuf:"bytes,14,rep,name=lineup_list,json=lineupList,proto3" json:"lineup_list,omitempty"`
	CurChallenge *CurChallenge `protobuf:"bytes,8,opt,name=cur_challenge,json=curChallenge,proto3" json:"cur_challenge,omitempty"`
	Retcode      uint32        `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *StartChallengeScRsp) Reset() {
	*x = StartChallengeScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StartChallengeScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartChallengeScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartChallengeScRsp) ProtoMessage() {}

func (x *StartChallengeScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_StartChallengeScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartChallengeScRsp.ProtoReflect.Descriptor instead.
func (*StartChallengeScRsp) Descriptor() ([]byte, []int) {
	return file_StartChallengeScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *StartChallengeScRsp) GetScene() *SceneInfo {
	if x != nil {
		return x.Scene
	}
	return nil
}

func (x *StartChallengeScRsp) GetPlayerInfo() *LMJLNMPCJJA {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

func (x *StartChallengeScRsp) GetLineupList() []*LineupInfo {
	if x != nil {
		return x.LineupList
	}
	return nil
}

func (x *StartChallengeScRsp) GetCurChallenge() *CurChallenge {
	if x != nil {
		return x.CurChallenge
	}
	return nil
}

func (x *StartChallengeScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_StartChallengeScRsp_proto protoreflect.FileDescriptor

var file_StartChallengeScRsp_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4c, 0x69, 0x6e,
	0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c,
	0x4d, 0x4a, 0x4c, 0x4e, 0x4d, 0x50, 0x43, 0x4a, 0x4a, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x43, 0x75, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a,
	0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x12,
	0x2d, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x4d, 0x4a, 0x4c, 0x4e, 0x4d, 0x50, 0x43, 0x4a,
	0x4a, 0x41, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c,
	0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0d,
	0x63, 0x75, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x75, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_StartChallengeScRsp_proto_rawDescOnce sync.Once
	file_StartChallengeScRsp_proto_rawDescData = file_StartChallengeScRsp_proto_rawDesc
)

func file_StartChallengeScRsp_proto_rawDescGZIP() []byte {
	file_StartChallengeScRsp_proto_rawDescOnce.Do(func() {
		file_StartChallengeScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_StartChallengeScRsp_proto_rawDescData)
	})
	return file_StartChallengeScRsp_proto_rawDescData
}

var file_StartChallengeScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StartChallengeScRsp_proto_goTypes = []interface{}{
	(*StartChallengeScRsp)(nil), // 0: StartChallengeScRsp
	(*SceneInfo)(nil),           // 1: SceneInfo
	(*LMJLNMPCJJA)(nil),         // 2: LMJLNMPCJJA
	(*LineupInfo)(nil),          // 3: LineupInfo
	(*CurChallenge)(nil),        // 4: CurChallenge
}
var file_StartChallengeScRsp_proto_depIdxs = []int32{
	1, // 0: StartChallengeScRsp.scene:type_name -> SceneInfo
	2, // 1: StartChallengeScRsp.player_info:type_name -> LMJLNMPCJJA
	3, // 2: StartChallengeScRsp.lineup_list:type_name -> LineupInfo
	4, // 3: StartChallengeScRsp.cur_challenge:type_name -> CurChallenge
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_StartChallengeScRsp_proto_init() }
func file_StartChallengeScRsp_proto_init() {
	if File_StartChallengeScRsp_proto != nil {
		return
	}
	file_LineupInfo_proto_init()
	file_LMJLNMPCJJA_proto_init()
	file_SceneInfo_proto_init()
	file_CurChallenge_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_StartChallengeScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartChallengeScRsp); i {
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
			RawDescriptor: file_StartChallengeScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StartChallengeScRsp_proto_goTypes,
		DependencyIndexes: file_StartChallengeScRsp_proto_depIdxs,
		MessageInfos:      file_StartChallengeScRsp_proto_msgTypes,
	}.Build()
	File_StartChallengeScRsp_proto = out.File
	file_StartChallengeScRsp_proto_rawDesc = nil
	file_StartChallengeScRsp_proto_goTypes = nil
	file_StartChallengeScRsp_proto_depIdxs = nil
}
