// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: OHAJNAEBFCO.proto

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

type OHAJNAEBFCO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NNJGPNFNENA []uint32 `protobuf:"varint,3,rep,packed,name=NNJGPNFNENA,proto3" json:"NNJGPNFNENA,omitempty"`
	DEKHELNPCAM bool     `protobuf:"varint,6,opt,name=DEKHELNPCAM,proto3" json:"DEKHELNPCAM,omitempty"`
	KIBEEKLHMBI uint32   `protobuf:"varint,15,opt,name=KIBEEKLHMBI,proto3" json:"KIBEEKLHMBI,omitempty"`
	AEEFHDJMEGD []uint32 `protobuf:"varint,11,rep,packed,name=AEEFHDJMEGD,proto3" json:"AEEFHDJMEGD,omitempty"`
	NEOIBJKPKKF uint32   `protobuf:"varint,8,opt,name=NEOIBJKPKKF,proto3" json:"NEOIBJKPKKF,omitempty"`
	HICGMJNEMDL uint32   `protobuf:"varint,2,opt,name=HICGMJNEMDL,proto3" json:"HICGMJNEMDL,omitempty"`
}

func (x *OHAJNAEBFCO) Reset() {
	*x = OHAJNAEBFCO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OHAJNAEBFCO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OHAJNAEBFCO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OHAJNAEBFCO) ProtoMessage() {}

func (x *OHAJNAEBFCO) ProtoReflect() protoreflect.Message {
	mi := &file_OHAJNAEBFCO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OHAJNAEBFCO.ProtoReflect.Descriptor instead.
func (*OHAJNAEBFCO) Descriptor() ([]byte, []int) {
	return file_OHAJNAEBFCO_proto_rawDescGZIP(), []int{0}
}

func (x *OHAJNAEBFCO) GetNNJGPNFNENA() []uint32 {
	if x != nil {
		return x.NNJGPNFNENA
	}
	return nil
}

func (x *OHAJNAEBFCO) GetDEKHELNPCAM() bool {
	if x != nil {
		return x.DEKHELNPCAM
	}
	return false
}

func (x *OHAJNAEBFCO) GetKIBEEKLHMBI() uint32 {
	if x != nil {
		return x.KIBEEKLHMBI
	}
	return 0
}

func (x *OHAJNAEBFCO) GetAEEFHDJMEGD() []uint32 {
	if x != nil {
		return x.AEEFHDJMEGD
	}
	return nil
}

func (x *OHAJNAEBFCO) GetNEOIBJKPKKF() uint32 {
	if x != nil {
		return x.NEOIBJKPKKF
	}
	return 0
}

func (x *OHAJNAEBFCO) GetHICGMJNEMDL() uint32 {
	if x != nil {
		return x.HICGMJNEMDL
	}
	return 0
}

var File_OHAJNAEBFCO_proto protoreflect.FileDescriptor

var file_OHAJNAEBFCO_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4f, 0x48, 0x41, 0x4a, 0x4e, 0x41, 0x45, 0x42, 0x46, 0x43, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x0b, 0x4f, 0x48, 0x41, 0x4a, 0x4e, 0x41, 0x45, 0x42,
	0x46, 0x43, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4e, 0x4a, 0x47, 0x50, 0x4e, 0x46, 0x4e, 0x45,
	0x4e, 0x41, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x4e, 0x4a, 0x47, 0x50, 0x4e,
	0x46, 0x4e, 0x45, 0x4e, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x45, 0x4b, 0x48, 0x45, 0x4c, 0x4e,
	0x50, 0x43, 0x41, 0x4d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x45, 0x4b, 0x48,
	0x45, 0x4c, 0x4e, 0x50, 0x43, 0x41, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x49, 0x42, 0x45, 0x45,
	0x4b, 0x4c, 0x48, 0x4d, 0x42, 0x49, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x49,
	0x42, 0x45, 0x45, 0x4b, 0x4c, 0x48, 0x4d, 0x42, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x45, 0x45,
	0x46, 0x48, 0x44, 0x4a, 0x4d, 0x45, 0x47, 0x44, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b,
	0x41, 0x45, 0x45, 0x46, 0x48, 0x44, 0x4a, 0x4d, 0x45, 0x47, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4e,
	0x45, 0x4f, 0x49, 0x42, 0x4a, 0x4b, 0x50, 0x4b, 0x4b, 0x46, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4e, 0x45, 0x4f, 0x49, 0x42, 0x4a, 0x4b, 0x50, 0x4b, 0x4b, 0x46, 0x12, 0x20, 0x0a,
	0x0b, 0x48, 0x49, 0x43, 0x47, 0x4d, 0x4a, 0x4e, 0x45, 0x4d, 0x44, 0x4c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x48, 0x49, 0x43, 0x47, 0x4d, 0x4a, 0x4e, 0x45, 0x4d, 0x44, 0x4c, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OHAJNAEBFCO_proto_rawDescOnce sync.Once
	file_OHAJNAEBFCO_proto_rawDescData = file_OHAJNAEBFCO_proto_rawDesc
)

func file_OHAJNAEBFCO_proto_rawDescGZIP() []byte {
	file_OHAJNAEBFCO_proto_rawDescOnce.Do(func() {
		file_OHAJNAEBFCO_proto_rawDescData = protoimpl.X.CompressGZIP(file_OHAJNAEBFCO_proto_rawDescData)
	})
	return file_OHAJNAEBFCO_proto_rawDescData
}

var file_OHAJNAEBFCO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OHAJNAEBFCO_proto_goTypes = []interface{}{
	(*OHAJNAEBFCO)(nil), // 0: OHAJNAEBFCO
}
var file_OHAJNAEBFCO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_OHAJNAEBFCO_proto_init() }
func file_OHAJNAEBFCO_proto_init() {
	if File_OHAJNAEBFCO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_OHAJNAEBFCO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OHAJNAEBFCO); i {
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
			RawDescriptor: file_OHAJNAEBFCO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OHAJNAEBFCO_proto_goTypes,
		DependencyIndexes: file_OHAJNAEBFCO_proto_depIdxs,
		MessageInfos:      file_OHAJNAEBFCO_proto_msgTypes,
	}.Build()
	File_OHAJNAEBFCO_proto = out.File
	file_OHAJNAEBFCO_proto_rawDesc = nil
	file_OHAJNAEBFCO_proto_goTypes = nil
	file_OHAJNAEBFCO_proto_depIdxs = nil
}
