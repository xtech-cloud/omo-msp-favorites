// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/favorite.proto

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

type FavoriteInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	Creator              string   `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,9,opt,name=operator,proto3" json:"operator,omitempty"`
	Cover                string   `protobuf:"bytes,10,opt,name=cover,proto3" json:"cover,omitempty"`
	Type                 uint32   `protobuf:"varint,12,opt,name=type,proto3" json:"type,omitempty"`
	Meta                 string   `protobuf:"bytes,13,opt,name=meta,proto3" json:"meta,omitempty"`
	Status               uint32   `protobuf:"varint,14,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []string `protobuf:"bytes,30,rep,name=tags,proto3" json:"tags,omitempty"`
	Keys                 []string `protobuf:"bytes,31,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FavoriteInfo) Reset()         { *m = FavoriteInfo{} }
func (m *FavoriteInfo) String() string { return proto.CompactTextString(m) }
func (*FavoriteInfo) ProtoMessage()    {}
func (*FavoriteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{0}
}

func (m *FavoriteInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FavoriteInfo.Unmarshal(m, b)
}
func (m *FavoriteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FavoriteInfo.Marshal(b, m, deterministic)
}
func (m *FavoriteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FavoriteInfo.Merge(m, src)
}
func (m *FavoriteInfo) XXX_Size() int {
	return xxx_messageInfo_FavoriteInfo.Size(m)
}
func (m *FavoriteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FavoriteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FavoriteInfo proto.InternalMessageInfo

func (m *FavoriteInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *FavoriteInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FavoriteInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *FavoriteInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *FavoriteInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FavoriteInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *FavoriteInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *FavoriteInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *FavoriteInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *FavoriteInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *FavoriteInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *FavoriteInfo) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *FavoriteInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *FavoriteInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *FavoriteInfo) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type ReqFavoriteAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Type                 int32    `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Origin               string   `protobuf:"bytes,7,opt,name=origin,proto3" json:"origin,omitempty"`
	Status               uint32   `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	Keys                 []string `protobuf:"bytes,30,rep,name=keys,proto3" json:"keys,omitempty"`
	Tags                 []string `protobuf:"bytes,31,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteAdd) Reset()         { *m = ReqFavoriteAdd{} }
func (m *ReqFavoriteAdd) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteAdd) ProtoMessage()    {}
func (*ReqFavoriteAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{1}
}

func (m *ReqFavoriteAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteAdd.Unmarshal(m, b)
}
func (m *ReqFavoriteAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteAdd.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteAdd.Merge(m, src)
}
func (m *ReqFavoriteAdd) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteAdd.Size(m)
}
func (m *ReqFavoriteAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteAdd proto.InternalMessageInfo

func (m *ReqFavoriteAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqFavoriteAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqFavoriteAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqFavoriteAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqFavoriteAdd) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqFavoriteAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFavoriteAdd) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *ReqFavoriteAdd) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReqFavoriteAdd) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *ReqFavoriteAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqFavoriteUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string   `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Origin               string   `protobuf:"bytes,6,opt,name=origin,proto3" json:"origin,omitempty"`
	Type                 uint32   `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteUpdate) Reset()         { *m = ReqFavoriteUpdate{} }
func (m *ReqFavoriteUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteUpdate) ProtoMessage()    {}
func (*ReqFavoriteUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{2}
}

func (m *ReqFavoriteUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteUpdate.Unmarshal(m, b)
}
func (m *ReqFavoriteUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteUpdate.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteUpdate.Merge(m, src)
}
func (m *ReqFavoriteUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteUpdate.Size(m)
}
func (m *ReqFavoriteUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteUpdate proto.InternalMessageInfo

func (m *ReqFavoriteUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *ReqFavoriteUpdate) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReqFavoriteTags struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Tags                 []string `protobuf:"bytes,21,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteTags) Reset()         { *m = ReqFavoriteTags{} }
func (m *ReqFavoriteTags) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteTags) ProtoMessage()    {}
func (*ReqFavoriteTags) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{3}
}

func (m *ReqFavoriteTags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteTags.Unmarshal(m, b)
}
func (m *ReqFavoriteTags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteTags.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteTags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteTags.Merge(m, src)
}
func (m *ReqFavoriteTags) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteTags.Size(m)
}
func (m *ReqFavoriteTags) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteTags.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteTags proto.InternalMessageInfo

