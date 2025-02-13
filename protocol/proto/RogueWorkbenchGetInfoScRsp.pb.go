// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueWorkbenchGetInfoScRsp.proto

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

type RogueWorkbenchGetInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32                  `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	IGJKECEIHAI map[uint32]*CMPBKIENJNF `protobuf:"bytes,2,rep,name=IGJKECEIHAI,proto3" json:"IGJKECEIHAI,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RogueWorkbenchGetInfoScRsp) Reset() {
	*x = RogueWorkbenchGetInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueWorkbenchGetInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueWorkbenchGetInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueWorkbenchGetInfoScRsp) ProtoMessage() {}

func (x *RogueWorkbenchGetInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_RogueWorkbenchGetInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueWorkbenchGetInfoScRsp.ProtoReflect.Descriptor instead.
func (*RogueWorkbenchGetInfoScRsp) Descriptor() ([]byte, []int) {
	return file_RogueWorkbenchGetInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *RogueWorkbenchGetInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *RogueWorkbenchGetInfoScRsp) GetIGJKECEIHAI() map[uint32]*CMPBKIENJNF {
	if x != nil {
		return x.IGJKECEIHAI
	}
	return nil
}

var File_RogueWorkbenchGetInfoScRsp_proto protoreflect.FileDescriptor

var file_RogueWorkbenchGetInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4d, 0x50, 0x42, 0x4b, 0x49, 0x45, 0x4e, 0x4a, 0x4e, 0x46, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x4e,
	0x0a, 0x0b, 0x49, 0x47, 0x4a, 0x4b, 0x45, 0x43, 0x45, 0x49, 0x48, 0x41, 0x49, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x62,
	0x65, 0x6e, 0x63, 0x68, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x2e, 0x49, 0x47, 0x4a, 0x4b, 0x45, 0x43, 0x45, 0x49, 0x48, 0x41, 0x49, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x49, 0x47, 0x4a, 0x4b, 0x45, 0x43, 0x45, 0x49, 0x48, 0x41, 0x49, 0x1a, 0x4c,
	0x0a, 0x10, 0x49, 0x47, 0x4a, 0x4b, 0x45, 0x43, 0x45, 0x49, 0x48, 0x41, 0x49, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x4d, 0x50, 0x42, 0x4b, 0x49, 0x45, 0x4e, 0x4a, 0x4e,
	0x46, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueWorkbenchGetInfoScRsp_proto_rawDescOnce sync.Once
	file_RogueWorkbenchGetInfoScRsp_proto_rawDescData = file_RogueWorkbenchGetInfoScRsp_proto_rawDesc
)

func file_RogueWorkbenchGetInfoScRsp_proto_rawDescGZIP() []byte {
	file_RogueWorkbenchGetInfoScRsp_proto_rawDescOnce.Do(func() {
		file_RogueWorkbenchGetInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueWorkbenchGetInfoScRsp_proto_rawDescData)
	})
	return file_RogueWorkbenchGetInfoScRsp_proto_rawDescData
}

var file_RogueWorkbenchGetInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RogueWorkbenchGetInfoScRsp_proto_goTypes = []interface{}{
	(*RogueWorkbenchGetInfoScRsp)(nil), // 0: RogueWorkbenchGetInfoScRsp
	nil,                                // 1: RogueWorkbenchGetInfoScRsp.IGJKECEIHAIEntry
	(*CMPBKIENJNF)(nil),                // 2: CMPBKIENJNF
}
var file_RogueWorkbenchGetInfoScRsp_proto_depIdxs = []int32{
	1, // 0: RogueWorkbenchGetInfoScRsp.IGJKECEIHAI:type_name -> RogueWorkbenchGetInfoScRsp.IGJKECEIHAIEntry
	2, // 1: RogueWorkbenchGetInfoScRsp.IGJKECEIHAIEntry.value:type_name -> CMPBKIENJNF
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RogueWorkbenchGetInfoScRsp_proto_init() }
func file_RogueWorkbenchGetInfoScRsp_proto_init() {
	if File_RogueWorkbenchGetInfoScRsp_proto != nil {
		return
	}
	file_CMPBKIENJNF_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueWorkbenchGetInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueWorkbenchGetInfoScRsp); i {
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
			RawDescriptor: file_RogueWorkbenchGetInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueWorkbenchGetInfoScRsp_proto_goTypes,
		DependencyIndexes: file_RogueWorkbenchGetInfoScRsp_proto_depIdxs,
		MessageInfos:      file_RogueWorkbenchGetInfoScRsp_proto_msgTypes,
	}.Build()
	File_RogueWorkbenchGetInfoScRsp_proto = out.File
	file_RogueWorkbenchGetInfoScRsp_proto_rawDesc = nil
	file_RogueWorkbenchGetInfoScRsp_proto_goTypes = nil
	file_RogueWorkbenchGetInfoScRsp_proto_depIdxs = nil
}
