// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SwordTrainingAction.proto

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

type SwordTrainingAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INMOIILMOJN *EFLKKNBMPHB `protobuf:"bytes,5,opt,name=INMOIILMOJN,proto3" json:"INMOIILMOJN,omitempty"`
	KABMHIGOCHM *LCEMAIAAPCA `protobuf:"bytes,1,opt,name=KABMHIGOCHM,proto3" json:"KABMHIGOCHM,omitempty"`
	// Types that are assignable to Action:
	//
	//	*SwordTrainingAction_PEIIECHJBOH
	//	*SwordTrainingAction_LLBPGOIFDAD
	//	*SwordTrainingAction_SubMood
	//	*SwordTrainingAction_BJLIOIKGPCK
	//	*SwordTrainingAction_BHCEKMIEFKC
	//	*SwordTrainingAction_LDMGDOGJKNC
	//	*SwordTrainingAction_NOJFNAKKOJO
	//	*SwordTrainingAction_DOBMEMOFBDE
	//	*SwordTrainingAction_FHFOFHBEIGI
	//	*SwordTrainingAction_DMOGPPBEFNA
	//	*SwordTrainingAction_MaxMood
	//	*SwordTrainingAction_IFBIMMDFOED
	//	*SwordTrainingAction_KENCPJJECKL
	//	*SwordTrainingAction_NHNNHJBBGAB
	Action isSwordTrainingAction_Action `protobuf_oneof:"action"`
}

func (x *SwordTrainingAction) Reset() {
	*x = SwordTrainingAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SwordTrainingAction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwordTrainingAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwordTrainingAction) ProtoMessage() {}

func (x *SwordTrainingAction) ProtoReflect() protoreflect.Message {
	mi := &file_SwordTrainingAction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwordTrainingAction.ProtoReflect.Descriptor instead.
func (*SwordTrainingAction) Descriptor() ([]byte, []int) {
	return file_SwordTrainingAction_proto_rawDescGZIP(), []int{0}
}

func (x *SwordTrainingAction) GetINMOIILMOJN() *EFLKKNBMPHB {
	if x != nil {
		return x.INMOIILMOJN
	}
	return nil
}

func (x *SwordTrainingAction) GetKABMHIGOCHM() *LCEMAIAAPCA {
	if x != nil {
		return x.KABMHIGOCHM
	}
	return nil
}

func (m *SwordTrainingAction) GetAction() isSwordTrainingAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *SwordTrainingAction) GetPEIIECHJBOH() *EJNNNBLNJIC {
	if x, ok := x.GetAction().(*SwordTrainingAction_PEIIECHJBOH); ok {
		return x.PEIIECHJBOH
	}
	return nil
}

func (x *SwordTrainingAction) GetLLBPGOIFDAD() *ChangeMoodInfo {
	if x, ok := x.GetAction().(*SwordTrainingAction_LLBPGOIFDAD); ok {
		return x.LLBPGOIFDAD
	}
	return nil
}

func (x *SwordTrainingAction) GetSubMood() *ChangeMoodInfo {
	if x, ok := x.GetAction().(*SwordTrainingAction_SubMood); ok {
		return x.SubMood
	}
	return nil
}

func (x *SwordTrainingAction) GetBJLIOIKGPCK() *ANDOCAGGDMH {
	if x, ok := x.GetAction().(*SwordTrainingAction_BJLIOIKGPCK); ok {
		return x.BJLIOIKGPCK
	}
	return nil
}

func (x *SwordTrainingAction) GetBHCEKMIEFKC() *ANDOCAGGDMH {
	if x, ok := x.GetAction().(*SwordTrainingAction_BHCEKMIEFKC); ok {
		return x.BHCEKMIEFKC
	}
	return nil
}

func (x *SwordTrainingAction) GetLDMGDOGJKNC() *HNHNFFFGFJC {
	if x, ok := x.GetAction().(*SwordTrainingAction_LDMGDOGJKNC); ok {
		return x.LDMGDOGJKNC
	}
	return nil
}

