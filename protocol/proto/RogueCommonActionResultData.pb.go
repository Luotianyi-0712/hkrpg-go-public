// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueCommonActionResultData.proto

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

type RogueCommonActionResultData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ResultData:
	//
	//	*RogueCommonActionResultData_GetItemList
	//	*RogueCommonActionResultData_RemoveItemList
	//	*RogueCommonActionResultData_GetBuffList
	//	*RogueCommonActionResultData_RemoveBuffList
	//	*RogueCommonActionResultData_GetMiracleList
	//	*RogueCommonActionResultData_KDGKHHFHFBJ
	//	*RogueCommonActionResultData_DHHFCKKDFNJ
	//	*RogueCommonActionResultData_ONGBGHBJBBF
	//	*RogueCommonActionResultData_EKMBBHEHEPC
	//	*RogueCommonActionResultData_HEIJFDDMMFJ
	//	*RogueCommonActionResultData_MICOCBNMEAJ
	//	*RogueCommonActionResultData_GetFormulaList
	//	*RogueCommonActionResultData_NKMIJHMBCJE
	//	*RogueCommonActionResultData_GACMJKONLPO
	//	*RogueCommonActionResultData_KDHCEKHEIEK
	//	*RogueCommonActionResultData_OCFHDHJHOLH
	//	*RogueCommonActionResultData_GetKeywordList
	//	*RogueCommonActionResultData_JBPJOAHOBGM
	ResultData isRogueCommonActionResultData_ResultData `protobuf_oneof:"result_data"`
}

func (x *RogueCommonActionResultData) Reset() {
	*x = RogueCommonActionResultData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueCommonActionResultData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueCommonActionResultData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueCommonActionResultData) ProtoMessage() {}

