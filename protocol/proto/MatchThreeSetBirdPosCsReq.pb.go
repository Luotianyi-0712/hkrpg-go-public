// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MatchThreeSetBirdPosCsReq.proto

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

type MatchThreeSetBirdPosCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos         uint32 `protobuf:"varint,11,opt,name=pos,proto3" json:"pos,omitempty"`
	JIELNNCBKOD uint32 `protobuf:"varint,3,opt,name=JIELNNCBKOD,proto3" json:"JIELNNCBKOD,omitempty"`
}

func (x *MatchThreeSetBirdPosCsReq) Reset() {
	*x = MatchThreeSetBirdPosCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MatchThreeSetBirdPosCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchThreeSetBirdPosCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchThreeSetBirdPosCsReq) ProtoMessage() {}

func (x *MatchThreeSetBirdPosCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_MatchThreeSetBirdPosCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchThreeSetBirdPosCsReq.ProtoReflect.Descriptor instead.
func (*MatchThreeSetBirdPosCsReq) Descriptor() ([]byte, []int) {
	return file_MatchThreeSetBirdPosCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *MatchThreeSetBirdPosCsReq) GetPos() uint32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

func (x *MatchThreeSetBirdPosCsReq) GetJIELNNCBKOD() uint32 {
	if x != nil {
		return x.JIELNNCBKOD
	}
	return 0
}

var File_MatchThreeSetBirdPosCsReq_proto protoreflect.FileDescriptor

var file_MatchThreeSetBirdPosCsReq_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x68, 0x72, 0x65, 0x65, 0x53, 0x65, 0x74, 0x42,
	0x69, 0x72, 0x64, 0x50, 0x6f, 0x73, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4f, 0x0a, 0x19, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x68, 0x72, 0x65, 0x65, 0x53,
	0x65, 0x74, 0x42, 0x69, 0x72, 0x64, 0x50, 0x6f, 0x73, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x6f, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x49, 0x45, 0x4c, 0x4e, 0x4e, 0x43, 0x42, 0x4b, 0x4f, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x49, 0x45, 0x4c, 0x4e, 0x4e, 0x43, 0x42, 0x4b,
	0x4f, 0x44, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MatchThreeSetBirdPosCsReq_proto_rawDescOnce sync.Once
	file_MatchThreeSetBirdPosCsReq_proto_rawDescData = file_MatchThreeSetBirdPosCsReq_proto_rawDesc
)

func file_MatchThreeSetBirdPosCsReq_proto_rawDescGZIP() []byte {
	file_MatchThreeSetBirdPosCsReq_proto_rawDescOnce.Do(func() {
		file_MatchThreeSetBirdPosCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_MatchThreeSetBirdPosCsReq_proto_rawDescData)
	})
	return file_MatchThreeSetBirdPosCsReq_proto_rawDescData
}

var file_MatchThreeSetBirdPosCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MatchThreeSetBirdPosCsReq_proto_goTypes = []interface{}{
	(*MatchThreeSetBirdPosCsReq)(nil), // 0: MatchThreeSetBirdPosCsReq
}
var file_MatchThreeSetBirdPosCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MatchThreeSetBirdPosCsReq_proto_init() }
func file_MatchThreeSetBirdPosCsReq_proto_init() {
	if File_MatchThreeSetBirdPosCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MatchThreeSetBirdPosCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchThreeSetBirdPosCsReq); i {
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
			RawDescriptor: file_MatchThreeSetBirdPosCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MatchThreeSetBirdPosCsReq_proto_goTypes,
		DependencyIndexes: file_MatchThreeSetBirdPosCsReq_proto_depIdxs,
		MessageInfos:      file_MatchThreeSetBirdPosCsReq_proto_msgTypes,
	}.Build()
	File_MatchThreeSetBirdPosCsReq_proto = out.File
	file_MatchThreeSetBirdPosCsReq_proto_rawDesc = nil
	file_MatchThreeSetBirdPosCsReq_proto_goTypes = nil
	file_MatchThreeSetBirdPosCsReq_proto_depIdxs = nil
}