func (x *SwordTrainingAction) GetNOJFNAKKOJO() *OEJIIDGGGBO {
	if x, ok := x.GetAction().(*SwordTrainingAction_NOJFNAKKOJO); ok {
		return x.NOJFNAKKOJO
	}
	return nil
}

func (x *SwordTrainingAction) GetDOBMEMOFBDE() *DLBAMGFIACN {
	if x, ok := x.GetAction().(*SwordTrainingAction_DOBMEMOFBDE); ok {
		return x.DOBMEMOFBDE
	}
	return nil
}

func (x *SwordTrainingAction) GetFHFOFHBEIGI() *DLBAMGFIACN {
	if x, ok := x.GetAction().(*SwordTrainingAction_FHFOFHBEIGI); ok {
		return x.FHFOFHBEIGI
	}
	return nil
}

func (x *SwordTrainingAction) GetDMOGPPBEFNA() uint32 {
	if x, ok := x.GetAction().(*SwordTrainingAction_DMOGPPBEFNA); ok {
		return x.DMOGPPBEFNA
	}
	return 0
}

func (x *SwordTrainingAction) GetMaxMood() uint32 {
	if x, ok := x.GetAction().(*SwordTrainingAction_MaxMood); ok {
		return x.MaxMood
	}
	return 0
}

func (x *SwordTrainingAction) GetIFBIMMDFOED() *OOGAHFDMHJM {
	if x, ok := x.GetAction().(*SwordTrainingAction_IFBIMMDFOED); ok {
		return x.IFBIMMDFOED
	}
	return nil
}

func (x *SwordTrainingAction) GetKENCPJJECKL() *ENGJLGMAOAC {
	if x, ok := x.GetAction().(*SwordTrainingAction_KENCPJJECKL); ok {
		return x.KENCPJJECKL
	}
	return nil
}

func (x *SwordTrainingAction) GetNHNNHJBBGAB() *BOIFCPNAOLC {
	if x, ok := x.GetAction().(*SwordTrainingAction_NHNNHJBBGAB); ok {
		return x.NHNNHJBBGAB
	}
	return nil
}

type isSwordTrainingAction_Action interface {
	isSwordTrainingAction_Action()
}

type SwordTrainingAction_PEIIECHJBOH struct {
	PEIIECHJBOH *EJNNNBLNJIC `protobuf:"bytes,3,opt,name=PEIIECHJBOH,proto3,oneof"`
}

type SwordTrainingAction_LLBPGOIFDAD struct {
	LLBPGOIFDAD *ChangeMoodInfo `protobuf:"bytes,8,opt,name=LLBPGOIFDAD,proto3,oneof"`
}

type SwordTrainingAction_SubMood struct {
	SubMood *ChangeMoodInfo `protobuf:"bytes,15,opt,name=sub_mood,json=subMood,proto3,oneof"`
}

type SwordTrainingAction_BJLIOIKGPCK struct {
	BJLIOIKGPCK *ANDOCAGGDMH `protobuf:"bytes,14,opt,name=BJLIOIKGPCK,proto3,oneof"`
}

type SwordTrainingAction_BHCEKMIEFKC struct {
	BHCEKMIEFKC *ANDOCAGGDMH `protobuf:"bytes,2,opt,name=BHCEKMIEFKC,proto3,oneof"`
}

type SwordTrainingAction_LDMGDOGJKNC struct {
	LDMGDOGJKNC *HNHNFFFGFJC `protobuf:"bytes,11,opt,name=LDMGDOGJKNC,proto3,oneof"`
}

type SwordTrainingAction_NOJFNAKKOJO struct {
	NOJFNAKKOJO *OEJIIDGGGBO `protobuf:"bytes,10,opt,name=NOJFNAKKOJO,proto3,oneof"`
}

type SwordTrainingAction_DOBMEMOFBDE struct {
	DOBMEMOFBDE *DLBAMGFIACN `protobuf:"bytes,12,opt,name=DOBMEMOFBDE,proto3,oneof"`
}

