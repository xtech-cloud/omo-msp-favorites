// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/message.proto

package favorite

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageInfo struct {
	Uid                  string    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Created              uint64    `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Name                 string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Date                 *DateInfo `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Remark               string    `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string    `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Organizer            string    `protobuf:"bytes,7,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Type                 uint32    `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	Entity               string    `protobuf:"bytes,9,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 []string  `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	Targets              []string  `protobuf:"bytes,11,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *MessageInfo) Reset()         { *m = MessageInfo{} }
func (m *MessageInfo) String() string { return proto.CompactTextString(m) }
func (*MessageInfo) ProtoMessage()    {}
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a2cdcb179ecaf9, []int{0}
}

func (m *MessageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageInfo.Unmarshal(m, b)
}
func (m *MessageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageInfo.Marshal(b, m, deterministic)
}
func (m *MessageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageInfo.Merge(m, src)
}
func (m *MessageInfo) XXX_Size() int {
	return xxx_messageInfo_MessageInfo.Size(m)
}
func (m *MessageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MessageInfo proto.InternalMessageInfo

func (m *MessageInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *MessageInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *MessageInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MessageInfo) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *MessageInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *MessageInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MessageInfo) GetOrganizer() string {
	if m != nil {
		return m.Organizer
	}
	return ""
}

func (m *MessageInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MessageInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *MessageInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MessageInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type TargetInfo struct {
	Time                 uint64   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Entity               string   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	Targets              []string `protobuf:"bytes,4,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetInfo) Reset()         { *m = TargetInfo{} }
func (m *TargetInfo) String() string { return proto.CompactTextString(m) }
func (*TargetInfo) ProtoMessage()    {}
func (*TargetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a2cdcb179ecaf9, []int{1}
}

func (m *TargetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetInfo.Unmarshal(m, b)
}
func (m *TargetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetInfo.Marshal(b, m, deterministic)
}
func (m *TargetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetInfo.Merge(m, src)
}
func (m *TargetInfo) XXX_Size() int {
	return xxx_messageInfo_TargetInfo.Size(m)
}
func (m *TargetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TargetInfo proto.InternalMessageInfo

func (m *TargetInfo) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *TargetInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TargetInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *TargetInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqMessageFilter struct {
	Key                  string        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Page                 uint32        `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32        `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	Array                []string      `protobuf:"bytes,5,rep,name=array,proto3" json:"array,omitempty"`
	List                 []*TargetInfo `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReqMessageFilter) Reset()         { *m = ReqMessageFilter{} }
func (m *ReqMessageFilter) String() string { return proto.CompactTextString(m) }
func (*ReqMessageFilter) ProtoMessage()    {}
func (*ReqMessageFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a2cdcb179ecaf9, []int{2}
}

func (m *ReqMessageFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMessageFilter.Unmarshal(m, b)
}
func (m *ReqMessageFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMessageFilter.Marshal(b, m, deterministic)
}
func (m *ReqMessageFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMessageFilter.Merge(m, src)
}
func (m *ReqMessageFilter) XXX_Size() int {
	return xxx_messageInfo_ReqMessageFilter.Size(m)
}
func (m *ReqMessageFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMessageFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMessageFilter proto.InternalMessageInfo

func (m *ReqMessageFilter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReqMessageFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ReqMessageFilter) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReqMessageFilter) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReqMessageFilter) GetArray() []string {
	if m != nil {
		return m.Array
	}
	return nil
}

func (m *ReqMessageFilter) GetList() []*TargetInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyMessages struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32         `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Count                uint32         `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	List                 []*MessageInfo `protobuf:"bytes,6,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyMessages) Reset()         { *m = ReplyMessages{} }
func (m *ReplyMessages) String() string { return proto.CompactTextString(m) }
func (*ReplyMessages) ProtoMessage()    {}
func (*ReplyMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a2cdcb179ecaf9, []int{3}
}

func (m *ReplyMessages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMessages.Unmarshal(m, b)
}
func (m *ReplyMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMessages.Marshal(b, m, deterministic)
}
func (m *ReplyMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMessages.Merge(m, src)
}
func (m *ReplyMessages) XXX_Size() int {
	return xxx_messageInfo_ReplyMessages.Size(m)
}
func (m *ReplyMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMessages.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMessages proto.InternalMessageInfo

func (m *ReplyMessages) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyMessages) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyMessages) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyMessages) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyMessages) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReplyMessages) GetList() []*MessageInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageInfo)(nil), "omo.msp.favorite.MessageInfo")
	proto.RegisterType((*TargetInfo)(nil), "omo.msp.favorite.TargetInfo")
	proto.RegisterType((*ReqMessageFilter)(nil), "omo.msp.favorite.ReqMessageFilter")
	proto.RegisterType((*ReplyMessages)(nil), "omo.msp.favorite.ReplyMessages")
}

