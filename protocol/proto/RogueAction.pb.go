// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueAction.proto

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

type RogueAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Action:
	//
	//	*RogueAction_BuffSelectInfo
	//	*RogueAction_HBHFHCNKIBA
	//	*RogueAction_JKEJALPLIBP
	//	*RogueAction_FBIBBHIBHAA
	//	*RogueAction_MiracleSelectInfo
	//	*RogueAction_GLLOCFABBPD
	//	*RogueAction_GNOCICNIFJH
	//	*RogueAction_DBOFJOJCDKK
	//	*RogueAction_MEMAHJEECMD
	//	*RogueAction_PIKAJBPNFPF
	//	*RogueAction_FKCGFGKLIID
	//	*RogueAction_GBDDILHDOEN
	//	*RogueAction_BonusSelectInfo
	//	*RogueAction_RogueFormulaSelectInfo
	//	*RogueAction_LMGIMCMGGJL
	//	*RogueAction_KIMHACECBKF
	Action isRogueAction_Action `protobuf_oneof:"action"`
}

func (x *RogueAction) Reset() {
	*x = RogueAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueAction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueAction) ProtoMessage() {}

func (x *RogueAction) ProtoReflect() protoreflect.Message {
	mi := &file_RogueAction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueAction.ProtoReflect.Descriptor instead.
func (*RogueAction) Descriptor() ([]byte, []int) {
	return file_RogueAction_proto_rawDescGZIP(), []int{0}
}

func (m *RogueAction) GetAction() isRogueAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *RogueAction) GetBuffSelectInfo() *RogueCommonBuffSelectInfo {
	if x, ok := x.GetAction().(*RogueAction_BuffSelectInfo); ok {
		return x.BuffSelectInfo
	}
	return nil
}

func (x *RogueAction) GetHBHFHCNKIBA() *EFPIHKDCMME {
	if x, ok := x.GetAction().(*RogueAction_HBHFHCNKIBA); ok {
		return x.HBHFHCNKIBA
	}
	return nil
}

func (x *RogueAction) GetJKEJALPLIBP() *ECBKGCFNANB {
	if x, ok := x.GetAction().(*RogueAction_JKEJALPLIBP); ok {
		return x.JKEJALPLIBP
	}
	return nil
}

func (x *RogueAction) GetFBIBBHIBHAA() *ECMOGCMOJJH {
	if x, ok := x.GetAction().(*RogueAction_FBIBBHIBHAA); ok {
		return x.FBIBBHIBHAA
	}
	return nil
}

func (x *RogueAction) GetMiracleSelectInfo() *RogueMiracleSelectInfo {
	if x, ok := x.GetAction().(*RogueAction_MiracleSelectInfo); ok {
		return x.MiracleSelectInfo
	}
	return nil
}

func (x *RogueAction) GetGLLOCFABBPD() *NCHBHHHFBPM {
	if x, ok := x.GetAction().(*RogueAction_GLLOCFABBPD); ok {
		return x.GLLOCFABBPD
	}
	return nil
}

func (x *RogueAction) GetGNOCICNIFJH() *MDALALGLJJC {
	if x, ok := x.GetAction().(*RogueAction_GNOCICNIFJH); ok {
		return x.GNOCICNIFJH
	}
	return nil
}

func (x *RogueAction) GetDBOFJOJCDKK() *EPECOIGOEAF {
	if x, ok := x.GetAction().(*RogueAction_DBOFJOJCDKK); ok {
		return x.DBOFJOJCDKK
	}
	return nil
}

func (x *RogueAction) GetMEMAHJEECMD() *LJFDHIAPLAI {
	if x, ok := x.GetAction().(*RogueAction_MEMAHJEECMD); ok {
		return x.MEMAHJEECMD
	}
	return nil
}

func (x *RogueAction) GetPIKAJBPNFPF() *AIAFLIFPJFP {
	if x, ok := x.GetAction().(*RogueAction_PIKAJBPNFPF); ok {
		return x.PIKAJBPNFPF
	}
	return nil
}

func (x *RogueAction) GetFKCGFGKLIID() *IEKOLCJKMML {
	if x, ok := x.GetAction().(*RogueAction_FKCGFGKLIID); ok {
		return x.FKCGFGKLIID
	}
	return nil
}

