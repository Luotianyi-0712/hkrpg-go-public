// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EnterSceneCsReq.proto

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

type EnterSceneCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryId         uint32 `protobuf:"varint,4,opt,name=entry_id,json=entryId,proto3" json:"entry_id,omitempty"`
	GameStoryLineId uint32 `protobuf:"varint,3,opt,name=game_story_line_id,json=gameStoryLineId,proto3" json:"game_story_line_id,omitempty"`
	TeleportId      uint32 `protobuf:"varint,10,opt,name=teleport_id,json=teleportId,proto3" json:"teleport_id,omitempty"`
	ContentId       uint32 `protobuf:"varint,13,opt,name=content_id,json=contentId,proto3" json:"content_id,omitempty"`
	IsCloseMap      bool   `protobuf:"varint,6,opt,name=is_close_map,json=isCloseMap,proto3" json:"is_close_map,omitempty"`
}

func (x *EnterSceneCsReq) Reset() {
	*x = EnterSceneCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnterSceneCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterSceneCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterSceneCsReq) ProtoMessage() {}

func (x *EnterSceneCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_EnterSceneCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterSceneCsReq.ProtoReflect.Descriptor instead.
func (*EnterSceneCsReq) Descriptor() ([]byte, []int) {
	return file_EnterSceneCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *EnterSceneCsReq) GetEntryId() uint32 {
	if x != nil {
		return x.EntryId
	}
	return 0
}

func (x *EnterSceneCsReq) GetGameStoryLineId() uint32 {
	if x != nil {
		return x.GameStoryLineId
	}
	return 0
}

func (x *EnterSceneCsReq) GetTeleportId() uint32 {
	if x != nil {
		return x.TeleportId
	}
	return 0
}

func (x *EnterSceneCsReq) GetContentId() uint32 {
	if x != nil {
		return x.ContentId
	}
	return 0
}

func (x *EnterSceneCsReq) GetIsCloseMap() bool {
	if x != nil {
		return x.IsCloseMap
	}
	return false
}

var File_EnterSceneCsReq_proto protoreflect.FileDescriptor

var file_EnterSceneCsReq_proto_rawDesc = []byte{
	0x0a, 0x15, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x73, 0x52, 0x65,
	0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x0f, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f,
	0x6d, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x4d, 0x61, 0x70, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EnterSceneCsReq_proto_rawDescOnce sync.Once
	file_EnterSceneCsReq_proto_rawDescData = file_EnterSceneCsReq_proto_rawDesc
)

func file_EnterSceneCsReq_proto_rawDescGZIP() []byte {
	file_EnterSceneCsReq_proto_rawDescOnce.Do(func() {
		file_EnterSceneCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterSceneCsReq_proto_rawDescData)
	})
	return file_EnterSceneCsReq_proto_rawDescData
}

var file_EnterSceneCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnterSceneCsReq_proto_goTypes = []interface{}{
	(*EnterSceneCsReq)(nil), // 0: EnterSceneCsReq
}
var file_EnterSceneCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EnterSceneCsReq_proto_init() }
func file_EnterSceneCsReq_proto_init() {
	if File_EnterSceneCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EnterSceneCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterSceneCsReq); i {
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
			RawDescriptor: file_EnterSceneCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterSceneCsReq_proto_goTypes,
		DependencyIndexes: file_EnterSceneCsReq_proto_depIdxs,
		MessageInfos:      file_EnterSceneCsReq_proto_msgTypes,
	}.Build()
	File_EnterSceneCsReq_proto = out.File
	file_EnterSceneCsReq_proto_rawDesc = nil
	file_EnterSceneCsReq_proto_goTypes = nil
	file_EnterSceneCsReq_proto_depIdxs = nil
}
