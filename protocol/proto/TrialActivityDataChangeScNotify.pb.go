// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TrialActivityDataChangeScNotify.proto

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

type TrialActivityDataChangeScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrialActivityInfo *TrialActivityInfo `protobuf:"bytes,3,opt,name=trial_activity_info,json=trialActivityInfo,proto3" json:"trial_activity_info,omitempty"`
}

func (x *TrialActivityDataChangeScNotify) Reset() {
	*x = TrialActivityDataChangeScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TrialActivityDataChangeScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrialActivityDataChangeScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrialActivityDataChangeScNotify) ProtoMessage() {}

func (x *TrialActivityDataChangeScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_TrialActivityDataChangeScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrialActivityDataChangeScNotify.ProtoReflect.Descriptor instead.
func (*TrialActivityDataChangeScNotify) Descriptor() ([]byte, []int) {
	return file_TrialActivityDataChangeScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *TrialActivityDataChangeScNotify) GetTrialActivityInfo() *TrialActivityInfo {
	if x != nil {
		return x.TrialActivityInfo
	}
	return nil
}

var File_TrialActivityDataChangeScNotify_proto protoreflect.FileDescriptor

var file_TrialActivityDataChangeScNotify_proto_rawDesc = []byte{
	0x0a, 0x25, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x65, 0x0a, 0x1f, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x42, 0x0a, 0x13, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TrialActivityDataChangeScNotify_proto_rawDescOnce sync.Once
	file_TrialActivityDataChangeScNotify_proto_rawDescData = file_TrialActivityDataChangeScNotify_proto_rawDesc
)

func file_TrialActivityDataChangeScNotify_proto_rawDescGZIP() []byte {
	file_TrialActivityDataChangeScNotify_proto_rawDescOnce.Do(func() {
		file_TrialActivityDataChangeScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_TrialActivityDataChangeScNotify_proto_rawDescData)
	})
	return file_TrialActivityDataChangeScNotify_proto_rawDescData
}

var file_TrialActivityDataChangeScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TrialActivityDataChangeScNotify_proto_goTypes = []interface{}{
	(*TrialActivityDataChangeScNotify)(nil), // 0: TrialActivityDataChangeScNotify
	(*TrialActivityInfo)(nil),               // 1: TrialActivityInfo
}
var file_TrialActivityDataChangeScNotify_proto_depIdxs = []int32{
	1, // 0: TrialActivityDataChangeScNotify.trial_activity_info:type_name -> TrialActivityInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TrialActivityDataChangeScNotify_proto_init() }
func file_TrialActivityDataChangeScNotify_proto_init() {
	if File_TrialActivityDataChangeScNotify_proto != nil {
		return
	}
	file_TrialActivityInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TrialActivityDataChangeScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrialActivityDataChangeScNotify); i {
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
			RawDescriptor: file_TrialActivityDataChangeScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TrialActivityDataChangeScNotify_proto_goTypes,
		DependencyIndexes: file_TrialActivityDataChangeScNotify_proto_depIdxs,
		MessageInfos:      file_TrialActivityDataChangeScNotify_proto_msgTypes,
	}.Build()
	File_TrialActivityDataChangeScNotify_proto = out.File
	file_TrialActivityDataChangeScNotify_proto_rawDesc = nil
	file_TrialActivityDataChangeScNotify_proto_goTypes = nil
	file_TrialActivityDataChangeScNotify_proto_depIdxs = nil
}
