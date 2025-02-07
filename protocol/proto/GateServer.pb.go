// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GateServer.proto

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

type GateServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JIHFOLPBFFL     string   `protobuf:"bytes,1706,opt,name=JIHFOLPBFFL,proto3" json:"JIHFOLPBFFL,omitempty"`
	AIFBLJAACHO     string   `protobuf:"bytes,383,opt,name=AIFBLJAACHO,proto3" json:"AIFBLJAACHO,omitempty"`
	Unk6            bool     `protobuf:"varint,737,opt,name=unk6,proto3" json:"unk6,omitempty"`
	NAHBLPGMKJC     string   `protobuf:"bytes,1829,opt,name=NAHBLPGMKJC,proto3" json:"NAHBLPGMKJC,omitempty"`
	Retcode         uint32   `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	AOEKIKFKMGA     bool     `protobuf:"varint,1851,opt,name=AOEKIKFKMGA,proto3" json:"AOEKIKFKMGA,omitempty"`
	CLJNMIKINLN     string   `protobuf:"bytes,1197,opt,name=CLJNMIKINLN,proto3" json:"CLJNMIKINLN,omitempty"`
	GJMFLPHIDFD     []string `protobuf:"bytes,1394,rep,name=GJMFLPHIDFD,proto3" json:"GJMFLPHIDFD,omitempty"`
	Unk4            bool     `protobuf:"varint,210,opt,name=unk4,proto3" json:"unk4,omitempty"`
	LuaUrl          string   `protobuf:"bytes,2,opt,name=lua_url,json=luaUrl,proto3" json:"lua_url,omitempty"`
	HJPAPNPPNNH     string   `protobuf:"bytes,935,opt,name=HJPAPNPPNNH,proto3" json:"HJPAPNPPNNH,omitempty"`
	CDHOIAAONHD     bool     `protobuf:"varint,1386,opt,name=CDHOIAAONHD,proto3" json:"CDHOIAAONHD,omitempty"`
	DHKIEGBPJML     string   `protobuf:"bytes,1818,opt,name=DHKIEGBPJML,proto3" json:"DHKIEGBPJML,omitempty"`
	RegionName      string   `protobuf:"bytes,10,opt,name=region_name,json=regionName,proto3" json:"region_name,omitempty"`
	OMDAJOMPIMG     string   `protobuf:"bytes,1188,opt,name=OMDAJOMPIMG,proto3" json:"OMDAJOMPIMG,omitempty"`
	KFIGINFCAKM     string   `protobuf:"bytes,1069,opt,name=KFIGINFCAKM,proto3" json:"KFIGINFCAKM,omitempty"`
	FPHMEDIMIOB     string   `protobuf:"bytes,1274,opt,name=FPHMEDIMIOB,proto3" json:"FPHMEDIMIOB,omitempty"`
	GOELMBPGOFC     bool     `protobuf:"varint,1006,opt,name=GOELMBPGOFC,proto3" json:"GOELMBPGOFC,omitempty"`
	Unk1            bool     `protobuf:"varint,6,opt,name=unk1,proto3" json:"unk1,omitempty"`
	FEDENKOLCMC     string   `protobuf:"bytes,1209,opt,name=FEDENKOLCMC,proto3" json:"FEDENKOLCMC,omitempty"`
	IfixUrl         string   `protobuf:"bytes,1895,opt,name=ifix_url,json=ifixUrl,proto3" json:"ifix_url,omitempty"`
	CIKOEPDNJBA     string   `protobuf:"bytes,75,opt,name=CIKOEPDNJBA,proto3" json:"CIKOEPDNJBA,omitempty"`
	DIECBLIJNPO     string   `protobuf:"bytes,387,opt,name=DIECBLIJNPO,proto3" json:"DIECBLIJNPO,omitempty"`
	IfixVersion     string   `protobuf:"bytes,1690,opt,name=ifix_version,json=ifixVersion,proto3" json:"ifix_version,omitempty"`
	ExResourceUrl   string   `protobuf:"bytes,14,opt,name=ex_resource_url,json=exResourceUrl,proto3" json:"ex_resource_url,omitempty"`
	CIGLGHGDBHB     string   `protobuf:"bytes,462,opt,name=CIGLGHGDBHB,proto3" json:"CIGLGHGDBHB,omitempty"`
	GAHDEBODOPK     string   `protobuf:"bytes,1609,opt,name=GAHDEBODOPK,proto3" json:"GAHDEBODOPK,omitempty"`
	Ip              string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	AGMKEMOJHHJ     int64    `protobuf:"varint,11,opt,name=AGMKEMOJHHJ,proto3" json:"AGMKEMOJHHJ,omitempty"`
	KDGOBNKMNKC     uint32   `protobuf:"varint,9,opt,name=KDGOBNKMNKC,proto3" json:"KDGOBNKMNKC,omitempty"`
	Msg             string   `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg,omitempty"`
	GEBNKMIPOMM     bool     `protobuf:"varint,989,opt,name=GEBNKMIPOMM,proto3" json:"GEBNKMIPOMM,omitempty"`
	EBNBNCHNOFH     string   `protobuf:"bytes,4,opt,name=EBNBNCHNOFH,proto3" json:"EBNBNCHNOFH,omitempty"`
	GLCOMFKHDGH     string   `protobuf:"bytes,929,opt,name=GLCOMFKHDGH,proto3" json:"GLCOMFKHDGH,omitempty"`
	GECDDMDAGCL     string   `protobuf:"bytes,1226,opt,name=GECDDMDAGCL,proto3" json:"GECDDMDAGCL,omitempty"`
	FABJDGBFAFF     string   `protobuf:"bytes,512,opt,name=FABJDGBFAFF,proto3" json:"FABJDGBFAFF,omitempty"`
	EGNMPIPOGLI     bool     `protobuf:"varint,1359,opt,name=EGNMPIPOGLI,proto3" json:"EGNMPIPOGLI,omitempty"`
	JPGHOAAFODJ     string   `protobuf:"bytes,194,opt,name=JPGHOAAFODJ,proto3" json:"JPGHOAAFODJ,omitempty"`
	BKMDGECDJLJ     string   `protobuf:"bytes,2007,opt,name=BKMDGECDJLJ,proto3" json:"BKMDGECDJLJ,omitempty"`
	CCFBLAONHAF     uint32   `protobuf:"varint,350,opt,name=CCFBLAONHAF,proto3" json:"CCFBLAONHAF,omitempty"`
	FKFKCDJNHFL     bool     `protobuf:"varint,1401,opt,name=FKFKCDJNHFL,proto3" json:"FKFKCDJNHFL,omitempty"`
	MFENPHLEGHJ     bool     `protobuf:"varint,1537,opt,name=MFENPHLEGHJ,proto3" json:"MFENPHLEGHJ,omitempty"`
	Unk7            bool     `protobuf:"varint,889,opt,name=unk7,proto3" json:"unk7,omitempty"`
	FOPDCGLBOOI     string   `protobuf:"bytes,1470,opt,name=FOPDCGLBOOI,proto3" json:"FOPDCGLBOOI,omitempty"`
	GNFPFKJHIDJ     bool     `protobuf:"varint,1037,opt,name=GNFPFKJHIDJ,proto3" json:"GNFPFKJHIDJ,omitempty"`
	LHKJHDCIDND     string   `protobuf:"bytes,886,opt,name=LHKJHDCIDND,proto3" json:"LHKJHDCIDND,omitempty"`
	ALNIFJLCIJJ     string   `protobuf:"bytes,1783,opt,name=ALNIFJLCIJJ,proto3" json:"ALNIFJLCIJJ,omitempty"`
	AGEBFKODEAE     string   `protobuf:"bytes,616,opt,name=AGEBFKODEAE,proto3" json:"AGEBFKODEAE,omitempty"`
	ClientSecretKey string   `protobuf:"bytes,1597,opt,name=client_secret_key,json=clientSecretKey,proto3" json:"client_secret_key,omitempty"`
	Unk3            bool     `protobuf:"varint,82,opt,name=unk3,proto3" json:"unk3,omitempty"`
	Unk5            bool     `protobuf:"varint,452,opt,name=unk5,proto3" json:"unk5,omitempty"`
	Unk2            bool     `protobuf:"varint,13,opt,name=unk2,proto3" json:"unk2,omitempty"`
	NCDIJNHGELK     uint32   `protobuf:"varint,8,opt,name=NCDIJNHGELK,proto3" json:"NCDIJNHGELK,omitempty"`
	MdkResVersion   string   `protobuf:"bytes,1798,opt,name=mdk_res_version,json=mdkResVersion,proto3" json:"mdk_res_version,omitempty"`
	NMKHBELIAOK     string   `protobuf:"bytes,1613,opt,name=NMKHBELIAOK,proto3" json:"NMKHBELIAOK,omitempty"`
	ALEILJPDNFF     string   `protobuf:"bytes,692,opt,name=ALEILJPDNFF,proto3" json:"ALEILJPDNFF,omitempty"`
	OFGPAONJADM     int64    `protobuf:"varint,5,opt,name=OFGPAONJADM,proto3" json:"OFGPAONJADM,omitempty"`
	Port            uint32   `protobuf:"varint,12,opt,name=port,proto3" json:"port,omitempty"`
	AssetBundleUrl  string   `protobuf:"bytes,15,opt,name=asset_bundle_url,json=assetBundleUrl,proto3" json:"asset_bundle_url,omitempty"`
}

