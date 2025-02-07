// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SubmitEmotionItemCsReq.proto

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

type SubmitEmotionItemCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PKBMEHAFBCH uint32    `protobuf:"varint,2,opt,name=PKBMEHAFBCH,proto3" json:"PKBMEHAFBCH,omitempty"`
	ScriptId    uint32    `protobuf:"varint,13,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty"`
	ItemList    *ItemList `protobuf:"bytes,6,opt,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
}

func (x *SubmitEmotionItemCsReq) Reset() {
	*x = SubmitEmotionItemCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SubmitEmotionItemCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitEmotionItemCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitEmotionItemCsReq) ProtoMessage() {}

func (x *SubmitEmotionItemCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SubmitEmotionItemCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitEmotionItemCsReq.ProtoReflect.Descriptor instead.
func (*SubmitEmotionItemCsReq) Descriptor() ([]byte, []int) {
	return file_SubmitEmotionItemCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitEmotionItemCsReq) GetPKBMEHAFBCH() uint32 {
	if x != nil {
		return x.PKBMEHAFBCH
	}
	return 0
}

func (x *SubmitEmotionItemCsReq) GetScriptId() uint32 {
	if x != nil {
		return x.ScriptId
	}
	return 0
}

func (x *SubmitEmotionItemCsReq) GetItemList() *ItemList {
	if x != nil {
		return x.ItemList
	}
	return nil
}

var File_SubmitEmotionItemCsReq_proto protoreflect.FileDescriptor

var file_SubmitEmotionItemCsReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f,
	0x0a, 0x16, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x4b, 0x42, 0x4d,
	0x45, 0x48, 0x41, 0x46, 0x42, 0x43, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50,
	0x4b, 0x42, 0x4d, 0x45, 0x48, 0x41, 0x46, 0x42, 0x43, 0x48, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SubmitEmotionItemCsReq_proto_rawDescOnce sync.Once
	file_SubmitEmotionItemCsReq_proto_rawDescData = file_SubmitEmotionItemCsReq_proto_rawDesc
)

func file_SubmitEmotionItemCsReq_proto_rawDescGZIP() []byte {
	file_SubmitEmotionItemCsReq_proto_rawDescOnce.Do(func() {
		file_SubmitEmotionItemCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SubmitEmotionItemCsReq_proto_rawDescData)
	})
	return file_SubmitEmotionItemCsReq_proto_rawDescData
}

var file_SubmitEmotionItemCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SubmitEmotionItemCsReq_proto_goTypes = []interface{}{
	(*SubmitEmotionItemCsReq)(nil), // 0: SubmitEmotionItemCsReq
	(*ItemList)(nil),               // 1: ItemList
}
var file_SubmitEmotionItemCsReq_proto_depIdxs = []int32{
	1, // 0: SubmitEmotionItemCsReq.item_list:type_name -> ItemList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SubmitEmotionItemCsReq_proto_init() }
func file_SubmitEmotionItemCsReq_proto_init() {
	if File_SubmitEmotionItemCsReq_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SubmitEmotionItemCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitEmotionItemCsReq); i {
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
			RawDescriptor: file_SubmitEmotionItemCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SubmitEmotionItemCsReq_proto_goTypes,
		DependencyIndexes: file_SubmitEmotionItemCsReq_proto_depIdxs,
		MessageInfos:      file_SubmitEmotionItemCsReq_proto_msgTypes,
	}.Build()
	File_SubmitEmotionItemCsReq_proto = out.File
	file_SubmitEmotionItemCsReq_proto_rawDesc = nil
	file_SubmitEmotionItemCsReq_proto_goTypes = nil
	file_SubmitEmotionItemCsReq_proto_depIdxs = nil
}
