// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CMPBKIENJNF.proto

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

type CMPBKIENJNF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to GBODJCDLFMO:
	//
	//	*CMPBKIENJNF_HPPKDAHGDKE
	//	*CMPBKIENJNF_OJFHPMBGPCC
	//	*CMPBKIENJNF_GGCIMCAFHKL
	//	*CMPBKIENJNF_NGLNGALEBJD
	//	*CMPBKIENJNF_MJGIKHCEBOG
	GBODJCDLFMO isCMPBKIENJNF_GBODJCDLFMO `protobuf_oneof:"GBODJCDLFMO"`
}

func (x *CMPBKIENJNF) Reset() {
	*x = CMPBKIENJNF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CMPBKIENJNF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMPBKIENJNF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMPBKIENJNF) ProtoMessage() {}

func (x *CMPBKIENJNF) ProtoReflect() protoreflect.Message {
	mi := &file_CMPBKIENJNF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMPBKIENJNF.ProtoReflect.Descriptor instead.
func (*CMPBKIENJNF) Descriptor() ([]byte, []int) {
	return file_CMPBKIENJNF_proto_rawDescGZIP(), []int{0}
}

func (m *CMPBKIENJNF) GetGBODJCDLFMO() isCMPBKIENJNF_GBODJCDLFMO {
	if m != nil {
		return m.GBODJCDLFMO
	}
	return nil
}

func (x *CMPBKIENJNF) GetHPPKDAHGDKE() *GJJIBMLPBPP {
	if x, ok := x.GetGBODJCDLFMO().(*CMPBKIENJNF_HPPKDAHGDKE); ok {
		return x.HPPKDAHGDKE
	}
	return nil
}

func (x *CMPBKIENJNF) GetOJFHPMBGPCC() *KFKAFAJGJFO {
	if x, ok := x.GetGBODJCDLFMO().(*CMPBKIENJNF_OJFHPMBGPCC); ok {
		return x.OJFHPMBGPCC
	}
	return nil
}

func (x *CMPBKIENJNF) GetGGCIMCAFHKL() *JNNOFOOFGJC {
	if x, ok := x.GetGBODJCDLFMO().(*CMPBKIENJNF_GGCIMCAFHKL); ok {
		return x.GGCIMCAFHKL
	}
	return nil
}

func (x *CMPBKIENJNF) GetNGLNGALEBJD() *MGKEPIFFDKK {
	if x, ok := x.GetGBODJCDLFMO().(*CMPBKIENJNF_NGLNGALEBJD); ok {
		return x.NGLNGALEBJD
	}
	return nil
}

func (x *CMPBKIENJNF) GetMJGIKHCEBOG() *JMFCLCEJIJA {
	if x, ok := x.GetGBODJCDLFMO().(*CMPBKIENJNF_MJGIKHCEBOG); ok {
		return x.MJGIKHCEBOG
	}
	return nil
}

type isCMPBKIENJNF_GBODJCDLFMO interface {
	isCMPBKIENJNF_GBODJCDLFMO()
}

type CMPBKIENJNF_HPPKDAHGDKE struct {
	HPPKDAHGDKE *GJJIBMLPBPP `protobuf:"bytes,2,opt,name=HPPKDAHGDKE,proto3,oneof"`
}

type CMPBKIENJNF_OJFHPMBGPCC struct {
	OJFHPMBGPCC *KFKAFAJGJFO `protobuf:"bytes,4,opt,name=OJFHPMBGPCC,proto3,oneof"`
}

type CMPBKIENJNF_GGCIMCAFHKL struct {
	GGCIMCAFHKL *JNNOFOOFGJC `protobuf:"bytes,11,opt,name=GGCIMCAFHKL,proto3,oneof"`
}

type CMPBKIENJNF_NGLNGALEBJD struct {
	NGLNGALEBJD *MGKEPIFFDKK `protobuf:"bytes,12,opt,name=NGLNGALEBJD,proto3,oneof"`
}

type CMPBKIENJNF_MJGIKHCEBOG struct {
	MJGIKHCEBOG *JMFCLCEJIJA `protobuf:"bytes,15,opt,name=MJGIKHCEBOG,proto3,oneof"`
}

func (*CMPBKIENJNF_HPPKDAHGDKE) isCMPBKIENJNF_GBODJCDLFMO() {}

func (*CMPBKIENJNF_OJFHPMBGPCC) isCMPBKIENJNF_GBODJCDLFMO() {}

func (*CMPBKIENJNF_GGCIMCAFHKL) isCMPBKIENJNF_GBODJCDLFMO() {}

func (*CMPBKIENJNF_NGLNGALEBJD) isCMPBKIENJNF_GBODJCDLFMO() {}

func (*CMPBKIENJNF_MJGIKHCEBOG) isCMPBKIENJNF_GBODJCDLFMO() {}

var File_CMPBKIENJNF_proto protoreflect.FileDescriptor

var file_CMPBKIENJNF_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x4d, 0x50, 0x42, 0x4b, 0x49, 0x45, 0x4e, 0x4a, 0x4e, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x4d, 0x46, 0x43, 0x4c, 0x43, 0x45, 0x4a, 0x49, 0x4a, 0x41,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x47, 0x4b, 0x45, 0x50, 0x49, 0x46, 0x46,
	0x44, 0x4b, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x4e, 0x4e, 0x4f, 0x46,
	0x4f, 0x4f, 0x46, 0x47, 0x4a, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x4a,
	0x4a, 0x49, 0x42, 0x4d, 0x4c, 0x50, 0x42, 0x50, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4b, 0x46, 0x4b, 0x41, 0x46, 0x41, 0x4a, 0x47, 0x4a, 0x46, 0x4f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x0b, 0x43, 0x4d, 0x50, 0x42, 0x4b, 0x49, 0x45, 0x4e, 0x4a,
	0x4e, 0x46, 0x12, 0x30, 0x0a, 0x0b, 0x48, 0x50, 0x50, 0x4b, 0x44, 0x41, 0x48, 0x47, 0x44, 0x4b,
	0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x4a, 0x4a, 0x49, 0x42, 0x4d,
	0x4c, 0x50, 0x42, 0x50, 0x50, 0x48, 0x00, 0x52, 0x0b, 0x48, 0x50, 0x50, 0x4b, 0x44, 0x41, 0x48,
	0x47, 0x44, 0x4b, 0x45, 0x12, 0x30, 0x0a, 0x0b, 0x4f, 0x4a, 0x46, 0x48, 0x50, 0x4d, 0x42, 0x47,
	0x50, 0x43, 0x43, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x46, 0x4b, 0x41,
	0x46, 0x41, 0x4a, 0x47, 0x4a, 0x46, 0x4f, 0x48, 0x00, 0x52, 0x0b, 0x4f, 0x4a, 0x46, 0x48, 0x50,
	0x4d, 0x42, 0x47, 0x50, 0x43, 0x43, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x47, 0x43, 0x49, 0x4d, 0x43,
	0x41, 0x46, 0x48, 0x4b, 0x4c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x4e,
	0x4e, 0x4f, 0x46, 0x4f, 0x4f, 0x46, 0x47, 0x4a, 0x43, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x47, 0x43,
	0x49, 0x4d, 0x43, 0x41, 0x46, 0x48, 0x4b, 0x4c, 0x12, 0x30, 0x0a, 0x0b, 0x4e, 0x47, 0x4c, 0x4e,
	0x47, 0x41, 0x4c, 0x45, 0x42, 0x4a, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4d, 0x47, 0x4b, 0x45, 0x50, 0x49, 0x46, 0x46, 0x44, 0x4b, 0x4b, 0x48, 0x00, 0x52, 0x0b, 0x4e,
	0x47, 0x4c, 0x4e, 0x47, 0x41, 0x4c, 0x45, 0x42, 0x4a, 0x44, 0x12, 0x30, 0x0a, 0x0b, 0x4d, 0x4a,
	0x47, 0x49, 0x4b, 0x48, 0x43, 0x45, 0x42, 0x4f, 0x47, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4a, 0x4d, 0x46, 0x43, 0x4c, 0x43, 0x45, 0x4a, 0x49, 0x4a, 0x41, 0x48, 0x00, 0x52,
	0x0b, 0x4d, 0x4a, 0x47, 0x49, 0x4b, 0x48, 0x43, 0x45, 0x42, 0x4f, 0x47, 0x42, 0x0d, 0x0a, 0x0b,
	0x47, 0x42, 0x4f, 0x44, 0x4a, 0x43, 0x44, 0x4c, 0x46, 0x4d, 0x4f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_CMPBKIENJNF_proto_rawDescOnce sync.Once
	file_CMPBKIENJNF_proto_rawDescData = file_CMPBKIENJNF_proto_rawDesc
)

