// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FEMNHLAPBOG.proto

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

type FEMNHLAPBOG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pos         uint32 `protobuf:"varint,11,opt,name=pos,proto3" json:"pos,omitempty"`
	JIELNNCBKOD uint32 `protobuf:"varint,10,opt,name=JIELNNCBKOD,proto3" json:"JIELNNCBKOD,omitempty"`
	CPKLPJGGEOM uint32 `protobuf:"varint,9,opt,name=CPKLPJGGEOM,proto3" json:"CPKLPJGGEOM,omitempty"`
	Count       uint32 `protobuf:"varint,12,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *FEMNHLAPBOG) Reset() {
	*x = FEMNHLAPBOG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FEMNHLAPBOG_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FEMNHLAPBOG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FEMNHLAPBOG) ProtoMessage() {}

func (x *FEMNHLAPBOG) ProtoReflect() protoreflect.Message {
	mi := &file_FEMNHLAPBOG_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FEMNHLAPBOG.ProtoReflect.Descriptor instead.
func (*FEMNHLAPBOG) Descriptor() ([]byte, []int) {
	return file_FEMNHLAPBOG_proto_rawDescGZIP(), []int{0}
}

func (x *FEMNHLAPBOG) GetPos() uint32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

func (x *FEMNHLAPBOG) GetJIELNNCBKOD() uint32 {
	if x != nil {
		return x.JIELNNCBKOD
	}
	return 0
}

func (x *FEMNHLAPBOG) GetCPKLPJGGEOM() uint32 {
	if x != nil {
		return x.CPKLPJGGEOM
	}
	return 0
}

func (x *FEMNHLAPBOG) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_FEMNHLAPBOG_proto protoreflect.FileDescriptor

var file_FEMNHLAPBOG_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x45, 0x4d, 0x4e, 0x48, 0x4c, 0x41, 0x50, 0x42, 0x4f, 0x47, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0b, 0x46, 0x45, 0x4d, 0x4e, 0x48, 0x4c, 0x41, 0x50, 0x42,
	0x4f, 0x47, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x70, 0x6f, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x49, 0x45, 0x4c, 0x4e, 0x4e, 0x43, 0x42,
	0x4b, 0x4f, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x49, 0x45, 0x4c, 0x4e,
	0x4e, 0x43, 0x42, 0x4b, 0x4f, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x50, 0x4b, 0x4c, 0x50, 0x4a,
	0x47, 0x47, 0x45, 0x4f, 0x4d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x50, 0x4b,
	0x4c, 0x50, 0x4a, 0x47, 0x47, 0x45, 0x4f, 0x4d, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FEMNHLAPBOG_proto_rawDescOnce sync.Once
	file_FEMNHLAPBOG_proto_rawDescData = file_FEMNHLAPBOG_proto_rawDesc
)

func file_FEMNHLAPBOG_proto_rawDescGZIP() []byte {
	file_FEMNHLAPBOG_proto_rawDescOnce.Do(func() {
		file_FEMNHLAPBOG_proto_rawDescData = protoimpl.X.CompressGZIP(file_FEMNHLAPBOG_proto_rawDescData)
	})
	return file_FEMNHLAPBOG_proto_rawDescData
}

var file_FEMNHLAPBOG_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FEMNHLAPBOG_proto_goTypes = []interface{}{
	(*FEMNHLAPBOG)(nil), // 0: FEMNHLAPBOG
}
var file_FEMNHLAPBOG_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FEMNHLAPBOG_proto_init() }
func file_FEMNHLAPBOG_proto_init() {
	if File_FEMNHLAPBOG_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FEMNHLAPBOG_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FEMNHLAPBOG); i {
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
			RawDescriptor: file_FEMNHLAPBOG_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FEMNHLAPBOG_proto_goTypes,
		DependencyIndexes: file_FEMNHLAPBOG_proto_depIdxs,
		MessageInfos:      file_FEMNHLAPBOG_proto_msgTypes,
	}.Build()
	File_FEMNHLAPBOG_proto = out.File
	file_FEMNHLAPBOG_proto_rawDesc = nil
	file_FEMNHLAPBOG_proto_goTypes = nil
	file_FEMNHLAPBOG_proto_depIdxs = nil
}
