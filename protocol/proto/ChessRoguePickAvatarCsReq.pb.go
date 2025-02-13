// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRoguePickAvatarCsReq.proto

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

type ChessRoguePickAvatarCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAvatarIdList []uint32 `protobuf:"varint,4,rep,packed,name=base_avatar_id_list,json=baseAvatarIdList,proto3" json:"base_avatar_id_list,omitempty"`
	PropEntityId     uint32   `protobuf:"varint,10,opt,name=prop_entity_id,json=propEntityId,proto3" json:"prop_entity_id,omitempty"`
}

func (x *ChessRoguePickAvatarCsReq) Reset() {
	*x = ChessRoguePickAvatarCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRoguePickAvatarCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRoguePickAvatarCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRoguePickAvatarCsReq) ProtoMessage() {}

func (x *ChessRoguePickAvatarCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRoguePickAvatarCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRoguePickAvatarCsReq.ProtoReflect.Descriptor instead.
func (*ChessRoguePickAvatarCsReq) Descriptor() ([]byte, []int) {
	return file_ChessRoguePickAvatarCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRoguePickAvatarCsReq) GetBaseAvatarIdList() []uint32 {
	if x != nil {
		return x.BaseAvatarIdList
	}
	return nil
}

func (x *ChessRoguePickAvatarCsReq) GetPropEntityId() uint32 {
	if x != nil {
		return x.PropEntityId
	}
	return 0
}

var File_ChessRoguePickAvatarCsReq_proto protoreflect.FileDescriptor

var file_ChessRoguePickAvatarCsReq_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x50, 0x69, 0x63, 0x6b,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x70, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x50,
	0x69, 0x63, 0x6b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2d,
	0x0a, 0x13, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x62, 0x61, 0x73,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRoguePickAvatarCsReq_proto_rawDescOnce sync.Once
	file_ChessRoguePickAvatarCsReq_proto_rawDescData = file_ChessRoguePickAvatarCsReq_proto_rawDesc
)

func file_ChessRoguePickAvatarCsReq_proto_rawDescGZIP() []byte {
	file_ChessRoguePickAvatarCsReq_proto_rawDescOnce.Do(func() {
		file_ChessRoguePickAvatarCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRoguePickAvatarCsReq_proto_rawDescData)
	})
	return file_ChessRoguePickAvatarCsReq_proto_rawDescData
}

var file_ChessRoguePickAvatarCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRoguePickAvatarCsReq_proto_goTypes = []interface{}{
	(*ChessRoguePickAvatarCsReq)(nil), // 0: ChessRoguePickAvatarCsReq
}
var file_ChessRoguePickAvatarCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChessRoguePickAvatarCsReq_proto_init() }
func file_ChessRoguePickAvatarCsReq_proto_init() {
	if File_ChessRoguePickAvatarCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChessRoguePickAvatarCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRoguePickAvatarCsReq); i {
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
			RawDescriptor: file_ChessRoguePickAvatarCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRoguePickAvatarCsReq_proto_goTypes,
		DependencyIndexes: file_ChessRoguePickAvatarCsReq_proto_depIdxs,
		MessageInfos:      file_ChessRoguePickAvatarCsReq_proto_msgTypes,
	}.Build()
	File_ChessRoguePickAvatarCsReq_proto = out.File
	file_ChessRoguePickAvatarCsReq_proto_rawDesc = nil
	file_ChessRoguePickAvatarCsReq_proto_goTypes = nil
	file_ChessRoguePickAvatarCsReq_proto_depIdxs = nil
}