func (x *GateServer) Reset() {
	*x = GateServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GateServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateServer) ProtoMessage() {}

func (x *GateServer) ProtoReflect() protoreflect.Message {
	mi := &file_GateServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateServer.ProtoReflect.Descriptor instead.
func (*GateServer) Descriptor() ([]byte, []int) {
	return file_GateServer_proto_rawDescGZIP(), []int{0}
}

func (x *GateServer) GetJIHFOLPBFFL() string {
	if x != nil {
		return x.JIHFOLPBFFL
	}
	return ""
}

func (x *GateServer) GetAIFBLJAACHO() string {
	if x != nil {
		return x.AIFBLJAACHO
	}
	return ""
}

func (x *GateServer) GetUnk6() bool {
	if x != nil {
		return x.Unk6
	}
	return false
}

func (x *GateServer) GetNAHBLPGMKJC() string {
	if x != nil {
		return x.NAHBLPGMKJC
	}
	return ""
}

func (x *GateServer) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GateServer) GetAOEKIKFKMGA() bool {
	if x != nil {
		return x.AOEKIKFKMGA
	}
	return false
}

func (x *GateServer) GetCLJNMIKINLN() string {
	if x != nil {
		return x.CLJNMIKINLN
	}
	return ""
}

func (x *GateServer) GetGJMFLPHIDFD() []string {
	if x != nil {
		return x.GJMFLPHIDFD
	}
	return nil
}

