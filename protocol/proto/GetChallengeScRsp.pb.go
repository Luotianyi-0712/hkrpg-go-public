// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetChallengeScRsp.proto

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

type GetChallengeScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeGroupList []*ChallengeGroup `protobuf:"bytes,4,rep,name=challenge_group_list,json=challengeGroupList,proto3" json:"challenge_group_list,omitempty"`
	ChallengeList      []*Challenge      `protobuf:"bytes,9,rep,name=challenge_list,json=challengeList,proto3" json:"challenge_list,omitempty"`
	OKGCOBHLIIM        uint32            `protobuf:"varint,8,opt,name=OKGCOBHLIIM,proto3" json:"OKGCOBHLIIM,omitempty"`
	BCGCOGHPHPP        []*NMHNANJAINM    `protobuf:"bytes,2,rep,name=BCGCOGHPHPP,proto3" json:"BCGCOGHPHPP,omitempty"`
	Retcode            uint32            `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetChallengeScRsp) Reset() {
	*x = GetChallengeScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetChallengeScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChallengeScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChallengeScRsp) ProtoMessage() {}

func (x *GetChallengeScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetChallengeScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChallengeScRsp.ProtoReflect.Descriptor instead.
func (*GetChallengeScRsp) Descriptor() ([]byte, []int) {
	return file_GetChallengeScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetChallengeScRsp) GetChallengeGroupList() []*ChallengeGroup {
	if x != nil {
		return x.ChallengeGroupList
	}
	return nil
}

func (x *GetChallengeScRsp) GetChallengeList() []*Challenge {
	if x != nil {
		return x.ChallengeList
	}
	return nil
}

func (x *GetChallengeScRsp) GetOKGCOBHLIIM() uint32 {
	if x != nil {
		return x.OKGCOBHLIIM
	}
	return 0
}

func (x *GetChallengeScRsp) GetBCGCOGHPHPP() []*NMHNANJAINM {
	if x != nil {
		return x.BCGCOGHPHPP
	}
	return nil
}

func (x *GetChallengeScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetChallengeScRsp_proto protoreflect.FileDescriptor

var file_GetChallengeScRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x4d, 0x48, 0x4e,
	0x41, 0x4e, 0x4a, 0x41, 0x49, 0x4e, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x14, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x12, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0e,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x52, 0x0d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x4b, 0x47, 0x43, 0x4f, 0x42, 0x48, 0x4c, 0x49, 0x49, 0x4d, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4b, 0x47, 0x43, 0x4f, 0x42, 0x48, 0x4c, 0x49, 0x49,
	0x4d, 0x12, 0x2e, 0x0a, 0x0b, 0x42, 0x43, 0x47, 0x43, 0x4f, 0x47, 0x48, 0x50, 0x48, 0x50, 0x50,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x4d, 0x48, 0x4e, 0x41, 0x4e, 0x4a,
	0x41, 0x49, 0x4e, 0x4d, 0x52, 0x0b, 0x42, 0x43, 0x47, 0x43, 0x4f, 0x47, 0x48, 0x50, 0x48, 0x50,
	0x50, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetChallengeScRsp_proto_rawDescOnce sync.Once
	file_GetChallengeScRsp_proto_rawDescData = file_GetChallengeScRsp_proto_rawDesc
)

func file_GetChallengeScRsp_proto_rawDescGZIP() []byte {
	file_GetChallengeScRsp_proto_rawDescOnce.Do(func() {
		file_GetChallengeScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetChallengeScRsp_proto_rawDescData)
	})
	return file_GetChallengeScRsp_proto_rawDescData
}

var file_GetChallengeScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetChallengeScRsp_proto_goTypes = []interface{}{
	(*GetChallengeScRsp)(nil), // 0: GetChallengeScRsp
	(*ChallengeGroup)(nil),    // 1: ChallengeGroup
	(*Challenge)(nil),         // 2: Challenge
	(*NMHNANJAINM)(nil),       // 3: NMHNANJAINM
}
var file_GetChallengeScRsp_proto_depIdxs = []int32{
	1, // 0: GetChallengeScRsp.challenge_group_list:type_name -> ChallengeGroup
	2, // 1: GetChallengeScRsp.challenge_list:type_name -> Challenge
	3, // 2: GetChallengeScRsp.BCGCOGHPHPP:type_name -> NMHNANJAINM
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetChallengeScRsp_proto_init() }
func file_GetChallengeScRsp_proto_init() {
	if File_GetChallengeScRsp_proto != nil {
		return
	}
	file_Challenge_proto_init()
	file_NMHNANJAINM_proto_init()
	file_ChallengeGroup_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetChallengeScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChallengeScRsp); i {
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
			RawDescriptor: file_GetChallengeScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetChallengeScRsp_proto_goTypes,
		DependencyIndexes: file_GetChallengeScRsp_proto_depIdxs,
		MessageInfos:      file_GetChallengeScRsp_proto_msgTypes,
	}.Build()
	File_GetChallengeScRsp_proto = out.File
	file_GetChallengeScRsp_proto_rawDesc = nil
	file_GetChallengeScRsp_proto_goTypes = nil
	file_GetChallengeScRsp_proto_depIdxs = nil
}