func (m *ReqFavoriteTags) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqFavoriteTags) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFavoriteTags) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReqFavoriteList struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Type                 uint32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteList) Reset()         { *m = ReqFavoriteList{} }
func (m *ReqFavoriteList) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteList) ProtoMessage()    {}
func (*ReqFavoriteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{4}
}

func (m *ReqFavoriteList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteList.Unmarshal(m, b)
}
func (m *ReqFavoriteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteList.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteList.Merge(m, src)
}
func (m *ReqFavoriteList) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteList.Size(m)
}
func (m *ReqFavoriteList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteList proto.InternalMessageInfo

func (m *ReqFavoriteList) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqFavoriteList) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReqFavoriteMeta struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Meta                 string   `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteMeta) Reset()         { *m = ReqFavoriteMeta{} }
func (m *ReqFavoriteMeta) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteMeta) ProtoMessage()    {}
func (*ReqFavoriteMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{5}
}

func (m *ReqFavoriteMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteMeta.Unmarshal(m, b)
}
func (m *ReqFavoriteMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteMeta.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteMeta.Merge(m, src)
}
func (m *ReqFavoriteMeta) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteMeta.Size(m)
}
func (m *ReqFavoriteMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteMeta proto.InternalMessageInfo

func (m *ReqFavoriteMeta) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqFavoriteMeta) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFavoriteMeta) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

type ReplyFavoriteInfo struct {
	Info                 *FavoriteInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyFavoriteInfo) Reset()         { *m = ReplyFavoriteInfo{} }
func (m *ReplyFavoriteInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyFavoriteInfo) ProtoMessage()    {}
func (*ReplyFavoriteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{6}
}

func (m *ReplyFavoriteInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFavoriteInfo.Unmarshal(m, b)
}
func (m *ReplyFavoriteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFavoriteInfo.Marshal(b, m, deterministic)
}
func (m *ReplyFavoriteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFavoriteInfo.Merge(m, src)
}
func (m *ReplyFavoriteInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyFavoriteInfo.Size(m)
}
func (m *ReplyFavoriteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFavoriteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFavoriteInfo proto.InternalMessageInfo

func (m *ReplyFavoriteInfo) GetInfo() *FavoriteInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyFavoriteInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyFavoriteList struct {
	Status               *ReplyStatus    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32          `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32          `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32          `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*FavoriteInfo `protobuf:"bytes,21,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReplyFavoriteList) Reset()         { *m = ReplyFavoriteList{} }
func (m *ReplyFavoriteList) String() string { return proto.CompactTextString(m) }
func (*ReplyFavoriteList) ProtoMessage()    {}
func (*ReplyFavoriteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{7}
}

func (m *ReplyFavoriteList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFavoriteList.Unmarshal(m, b)
}
func (m *ReplyFavoriteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFavoriteList.Marshal(b, m, deterministic)
}
func (m *ReplyFavoriteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFavoriteList.Merge(m, src)
}
func (m *ReplyFavoriteList) XXX_Size() int {
	return xxx_messageInfo_ReplyFavoriteList.Size(m)
}
func (m *ReplyFavoriteList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFavoriteList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFavoriteList proto.InternalMessageInfo

func (m *ReplyFavoriteList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyFavoriteList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyFavoriteList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyFavoriteList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyFavoriteList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyFavoriteList) GetList() []*FavoriteInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReqFavoriteKeys struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Keys                 []string `protobuf:"bytes,21,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFavoriteKeys) Reset()         { *m = ReqFavoriteKeys{} }
func (m *ReqFavoriteKeys) String() string { return proto.CompactTextString(m) }
func (*ReqFavoriteKeys) ProtoMessage()    {}
func (*ReqFavoriteKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{8}
}

