// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueCommonFormula.proto

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

type RogueCommonFormula struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormulaInfo *FormulaInfo `protobuf:"bytes,4,opt,name=formula_info,json=formulaInfo,proto3" json:"formula_info,omitempty"`
}

func (x *RogueCommonFormula) Reset() {
	*x = RogueCommonFormula{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueCommonFormula_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueCommonFormula) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueCommonFormula) ProtoMessage() {}

func (x *RogueCommonFormula) ProtoReflect() protoreflect.Message {
	mi := &file_RogueCommonFormula_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueCommonFormula.ProtoReflect.Descriptor instead.
func (*RogueCommonFormula) Descriptor() ([]byte, []int) {
	return file_RogueCommonFormula_proto_rawDescGZIP(), []int{0}
}

func (x *RogueCommonFormula) GetFormulaInfo() *FormulaInfo {
	if x != nil {
		return x.FormulaInfo
	}
	return nil
}

var File_RogueCommonFormula_proto protoreflect.FileDescriptor

var file_RogueCommonFormula_proto_rawDesc = []byte{
	0x0a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x46, 0x6f, 0x72,
	0x6d, 0x75, 0x6c, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x6f, 0x72, 0x6d,
	0x75, 0x6c, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a,
	0x12, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d,
	0x75, 0x6c, 0x61, 0x12, 0x2f, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x6f, 0x72, 0x6d,
	0x75, 0x6c, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueCommonFormula_proto_rawDescOnce sync.Once
	file_RogueCommonFormula_proto_rawDescData = file_RogueCommonFormula_proto_rawDesc
)

func file_RogueCommonFormula_proto_rawDescGZIP() []byte {
	file_RogueCommonFormula_proto_rawDescOnce.Do(func() {
		file_RogueCommonFormula_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueCommonFormula_proto_rawDescData)
	})
	return file_RogueCommonFormula_proto_rawDescData
}

var file_RogueCommonFormula_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueCommonFormula_proto_goTypes = []interface{}{
	(*RogueCommonFormula)(nil), // 0: RogueCommonFormula
	(*FormulaInfo)(nil),        // 1: FormulaInfo
}
var file_RogueCommonFormula_proto_depIdxs = []int32{
	1, // 0: RogueCommonFormula.formula_info:type_name -> FormulaInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueCommonFormula_proto_init() }
func file_RogueCommonFormula_proto_init() {
	if File_RogueCommonFormula_proto != nil {
		return
	}
	file_FormulaInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueCommonFormula_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueCommonFormula); i {
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
			RawDescriptor: file_RogueCommonFormula_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueCommonFormula_proto_goTypes,
		DependencyIndexes: file_RogueCommonFormula_proto_depIdxs,
		MessageInfos:      file_RogueCommonFormula_proto_msgTypes,
	}.Build()
	File_RogueCommonFormula_proto = out.File
	file_RogueCommonFormula_proto_rawDesc = nil
	file_RogueCommonFormula_proto_goTypes = nil
	file_RogueCommonFormula_proto_depIdxs = nil
}