func (x *RogueCommonActionResultData) ProtoReflect() protoreflect.Message {
	mi := &file_RogueCommonActionResultData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueCommonActionResultData.ProtoReflect.Descriptor instead.
func (*RogueCommonActionResultData) Descriptor() ([]byte, []int) {
	return file_RogueCommonActionResultData_proto_rawDescGZIP(), []int{0}
}

func (m *RogueCommonActionResultData) GetResultData() isRogueCommonActionResultData_ResultData {
	if m != nil {
		return m.ResultData
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGetItemList() *RogueCommonMoney {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GetItemList); ok {
		return x.GetItemList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetRemoveItemList() *RogueCommonMoney {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_RemoveItemList); ok {
		return x.RemoveItemList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGetBuffList() *RogueCommonBuff {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GetBuffList); ok {
		return x.GetBuffList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetRemoveBuffList() *RogueCommonBuff {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_RemoveBuffList); ok {
		return x.RemoveBuffList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGetMiracleList() *RogueCommonMiracle {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GetMiracleList); ok {
		return x.GetMiracleList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetKDGKHHFHFBJ() *EHLDCELKDKB {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_KDGKHHFHFBJ); ok {
		return x.KDGKHHFHFBJ
	}
	return nil
}

func (x *RogueCommonActionResultData) GetDHHFCKKDFNJ() *AFKGDNPFJAA {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_DHHFCKKDFNJ); ok {
		return x.DHHFCKKDFNJ
	}
	return nil
}

func (x *RogueCommonActionResultData) GetONGBGHBJBBF() *FGOIJBNMIEP {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_ONGBGHBJBBF); ok {
		return x.ONGBGHBJBBF
	}
	return nil
}

func (x *RogueCommonActionResultData) GetEKMBBHEHEPC() *FBIOOAHFDEM {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_EKMBBHEHEPC); ok {
		return x.EKMBBHEHEPC
	}
	return nil
}

func (x *RogueCommonActionResultData) GetHEIJFDDMMFJ() *LPPPBHGHJDF {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_HEIJFDDMMFJ); ok {
		return x.HEIJFDDMMFJ
	}
	return nil
}

func (x *RogueCommonActionResultData) GetMICOCBNMEAJ() *KOMCCJGJHNM {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_MICOCBNMEAJ); ok {
		return x.MICOCBNMEAJ
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGetFormulaList() *RogueCommonFormula {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GetFormulaList); ok {
		return x.GetFormulaList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetNKMIJHMBCJE() *BANENCLAHFH {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_NKMIJHMBCJE); ok {
		return x.NKMIJHMBCJE
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGACMJKONLPO() *JCEGOHOHIGK {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GACMJKONLPO); ok {
		return x.GACMJKONLPO
	}
	return nil
}

func (x *RogueCommonActionResultData) GetKDHCEKHEIEK() *FCKINKMKFOJ {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_KDHCEKHEIEK); ok {
		return x.KDHCEKHEIEK
	}
	return nil
}

func (x *RogueCommonActionResultData) GetOCFHDHJHOLH() *MEPPCFCOCMC {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_OCFHDHJHOLH); ok {
		return x.OCFHDHJHOLH
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGetKeywordList() *RogueCommonKeyword {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GetKeywordList); ok {
		return x.GetKeywordList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetJBPJOAHOBGM() *DCCENMIJJDG {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_JBPJOAHOBGM); ok {
		return x.JBPJOAHOBGM
	}
	return nil
}

type isRogueCommonActionResultData_ResultData interface {
	isRogueCommonActionResultData_ResultData()
}

type RogueCommonActionResultData_GetItemList struct {
	GetItemList *RogueCommonMoney `protobuf:"bytes,6,opt,name=get_item_list,json=getItemList,proto3,oneof"`
}

type RogueCommonActionResultData_RemoveItemList struct {
	RemoveItemList *RogueCommonMoney `protobuf:"bytes,2,opt,name=remove_item_list,json=removeItemList,proto3,oneof"`
}

type RogueCommonActionResultData_GetBuffList struct {
	GetBuffList *RogueCommonBuff `protobuf:"bytes,1884,opt,name=get_buff_list,json=getBuffList,proto3,oneof"`
}

type RogueCommonActionResultData_RemoveBuffList struct {
	RemoveBuffList *RogueCommonBuff `protobuf:"bytes,1544,opt,name=remove_buff_list,json=removeBuffList,proto3,oneof"`
}

type RogueCommonActionResultData_GetMiracleList struct {
	GetMiracleList *RogueCommonMiracle `protobuf:"bytes,1263,opt,name=get_miracle_list,json=getMiracleList,proto3,oneof"`
}

type RogueCommonActionResultData_KDGKHHFHFBJ struct {
	KDGKHHFHFBJ *EHLDCELKDKB `protobuf:"bytes,1017,opt,name=KDGKHHFHFBJ,proto3,oneof"`
}

type RogueCommonActionResultData_DHHFCKKDFNJ struct {
	DHHFCKKDFNJ *AFKGDNPFJAA `protobuf:"bytes,861,opt,name=DHHFCKKDFNJ,proto3,oneof"`
}

type RogueCommonActionResultData_ONGBGHBJBBF struct {
	ONGBGHBJBBF *FGOIJBNMIEP `protobuf:"bytes,217,opt,name=ONGBGHBJBBF,proto3,oneof"`
}

type RogueCommonActionResultData_EKMBBHEHEPC struct {
	EKMBBHEHEPC *FBIOOAHFDEM `protobuf:"bytes,1386,opt,name=EKMBBHEHEPC,proto3,oneof"`
}

type RogueCommonActionResultData_HEIJFDDMMFJ struct {
	HEIJFDDMMFJ *LPPPBHGHJDF `protobuf:"bytes,265,opt,name=HEIJFDDMMFJ,proto3,oneof"`
}

type RogueCommonActionResultData_MICOCBNMEAJ struct {
	MICOCBNMEAJ *KOMCCJGJHNM `protobuf:"bytes,330,opt,name=MICOCBNMEAJ,proto3,oneof"`
}

type RogueCommonActionResultData_GetFormulaList struct {
	GetFormulaList *RogueCommonFormula `protobuf:"bytes,776,opt,name=get_formula_list,json=getFormulaList,proto3,oneof"`
}

type RogueCommonActionResultData_NKMIJHMBCJE struct {
	NKMIJHMBCJE *BANENCLAHFH `protobuf:"bytes,2044,opt,name=NKMIJHMBCJE,proto3,oneof"`
}

type RogueCommonActionResultData_GACMJKONLPO struct {
	GACMJKONLPO *JCEGOHOHIGK `protobuf:"bytes,1519,opt,name=GACMJKONLPO,proto3,oneof"`
}

type RogueCommonActionResultData_KDHCEKHEIEK struct {
	KDHCEKHEIEK *FCKINKMKFOJ `protobuf:"bytes,1280,opt,name=KDHCEKHEIEK,proto3,oneof"`
}

type RogueCommonActionResultData_OCFHDHJHOLH struct {
	OCFHDHJHOLH *MEPPCFCOCMC `protobuf:"bytes,1843,opt,name=OCFHDHJHOLH,proto3,oneof"`
}

type RogueCommonActionResultData_GetKeywordList struct {
	GetKeywordList *RogueCommonKeyword `protobuf:"bytes,1434,opt,name=get_keyword_list,json=getKeywordList,proto3,oneof"`
}

type RogueCommonActionResultData_JBPJOAHOBGM struct {
	JBPJOAHOBGM *DCCENMIJJDG `protobuf:"bytes,1346,opt,name=JBPJOAHOBGM,proto3,oneof"`
}

func (*RogueCommonActionResultData_GetItemList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_RemoveItemList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_GetBuffList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_RemoveBuffList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_GetMiracleList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_KDGKHHFHFBJ) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_DHHFCKKDFNJ) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_ONGBGHBJBBF) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_EKMBBHEHEPC) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_HEIJFDDMMFJ) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_MICOCBNMEAJ) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_GetFormulaList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_NKMIJHMBCJE) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_GACMJKONLPO) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_KDHCEKHEIEK) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_OCFHDHJHOLH) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_GetKeywordList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_JBPJOAHOBGM) isRogueCommonActionResultData_ResultData() {}

