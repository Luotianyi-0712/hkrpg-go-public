// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: OKANJDMIODN.proto

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

type OKANJDMIODN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NBHPDMDNDKL                 uint32       `protobuf:"varint,9,opt,name=NBHPDMDNDKL,proto3" json:"NBHPDMDNDKL,omitempty"`
	PUNK_LORD_SHARE_TYPE_FRIEND *CHECMAAPOJC `protobuf:"bytes,2,opt,name=PUNK_LORD_SHARE_TYPE_FRIEND,json=PUNKLORDSHARETYPEFRIEND,proto3" json:"PUNK_LORD_SHARE_TYPE_FRIEND,omitempty"`
	// Types that are assignable to KJHHJKKOAJA:
	//
	//	*OKANJDMIODN_OGEPOCCCEBK
	//	*OKANJDMIODN_MJFDOAHFAAF
	//	*OKANJDMIODN_JOANGDLCHKM
	//	*OKANJDMIODN_NOIFOCJDLGC
	//	*OKANJDMIODN_FOHJOGGFIII
	//	*OKANJDMIODN_SwordTrainingGameInfo
	//	*OKANJDMIODN_DLDADINBFNF
	//	*OKANJDMIODN_ADDNICELMIC
	KJHHJKKOAJA isOKANJDMIODN_KJHHJKKOAJA `protobuf_oneof:"KJHHJKKOAJA"`
}

func (x *OKANJDMIODN) Reset() {
	*x = OKANJDMIODN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OKANJDMIODN_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OKANJDMIODN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OKANJDMIODN) ProtoMessage() {}

