// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: OHBCINICBHP.proto

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

type OHBCINICBHP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FGPDKJPAKNC MEMCBKNIBAJ    `protobuf:"varint,1,opt,name=FGPDKJPAKNC,proto3,enum=MEMCBKNIBAJ" json:"FGPDKJPAKNC,omitempty"`
	HKOOAJAOJFJ DAADHEDHJNH    `protobuf:"varint,9,opt,name=HKOOAJAOJFJ,proto3,enum=DAADHEDHJNH" json:"HKOOAJAOJFJ,omitempty"`
	HHMPIBMIHJA uint32         `protobuf:"varint,2,opt,name=HHMPIBMIHJA,proto3" json:"HHMPIBMIHJA,omitempty"`
	OCKLJIFBFIN uint32         `protobuf:"varint,6,opt,name=OCKLJIFBFIN,proto3" json:"OCKLJIFBFIN,omitempty"`
	DBKFLOBAMAE []*OOOJKGDMFOK `protobuf:"bytes,4,rep,name=DBKFLOBAMAE,proto3" json:"DBKFLOBAMAE,omitempty"`
}

func (x *OHBCINICBHP) Reset() {
	*x = OHBCINICBHP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OHBCINICBHP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OHBCINICBHP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OHBCINICBHP) ProtoMessage() {}

func (x *OHBCINICBHP) ProtoReflect() protoreflect.Message {
	mi := &file_OHBCINICBHP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OHBCINICBHP.ProtoReflect.Descriptor instead.
func (*OHBCINICBHP) Descriptor() ([]byte, []int) {
	return file_OHBCINICBHP_proto_rawDescGZIP(), []int{0}
}

func (x *OHBCINICBHP) GetFGPDKJPAKNC() MEMCBKNIBAJ {
	if x != nil {
		return x.FGPDKJPAKNC
	}
	return MEMCBKNIBAJ_PAGE_NONE
}

func (x *OHBCINICBHP) GetHKOOAJAOJFJ() DAADHEDHJNH {
	if x != nil {
		return x.HKOOAJAOJFJ
	}
	return DAADHEDHJNH_PAGE_DESC_NONE
}

func (x *OHBCINICBHP) GetHHMPIBMIHJA() uint32 {
	if x != nil {
		return x.HHMPIBMIHJA
	}
	return 0
}

func (x *OHBCINICBHP) GetOCKLJIFBFIN() uint32 {
	if x != nil {
		return x.OCKLJIFBFIN
	}
	return 0
}

func (x *OHBCINICBHP) GetDBKFLOBAMAE() []*OOOJKGDMFOK {
	if x != nil {
		return x.DBKFLOBAMAE
	}
	return nil
}

var File_OHBCINICBHP_proto protoreflect.FileDescriptor

var file_OHBCINICBHP_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4f, 0x48, 0x42, 0x43, 0x49, 0x4e, 0x49, 0x43, 0x42, 0x48, 0x50, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x45, 0x4d, 0x43, 0x42, 0x4b, 0x4e, 0x49, 0x42, 0x41, 0x4a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x41, 0x41, 0x44, 0x48, 0x45, 0x44, 0x48,
	0x4a, 0x4e, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4f, 0x4f, 0x4a, 0x4b,
	0x47, 0x44, 0x4d, 0x46, 0x4f, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a,
	0x0b, 0x4f, 0x48, 0x42, 0x43, 0x49, 0x4e, 0x49, 0x43, 0x42, 0x48, 0x50, 0x12, 0x2e, 0x0a, 0x0b,
	0x46, 0x47, 0x50, 0x44, 0x4b, 0x4a, 0x50, 0x41, 0x4b, 0x4e, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x4d, 0x45, 0x4d, 0x43, 0x42, 0x4b, 0x4e, 0x49, 0x42, 0x41, 0x4a, 0x52,
	0x0b, 0x46, 0x47, 0x50, 0x44, 0x4b, 0x4a, 0x50, 0x41, 0x4b, 0x4e, 0x43, 0x12, 0x2e, 0x0a, 0x0b,
	0x48, 0x4b, 0x4f, 0x4f, 0x41, 0x4a, 0x41, 0x4f, 0x4a, 0x46, 0x4a, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x44, 0x41, 0x41, 0x44, 0x48, 0x45, 0x44, 0x48, 0x4a, 0x4e, 0x48, 0x52,
	0x0b, 0x48, 0x4b, 0x4f, 0x4f, 0x41, 0x4a, 0x41, 0x4f, 0x4a, 0x46, 0x4a, 0x12, 0x20, 0x0a, 0x0b,
	0x48, 0x48, 0x4d, 0x50, 0x49, 0x42, 0x4d, 0x49, 0x48, 0x4a, 0x41, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x48, 0x48, 0x4d, 0x50, 0x49, 0x42, 0x4d, 0x49, 0x48, 0x4a, 0x41, 0x12, 0x20,
	0x0a, 0x0b, 0x4f, 0x43, 0x4b, 0x4c, 0x4a, 0x49, 0x46, 0x42, 0x46, 0x49, 0x4e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x43, 0x4b, 0x4c, 0x4a, 0x49, 0x46, 0x42, 0x46, 0x49, 0x4e,
	0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x42, 0x4b, 0x46, 0x4c, 0x4f, 0x42, 0x41, 0x4d, 0x41, 0x45, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4f, 0x4f, 0x4a, 0x4b, 0x47, 0x44, 0x4d,
	0x46, 0x4f, 0x4b, 0x52, 0x0b, 0x44, 0x42, 0x4b, 0x46, 0x4c, 0x4f, 0x42, 0x41, 0x4d, 0x41, 0x45,
	0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OHBCINICBHP_proto_rawDescOnce sync.Once
	file_OHBCINICBHP_proto_rawDescData = file_OHBCINICBHP_proto_rawDesc
)

