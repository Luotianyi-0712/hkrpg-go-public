// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SetRedPointStatusScNotify.proto

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

type SetRedPointStatusScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         uint32         `protobuf:"varint,6,opt,name=uid,proto3" json:"uid,omitempty"`
	JGMIDHMLNFK uint32         `protobuf:"varint,12,opt,name=JGMIDHMLNFK,proto3" json:"JGMIDHMLNFK,omitempty"`
	DFGNMBNLMJI uint32         `protobuf:"varint,11,opt,name=DFGNMBNLMJI,proto3" json:"DFGNMBNLMJI,omitempty"`
	NBFOOFHBOPC []*ALLNGMEDPMD `protobuf:"bytes,14,rep,name=NBFOOFHBOPC,proto3" json:"NBFOOFHBOPC,omitempty"`
	ContentId   uint32         `protobuf:"varint,9,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
}

func (x *SetRedPointStatusScNotify) Reset() {
	*x = SetRedPointStatusScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetRedPointStatusScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRedPointStatusScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRedPointStatusScNotify) ProtoMessage() {}

func (x *SetRedPointStatusScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SetRedPointStatusScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRedPointStatusScNotify.ProtoReflect.Descriptor instead.
func (*SetRedPointStatusScNotify) Descriptor() ([]byte, []int) {
	return file_SetRedPointStatusScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SetRedPointStatusScNotify) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SetRedPointStatusScNotify) GetJGMIDHMLNFK() uint32 {
	if x != nil {
		return x.JGMIDHMLNFK
	}
	return 0
}

func (x *SetRedPointStatusScNotify) GetDFGNMBNLMJI() uint32 {
	if x != nil {
		return x.DFGNMBNLMJI
	}
	return 0
}

func (x *SetRedPointStatusScNotify) GetNBFOOFHBOPC() []*ALLNGMEDPMD {
	if x != nil {
		return x.NBFOOFHBOPC
	}
	return nil
}

func (x *SetRedPointStatusScNotify) GetContentId() uint32 {
	if x != nil {
		return x.ContentId
	}
	return 0
}

var File_SetRedPointStatusScNotify_proto protoreflect.FileDescriptor

var file_SetRedPointStatusScNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x53, 0x65, 0x74, 0x52, 0x65, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x41, 0x4c, 0x4c, 0x4e, 0x47, 0x4d, 0x45, 0x44, 0x50, 0x4d, 0x44, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x19, 0x53, 0x65, 0x74, 0x52, 0x65, 0x64, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x47, 0x4d, 0x49, 0x44, 0x48, 0x4d, 0x4c,
	0x4e, 0x46, 0x4b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x47, 0x4d, 0x49, 0x44,
	0x48, 0x4d, 0x4c, 0x4e, 0x46, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x46, 0x47, 0x4e, 0x4d, 0x42,
	0x4e, 0x4c, 0x4d, 0x4a, 0x49, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x46, 0x47,
	0x4e, 0x4d, 0x42, 0x4e, 0x4c, 0x4d, 0x4a, 0x49, 0x12, 0x2e, 0x0a, 0x0b, 0x4e, 0x42, 0x46, 0x4f,
	0x4f, 0x46, 0x48, 0x42, 0x4f, 0x50, 0x43, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x41, 0x4c, 0x4c, 0x4e, 0x47, 0x4d, 0x45, 0x44, 0x50, 0x4d, 0x44, 0x52, 0x0b, 0x4e, 0x42, 0x46,
	0x4f, 0x4f, 0x46, 0x48, 0x42, 0x4f, 0x50, 0x43, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetRedPointStatusScNotify_proto_rawDescOnce sync.Once
	file_SetRedPointStatusScNotify_proto_rawDescData = file_SetRedPointStatusScNotify_proto_rawDesc
)

func file_SetRedPointStatusScNotify_proto_rawDescGZIP() []byte {
	file_SetRedPointStatusScNotify_proto_rawDescOnce.Do(func() {
		file_SetRedPointStatusScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetRedPointStatusScNotify_proto_rawDescData)
	})
	return file_SetRedPointStatusScNotify_proto_rawDescData
}

var file_SetRedPointStatusScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetRedPointStatusScNotify_proto_goTypes = []interface{}{
	(*SetRedPointStatusScNotify)(nil), // 0: SetRedPointStatusScNotify
	(*ALLNGMEDPMD)(nil),               // 1: ALLNGMEDPMD
}
var file_SetRedPointStatusScNotify_proto_depIdxs = []int32{
	1, // 0: SetRedPointStatusScNotify.NBFOOFHBOPC:type_name -> ALLNGMEDPMD
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SetRedPointStatusScNotify_proto_init() }
func file_SetRedPointStatusScNotify_proto_init() {
	if File_SetRedPointStatusScNotify_proto != nil {
		return
	}
	file_ALLNGMEDPMD_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SetRedPointStatusScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRedPointStatusScNotify); i {
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
			RawDescriptor: file_SetRedPointStatusScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetRedPointStatusScNotify_proto_goTypes,
		DependencyIndexes: file_SetRedPointStatusScNotify_proto_depIdxs,
		MessageInfos:      file_SetRedPointStatusScNotify_proto_msgTypes,
	}.Build()
	File_SetRedPointStatusScNotify_proto = out.File
	file_SetRedPointStatusScNotify_proto_rawDesc = nil
	file_SetRedPointStatusScNotify_proto_goTypes = nil
	file_SetRedPointStatusScNotify_proto_depIdxs = nil
}
