// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetPhoneDataScRsp.proto

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

type GetPhoneDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurPhoneTheme    uint32   `protobuf:"varint,14,opt,name=cur_phone_theme,json=curPhoneTheme,proto3" json:"cur_phone_theme,omitempty"`
	OwnedPhoneThemes []uint32 `protobuf:"varint,13,rep,packed,name=owned_phone_themes,json=ownedPhoneThemes,proto3" json:"owned_phone_themes,omitempty"`
	CurChatBubble    uint32   `protobuf:"varint,15,opt,name=cur_chat_bubble,json=curChatBubble,proto3" json:"cur_chat_bubble,omitempty"`
	OwnedChatBubbles []uint32 `protobuf:"varint,5,rep,packed,name=owned_chat_bubbles,json=ownedChatBubbles,proto3" json:"owned_chat_bubbles,omitempty"`
	Retcode          uint32   `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetPhoneDataScRsp) Reset() {
	*x = GetPhoneDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPhoneDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPhoneDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPhoneDataScRsp) ProtoMessage() {}

func (x *GetPhoneDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetPhoneDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPhoneDataScRsp.ProtoReflect.Descriptor instead.
func (*GetPhoneDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetPhoneDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPhoneDataScRsp) GetCurPhoneTheme() uint32 {
	if x != nil {
		return x.CurPhoneTheme
	}
	return 0
}

func (x *GetPhoneDataScRsp) GetOwnedPhoneThemes() []uint32 {
	if x != nil {
		return x.OwnedPhoneThemes
	}
	return nil
}

func (x *GetPhoneDataScRsp) GetCurChatBubble() uint32 {
	if x != nil {
		return x.CurChatBubble
	}
	return 0
}

func (x *GetPhoneDataScRsp) GetOwnedChatBubbles() []uint32 {
	if x != nil {
		return x.OwnedChatBubbles
	}
	return nil
}

func (x *GetPhoneDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetPhoneDataScRsp_proto protoreflect.FileDescriptor

var file_GetPhoneDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12,
	0x26, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x74, 0x68, 0x65,
	0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x77, 0x6e, 0x65, 0x64,
	0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x10, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54,
	0x68, 0x65, 0x6d, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x5f, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x62, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x63, 0x75, 0x72, 0x43, 0x68, 0x61, 0x74, 0x42, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x12, 0x2c, 0x0a,
	0x12, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x62, 0x75, 0x62, 0x62,
	0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x6f, 0x77, 0x6e, 0x65, 0x64,
	0x43, 0x68, 0x61, 0x74, 0x42, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetPhoneDataScRsp_proto_rawDescOnce sync.Once
	file_GetPhoneDataScRsp_proto_rawDescData = file_GetPhoneDataScRsp_proto_rawDesc
)

func file_GetPhoneDataScRsp_proto_rawDescGZIP() []byte {
	file_GetPhoneDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetPhoneDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPhoneDataScRsp_proto_rawDescData)
	})
	return file_GetPhoneDataScRsp_proto_rawDescData
}

var file_GetPhoneDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPhoneDataScRsp_proto_goTypes = []interface{}{
	(*GetPhoneDataScRsp)(nil), // 0: GetPhoneDataScRsp
}
var file_GetPhoneDataScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetPhoneDataScRsp_proto_init() }
func file_GetPhoneDataScRsp_proto_init() {
	if File_GetPhoneDataScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetPhoneDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPhoneDataScRsp); i {
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
			RawDescriptor: file_GetPhoneDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPhoneDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetPhoneDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetPhoneDataScRsp_proto_msgTypes,
	}.Build()
	File_GetPhoneDataScRsp_proto = out.File
	file_GetPhoneDataScRsp_proto_rawDesc = nil
	file_GetPhoneDataScRsp_proto_goTypes = nil
	file_GetPhoneDataScRsp_proto_depIdxs = nil
}
