// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetAlleyInfoScRsp.proto

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

type GetAlleyInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FIFHMONIDBL *DMCFKOHNIHB      `protobuf:"bytes,6,opt,name=FIFHMONIDBL,proto3" json:"FIFHMONIDBL,omitempty"`
	KHOFMPHEFNC []*NABCIIFFDCE    `protobuf:"bytes,7,rep,name=KHOFMPHEFNC,proto3" json:"KHOFMPHEFNC,omitempty"`
	EBNNBEEGJFN uint32            `protobuf:"varint,12,opt,name=EBNNBEEGJFN,proto3" json:"EBNNBEEGJFN,omitempty"`
	KFHDOFNDNAH []uint32          `protobuf:"varint,1,rep,packed,name=KFHDOFNDNAH,proto3" json:"KFHDOFNDNAH,omitempty"`
	IGOGPLHKHDA *IKBABPKIKAF      `protobuf:"bytes,5,opt,name=IGOGPLHKHDA,proto3" json:"IGOGPLHKHDA,omitempty"`
	PCFBNFBDLAH map[uint32]uint32 `protobuf:"bytes,8,rep,name=PCFBNFBDLAH,proto3" json:"PCFBNFBDLAH,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	EIFJOOEEELK []uint32          `protobuf:"varint,14,rep,packed,name=EIFJOOEEELK,proto3" json:"EIFJOOEEELK,omitempty"`
	CDMHBJPKLJO *OHAJNAEBFCO      `protobuf:"bytes,3,opt,name=CDMHBJPKLJO,proto3" json:"CDMHBJPKLJO,omitempty"`
	Level       uint32            `protobuf:"varint,10,opt,name=level,proto3" json:"level,omitempty"`
	POLIIKOPAEB uint32            `protobuf:"varint,11,opt,name=POLIIKOPAEB,proto3" json:"POLIIKOPAEB,omitempty"`
	Retcode     uint32            `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	CCJNIIDBKCE []uint32          `protobuf:"varint,15,rep,packed,name=CCJNIIDBKCE,proto3" json:"CCJNIIDBKCE,omitempty"`
}

func (x *GetAlleyInfoScRsp) Reset() {
	*x = GetAlleyInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetAlleyInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlleyInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlleyInfoScRsp) ProtoMessage() {}

