// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CancelCacheNotifyCsReq.proto

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

type CancelCacheNotifyCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ADENLFJOOPL []uint32    `protobuf:"varint,9,rep,packed,name=ADENLFJOOPL,proto3" json:"ADENLFJOOPL,omitempty"`
	LDELHKCCHHN []string    `protobuf:"bytes,13,rep,name=LDELHKCCHHN,proto3" json:"LDELHKCCHHN,omitempty"`
	Type        HDCNOKOKDHC `protobuf:"varint,2,opt,name=type,proto3,enum=HDCNOKOKDHC" json:"type,omitempty"`
}

func (x *CancelCacheNotifyCsReq) Reset() {
	*x = CancelCacheNotifyCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CancelCacheNotifyCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelCacheNotifyCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelCacheNotifyCsReq) ProtoMessage() {}

func (x *CancelCacheNotifyCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_CancelCacheNotifyCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelCacheNotifyCsReq.ProtoReflect.Descriptor instead.
func (*CancelCacheNotifyCsReq) Descriptor() ([]byte, []int) {
	return file_CancelCacheNotifyCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *CancelCacheNotifyCsReq) GetADENLFJOOPL() []uint32 {
	if x != nil {
		return x.ADENLFJOOPL
	}
	return nil
}

func (x *CancelCacheNotifyCsReq) GetLDELHKCCHHN() []string {
	if x != nil {
		return x.LDELHKCCHHN
	}
	return nil
}

func (x *CancelCacheNotifyCsReq) GetType() HDCNOKOKDHC {
	if x != nil {
		return x.Type
	}
	return HDCNOKOKDHC_CACHE_NOTIFY_TYPE_NONE
}

var File_CancelCacheNotifyCsReq_proto protoreflect.FileDescriptor

var file_CancelCacheNotifyCsReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x48, 0x44, 0x43, 0x4e, 0x4f, 0x4b, 0x4f, 0x4b, 0x44, 0x48, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7e, 0x0a, 0x16, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x41,
	0x44, 0x45, 0x4e, 0x4c, 0x46, 0x4a, 0x4f, 0x4f, 0x50, 0x4c, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0b, 0x41, 0x44, 0x45, 0x4e, 0x4c, 0x46, 0x4a, 0x4f, 0x4f, 0x50, 0x4c, 0x12, 0x20, 0x0a,
	0x0b, 0x4c, 0x44, 0x45, 0x4c, 0x48, 0x4b, 0x43, 0x43, 0x48, 0x48, 0x4e, 0x18, 0x0d, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x4c, 0x44, 0x45, 0x4c, 0x48, 0x4b, 0x43, 0x43, 0x48, 0x48, 0x4e, 0x12,
	0x20, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e,
	0x48, 0x44, 0x43, 0x4e, 0x4f, 0x4b, 0x4f, 0x4b, 0x44, 0x48, 0x43, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CancelCacheNotifyCsReq_proto_rawDescOnce sync.Once
	file_CancelCacheNotifyCsReq_proto_rawDescData = file_CancelCacheNotifyCsReq_proto_rawDesc
)

func file_CancelCacheNotifyCsReq_proto_rawDescGZIP() []byte {
	file_CancelCacheNotifyCsReq_proto_rawDescOnce.Do(func() {
		file_CancelCacheNotifyCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_CancelCacheNotifyCsReq_proto_rawDescData)
	})
	return file_CancelCacheNotifyCsReq_proto_rawDescData
}

var file_CancelCacheNotifyCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CancelCacheNotifyCsReq_proto_goTypes = []interface{}{
	(*CancelCacheNotifyCsReq)(nil), // 0: CancelCacheNotifyCsReq
	(HDCNOKOKDHC)(0),               // 1: HDCNOKOKDHC
}
var file_CancelCacheNotifyCsReq_proto_depIdxs = []int32{
	1, // 0: CancelCacheNotifyCsReq.type:type_name -> HDCNOKOKDHC
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CancelCacheNotifyCsReq_proto_init() }
func file_CancelCacheNotifyCsReq_proto_init() {
	if File_CancelCacheNotifyCsReq_proto != nil {
		return
	}
	file_HDCNOKOKDHC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CancelCacheNotifyCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelCacheNotifyCsReq); i {
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
			RawDescriptor: file_CancelCacheNotifyCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CancelCacheNotifyCsReq_proto_goTypes,
		DependencyIndexes: file_CancelCacheNotifyCsReq_proto_depIdxs,
		MessageInfos:      file_CancelCacheNotifyCsReq_proto_msgTypes,
	}.Build()
	File_CancelCacheNotifyCsReq_proto = out.File
	file_CancelCacheNotifyCsReq_proto_rawDesc = nil
	file_CancelCacheNotifyCsReq_proto_goTypes = nil
	file_CancelCacheNotifyCsReq_proto_depIdxs = nil
}
