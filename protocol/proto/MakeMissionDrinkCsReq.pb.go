// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MakeMissionDrinkCsReq.proto

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

type MakeMissionDrinkCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FCLAFDMNGHB uint32       `protobuf:"varint,3,opt,name=FCLAFDMNGHB,proto3" json:"FCLAFDMNGHB,omitempty"`
	IsSaveData  bool         `protobuf:"varint,2,opt,name=is_save_data,json=isSaveData,proto3" json:"is_save_data,omitempty"`
	LIKBNKOPFIF *EICPBAEMNEC `protobuf:"bytes,9,opt,name=LIKBNKOPFIF,proto3" json:"LIKBNKOPFIF,omitempty"`
}

func (x *MakeMissionDrinkCsReq) Reset() {
	*x = MakeMissionDrinkCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MakeMissionDrinkCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeMissionDrinkCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeMissionDrinkCsReq) ProtoMessage() {}

func (x *MakeMissionDrinkCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_MakeMissionDrinkCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeMissionDrinkCsReq.ProtoReflect.Descriptor instead.
func (*MakeMissionDrinkCsReq) Descriptor() ([]byte, []int) {
	return file_MakeMissionDrinkCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *MakeMissionDrinkCsReq) GetFCLAFDMNGHB() uint32 {
	if x != nil {
		return x.FCLAFDMNGHB
	}
	return 0
}

func (x *MakeMissionDrinkCsReq) GetIsSaveData() bool {
	if x != nil {
		return x.IsSaveData
	}
	return false
}

func (x *MakeMissionDrinkCsReq) GetLIKBNKOPFIF() *EICPBAEMNEC {
	if x != nil {
		return x.LIKBNKOPFIF
	}
	return nil
}

var File_MakeMissionDrinkCsReq_proto protoreflect.FileDescriptor

var file_MakeMissionDrinkCsReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x4d, 0x61, 0x6b, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x72, 0x69,
	0x6e, 0x6b, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45,
	0x49, 0x43, 0x50, 0x42, 0x41, 0x45, 0x4d, 0x4e, 0x45, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8b, 0x01, 0x0a, 0x15, 0x4d, 0x61, 0x6b, 0x65, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x44, 0x72, 0x69, 0x6e, 0x6b, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x43,
	0x4c, 0x41, 0x46, 0x44, 0x4d, 0x4e, 0x47, 0x48, 0x42, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x46, 0x43, 0x4c, 0x41, 0x46, 0x44, 0x4d, 0x4e, 0x47, 0x48, 0x42, 0x12, 0x20, 0x0a, 0x0c,
	0x69, 0x73, 0x5f, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e,
	0x0a, 0x0b, 0x4c, 0x49, 0x4b, 0x42, 0x4e, 0x4b, 0x4f, 0x50, 0x46, 0x49, 0x46, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x49, 0x43, 0x50, 0x42, 0x41, 0x45, 0x4d, 0x4e, 0x45,
	0x43, 0x52, 0x0b, 0x4c, 0x49, 0x4b, 0x42, 0x4e, 0x4b, 0x4f, 0x50, 0x46, 0x49, 0x46, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MakeMissionDrinkCsReq_proto_rawDescOnce sync.Once
	file_MakeMissionDrinkCsReq_proto_rawDescData = file_MakeMissionDrinkCsReq_proto_rawDesc
)

func file_MakeMissionDrinkCsReq_proto_rawDescGZIP() []byte {
	file_MakeMissionDrinkCsReq_proto_rawDescOnce.Do(func() {
		file_MakeMissionDrinkCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_MakeMissionDrinkCsReq_proto_rawDescData)
	})
	return file_MakeMissionDrinkCsReq_proto_rawDescData
}

var file_MakeMissionDrinkCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MakeMissionDrinkCsReq_proto_goTypes = []interface{}{
	(*MakeMissionDrinkCsReq)(nil), // 0: MakeMissionDrinkCsReq
	(*EICPBAEMNEC)(nil),           // 1: EICPBAEMNEC
}
var file_MakeMissionDrinkCsReq_proto_depIdxs = []int32{
	1, // 0: MakeMissionDrinkCsReq.LIKBNKOPFIF:type_name -> EICPBAEMNEC
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MakeMissionDrinkCsReq_proto_init() }
func file_MakeMissionDrinkCsReq_proto_init() {
	if File_MakeMissionDrinkCsReq_proto != nil {
		return
	}
	file_EICPBAEMNEC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MakeMissionDrinkCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeMissionDrinkCsReq); i {
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
			RawDescriptor: file_MakeMissionDrinkCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MakeMissionDrinkCsReq_proto_goTypes,
		DependencyIndexes: file_MakeMissionDrinkCsReq_proto_depIdxs,
		MessageInfos:      file_MakeMissionDrinkCsReq_proto_msgTypes,
	}.Build()
	File_MakeMissionDrinkCsReq_proto = out.File
	file_MakeMissionDrinkCsReq_proto_rawDesc = nil
	file_MakeMissionDrinkCsReq_proto_goTypes = nil
	file_MakeMissionDrinkCsReq_proto_depIdxs = nil
}
