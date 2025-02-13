// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RotateMapCsReq.proto

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

type RotateMapCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId     uint32       `protobuf:"varint,8,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	IOKBKAOEIPC uint32       `protobuf:"varint,7,opt,name=IOKBKAOEIPC,proto3" json:"IOKBKAOEIPC,omitempty"`
	RoomMap     *OAKOJBOEMAA `protobuf:"bytes,3,opt,name=room_map,json=roomMap,proto3" json:"room_map,omitempty"`
	Motion      *MotionInfo  `protobuf:"bytes,13,opt,name=motion,proto3" json:"motion,omitempty"`
}

func (x *RotateMapCsReq) Reset() {
	*x = RotateMapCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RotateMapCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RotateMapCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RotateMapCsReq) ProtoMessage() {}

func (x *RotateMapCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_RotateMapCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RotateMapCsReq.ProtoReflect.Descriptor instead.
func (*RotateMapCsReq) Descriptor() ([]byte, []int) {
	return file_RotateMapCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *RotateMapCsReq) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *RotateMapCsReq) GetIOKBKAOEIPC() uint32 {
	if x != nil {
		return x.IOKBKAOEIPC
	}
	return 0
}

func (x *RotateMapCsReq) GetRoomMap() *OAKOJBOEMAA {
	if x != nil {
		return x.RoomMap
	}
	return nil
}

func (x *RotateMapCsReq) GetMotion() *MotionInfo {
	if x != nil {
		return x.Motion
	}
	return nil
}

var File_RotateMapCsReq_proto protoreflect.FileDescriptor

var file_RotateMapCsReq_proto_rawDesc = []byte{
	0x0a, 0x14, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x43, 0x73, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x41, 0x4b, 0x4f, 0x4a, 0x42,
	0x4f, 0x45, 0x4d, 0x41, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x0e,
	0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4f, 0x4b,
	0x42, 0x4b, 0x41, 0x4f, 0x45, 0x49, 0x50, 0x43, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x49, 0x4f, 0x4b, 0x42, 0x4b, 0x41, 0x4f, 0x45, 0x49, 0x50, 0x43, 0x12, 0x27, 0x0a, 0x08, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4f, 0x41, 0x4b, 0x4f, 0x4a, 0x42, 0x4f, 0x45, 0x4d, 0x41, 0x41, 0x52, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x4d, 0x61, 0x70, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_RotateMapCsReq_proto_rawDescOnce sync.Once
	file_RotateMapCsReq_proto_rawDescData = file_RotateMapCsReq_proto_rawDesc
)

func file_RotateMapCsReq_proto_rawDescGZIP() []byte {
	file_RotateMapCsReq_proto_rawDescOnce.Do(func() {
		file_RotateMapCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_RotateMapCsReq_proto_rawDescData)
	})
	return file_RotateMapCsReq_proto_rawDescData
}

var file_RotateMapCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RotateMapCsReq_proto_goTypes = []interface{}{
	(*RotateMapCsReq)(nil), // 0: RotateMapCsReq
	(*OAKOJBOEMAA)(nil),    // 1: OAKOJBOEMAA
	(*MotionInfo)(nil),     // 2: MotionInfo
}
var file_RotateMapCsReq_proto_depIdxs = []int32{
	1, // 0: RotateMapCsReq.room_map:type_name -> OAKOJBOEMAA
	2, // 1: RotateMapCsReq.motion:type_name -> MotionInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RotateMapCsReq_proto_init() }
func file_RotateMapCsReq_proto_init() {
	if File_RotateMapCsReq_proto != nil {
		return
	}
	file_MotionInfo_proto_init()
	file_OAKOJBOEMAA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RotateMapCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RotateMapCsReq); i {
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
			RawDescriptor: file_RotateMapCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RotateMapCsReq_proto_goTypes,
		DependencyIndexes: file_RotateMapCsReq_proto_depIdxs,
		MessageInfos:      file_RotateMapCsReq_proto_msgTypes,
	}.Build()
	File_RotateMapCsReq_proto = out.File
	file_RotateMapCsReq_proto_rawDesc = nil
	file_RotateMapCsReq_proto_goTypes = nil
	file_RotateMapCsReq_proto_depIdxs = nil
}
