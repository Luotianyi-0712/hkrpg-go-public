// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueCommonDialogueOptionResultInfo.proto

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

type RogueCommonDialogueOptionResultInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OptionResult:
	//
	//	*RogueCommonDialogueOptionResultInfo_BattleResultInfo
	//	*RogueCommonDialogueOptionResultInfo_ACHLGHPPMNO
	OptionResult isRogueCommonDialogueOptionResultInfo_OptionResult `protobuf_oneof:"option_result"`
}

func (x *RogueCommonDialogueOptionResultInfo) Reset() {
	*x = RogueCommonDialogueOptionResultInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueCommonDialogueOptionResultInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueCommonDialogueOptionResultInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueCommonDialogueOptionResultInfo) ProtoMessage() {}

func (x *RogueCommonDialogueOptionResultInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueCommonDialogueOptionResultInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueCommonDialogueOptionResultInfo.ProtoReflect.Descriptor instead.
func (*RogueCommonDialogueOptionResultInfo) Descriptor() ([]byte, []int) {
	return file_RogueCommonDialogueOptionResultInfo_proto_rawDescGZIP(), []int{0}
}

func (m *RogueCommonDialogueOptionResultInfo) GetOptionResult() isRogueCommonDialogueOptionResultInfo_OptionResult {
	if m != nil {
		return m.OptionResult
	}
	return nil
}

func (x *RogueCommonDialogueOptionResultInfo) GetBattleResultInfo() *RogueCommonDialogueOptionBattleResultInfo {
	if x, ok := x.GetOptionResult().(*RogueCommonDialogueOptionResultInfo_BattleResultInfo); ok {
		return x.BattleResultInfo
	}
	return nil
}

func (x *RogueCommonDialogueOptionResultInfo) GetACHLGHPPMNO() *OJAHMDLNMAK {
	if x, ok := x.GetOptionResult().(*RogueCommonDialogueOptionResultInfo_ACHLGHPPMNO); ok {
		return x.ACHLGHPPMNO
	}
	return nil
}

type isRogueCommonDialogueOptionResultInfo_OptionResult interface {
	isRogueCommonDialogueOptionResultInfo_OptionResult()
}

type RogueCommonDialogueOptionResultInfo_BattleResultInfo struct {
	BattleResultInfo *RogueCommonDialogueOptionBattleResultInfo `protobuf:"bytes,11,opt,name=battle_result_info,json=battleResultInfo,proto3,oneof"`
}

type RogueCommonDialogueOptionResultInfo_ACHLGHPPMNO struct {
	ACHLGHPPMNO *OJAHMDLNMAK `protobuf:"bytes,3,opt,name=ACHLGHPPMNO,proto3,oneof"`
}

func (*RogueCommonDialogueOptionResultInfo_BattleResultInfo) isRogueCommonDialogueOptionResultInfo_OptionResult() {
}

func (*RogueCommonDialogueOptionResultInfo_ACHLGHPPMNO) isRogueCommonDialogueOptionResultInfo_OptionResult() {
}

var File_RogueCommonDialogueOptionResultInfo_proto protoreflect.FileDescriptor

var file_RogueCommonDialogueOptionResultInfo_proto_rawDesc = []byte{
	0x0a, 0x29, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4a, 0x41,
	0x48, 0x4d, 0x44, 0x4c, 0x4e, 0x4d, 0x41, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f,
	0x67, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc4, 0x01, 0x0a, 0x23, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5a, 0x0a, 0x12, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x10, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x0b, 0x41, 0x43, 0x48, 0x4c, 0x47, 0x48, 0x50, 0x50, 0x4d,
	0x4e, 0x4f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4a, 0x41, 0x48, 0x4d,
	0x44, 0x4c, 0x4e, 0x4d, 0x41, 0x4b, 0x48, 0x00, 0x52, 0x0b, 0x41, 0x43, 0x48, 0x4c, 0x47, 0x48,
	0x50, 0x50, 0x4d, 0x4e, 0x4f, 0x42, 0x0f, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueCommonDialogueOptionResultInfo_proto_rawDescOnce sync.Once
	file_RogueCommonDialogueOptionResultInfo_proto_rawDescData = file_RogueCommonDialogueOptionResultInfo_proto_rawDesc
)

func file_RogueCommonDialogueOptionResultInfo_proto_rawDescGZIP() []byte {
	file_RogueCommonDialogueOptionResultInfo_proto_rawDescOnce.Do(func() {
		file_RogueCommonDialogueOptionResultInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueCommonDialogueOptionResultInfo_proto_rawDescData)
	})
	return file_RogueCommonDialogueOptionResultInfo_proto_rawDescData
}

var file_RogueCommonDialogueOptionResultInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueCommonDialogueOptionResultInfo_proto_goTypes = []interface{}{
	(*RogueCommonDialogueOptionResultInfo)(nil),       // 0: RogueCommonDialogueOptionResultInfo
	(*RogueCommonDialogueOptionBattleResultInfo)(nil), // 1: RogueCommonDialogueOptionBattleResultInfo
	(*OJAHMDLNMAK)(nil),                               // 2: OJAHMDLNMAK
}
var file_RogueCommonDialogueOptionResultInfo_proto_depIdxs = []int32{
	1, // 0: RogueCommonDialogueOptionResultInfo.battle_result_info:type_name -> RogueCommonDialogueOptionBattleResultInfo
	2, // 1: RogueCommonDialogueOptionResultInfo.ACHLGHPPMNO:type_name -> OJAHMDLNMAK
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RogueCommonDialogueOptionResultInfo_proto_init() }
func file_RogueCommonDialogueOptionResultInfo_proto_init() {
	if File_RogueCommonDialogueOptionResultInfo_proto != nil {
		return
	}
	file_OJAHMDLNMAK_proto_init()
	file_RogueCommonDialogueOptionBattleResultInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueCommonDialogueOptionResultInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueCommonDialogueOptionResultInfo); i {
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
	file_RogueCommonDialogueOptionResultInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RogueCommonDialogueOptionResultInfo_BattleResultInfo)(nil),
		(*RogueCommonDialogueOptionResultInfo_ACHLGHPPMNO)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueCommonDialogueOptionResultInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueCommonDialogueOptionResultInfo_proto_goTypes,
		DependencyIndexes: file_RogueCommonDialogueOptionResultInfo_proto_depIdxs,
		MessageInfos:      file_RogueCommonDialogueOptionResultInfo_proto_msgTypes,
	}.Build()
	File_RogueCommonDialogueOptionResultInfo_proto = out.File
	file_RogueCommonDialogueOptionResultInfo_proto_rawDesc = nil
	file_RogueCommonDialogueOptionResultInfo_proto_goTypes = nil
	file_RogueCommonDialogueOptionResultInfo_proto_depIdxs = nil
}