func (x *GateServer) GetUnk4() bool {
	if x != nil {
		return x.Unk4
	}
	return false
}

func (x *GateServer) GetLuaUrl() string {
	if x != nil {
		return x.LuaUrl
	}
	return ""
}

func (x *GateServer) GetHJPAPNPPNNH() string {
	if x != nil {
		return x.HJPAPNPPNNH
	}
	return ""
}

func (x *GateServer) GetCDHOIAAONHD() bool {
	if x != nil {
		return x.CDHOIAAONHD
	}
	return false
}

func (x *GateServer) GetDHKIEGBPJML() string {
	if x != nil {
		return x.DHKIEGBPJML
	}
	return ""
}

func (x *GateServer) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *GateServer) GetOMDAJOMPIMG() string {
	if x != nil {
		return x.OMDAJOMPIMG
	}
	return ""
}

func (x *GateServer) GetKFIGINFCAKM() string {
	if x != nil {
		return x.KFIGINFCAKM
	}
	return ""
}

func (x *GateServer) GetFPHMEDIMIOB() string {
	if x != nil {
		return x.FPHMEDIMIOB
	}
	return ""
}

func (x *GateServer) GetGOELMBPGOFC() bool {
	if x != nil {
		return x.GOELMBPGOFC
	}
	return false
}

func (x *GateServer) GetUnk1() bool {
	if x != nil {
		return x.Unk1
	}
	return false
}

func (x *GateServer) GetFEDENKOLCMC() string {
	if x != nil {
		return x.FEDENKOLCMC
	}
	return ""
}

func (x *GateServer) GetIfixUrl() string {
	if x != nil {
		return x.IfixUrl
	}
	return ""
}