func (x *RogueAction) GetGBDDILHDOEN() *NILLPIGHKGK {
	if x, ok := x.GetAction().(*RogueAction_GBDDILHDOEN); ok {
		return x.GBDDILHDOEN
	}
	return nil
}

func (x *RogueAction) GetBonusSelectInfo() *RogueBonusSelectInfo {
	if x, ok := x.GetAction().(*RogueAction_BonusSelectInfo); ok {
		return x.BonusSelectInfo
	}
	return nil
}

func (x *RogueAction) GetRogueFormulaSelectInfo() *RogueFormulaSelectInfo {
	if x, ok := x.GetAction().(*RogueAction_RogueFormulaSelectInfo); ok {
		return x.RogueFormulaSelectInfo
	}
	return nil
}

func (x *RogueAction) GetLMGIMCMGGJL() *OABIBGMGIIE {
	if x, ok := x.GetAction().(*RogueAction_LMGIMCMGGJL); ok {
		return x.LMGIMCMGGJL
	}
	return nil
}

func (x *RogueAction) GetKIMHACECBKF() *CHKJICPBHPA {
	if x, ok := x.GetAction().(*RogueAction_KIMHACECBKF); ok {
		return x.KIMHACECBKF
	}
	return nil
}

type isRogueAction_Action interface {
	isRogueAction_Action()
}

type RogueAction_BuffSelectInfo struct {
	BuffSelectInfo *RogueCommonBuffSelectInfo `protobuf:"bytes,820,opt,name=buff_select_info,json=buffSelectInfo,proto3,oneof"`
}

type RogueAction_HBHFHCNKIBA struct {
	HBHFHCNKIBA *EFPIHKDCMME `protobuf:"bytes,1471,opt,name=HBHFHCNKIBA,proto3,oneof"`
}

type RogueAction_JKEJALPLIBP struct {
	JKEJALPLIBP *ECBKGCFNANB `protobuf:"bytes,638,opt,name=JKEJALPLIBP,proto3,oneof"`
}

type RogueAction_FBIBBHIBHAA struct {
	FBIBBHIBHAA *ECMOGCMOJJH `protobuf:"bytes,1102,opt,name=FBIBBHIBHAA,proto3,oneof"`
}

type RogueAction_MiracleSelectInfo struct {
	MiracleSelectInfo *RogueMiracleSelectInfo `protobuf:"bytes,1995,opt,name=miracle_select_info,json=miracleSelectInfo,proto3,oneof"`
}

type RogueAction_GLLOCFABBPD struct {
	GLLOCFABBPD *NCHBHHHFBPM `protobuf:"bytes,240,opt,name=GLLOCFABBPD,proto3,oneof"`
}

type RogueAction_GNOCICNIFJH struct {
	GNOCICNIFJH *MDALALGLJJC `protobuf:"bytes,1263,opt,name=GNOCICNIFJH,proto3,oneof"`
}

type RogueAction_DBOFJOJCDKK struct {
	DBOFJOJCDKK *EPECOIGOEAF `protobuf:"bytes,674,opt,name=DBOFJOJCDKK,proto3,oneof"`
}

type RogueAction_MEMAHJEECMD struct {
	MEMAHJEECMD *LJFDHIAPLAI `protobuf:"bytes,1971,opt,name=MEMAHJEECMD,proto3,oneof"`
}

type RogueAction_PIKAJBPNFPF struct {
	PIKAJBPNFPF *AIAFLIFPJFP `protobuf:"bytes,133,opt,name=PIKAJBPNFPF,proto3,oneof"`
}

type RogueAction_FKCGFGKLIID struct {
	FKCGFGKLIID *IEKOLCJKMML `protobuf:"bytes,144,opt,name=FKCGFGKLIID,proto3,oneof"`
}

type RogueAction_GBDDILHDOEN struct {
	GBDDILHDOEN *NILLPIGHKGK `protobuf:"bytes,451,opt,name=GBDDILHDOEN,proto3,oneof"`
}

