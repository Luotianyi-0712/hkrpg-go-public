// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: UpdateFloorSavedValueNotify.proto

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

type UpdateFloorSavedValueNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SavedValue map[string]int32 `protobuf:"bytes,13,rep,name=saved_value,json=savedValue,proto3" json:"saved_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *UpdateFloorSavedValueNotify) Reset() {
	*x = UpdateFloorSavedValueNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpdateFloorSavedValueNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFloorSavedValueNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFloorSavedValueNotify) ProtoMessage() {}

func (x *UpdateFloorSavedValueNotify) ProtoReflect() protoreflect.Message {
	mi := &file_UpdateFloorSavedValueNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFloorSavedValueNotify.ProtoReflect.Descriptor instead.
func (*UpdateFloorSavedValueNotify) Descriptor() ([]byte, []int) {
	return file_UpdateFloorSavedValueNotify_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateFloorSavedValueNotify) GetSavedValue() map[string]int32 {
	if x != nil {
		return x.SavedValue
	}
	return nil
}

var File_UpdateFloorSavedValueNotify_proto protoreflect.FileDescriptor

var file_UpdateFloorSavedValueNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x53, 0x61, 0x76,
	0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c,
	0x6f, 0x6f, 0x72, 0x53, 0x61, 0x76, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x4d, 0x0a, 0x0b, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x53, 0x61, 0x76, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x73, 0x61, 0x76, 0x65, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x3d, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UpdateFloorSavedValueNotify_proto_rawDescOnce sync.Once
	file_UpdateFloorSavedValueNotify_proto_rawDescData = file_UpdateFloorSavedValueNotify_proto_rawDesc
)

func file_UpdateFloorSavedValueNotify_proto_rawDescGZIP() []byte {
	file_UpdateFloorSavedValueNotify_proto_rawDescOnce.Do(func() {
		file_UpdateFloorSavedValueNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_UpdateFloorSavedValueNotify_proto_rawDescData)
	})
	return file_UpdateFloorSavedValueNotify_proto_rawDescData
}

var file_UpdateFloorSavedValueNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_UpdateFloorSavedValueNotify_proto_goTypes = []interface{}{
	(*UpdateFloorSavedValueNotify)(nil), // 0: UpdateFloorSavedValueNotify
	nil,                                 // 1: UpdateFloorSavedValueNotify.SavedValueEntry
}
var file_UpdateFloorSavedValueNotify_proto_depIdxs = []int32{
	1, // 0: UpdateFloorSavedValueNotify.saved_value:type_name -> UpdateFloorSavedValueNotify.SavedValueEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UpdateFloorSavedValueNotify_proto_init() }
func file_UpdateFloorSavedValueNotify_proto_init() {
	if File_UpdateFloorSavedValueNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_UpdateFloorSavedValueNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFloorSavedValueNotify); i {
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
			RawDescriptor: file_UpdateFloorSavedValueNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UpdateFloorSavedValueNotify_proto_goTypes,
		DependencyIndexes: file_UpdateFloorSavedValueNotify_proto_depIdxs,
		MessageInfos:      file_UpdateFloorSavedValueNotify_proto_msgTypes,
	}.Build()
	File_UpdateFloorSavedValueNotify_proto = out.File
	file_UpdateFloorSavedValueNotify_proto_rawDesc = nil
	file_UpdateFloorSavedValueNotify_proto_goTypes = nil
	file_UpdateFloorSavedValueNotify_proto_depIdxs = nil
}