func (x *OKANJDMIODN) ProtoReflect() protoreflect.Message {
	mi := &file_OKANJDMIODN_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OKANJDMIODN.ProtoReflect.Descriptor instead.
func (*OKANJDMIODN) Descriptor() ([]byte, []int) {
	return file_OKANJDMIODN_proto_rawDescGZIP(), []int{0}
}

func (x *OKANJDMIODN) GetNBHPDMDNDKL() uint32 {
	if x != nil {
		return x.NBHPDMDNDKL
	}
	return 0
}

func (x *OKANJDMIODN) GetPUNK_LORD_SHARE_TYPE_FRIEND() *CHECMAAPOJC {
	if x != nil {
		return x.PUNK_LORD_SHARE_TYPE_FRIEND
	}
	return nil
}

func (m *OKANJDMIODN) GetKJHHJKKOAJA() isOKANJDMIODN_KJHHJKKOAJA {
	if m != nil {
		return m.KJHHJKKOAJA
	}
	return nil
}

func (x *OKANJDMIODN) GetOGEPOCCCEBK() *JNOEPBPBADE {
	if x, ok := x.GetKJHHJKKOAJA().(*OKANJDMIODN_OGEPOCCCEBK); ok {
		return x.OGEPOCCCEBK
	}
	return nil
}

func (x *OKANJDMIODN) GetMJFDOAHFAAF() *CEHEDCEICDE {
	if x, ok := x.GetKJHHJKKOAJA().(*OKANJDMIODN_MJFDOAHFAAF); ok {
		return x.MJFDOAHFAAF
	}
	return nil
}

func (x *OKANJDMIODN) GetJOANGDLCHKM() *HHGCBNHLGGN {
	if x, ok := x.GetKJHHJKKOAJA().(*OKANJDMIODN_JOANGDLCHKM); ok {
		return x.JOANGDLCHKM
	}
	return nil
}

func (x *OKANJDMIODN) GetNOIFOCJDLGC() *LCIHCFHIJEN {
	if x, ok := x.GetKJHHJKKOAJA().(*OKANJDMIODN_NOIFOCJDLGC); ok {
		return x.NOIFOCJDLGC
	}
	return nil
}

func (x *OKANJDMIODN) GetFOHJOGGFIII() *KMBOMDFAMEM {
	if x, ok := x.GetKJHHJKKOAJA().(*OKANJDMIODN_FOHJOGGFIII); ok {
		return x.FOHJOGGFIII
	}
	return nil
}

func (x *OKANJDMIODN) GetSwordTrainingGameInfo() *LCIPLPHHHJC {
	if x, ok := x.GetKJHHJKKOAJA().(*OKANJDMIODN_SwordTrainingGameInfo); ok {
		return x.SwordTrainingGameInfo
	}
	return nil
}

func (x *OKANJDMIODN) GetDLDADINBFNF() *PMDKLNBLFMI {
	if x, ok := x.GetKJHHJKKOAJA().(*OKANJDMIODN_DLDADINBFNF); ok {
		return x.DLDADINBFNF
	}
	return nil
}

func (x *OKANJDMIODN) GetADDNICELMIC() *POHKGHPBLFB {
	if x, ok := x.GetKJHHJKKOAJA().(*OKANJDMIODN_ADDNICELMIC); ok {
		return x.ADDNICELMIC
	}
	return nil
}

type isOKANJDMIODN_KJHHJKKOAJA interface {
	isOKANJDMIODN_KJHHJKKOAJA()
}

type OKANJDMIODN_OGEPOCCCEBK struct {
	OGEPOCCCEBK *JNOEPBPBADE `protobuf:"bytes,7,opt,name=OGEPOCCCEBK,proto3,oneof"`
}

type OKANJDMIODN_MJFDOAHFAAF struct {
	MJFDOAHFAAF *CEHEDCEICDE `protobuf:"bytes,1,opt,name=MJFDOAHFAAF,proto3,oneof"`
}

type OKANJDMIODN_JOANGDLCHKM struct {
	JOANGDLCHKM *HHGCBNHLGGN `protobuf:"bytes,4,opt,name=JOANGDLCHKM,proto3,oneof"`
}

type OKANJDMIODN_NOIFOCJDLGC struct {
	NOIFOCJDLGC *LCIHCFHIJEN `protobuf:"bytes,10,opt,name=NOIFOCJDLGC,proto3,oneof"`
}

type OKANJDMIODN_FOHJOGGFIII struct {
	FOHJOGGFIII *KMBOMDFAMEM `protobuf:"bytes,6,opt,name=FOHJOGGFIII,proto3,oneof"`
}

type OKANJDMIODN_SwordTrainingGameInfo struct {
	SwordTrainingGameInfo *LCIPLPHHHJC `protobuf:"bytes,12,opt,name=sword_training_game_info,json=swordTrainingGameInfo,proto3,oneof"`
}

type OKANJDMIODN_DLDADINBFNF struct {
	DLDADINBFNF *PMDKLNBLFMI `protobuf:"bytes,8,opt,name=DLDADINBFNF,proto3,oneof"`
}

type OKANJDMIODN_ADDNICELMIC struct {
	ADDNICELMIC *POHKGHPBLFB `protobuf:"bytes,11,opt,name=ADDNICELMIC,proto3,oneof"`
}

func (*OKANJDMIODN_OGEPOCCCEBK) isOKANJDMIODN_KJHHJKKOAJA() {}

func (*OKANJDMIODN_MJFDOAHFAAF) isOKANJDMIODN_KJHHJKKOAJA() {}

func (*OKANJDMIODN_JOANGDLCHKM) isOKANJDMIODN_KJHHJKKOAJA() {}

func (*OKANJDMIODN_NOIFOCJDLGC) isOKANJDMIODN_KJHHJKKOAJA() {}

func (*OKANJDMIODN_FOHJOGGFIII) isOKANJDMIODN_KJHHJKKOAJA() {}

func (*OKANJDMIODN_SwordTrainingGameInfo) isOKANJDMIODN_KJHHJKKOAJA() {}

func (*OKANJDMIODN_DLDADINBFNF) isOKANJDMIODN_KJHHJKKOAJA() {}

func (*OKANJDMIODN_ADDNICELMIC) isOKANJDMIODN_KJHHJKKOAJA() {}

var File_OKANJDMIODN_proto protoreflect.FileDescriptor

var file_OKANJDMIODN_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4f, 0x4b, 0x41, 0x4e, 0x4a, 0x44, 0x4d, 0x49, 0x4f, 0x44, 0x4e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x43, 0x49, 0x48, 0x43, 0x46, 0x48, 0x49, 0x4a, 0x45, 0x4e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x4e, 0x4f, 0x45, 0x50, 0x42, 0x50, 0x42,
	0x41, 0x44, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x45, 0x48, 0x45, 0x44,
	0x43, 0x45, 0x49, 0x43, 0x44, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x48,
	0x47, 0x43, 0x42, 0x4e, 0x48, 0x4c, 0x47, 0x47, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x43, 0x48, 0x45, 0x43, 0x4d, 0x41, 0x41, 0x50, 0x4f, 0x4a, 0x43, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4d, 0x42, 0x4f, 0x4d, 0x44, 0x46, 0x41, 0x4d, 0x45, 0x4d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x4f, 0x48, 0x4b, 0x47, 0x48, 0x50, 0x42, 0x4c,
	0x46, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x43, 0x49, 0x50, 0x4c, 0x50,
	0x48, 0x48, 0x48, 0x4a, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x4d, 0x44,
	0x4b, 0x4c, 0x4e, 0x42, 0x4c, 0x46, 0x4d, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1,
	0x04, 0x0a, 0x0b, 0x4f, 0x4b, 0x41, 0x4e, 0x4a, 0x44, 0x4d, 0x49, 0x4f, 0x44, 0x4e, 0x12, 0x20,
	0x0a, 0x0b, 0x4e, 0x42, 0x48, 0x50, 0x44, 0x4d, 0x44, 0x4e, 0x44, 0x4b, 0x4c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x42, 0x48, 0x50, 0x44, 0x4d, 0x44, 0x4e, 0x44, 0x4b, 0x4c,
	0x12, 0x4a, 0x0a, 0x1b, 0x50, 0x55, 0x4e, 0x4b, 0x5f, 0x4c, 0x4f, 0x52, 0x44, 0x5f, 0x53, 0x48,
	0x41, 0x52, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x48, 0x45, 0x43, 0x4d, 0x41, 0x41, 0x50,
	0x4f, 0x4a, 0x43, 0x52, 0x17, 0x50, 0x55, 0x4e, 0x4b, 0x4c, 0x4f, 0x52, 0x44, 0x53, 0x48, 0x41,
	0x52, 0x45, 0x54, 0x59, 0x50, 0x45, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x12, 0x30, 0x0a, 0x0b,
	0x4f, 0x47, 0x45, 0x50, 0x4f, 0x43, 0x43, 0x43, 0x45, 0x42, 0x4b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x4e, 0x4f, 0x45, 0x50, 0x42, 0x50, 0x42, 0x41, 0x44, 0x45, 0x48,
	0x00, 0x52, 0x0b, 0x4f, 0x47, 0x45, 0x50, 0x4f, 0x43, 0x43, 0x43, 0x45, 0x42, 0x4b, 0x12, 0x30,
	0x0a, 0x0b, 0x4d, 0x4a, 0x46, 0x44, 0x4f, 0x41, 0x48, 0x46, 0x41, 0x41, 0x46, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x45, 0x48, 0x45, 0x44, 0x43, 0x45, 0x49, 0x43, 0x44,
	0x45, 0x48, 0x00, 0x52, 0x0b, 0x4d, 0x4a, 0x46, 0x44, 0x4f, 0x41, 0x48, 0x46, 0x41, 0x41, 0x46,
	0x12, 0x30, 0x0a, 0x0b, 0x4a, 0x4f, 0x41, 0x4e, 0x47, 0x44, 0x4c, 0x43, 0x48, 0x4b, 0x4d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x48, 0x47, 0x43, 0x42, 0x4e, 0x48, 0x4c,
	0x47, 0x47, 0x4e, 0x48, 0x00, 0x52, 0x0b, 0x4a, 0x4f, 0x41, 0x4e, 0x47, 0x44, 0x4c, 0x43, 0x48,
	0x4b, 0x4d, 0x12, 0x30, 0x0a, 0x0b, 0x4e, 0x4f, 0x49, 0x46, 0x4f, 0x43, 0x4a, 0x44, 0x4c, 0x47,
	0x43, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x43, 0x49, 0x48, 0x43, 0x46,
	0x48, 0x49, 0x4a, 0x45, 0x4e, 0x48, 0x00, 0x52, 0x0b, 0x4e, 0x4f, 0x49, 0x46, 0x4f, 0x43, 0x4a,
	0x44, 0x4c, 0x47, 0x43, 0x12, 0x30, 0x0a, 0x0b, 0x46, 0x4f, 0x48, 0x4a, 0x4f, 0x47, 0x47, 0x46,
	0x49, 0x49, 0x49, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4d, 0x42, 0x4f,
	0x4d, 0x44, 0x46, 0x41, 0x4d, 0x45, 0x4d, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x4f, 0x48, 0x4a, 0x4f,
	0x47, 0x47, 0x46, 0x49, 0x49, 0x49, 0x12, 0x47, 0x0a, 0x18, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f,
	0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x43, 0x49, 0x50, 0x4c,
	0x50, 0x48, 0x48, 0x48, 0x4a, 0x43, 0x48, 0x00, 0x52, 0x15, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x54,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x30, 0x0a, 0x0b, 0x44, 0x4c, 0x44, 0x41, 0x44, 0x49, 0x4e, 0x42, 0x46, 0x4e, 0x46, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x4d, 0x44, 0x4b, 0x4c, 0x4e, 0x42, 0x4c, 0x46,
	0x4d, 0x49, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x4c, 0x44, 0x41, 0x44, 0x49, 0x4e, 0x42, 0x46, 0x4e,
	0x46, 0x12, 0x30, 0x0a, 0x0b, 0x41, 0x44, 0x44, 0x4e, 0x49, 0x43, 0x45, 0x4c, 0x4d, 0x49, 0x43,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x4f, 0x48, 0x4b, 0x47, 0x48, 0x50,
	0x42, 0x4c, 0x46, 0x42, 0x48, 0x00, 0x52, 0x0b, 0x41, 0x44, 0x44, 0x4e, 0x49, 0x43, 0x45, 0x4c,
	0x4d, 0x49, 0x43, 0x42, 0x0d, 0x0a, 0x0b, 0x4b, 0x4a, 0x48, 0x48, 0x4a, 0x4b, 0x4b, 0x4f, 0x41,
	0x4a, 0x41, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OKANJDMIODN_proto_rawDescOnce sync.Once
	file_OKANJDMIODN_proto_rawDescData = file_OKANJDMIODN_proto_rawDesc
)