type SwordTrainingAction_FHFOFHBEIGI struct {
	FHFOFHBEIGI *DLBAMGFIACN `protobuf:"bytes,9,opt,name=FHFOFHBEIGI,proto3,oneof"`
}

type SwordTrainingAction_DMOGPPBEFNA struct {
	DMOGPPBEFNA uint32 `protobuf:"varint,7,opt,name=DMOGPPBEFNA,proto3,oneof"`
}

type SwordTrainingAction_MaxMood struct {
	MaxMood uint32 `protobuf:"varint,623,opt,name=max_mood,json=maxMood,proto3,oneof"`
}

type SwordTrainingAction_IFBIMMDFOED struct {
	IFBIMMDFOED *OOGAHFDMHJM `protobuf:"bytes,1133,opt,name=IFBIMMDFOED,proto3,oneof"`
}

type SwordTrainingAction_KENCPJJECKL struct {
	KENCPJJECKL *ENGJLGMAOAC `protobuf:"bytes,343,opt,name=KENCPJJECKL,proto3,oneof"`
}

type SwordTrainingAction_NHNNHJBBGAB struct {
	NHNNHJBBGAB *BOIFCPNAOLC `protobuf:"bytes,747,opt,name=NHNNHJBBGAB,proto3,oneof"`
}

func (*SwordTrainingAction_PEIIECHJBOH) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_LLBPGOIFDAD) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_SubMood) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_BJLIOIKGPCK) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_BHCEKMIEFKC) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_LDMGDOGJKNC) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_NOJFNAKKOJO) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_DOBMEMOFBDE) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_FHFOFHBEIGI) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_DMOGPPBEFNA) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_MaxMood) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_IFBIMMDFOED) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_KENCPJJECKL) isSwordTrainingAction_Action() {}

func (*SwordTrainingAction_NHNNHJBBGAB) isSwordTrainingAction_Action() {}

var File_SwordTrainingAction_proto protoreflect.FileDescriptor

