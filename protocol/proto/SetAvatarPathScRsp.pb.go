// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SetAvatarPathScRsp.proto

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

type SetAvatarPathScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId MultiPathAvatarType `protobuf:"varint,9,opt,name=avatar_id,json=avatarId,proto3,enum=MultiPathAvatarType" json:"avatar_id,omitempty"`
	Retcode  uint32              `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SetAvatarPathScRsp) Reset() {
	*x = SetAvatarPathScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetAvatarPathScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAvatarPathScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAvatarPathScRsp) ProtoMessage() {}

func (x *SetAvatarPathScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetAvatarPathScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAvatarPathScRsp.ProtoReflect.Descriptor instead.
func (*SetAvatarPathScRsp) Descriptor() ([]byte, []int) {
	return file_SetAvatarPathScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetAvatarPathScRsp) GetAvatarId() MultiPathAvatarType {
	if x != nil {
		return x.AvatarId
	}
	return MultiPathAvatarType_MultiPathAvatarTypeNone
}

func (x *SetAvatarPathScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SetAvatarPathScRsp_proto protoreflect.FileDescriptor

var file_SetAvatarPathScRsp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x50, 0x61, 0x74, 0x68, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x09, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x50, 0x61, 0x74, 0x68, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetAvatarPathScRsp_proto_rawDescOnce sync.Once
	file_SetAvatarPathScRsp_proto_rawDescData = file_SetAvatarPathScRsp_proto_rawDesc
)

func file_SetAvatarPathScRsp_proto_rawDescGZIP() []byte {
	file_SetAvatarPathScRsp_proto_rawDescOnce.Do(func() {
		file_SetAvatarPathScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetAvatarPathScRsp_proto_rawDescData)
	})
	return file_SetAvatarPathScRsp_proto_rawDescData
}

var file_SetAvatarPathScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetAvatarPathScRsp_proto_goTypes = []interface{}{
	(*SetAvatarPathScRsp)(nil), // 0: SetAvatarPathScRsp
	(MultiPathAvatarType)(0),   // 1: MultiPathAvatarType
}
var file_SetAvatarPathScRsp_proto_depIdxs = []int32{
	1, // 0: SetAvatarPathScRsp.avatar_id:type_name -> MultiPathAvatarType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SetAvatarPathScRsp_proto_init() }
func file_SetAvatarPathScRsp_proto_init() {
	if File_SetAvatarPathScRsp_proto != nil {
		return
	}
	file_MultiPathAvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SetAvatarPathScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAvatarPathScRsp); i {
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
			RawDescriptor: file_SetAvatarPathScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetAvatarPathScRsp_proto_goTypes,
		DependencyIndexes: file_SetAvatarPathScRsp_proto_depIdxs,
		MessageInfos:      file_SetAvatarPathScRsp_proto_msgTypes,
	}.Build()
	File_SetAvatarPathScRsp_proto = out.File
	file_SetAvatarPathScRsp_proto_rawDesc = nil
	file_SetAvatarPathScRsp_proto_goTypes = nil
	file_SetAvatarPathScRsp_proto_depIdxs = nil
}
