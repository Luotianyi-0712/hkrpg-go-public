// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MHBCFFNEBEL.proto

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

type MHBCFFNEBEL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JHINDCDBNML []uint32                    `protobuf:"varint,14,rep,packed,name=JHINDCDBNML,proto3" json:"JHINDCDBNML,omitempty"`
	Type        SwordTrainingDailyPhaseType `protobuf:"varint,11,opt,name=type,proto3,enum=SwordTrainingDailyPhaseType" json:"type,omitempty"`
	LOJFKFFCEAA []*JDNNEFNIEGL              `protobuf:"bytes,10,rep,name=LOJFKFFCEAA,proto3" json:"LOJFKFFCEAA,omitempty"`
}

func (x *MHBCFFNEBEL) Reset() {
	*x = MHBCFFNEBEL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MHBCFFNEBEL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MHBCFFNEBEL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MHBCFFNEBEL) ProtoMessage() {}

func (x *MHBCFFNEBEL) ProtoReflect() protoreflect.Message {
	mi := &file_MHBCFFNEBEL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MHBCFFNEBEL.ProtoReflect.Descriptor instead.
func (*MHBCFFNEBEL) Descriptor() ([]byte, []int) {
	return file_MHBCFFNEBEL_proto_rawDescGZIP(), []int{0}
}

func (x *MHBCFFNEBEL) GetJHINDCDBNML() []uint32 {
	if x != nil {
		return x.JHINDCDBNML
	}
	return nil
}

func (x *MHBCFFNEBEL) GetType() SwordTrainingDailyPhaseType {
	if x != nil {
		return x.Type
	}
	return SwordTrainingDailyPhaseType_SWORD_TRAINING_DAILY_PHASE_TYPE_NONE
}

func (x *MHBCFFNEBEL) GetLOJFKFFCEAA() []*JDNNEFNIEGL {
	if x != nil {
		return x.LOJFKFFCEAA
	}
	return nil
}

var File_MHBCFFNEBEL_proto protoreflect.FileDescriptor

var file_MHBCFFNEBEL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x48, 0x42, 0x43, 0x46, 0x46, 0x4e, 0x45, 0x42, 0x45, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x44, 0x4e, 0x4e, 0x45, 0x46, 0x4e, 0x49, 0x45, 0x47, 0x4c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x68, 0x61, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0b, 0x4d, 0x48,
	0x42, 0x43, 0x46, 0x46, 0x4e, 0x45, 0x42, 0x45, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x48, 0x49,
	0x4e, 0x44, 0x43, 0x44, 0x42, 0x4e, 0x4d, 0x4c, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b,
	0x4a, 0x48, 0x49, 0x4e, 0x44, 0x43, 0x44, 0x42, 0x4e, 0x4d, 0x4c, 0x12, 0x30, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x53, 0x77, 0x6f, 0x72,
	0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x68,
	0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a,
	0x0b, 0x4c, 0x4f, 0x4a, 0x46, 0x4b, 0x46, 0x46, 0x43, 0x45, 0x41, 0x41, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x44, 0x4e, 0x4e, 0x45, 0x46, 0x4e, 0x49, 0x45, 0x47, 0x4c,
	0x52, 0x0b, 0x4c, 0x4f, 0x4a, 0x46, 0x4b, 0x46, 0x46, 0x43, 0x45, 0x41, 0x41, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MHBCFFNEBEL_proto_rawDescOnce sync.Once
	file_MHBCFFNEBEL_proto_rawDescData = file_MHBCFFNEBEL_proto_rawDesc
)

func file_MHBCFFNEBEL_proto_rawDescGZIP() []byte {
	file_MHBCFFNEBEL_proto_rawDescOnce.Do(func() {
		file_MHBCFFNEBEL_proto_rawDescData = protoimpl.X.CompressGZIP(file_MHBCFFNEBEL_proto_rawDescData)
	})
	return file_MHBCFFNEBEL_proto_rawDescData
}

var file_MHBCFFNEBEL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MHBCFFNEBEL_proto_goTypes = []interface{}{
	(*MHBCFFNEBEL)(nil),              // 0: MHBCFFNEBEL
	(SwordTrainingDailyPhaseType)(0), // 1: SwordTrainingDailyPhaseType
	(*JDNNEFNIEGL)(nil),              // 2: JDNNEFNIEGL
}
var file_MHBCFFNEBEL_proto_depIdxs = []int32{
	1, // 0: MHBCFFNEBEL.type:type_name -> SwordTrainingDailyPhaseType
	2, // 1: MHBCFFNEBEL.LOJFKFFCEAA:type_name -> JDNNEFNIEGL
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MHBCFFNEBEL_proto_init() }
func file_MHBCFFNEBEL_proto_init() {
	if File_MHBCFFNEBEL_proto != nil {
		return
	}
	file_JDNNEFNIEGL_proto_init()
	file_SwordTrainingDailyPhaseType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MHBCFFNEBEL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MHBCFFNEBEL); i {
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
			RawDescriptor: file_MHBCFFNEBEL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MHBCFFNEBEL_proto_goTypes,
		DependencyIndexes: file_MHBCFFNEBEL_proto_depIdxs,
		MessageInfos:      file_MHBCFFNEBEL_proto_msgTypes,
	}.Build()
	File_MHBCFFNEBEL_proto = out.File
	file_MHBCFFNEBEL_proto_rawDesc = nil
	file_MHBCFFNEBEL_proto_goTypes = nil
	file_MHBCFFNEBEL_proto_depIdxs = nil
}