var file_SwordTrainingAction_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x43, 0x45,
	0x4d, 0x41, 0x49, 0x41, 0x41, 0x50, 0x43, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x48, 0x4e, 0x48, 0x4e, 0x46, 0x46, 0x46, 0x47, 0x46, 0x4a, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x45, 0x4e, 0x47, 0x4a, 0x4c, 0x47, 0x4d, 0x41, 0x4f, 0x41, 0x43, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x4f, 0x49, 0x46, 0x43, 0x50, 0x4e, 0x41, 0x4f, 0x4c,
	0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x4e, 0x44, 0x4f, 0x43, 0x41, 0x47,
	0x47, 0x44, 0x4d, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x4a, 0x4e, 0x4e,
	0x4e, 0x42, 0x4c, 0x4e, 0x4a, 0x49, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45,
	0x46, 0x4c, 0x4b, 0x4b, 0x4e, 0x42, 0x4d, 0x50, 0x48, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x4f, 0x45, 0x4a, 0x49, 0x49, 0x44, 0x47, 0x47, 0x47, 0x42, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x4c, 0x42, 0x41, 0x4d, 0x47, 0x46, 0x49, 0x41, 0x43, 0x4e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x4f, 0x47, 0x41, 0x48, 0x46, 0x44, 0x4d,
	0x48, 0x4a, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4d, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9b, 0x06, 0x0a, 0x13, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x4e, 0x4d, 0x4f, 0x49,
	0x49, 0x4c, 0x4d, 0x4f, 0x4a, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45,
	0x46, 0x4c, 0x4b, 0x4b, 0x4e, 0x42, 0x4d, 0x50, 0x48, 0x42, 0x52, 0x0b, 0x49, 0x4e, 0x4d, 0x4f,
	0x49, 0x49, 0x4c, 0x4d, 0x4f, 0x4a, 0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x4b, 0x41, 0x42, 0x4d, 0x48,
	0x49, 0x47, 0x4f, 0x43, 0x48, 0x4d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c,
	0x43, 0x45, 0x4d, 0x41, 0x49, 0x41, 0x41, 0x50, 0x43, 0x41, 0x52, 0x0b, 0x4b, 0x41, 0x42, 0x4d,
	0x48, 0x49, 0x47, 0x4f, 0x43, 0x48, 0x4d, 0x12, 0x30, 0x0a, 0x0b, 0x50, 0x45, 0x49, 0x49, 0x45,
	0x43, 0x48, 0x4a, 0x42, 0x4f, 0x48, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45,
	0x4a, 0x4e, 0x4e, 0x4e, 0x42, 0x4c, 0x4e, 0x4a, 0x49, 0x43, 0x48, 0x00, 0x52, 0x0b, 0x50, 0x45,
	0x49, 0x49, 0x45, 0x43, 0x48, 0x4a, 0x42, 0x4f, 0x48, 0x12, 0x33, 0x0a, 0x0b, 0x4c, 0x4c, 0x42,
	0x50, 0x47, 0x4f, 0x49, 0x46, 0x44, 0x41, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x0b, 0x4c, 0x4c, 0x42, 0x50, 0x47, 0x4f, 0x49, 0x46, 0x44, 0x41, 0x44, 0x12, 0x2c,
	0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x6d, 0x6f, 0x6f, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75, 0x62, 0x4d, 0x6f, 0x6f, 0x64, 0x12, 0x30, 0x0a, 0x0b,
	0x42, 0x4a, 0x4c, 0x49, 0x4f, 0x49, 0x4b, 0x47, 0x50, 0x43, 0x4b, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x4e, 0x44, 0x4f, 0x43, 0x41, 0x47, 0x47, 0x44, 0x4d, 0x48, 0x48,
	0x00, 0x52, 0x0b, 0x42, 0x4a, 0x4c, 0x49, 0x4f, 0x49, 0x4b, 0x47, 0x50, 0x43, 0x4b, 0x12, 0x30,
	0x0a, 0x0b, 0x42, 0x48, 0x43, 0x45, 0x4b, 0x4d, 0x49, 0x45, 0x46, 0x4b, 0x43, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x4e, 0x44, 0x4f, 0x43, 0x41, 0x47, 0x47, 0x44, 0x4d,
	0x48, 0x48, 0x00, 0x52, 0x0b, 0x42, 0x48, 0x43, 0x45, 0x4b, 0x4d, 0x49, 0x45, 0x46, 0x4b, 0x43,
	0x12, 0x30, 0x0a, 0x0b, 0x4c, 0x44, 0x4d, 0x47, 0x44, 0x4f, 0x47, 0x4a, 0x4b, 0x4e, 0x43, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x4e, 0x48, 0x4e, 0x46, 0x46, 0x46, 0x47,
	0x46, 0x4a, 0x43, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x44, 0x4d, 0x47, 0x44, 0x4f, 0x47, 0x4a, 0x4b,
	0x4e, 0x43, 0x12, 0x30, 0x0a, 0x0b, 0x4e, 0x4f, 0x4a, 0x46, 0x4e, 0x41, 0x4b, 0x4b, 0x4f, 0x4a,
	0x4f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x45, 0x4a, 0x49, 0x49, 0x44,
	0x47, 0x47, 0x47, 0x42, 0x4f, 0x48, 0x00, 0x52, 0x0b, 0x4e, 0x4f, 0x4a, 0x46, 0x4e, 0x41, 0x4b,
	0x4b, 0x4f, 0x4a, 0x4f, 0x12, 0x30, 0x0a, 0x0b, 0x44, 0x4f, 0x42, 0x4d, 0x45, 0x4d, 0x4f, 0x46,
	0x42, 0x44, 0x45, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4c, 0x42, 0x41,
	0x4d, 0x47, 0x46, 0x49, 0x41, 0x43, 0x4e, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x4f, 0x42, 0x4d, 0x45,
	0x4d, 0x4f, 0x46, 0x42, 0x44, 0x45, 0x12, 0x30, 0x0a, 0x0b, 0x46, 0x48, 0x46, 0x4f, 0x46, 0x48,
	0x42, 0x45, 0x49, 0x47, 0x49, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4c,
	0x42, 0x41, 0x4d, 0x47, 0x46, 0x49, 0x41, 0x43, 0x4e, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x48, 0x46,
	0x4f, 0x46, 0x48, 0x42, 0x45, 0x49, 0x47, 0x49, 0x12, 0x22, 0x0a, 0x0b, 0x44, 0x4d, 0x4f, 0x47,
	0x50, 0x50, 0x42, 0x45, 0x46, 0x4e, 0x41, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x0b, 0x44, 0x4d, 0x4f, 0x47, 0x50, 0x50, 0x42, 0x45, 0x46, 0x4e, 0x41, 0x12, 0x1c, 0x0a, 0x08,
	0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x6f, 0x6f, 0x64, 0x18, 0xef, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x4d, 0x6f, 0x6f, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x49, 0x46,
	0x42, 0x49, 0x4d, 0x4d, 0x44, 0x46, 0x4f, 0x45, 0x44, 0x18, 0xed, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4f, 0x4f, 0x47, 0x41, 0x48, 0x46, 0x44, 0x4d, 0x48, 0x4a, 0x4d, 0x48, 0x00,
	0x52, 0x0b, 0x49, 0x46, 0x42, 0x49, 0x4d, 0x4d, 0x44, 0x46, 0x4f, 0x45, 0x44, 0x12, 0x31, 0x0a,
	0x0b, 0x4b, 0x45, 0x4e, 0x43, 0x50, 0x4a, 0x4a, 0x45, 0x43, 0x4b, 0x4c, 0x18, 0xd7, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x4e, 0x47, 0x4a, 0x4c, 0x47, 0x4d, 0x41, 0x4f, 0x41,
	0x43, 0x48, 0x00, 0x52, 0x0b, 0x4b, 0x45, 0x4e, 0x43, 0x50, 0x4a, 0x4a, 0x45, 0x43, 0x4b, 0x4c,
	0x12, 0x31, 0x0a, 0x0b, 0x4e, 0x48, 0x4e, 0x4e, 0x48, 0x4a, 0x42, 0x42, 0x47, 0x41, 0x42, 0x18,
	0xeb, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x4f, 0x49, 0x46, 0x43, 0x50, 0x4e,
	0x41, 0x4f, 0x4c, 0x43, 0x48, 0x00, 0x52, 0x0b, 0x4e, 0x48, 0x4e, 0x4e, 0x48, 0x4a, 0x42, 0x42,
	0x47, 0x41, 0x42, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SwordTrainingAction_proto_rawDescOnce sync.Once
	file_SwordTrainingAction_proto_rawDescData = file_SwordTrainingAction_proto_rawDesc
)