func file_CMPBKIENJNF_proto_rawDescGZIP() []byte {
	file_CMPBKIENJNF_proto_rawDescOnce.Do(func() {
		file_CMPBKIENJNF_proto_rawDescData = protoimpl.X.CompressGZIP(file_CMPBKIENJNF_proto_rawDescData)
	})
	return file_CMPBKIENJNF_proto_rawDescData
}

var file_CMPBKIENJNF_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CMPBKIENJNF_proto_goTypes = []interface{}{
	(*CMPBKIENJNF)(nil), // 0: CMPBKIENJNF
	(*GJJIBMLPBPP)(nil), // 1: GJJIBMLPBPP
	(*KFKAFAJGJFO)(nil), // 2: KFKAFAJGJFO
	(*JNNOFOOFGJC)(nil), // 3: JNNOFOOFGJC
	(*MGKEPIFFDKK)(nil), // 4: MGKEPIFFDKK
	(*JMFCLCEJIJA)(nil), // 5: JMFCLCEJIJA
}
var file_CMPBKIENJNF_proto_depIdxs = []int32{
	1, // 0: CMPBKIENJNF.HPPKDAHGDKE:type_name -> GJJIBMLPBPP
	2, // 1: CMPBKIENJNF.OJFHPMBGPCC:type_name -> KFKAFAJGJFO
	3, // 2: CMPBKIENJNF.GGCIMCAFHKL:type_name -> JNNOFOOFGJC
	4, // 3: CMPBKIENJNF.NGLNGALEBJD:type_name -> MGKEPIFFDKK
	5, // 4: CMPBKIENJNF.MJGIKHCEBOG:type_name -> JMFCLCEJIJA
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_CMPBKIENJNF_proto_init() }
func file_CMPBKIENJNF_proto_init() {
	if File_CMPBKIENJNF_proto != nil {
		return
	}
	file_JMFCLCEJIJA_proto_init()
	file_MGKEPIFFDKK_proto_init()
	file_JNNOFOOFGJC_proto_init()
	file_GJJIBMLPBPP_proto_init()
	file_KFKAFAJGJFO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CMPBKIENJNF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CMPBKIENJNF); i {
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
	file_CMPBKIENJNF_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CMPBKIENJNF_HPPKDAHGDKE)(nil),
		(*CMPBKIENJNF_OJFHPMBGPCC)(nil),
		(*CMPBKIENJNF_GGCIMCAFHKL)(nil),
		(*CMPBKIENJNF_NGLNGALEBJD)(nil),
		(*CMPBKIENJNF_MJGIKHCEBOG)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CMPBKIENJNF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CMPBKIENJNF_proto_goTypes,
		DependencyIndexes: file_CMPBKIENJNF_proto_depIdxs,
		MessageInfos:      file_CMPBKIENJNF_proto_msgTypes,
	}.Build()
	File_CMPBKIENJNF_proto = out.File
	file_CMPBKIENJNF_proto_rawDesc = nil
	file_CMPBKIENJNF_proto_goTypes = nil
	file_CMPBKIENJNF_proto_depIdxs = nil
}
