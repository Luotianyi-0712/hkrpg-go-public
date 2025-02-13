// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FinishPerformSectionIdCsReq.proto

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

type FinishPerformSectionIdCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SectionId uint32         `protobuf:"varint,4,opt,name=section_id,json=sectionId,proto3" json:"section_id,omitempty"`
	ItemList  []*MessageItem `protobuf:"bytes,3,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
}

func (x *FinishPerformSectionIdCsReq) Reset() {
	*x = FinishPerformSectionIdCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FinishPerformSectionIdCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishPerformSectionIdCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishPerformSectionIdCsReq) ProtoMessage() {}

func (x *FinishPerformSectionIdCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_FinishPerformSectionIdCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishPerformSectionIdCsReq.ProtoReflect.Descriptor instead.
func (*FinishPerformSectionIdCsReq) Descriptor() ([]byte, []int) {
	return file_FinishPerformSectionIdCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *FinishPerformSectionIdCsReq) GetSectionId() uint32 {
	if x != nil {
		return x.SectionId
	}
	return 0
}

func (x *FinishPerformSectionIdCsReq) GetItemList() []*MessageItem {
	if x != nil {
		return x.ItemList
	}
	return nil
}

var File_FinishPerformSectionIdCsReq_proto protoreflect.FileDescriptor

var file_FinishPerformSectionIdCsReq_proto_rawDesc = []byte{
	0x0a, 0x21, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FinishPerformSectionIdCsReq_proto_rawDescOnce sync.Once
	file_FinishPerformSectionIdCsReq_proto_rawDescData = file_FinishPerformSectionIdCsReq_proto_rawDesc
)

func file_FinishPerformSectionIdCsReq_proto_rawDescGZIP() []byte {
	file_FinishPerformSectionIdCsReq_proto_rawDescOnce.Do(func() {
		file_FinishPerformSectionIdCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_FinishPerformSectionIdCsReq_proto_rawDescData)
	})
	return file_FinishPerformSectionIdCsReq_proto_rawDescData
}

var file_FinishPerformSectionIdCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FinishPerformSectionIdCsReq_proto_goTypes = []interface{}{
	(*FinishPerformSectionIdCsReq)(nil), // 0: FinishPerformSectionIdCsReq
	(*MessageItem)(nil),                 // 1: MessageItem
}
var file_FinishPerformSectionIdCsReq_proto_depIdxs = []int32{
	1, // 0: FinishPerformSectionIdCsReq.item_list:type_name -> MessageItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FinishPerformSectionIdCsReq_proto_init() }
func file_FinishPerformSectionIdCsReq_proto_init() {
	if File_FinishPerformSectionIdCsReq_proto != nil {
		return
	}
	file_MessageItem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FinishPerformSectionIdCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishPerformSectionIdCsReq); i {
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
			RawDescriptor: file_FinishPerformSectionIdCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FinishPerformSectionIdCsReq_proto_goTypes,
		DependencyIndexes: file_FinishPerformSectionIdCsReq_proto_depIdxs,
		MessageInfos:      file_FinishPerformSectionIdCsReq_proto_msgTypes,
	}.Build()
	File_FinishPerformSectionIdCsReq_proto = out.File
	file_FinishPerformSectionIdCsReq_proto_rawDesc = nil
	file_FinishPerformSectionIdCsReq_proto_goTypes = nil
	file_FinishPerformSectionIdCsReq_proto_depIdxs = nil
}
