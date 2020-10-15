// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.2
// source: proto/favorite/repertory.proto

package favorite

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RepertoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Created  int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated  int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name     string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Owner    string   `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark   string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	Creator  string   `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator string   `protobuf:"bytes,9,opt,name=operator,proto3" json:"operator,omitempty"`
	Bags     []string `protobuf:"bytes,10,rep,name=bags,proto3" json:"bags,omitempty"`
}

func (x *RepertoryInfo) Reset() {
	*x = RepertoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_favorite_repertory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepertoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepertoryInfo) ProtoMessage() {}

func (x *RepertoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_favorite_repertory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepertoryInfo.ProtoReflect.Descriptor instead.
func (*RepertoryInfo) Descriptor() ([]byte, []int) {
	return file_proto_favorite_repertory_proto_rawDescGZIP(), []int{0}
}

func (x *RepertoryInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RepertoryInfo) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *RepertoryInfo) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *RepertoryInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RepertoryInfo) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *RepertoryInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *RepertoryInfo) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *RepertoryInfo) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *RepertoryInfo) GetBags() []string {
	if x != nil {
		return x.Bags
	}
	return nil
}

type ReplyRepertoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info   *RepertoryInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status *ReplyStatus   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReplyRepertoryInfo) Reset() {
	*x = ReplyRepertoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_favorite_repertory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyRepertoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyRepertoryInfo) ProtoMessage() {}

func (x *ReplyRepertoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_favorite_repertory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyRepertoryInfo.ProtoReflect.Descriptor instead.
func (*ReplyRepertoryInfo) Descriptor() ([]byte, []int) {
	return file_proto_favorite_repertory_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyRepertoryInfo) GetInfo() *RepertoryInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ReplyRepertoryInfo) GetStatus() *ReplyStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type ReqRepertoryBags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	List []string `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ReqRepertoryBags) Reset() {
	*x = ReqRepertoryBags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_favorite_repertory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqRepertoryBags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqRepertoryBags) ProtoMessage() {}

func (x *ReqRepertoryBags) ProtoReflect() protoreflect.Message {
	mi := &file_proto_favorite_repertory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqRepertoryBags.ProtoReflect.Descriptor instead.
func (*ReqRepertoryBags) Descriptor() ([]byte, []int) {
	return file_proto_favorite_repertory_proto_rawDescGZIP(), []int{2}
}

func (x *ReqRepertoryBags) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ReqRepertoryBags) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

var File_proto_favorite_repertory_proto protoreflect.FileDescriptor

var file_proto_favorite_repertory_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x2f, 0x72, 0x65, 0x70, 0x65, 0x72, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe1, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x65, 0x72, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x61, 0x67, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70,
	0x65, 0x72, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x65,
	0x72, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x38, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x52, 0x65, 0x70,
	0x65, 0x72, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x61, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x32, 0xe7, 0x02, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x65, 0x72, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12,
	0x1d, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x24,
	0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x65, 0x72, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64,
	0x4f, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x24, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x65, 0x72,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0b, 0x53, 0x75,
	0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x24, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x70, 0x65, 0x72, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x12, 0x58, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22,
	0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x2e, 0x52, 0x65, 0x71, 0x52, 0x65, 0x70, 0x65, 0x72, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x61,
	0x67, 0x73, 0x1a, 0x24, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x65, 0x72,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_favorite_repertory_proto_rawDescOnce sync.Once
	file_proto_favorite_repertory_proto_rawDescData = file_proto_favorite_repertory_proto_rawDesc
)

func file_proto_favorite_repertory_proto_rawDescGZIP() []byte {
	file_proto_favorite_repertory_proto_rawDescOnce.Do(func() {
		file_proto_favorite_repertory_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_favorite_repertory_proto_rawDescData)
	})
	return file_proto_favorite_repertory_proto_rawDescData
}

var file_proto_favorite_repertory_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_favorite_repertory_proto_goTypes = []interface{}{
	(*RepertoryInfo)(nil),      // 0: omo.msp.Favorite.RepertoryInfo
	(*ReplyRepertoryInfo)(nil), // 1: omo.msp.Favorite.ReplyRepertoryInfo
	(*ReqRepertoryBags)(nil),   // 2: omo.msp.Favorite.ReqRepertoryBags
	(*ReplyStatus)(nil),        // 3: omo.msp.Favorite.ReplyStatus
	(*RequestInfo)(nil),        // 4: omo.msp.Favorite.RequestInfo
}
var file_proto_favorite_repertory_proto_depIdxs = []int32{
	0, // 0: omo.msp.Favorite.ReplyRepertoryInfo.info:type_name -> omo.msp.Favorite.RepertoryInfo
	3, // 1: omo.msp.Favorite.ReplyRepertoryInfo.status:type_name -> omo.msp.Favorite.ReplyStatus
	4, // 2: omo.msp.Favorite.RepertoryService.GetOne:input_type -> omo.msp.Favorite.RequestInfo
	4, // 3: omo.msp.Favorite.RepertoryService.AppendOne:input_type -> omo.msp.Favorite.RequestInfo
	4, // 4: omo.msp.Favorite.RepertoryService.SubtractOne:input_type -> omo.msp.Favorite.RequestInfo
	2, // 5: omo.msp.Favorite.RepertoryService.UpdateList:input_type -> omo.msp.Favorite.ReqRepertoryBags
	1, // 6: omo.msp.Favorite.RepertoryService.GetOne:output_type -> omo.msp.Favorite.ReplyRepertoryInfo
	1, // 7: omo.msp.Favorite.RepertoryService.AppendOne:output_type -> omo.msp.Favorite.ReplyRepertoryInfo
	1, // 8: omo.msp.Favorite.RepertoryService.SubtractOne:output_type -> omo.msp.Favorite.ReplyRepertoryInfo
	1, // 9: omo.msp.Favorite.RepertoryService.UpdateList:output_type -> omo.msp.Favorite.ReplyRepertoryInfo
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_favorite_repertory_proto_init() }
func file_proto_favorite_repertory_proto_init() {
	if File_proto_favorite_repertory_proto != nil {
		return
	}
	file_proto_favorite_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_favorite_repertory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepertoryInfo); i {
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
		file_proto_favorite_repertory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyRepertoryInfo); i {
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
		file_proto_favorite_repertory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqRepertoryBags); i {
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
			RawDescriptor: file_proto_favorite_repertory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_favorite_repertory_proto_goTypes,
		DependencyIndexes: file_proto_favorite_repertory_proto_depIdxs,
		MessageInfos:      file_proto_favorite_repertory_proto_msgTypes,
	}.Build()
	File_proto_favorite_repertory_proto = out.File
	file_proto_favorite_repertory_proto_rawDesc = nil
	file_proto_favorite_repertory_proto_goTypes = nil
	file_proto_favorite_repertory_proto_depIdxs = nil
}