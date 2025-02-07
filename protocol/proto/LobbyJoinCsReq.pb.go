// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LobbyJoinCsReq.proto

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

type LobbyJoinCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MHDJJGDLLIG *JBEBLKIKGMP `protobuf:"bytes,14,opt,name=MHDJJGDLLIG,proto3" json:"MHDJJGDLLIG,omitempty"`
	RoomId      uint64       `protobuf:"varint,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *LobbyJoinCsReq) Reset() {
	*x = LobbyJoinCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LobbyJoinCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LobbyJoinCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LobbyJoinCsReq) ProtoMessage() {}

func (x *LobbyJoinCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_LobbyJoinCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LobbyJoinCsReq.ProtoReflect.Descriptor instead.
func (*LobbyJoinCsReq) Descriptor() ([]byte, []int) {
	return file_LobbyJoinCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *LobbyJoinCsReq) GetMHDJJGDLLIG() *JBEBLKIKGMP {
	if x != nil {
		return x.MHDJJGDLLIG
	}
	return nil
}

func (x *LobbyJoinCsReq) GetRoomId() uint64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

var File_LobbyJoinCsReq_proto protoreflect.FileDescriptor

var file_LobbyJoinCsReq_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x42, 0x45, 0x42, 0x4c, 0x4b, 0x49, 0x4b,
	0x47, 0x4d, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0e, 0x4c, 0x6f, 0x62,
	0x62, 0x79, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x0b, 0x4d,
	0x48, 0x44, 0x4a, 0x4a, 0x47, 0x44, 0x4c, 0x4c, 0x49, 0x47, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4a, 0x42, 0x45, 0x42, 0x4c, 0x4b, 0x49, 0x4b, 0x47, 0x4d, 0x50, 0x52, 0x0b,
	0x4d, 0x48, 0x44, 0x4a, 0x4a, 0x47, 0x44, 0x4c, 0x4c, 0x49, 0x47, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LobbyJoinCsReq_proto_rawDescOnce sync.Once
	file_LobbyJoinCsReq_proto_rawDescData = file_LobbyJoinCsReq_proto_rawDesc
)

func file_LobbyJoinCsReq_proto_rawDescGZIP() []byte {
	file_LobbyJoinCsReq_proto_rawDescOnce.Do(func() {
		file_LobbyJoinCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_LobbyJoinCsReq_proto_rawDescData)
	})
	return file_LobbyJoinCsReq_proto_rawDescData
}

var file_LobbyJoinCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LobbyJoinCsReq_proto_goTypes = []interface{}{
	(*LobbyJoinCsReq)(nil), // 0: LobbyJoinCsReq
	(*JBEBLKIKGMP)(nil),    // 1: JBEBLKIKGMP
}
var file_LobbyJoinCsReq_proto_depIdxs = []int32{
	1, // 0: LobbyJoinCsReq.MHDJJGDLLIG:type_name -> JBEBLKIKGMP
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LobbyJoinCsReq_proto_init() }
func file_LobbyJoinCsReq_proto_init() {
	if File_LobbyJoinCsReq_proto != nil {
		return
	}
	file_JBEBLKIKGMP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LobbyJoinCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LobbyJoinCsReq); i {
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
			RawDescriptor: file_LobbyJoinCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LobbyJoinCsReq_proto_goTypes,
		DependencyIndexes: file_LobbyJoinCsReq_proto_depIdxs,
		MessageInfos:      file_LobbyJoinCsReq_proto_msgTypes,
	}.Build()
	File_LobbyJoinCsReq_proto = out.File
	file_LobbyJoinCsReq_proto_rawDesc = nil
	file_LobbyJoinCsReq_proto_goTypes = nil
	file_LobbyJoinCsReq_proto_depIdxs = nil
}
