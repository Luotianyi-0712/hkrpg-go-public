// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LobbyModifyPlayerInfoCsReq.proto

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

type LobbyModifyPlayerInfoCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        LobbyModifyType `protobuf:"varint,8,opt,name=type,proto3,enum=LobbyModifyType" json:"type,omitempty"`
	FMIJGJANHFI uint32          `protobuf:"varint,9,opt,name=FMIJGJANHFI,proto3" json:"FMIJGJANHFI,omitempty"`
	MHDJJGDLLIG *JBEBLKIKGMP    `protobuf:"bytes,3,opt,name=MHDJJGDLLIG,proto3" json:"MHDJJGDLLIG,omitempty"`
}

func (x *LobbyModifyPlayerInfoCsReq) Reset() {
	*x = LobbyModifyPlayerInfoCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LobbyModifyPlayerInfoCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LobbyModifyPlayerInfoCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyModifyPlayerInfoCsReq) ProtoMessage() {}

func (x *LobbyModifyPlayerInfoCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_LobbyModifyPlayerInfoCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyModifyPlayerInfoCsReq.ProtoReflect.Descriptor instead.
func (*LobbyModifyPlayerInfoCsReq) Descriptor() ([]byte, []int) {
	return file_LobbyModifyPlayerInfoCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *LobbyModifyPlayerInfoCsReq) GetType() LobbyModifyType {
	if x != nil {
		return x.Type
	}
	return LobbyModifyType_LobbyModifyType_None
}

func (x *LobbyModifyPlayerInfoCsReq) GetFMIJGJANHFI() uint32 {
	if x != nil {
		return x.FMIJGJANHFI
	}
	return 0
}

func (x *LobbyModifyPlayerInfoCsReq) GetMHDJJGDLLIG() *JBEBLKIKGMP {
	if x != nil {
		return x.MHDJJGDLLIG
	}
	return nil
}

var File_LobbyModifyPlayerInfoCsReq_proto protoreflect.FileDescriptor

var file_LobbyModifyPlayerInfoCsReq_proto_rawDesc = []byte{
	0x0a, 0x20, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x42, 0x45, 0x42, 0x4c,
	0x4b, 0x49, 0x4b, 0x47, 0x4d, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a,
	0x1a, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x4c, 0x6f, 0x62, 0x62,
	0x79, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4d, 0x49, 0x4a, 0x47, 0x4a, 0x41, 0x4e, 0x48, 0x46, 0x49,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x4d, 0x49, 0x4a, 0x47, 0x4a, 0x41, 0x4e,
	0x48, 0x46, 0x49, 0x12, 0x2e, 0x0a, 0x0b, 0x4d, 0x48, 0x44, 0x4a, 0x4a, 0x47, 0x44, 0x4c, 0x4c,
	0x49, 0x47, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x42, 0x45, 0x42, 0x4c,
	0x4b, 0x49, 0x4b, 0x47, 0x4d, 0x50, 0x52, 0x0b, 0x4d, 0x48, 0x44, 0x4a, 0x4a, 0x47, 0x44, 0x4c,
	0x4c, 0x49, 0x47, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LobbyModifyPlayerInfoCsReq_proto_rawDescOnce sync.Once
	file_LobbyModifyPlayerInfoCsReq_proto_rawDescData = file_LobbyModifyPlayerInfoCsReq_proto_rawDesc
)

func file_LobbyModifyPlayerInfoCsReq_proto_rawDescGZIP() []byte {
	file_LobbyModifyPlayerInfoCsReq_proto_rawDescOnce.Do(func() {
		file_LobbyModifyPlayerInfoCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_LobbyModifyPlayerInfoCsReq_proto_rawDescData)
	})
	return file_LobbyModifyPlayerInfoCsReq_proto_rawDescData
}

var file_LobbyModifyPlayerInfoCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LobbyModifyPlayerInfoCsReq_proto_goTypes = []interface{}{
	(*LobbyModifyPlayerInfoCsReq)(nil), // 0: LobbyModifyPlayerInfoCsReq
	(LobbyModifyType)(0),               // 1: LobbyModifyType
	(*JBEBLKIKGMP)(nil),                // 2: JBEBLKIKGMP
}
var file_LobbyModifyPlayerInfoCsReq_proto_depIdxs = []int32{
	1, // 0: LobbyModifyPlayerInfoCsReq.type:type_name -> LobbyModifyType
	2, // 1: LobbyModifyPlayerInfoCsReq.MHDJJGDLLIG:type_name -> JBEBLKIKGMP
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_LobbyModifyPlayerInfoCsReq_proto_init() }
func file_LobbyModifyPlayerInfoCsReq_proto_init() {
	if File_LobbyModifyPlayerInfoCsReq_proto != nil {
		return
	}
	file_LobbyModifyType_proto_init()
	file_JBEBLKIKGMP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LobbyModifyPlayerInfoCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LobbyModifyPlayerInfoCsReq); i {
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
			RawDescriptor: file_LobbyModifyPlayerInfoCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LobbyModifyPlayerInfoCsReq_proto_goTypes,
		DependencyIndexes: file_LobbyModifyPlayerInfoCsReq_proto_depIdxs,
		MessageInfos:      file_LobbyModifyPlayerInfoCsReq_proto_msgTypes,
	}.Build()
	File_LobbyModifyPlayerInfoCsReq_proto = out.File
	file_LobbyModifyPlayerInfoCsReq_proto_rawDesc = nil
	file_LobbyModifyPlayerInfoCsReq_proto_goTypes = nil
	file_LobbyModifyPlayerInfoCsReq_proto_depIdxs = nil
}
