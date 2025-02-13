// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EnhanceRogueBuffScRsp.proto

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

type EnhanceRogueBuffScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode   uint32     `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	IsSuccess bool       `protobuf:"varint,3,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	RogueBuff *RogueBuff `protobuf:"bytes,11,opt,name=rogue_buff,json=rogueBuff,proto3" json:"rogue_buff,omitempty"`
}

func (x *EnhanceRogueBuffScRsp) Reset() {
	*x = EnhanceRogueBuffScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnhanceRogueBuffScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnhanceRogueBuffScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnhanceRogueBuffScRsp) ProtoMessage() {}

func (x *EnhanceRogueBuffScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_EnhanceRogueBuffScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnhanceRogueBuffScRsp.ProtoReflect.Descriptor instead.
func (*EnhanceRogueBuffScRsp) Descriptor() ([]byte, []int) {
	return file_EnhanceRogueBuffScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *EnhanceRogueBuffScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *EnhanceRogueBuffScRsp) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *EnhanceRogueBuffScRsp) GetRogueBuff() *RogueBuff {
	if x != nil {
		return x.RogueBuff
	}
	return nil
}

var File_EnhanceRogueBuffScRsp_proto protoreflect.FileDescriptor

var file_EnhanceRogueBuffScRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75,
	0x66, 0x66, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b,
	0x0a, 0x15, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75,
	0x66, 0x66, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x29, 0x0a, 0x0a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66,
	0x52, 0x09, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_EnhanceRogueBuffScRsp_proto_rawDescOnce sync.Once
	file_EnhanceRogueBuffScRsp_proto_rawDescData = file_EnhanceRogueBuffScRsp_proto_rawDesc
)

func file_EnhanceRogueBuffScRsp_proto_rawDescGZIP() []byte {
	file_EnhanceRogueBuffScRsp_proto_rawDescOnce.Do(func() {
		file_EnhanceRogueBuffScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnhanceRogueBuffScRsp_proto_rawDescData)
	})
	return file_EnhanceRogueBuffScRsp_proto_rawDescData
}

var file_EnhanceRogueBuffScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnhanceRogueBuffScRsp_proto_goTypes = []interface{}{
	(*EnhanceRogueBuffScRsp)(nil), // 0: EnhanceRogueBuffScRsp
	(*RogueBuff)(nil),             // 1: RogueBuff
}
var file_EnhanceRogueBuffScRsp_proto_depIdxs = []int32{
	1, // 0: EnhanceRogueBuffScRsp.rogue_buff:type_name -> RogueBuff
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EnhanceRogueBuffScRsp_proto_init() }
func file_EnhanceRogueBuffScRsp_proto_init() {
	if File_EnhanceRogueBuffScRsp_proto != nil {
		return
	}
	file_RogueBuff_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EnhanceRogueBuffScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnhanceRogueBuffScRsp); i {
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
			RawDescriptor: file_EnhanceRogueBuffScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnhanceRogueBuffScRsp_proto_goTypes,
		DependencyIndexes: file_EnhanceRogueBuffScRsp_proto_depIdxs,
		MessageInfos:      file_EnhanceRogueBuffScRsp_proto_msgTypes,
	}.Build()
	File_EnhanceRogueBuffScRsp_proto = out.File
	file_EnhanceRogueBuffScRsp_proto_rawDesc = nil
	file_EnhanceRogueBuffScRsp_proto_goTypes = nil
	file_EnhanceRogueBuffScRsp_proto_depIdxs = nil
}