var File_RogueCommonActionResultData_proto protoreflect.FileDescriptor

var file_RogueCommonActionResultData_proto_rawDesc = []byte{
	0x0a, 0x21, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x47, 0x4f, 0x49, 0x4a, 0x42, 0x4e, 0x4d, 0x49, 0x45, 0x50,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x46, 0x43, 0x4b, 0x49, 0x4e, 0x4b, 0x4d, 0x4b, 0x46, 0x4f, 0x4a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x43, 0x43, 0x45, 0x4e, 0x4d, 0x49, 0x4a, 0x4a, 0x44, 0x47,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x42, 0x49, 0x4f, 0x4f, 0x41, 0x48, 0x46,
	0x44, 0x45, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x41, 0x4e, 0x45, 0x4e,
	0x43, 0x4c, 0x41, 0x48, 0x46, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x50,
	0x50, 0x50, 0x42, 0x48, 0x47, 0x48, 0x4a, 0x44, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x45, 0x50, 0x50, 0x43,
	0x46, 0x43, 0x4f, 0x43, 0x4d, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4f,
	0x4d, 0x43, 0x43, 0x4a, 0x47, 0x4a, 0x48, 0x4e, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4a, 0x43, 0x45, 0x47, 0x4f, 0x48, 0x4f, 0x48, 0x49, 0x47, 0x4b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x48, 0x4c, 0x44, 0x43, 0x45, 0x4c, 0x4b, 0x44, 0x4b, 0x42,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x46, 0x4b, 0x47, 0x44, 0x4e, 0x50, 0x46,
	0x4a, 0x41, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x93, 0x08, 0x0a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x37, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x67,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x10, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x67, 0x65, 0x74,
	0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xdc, 0x0e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42,
	0x75, 0x66, 0x66, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x65, 0x74, 0x42, 0x75, 0x66, 0x66, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x62, 0x75, 0x66,
	0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x88, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x48,
	0x00, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xef, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x48, 0x00, 0x52, 0x0e, 0x67, 0x65, 0x74, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x4b, 0x44, 0x47, 0x4b, 0x48, 0x48, 0x46, 0x48, 0x46,
	0x42, 0x4a, 0x18, 0xf9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x48, 0x4c, 0x44,
	0x43, 0x45, 0x4c, 0x4b, 0x44, 0x4b, 0x42, 0x48, 0x00, 0x52, 0x0b, 0x4b, 0x44, 0x47, 0x4b, 0x48,
	0x48, 0x46, 0x48, 0x46, 0x42, 0x4a, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x48, 0x48, 0x46, 0x43, 0x4b,
	0x4b, 0x44, 0x46, 0x4e, 0x4a, 0x18, 0xdd, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41,
	0x46, 0x4b, 0x47, 0x44, 0x4e, 0x50, 0x46, 0x4a, 0x41, 0x41, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x48,
	0x48, 0x46, 0x43, 0x4b, 0x4b, 0x44, 0x46, 0x4e, 0x4a, 0x12, 0x31, 0x0a, 0x0b, 0x4f, 0x4e, 0x47,
	0x42, 0x47, 0x48, 0x42, 0x4a, 0x42, 0x42, 0x46, 0x18, 0xd9, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x46, 0x47, 0x4f, 0x49, 0x4a, 0x42, 0x4e, 0x4d, 0x49, 0x45, 0x50, 0x48, 0x00, 0x52,
	0x0b, 0x4f, 0x4e, 0x47, 0x42, 0x47, 0x48, 0x42, 0x4a, 0x42, 0x42, 0x46, 0x12, 0x31, 0x0a, 0x0b,
	0x45, 0x4b, 0x4d, 0x42, 0x42, 0x48, 0x45, 0x48, 0x45, 0x50, 0x43, 0x18, 0xea, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x42, 0x49, 0x4f, 0x4f, 0x41, 0x48, 0x46, 0x44, 0x45, 0x4d,
	0x48, 0x00, 0x52, 0x0b, 0x45, 0x4b, 0x4d, 0x42, 0x42, 0x48, 0x45, 0x48, 0x45, 0x50, 0x43, 0x12,
	0x31, 0x0a, 0x0b, 0x48, 0x45, 0x49, 0x4a, 0x46, 0x44, 0x44, 0x4d, 0x4d, 0x46, 0x4a, 0x18, 0x89,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x50, 0x50, 0x50, 0x42, 0x48, 0x47, 0x48,
	0x4a, 0x44, 0x46, 0x48, 0x00, 0x52, 0x0b, 0x48, 0x45, 0x49, 0x4a, 0x46, 0x44, 0x44, 0x4d, 0x4d,
	0x46, 0x4a, 0x12, 0x31, 0x0a, 0x0b, 0x4d, 0x49, 0x43, 0x4f, 0x43, 0x42, 0x4e, 0x4d, 0x45, 0x41,
	0x4a, 0x18, 0xca, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4f, 0x4d, 0x43, 0x43,
	0x4a, 0x47, 0x4a, 0x48, 0x4e, 0x4d, 0x48, 0x00, 0x52, 0x0b, 0x4d, 0x49, 0x43, 0x4f, 0x43, 0x42,
	0x4e, 0x4d, 0x45, 0x41, 0x4a, 0x12, 0x40, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x72,
	0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x88, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x46, 0x6f,
	0x72, 0x6d, 0x75, 0x6c, 0x61, 0x48, 0x00, 0x52, 0x0e, 0x67, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d,
	0x75, 0x6c, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x4e, 0x4b, 0x4d, 0x49, 0x4a,
	0x48, 0x4d, 0x42, 0x43, 0x4a, 0x45, 0x18, 0xfc, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x42, 0x41, 0x4e, 0x45, 0x4e, 0x43, 0x4c, 0x41, 0x48, 0x46, 0x48, 0x48, 0x00, 0x52, 0x0b, 0x4e,
	0x4b, 0x4d, 0x49, 0x4a, 0x48, 0x4d, 0x42, 0x43, 0x4a, 0x45, 0x12, 0x31, 0x0a, 0x0b, 0x47, 0x41,
	0x43, 0x4d, 0x4a, 0x4b, 0x4f, 0x4e, 0x4c, 0x50, 0x4f, 0x18, 0xef, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4a, 0x43, 0x45, 0x47, 0x4f, 0x48, 0x4f, 0x48, 0x49, 0x47, 0x4b, 0x48, 0x00,
	0x52, 0x0b, 0x47, 0x41, 0x43, 0x4d, 0x4a, 0x4b, 0x4f, 0x4e, 0x4c, 0x50, 0x4f, 0x12, 0x31, 0x0a,
	0x0b, 0x4b, 0x44, 0x48, 0x43, 0x45, 0x4b, 0x48, 0x45, 0x49, 0x45, 0x4b, 0x18, 0x80, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x46, 0x43, 0x4b, 0x49, 0x4e, 0x4b, 0x4d, 0x4b, 0x46, 0x4f,
	0x4a, 0x48, 0x00, 0x52, 0x0b, 0x4b, 0x44, 0x48, 0x43, 0x45, 0x4b, 0x48, 0x45, 0x49, 0x45, 0x4b,
	0x12, 0x31, 0x0a, 0x0b, 0x4f, 0x43, 0x46, 0x48, 0x44, 0x48, 0x4a, 0x48, 0x4f, 0x4c, 0x48, 0x18,
	0xb3, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x45, 0x50, 0x50, 0x43, 0x46, 0x43,
	0x4f, 0x43, 0x4d, 0x43, 0x48, 0x00, 0x52, 0x0b, 0x4f, 0x43, 0x46, 0x48, 0x44, 0x48, 0x4a, 0x48,
	0x4f, 0x4c, 0x48, 0x12, 0x40, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x9a, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x0e, 0x67, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x4a, 0x42, 0x50, 0x4a, 0x4f, 0x41, 0x48,
	0x4f, 0x42, 0x47, 0x4d, 0x18, 0xc2, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x43,
	0x43, 0x45, 0x4e, 0x4d, 0x49, 0x4a, 0x4a, 0x44, 0x47, 0x48, 0x00, 0x52, 0x0b, 0x4a, 0x42, 0x50,
	0x4a, 0x4f, 0x41, 0x48, 0x4f, 0x42, 0x47, 0x4d, 0x42, 0x0d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueCommonActionResultData_proto_rawDescOnce sync.Once
	file_RogueCommonActionResultData_proto_rawDescData = file_RogueCommonActionResultData_proto_rawDesc
)

