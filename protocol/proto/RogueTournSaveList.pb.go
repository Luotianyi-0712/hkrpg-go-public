// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournSaveList.proto

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

type RogueTournSaveList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxTimes          uint32                `protobuf:"varint,7,opt,name=max_times,json=maxTimes,proto3" json:"max_times,omitempty"`
	Time              int64                 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	RogueSeasonInfo   *RogueTournSeasonInfo `protobuf:"bytes,4,opt,name=rogue_season_info,json=rogueSeasonInfo,proto3" json:"rogue_season_info,omitempty"`
	EndTime           int64                 `protobuf:"varint,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Name              string                `protobuf:"bytes,14,opt,name=name,proto3" json:"name,omitempty"`
	RogueTournCurInfo *RogueTournCurInfo    `protobuf:"bytes,15,opt,name=rogue_tourn_cur_info,json=rogueTournCurInfo,proto3" json:"rogue_tourn_cur_info,omitempty"`
	Data              *JHMIILIPJNA          `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RogueTournSaveList) Reset() {
	*x = RogueTournSaveList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournSaveList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournSaveList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournSaveList) ProtoMessage() {}

func (x *RogueTournSaveList) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournSaveList_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournSaveList.ProtoReflect.Descriptor instead.
func (*RogueTournSaveList) Descriptor() ([]byte, []int) {
	return file_RogueTournSaveList_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournSaveList) GetMaxTimes() uint32 {
	if x != nil {
		return x.MaxTimes
	}
	return 0
}

func (x *RogueTournSaveList) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *RogueTournSaveList) GetRogueSeasonInfo() *RogueTournSeasonInfo {
	if x != nil {
		return x.RogueSeasonInfo
	}
	return nil
}

func (x *RogueTournSaveList) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *RogueTournSaveList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RogueTournSaveList) GetRogueTournCurInfo() *RogueTournCurInfo {
	if x != nil {
		return x.RogueTournCurInfo
	}
	return nil
}

func (x *RogueTournSaveList) GetData() *JHMIILIPJNA {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_RogueTournSaveList_proto protoreflect.FileDescriptor

var file_RogueTournSaveList_proto_rawDesc = []byte{
	0x0a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x53, 0x61, 0x76, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x48, 0x4d, 0x49,
	0x49, 0x4c, 0x49, 0x50, 0x4a, 0x4e, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x12, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x53, 0x61, 0x76, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x11, 0x72, 0x6f,
	0x67, 0x75, 0x65, 0x5f, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75,
	0x72, 0x6e, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x72, 0x6f,
	0x67, 0x75, 0x65, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x14,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x63, 0x75, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11,
	0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4a, 0x48, 0x4d, 0x49, 0x49, 0x4c, 0x49, 0x50, 0x4a, 0x4e, 0x41, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournSaveList_proto_rawDescOnce sync.Once
	file_RogueTournSaveList_proto_rawDescData = file_RogueTournSaveList_proto_rawDesc
)

func file_RogueTournSaveList_proto_rawDescGZIP() []byte {
	file_RogueTournSaveList_proto_rawDescOnce.Do(func() {
		file_RogueTournSaveList_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournSaveList_proto_rawDescData)
	})
	return file_RogueTournSaveList_proto_rawDescData
}

var file_RogueTournSaveList_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournSaveList_proto_goTypes = []interface{}{
	(*RogueTournSaveList)(nil),   // 0: RogueTournSaveList
	(*RogueTournSeasonInfo)(nil), // 1: RogueTournSeasonInfo
	(*RogueTournCurInfo)(nil),    // 2: RogueTournCurInfo
	(*JHMIILIPJNA)(nil),          // 3: JHMIILIPJNA
}
var file_RogueTournSaveList_proto_depIdxs = []int32{
	1, // 0: RogueTournSaveList.rogue_season_info:type_name -> RogueTournSeasonInfo
	2, // 1: RogueTournSaveList.rogue_tourn_cur_info:type_name -> RogueTournCurInfo
	3, // 2: RogueTournSaveList.data:type_name -> JHMIILIPJNA
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_RogueTournSaveList_proto_init() }
func file_RogueTournSaveList_proto_init() {
	if File_RogueTournSaveList_proto != nil {
		return
	}
	file_JHMIILIPJNA_proto_init()
	file_RogueTournSeasonInfo_proto_init()
	file_RogueTournCurInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTournSaveList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournSaveList); i {
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
			RawDescriptor: file_RogueTournSaveList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournSaveList_proto_goTypes,
		DependencyIndexes: file_RogueTournSaveList_proto_depIdxs,
		MessageInfos:      file_RogueTournSaveList_proto_msgTypes,
	}.Build()
	File_RogueTournSaveList_proto = out.File
	file_RogueTournSaveList_proto_rawDesc = nil
	file_RogueTournSaveList_proto_goTypes = nil
	file_RogueTournSaveList_proto_depIdxs = nil
}