type RogueAction_BonusSelectInfo struct {
	BonusSelectInfo *RogueBonusSelectInfo `protobuf:"bytes,1849,opt,name=bonus_select_info,json=bonusSelectInfo,proto3,oneof"`
}

type RogueAction_RogueFormulaSelectInfo struct {
	RogueFormulaSelectInfo *RogueFormulaSelectInfo `protobuf:"bytes,113,opt,name=rogue_formula_select_info,json=rogueFormulaSelectInfo,proto3,oneof"`
}

type RogueAction_LMGIMCMGGJL struct {
	LMGIMCMGGJL *OABIBGMGIIE `protobuf:"bytes,305,opt,name=LMGIMCMGGJL,proto3,oneof"`
}

type RogueAction_KIMHACECBKF struct {
	KIMHACECBKF *CHKJICPBHPA `protobuf:"bytes,207,opt,name=KIMHACECBKF,proto3,oneof"`
}

func (*RogueAction_BuffSelectInfo) isRogueAction_Action() {}

func (*RogueAction_HBHFHCNKIBA) isRogueAction_Action() {}

func (*RogueAction_JKEJALPLIBP) isRogueAction_Action() {}

func (*RogueAction_FBIBBHIBHAA) isRogueAction_Action() {}

func (*RogueAction_MiracleSelectInfo) isRogueAction_Action() {}

func (*RogueAction_GLLOCFABBPD) isRogueAction_Action() {}

func (*RogueAction_GNOCICNIFJH) isRogueAction_Action() {}

func (*RogueAction_DBOFJOJCDKK) isRogueAction_Action() {}

func (*RogueAction_MEMAHJEECMD) isRogueAction_Action() {}

func (*RogueAction_PIKAJBPNFPF) isRogueAction_Action() {}

func (*RogueAction_FKCGFGKLIID) isRogueAction_Action() {}

func (*RogueAction_GBDDILHDOEN) isRogueAction_Action() {}

func (*RogueAction_BonusSelectInfo) isRogueAction_Action() {}

func (*RogueAction_RogueFormulaSelectInfo) isRogueAction_Action() {}

func (*RogueAction_LMGIMCMGGJL) isRogueAction_Action() {}

func (*RogueAction_KIMHACECBKF) isRogueAction_Action() {}

var File_RogueAction_proto protoreflect.FileDescriptor