func file_SwordTrainingAction_proto_rawDescGZIP() []byte {
	file_SwordTrainingAction_proto_rawDescOnce.Do(func() {
		file_SwordTrainingAction_proto_rawDescData = protoimpl.X.CompressGZIP(file_SwordTrainingAction_proto_rawDescData)
	})
	return file_SwordTrainingAction_proto_rawDescData
}

var file_SwordTrainingAction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SwordTrainingAction_proto_goTypes = []interface{}{
	(*SwordTrainingAction)(nil), // 0: SwordTrainingAction
	(*EFLKKNBMPHB)(nil),         // 1: EFLKKNBMPHB
	(*LCEMAIAAPCA)(nil),         // 2: LCEMAIAAPCA
	(*EJNNNBLNJIC)(nil),         // 3: EJNNNBLNJIC
	(*ChangeMoodInfo)(nil),      // 4: ChangeMoodInfo
	(*ANDOCAGGDMH)(nil),         // 5: ANDOCAGGDMH
	(*HNHNFFFGFJC)(nil),         // 6: HNHNFFFGFJC
	(*OEJIIDGGGBO)(nil),         // 7: OEJIIDGGGBO
	(*DLBAMGFIACN)(nil),         // 8: DLBAMGFIACN
	(*OOGAHFDMHJM)(nil),         // 9: OOGAHFDMHJM
	(*ENGJLGMAOAC)(nil),         // 10: ENGJLGMAOAC
	(*BOIFCPNAOLC)(nil),         // 11: BOIFCPNAOLC
}
var file_SwordTrainingAction_proto_depIdxs = []int32{
	1,  // 0: SwordTrainingAction.INMOIILMOJN:type_name -> EFLKKNBMPHB
	2,  // 1: SwordTrainingAction.KABMHIGOCHM:type_name -> LCEMAIAAPCA
	3,  // 2: SwordTrainingAction.PEIIECHJBOH:type_name -> EJNNNBLNJIC
	4,  // 3: SwordTrainingAction.LLBPGOIFDAD:type_name -> ChangeMoodInfo
	4,  // 4: SwordTrainingAction.sub_mood:type_name -> ChangeMoodInfo
	5,  // 5: SwordTrainingAction.BJLIOIKGPCK:type_name -> ANDOCAGGDMH
	5,  // 6: SwordTrainingAction.BHCEKMIEFKC:type_name -> ANDOCAGGDMH
	6,  // 7: SwordTrainingAction.LDMGDOGJKNC:type_name -> HNHNFFFGFJC
	7,  // 8: SwordTrainingAction.NOJFNAKKOJO:type_name -> OEJIIDGGGBO
	8,  // 9: SwordTrainingAction.DOBMEMOFBDE:type_name -> DLBAMGFIACN
	8,  // 10: SwordTrainingAction.FHFOFHBEIGI:type_name -> DLBAMGFIACN
	9,  // 11: SwordTrainingAction.IFBIMMDFOED:type_name -> OOGAHFDMHJM
	10, // 12: SwordTrainingAction.KENCPJJECKL:type_name -> ENGJLGMAOAC
	11, // 13: SwordTrainingAction.NHNNHJBBGAB:type_name -> BOIFCPNAOLC
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_SwordTrainingAction_proto_init() }
func file_SwordTrainingAction_proto_init() {
	if File_SwordTrainingAction_proto != nil {
		return
	}
	file_LCEMAIAAPCA_proto_init()
	file_HNHNFFFGFJC_proto_init()
	file_ENGJLGMAOAC_proto_init()
	file_BOIFCPNAOLC_proto_init()
	file_ANDOCAGGDMH_proto_init()
	file_EJNNNBLNJIC_proto_init()
	file_EFLKKNBMPHB_proto_init()
	file_OEJIIDGGGBO_proto_init()
	file_DLBAMGFIACN_proto_init()
	file_OOGAHFDMHJM_proto_init()
	file_ChangeMoodInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SwordTrainingAction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwordTrainingAction); i {
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
	file_SwordTrainingAction_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SwordTrainingAction_PEIIECHJBOH)(nil),
		(*SwordTrainingAction_LLBPGOIFDAD)(nil),
		(*SwordTrainingAction_SubMood)(nil),
		(*SwordTrainingAction_BJLIOIKGPCK)(nil),
		(*SwordTrainingAction_BHCEKMIEFKC)(nil),
		(*SwordTrainingAction_LDMGDOGJKNC)(nil),
		(*SwordTrainingAction_NOJFNAKKOJO)(nil),
		(*SwordTrainingAction_DOBMEMOFBDE)(nil),
		(*SwordTrainingAction_FHFOFHBEIGI)(nil),
		(*SwordTrainingAction_DMOGPPBEFNA)(nil),
		(*SwordTrainingAction_MaxMood)(nil),
		(*SwordTrainingAction_IFBIMMDFOED)(nil),
		(*SwordTrainingAction_KENCPJJECKL)(nil),
		(*SwordTrainingAction_NHNNHJBBGAB)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SwordTrainingAction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SwordTrainingAction_proto_goTypes,
		DependencyIndexes: file_SwordTrainingAction_proto_depIdxs,
		MessageInfos:      file_SwordTrainingAction_proto_msgTypes,
	}.Build()
	File_SwordTrainingAction_proto = out.File
	file_SwordTrainingAction_proto_rawDesc = nil
	file_SwordTrainingAction_proto_goTypes = nil
	file_SwordTrainingAction_proto_depIdxs = nil
}
