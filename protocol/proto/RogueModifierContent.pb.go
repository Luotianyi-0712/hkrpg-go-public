// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueModifierContent.proto

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

type RogueModifierContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentModifierEffectId uint32                   `protobuf:"varint,5,opt,name=content_modifier_effect_id,json=contentModifierEffectId,proto3" json:"content_modifier_effect_id,omitempty"`
	ModifierContentType     RogueModifierContentType `protobuf:"varint,12,opt,name=modifier_content_type,json=modifierContentType,proto3,enum=RogueModifierContentType" json:"modifier_content_type,omitempty"`
	BKIEKCHKBHH             uint32                   `protobuf:"varint,13,opt,name=BKIEKCHKBHH,proto3" json:"BKIEKCHKBHH,omitempty"`
}

func (x *RogueModifierContent) Reset() {
	*x = RogueModifierContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueModifierContent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueModifierContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueModifierContent) ProtoMessage() {}

func (x *RogueModifierContent) ProtoReflect() protoreflect.Message {
	mi := &file_RogueModifierContent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueModifierContent.ProtoReflect.Descriptor instead.
func (*RogueModifierContent) Descriptor() ([]byte, []int) {
	return file_RogueModifierContent_proto_rawDescGZIP(), []int{0}
}

func (x *RogueModifierContent) GetContentModifierEffectId() uint32 {
	if x != nil {
		return x.ContentModifierEffectId
	}
	return 0
}

func (x *RogueModifierContent) GetModifierContentType() RogueModifierContentType {
	if x != nil {
		return x.ModifierContentType
	}
	return RogueModifierContentType_ROGUE_MODIFIER_CONTENT_DEFINITE
}

func (x *RogueModifierContent) GetBKIEKCHKBHH() uint32 {
	if x != nil {
		return x.BKIEKCHKBHH
	}
	return 0
}

var File_RogueModifierContent_proto protoreflect.FileDescriptor

var file_RogueModifierContent_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a,
	0x14, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x4d, 0x0a, 0x15, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x4b, 0x49, 0x45, 0x4b, 0x43, 0x48, 0x4b, 0x42, 0x48, 0x48,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x4b, 0x49, 0x45, 0x4b, 0x43, 0x48, 0x4b,
	0x42, 0x48, 0x48, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueModifierContent_proto_rawDescOnce sync.Once
	file_RogueModifierContent_proto_rawDescData = file_RogueModifierContent_proto_rawDesc
)

func file_RogueModifierContent_proto_rawDescGZIP() []byte {
	file_RogueModifierContent_proto_rawDescOnce.Do(func() {
		file_RogueModifierContent_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueModifierContent_proto_rawDescData)
	})
	return file_RogueModifierContent_proto_rawDescData
}

var file_RogueModifierContent_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueModifierContent_proto_goTypes = []interface{}{
	(*RogueModifierContent)(nil),  // 0: RogueModifierContent
	(RogueModifierContentType)(0), // 1: RogueModifierContentType
}
var file_RogueModifierContent_proto_depIdxs = []int32{
	1, // 0: RogueModifierContent.modifier_content_type:type_name -> RogueModifierContentType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueModifierContent_proto_init() }
func file_RogueModifierContent_proto_init() {
	if File_RogueModifierContent_proto != nil {
		return
	}
	file_RogueModifierContentType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueModifierContent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueModifierContent); i {
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
			RawDescriptor: file_RogueModifierContent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueModifierContent_proto_goTypes,
		DependencyIndexes: file_RogueModifierContent_proto_depIdxs,
		MessageInfos:      file_RogueModifierContent_proto_msgTypes,
	}.Build()
	File_RogueModifierContent_proto = out.File
	file_RogueModifierContent_proto_rawDesc = nil
	file_RogueModifierContent_proto_goTypes = nil
	file_RogueModifierContent_proto_depIdxs = nil
}