var file_RogueAction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x48, 0x4b, 0x4a, 0x49, 0x43, 0x50, 0x42, 0x48, 0x50, 0x41,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x43, 0x42, 0x4b, 0x47, 0x43, 0x46, 0x4e,
	0x41, 0x4e, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x49, 0x4c, 0x4c, 0x50,
	0x49, 0x47, 0x48, 0x4b, 0x47, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x49, 0x41, 0x46, 0x4c,
	0x49, 0x46, 0x50, 0x4a, 0x46, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x46,
	0x50, 0x49, 0x48, 0x4b, 0x44, 0x43, 0x4d, 0x4d, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45,
	0x50, 0x45, 0x43, 0x4f, 0x49, 0x47, 0x4f, 0x45, 0x41, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x4f, 0x41, 0x42, 0x49, 0x42, 0x47, 0x4d, 0x47, 0x49, 0x49, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x4a, 0x46, 0x44, 0x48, 0x49, 0x41, 0x50, 0x4c, 0x41, 0x49,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x45, 0x4b, 0x4f, 0x4c, 0x43, 0x4a, 0x4b,
	0x4d, 0x4d, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x44, 0x41, 0x4c, 0x41, 0x4c, 0x47,
	0x4c, 0x4a, 0x4a, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x43, 0x4d, 0x4f,
	0x47, 0x43, 0x4d, 0x4f, 0x4a, 0x4a, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e,
	0x43, 0x48, 0x42, 0x48, 0x48, 0x48, 0x46, 0x42, 0x50, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xac, 0x07, 0x0a, 0x0b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x47, 0x0a, 0x10, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0xb4, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0b, 0x48, 0x42, 0x48,
	0x46, 0x48, 0x43, 0x4e, 0x4b, 0x49, 0x42, 0x41, 0x18, 0xbf, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x45, 0x46, 0x50, 0x49, 0x48, 0x4b, 0x44, 0x43, 0x4d, 0x4d, 0x45, 0x48, 0x00, 0x52,
	0x0b, 0x48, 0x42, 0x48, 0x46, 0x48, 0x43, 0x4e, 0x4b, 0x49, 0x42, 0x41, 0x12, 0x31, 0x0a, 0x0b,
	0x4a, 0x4b, 0x45, 0x4a, 0x41, 0x4c, 0x50, 0x4c, 0x49, 0x42, 0x50, 0x18, 0xfe, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x43, 0x42, 0x4b, 0x47, 0x43, 0x46, 0x4e, 0x41, 0x4e, 0x42,
	0x48, 0x00, 0x52, 0x0b, 0x4a, 0x4b, 0x45, 0x4a, 0x41, 0x4c, 0x50, 0x4c, 0x49, 0x42, 0x50, 0x12,
	0x31, 0x0a, 0x0b, 0x46, 0x42, 0x49, 0x42, 0x42, 0x48, 0x49, 0x42, 0x48, 0x41, 0x41, 0x18, 0xce,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x43, 0x4d, 0x4f, 0x47, 0x43, 0x4d, 0x4f,
	0x4a, 0x4a, 0x48, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x42, 0x49, 0x42, 0x42, 0x48, 0x49, 0x42, 0x48,
	0x41, 0x41, 0x12, 0x4a, 0x0a, 0x13, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xcb, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x11, 0x6d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31,
	0x0a, 0x0b, 0x47, 0x4c, 0x4c, 0x4f, 0x43, 0x46, 0x41, 0x42, 0x42, 0x50, 0x44, 0x18, 0xf0, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x43, 0x48, 0x42, 0x48, 0x48, 0x48, 0x46, 0x42,
	0x50, 0x4d, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x4c, 0x4c, 0x4f, 0x43, 0x46, 0x41, 0x42, 0x42, 0x50,
	0x44, 0x12, 0x31, 0x0a, 0x0b, 0x47, 0x4e, 0x4f, 0x43, 0x49, 0x43, 0x4e, 0x49, 0x46, 0x4a, 0x48,
	0x18, 0xef, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x44, 0x41, 0x4c, 0x41, 0x4c,
	0x47, 0x4c, 0x4a, 0x4a, 0x43, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x4e, 0x4f, 0x43, 0x49, 0x43, 0x4e,
	0x49, 0x46, 0x4a, 0x48, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x42, 0x4f, 0x46, 0x4a, 0x4f, 0x4a, 0x43,
	0x44, 0x4b, 0x4b, 0x18, 0xa2, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x50, 0x45,
	0x43, 0x4f, 0x49, 0x47, 0x4f, 0x45, 0x41, 0x46, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x42, 0x4f, 0x46,
	0x4a, 0x4f, 0x4a, 0x43, 0x44, 0x4b, 0x4b, 0x12, 0x31, 0x0a, 0x0b, 0x4d, 0x45, 0x4d, 0x41, 0x48,
	0x4a, 0x45, 0x45, 0x43, 0x4d, 0x44, 0x18, 0xb3, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4c, 0x4a, 0x46, 0x44, 0x48, 0x49, 0x41, 0x50, 0x4c, 0x41, 0x49, 0x48, 0x00, 0x52, 0x0b, 0x4d,
	0x45, 0x4d, 0x41, 0x48, 0x4a, 0x45, 0x45, 0x43, 0x4d, 0x44, 0x12, 0x31, 0x0a, 0x0b, 0x50, 0x49,
	0x4b, 0x41, 0x4a, 0x42, 0x50, 0x4e, 0x46, 0x50, 0x46, 0x18, 0x85, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x41, 0x49, 0x41, 0x46, 0x4c, 0x49, 0x46, 0x50, 0x4a, 0x46, 0x50, 0x48, 0x00,
	0x52, 0x0b, 0x50, 0x49, 0x4b, 0x41, 0x4a, 0x42, 0x50, 0x4e, 0x46, 0x50, 0x46, 0x12, 0x31, 0x0a,
	0x0b, 0x46, 0x4b, 0x43, 0x47, 0x46, 0x47, 0x4b, 0x4c, 0x49, 0x49, 0x44, 0x18, 0x90, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x45, 0x4b, 0x4f, 0x4c, 0x43, 0x4a, 0x4b, 0x4d, 0x4d,
	0x4c, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x4b, 0x43, 0x47, 0x46, 0x47, 0x4b, 0x4c, 0x49, 0x49, 0x44,
	0x12, 0x31, 0x0a, 0x0b, 0x47, 0x42, 0x44, 0x44, 0x49, 0x4c, 0x48, 0x44, 0x4f, 0x45, 0x4e, 0x18,
	0xc3, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x49, 0x4c, 0x4c, 0x50, 0x49, 0x47,
	0x48, 0x4b, 0x47, 0x4b, 0x48, 0x00, 0x52, 0x0b, 0x47, 0x42, 0x44, 0x44, 0x49, 0x4c, 0x48, 0x44,
	0x4f, 0x45, 0x4e, 0x12, 0x44, 0x0a, 0x11, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xb9, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0f, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x54, 0x0a, 0x19, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x71, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x16, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x46, 0x6f,
	0x72, 0x6d, 0x75, 0x6c, 0x61, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x31, 0x0a, 0x0b, 0x4c, 0x4d, 0x47, 0x49, 0x4d, 0x43, 0x4d, 0x47, 0x47, 0x4a, 0x4c, 0x18, 0xb1,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x41, 0x42, 0x49, 0x42, 0x47, 0x4d, 0x47,
	0x49, 0x49, 0x45, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x4d, 0x47, 0x49, 0x4d, 0x43, 0x4d, 0x47, 0x47,
	0x4a, 0x4c, 0x12, 0x31, 0x0a, 0x0b, 0x4b, 0x49, 0x4d, 0x48, 0x41, 0x43, 0x45, 0x43, 0x42, 0x4b,
	0x46, 0x18, 0xcf, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x48, 0x4b, 0x4a, 0x49,
	0x43, 0x50, 0x42, 0x48, 0x50, 0x41, 0x48, 0x00, 0x52, 0x0b, 0x4b, 0x49, 0x4d, 0x48, 0x41, 0x43,
	0x45, 0x43, 0x42, 0x4b, 0x46, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_RogueAction_proto_rawDescOnce sync.Once
	file_RogueAction_proto_rawDescData = file_RogueAction_proto_rawDesc
)