func init() {
	proto.RegisterFile("proto/favorite/message.proto", fileDescriptor_12a2cdcb179ecaf9)
}

var fileDescriptor_12a2cdcb179ecaf9 = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xc1, 0x8e, 0x94, 0x40,
	0x10, 0x95, 0x81, 0x61, 0x9d, 0x26, 0x6c, 0x26, 0x9d, 0x8d, 0xe9, 0x8c, 0x63, 0x24, 0x9c, 0x38,
	0xb1, 0x3a, 0xc6, 0x1f, 0xd8, 0x18, 0x8d, 0x07, 0x2f, 0xbd, 0x7b, 0xf2, 0xd6, 0x3b, 0x53, 0x83,
	0x64, 0x80, 0xc6, 0xee, 0x66, 0x0c, 0xfe, 0x8e, 0x77, 0x7f, 0xc5, 0x5f, 0x32, 0x5d, 0x0d, 0xce,
	0xb0, 0xbb, 0xde, 0xea, 0x15, 0x8f, 0xaa, 0xf7, 0xea, 0x01, 0x59, 0xb7, 0x4a, 0x1a, 0x79, 0xbd,
	0x17, 0x47, 0xa9, 0x4a, 0x03, 0xd7, 0x35, 0x68, 0x2d, 0x0a, 0xc8, 0xb1, 0x4d, 0x97, 0xb2, 0x96,
	0x79, 0xad, 0xdb, 0x7c, 0x7c, 0xbe, 0x7a, 0xf9, 0x80, 0xbf, 0x95, 0x75, 0x2d, 0x1b, 0x47, 0x4f,
	0x7f, 0xcd, 0x48, 0xf4, 0xc5, 0x0d, 0xf8, 0xdc, 0xec, 0x25, 0x5d, 0x12, 0xbf, 0x2b, 0x77, 0xcc,
	0x4b, 0xbc, 0x6c, 0xc1, 0x6d, 0x49, 0x19, 0xb9, 0xd8, 0x2a, 0x10, 0x06, 0x76, 0x6c, 0x96, 0x78,
	0x59, 0xc0, 0x47, 0x48, 0x29, 0x09, 0x1a, 0x51, 0x03, 0xf3, 0x91, 0x8c, 0x35, 0xcd, 0x49, 0xb0,
	0x13, 0x06, 0x58, 0x90, 0x78, 0x59, 0xb4, 0x59, 0xe5, 0x0f, 0xd5, 0xe4, 0x1f, 0x84, 0xc1, 0x4d,
	0x1c, 0x79, 0xf4, 0x05, 0x09, 0x15, 0xd4, 0x42, 0x1d, 0xd8, 0x1c, 0xa7, 0x0c, 0x88, 0x5e, 0x91,
	0xb9, 0xfc, 0xd1, 0x80, 0x62, 0x21, 0xb6, 0x1d, 0xa0, 0x6b, 0xb2, 0x90, 0xaa, 0x10, 0x4d, 0xf9,
	0x13, 0x14, 0xbb, 0xc0, 0x27, 0xa7, 0x86, 0xd5, 0x63, 0xfa, 0x16, 0xd8, 0xf3, 0xc4, 0xcb, 0x62,
	0x8e, 0xb5, 0x9d, 0x0f, 0x8d, 0x29, 0x4d, 0xcf, 0x16, 0x6e, 0xbe, 0x43, 0xc8, 0x15, 0x85, 0x66,
	0x24, 0xf1, 0xad, 0x76, 0x5b, 0x5b, 0xa7, 0x46, 0xa8, 0x02, 0x8c, 0x66, 0x11, 0xb6, 0x47, 0x98,
	0x7e, 0x23, 0xe4, 0x0e, 0x4b, 0xbc, 0x91, 0x7d, 0xb7, 0xac, 0x01, 0x8f, 0x14, 0x70, 0xac, 0x4f,
	0x7a, 0x67, 0xe7, 0x7a, 0x4f, 0xdb, 0xfd, 0xc9, 0xf6, 0xb3, 0x4d, 0xc1, 0x74, 0xd3, 0x6f, 0x8f,
	0x2c, 0x39, 0x7c, 0x1f, 0x22, 0xf9, 0x58, 0x56, 0x06, 0x94, 0x0d, 0xe5, 0x00, 0xfd, 0x18, 0xca,
	0x01, 0x7a, 0xbb, 0xee, 0x28, 0xaa, 0x0e, 0xc6, 0x75, 0x08, 0xac, 0xb0, 0x56, 0x14, 0x2e, 0x90,
	0x98, 0x63, 0x6d, 0x25, 0x34, 0x5d, 0x7d, 0x0f, 0x0a, 0x23, 0x89, 0xf9, 0x80, 0xec, 0x04, 0xa1,
	0x94, 0xe8, 0xd9, 0x1c, 0x05, 0x38, 0x40, 0xdf, 0x90, 0xa0, 0x2a, 0xb5, 0x61, 0x61, 0xe2, 0x67,
	0xd1, 0x66, 0xfd, 0x38, 0xbe, 0xd3, 0x19, 0x38, 0x32, 0xd3, 0x3f, 0x1e, 0x89, 0x39, 0xb4, 0x55,
	0x3f, 0x48, 0xd6, 0xf4, 0x3d, 0x09, 0xb5, 0x11, 0xa6, 0xd3, 0x28, 0x38, 0xda, 0xbc, 0x7a, 0x3c,
	0x05, 0x5f, 0xb8, 0x45, 0x12, 0x1f, 0xc8, 0x56, 0x90, 0x91, 0x46, 0x54, 0x68, 0x29, 0xe6, 0x0e,
	0xd8, 0xae, 0xb5, 0xa1, 0x07, 0x4f, 0x0e, 0xfc, 0x33, 0x1a, 0x9c, 0x19, 0xbd, 0x22, 0xf3, 0xad,
	0xec, 0x1a, 0x83, 0x1f, 0x52, 0xcc, 0x1d, 0xa0, 0x6f, 0x27, 0x86, 0x9e, 0x90, 0x72, 0xf6, 0xf1,
	0x3b, 0x47, 0x9b, 0x3d, 0xb9, 0x1c, 0x9a, 0xb7, 0xa0, 0x8e, 0xe5, 0x16, 0xe8, 0x1d, 0x89, 0x3e,
	0x81, 0xb9, 0xe9, 0x87, 0x38, 0xd2, 0xa7, 0x0c, 0x4d, 0x23, 0x5b, 0xbd, 0xfe, 0x8f, 0xe9, 0xf1,
	0x4a, 0xe9, 0xb3, 0x9b, 0xe5, 0xd7, 0xcb, 0xe9, 0x9f, 0x79, 0x1f, 0x22, 0x7e, 0xf7, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x3a, 0x28, 0x5f, 0xde, 0xe2, 0x03, 0x00, 0x00,
}