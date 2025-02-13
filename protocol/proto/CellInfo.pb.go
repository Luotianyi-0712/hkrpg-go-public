// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CellInfo.proto

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

type CellInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndCellId   uint32            `protobuf:"varint,10,opt,name=end_cell_id,json=endCellId,proto3" json:"end_cell_id,omitempty"`
	CellList    []*ChessRogueCell `protobuf:"bytes,14,rep,name=cell_list,json=cellList,proto3" json:"cell_list,omitempty"`
	EPOELFJDPNI uint32            `protobuf:"varint,9,opt,name=EPOELFJDPNI,proto3" json:"EPOELFJDPNI,omitempty"`
	StartCellId uint32            `protobuf:"varint,7,opt,name=start_cell_id,json=startCellId,proto3" json:"start_cell_id,omitempty"`
	FGEOEOPLMCP uint32            `protobuf:"varint,5,opt,name=FGEOEOPLMCP,proto3" json:"FGEOEOPLMCP,omitempty"`
}

func (x *CellInfo) Reset() {
	*x = CellInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CellInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CellInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CellInfo) ProtoMessage() {}

func (x *CellInfo) ProtoReflect() protoreflect.Message {
	mi := &file_CellInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CellInfo.ProtoReflect.Descriptor instead.
func (*CellInfo) Descriptor() ([]byte, []int) {
	return file_CellInfo_proto_rawDescGZIP(), []int{0}
}

func (x *CellInfo) GetEndCellId() uint32 {
	if x != nil {
		return x.EndCellId
	}
	return 0
}

func (x *CellInfo) GetCellList() []*ChessRogueCell {
	if x != nil {
		return x.CellList
	}
	return nil
}

func (x *CellInfo) GetEPOELFJDPNI() uint32 {
	if x != nil {
		return x.EPOELFJDPNI
	}
	return 0
}

func (x *CellInfo) GetStartCellId() uint32 {
	if x != nil {
		return x.StartCellId
	}
	return 0
}

func (x *CellInfo) GetFGEOEOPLMCP() uint32 {
	if x != nil {
		return x.FGEOEOPLMCP
	}
	return 0
}

var File_CellInfo_proto protoreflect.FileDescriptor

var file_CellInfo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x65, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x43, 0x65, 0x6c, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x43, 0x65, 0x6c,
	0x6c, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x50, 0x4f, 0x45, 0x4c, 0x46, 0x4a, 0x44, 0x50, 0x4e, 0x49,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x50, 0x4f, 0x45, 0x4c, 0x46, 0x4a, 0x44,
	0x50, 0x4e, 0x49, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x65, 0x6c,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x47, 0x45, 0x4f, 0x45,
	0x4f, 0x50, 0x4c, 0x4d, 0x43, 0x50, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x47,
	0x45, 0x4f, 0x45, 0x4f, 0x50, 0x4c, 0x4d, 0x43, 0x50, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_CellInfo_proto_rawDescOnce sync.Once
	file_CellInfo_proto_rawDescData = file_CellInfo_proto_rawDesc
)

func file_CellInfo_proto_rawDescGZIP() []byte {
	file_CellInfo_proto_rawDescOnce.Do(func() {
		file_CellInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_CellInfo_proto_rawDescData)
	})
	return file_CellInfo_proto_rawDescData
}

var file_CellInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CellInfo_proto_goTypes = []interface{}{
	(*CellInfo)(nil),       // 0: CellInfo
	(*ChessRogueCell)(nil), // 1: ChessRogueCell
}
var file_CellInfo_proto_depIdxs = []int32{
	1, // 0: CellInfo.cell_list:type_name -> ChessRogueCell
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CellInfo_proto_init() }
func file_CellInfo_proto_init() {
	if File_CellInfo_proto != nil {
		return
	}
	file_ChessRogueCell_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CellInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CellInfo); i {
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
			RawDescriptor: file_CellInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CellInfo_proto_goTypes,
		DependencyIndexes: file_CellInfo_proto_depIdxs,
		MessageInfos:      file_CellInfo_proto_msgTypes,
	}.Build()
	File_CellInfo_proto = out.File
	file_CellInfo_proto_rawDesc = nil
	file_CellInfo_proto_goTypes = nil
	file_CellInfo_proto_depIdxs = nil
}
