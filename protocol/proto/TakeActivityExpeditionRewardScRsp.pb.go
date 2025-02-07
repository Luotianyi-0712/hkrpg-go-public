// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TakeActivityExpeditionRewardScRsp.proto

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

type TakeActivityExpeditionRewardScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32    `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Reward      *ItemList `protobuf:"bytes,12,opt,name=reward,proto3" json:"reward,omitempty"`
	NMDJEDIJDKA uint32    `protobuf:"varint,6,opt,name=NMDJEDIJDKA,proto3" json:"NMDJEDIJDKA,omitempty"`
	JBNOBKKJMNN *ItemList `protobuf:"bytes,4,opt,name=JBNOBKKJMNN,proto3" json:"JBNOBKKJMNN,omitempty"`
	ScoreId     uint32    `protobuf:"varint,5,opt,name=score_id,json=scoreId,proto3" json:"score_id,omitempty"`
}

func (x *TakeActivityExpeditionRewardScRsp) Reset() {
	*x = TakeActivityExpeditionRewardScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeActivityExpeditionRewardScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeActivityExpeditionRewardScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeActivityExpeditionRewardScRsp) ProtoMessage() {}

func (x *TakeActivityExpeditionRewardScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TakeActivityExpeditionRewardScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeActivityExpeditionRewardScRsp.ProtoReflect.Descriptor instead.
func (*TakeActivityExpeditionRewardScRsp) Descriptor() ([]byte, []int) {
	return file_TakeActivityExpeditionRewardScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TakeActivityExpeditionRewardScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *TakeActivityExpeditionRewardScRsp) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *TakeActivityExpeditionRewardScRsp) GetNMDJEDIJDKA() uint32 {
	if x != nil {
		return x.NMDJEDIJDKA
	}
	return 0
}

func (x *TakeActivityExpeditionRewardScRsp) GetJBNOBKKJMNN() *ItemList {
	if x != nil {
		return x.JBNOBKKJMNN
	}
	return nil
}

func (x *TakeActivityExpeditionRewardScRsp) GetScoreId() uint32 {
	if x != nil {
		return x.ScoreId
	}
	return 0
}

var File_TakeActivityExpeditionRewardScRsp_proto protoreflect.FileDescriptor

var file_TakeActivityExpeditionRewardScRsp_proto_rawDesc = []byte{
	0x0a, 0x27, 0x54, 0x61, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x78,
	0x70, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x21, 0x54, 0x61,
	0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x45, 0x78, 0x70, 0x65, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x4e, 0x4d, 0x44, 0x4a, 0x45, 0x44, 0x49, 0x4a, 0x44, 0x4b, 0x41, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4e, 0x4d, 0x44, 0x4a, 0x45, 0x44, 0x49, 0x4a, 0x44, 0x4b, 0x41, 0x12, 0x2b,
	0x0a, 0x0b, 0x4a, 0x42, 0x4e, 0x4f, 0x42, 0x4b, 0x4b, 0x4a, 0x4d, 0x4e, 0x4e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b,
	0x4a, 0x42, 0x4e, 0x4f, 0x42, 0x4b, 0x4b, 0x4a, 0x4d, 0x4e, 0x4e, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TakeActivityExpeditionRewardScRsp_proto_rawDescOnce sync.Once
	file_TakeActivityExpeditionRewardScRsp_proto_rawDescData = file_TakeActivityExpeditionRewardScRsp_proto_rawDesc
)

func file_TakeActivityExpeditionRewardScRsp_proto_rawDescGZIP() []byte {
	file_TakeActivityExpeditionRewardScRsp_proto_rawDescOnce.Do(func() {
		file_TakeActivityExpeditionRewardScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeActivityExpeditionRewardScRsp_proto_rawDescData)
	})
	return file_TakeActivityExpeditionRewardScRsp_proto_rawDescData
}

var file_TakeActivityExpeditionRewardScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeActivityExpeditionRewardScRsp_proto_goTypes = []interface{}{
	(*TakeActivityExpeditionRewardScRsp)(nil), // 0: TakeActivityExpeditionRewardScRsp
	(*ItemList)(nil), // 1: ItemList
}
var file_TakeActivityExpeditionRewardScRsp_proto_depIdxs = []int32{
	1, // 0: TakeActivityExpeditionRewardScRsp.reward:type_name -> ItemList
	1, // 1: TakeActivityExpeditionRewardScRsp.JBNOBKKJMNN:type_name -> ItemList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TakeActivityExpeditionRewardScRsp_proto_init() }
func file_TakeActivityExpeditionRewardScRsp_proto_init() {
	if File_TakeActivityExpeditionRewardScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TakeActivityExpeditionRewardScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeActivityExpeditionRewardScRsp); i {
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
			RawDescriptor: file_TakeActivityExpeditionRewardScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeActivityExpeditionRewardScRsp_proto_goTypes,
		DependencyIndexes: file_TakeActivityExpeditionRewardScRsp_proto_depIdxs,
		MessageInfos:      file_TakeActivityExpeditionRewardScRsp_proto_msgTypes,
	}.Build()
	File_TakeActivityExpeditionRewardScRsp_proto = out.File
	file_TakeActivityExpeditionRewardScRsp_proto_rawDesc = nil
	file_TakeActivityExpeditionRewardScRsp_proto_goTypes = nil
	file_TakeActivityExpeditionRewardScRsp_proto_depIdxs = nil
}