func file_RogueAction_proto_rawDescGZIP() []byte {
	file_RogueAction_proto_rawDescOnce.Do(func() {
		file_RogueAction_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueAction_proto_rawDescData)
	})
	return file_RogueAction_proto_rawDescData
}

var file_RogueAction_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueAction_proto_goTypes = []interface{}{
	(*RogueAction)(nil),               // 0: RogueAction
	(*RogueCommonBuffSelectInfo)(nil), // 1: RogueCommonBuffSelectInfo
	(*EFPIHKDCMME)(nil),               // 2: EFPIHKDCMME
	(*ECBKGCFNANB)(nil),               // 3: ECBKGCFNANB
	(*ECMOGCMOJJH)(nil),               // 4: ECMOGCMOJJH
	(*RogueMiracleSelectInfo)(nil),    // 5: RogueMiracleSelectInfo
	(*NCHBHHHFBPM)(nil),               // 6: NCHBHHHFBPM
	(*MDALALGLJJC)(nil),               // 7: MDALALGLJJC
	(*EPECOIGOEAF)(nil),               // 8: EPECOIGOEAF
	(*LJFDHIAPLAI)(nil),               // 9: LJFDHIAPLAI
	(*AIAFLIFPJFP)(nil),               // 10: AIAFLIFPJFP
	(*IEKOLCJKMML)(nil),               // 11: IEKOLCJKMML
	(*NILLPIGHKGK)(nil),               // 12: NILLPIGHKGK
	(*RogueBonusSelectInfo)(nil),      // 13: RogueBonusSelectInfo
	(*RogueFormulaSelectInfo)(nil),    // 14: RogueFormulaSelectInfo
	(*OABIBGMGIIE)(nil),               // 15: OABIBGMGIIE
	(*CHKJICPBHPA)(nil),               // 16: CHKJICPBHPA
}
var file_RogueAction_proto_depIdxs = []int32{
	1,  // 0: RogueAction.buff_select_info:type_name -> RogueCommonBuffSelectInfo
	2,  // 1: RogueAction.HBHFHCNKIBA:type_name -> EFPIHKDCMME
	3,  // 2: RogueAction.JKEJALPLIBP:type_name -> ECBKGCFNANB
	4,  // 3: RogueAction.FBIBBHIBHAA:type_name -> ECMOGCMOJJH
	5,  // 4: RogueAction.miracle_select_info:type_name -> RogueMiracleSelectInfo
	6,  // 5: RogueAction.GLLOCFABBPD:type_name -> NCHBHHHFBPM
	7,  // 6: RogueAction.GNOCICNIFJH:type_name -> MDALALGLJJC
	8,  // 7: RogueAction.DBOFJOJCDKK:type_name -> EPECOIGOEAF
	9,  // 8: RogueAction.MEMAHJEECMD:type_name -> LJFDHIAPLAI
	10, // 9: RogueAction.PIKAJBPNFPF:type_name -> AIAFLIFPJFP
	11, // 10: RogueAction.FKCGFGKLIID:type_name -> IEKOLCJKMML
	12, // 11: RogueAction.GBDDILHDOEN:type_name -> NILLPIGHKGK
	13, // 12: RogueAction.bonus_select_info:type_name -> RogueBonusSelectInfo
	14, // 13: RogueAction.rogue_formula_select_info:type_name -> RogueFormulaSelectInfo
	15, // 14: RogueAction.LMGIMCMGGJL:type_name -> OABIBGMGIIE
	16, // 15: RogueAction.KIMHACECBKF:type_name -> CHKJICPBHPA
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_RogueAction_proto_init() }
func file_RogueAction_proto_init() {
	if File_RogueAction_proto != nil {
		return
	}
	file_CHKJICPBHPA_proto_init()
	file_ECBKGCFNANB_proto_init()
	file_NILLPIGHKGK_proto_init()
	file_RogueCommonBuffSelectInfo_proto_init()
	file_RogueBonusSelectInfo_proto_init()
	file_AIAFLIFPJFP_proto_init()
	file_EFPIHKDCMME_proto_init()
	file_RogueFormulaSelectInfo_proto_init()
	file_EPECOIGOEAF_proto_init()
	file_OABIBGMGIIE_proto_init()
	file_LJFDHIAPLAI_proto_init()
	file_IEKOLCJKMML_proto_init()
	file_RogueMiracleSelectInfo_proto_init()
	file_MDALALGLJJC_proto_init()
	file_ECMOGCMOJJH_proto_init()
	file_NCHBHHHFBPM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueAction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueAction); i {
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
	file_RogueAction_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RogueAction_BuffSelectInfo)(nil),
		(*RogueAction_HBHFHCNKIBA)(nil),
		(*RogueAction_JKEJALPLIBP)(nil),
		(*RogueAction_FBIBBHIBHAA)(nil),
		(*RogueAction_MiracleSelectInfo)(nil),
		(*RogueAction_GLLOCFABBPD)(nil),
		(*RogueAction_GNOCICNIFJH)(nil),
		(*RogueAction_DBOFJOJCDKK)(nil),
		(*RogueAction_MEMAHJEECMD)(nil),
		(*RogueAction_PIKAJBPNFPF)(nil),
		(*RogueAction_FKCGFGKLIID)(nil),
		(*RogueAction_GBDDILHDOEN)(nil),
		(*RogueAction_BonusSelectInfo)(nil),
		(*RogueAction_RogueFormulaSelectInfo)(nil),
		(*RogueAction_LMGIMCMGGJL)(nil),
		(*RogueAction_KIMHACECBKF)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueAction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueAction_proto_goTypes,
		DependencyIndexes: file_RogueAction_proto_depIdxs,
		MessageInfos:      file_RogueAction_proto_msgTypes,
	}.Build()
	File_RogueAction_proto = out.File
	file_RogueAction_proto_rawDesc = nil
	file_RogueAction_proto_goTypes = nil
	file_RogueAction_proto_depIdxs = nil
}
