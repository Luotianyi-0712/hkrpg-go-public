// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MultiplayerFightGameStateScRsp.proto

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

type MultiplayerFightGameStateScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AOHJJJHMGHM *PEGAEGMBJBI   `protobuf:"bytes,13,opt,name=AOHJJJHMGHM,proto3" json:"AOHJJJHMGHM,omitempty"`
	DOLKBPLJGNO []*NJMGIJBHCPE `protobuf:"bytes,2,rep,name=DOLKBPLJGNO,proto3" json:"DOLKBPLJGNO,omitempty"`
	Retcode     uint32         `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *MultiplayerFightGameStateScRsp) Reset() {
	*x = MultiplayerFightGameStateScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MultiplayerFightGameStateScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplayerFightGameStateScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplayerFightGameStateScRsp) ProtoMessage() {}

func (x *MultiplayerFightGameStateScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_MultiplayerFightGameStateScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplayerFightGameStateScRsp.ProtoReflect.Descriptor instead.
func (*MultiplayerFightGameStateScRsp) Descriptor() ([]byte, []int) {
	return file_MultiplayerFightGameStateScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *MultiplayerFightGameStateScRsp) GetAOHJJJHMGHM() *PEGAEGMBJBI {
	if x != nil {
		return x.AOHJJJHMGHM
	}
	return nil
}

func (x *MultiplayerFightGameStateScRsp) GetDOLKBPLJGNO() []*NJMGIJBHCPE {
	if x != nil {
		return x.DOLKBPLJGNO
	}
	return nil
}

func (x *MultiplayerFightGameStateScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_MultiplayerFightGameStateScRsp_proto protoreflect.FileDescriptor

var file_MultiplayerFightGameStateScRsp_proto_rawDesc = []byte{
	0x0a, 0x24, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x67,
	0x68, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x45, 0x47, 0x41, 0x45, 0x47, 0x4d, 0x42,
	0x4a, 0x42, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x4a, 0x4d, 0x47, 0x49,
	0x4a, 0x42, 0x48, 0x43, 0x50, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a,
	0x1e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x67, 0x68,
	0x74, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12,
	0x2e, 0x0a, 0x0b, 0x41, 0x4f, 0x48, 0x4a, 0x4a, 0x4a, 0x48, 0x4d, 0x47, 0x48, 0x4d, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x45, 0x47, 0x41, 0x45, 0x47, 0x4d, 0x42, 0x4a,
	0x42, 0x49, 0x52, 0x0b, 0x41, 0x4f, 0x48, 0x4a, 0x4a, 0x4a, 0x48, 0x4d, 0x47, 0x48, 0x4d, 0x12,
	0x2e, 0x0a, 0x0b, 0x44, 0x4f, 0x4c, 0x4b, 0x42, 0x50, 0x4c, 0x4a, 0x47, 0x4e, 0x4f, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x4a, 0x4d, 0x47, 0x49, 0x4a, 0x42, 0x48, 0x43,
	0x50, 0x45, 0x52, 0x0b, 0x44, 0x4f, 0x4c, 0x4b, 0x42, 0x50, 0x4c, 0x4a, 0x47, 0x4e, 0x4f, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MultiplayerFightGameStateScRsp_proto_rawDescOnce sync.Once
	file_MultiplayerFightGameStateScRsp_proto_rawDescData = file_MultiplayerFightGameStateScRsp_proto_rawDesc
)

func file_MultiplayerFightGameStateScRsp_proto_rawDescGZIP() []byte {
	file_MultiplayerFightGameStateScRsp_proto_rawDescOnce.Do(func() {
		file_MultiplayerFightGameStateScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_MultiplayerFightGameStateScRsp_proto_rawDescData)
	})
	return file_MultiplayerFightGameStateScRsp_proto_rawDescData
}

var file_MultiplayerFightGameStateScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MultiplayerFightGameStateScRsp_proto_goTypes = []interface{}{
	(*MultiplayerFightGameStateScRsp)(nil), // 0: MultiplayerFightGameStateScRsp
	(*PEGAEGMBJBI)(nil),                    // 1: PEGAEGMBJBI
	(*NJMGIJBHCPE)(nil),                    // 2: NJMGIJBHCPE
}
var file_MultiplayerFightGameStateScRsp_proto_depIdxs = []int32{
	1, // 0: MultiplayerFightGameStateScRsp.AOHJJJHMGHM:type_name -> PEGAEGMBJBI
	2, // 1: MultiplayerFightGameStateScRsp.DOLKBPLJGNO:type_name -> NJMGIJBHCPE
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MultiplayerFightGameStateScRsp_proto_init() }
func file_MultiplayerFightGameStateScRsp_proto_init() {
	if File_MultiplayerFightGameStateScRsp_proto != nil {
		return
	}
	file_PEGAEGMBJBI_proto_init()
	file_NJMGIJBHCPE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MultiplayerFightGameStateScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplayerFightGameStateScRsp); i {
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
			RawDescriptor: file_MultiplayerFightGameStateScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MultiplayerFightGameStateScRsp_proto_goTypes,
		DependencyIndexes: file_MultiplayerFightGameStateScRsp_proto_depIdxs,
		MessageInfos:      file_MultiplayerFightGameStateScRsp_proto_msgTypes,
	}.Build()
	File_MultiplayerFightGameStateScRsp_proto = out.File
	file_MultiplayerFightGameStateScRsp_proto_rawDesc = nil
	file_MultiplayerFightGameStateScRsp_proto_goTypes = nil
	file_MultiplayerFightGameStateScRsp_proto_depIdxs = nil
}