func (m *ReqFavoriteKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFavoriteKeys.Unmarshal(m, b)
}
func (m *ReqFavoriteKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFavoriteKeys.Marshal(b, m, deterministic)
}
func (m *ReqFavoriteKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFavoriteKeys.Merge(m, src)
}
func (m *ReqFavoriteKeys) XXX_Size() int {
	return xxx_messageInfo_ReqFavoriteKeys.Size(m)
}
func (m *ReqFavoriteKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFavoriteKeys.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFavoriteKeys proto.InternalMessageInfo

func (m *ReqFavoriteKeys) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqFavoriteKeys) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqFavoriteKeys) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type ReplyFavoriteKeys struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Keys                 []string     `protobuf:"bytes,21,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyFavoriteKeys) Reset()         { *m = ReplyFavoriteKeys{} }
func (m *ReplyFavoriteKeys) String() string { return proto.CompactTextString(m) }
func (*ReplyFavoriteKeys) ProtoMessage()    {}
func (*ReplyFavoriteKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdf9f6d1ce804485, []int{9}
}

func (m *ReplyFavoriteKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyFavoriteKeys.Unmarshal(m, b)
}
func (m *ReplyFavoriteKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyFavoriteKeys.Marshal(b, m, deterministic)
}
func (m *ReplyFavoriteKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyFavoriteKeys.Merge(m, src)
}
func (m *ReplyFavoriteKeys) XXX_Size() int {
	return xxx_messageInfo_ReplyFavoriteKeys.Size(m)
}
func (m *ReplyFavoriteKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyFavoriteKeys.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyFavoriteKeys proto.InternalMessageInfo

func (m *ReplyFavoriteKeys) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyFavoriteKeys) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyFavoriteKeys) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*FavoriteInfo)(nil), "omo.msp.favorite.FavoriteInfo")
	proto.RegisterType((*ReqFavoriteAdd)(nil), "omo.msp.favorite.ReqFavoriteAdd")
	proto.RegisterType((*ReqFavoriteUpdate)(nil), "omo.msp.favorite.ReqFavoriteUpdate")
	proto.RegisterType((*ReqFavoriteTags)(nil), "omo.msp.favorite.ReqFavoriteTags")
	proto.RegisterType((*ReqFavoriteList)(nil), "omo.msp.favorite.ReqFavoriteList")
	proto.RegisterType((*ReqFavoriteMeta)(nil), "omo.msp.favorite.ReqFavoriteMeta")
	proto.RegisterType((*ReplyFavoriteInfo)(nil), "omo.msp.favorite.ReplyFavoriteInfo")
	proto.RegisterType((*ReplyFavoriteList)(nil), "omo.msp.favorite.ReplyFavoriteList")
	proto.RegisterType((*ReqFavoriteKeys)(nil), "omo.msp.favorite.ReqFavoriteKeys")
	proto.RegisterType((*ReplyFavoriteKeys)(nil), "omo.msp.favorite.ReplyFavoriteKeys")
}

func init() {
	proto.RegisterFile("proto/favorite/favorite.proto", fileDescriptor_bdf9f6d1ce804485)
}

var fileDescriptor_bdf9f6d1ce804485 = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x4f, 0xdb, 0x4a,
	0x10, 0x8f, 0x9d, 0xc4, 0x21, 0x4b, 0x12, 0x78, 0xab, 0x07, 0x5a, 0x05, 0x11, 0xf2, 0xcc, 0x25,
	0xa7, 0x3c, 0x89, 0xa7, 0x77, 0xea, 0x09, 0x0e, 0x44, 0x88, 0xd2, 0x16, 0xa7, 0x54, 0x55, 0x6f,
	0x26, 0x5e, 0x90, 0x45, 0xec, 0x35, 0xf6, 0x26, 0x55, 0x2e, 0xfd, 0x42, 0x3d, 0xf4, 0xdb, 0xf4,
	0xd0, 0x4f, 0xd1, 0x8f, 0x50, 0xcd, 0xac, 0x1d, 0x3b, 0x26, 0x7f, 0x1c, 0xb5, 0xb7, 0x99, 0xd9,
	0x9f, 0x7f, 0x33, 0xf3, 0x9b, 0xd9, 0x95, 0xc9, 0x71, 0x10, 0x0a, 0x29, 0xfe, 0x7d, 0xb0, 0xa7,
	0x22, 0x74, 0x25, 0x9f, 0x1b, 0x7d, 0x8c, 0xd3, 0x7d, 0xe1, 0x89, 0xbe, 0x17, 0x05, 0xfd, 0x24,
	0xde, 0x3e, 0xca, 0x7d, 0x30, 0x12, 0x9e, 0x27, 0x7c, 0x05, 0x37, 0x7f, 0xe8, 0xa4, 0x71, 0x19,
	0x9f, 0x5c, 0xf9, 0x0f, 0x82, 0xee, 0x93, 0xf2, 0xc4, 0x75, 0x98, 0xd6, 0xd5, 0x7a, 0x75, 0x0b,
	0x4c, 0xda, 0x22, 0xba, 0xeb, 0x30, 0xbd, 0xab, 0xf5, 0x2a, 0x96, 0xee, 0x3a, 0x94, 0x91, 0xda,
	0x28, 0xe4, 0xb6, 0xe4, 0x0e, 0x2b, 0x77, 0xb5, 0x5e, 0xd9, 0x4a, 0x5c, 0x38, 0x99, 0x04, 0x0e,
	0x9e, 0x54, 0xd4, 0x49, 0xec, 0x52, 0x4a, 0x2a, 0xbe, 0xed, 0x71, 0x56, 0x45, 0x5a, 0xb4, 0xe9,
	0xdf, 0xa4, 0x2a, 0x3e, 0xfb, 0x3c, 0x64, 0x06, 0x06, 0x95, 0x43, 0x0f, 0x89, 0x11, 0x72, 0xcf,
	0x0e, 0x9f, 0x58, 0x0d, 0xc3, 0xb1, 0x37, 0xcf, 0x2a, 0x42, 0xb6, 0x83, 0x07, 0x89, 0x4b, 0xdb,
	0x64, 0x47, 0x04, 0x3c, 0xc4, 0xa3, 0x3a, 0x1e, 0xcd, 0x7d, 0xc8, 0x31, 0x12, 0x53, 0x1e, 0x32,
	0xa2, 0x72, 0xa0, 0x03, 0xd5, 0xc8, 0x59, 0xc0, 0x59, 0xa3, 0xab, 0xf5, 0x9a, 0x16, 0xda, 0x10,
	0xf3, 0xb8, 0xb4, 0x59, 0x53, 0x55, 0x08, 0x36, 0xd4, 0x12, 0x49, 0x5b, 0x4e, 0x22, 0xd6, 0x42,
	0x64, 0xec, 0xe1, 0xf7, 0xf6, 0x63, 0xc4, 0x3a, 0xdd, 0x32, 0x60, 0xc1, 0x86, 0xd8, 0x13, 0x9f,
	0x45, 0xec, 0x44, 0xc5, 0xc0, 0x36, 0x7f, 0x6a, 0xa4, 0x65, 0xf1, 0xe7, 0x44, 0xdf, 0x73, 0x27,
	0x15, 0x42, 0x5b, 0x26, 0x84, 0xbe, 0x5c, 0x88, 0xf2, 0x82, 0x10, 0xf3, 0x96, 0x2a, 0xcb, 0x5a,
	0x02, 0x81, 0xab, 0x71, 0x4b, 0x59, 0x61, 0x8c, 0x9c, 0x30, 0x87, 0xc4, 0x10, 0xa1, 0xfb, 0xe8,
	0xfa, 0x89, 0xcc, 0xca, 0xcb, 0xb4, 0x5c, 0xcf, 0xb7, 0x8c, 0xed, 0x75, 0xd2, 0xf6, 0xe6, 0x32,
	0x9c, 0xa4, 0x32, 0x98, 0xdf, 0x34, 0xf2, 0x57, 0xa6, 0xe5, 0x3b, 0x9c, 0xff, 0x92, 0xa5, 0x4a,
	0x74, 0xd0, 0x33, 0x3a, 0x6c, 0xd7, 0x71, 0xb6, 0xbb, 0xea, 0xca, 0xee, 0x8c, 0x85, 0xee, 0x12,
	0x95, 0x6a, 0xe9, 0xe0, 0xcd, 0x21, 0xd9, 0xcb, 0x14, 0xfc, 0x1e, 0x66, 0xf9, 0xb2, 0xdc, 0x6c,
	0x32, 0x3d, 0x97, 0x2c, 0x91, 0xe1, 0x20, 0x23, 0xc3, 0xab, 0x05, 0xd2, 0xd7, 0x6e, 0x24, 0xd3,
	0x29, 0x6b, 0xd9, 0x29, 0x27, 0x15, 0xe9, 0x2b, 0x2b, 0xba, 0x81, 0x4d, 0xdc, 0xba, 0x22, 0xdc,
	0xe5, 0x4a, 0xba, 0xcb, 0xe6, 0x17, 0x98, 0x4b, 0x30, 0x9e, 0x2d, 0x5c, 0xf6, 0x33, 0x52, 0x71,
	0xfd, 0x07, 0x81, 0xbc, 0xbb, 0x67, 0x9d, 0x7e, 0xfe, 0xed, 0xe8, 0x67, 0xd1, 0x16, 0x62, 0xe9,
	0xff, 0xf3, 0x0d, 0xd1, 0xf1, 0xab, 0xe3, 0x97, 0x5f, 0x61, 0xa2, 0x21, 0x82, 0x92, 0x05, 0x32,
	0xbf, 0x6b, 0xb9, 0x02, 0x50, 0x94, 0x94, 0x4c, 0xdb, 0x82, 0x0c, 0xb4, 0x94, 0x42, 0xda, 0xe3,
	0x58, 0x36, 0xe5, 0x40, 0x34, 0xb0, 0x1f, 0x79, 0x84, 0xeb, 0xd3, 0xb4, 0x94, 0x03, 0x62, 0x80,
	0x81, 0x62, 0x34, 0x2d, 0xb4, 0x61, 0x3f, 0xfc, 0x89, 0x77, 0xcf, 0xd5, 0xe6, 0x34, 0xad, 0xd8,
	0x03, 0x3d, 0xc6, 0x6e, 0x24, 0x71, 0x94, 0x05, 0xf4, 0x00, 0x6c, 0x6e, 0x5a, 0xd7, 0x70, 0x31,
	0xb6, 0x9e, 0x16, 0x5e, 0xad, 0x83, 0xcc, 0xcb, 0x11, 0xe4, 0xc4, 0x5a, 0x41, 0x9b, 0xca, 0x57,
	0xde, 0x46, 0xbe, 0x25, 0x19, 0xcf, 0xbe, 0xd6, 0xc9, 0x5e, 0x92, 0x6d, 0xc8, 0xc3, 0xa9, 0x3b,
	0xe2, 0xf4, 0x96, 0x18, 0xe7, 0x8e, 0xf3, 0xd6, 0xe7, 0xb4, 0xbb, 0x8c, 0x38, 0xfb, 0xb0, 0xb5,
	0x4f, 0x57, 0xa4, 0xce, 0x2a, 0x66, 0x96, 0xe8, 0x1b, 0x62, 0x0c, 0xb8, 0x04, 0xca, 0xa5, 0xb5,
	0x3e, 0x4f, 0x78, 0x24, 0x01, 0x5a, 0x94, 0x6f, 0x48, 0x6a, 0x03, 0x2e, 0x71, 0x97, 0xfe, 0x59,
	0x5b, 0x23, 0x40, 0x36, 0x92, 0x02, 0xc8, 0x2c, 0xd1, 0x3b, 0xb2, 0x3b, 0xe0, 0xf2, 0x62, 0x76,
	0xe9, 0x8e, 0x25, 0x0f, 0xe9, 0xc9, 0xca, 0x4a, 0x15, 0xa0, 0x28, 0xed, 0x90, 0x34, 0x06, 0x5c,
	0xc2, 0x2c, 0xdc, 0x48, 0xba, 0xa3, 0xcd, 0xbc, 0xdd, 0x35, 0xe3, 0x44, 0x0a, 0xb3, 0x44, 0xdf,
	0x91, 0x96, 0x7a, 0x64, 0x0b, 0x94, 0xab, 0x80, 0xed, 0xa3, 0x15, 0xb4, 0xb1, 0xa4, 0x1f, 0x09,
	0x89, 0x19, 0xed, 0x88, 0xd3, 0xd3, 0xb5, 0xaa, 0xc6, 0x8c, 0x05, 0x87, 0xf5, 0x21, 0x61, 0xc6,
	0x37, 0x6d, 0xfd, 0xbc, 0x00, 0xb2, 0x35, 0x2f, 0xbe, 0xde, 0xeb, 0x79, 0x01, 0x52, 0x94, 0xf7,
	0x86, 0x34, 0x14, 0xaf, 0xba, 0x3f, 0xb4, 0xb3, 0x52, 0x59, 0x00, 0x6c, 0x14, 0xf6, 0x8a, 0xd4,
	0x2d, 0xee, 0x89, 0x29, 0x2f, 0xb0, 0xfe, 0x1b, 0xa8, 0xe6, 0x1d, 0xe3, 0xc3, 0xb0, 0xbe, 0x63,
	0x80, 0x6c, 0xec, 0x18, 0x40, 0x66, 0x89, 0xde, 0x92, 0xfa, 0x79, 0x10, 0x70, 0xdf, 0xb9, 0xe6,
	0xb3, 0xdf, 0xbd, 0xa1, 0x31, 0xe5, 0x90, 0xec, 0x0e, 0x27, 0xf7, 0x32, 0xb4, 0x47, 0xf2, 0x8f,
	0x91, 0x5e, 0xec, 0x7f, 0x6a, 0x2d, 0xfe, 0xd5, 0xde, 0x1b, 0xe8, 0xff, 0xf7, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x69, 0x28, 0x93, 0xe9, 0x1f, 0x0b, 0x00, 0x00,
}
