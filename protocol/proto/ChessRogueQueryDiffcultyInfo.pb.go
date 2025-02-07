// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueQueryDiffcultyInfo.proto

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

type ChessRogueQueryDiffcultyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryDifficultyIdList []uint32 `protobuf:"varint,14,rep,packed,name=query_difficulty_id_list,json=queryDifficultyIdList,proto3" json:"query_difficulty_id_list,omitempty"`
}

func (x *ChessRogueQueryDiffcultyInfo) Reset() {
	*x = ChessRogueQueryDiffcultyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueQueryDiffcultyInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueQueryDiffcultyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueQueryDiffcultyInfo) ProtoMessage() {}

func (x *ChessRogueQueryDiffcultyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueQueryDiffcultyInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueQueryDiffcultyInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueQueryDiffcultyInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueQueryDiffcultyInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueQueryDiffcultyInfo) GetQueryDifficultyIdList() []uint32 {
	if x != nil {
		return x.QueryDifficultyIdList
	}
	return nil
}

var File_ChessRogueQueryDiffcultyInfo_proto protoreflect.FileDescriptor

var file_ChessRogueQueryDiffcultyInfo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x44, 0x69, 0x66, 0x66, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x1c, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69, 0x66, 0x66, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x18, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x15, 0x71, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69, 0x66,
	0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueQueryDiffcultyInfo_proto_rawDescOnce sync.Once
	file_ChessRogueQueryDiffcultyInfo_proto_rawDescData = file_ChessRogueQueryDiffcultyInfo_proto_rawDesc
)

func file_ChessRogueQueryDiffcultyInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueQueryDiffcultyInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueQueryDiffcultyInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueQueryDiffcultyInfo_proto_rawDescData)
	})
	return file_ChessRogueQueryDiffcultyInfo_proto_rawDescData
}

var file_ChessRogueQueryDiffcultyInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueQueryDiffcultyInfo_proto_goTypes = []interface{}{
	(*ChessRogueQueryDiffcultyInfo)(nil), // 0: ChessRogueQueryDiffcultyInfo
}
var file_ChessRogueQueryDiffcultyInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChessRogueQueryDiffcultyInfo_proto_init() }
func file_ChessRogueQueryDiffcultyInfo_proto_init() {
	if File_ChessRogueQueryDiffcultyInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueQueryDiffcultyInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueQueryDiffcultyInfo); i {
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
			RawDescriptor: file_ChessRogueQueryDiffcultyInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueQueryDiffcultyInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueQueryDiffcultyInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueQueryDiffcultyInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueQueryDiffcultyInfo_proto = out.File
	file_ChessRogueQueryDiffcultyInfo_proto_rawDesc = nil
	file_ChessRogueQueryDiffcultyInfo_proto_goTypes = nil
	file_ChessRogueQueryDiffcultyInfo_proto_depIdxs = nil
}
