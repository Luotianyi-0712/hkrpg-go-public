// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ClockParkGetOngoingScriptInfoScRsp.proto

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

type ClockParkGetOngoingScriptInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MDBFIEDBPMO   uint32       `protobuf:"varint,1,opt,name=MDBFIEDBPMO,proto3" json:"MDBFIEDBPMO,omitempty"`
	ScriptId      uint32       `protobuf:"varint,14,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty"`
	PGBMJLNICPH   uint32       `protobuf:"varint,11,opt,name=PGBMJLNICPH,proto3" json:"PGBMJLNICPH,omitempty"`
	RogueBuffInfo *PKIPPPBNMLP `protobuf:"bytes,13,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
	KFKBNMFAJKJ   []uint32     `protobuf:"varint,7,rep,packed,name=KFKBNMFAJKJ,proto3" json:"KFKBNMFAJKJ,omitempty"`
	CHDEKEMNIEA   uint32       `protobuf:"varint,9,opt,name=CHDEKEMNIEA,proto3" json:"CHDEKEMNIEA,omitempty"`
	GJAFBHKJIOK   *MANNPANJCLL `protobuf:"bytes,6,opt,name=GJAFBHKJIOK,proto3" json:"GJAFBHKJIOK,omitempty"`
	IIHCNNGOLBE   uint32       `protobuf:"varint,5,opt,name=IIHCNNGOLBE,proto3" json:"IIHCNNGOLBE,omitempty"`
	BJDGPNJCAJC   string       `protobuf:"bytes,10,opt,name=BJDGPNJCAJC,proto3" json:"BJDGPNJCAJC,omitempty"`
	JEGHBJMOELL   uint32       `protobuf:"varint,12,opt,name=JEGHBJMOELL,proto3" json:"JEGHBJMOELL,omitempty"`
	JBBDIEMBGMG   *GFJBDNNGECB `protobuf:"bytes,2,opt,name=JBBDIEMBGMG,proto3" json:"JBBDIEMBGMG,omitempty"`
	Retcode       uint32       `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *ClockParkGetOngoingScriptInfoScRsp) Reset() {
	*x = ClockParkGetOngoingScriptInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClockParkGetOngoingScriptInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockParkGetOngoingScriptInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockParkGetOngoingScriptInfoScRsp) ProtoMessage() {}

func (x *ClockParkGetOngoingScriptInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ClockParkGetOngoingScriptInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockParkGetOngoingScriptInfoScRsp.ProtoReflect.Descriptor instead.
func (*ClockParkGetOngoingScriptInfoScRsp) Descriptor() ([]byte, []int) {
	return file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetMDBFIEDBPMO() uint32 {
	if x != nil {
		return x.MDBFIEDBPMO
	}
	return 0
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetScriptId() uint32 {
	if x != nil {
		return x.ScriptId
	}
	return 0
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetPGBMJLNICPH() uint32 {
	if x != nil {
		return x.PGBMJLNICPH
	}
	return 0
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetRogueBuffInfo() *PKIPPPBNMLP {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetKFKBNMFAJKJ() []uint32 {
	if x != nil {
		return x.KFKBNMFAJKJ
	}
	return nil
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetCHDEKEMNIEA() uint32 {
	if x != nil {
		return x.CHDEKEMNIEA
	}
	return 0
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetGJAFBHKJIOK() *MANNPANJCLL {
	if x != nil {
		return x.GJAFBHKJIOK
	}
	return nil
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetIIHCNNGOLBE() uint32 {
	if x != nil {
		return x.IIHCNNGOLBE
	}
	return 0
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetBJDGPNJCAJC() string {
	if x != nil {
		return x.BJDGPNJCAJC
	}
	return ""
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetJEGHBJMOELL() uint32 {
	if x != nil {
		return x.JEGHBJMOELL
	}
	return 0
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetJBBDIEMBGMG() *GFJBDNNGECB {
	if x != nil {
		return x.JBBDIEMBGMG
	}
	return nil
}

func (x *ClockParkGetOngoingScriptInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_ClockParkGetOngoingScriptInfoScRsp_proto protoreflect.FileDescriptor

var file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x28, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x67, 0x6f, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x46, 0x4a, 0x42,
	0x44, 0x4e, 0x4e, 0x47, 0x45, 0x43, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d,
	0x41, 0x4e, 0x4e, 0x50, 0x41, 0x4e, 0x4a, 0x43, 0x4c, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x50, 0x4b, 0x49, 0x50, 0x50, 0x50, 0x42, 0x4e, 0x4d, 0x4c, 0x50, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x03, 0x0a, 0x22, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72,
	0x6b, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x44,
	0x42, 0x46, 0x49, 0x45, 0x44, 0x42, 0x50, 0x4d, 0x4f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x4d, 0x44, 0x42, 0x46, 0x49, 0x45, 0x44, 0x42, 0x50, 0x4d, 0x4f, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x47, 0x42,
	0x4d, 0x4a, 0x4c, 0x4e, 0x49, 0x43, 0x50, 0x48, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x50, 0x47, 0x42, 0x4d, 0x4a, 0x4c, 0x4e, 0x49, 0x43, 0x50, 0x48, 0x12, 0x34, 0x0a, 0x0f, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x4b, 0x49, 0x50, 0x50, 0x50, 0x42, 0x4e, 0x4d,
	0x4c, 0x50, 0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x46, 0x4b, 0x42, 0x4e, 0x4d, 0x46, 0x41, 0x4a, 0x4b, 0x4a,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x46, 0x4b, 0x42, 0x4e, 0x4d, 0x46, 0x41,
	0x4a, 0x4b, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x48, 0x44, 0x45, 0x4b, 0x45, 0x4d, 0x4e, 0x49,
	0x45, 0x41, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x48, 0x44, 0x45, 0x4b, 0x45,
	0x4d, 0x4e, 0x49, 0x45, 0x41, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x4a, 0x41, 0x46, 0x42, 0x48, 0x4b,
	0x4a, 0x49, 0x4f, 0x4b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x41, 0x4e,
	0x4e, 0x50, 0x41, 0x4e, 0x4a, 0x43, 0x4c, 0x4c, 0x52, 0x0b, 0x47, 0x4a, 0x41, 0x46, 0x42, 0x48,
	0x4b, 0x4a, 0x49, 0x4f, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x49, 0x48, 0x43, 0x4e, 0x4e, 0x47,
	0x4f, 0x4c, 0x42, 0x45, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x49, 0x48, 0x43,
	0x4e, 0x4e, 0x47, 0x4f, 0x4c, 0x42, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x4a, 0x44, 0x47, 0x50,
	0x4e, 0x4a, 0x43, 0x41, 0x4a, 0x43, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x4a,
	0x44, 0x47, 0x50, 0x4e, 0x4a, 0x43, 0x41, 0x4a, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x45, 0x47,
	0x48, 0x42, 0x4a, 0x4d, 0x4f, 0x45, 0x4c, 0x4c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4a, 0x45, 0x47, 0x48, 0x42, 0x4a, 0x4d, 0x4f, 0x45, 0x4c, 0x4c, 0x12, 0x2e, 0x0a, 0x0b, 0x4a,
	0x42, 0x42, 0x44, 0x49, 0x45, 0x4d, 0x42, 0x47, 0x4d, 0x47, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x47, 0x46, 0x4a, 0x42, 0x44, 0x4e, 0x4e, 0x47, 0x45, 0x43, 0x42, 0x52, 0x0b,
	0x4a, 0x42, 0x42, 0x44, 0x49, 0x45, 0x4d, 0x42, 0x47, 0x4d, 0x47, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDescOnce sync.Once
	file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDescData = file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDesc
)

func file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDescGZIP() []byte {
	file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDescOnce.Do(func() {
		file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDescData)
	})
	return file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDescData
}

var file_ClockParkGetOngoingScriptInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClockParkGetOngoingScriptInfoScRsp_proto_goTypes = []interface{}{
	(*ClockParkGetOngoingScriptInfoScRsp)(nil), // 0: ClockParkGetOngoingScriptInfoScRsp
	(*PKIPPPBNMLP)(nil),                        // 1: PKIPPPBNMLP
	(*MANNPANJCLL)(nil),                        // 2: MANNPANJCLL
	(*GFJBDNNGECB)(nil),                        // 3: GFJBDNNGECB
}
var file_ClockParkGetOngoingScriptInfoScRsp_proto_depIdxs = []int32{
	1, // 0: ClockParkGetOngoingScriptInfoScRsp.rogue_buff_info:type_name -> PKIPPPBNMLP
	2, // 1: ClockParkGetOngoingScriptInfoScRsp.GJAFBHKJIOK:type_name -> MANNPANJCLL
	3, // 2: ClockParkGetOngoingScriptInfoScRsp.JBBDIEMBGMG:type_name -> GFJBDNNGECB
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ClockParkGetOngoingScriptInfoScRsp_proto_init() }
func file_ClockParkGetOngoingScriptInfoScRsp_proto_init() {
	if File_ClockParkGetOngoingScriptInfoScRsp_proto != nil {
		return
	}
	file_GFJBDNNGECB_proto_init()
	file_MANNPANJCLL_proto_init()
	file_PKIPPPBNMLP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ClockParkGetOngoingScriptInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockParkGetOngoingScriptInfoScRsp); i {
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
			RawDescriptor: file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClockParkGetOngoingScriptInfoScRsp_proto_goTypes,
		DependencyIndexes: file_ClockParkGetOngoingScriptInfoScRsp_proto_depIdxs,
		MessageInfos:      file_ClockParkGetOngoingScriptInfoScRsp_proto_msgTypes,
	}.Build()
	File_ClockParkGetOngoingScriptInfoScRsp_proto = out.File
	file_ClockParkGetOngoingScriptInfoScRsp_proto_rawDesc = nil
	file_ClockParkGetOngoingScriptInfoScRsp_proto_goTypes = nil
	file_ClockParkGetOngoingScriptInfoScRsp_proto_depIdxs = nil
}
