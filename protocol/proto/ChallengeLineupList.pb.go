// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChallengeLineupList.proto

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

type ChallengeLineupList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarList []*ChallengeAvatarInfo `protobuf:"bytes,4,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
}

func (x *ChallengeLineupList) Reset() {
	*x = ChallengeLineupList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeLineupList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeLineupList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeLineupList) ProtoMessage() {}

func (x *ChallengeLineupList) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeLineupList_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeLineupList.ProtoReflect.Descriptor instead.
func (*ChallengeLineupList) Descriptor() ([]byte, []int) {
	return file_ChallengeLineupList_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeLineupList) GetAvatarList() []*ChallengeAvatarInfo {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

var File_ChallengeLineupList_proto protoreflect.FileDescriptor

var file_ChallengeLineupList_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a,
	0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeLineupList_proto_rawDescOnce sync.Once
	file_ChallengeLineupList_proto_rawDescData = file_ChallengeLineupList_proto_rawDesc
)

func file_ChallengeLineupList_proto_rawDescGZIP() []byte {
	file_ChallengeLineupList_proto_rawDescOnce.Do(func() {
		file_ChallengeLineupList_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeLineupList_proto_rawDescData)
	})
	return file_ChallengeLineupList_proto_rawDescData
}

var file_ChallengeLineupList_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChallengeLineupList_proto_goTypes = []interface{}{
	(*ChallengeLineupList)(nil), // 0: ChallengeLineupList
	(*ChallengeAvatarInfo)(nil), // 1: ChallengeAvatarInfo
}
var file_ChallengeLineupList_proto_depIdxs = []int32{
	1, // 0: ChallengeLineupList.avatar_list:type_name -> ChallengeAvatarInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChallengeLineupList_proto_init() }
func file_ChallengeLineupList_proto_init() {
	if File_ChallengeLineupList_proto != nil {
		return
	}
	file_ChallengeAvatarInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChallengeLineupList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeLineupList); i {
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
			RawDescriptor: file_ChallengeLineupList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeLineupList_proto_goTypes,
		DependencyIndexes: file_ChallengeLineupList_proto_depIdxs,
		MessageInfos:      file_ChallengeLineupList_proto_msgTypes,
	}.Build()
	File_ChallengeLineupList_proto = out.File
	file_ChallengeLineupList_proto_rawDesc = nil
	file_ChallengeLineupList_proto_goTypes = nil
	file_ChallengeLineupList_proto_depIdxs = nil
}
