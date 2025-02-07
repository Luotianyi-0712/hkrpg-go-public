// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SpaceZooBornScRsp.proto

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

type SpaceZooBornScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BPEGLOJENEN *KMEDPBBAOHC   `protobuf:"bytes,8,opt,name=BPEGLOJENEN,proto3" json:"BPEGLOJENEN,omitempty"`
	NNICHBFOECC []*OOHKNIOIIGC `protobuf:"bytes,13,rep,name=NNICHBFOECC,proto3" json:"NNICHBFOECC,omitempty"`
	LOJOEDFDIEE bool           `protobuf:"varint,12,opt,name=LOJOEDFDIEE,proto3" json:"LOJOEDFDIEE,omitempty"`
	Retcode     uint32         `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SpaceZooBornScRsp) Reset() {
	*x = SpaceZooBornScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SpaceZooBornScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceZooBornScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceZooBornScRsp) ProtoMessage() {}

func (x *SpaceZooBornScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SpaceZooBornScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceZooBornScRsp.ProtoReflect.Descriptor instead.
func (*SpaceZooBornScRsp) Descriptor() ([]byte, []int) {
	return file_SpaceZooBornScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SpaceZooBornScRsp) GetBPEGLOJENEN() *KMEDPBBAOHC {
	if x != nil {
		return x.BPEGLOJENEN
	}
	return nil
}

func (x *SpaceZooBornScRsp) GetNNICHBFOECC() []*OOHKNIOIIGC {
	if x != nil {
		return x.NNICHBFOECC
	}
	return nil
}

func (x *SpaceZooBornScRsp) GetLOJOEDFDIEE() bool {
	if x != nil {
		return x.LOJOEDFDIEE
	}
	return false
}

func (x *SpaceZooBornScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SpaceZooBornScRsp_proto protoreflect.FileDescriptor

var file_SpaceZooBornScRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x53, 0x70, 0x61, 0x63, 0x65, 0x5a, 0x6f, 0x6f, 0x42, 0x6f, 0x72, 0x6e, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4f, 0x48, 0x4b, 0x4e,
	0x49, 0x4f, 0x49, 0x49, 0x47, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4d,
	0x45, 0x44, 0x50, 0x42, 0x42, 0x41, 0x4f, 0x48, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xaf, 0x01, 0x0a, 0x11, 0x53, 0x70, 0x61, 0x63, 0x65, 0x5a, 0x6f, 0x6f, 0x42, 0x6f, 0x72, 0x6e,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x42, 0x50, 0x45, 0x47, 0x4c, 0x4f, 0x4a,
	0x45, 0x4e, 0x45, 0x4e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4d, 0x45,
	0x44, 0x50, 0x42, 0x42, 0x41, 0x4f, 0x48, 0x43, 0x52, 0x0b, 0x42, 0x50, 0x45, 0x47, 0x4c, 0x4f,
	0x4a, 0x45, 0x4e, 0x45, 0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x4e, 0x4e, 0x49, 0x43, 0x48, 0x42, 0x46,
	0x4f, 0x45, 0x43, 0x43, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4f, 0x48,
	0x4b, 0x4e, 0x49, 0x4f, 0x49, 0x49, 0x47, 0x43, 0x52, 0x0b, 0x4e, 0x4e, 0x49, 0x43, 0x48, 0x42,
	0x46, 0x4f, 0x45, 0x43, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4f, 0x4a, 0x4f, 0x45, 0x44, 0x46,
	0x44, 0x49, 0x45, 0x45, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x4f, 0x4a, 0x4f,
	0x45, 0x44, 0x46, 0x44, 0x49, 0x45, 0x45, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SpaceZooBornScRsp_proto_rawDescOnce sync.Once
	file_SpaceZooBornScRsp_proto_rawDescData = file_SpaceZooBornScRsp_proto_rawDesc
)

func file_SpaceZooBornScRsp_proto_rawDescGZIP() []byte {
	file_SpaceZooBornScRsp_proto_rawDescOnce.Do(func() {
		file_SpaceZooBornScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SpaceZooBornScRsp_proto_rawDescData)
	})
	return file_SpaceZooBornScRsp_proto_rawDescData
}

var file_SpaceZooBornScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SpaceZooBornScRsp_proto_goTypes = []interface{}{
	(*SpaceZooBornScRsp)(nil), // 0: SpaceZooBornScRsp
	(*KMEDPBBAOHC)(nil),       // 1: KMEDPBBAOHC
	(*OOHKNIOIIGC)(nil),       // 2: OOHKNIOIIGC
}
var file_SpaceZooBornScRsp_proto_depIdxs = []int32{
	1, // 0: SpaceZooBornScRsp.BPEGLOJENEN:type_name -> KMEDPBBAOHC
	2, // 1: SpaceZooBornScRsp.NNICHBFOECC:type_name -> OOHKNIOIIGC
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SpaceZooBornScRsp_proto_init() }
func file_SpaceZooBornScRsp_proto_init() {
	if File_SpaceZooBornScRsp_proto != nil {
		return
	}
	file_OOHKNIOIIGC_proto_init()
	file_KMEDPBBAOHC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SpaceZooBornScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceZooBornScRsp); i {
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
			RawDescriptor: file_SpaceZooBornScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SpaceZooBornScRsp_proto_goTypes,
		DependencyIndexes: file_SpaceZooBornScRsp_proto_depIdxs,
		MessageInfos:      file_SpaceZooBornScRsp_proto_msgTypes,
	}.Build()
	File_SpaceZooBornScRsp_proto = out.File
	file_SpaceZooBornScRsp_proto_rawDesc = nil
	file_SpaceZooBornScRsp_proto_goTypes = nil
	file_SpaceZooBornScRsp_proto_depIdxs = nil
}
