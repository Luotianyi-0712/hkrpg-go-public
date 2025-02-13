// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetAuthkeyScRsp.proto

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

type GetAuthkeyScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32 `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
	FPLJABGPLBA string `protobuf:"bytes,3,opt,name=FPLJABGPLBA,proto3" json:"FPLJABGPLBA,omitempty"`
	MMGCPODCFLP string `protobuf:"bytes,15,opt,name=MMGCPODCFLP,proto3" json:"MMGCPODCFLP,omitempty"`
	LIKAEKCNMHN uint32 `protobuf:"varint,4,opt,name=LIKAEKCNMHN,proto3" json:"LIKAEKCNMHN,omitempty"`
	ADIOHJGIBLP uint32 `protobuf:"varint,6,opt,name=ADIOHJGIBLP,proto3" json:"ADIOHJGIBLP,omitempty"`
}

func (x *GetAuthkeyScRsp) Reset() {
	*x = GetAuthkeyScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetAuthkeyScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthkeyScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthkeyScRsp) ProtoMessage() {}

func (x *GetAuthkeyScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetAuthkeyScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthkeyScRsp.ProtoReflect.Descriptor instead.
func (*GetAuthkeyScRsp) Descriptor() ([]byte, []int) {
	return file_GetAuthkeyScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetAuthkeyScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetAuthkeyScRsp) GetFPLJABGPLBA() string {
	if x != nil {
		return x.FPLJABGPLBA
	}
	return ""
}

func (x *GetAuthkeyScRsp) GetMMGCPODCFLP() string {
	if x != nil {
		return x.MMGCPODCFLP
	}
	return ""
}

func (x *GetAuthkeyScRsp) GetLIKAEKCNMHN() uint32 {
	if x != nil {
		return x.LIKAEKCNMHN
	}
	return 0
}

func (x *GetAuthkeyScRsp) GetADIOHJGIBLP() uint32 {
	if x != nil {
		return x.ADIOHJGIBLP
	}
	return 0
}

var File_GetAuthkeyScRsp_proto protoreflect.FileDescriptor

var file_GetAuthkeyScRsp_proto_rawDesc = []byte{
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x53, 0x63, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x50, 0x4c, 0x4a, 0x41, 0x42, 0x47,
	0x50, 0x4c, 0x42, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x50, 0x4c, 0x4a,
	0x41, 0x42, 0x47, 0x50, 0x4c, 0x42, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4d, 0x47, 0x43, 0x50,
	0x4f, 0x44, 0x43, 0x46, 0x4c, 0x50, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x4d,
	0x47, 0x43, 0x50, 0x4f, 0x44, 0x43, 0x46, 0x4c, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x49, 0x4b,
	0x41, 0x45, 0x4b, 0x43, 0x4e, 0x4d, 0x48, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4c, 0x49, 0x4b, 0x41, 0x45, 0x4b, 0x43, 0x4e, 0x4d, 0x48, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x41,
	0x44, 0x49, 0x4f, 0x48, 0x4a, 0x47, 0x49, 0x42, 0x4c, 0x50, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x41, 0x44, 0x49, 0x4f, 0x48, 0x4a, 0x47, 0x49, 0x42, 0x4c, 0x50, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetAuthkeyScRsp_proto_rawDescOnce sync.Once
	file_GetAuthkeyScRsp_proto_rawDescData = file_GetAuthkeyScRsp_proto_rawDesc
)

func file_GetAuthkeyScRsp_proto_rawDescGZIP() []byte {
	file_GetAuthkeyScRsp_proto_rawDescOnce.Do(func() {
		file_GetAuthkeyScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetAuthkeyScRsp_proto_rawDescData)
	})
	return file_GetAuthkeyScRsp_proto_rawDescData
}

var file_GetAuthkeyScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetAuthkeyScRsp_proto_goTypes = []interface{}{
	(*GetAuthkeyScRsp)(nil), // 0: GetAuthkeyScRsp
}
var file_GetAuthkeyScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetAuthkeyScRsp_proto_init() }
func file_GetAuthkeyScRsp_proto_init() {
	if File_GetAuthkeyScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetAuthkeyScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthkeyScRsp); i {
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
			RawDescriptor: file_GetAuthkeyScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetAuthkeyScRsp_proto_goTypes,
		DependencyIndexes: file_GetAuthkeyScRsp_proto_depIdxs,
		MessageInfos:      file_GetAuthkeyScRsp_proto_msgTypes,
	}.Build()
	File_GetAuthkeyScRsp_proto = out.File
	file_GetAuthkeyScRsp_proto_rawDesc = nil
	file_GetAuthkeyScRsp_proto_goTypes = nil
	file_GetAuthkeyScRsp_proto_depIdxs = nil
}