func (x *GateServer) GetCIKOEPDNJBA() string {
	if x != nil {
		return x.CIKOEPDNJBA
	}
	return ""
}

func (x *GateServer) GetDIECBLIJNPO() string {
	if x != nil {
		return x.DIECBLIJNPO
	}
	return ""
}

func (x *GateServer) GetIfixVersion() string {
	if x != nil {
		return x.IfixVersion
	}
	return ""
}

func (x *GateServer) GetExResourceUrl() string {
	if x != nil {
		return x.ExResourceUrl
	}
	return ""
}

func (x *GateServer) GetCIGLGHGDBHB() string {
	if x != nil {
		return x.CIGLGHGDBHB
	}
	return ""
}

func (x *GateServer) GetGAHDEBODOPK() string {
	if x != nil {
		return x.GAHDEBODOPK
	}
	return ""
}

func (x *GateServer) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *GateServer) GetAGMKEMOJHHJ() int64 {
	if x != nil {
		return x.AGMKEMOJHHJ
	}
	return 0
}

func (x *GateServer) GetKDGOBNKMNKC() uint32 {
	if x != nil {
		return x.KDGOBNKMNKC
	}
	return 0
}

func (x *GateServer) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GateServer) GetGEBNKMIPOMM() bool {
	if x != nil {
		return x.GEBNKMIPOMM
	}
	return false
}

func (x *GateServer) GetEBNBNCHNOFH() string {
	if x != nil {
		return x.EBNBNCHNOFH
	}
	return ""
}

func (x *GateServer) GetGLCOMFKHDGH() string {
	if x != nil {
		return x.GLCOMFKHDGH
	}
	return ""
}

func (x *GateServer) GetGECDDMDAGCL() string {
	if x != nil {
		return x.GECDDMDAGCL
	}
	return ""
}

func (x *GateServer) GetFABJDGBFAFF() string {
	if x != nil {
		return x.FABJDGBFAFF
	}
	return ""
}

func (x *GateServer) GetEGNMPIPOGLI() bool {
	if x != nil {
		return x.EGNMPIPOGLI
	}
	return false
}

func (x *GateServer) GetJPGHOAAFODJ() string {
	if x != nil {
		return x.JPGHOAAFODJ
	}
	return ""
}

func (x *GateServer) GetBKMDGECDJLJ() string {
	if x != nil {
		return x.BKMDGECDJLJ
	}
	return ""
}

func (x *GateServer) GetCCFBLAONHAF() uint32 {
	if x != nil {
		return x.CCFBLAONHAF
	}
	return 0
}

func (x *GateServer) GetFKFKCDJNHFL() bool {
	if x != nil {
		return x.FKFKCDJNHFL
	}
	return false
}

func (x *GateServer) GetMFENPHLEGHJ() bool {
	if x != nil {
		return x.MFENPHLEGHJ
	}
	return false
}

func (x *GateServer) GetUnk7() bool {
	if x != nil {
		return x.Unk7
	}
	return false
}

func (x *GateServer) GetFOPDCGLBOOI() string {
	if x != nil {
		return x.FOPDCGLBOOI
	}
	return ""
}

func (x *GateServer) GetGNFPFKJHIDJ() bool {
	if x != nil {
		return x.GNFPFKJHIDJ
	}
	return false
}

func (x *GateServer) GetLHKJHDCIDND() string {
	if x != nil {
		return x.LHKJHDCIDND
	}
	return ""
}

func (x *GateServer) GetALNIFJLCIJJ() string {
	if x != nil {
		return x.ALNIFJLCIJJ
	}
	return ""
}

func (x *GateServer) GetAGEBFKODEAE() string {
	if x != nil {
		return x.AGEBFKODEAE
	}
	return ""
}

func (x *GateServer) GetClientSecretKey() string {
	if x != nil {
		return x.ClientSecretKey
	}
	return ""
}

func (x *GateServer) GetUnk3() bool {
	if x != nil {
		return x.Unk3
	}
	return false
}

func (x *GateServer) GetUnk5() bool {
	if x != nil {
		return x.Unk5
	}
	return false
}

func (x *GateServer) GetUnk2() bool {
	if x != nil {
		return x.Unk2
	}
	return false
}