func file_OHBCINICBHP_proto_rawDescGZIP() []byte {
	file_OHBCINICBHP_proto_rawDescOnce.Do(func() {
		file_OHBCINICBHP_proto_rawDescData = protoimpl.X.CompressGZIP(file_OHBCINICBHP_proto_rawDescData)
	})
	return file_OHBCINICBHP_proto_rawDescData
}

var file_OHBCINICBHP_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OHBCINICBHP_proto_goTypes = []interface{}{
	(*OHBCINICBHP)(nil), // 0: OHBCINICBHP
	(MEMCBKNIBAJ)(0),    // 1: MEMCBKNIBAJ
	(DAADHEDHJNH)(0),    // 2: DAADHEDHJNH
	(*OOOJKGDMFOK)(nil), // 3: OOOJKGDMFOK
}
var file_OHBCINICBHP_proto_depIdxs = []int32{
	1, // 0: OHBCINICBHP.FGPDKJPAKNC:type_name -> MEMCBKNIBAJ
	2, // 1: OHBCINICBHP.HKOOAJAOJFJ:type_name -> DAADHEDHJNH
	3, // 2: OHBCINICBHP.DBKFLOBAMAE:type_name -> OOOJKGDMFOK
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_OHBCINICBHP_proto_init() }
func file_OHBCINICBHP_proto_init() {
	if File_OHBCINICBHP_proto != nil {
		return
	}
	file_MEMCBKNIBAJ_proto_init()
	file_DAADHEDHJNH_proto_init()
	file_OOOJKGDMFOK_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_OHBCINICBHP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OHBCINICBHP); i {
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
			RawDescriptor: file_OHBCINICBHP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OHBCINICBHP_proto_goTypes,
		DependencyIndexes: file_OHBCINICBHP_proto_depIdxs,
		MessageInfos:      file_OHBCINICBHP_proto_msgTypes,
	}.Build()
	File_OHBCINICBHP_proto = out.File
	file_OHBCINICBHP_proto_rawDesc = nil
	file_OHBCINICBHP_proto_goTypes = nil
	file_OHBCINICBHP_proto_depIdxs = nil
}
