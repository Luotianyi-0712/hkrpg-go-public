// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MonopolyGetDailyInitItemScRsp.proto

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

type MonopolyGetDailyInitItemScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LPOKPLLNHPE uint32 `protobuf:"varint,12,opt,name=LPOKPLLNHPE,proto3" json:"LPOKPLLNHPE,omitempty"`
	HCIBLLAPPMG uint32 `protobuf:"varint,1,opt,name=HCIBLLAPPMG,proto3" json:"HCIBLLAPPMG,omitempty"`
	PLABPPOECFN int64  `protobuf:"varint,11,opt,name=PLABPPOECFN,proto3" json:"PLABPPOECFN,omitempty"`
	Retcode     uint32 `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	KBPCANEJHKO uint32 `protobuf:"varint,8,opt,name=KBPCANEJHKO,proto3" json:"KBPCANEJHKO,omitempty"`
	AFIMBFONAHC uint32 `protobuf:"varint,6,opt,name=AFIMBFONAHC,proto3" json:"AFIMBFONAHC,omitempty"`
	HAHCADMIMHJ uint32 `protobuf:"varint,14,opt,name=HAHCADMIMHJ,proto3" json:"HAHCADMIMHJ,omitempty"`
}

func (x *MonopolyGetDailyInitItemScRsp) Reset() {
	*x = MonopolyGetDailyInitItemScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MonopolyGetDailyInitItemScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonopolyGetDailyInitItemScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonopolyGetDailyInitItemScRsp) ProtoMessage() {}

func (x *MonopolyGetDailyInitItemScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_MonopolyGetDailyInitItemScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonopolyGetDailyInitItemScRsp.ProtoReflect.Descriptor instead.
func (*MonopolyGetDailyInitItemScRsp) Descriptor() ([]byte, []int) {
	return file_MonopolyGetDailyInitItemScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *MonopolyGetDailyInitItemScRsp) GetLPOKPLLNHPE() uint32 {
	if x != nil {
		return x.LPOKPLLNHPE
	}
	return 0
}

func (x *MonopolyGetDailyInitItemScRsp) GetHCIBLLAPPMG() uint32 {
	if x != nil {
		return x.HCIBLLAPPMG
	}
	return 0
}

func (x *MonopolyGetDailyInitItemScRsp) GetPLABPPOECFN() int64 {
	if x != nil {
		return x.PLABPPOECFN
	}
	return 0
}

func (x *MonopolyGetDailyInitItemScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *MonopolyGetDailyInitItemScRsp) GetKBPCANEJHKO() uint32 {
	if x != nil {
		return x.KBPCANEJHKO
	}
	return 0
}

func (x *MonopolyGetDailyInitItemScRsp) GetAFIMBFONAHC() uint32 {
	if x != nil {
		return x.AFIMBFONAHC
	}
	return 0
}

func (x *MonopolyGetDailyInitItemScRsp) GetHAHCADMIMHJ() uint32 {
	if x != nil {
		return x.HAHCADMIMHJ
	}
	return 0
}

var File_MonopolyGetDailyInitItemScRsp_proto protoreflect.FileDescriptor

var file_MonopolyGetDailyInitItemScRsp_proto_rawDesc = []byte{
	0x0a, 0x23, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x1d, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f,
	0x6c, 0x79, 0x47, 0x65, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x50, 0x4f, 0x4b, 0x50,
	0x4c, 0x4c, 0x4e, 0x48, 0x50, 0x45, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x50,
	0x4f, 0x4b, 0x50, 0x4c, 0x4c, 0x4e, 0x48, 0x50, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x43, 0x49,
	0x42, 0x4c, 0x4c, 0x41, 0x50, 0x50, 0x4d, 0x47, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x48, 0x43, 0x49, 0x42, 0x4c, 0x4c, 0x41, 0x50, 0x50, 0x4d, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x4c, 0x41, 0x42, 0x50, 0x50, 0x4f, 0x45, 0x43, 0x46, 0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x50, 0x4c, 0x41, 0x42, 0x50, 0x50, 0x4f, 0x45, 0x43, 0x46, 0x4e, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x42, 0x50, 0x43, 0x41,
	0x4e, 0x45, 0x4a, 0x48, 0x4b, 0x4f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x42,
	0x50, 0x43, 0x41, 0x4e, 0x45, 0x4a, 0x48, 0x4b, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x46, 0x49,
	0x4d, 0x42, 0x46, 0x4f, 0x4e, 0x41, 0x48, 0x43, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x41, 0x46, 0x49, 0x4d, 0x42, 0x46, 0x4f, 0x4e, 0x41, 0x48, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x48,
	0x41, 0x48, 0x43, 0x41, 0x44, 0x4d, 0x49, 0x4d, 0x48, 0x4a, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x48, 0x41, 0x48, 0x43, 0x41, 0x44, 0x4d, 0x49, 0x4d, 0x48, 0x4a, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MonopolyGetDailyInitItemScRsp_proto_rawDescOnce sync.Once
	file_MonopolyGetDailyInitItemScRsp_proto_rawDescData = file_MonopolyGetDailyInitItemScRsp_proto_rawDesc
)

func file_MonopolyGetDailyInitItemScRsp_proto_rawDescGZIP() []byte {
	file_MonopolyGetDailyInitItemScRsp_proto_rawDescOnce.Do(func() {
		file_MonopolyGetDailyInitItemScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_MonopolyGetDailyInitItemScRsp_proto_rawDescData)
	})
	return file_MonopolyGetDailyInitItemScRsp_proto_rawDescData
}

var file_MonopolyGetDailyInitItemScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MonopolyGetDailyInitItemScRsp_proto_goTypes = []interface{}{
	(*MonopolyGetDailyInitItemScRsp)(nil), // 0: MonopolyGetDailyInitItemScRsp
}
var file_MonopolyGetDailyInitItemScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MonopolyGetDailyInitItemScRsp_proto_init() }
func file_MonopolyGetDailyInitItemScRsp_proto_init() {
	if File_MonopolyGetDailyInitItemScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MonopolyGetDailyInitItemScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonopolyGetDailyInitItemScRsp); i {
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
			RawDescriptor: file_MonopolyGetDailyInitItemScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MonopolyGetDailyInitItemScRsp_proto_goTypes,
		DependencyIndexes: file_MonopolyGetDailyInitItemScRsp_proto_depIdxs,
		MessageInfos:      file_MonopolyGetDailyInitItemScRsp_proto_msgTypes,
	}.Build()
	File_MonopolyGetDailyInitItemScRsp_proto = out.File
	file_MonopolyGetDailyInitItemScRsp_proto_rawDesc = nil
	file_MonopolyGetDailyInitItemScRsp_proto_goTypes = nil
	file_MonopolyGetDailyInitItemScRsp_proto_depIdxs = nil
}