func (x *GateServer) GetNCDIJNHGELK() uint32 {
	if x != nil {
		return x.NCDIJNHGELK
	}
	return 0
}

func (x *GateServer) GetMdkResVersion() string {
	if x != nil {
		return x.MdkResVersion
	}
	return ""
}

func (x *GateServer) GetNMKHBELIAOK() string {
	if x != nil {
		return x.NMKHBELIAOK
	}
	return ""
}

func (x *GateServer) GetALEILJPDNFF() string {
	if x != nil {
		return x.ALEILJPDNFF
	}
	return ""
}

func (x *GateServer) GetOFGPAONJADM() int64 {
	if x != nil {
		return x.OFGPAONJADM
	}
	return 0
}

func (x *GateServer) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *GateServer) GetAssetBundleUrl() string {
	if x != nil {
		return x.AssetBundleUrl
	}
	return ""
}

var File_GateServer_proto protoreflect.FileDescriptor

var file_GateServer_proto_rawDesc = []byte{
	0x0a, 0x10, 0x47, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x80, 0x0f, 0x0a, 0x0a, 0x47, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0b, 0x4a, 0x49, 0x48, 0x46, 0x4f, 0x4c, 0x50, 0x42, 0x46, 0x46, 0x4c,
	0x18, 0xaa, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4a, 0x49, 0x48, 0x46, 0x4f, 0x4c, 0x50,
	0x42, 0x46, 0x46, 0x4c, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x49, 0x46, 0x42, 0x4c, 0x4a, 0x41, 0x41,
	0x43, 0x48, 0x4f, 0x18, 0xff, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x49, 0x46, 0x42,
	0x4c, 0x4a, 0x41, 0x41, 0x43, 0x48, 0x4f, 0x12, 0x13, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x36, 0x18,
	0xe1, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x36, 0x12, 0x21, 0x0a, 0x0b,
	0x4e, 0x41, 0x48, 0x42, 0x4c, 0x50, 0x47, 0x4d, 0x4b, 0x4a, 0x43, 0x18, 0xa5, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x4e, 0x41, 0x48, 0x42, 0x4c, 0x50, 0x47, 0x4d, 0x4b, 0x4a, 0x43, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x4f, 0x45,
	0x4b, 0x49, 0x4b, 0x46, 0x4b, 0x4d, 0x47, 0x41, 0x18, 0xbb, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x41, 0x4f, 0x45, 0x4b, 0x49, 0x4b, 0x46, 0x4b, 0x4d, 0x47, 0x41, 0x12, 0x21, 0x0a, 0x0b,
	0x43, 0x4c, 0x4a, 0x4e, 0x4d, 0x49, 0x4b, 0x49, 0x4e, 0x4c, 0x4e, 0x18, 0xad, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x43, 0x4c, 0x4a, 0x4e, 0x4d, 0x49, 0x4b, 0x49, 0x4e, 0x4c, 0x4e, 0x12,
	0x21, 0x0a, 0x0b, 0x47, 0x4a, 0x4d, 0x46, 0x4c, 0x50, 0x48, 0x49, 0x44, 0x46, 0x44, 0x18, 0xf2,
	0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x4a, 0x4d, 0x46, 0x4c, 0x50, 0x48, 0x49, 0x44,
	0x46, 0x44, 0x12, 0x13, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x34, 0x18, 0xd2, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x34, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x75, 0x61, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x75, 0x61, 0x55, 0x72, 0x6c,
	0x12, 0x21, 0x0a, 0x0b, 0x48, 0x4a, 0x50, 0x41, 0x50, 0x4e, 0x50, 0x50, 0x4e, 0x4e, 0x48, 0x18,
	0xa7, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x48, 0x4a, 0x50, 0x41, 0x50, 0x4e, 0x50, 0x50,
	0x4e, 0x4e, 0x48, 0x12, 0x21, 0x0a, 0x0b, 0x43, 0x44, 0x48, 0x4f, 0x49, 0x41, 0x41, 0x4f, 0x4e,
	0x48, 0x44, 0x18, 0xea, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x43, 0x44, 0x48, 0x4f, 0x49,
	0x41, 0x41, 0x4f, 0x4e, 0x48, 0x44, 0x12, 0x21, 0x0a, 0x0b, 0x44, 0x48, 0x4b, 0x49, 0x45, 0x47,
	0x42, 0x50, 0x4a, 0x4d, 0x4c, 0x18, 0x9a, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x48,
	0x4b, 0x49, 0x45, 0x47, 0x42, 0x50, 0x4a, 0x4d, 0x4c, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x4f, 0x4d,
	0x44, 0x41, 0x4a, 0x4f, 0x4d, 0x50, 0x49, 0x4d, 0x47, 0x18, 0xa4, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x4f, 0x4d, 0x44, 0x41, 0x4a, 0x4f, 0x4d, 0x50, 0x49, 0x4d, 0x47, 0x12, 0x21, 0x0a,
	0x0b, 0x4b, 0x46, 0x49, 0x47, 0x49, 0x4e, 0x46, 0x43, 0x41, 0x4b, 0x4d, 0x18, 0xad, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x4b, 0x46, 0x49, 0x47, 0x49, 0x4e, 0x46, 0x43, 0x41, 0x4b, 0x4d,
	0x12, 0x21, 0x0a, 0x0b, 0x46, 0x50, 0x48, 0x4d, 0x45, 0x44, 0x49, 0x4d, 0x49, 0x4f, 0x42, 0x18,
	0xfa, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x50, 0x48, 0x4d, 0x45, 0x44, 0x49, 0x4d,
	0x49, 0x4f, 0x42, 0x12, 0x21, 0x0a, 0x0b, 0x47, 0x4f, 0x45, 0x4c, 0x4d, 0x42, 0x50, 0x47, 0x4f,
	0x46, 0x43, 0x18, 0xee, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x47, 0x4f, 0x45, 0x4c, 0x4d,
	0x42, 0x50, 0x47, 0x4f, 0x46, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x31, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x31, 0x12, 0x21, 0x0a, 0x0b, 0x46, 0x45,
	0x44, 0x45, 0x4e, 0x4b, 0x4f, 0x4c, 0x43, 0x4d, 0x43, 0x18, 0xb9, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x46, 0x45, 0x44, 0x45, 0x4e, 0x4b, 0x4f, 0x4c, 0x43, 0x4d, 0x43, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x66, 0x69, 0x78, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0xe7, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x69, 0x66, 0x69, 0x78, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x49, 0x4b,
	0x4f, 0x45, 0x50, 0x44, 0x4e, 0x4a, 0x42, 0x41, 0x18, 0x4b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x43, 0x49, 0x4b, 0x4f, 0x45, 0x50, 0x44, 0x4e, 0x4a, 0x42, 0x41, 0x12, 0x21, 0x0a, 0x0b, 0x44,
	0x49, 0x45, 0x43, 0x42, 0x4c, 0x49, 0x4a, 0x4e, 0x50, 0x4f, 0x18, 0x83, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x49, 0x45, 0x43, 0x42, 0x4c, 0x49, 0x4a, 0x4e, 0x50, 0x4f, 0x12, 0x22,
	0x0a, 0x0c, 0x69, 0x66, 0x69, 0x78, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x9a,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x66, 0x69, 0x78, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0b, 0x43, 0x49,
	0x47, 0x4c, 0x47, 0x48, 0x47, 0x44, 0x42, 0x48, 0x42, 0x18, 0xce, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x43, 0x49, 0x47, 0x4c, 0x47, 0x48, 0x47, 0x44, 0x42, 0x48, 0x42, 0x12, 0x21, 0x0a,
	0x0b, 0x47, 0x41, 0x48, 0x44, 0x45, 0x42, 0x4f, 0x44, 0x4f, 0x50, 0x4b, 0x18, 0xc9, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x41, 0x48, 0x44, 0x45, 0x42, 0x4f, 0x44, 0x4f, 0x50, 0x4b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x20, 0x0a, 0x0b, 0x41, 0x47, 0x4d, 0x4b, 0x45, 0x4d, 0x4f, 0x4a, 0x48, 0x48, 0x4a, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x41, 0x47, 0x4d, 0x4b, 0x45, 0x4d, 0x4f, 0x4a, 0x48,
	0x48, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x44, 0x47, 0x4f, 0x42, 0x4e, 0x4b, 0x4d, 0x4e, 0x4b,
	0x43, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x44, 0x47, 0x4f, 0x42, 0x4e, 0x4b,
	0x4d, 0x4e, 0x4b, 0x43, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x21, 0x0a, 0x0b, 0x47, 0x45, 0x42, 0x4e, 0x4b, 0x4d,
	0x49, 0x50, 0x4f, 0x4d, 0x4d, 0x18, 0xdd, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x47, 0x45,
	0x42, 0x4e, 0x4b, 0x4d, 0x49, 0x50, 0x4f, 0x4d, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x42, 0x4e,
	0x42, 0x4e, 0x43, 0x48, 0x4e, 0x4f, 0x46, 0x48, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x45, 0x42, 0x4e, 0x42, 0x4e, 0x43, 0x48, 0x4e, 0x4f, 0x46, 0x48, 0x12, 0x21, 0x0a, 0x0b, 0x47,
	0x4c, 0x43, 0x4f, 0x4d, 0x46, 0x4b, 0x48, 0x44, 0x47, 0x48, 0x18, 0xa1, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x47, 0x4c, 0x43, 0x4f, 0x4d, 0x46, 0x4b, 0x48, 0x44, 0x47, 0x48, 0x12, 0x21,
	0x0a, 0x0b, 0x47, 0x45, 0x43, 0x44, 0x44, 0x4d, 0x44, 0x41, 0x47, 0x43, 0x4c, 0x18, 0xca, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x45, 0x43, 0x44, 0x44, 0x4d, 0x44, 0x41, 0x47, 0x43,
	0x4c, 0x12, 0x21, 0x0a, 0x0b, 0x46, 0x41, 0x42, 0x4a, 0x44, 0x47, 0x42, 0x46, 0x41, 0x46, 0x46,
	0x18, 0x80, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x41, 0x42, 0x4a, 0x44, 0x47, 0x42,
	0x46, 0x41, 0x46, 0x46, 0x12, 0x21, 0x0a, 0x0b, 0x45, 0x47, 0x4e, 0x4d, 0x50, 0x49, 0x50, 0x4f,
	0x47, 0x4c, 0x49, 0x18, 0xcf, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x45, 0x47, 0x4e, 0x4d,
	0x50, 0x49, 0x50, 0x4f, 0x47, 0x4c, 0x49, 0x12, 0x21, 0x0a, 0x0b, 0x4a, 0x50, 0x47, 0x48, 0x4f,
	0x41, 0x41, 0x46, 0x4f, 0x44, 0x4a, 0x18, 0xc2, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4a,
	0x50, 0x47, 0x48, 0x4f, 0x41, 0x41, 0x46, 0x4f, 0x44, 0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x42, 0x4b,
	0x4d, 0x44, 0x47, 0x45, 0x43, 0x44, 0x4a, 0x4c, 0x4a, 0x18, 0xd7, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x42, 0x4b, 0x4d, 0x44, 0x47, 0x45, 0x43, 0x44, 0x4a, 0x4c, 0x4a, 0x12, 0x21, 0x0a,
	0x0b, 0x43, 0x43, 0x46, 0x42, 0x4c, 0x41, 0x4f, 0x4e, 0x48, 0x41, 0x46, 0x18, 0xde, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x43, 0x46, 0x42, 0x4c, 0x41, 0x4f, 0x4e, 0x48, 0x41, 0x46,
	0x12, 0x21, 0x0a, 0x0b, 0x46, 0x4b, 0x46, 0x4b, 0x43, 0x44, 0x4a, 0x4e, 0x48, 0x46, 0x4c, 0x18,
	0xf9, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x46, 0x4b, 0x46, 0x4b, 0x43, 0x44, 0x4a, 0x4e,
	0x48, 0x46, 0x4c, 0x12, 0x21, 0x0a, 0x0b, 0x4d, 0x46, 0x45, 0x4e, 0x50, 0x48, 0x4c, 0x45, 0x47,
	0x48, 0x4a, 0x18, 0x81, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4d, 0x46, 0x45, 0x4e, 0x50,
	0x48, 0x4c, 0x45, 0x47, 0x48, 0x4a, 0x12, 0x13, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x37, 0x18, 0xf9,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x37, 0x12, 0x21, 0x0a, 0x0b, 0x46,
	0x4f, 0x50, 0x44, 0x43, 0x47, 0x4c, 0x42, 0x4f, 0x4f, 0x49, 0x18, 0xbe, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x46, 0x4f, 0x50, 0x44, 0x43, 0x47, 0x4c, 0x42, 0x4f, 0x4f, 0x49, 0x12, 0x21,
	0x0a, 0x0b, 0x47, 0x4e, 0x46, 0x50, 0x46, 0x4b, 0x4a, 0x48, 0x49, 0x44, 0x4a, 0x18, 0x8d, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x47, 0x4e, 0x46, 0x50, 0x46, 0x4b, 0x4a, 0x48, 0x49, 0x44,
	0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x4c, 0x48, 0x4b, 0x4a, 0x48, 0x44, 0x43, 0x49, 0x44, 0x4e, 0x44,
	0x18, 0xf6, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x48, 0x4b, 0x4a, 0x48, 0x44, 0x43,
	0x49, 0x44, 0x4e, 0x44, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x4c, 0x4e, 0x49, 0x46, 0x4a, 0x4c, 0x43,
	0x49, 0x4a, 0x4a, 0x18, 0xf7, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x4c, 0x4e, 0x49,
	0x46, 0x4a, 0x4c, 0x43, 0x49, 0x4a, 0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x47, 0x45, 0x42, 0x46,
	0x4b, 0x4f, 0x44, 0x45, 0x41, 0x45, 0x18, 0xe8, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41,
	0x47, 0x45, 0x42, 0x46, 0x4b, 0x4f, 0x44, 0x45, 0x41, 0x45, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0xbd, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x33, 0x18,
	0x52, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x33, 0x12, 0x13, 0x0a, 0x04, 0x75,
	0x6e, 0x6b, 0x35, 0x18, 0xc4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x6e, 0x6b, 0x35,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x6b, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x75, 0x6e, 0x6b, 0x32, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x43, 0x44, 0x49, 0x4a, 0x4e, 0x48, 0x47,
	0x45, 0x4c, 0x4b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x43, 0x44, 0x49, 0x4a,
	0x4e, 0x48, 0x47, 0x45, 0x4c, 0x4b, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x64, 0x6b, 0x5f, 0x72, 0x65,
	0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x86, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6d, 0x64, 0x6b, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0b, 0x4e, 0x4d, 0x4b, 0x48, 0x42, 0x45, 0x4c, 0x49, 0x41, 0x4f, 0x4b, 0x18, 0xcd,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x4d, 0x4b, 0x48, 0x42, 0x45, 0x4c, 0x49, 0x41,
	0x4f, 0x4b, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x4c, 0x45, 0x49, 0x4c, 0x4a, 0x50, 0x44, 0x4e, 0x46,
	0x46, 0x18, 0xb4, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x4c, 0x45, 0x49, 0x4c, 0x4a,
	0x50, 0x44, 0x4e, 0x46, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x46, 0x47, 0x50, 0x41, 0x4f, 0x4e,
	0x4a, 0x41, 0x44, 0x4d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x4f, 0x46, 0x47, 0x50,
	0x41, 0x4f, 0x4e, 0x4a, 0x41, 0x44, 0x4d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GateServer_proto_rawDescOnce sync.Once
	file_GateServer_proto_rawDescData = file_GateServer_proto_rawDesc
)

func file_GateServer_proto_rawDescGZIP() []byte {
	file_GateServer_proto_rawDescOnce.Do(func() {
		file_GateServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_GateServer_proto_rawDescData)
	})
	return file_GateServer_proto_rawDescData
}

var file_GateServer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GateServer_proto_goTypes = []interface{}{
	(*GateServer)(nil), // 0: GateServer
}
var file_GateServer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GateServer_proto_init() }
func file_GateServer_proto_init() {
	if File_GateServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GateServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateServer); i {
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
			RawDescriptor: file_GateServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GateServer_proto_goTypes,
		DependencyIndexes: file_GateServer_proto_depIdxs,
		MessageInfos:      file_GateServer_proto_msgTypes,
	}.Build()
	File_GateServer_proto = out.File
	file_GateServer_proto_rawDesc = nil
	file_GateServer_proto_goTypes = nil
	file_GateServer_proto_depIdxs = nil
}
