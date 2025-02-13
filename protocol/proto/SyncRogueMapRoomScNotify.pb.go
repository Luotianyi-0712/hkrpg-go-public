// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SyncRogueMapRoomScNotify.proto

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

type SyncRogueMapRoomScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapId   uint32     `protobuf:"varint,6,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	CurRoom *RogueRoom `protobuf:"bytes,8,opt,name=cur_room,json=curRoom,proto3" json:"cur_room,omitempty"`
}

func (x *SyncRogueMapRoomScNotify) Reset() {
	*x = SyncRogueMapRoomScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SyncRogueMapRoomScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRogueMapRoomScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRogueMapRoomScNotify) ProtoMessage() {}

func (x *SyncRogueMapRoomScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SyncRogueMapRoomScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRogueMapRoomScNotify.ProtoReflect.Descriptor instead.
func (*SyncRogueMapRoomScNotify) Descriptor() ([]byte, []int) {
	return file_SyncRogueMapRoomScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRogueMapRoomScNotify) GetMapId() uint32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *SyncRogueMapRoomScNotify) GetCurRoom() *RogueRoom {
	if x != nil {
		return x.CurRoom
	}
	return nil
}

var File_SyncRogueMapRoomScNotify_proto protoreflect.FileDescriptor

var file_SyncRogueMapRoomScNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x6f,
	0x6f, 0x6d, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x58, 0x0a, 0x18, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61,
	0x70, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x15, 0x0a,
	0x06, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d,
	0x61, 0x70, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x6d,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x07, 0x63, 0x75, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_SyncRogueMapRoomScNotify_proto_rawDescOnce sync.Once
	file_SyncRogueMapRoomScNotify_proto_rawDescData = file_SyncRogueMapRoomScNotify_proto_rawDesc
)

func file_SyncRogueMapRoomScNotify_proto_rawDescGZIP() []byte {
	file_SyncRogueMapRoomScNotify_proto_rawDescOnce.Do(func() {
		file_SyncRogueMapRoomScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SyncRogueMapRoomScNotify_proto_rawDescData)
	})
	return file_SyncRogueMapRoomScNotify_proto_rawDescData
}

var file_SyncRogueMapRoomScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SyncRogueMapRoomScNotify_proto_goTypes = []interface{}{
	(*SyncRogueMapRoomScNotify)(nil), // 0: SyncRogueMapRoomScNotify
	(*RogueRoom)(nil),                // 1: RogueRoom
}
var file_SyncRogueMapRoomScNotify_proto_depIdxs = []int32{
	1, // 0: SyncRogueMapRoomScNotify.cur_room:type_name -> RogueRoom
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SyncRogueMapRoomScNotify_proto_init() }
func file_SyncRogueMapRoomScNotify_proto_init() {
	if File_SyncRogueMapRoomScNotify_proto != nil {
		return
	}
	file_RogueRoom_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SyncRogueMapRoomScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRogueMapRoomScNotify); i {
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
			RawDescriptor: file_SyncRogueMapRoomScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SyncRogueMapRoomScNotify_proto_goTypes,
		DependencyIndexes: file_SyncRogueMapRoomScNotify_proto_depIdxs,
		MessageInfos:      file_SyncRogueMapRoomScNotify_proto_msgTypes,
	}.Build()
	File_SyncRogueMapRoomScNotify_proto = out.File
	file_SyncRogueMapRoomScNotify_proto_rawDesc = nil
	file_SyncRogueMapRoomScNotify_proto_goTypes = nil
	file_SyncRogueMapRoomScNotify_proto_depIdxs = nil
}