func (x *GetAlleyInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetAlleyInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlleyInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetAlleyInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetAlleyInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetAlleyInfoScRsp) GetFIFHMONIDBL() *DMCFKOHNIHB {
	if x != nil {
		return x.FIFHMONIDBL
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetKHOFMPHEFNC() []*NABCIIFFDCE {
	if x != nil {
		return x.KHOFMPHEFNC
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetEBNNBEEGJFN() uint32 {
	if x != nil {
		return x.EBNNBEEGJFN
	}
	return 0
}

func (x *GetAlleyInfoScRsp) GetKFHDOFNDNAH() []uint32 {
	if x != nil {
		return x.KFHDOFNDNAH
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetIGOGPLHKHDA() *IKBABPKIKAF {
	if x != nil {
		return x.IGOGPLHKHDA
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetPCFBNFBDLAH() map[uint32]uint32 {
	if x != nil {
		return x.PCFBNFBDLAH
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetEIFJOOEEELK() []uint32 {
	if x != nil {
		return x.EIFJOOEEELK
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetCDMHBJPKLJO() *OHAJNAEBFCO {
	if x != nil {
		return x.CDMHBJPKLJO
	}
	return nil
}

func (x *GetAlleyInfoScRsp) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *GetAlleyInfoScRsp) GetPOLIIKOPAEB() uint32 {
	if x != nil {
		return x.POLIIKOPAEB
	}
	return 0
}

func (x *GetAlleyInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetAlleyInfoScRsp) GetCCJNIIDBKCE() []uint32 {
	if x != nil {
		return x.CCJNIIDBKCE
	}
	return nil
}

var File_GetAlleyInfoScRsp_proto protoreflect.FileDescriptor

var file_GetAlleyInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x4d, 0x43, 0x46, 0x4b,
	0x4f, 0x48, 0x4e, 0x49, 0x48, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x41,
	0x42, 0x43, 0x49, 0x49, 0x46, 0x46, 0x44, 0x43, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x49, 0x4b, 0x42, 0x41, 0x42, 0x50, 0x4b, 0x49, 0x4b, 0x41, 0x46, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x48, 0x41, 0x4a, 0x4e, 0x41, 0x45, 0x42, 0x46, 0x43, 0x4f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x04, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x46,
	0x49, 0x46, 0x48, 0x4d, 0x4f, 0x4e, 0x49, 0x44, 0x42, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x44, 0x4d, 0x43, 0x46, 0x4b, 0x4f, 0x48, 0x4e, 0x49, 0x48, 0x42, 0x52, 0x0b,
	0x46, 0x49, 0x46, 0x48, 0x4d, 0x4f, 0x4e, 0x49, 0x44, 0x42, 0x4c, 0x12, 0x2e, 0x0a, 0x0b, 0x4b,
	0x48, 0x4f, 0x46, 0x4d, 0x50, 0x48, 0x45, 0x46, 0x4e, 0x43, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4e, 0x41, 0x42, 0x43, 0x49, 0x49, 0x46, 0x46, 0x44, 0x43, 0x45, 0x52, 0x0b,
	0x4b, 0x48, 0x4f, 0x46, 0x4d, 0x50, 0x48, 0x45, 0x46, 0x4e, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x45,
	0x42, 0x4e, 0x4e, 0x42, 0x45, 0x45, 0x47, 0x4a, 0x46, 0x4e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x45, 0x42, 0x4e, 0x4e, 0x42, 0x45, 0x45, 0x47, 0x4a, 0x46, 0x4e, 0x12, 0x20, 0x0a,
	0x0b, 0x4b, 0x46, 0x48, 0x44, 0x4f, 0x46, 0x4e, 0x44, 0x4e, 0x41, 0x48, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x46, 0x48, 0x44, 0x4f, 0x46, 0x4e, 0x44, 0x4e, 0x41, 0x48, 0x12,
	0x2e, 0x0a, 0x0b, 0x49, 0x47, 0x4f, 0x47, 0x50, 0x4c, 0x48, 0x4b, 0x48, 0x44, 0x41, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x4b, 0x42, 0x41, 0x42, 0x50, 0x4b, 0x49, 0x4b,
	0x41, 0x46, 0x52, 0x0b, 0x49, 0x47, 0x4f, 0x47, 0x50, 0x4c, 0x48, 0x4b, 0x48, 0x44, 0x41, 0x12,
	0x45, 0x0a, 0x0b, 0x50, 0x43, 0x46, 0x42, 0x4e, 0x46, 0x42, 0x44, 0x4c, 0x41, 0x48, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x50, 0x43, 0x46, 0x42, 0x4e, 0x46, 0x42,
	0x44, 0x4c, 0x41, 0x48, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x50, 0x43, 0x46, 0x42, 0x4e,
	0x46, 0x42, 0x44, 0x4c, 0x41, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x49, 0x46, 0x4a, 0x4f, 0x4f,
	0x45, 0x45, 0x45, 0x4c, 0x4b, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x49, 0x46,
	0x4a, 0x4f, 0x4f, 0x45, 0x45, 0x45, 0x4c, 0x4b, 0x12, 0x2e, 0x0a, 0x0b, 0x43, 0x44, 0x4d, 0x48,
	0x42, 0x4a, 0x50, 0x4b, 0x4c, 0x4a, 0x4f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4f, 0x48, 0x41, 0x4a, 0x4e, 0x41, 0x45, 0x42, 0x46, 0x43, 0x4f, 0x52, 0x0b, 0x43, 0x44, 0x4d,
	0x48, 0x42, 0x4a, 0x50, 0x4b, 0x4c, 0x4a, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x4f, 0x4c, 0x49, 0x49, 0x4b, 0x4f, 0x50, 0x41, 0x45, 0x42, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x4f, 0x4c, 0x49, 0x49, 0x4b, 0x4f, 0x50, 0x41, 0x45, 0x42,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x43,
	0x4a, 0x4e, 0x49, 0x49, 0x44, 0x42, 0x4b, 0x43, 0x45, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x0b, 0x43, 0x43, 0x4a, 0x4e, 0x49, 0x49, 0x44, 0x42, 0x4b, 0x43, 0x45, 0x1a, 0x3e, 0x0a, 0x10,
	0x50, 0x43, 0x46, 0x42, 0x4e, 0x46, 0x42, 0x44, 0x4c, 0x41, 0x48, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetAlleyInfoScRsp_proto_rawDescOnce sync.Once
	file_GetAlleyInfoScRsp_proto_rawDescData = file_GetAlleyInfoScRsp_proto_rawDesc
)

func file_GetAlleyInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetAlleyInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetAlleyInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetAlleyInfoScRsp_proto_rawDescData)
	})
	return file_GetAlleyInfoScRsp_proto_rawDescData
}

var file_GetAlleyInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GetAlleyInfoScRsp_proto_goTypes = []interface{}{
	(*GetAlleyInfoScRsp)(nil), // 0: GetAlleyInfoScRsp
	nil,                       // 1: GetAlleyInfoScRsp.PCFBNFBDLAHEntry
	(*DMCFKOHNIHB)(nil),       // 2: DMCFKOHNIHB
	(*NABCIIFFDCE)(nil),       // 3: NABCIIFFDCE
	(*IKBABPKIKAF)(nil),       // 4: IKBABPKIKAF
	(*OHAJNAEBFCO)(nil),       // 5: OHAJNAEBFCO
}
var file_GetAlleyInfoScRsp_proto_depIdxs = []int32{
	2, // 0: GetAlleyInfoScRsp.FIFHMONIDBL:type_name -> DMCFKOHNIHB
	3, // 1: GetAlleyInfoScRsp.KHOFMPHEFNC:type_name -> NABCIIFFDCE
	4, // 2: GetAlleyInfoScRsp.IGOGPLHKHDA:type_name -> IKBABPKIKAF
	1, // 3: GetAlleyInfoScRsp.PCFBNFBDLAH:type_name -> GetAlleyInfoScRsp.PCFBNFBDLAHEntry
	5, // 4: GetAlleyInfoScRsp.CDMHBJPKLJO:type_name -> OHAJNAEBFCO
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_GetAlleyInfoScRsp_proto_init() }
func file_GetAlleyInfoScRsp_proto_init() {
	if File_GetAlleyInfoScRsp_proto != nil {
		return
	}
	file_DMCFKOHNIHB_proto_init()
	file_NABCIIFFDCE_proto_init()
	file_IKBABPKIKAF_proto_init()
	file_OHAJNAEBFCO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetAlleyInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlleyInfoScRsp); i {
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
			RawDescriptor: file_GetAlleyInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetAlleyInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetAlleyInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetAlleyInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetAlleyInfoScRsp_proto = out.File
	file_GetAlleyInfoScRsp_proto_rawDesc = nil
	file_GetAlleyInfoScRsp_proto_goTypes = nil
	file_GetAlleyInfoScRsp_proto_depIdxs = nil
}