func file_OKANJDMIODN_proto_rawDescGZIP() []byte {
	file_OKANJDMIODN_proto_rawDescOnce.Do(func() {
		file_OKANJDMIODN_proto_rawDescData = protoimpl.X.CompressGZIP(file_OKANJDMIODN_proto_rawDescData)
	})
	return file_OKANJDMIODN_proto_rawDescData
}

var file_OKANJDMIODN_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OKANJDMIODN_proto_goTypes = []interface{}{
	(*OKANJDMIODN)(nil), // 0: OKANJDMIODN
	(*CHECMAAPOJC)(nil), // 1: CHECMAAPOJC
	(*JNOEPBPBADE)(nil), // 2: JNOEPBPBADE
	(*CEHEDCEICDE)(nil), // 3: CEHEDCEICDE
	(*HHGCBNHLGGN)(nil), // 4: HHGCBNHLGGN
	(*LCIHCFHIJEN)(nil), // 5: LCIHCFHIJEN
	(*KMBOMDFAMEM)(nil), // 6: KMBOMDFAMEM
	(*LCIPLPHHHJC)(nil), // 7: LCIPLPHHHJC
	(*PMDKLNBLFMI)(nil), // 8: PMDKLNBLFMI
	(*POHKGHPBLFB)(nil), // 9: POHKGHPBLFB
}
var file_OKANJDMIODN_proto_depIdxs = []int32{
	1, // 0: OKANJDMIODN.PUNK_LORD_SHARE_TYPE_FRIEND:type_name -> CHECMAAPOJC
	2, // 1: OKANJDMIODN.OGEPOCCCEBK:type_name -> JNOEPBPBADE
	3, // 2: OKANJDMIODN.MJFDOAHFAAF:type_name -> CEHEDCEICDE
	4, // 3: OKANJDMIODN.JOANGDLCHKM:type_name -> HHGCBNHLGGN
	5, // 4: OKANJDMIODN.NOIFOCJDLGC:type_name -> LCIHCFHIJEN
	6, // 5: OKANJDMIODN.FOHJOGGFIII:type_name -> KMBOMDFAMEM
	7, // 6: OKANJDMIODN.sword_training_game_info:type_name -> LCIPLPHHHJC
	8, // 7: OKANJDMIODN.DLDADINBFNF:type_name -> PMDKLNBLFMI
	9, // 8: OKANJDMIODN.ADDNICELMIC:type_name -> POHKGHPBLFB
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_OKANJDMIODN_proto_init() }
func file_OKANJDMIODN_proto_init() {
	if File_OKANJDMIODN_proto != nil {
		return
	}
	file_LCIHCFHIJEN_proto_init()
	file_JNOEPBPBADE_proto_init()
	file_CEHEDCEICDE_proto_init()
	file_HHGCBNHLGGN_proto_init()
	file_CHECMAAPOJC_proto_init()
	file_KMBOMDFAMEM_proto_init()
	file_POHKGHPBLFB_proto_init()
	file_LCIPLPHHHJC_proto_init()
	file_PMDKLNBLFMI_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_OKANJDMIODN_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OKANJDMIODN); i {
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
	file_OKANJDMIODN_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OKANJDMIODN_OGEPOCCCEBK)(nil),
		(*OKANJDMIODN_MJFDOAHFAAF)(nil),
		(*OKANJDMIODN_JOANGDLCHKM)(nil),
		(*OKANJDMIODN_NOIFOCJDLGC)(nil),
		(*OKANJDMIODN_FOHJOGGFIII)(nil),
		(*OKANJDMIODN_SwordTrainingGameInfo)(nil),
		(*OKANJDMIODN_DLDADINBFNF)(nil),
		(*OKANJDMIODN_ADDNICELMIC)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_OKANJDMIODN_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OKANJDMIODN_proto_goTypes,
		DependencyIndexes: file_OKANJDMIODN_proto_depIdxs,
		MessageInfos:      file_OKANJDMIODN_proto_msgTypes,
	}.Build()
	File_OKANJDMIODN_proto = out.File
	file_OKANJDMIODN_proto_rawDesc = nil
	file_OKANJDMIODN_proto_goTypes = nil
	file_OKANJDMIODN_proto_depIdxs = nil
}
