// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EvolveBuildShopAbilityResetScRsp.proto

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

type EvolveBuildShopAbilityResetScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ENDPIBIJMPO []*GNKBFNAAEED `protobuf:"bytes,9,rep,name=ENDPIBIJMPO,proto3" json:"ENDPIBIJMPO,omitempty"`
	Retcode     uint32         `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ItemValue   uint32         `protobuf:"varint,7,opt,name=item_value,json=itemValue,proto3" json:"item_value,omitempty"`
}

func (x *EvolveBuildShopAbilityResetScRsp) Reset() {
	*x = EvolveBuildShopAbilityResetScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvolveBuildShopAbilityResetScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvolveBuildShopAbilityResetScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvolveBuildShopAbilityResetScRsp) ProtoMessage() {}

func (x *EvolveBuildShopAbilityResetScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_EvolveBuildShopAbilityResetScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvolveBuildShopAbilityResetScRsp.ProtoReflect.Descriptor instead.
func (*EvolveBuildShopAbilityResetScRsp) Descriptor() ([]byte, []int) {
	return file_EvolveBuildShopAbilityResetScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *EvolveBuildShopAbilityResetScRsp) GetENDPIBIJMPO() []*GNKBFNAAEED {
	if x != nil {
		return x.ENDPIBIJMPO
	}
	return nil
}

func (x *EvolveBuildShopAbilityResetScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *EvolveBuildShopAbilityResetScRsp) GetItemValue() uint32 {
	if x != nil {
		return x.ItemValue
	}
	return 0
}

var File_EvolveBuildShopAbilityResetScRsp_proto protoreflect.FileDescriptor

var file_EvolveBuildShopAbilityResetScRsp_proto_rawDesc = []byte{
	0x0a, 0x26, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x68, 0x6f,
	0x70, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x4e, 0x4b, 0x42, 0x46, 0x4e,
	0x41, 0x41, 0x45, 0x45, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x20,
	0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x68, 0x6f, 0x70, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x4e, 0x44, 0x50, 0x49, 0x42, 0x49, 0x4a, 0x4d, 0x50, 0x4f, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x4e, 0x4b, 0x42, 0x46, 0x4e, 0x41, 0x41,
	0x45, 0x45, 0x44, 0x52, 0x0b, 0x45, 0x4e, 0x44, 0x50, 0x49, 0x42, 0x49, 0x4a, 0x4d, 0x50, 0x4f,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x69, 0x74, 0x65, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_EvolveBuildShopAbilityResetScRsp_proto_rawDescOnce sync.Once
	file_EvolveBuildShopAbilityResetScRsp_proto_rawDescData = file_EvolveBuildShopAbilityResetScRsp_proto_rawDesc
)

func file_EvolveBuildShopAbilityResetScRsp_proto_rawDescGZIP() []byte {
	file_EvolveBuildShopAbilityResetScRsp_proto_rawDescOnce.Do(func() {
		file_EvolveBuildShopAbilityResetScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvolveBuildShopAbilityResetScRsp_proto_rawDescData)
	})
	return file_EvolveBuildShopAbilityResetScRsp_proto_rawDescData
}

var file_EvolveBuildShopAbilityResetScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvolveBuildShopAbilityResetScRsp_proto_goTypes = []interface{}{
	(*EvolveBuildShopAbilityResetScRsp)(nil), // 0: EvolveBuildShopAbilityResetScRsp
	(*GNKBFNAAEED)(nil),                      // 1: GNKBFNAAEED
}
var file_EvolveBuildShopAbilityResetScRsp_proto_depIdxs = []int32{
	1, // 0: EvolveBuildShopAbilityResetScRsp.ENDPIBIJMPO:type_name -> GNKBFNAAEED
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EvolveBuildShopAbilityResetScRsp_proto_init() }
func file_EvolveBuildShopAbilityResetScRsp_proto_init() {
	if File_EvolveBuildShopAbilityResetScRsp_proto != nil {
		return
	}
	file_GNKBFNAAEED_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvolveBuildShopAbilityResetScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvolveBuildShopAbilityResetScRsp); i {
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
			RawDescriptor: file_EvolveBuildShopAbilityResetScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvolveBuildShopAbilityResetScRsp_proto_goTypes,
		DependencyIndexes: file_EvolveBuildShopAbilityResetScRsp_proto_depIdxs,
		MessageInfos:      file_EvolveBuildShopAbilityResetScRsp_proto_msgTypes,
	}.Build()
	File_EvolveBuildShopAbilityResetScRsp_proto = out.File
	file_EvolveBuildShopAbilityResetScRsp_proto_rawDesc = nil
	file_EvolveBuildShopAbilityResetScRsp_proto_goTypes = nil
	file_EvolveBuildShopAbilityResetScRsp_proto_depIdxs = nil
}