func file_RogueCommonActionResultData_proto_rawDescGZIP() []byte {
	file_RogueCommonActionResultData_proto_rawDescOnce.Do(func() {
		file_RogueCommonActionResultData_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueCommonActionResultData_proto_rawDescData)
	})
	return file_RogueCommonActionResultData_proto_rawDescData
}

var file_RogueCommonActionResultData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueCommonActionResultData_proto_goTypes = []interface{}{
	(*RogueCommonActionResultData)(nil), // 0: RogueCommonActionResultData
	(*RogueCommonMoney)(nil),            // 1: RogueCommonMoney
	(*RogueCommonBuff)(nil),             // 2: RogueCommonBuff
	(*RogueCommonMiracle)(nil),          // 3: RogueCommonMiracle
	(*EHLDCELKDKB)(nil),                 // 4: EHLDCELKDKB
	(*AFKGDNPFJAA)(nil),                 // 5: AFKGDNPFJAA
	(*FGOIJBNMIEP)(nil),                 // 6: FGOIJBNMIEP
	(*FBIOOAHFDEM)(nil),                 // 7: FBIOOAHFDEM
	(*LPPPBHGHJDF)(nil),                 // 8: LPPPBHGHJDF
	(*KOMCCJGJHNM)(nil),                 // 9: KOMCCJGJHNM
	(*RogueCommonFormula)(nil),          // 10: RogueCommonFormula
	(*BANENCLAHFH)(nil),                 // 11: BANENCLAHFH
	(*JCEGOHOHIGK)(nil),                 // 12: JCEGOHOHIGK
	(*FCKINKMKFOJ)(nil),                 // 13: FCKINKMKFOJ
	(*MEPPCFCOCMC)(nil),                 // 14: MEPPCFCOCMC
	(*RogueCommonKeyword)(nil),          // 15: RogueCommonKeyword
	(*DCCENMIJJDG)(nil),                 // 16: DCCENMIJJDG
}
var file_RogueCommonActionResultData_proto_depIdxs = []int32{
	1,  // 0: RogueCommonActionResultData.get_item_list:type_name -> RogueCommonMoney
	1,  // 1: RogueCommonActionResultData.remove_item_list:type_name -> RogueCommonMoney
	2,  // 2: RogueCommonActionResultData.get_buff_list:type_name -> RogueCommonBuff
	2,  // 3: RogueCommonActionResultData.remove_buff_list:type_name -> RogueCommonBuff
	3,  // 4: RogueCommonActionResultData.get_miracle_list:type_name -> RogueCommonMiracle
	4,  // 5: RogueCommonActionResultData.KDGKHHFHFBJ:type_name -> EHLDCELKDKB
	5,  // 6: RogueCommonActionResultData.DHHFCKKDFNJ:type_name -> AFKGDNPFJAA
	6,  // 7: RogueCommonActionResultData.ONGBGHBJBBF:type_name -> FGOIJBNMIEP
	7,  // 8: RogueCommonActionResultData.EKMBBHEHEPC:type_name -> FBIOOAHFDEM
	8,  // 9: RogueCommonActionResultData.HEIJFDDMMFJ:type_name -> LPPPBHGHJDF
	9,  // 10: RogueCommonActionResultData.MICOCBNMEAJ:type_name -> KOMCCJGJHNM
	10, // 11: RogueCommonActionResultData.get_formula_list:type_name -> RogueCommonFormula
	11, // 12: RogueCommonActionResultData.NKMIJHMBCJE:type_name -> BANENCLAHFH
	12, // 13: RogueCommonActionResultData.GACMJKONLPO:type_name -> JCEGOHOHIGK
	13, // 14: RogueCommonActionResultData.KDHCEKHEIEK:type_name -> FCKINKMKFOJ
	14, // 15: RogueCommonActionResultData.OCFHDHJHOLH:type_name -> MEPPCFCOCMC
	15, // 16: RogueCommonActionResultData.get_keyword_list:type_name -> RogueCommonKeyword
	16, // 17: RogueCommonActionResultData.JBPJOAHOBGM:type_name -> DCCENMIJJDG
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_RogueCommonActionResultData_proto_init() }
func file_RogueCommonActionResultData_proto_init() {
	if File_RogueCommonActionResultData_proto != nil {
		return
	}
	file_FGOIJBNMIEP_proto_init()
	file_RogueCommonFormula_proto_init()
	file_FCKINKMKFOJ_proto_init()
	file_DCCENMIJJDG_proto_init()
	file_FBIOOAHFDEM_proto_init()
	file_BANENCLAHFH_proto_init()
	file_LPPPBHGHJDF_proto_init()
	file_RogueCommonMiracle_proto_init()
	file_MEPPCFCOCMC_proto_init()
	file_KOMCCJGJHNM_proto_init()
	file_JCEGOHOHIGK_proto_init()
	file_RogueCommonKeyword_proto_init()
	file_RogueCommonBuff_proto_init()
	file_EHLDCELKDKB_proto_init()
	file_AFKGDNPFJAA_proto_init()
	file_RogueCommonMoney_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueCommonActionResultData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueCommonActionResultData); i {
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
	file_RogueCommonActionResultData_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RogueCommonActionResultData_GetItemList)(nil),
		(*RogueCommonActionResultData_RemoveItemList)(nil),
		(*RogueCommonActionResultData_GetBuffList)(nil),
		(*RogueCommonActionResultData_RemoveBuffList)(nil),
		(*RogueCommonActionResultData_GetMiracleList)(nil),
		(*RogueCommonActionResultData_KDGKHHFHFBJ)(nil),
		(*RogueCommonActionResultData_DHHFCKKDFNJ)(nil),
		(*RogueCommonActionResultData_ONGBGHBJBBF)(nil),
		(*RogueCommonActionResultData_EKMBBHEHEPC)(nil),
		(*RogueCommonActionResultData_HEIJFDDMMFJ)(nil),
		(*RogueCommonActionResultData_MICOCBNMEAJ)(nil),
		(*RogueCommonActionResultData_GetFormulaList)(nil),
		(*RogueCommonActionResultData_NKMIJHMBCJE)(nil),
		(*RogueCommonActionResultData_GACMJKONLPO)(nil),
		(*RogueCommonActionResultData_KDHCEKHEIEK)(nil),
		(*RogueCommonActionResultData_OCFHDHJHOLH)(nil),
		(*RogueCommonActionResultData_GetKeywordList)(nil),
		(*RogueCommonActionResultData_JBPJOAHOBGM)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueCommonActionResultData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueCommonActionResultData_proto_goTypes,
		DependencyIndexes: file_RogueCommonActionResultData_proto_depIdxs,
		MessageInfos:      file_RogueCommonActionResultData_proto_msgTypes,
	}.Build()
	File_RogueCommonActionResultData_proto = out.File
	file_RogueCommonActionResultData_proto_rawDesc = nil
	file_RogueCommonActionResultData_proto_goTypes = nil
	file_RogueCommonActionResultData_proto_depIdxs = nil
}
