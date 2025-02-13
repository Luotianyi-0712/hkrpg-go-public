// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DOJCGNJFHHD.proto

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

type DOJCGNJFHHD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeadIcon    uint32       `protobuf:"varint,4,opt,name=head_icon,json=headIcon,proto3" json:"head_icon,omitempty"`
	EKONPNJPPAF uint32       `protobuf:"varint,7,opt,name=EKONPNJPPAF,proto3" json:"EKONPNJPPAF,omitempty"`
	Platform    PlatformType `protobuf:"varint,9,opt,name=platform,proto3,enum=PlatformType" json:"platform,omitempty"`
	MCKKBKPOMLI string       `protobuf:"bytes,1,opt,name=MCKKBKPOMLI,proto3" json:"MCKKBKPOMLI,omitempty"`
	Level       uint32       `protobuf:"varint,10,opt,name=level,proto3" json:"level,omitempty"`
	RemarkName  string       `protobuf:"bytes,15,opt,name=remark_name,json=remarkName,proto3" json:"remark_name,omitempty"`
	Uid         uint32       `protobuf:"varint,13,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname    string       `protobuf:"bytes,12,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *DOJCGNJFHHD) Reset() {
	*x = DOJCGNJFHHD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DOJCGNJFHHD_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DOJCGNJFHHD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DOJCGNJFHHD) ProtoMessage() {}

func (x *DOJCGNJFHHD) ProtoReflect() protoreflect.Message {
	mi := &file_DOJCGNJFHHD_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DOJCGNJFHHD.ProtoReflect.Descriptor instead.
func (*DOJCGNJFHHD) Descriptor() ([]byte, []int) {
	return file_DOJCGNJFHHD_proto_rawDescGZIP(), []int{0}
}

func (x *DOJCGNJFHHD) GetHeadIcon() uint32 {
	if x != nil {
		return x.HeadIcon
	}
	return 0
}

func (x *DOJCGNJFHHD) GetEKONPNJPPAF() uint32 {
	if x != nil {
		return x.EKONPNJPPAF
	}
	return 0
}

func (x *DOJCGNJFHHD) GetPlatform() PlatformType {
	if x != nil {
		return x.Platform
	}
	return PlatformType_EDITOR
}

func (x *DOJCGNJFHHD) GetMCKKBKPOMLI() string {
	if x != nil {
		return x.MCKKBKPOMLI
	}
	return ""
}

func (x *DOJCGNJFHHD) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *DOJCGNJFHHD) GetRemarkName() string {
	if x != nil {
		return x.RemarkName
	}
	return ""
}

func (x *DOJCGNJFHHD) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DOJCGNJFHHD) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

var File_DOJCGNJFHHD_proto protoreflect.FileDescriptor

var file_DOJCGNJFHHD_proto_rawDesc = []byte{
	0x0a, 0x11, 0x44, 0x4f, 0x4a, 0x43, 0x47, 0x4e, 0x4a, 0x46, 0x48, 0x48, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x0b, 0x44, 0x4f, 0x4a, 0x43,
	0x47, 0x4e, 0x4a, 0x46, 0x48, 0x48, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x64, 0x5f,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64,
	0x49, 0x63, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x4b, 0x4f, 0x4e, 0x50, 0x4e, 0x4a, 0x50,
	0x50, 0x41, 0x46, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x4b, 0x4f, 0x4e, 0x50,
	0x4e, 0x4a, 0x50, 0x50, 0x41, 0x46, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x43, 0x4b, 0x4b, 0x42, 0x4b, 0x50, 0x4f, 0x4d, 0x4c, 0x49,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x43, 0x4b, 0x4b, 0x42, 0x4b, 0x50, 0x4f,
	0x4d, 0x4c, 0x49, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DOJCGNJFHHD_proto_rawDescOnce sync.Once
	file_DOJCGNJFHHD_proto_rawDescData = file_DOJCGNJFHHD_proto_rawDesc
)

func file_DOJCGNJFHHD_proto_rawDescGZIP() []byte {
	file_DOJCGNJFHHD_proto_rawDescOnce.Do(func() {
		file_DOJCGNJFHHD_proto_rawDescData = protoimpl.X.CompressGZIP(file_DOJCGNJFHHD_proto_rawDescData)
	})
	return file_DOJCGNJFHHD_proto_rawDescData
}

var file_DOJCGNJFHHD_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DOJCGNJFHHD_proto_goTypes = []interface{}{
	(*DOJCGNJFHHD)(nil), // 0: DOJCGNJFHHD
	(PlatformType)(0),   // 1: PlatformType
}
var file_DOJCGNJFHHD_proto_depIdxs = []int32{
	1, // 0: DOJCGNJFHHD.platform:type_name -> PlatformType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DOJCGNJFHHD_proto_init() }
func file_DOJCGNJFHHD_proto_init() {
	if File_DOJCGNJFHHD_proto != nil {
		return
	}
	file_PlatformType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DOJCGNJFHHD_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DOJCGNJFHHD); i {
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
			RawDescriptor: file_DOJCGNJFHHD_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DOJCGNJFHHD_proto_goTypes,
		DependencyIndexes: file_DOJCGNJFHHD_proto_depIdxs,
		MessageInfos:      file_DOJCGNJFHHD_proto_msgTypes,
	}.Build()
	File_DOJCGNJFHHD_proto = out.File
	file_DOJCGNJFHHD_proto_rawDesc = nil
	file_DOJCGNJFHHD_proto_goTypes = nil
	file_DOJCGNJFHHD_proto_depIdxs = nil
}
