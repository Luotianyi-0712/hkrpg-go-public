// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueFinishInfo.proto

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

type RogueFinishInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OJBOCHEBJFO uint32                 `protobuf:"varint,1,opt,name=OJBOCHEBJFO,proto3" json:"OJBOCHEBJFO,omitempty"`
	CDGDNPLKHDN uint32                 `protobuf:"varint,1818,opt,name=CDGDNPLKHDN,proto3" json:"CDGDNPLKHDN,omitempty"`
	KAGHBIBAIPP *MLPPNJOPPML           `protobuf:"bytes,13,opt,name=KAGHBIBAIPP,proto3" json:"KAGHBIBAIPP,omitempty"`
	FJHJDMNBKDO *ItemList              `protobuf:"bytes,12,opt,name=FJHJDMNBKDO,proto3" json:"FJHJDMNBKDO,omitempty"`
	OCMOCECJGHJ *ItemList              `protobuf:"bytes,9,opt,name=OCMOCECJGHJ,proto3" json:"OCMOCECJGHJ,omitempty"`
	DNNHHDMJPHO *RogueScoreRewardInfo  `protobuf:"bytes,2,opt,name=DNNHHDMJPHO,proto3" json:"DNNHHDMJPHO,omitempty"`
	ABOGJOHIGJL uint32                 `protobuf:"varint,7,opt,name=ABOGJOHIGJL,proto3" json:"ABOGJOHIGJL,omitempty"`
	RecordInfo  *RogueRecordInfo       `protobuf:"bytes,14,opt,name=record_info,json=recordInfo,proto3" json:"record_info,omitempty"`
	ScoreInfo   *RogueExploreScoreInfo `protobuf:"bytes,11,opt,name=score_info,json=scoreInfo,proto3" json:"score_info,omitempty"`
	MAGAIHIELEN uint32                 `protobuf:"varint,15,opt,name=MAGAIHIELEN,proto3" json:"MAGAIHIELEN,omitempty"`
	LEJJFNPJAOA uint32                 `protobuf:"varint,8,opt,name=LEJJFNPJAOA,proto3" json:"LEJJFNPJAOA,omitempty"`
	KFJLHEBDPNG *ItemList              `protobuf:"bytes,3,opt,name=KFJLHEBDPNG,proto3" json:"KFJLHEBDPNG,omitempty"`
	IsWin       bool                   `protobuf:"varint,4,opt,name=is_win,json=isWin,proto3" json:"is_win,omitempty"`
	AreaId      uint32                 `protobuf:"varint,1131,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	KKJNIACJBFK bool                   `protobuf:"varint,10,opt,name=KKJNIACJBFK,proto3" json:"KKJNIACJBFK,omitempty"`
	NCBHLBJAAJK *RogueScoreRewardInfo  `protobuf:"bytes,5,opt,name=NCBHLBJAAJK,proto3" json:"NCBHLBJAAJK,omitempty"`
	ScoreId     uint32                 `protobuf:"varint,6,opt,name=score_id,json=scoreId,proto3" json:"score_id,omitempty"`
}

func (x *RogueFinishInfo) Reset() {
	*x = RogueFinishInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueFinishInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueFinishInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueFinishInfo) ProtoMessage() {}

func (x *RogueFinishInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueFinishInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueFinishInfo.ProtoReflect.Descriptor instead.
func (*RogueFinishInfo) Descriptor() ([]byte, []int) {
	return file_RogueFinishInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueFinishInfo) GetOJBOCHEBJFO() uint32 {
	if x != nil {
		return x.OJBOCHEBJFO
	}
	return 0
}

func (x *RogueFinishInfo) GetCDGDNPLKHDN() uint32 {
	if x != nil {
		return x.CDGDNPLKHDN
	}
	return 0
}

func (x *RogueFinishInfo) GetKAGHBIBAIPP() *MLPPNJOPPML {
	if x != nil {
		return x.KAGHBIBAIPP
	}
	return nil
}

func (x *RogueFinishInfo) GetFJHJDMNBKDO() *ItemList {
	if x != nil {
		return x.FJHJDMNBKDO
	}
	return nil
}

func (x *RogueFinishInfo) GetOCMOCECJGHJ() *ItemList {
	if x != nil {
		return x.OCMOCECJGHJ
	}
	return nil
}

func (x *RogueFinishInfo) GetDNNHHDMJPHO() *RogueScoreRewardInfo {
	if x != nil {
		return x.DNNHHDMJPHO
	}
	return nil
}

func (x *RogueFinishInfo) GetABOGJOHIGJL() uint32 {
	if x != nil {
		return x.ABOGJOHIGJL
	}
	return 0
}

func (x *RogueFinishInfo) GetRecordInfo() *RogueRecordInfo {
	if x != nil {
		return x.RecordInfo
	}
	return nil
}

func (x *RogueFinishInfo) GetScoreInfo() *RogueExploreScoreInfo {
	if x != nil {
		return x.ScoreInfo
	}
	return nil
}

func (x *RogueFinishInfo) GetMAGAIHIELEN() uint32 {
	if x != nil {
		return x.MAGAIHIELEN
	}
	return 0
}

func (x *RogueFinishInfo) GetLEJJFNPJAOA() uint32 {
	if x != nil {
		return x.LEJJFNPJAOA
	}
	return 0
}

func (x *RogueFinishInfo) GetKFJLHEBDPNG() *ItemList {
	if x != nil {
		return x.KFJLHEBDPNG
	}
	return nil
}

func (x *RogueFinishInfo) GetIsWin() bool {
	if x != nil {
		return x.IsWin
	}
	return false
}

func (x *RogueFinishInfo) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *RogueFinishInfo) GetKKJNIACJBFK() bool {
	if x != nil {
		return x.KKJNIACJBFK
	}
	return false
}

func (x *RogueFinishInfo) GetNCBHLBJAAJK() *RogueScoreRewardInfo {
	if x != nil {
		return x.NCBHLBJAAJK
	}
	return nil
}

func (x *RogueFinishInfo) GetScoreId() uint32 {
	if x != nil {
		return x.ScoreId
	}
	return 0
}

var File_RogueFinishInfo_proto protoreflect.FileDescriptor

var file_RogueFinishInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x4c, 0x50, 0x50,
	0x4e, 0x4a, 0x4f, 0x50, 0x50, 0x4d, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x05, 0x0a, 0x0f, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x4f, 0x4a, 0x42, 0x4f, 0x43, 0x48, 0x45, 0x42, 0x4a, 0x46, 0x4f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4a, 0x42, 0x4f, 0x43, 0x48, 0x45, 0x42, 0x4a, 0x46, 0x4f,
	0x12, 0x21, 0x0a, 0x0b, 0x43, 0x44, 0x47, 0x44, 0x4e, 0x50, 0x4c, 0x4b, 0x48, 0x44, 0x4e, 0x18,
	0x9a, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x44, 0x47, 0x44, 0x4e, 0x50, 0x4c, 0x4b,
	0x48, 0x44, 0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x4b, 0x41, 0x47, 0x48, 0x42, 0x49, 0x42, 0x41, 0x49,
	0x50, 0x50, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x4c, 0x50, 0x50, 0x4e,
	0x4a, 0x4f, 0x50, 0x50, 0x4d, 0x4c, 0x52, 0x0b, 0x4b, 0x41, 0x47, 0x48, 0x42, 0x49, 0x42, 0x41,
	0x49, 0x50, 0x50, 0x12, 0x2b, 0x0a, 0x0b, 0x46, 0x4a, 0x48, 0x4a, 0x44, 0x4d, 0x4e, 0x42, 0x4b,
	0x44, 0x4f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x0b, 0x46, 0x4a, 0x48, 0x4a, 0x44, 0x4d, 0x4e, 0x42, 0x4b, 0x44, 0x4f,
	0x12, 0x2b, 0x0a, 0x0b, 0x4f, 0x43, 0x4d, 0x4f, 0x43, 0x45, 0x43, 0x4a, 0x47, 0x48, 0x4a, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x0b, 0x4f, 0x43, 0x4d, 0x4f, 0x43, 0x45, 0x43, 0x4a, 0x47, 0x48, 0x4a, 0x12, 0x37, 0x0a,
	0x0b, 0x44, 0x4e, 0x4e, 0x48, 0x48, 0x44, 0x4d, 0x4a, 0x50, 0x48, 0x4f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x44, 0x4e, 0x4e, 0x48, 0x48,
	0x44, 0x4d, 0x4a, 0x50, 0x48, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x42, 0x4f, 0x47, 0x4a, 0x4f,
	0x48, 0x49, 0x47, 0x4a, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x42, 0x4f,
	0x47, 0x4a, 0x4f, 0x48, 0x49, 0x47, 0x4a, 0x4c, 0x12, 0x31, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x0a, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x41, 0x47, 0x41, 0x49, 0x48, 0x49, 0x45, 0x4c, 0x45,
	0x4e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x41, 0x47, 0x41, 0x49, 0x48, 0x49,
	0x45, 0x4c, 0x45, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x45, 0x4a, 0x4a, 0x46, 0x4e, 0x50, 0x4a,
	0x41, 0x4f, 0x41, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x45, 0x4a, 0x4a, 0x46,
	0x4e, 0x50, 0x4a, 0x41, 0x4f, 0x41, 0x12, 0x2b, 0x0a, 0x0b, 0x4b, 0x46, 0x4a, 0x4c, 0x48, 0x45,
	0x42, 0x44, 0x50, 0x4e, 0x47, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x4b, 0x46, 0x4a, 0x4c, 0x48, 0x45, 0x42, 0x44,
	0x50, 0x4e, 0x47, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x77, 0x69, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x57, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x72,
	0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0xeb, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x72,
	0x65, 0x61, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x4b, 0x4a, 0x4e, 0x49, 0x41, 0x43, 0x4a,
	0x42, 0x46, 0x4b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4b, 0x4b, 0x4a, 0x4e, 0x49,
	0x41, 0x43, 0x4a, 0x42, 0x46, 0x4b, 0x12, 0x37, 0x0a, 0x0b, 0x4e, 0x43, 0x42, 0x48, 0x4c, 0x42,
	0x4a, 0x41, 0x41, 0x4a, 0x4b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0b, 0x4e, 0x43, 0x42, 0x48, 0x4c, 0x42, 0x4a, 0x41, 0x41, 0x4a, 0x4b, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_RogueFinishInfo_proto_rawDescOnce sync.Once
	file_RogueFinishInfo_proto_rawDescData = file_RogueFinishInfo_proto_rawDesc
)

func file_RogueFinishInfo_proto_rawDescGZIP() []byte {
	file_RogueFinishInfo_proto_rawDescOnce.Do(func() {
		file_RogueFinishInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueFinishInfo_proto_rawDescData)
	})
	return file_RogueFinishInfo_proto_rawDescData
}

var file_RogueFinishInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueFinishInfo_proto_goTypes = []interface{}{
	(*RogueFinishInfo)(nil),       // 0: RogueFinishInfo
	(*MLPPNJOPPML)(nil),           // 1: MLPPNJOPPML
	(*ItemList)(nil),              // 2: ItemList
	(*RogueScoreRewardInfo)(nil),  // 3: RogueScoreRewardInfo
	(*RogueRecordInfo)(nil),       // 4: RogueRecordInfo
	(*RogueExploreScoreInfo)(nil), // 5: RogueExploreScoreInfo
}
var file_RogueFinishInfo_proto_depIdxs = []int32{
	1, // 0: RogueFinishInfo.KAGHBIBAIPP:type_name -> MLPPNJOPPML
	2, // 1: RogueFinishInfo.FJHJDMNBKDO:type_name -> ItemList
	2, // 2: RogueFinishInfo.OCMOCECJGHJ:type_name -> ItemList
	3, // 3: RogueFinishInfo.DNNHHDMJPHO:type_name -> RogueScoreRewardInfo
	4, // 4: RogueFinishInfo.record_info:type_name -> RogueRecordInfo
	5, // 5: RogueFinishInfo.score_info:type_name -> RogueExploreScoreInfo
	2, // 6: RogueFinishInfo.KFJLHEBDPNG:type_name -> ItemList
	3, // 7: RogueFinishInfo.NCBHLBJAAJK:type_name -> RogueScoreRewardInfo
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_RogueFinishInfo_proto_init() }
func file_RogueFinishInfo_proto_init() {
	if File_RogueFinishInfo_proto != nil {
		return
	}
	file_ItemList_proto_init()
	file_RogueRecordInfo_proto_init()
	file_RogueScoreRewardInfo_proto_init()
	file_MLPPNJOPPML_proto_init()
	file_RogueExploreScoreInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueFinishInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueFinishInfo); i {
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
			RawDescriptor: file_RogueFinishInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueFinishInfo_proto_goTypes,
		DependencyIndexes: file_RogueFinishInfo_proto_depIdxs,
		MessageInfos:      file_RogueFinishInfo_proto_msgTypes,
	}.Build()
	File_RogueFinishInfo_proto = out.File
	file_RogueFinishInfo_proto_rawDesc = nil
	file_RogueFinishInfo_proto_goTypes = nil
	file_RogueFinishInfo_proto_depIdxs = nil
}
