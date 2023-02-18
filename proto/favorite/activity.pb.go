// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/activity.proto

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

type ActivityInfo struct {
	Uid                  string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64        `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 uint32        `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64         `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64         `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string        `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string        `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string        `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string        `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Require              string        `protobuf:"bytes,10,opt,name=require,proto3" json:"require,omitempty"`
	Cover                string        `protobuf:"bytes,11,opt,name=cover,proto3" json:"cover,omitempty"`
	Owner                string        `protobuf:"bytes,12,opt,name=owner,proto3" json:"owner,omitempty"`
	Date                 *DateInfo     `protobuf:"bytes,13,opt,name=date,proto3" json:"date,omitempty"`
	Place                *PlaceInfo    `protobuf:"bytes,14,opt,name=place,proto3" json:"place,omitempty"`
	Organizer            string        `protobuf:"bytes,15,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Limit                uint32        `protobuf:"varint,16,opt,name=limit,proto3" json:"limit,omitempty"`
	Status               uint32        `protobuf:"varint,17,opt,name=status,proto3" json:"status,omitempty"`
	Show                 uint32        `protobuf:"varint,18,opt,name=show,proto3" json:"show,omitempty"`
	Participant          uint32        `protobuf:"varint,19,opt,name=participant,proto3" json:"participant,omitempty"`
	Template             string        `protobuf:"bytes,20,opt,name=template,proto3" json:"template,omitempty"`
	Prize                *PrizeInfo    `protobuf:"bytes,21,opt,name=prize,proto3" json:"prize,omitempty"`
	Tags                 []string      `protobuf:"bytes,41,rep,name=tags,proto3" json:"tags,omitempty"`
	Assets               []string      `protobuf:"bytes,42,rep,name=assets,proto3" json:"assets,omitempty"`
	Targets              []string      `protobuf:"bytes,43,rep,name=targets,proto3" json:"targets,omitempty"`
	Opuses               []*OpusInfo   `protobuf:"bytes,44,rep,name=opuses,proto3" json:"opuses,omitempty"`
	Records              []*RecordInfo `protobuf:"bytes,45,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ActivityInfo) Reset()         { *m = ActivityInfo{} }
func (m *ActivityInfo) String() string { return proto.CompactTextString(m) }
func (*ActivityInfo) ProtoMessage()    {}
func (*ActivityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{0}
}

func (m *ActivityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityInfo.Unmarshal(m, b)
}
func (m *ActivityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityInfo.Marshal(b, m, deterministic)
}
func (m *ActivityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityInfo.Merge(m, src)
}
func (m *ActivityInfo) XXX_Size() int {
	return xxx_messageInfo_ActivityInfo.Size(m)
}
func (m *ActivityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityInfo proto.InternalMessageInfo

func (m *ActivityInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ActivityInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActivityInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ActivityInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ActivityInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ActivityInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ActivityInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ActivityInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActivityInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ActivityInfo) GetRequire() string {
	if m != nil {
		return m.Require
	}
	return ""
}

func (m *ActivityInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ActivityInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ActivityInfo) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ActivityInfo) GetPlace() *PlaceInfo {
	if m != nil {
		return m.Place
	}
	return nil
}

func (m *ActivityInfo) GetOrganizer() string {
	if m != nil {
		return m.Organizer
	}
	return ""
}

func (m *ActivityInfo) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ActivityInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ActivityInfo) GetShow() uint32 {
	if m != nil {
		return m.Show
	}
	return 0
}

func (m *ActivityInfo) GetParticipant() uint32 {
	if m != nil {
		return m.Participant
	}
	return 0
}

func (m *ActivityInfo) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *ActivityInfo) GetPrize() *PrizeInfo {
	if m != nil {
		return m.Prize
	}
	return nil
}

func (m *ActivityInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ActivityInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *ActivityInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *ActivityInfo) GetOpuses() []*OpusInfo {
	if m != nil {
		return m.Opuses
	}
	return nil
}

func (m *ActivityInfo) GetRecords() []*RecordInfo {
	if m != nil {
		return m.Records
	}
	return nil
}

type PrizeInfo struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string      `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Ranks                []*RankInfo `protobuf:"bytes,3,rep,name=ranks,proto3" json:"ranks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PrizeInfo) Reset()         { *m = PrizeInfo{} }
func (m *PrizeInfo) String() string { return proto.CompactTextString(m) }
func (*PrizeInfo) ProtoMessage()    {}
func (*PrizeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{1}
}

func (m *PrizeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrizeInfo.Unmarshal(m, b)
}
func (m *PrizeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrizeInfo.Marshal(b, m, deterministic)
}
func (m *PrizeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrizeInfo.Merge(m, src)
}
func (m *PrizeInfo) XXX_Size() int {
	return xxx_messageInfo_PrizeInfo.Size(m)
}
func (m *PrizeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PrizeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PrizeInfo proto.InternalMessageInfo

func (m *PrizeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PrizeInfo) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *PrizeInfo) GetRanks() []*RankInfo {
	if m != nil {
		return m.Ranks
	}
	return nil
}

type RecordInfo struct {
	Option               uint32   `protobuf:"varint,1,opt,name=option,proto3" json:"option,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordInfo) Reset()         { *m = RecordInfo{} }
func (m *RecordInfo) String() string { return proto.CompactTextString(m) }
func (*RecordInfo) ProtoMessage()    {}
func (*RecordInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{2}
}

func (m *RecordInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordInfo.Unmarshal(m, b)
}
func (m *RecordInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordInfo.Marshal(b, m, deterministic)
}
func (m *RecordInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordInfo.Merge(m, src)
}
func (m *RecordInfo) XXX_Size() int {
	return xxx_messageInfo_RecordInfo.Size(m)
}
func (m *RecordInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RecordInfo proto.InternalMessageInfo

func (m *RecordInfo) GetOption() uint32 {
	if m != nil {
		return m.Option
	}
	return 0
}

func (m *RecordInfo) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *RecordInfo) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *RecordInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *RecordInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RankInfo struct {
	Index                uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Count                uint32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankInfo) Reset()         { *m = RankInfo{} }
func (m *RankInfo) String() string { return proto.CompactTextString(m) }
func (*RankInfo) ProtoMessage()    {}
func (*RankInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{3}
}

func (m *RankInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankInfo.Unmarshal(m, b)
}
func (m *RankInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankInfo.Marshal(b, m, deterministic)
}
func (m *RankInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankInfo.Merge(m, src)
}
func (m *RankInfo) XXX_Size() int {
	return xxx_messageInfo_RankInfo.Size(m)
}
func (m *RankInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RankInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RankInfo proto.InternalMessageInfo

func (m *RankInfo) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *RankInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RankInfo) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type OpusInfo struct {
	Rank                 uint32   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Asset                string   `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpusInfo) Reset()         { *m = OpusInfo{} }
func (m *OpusInfo) String() string { return proto.CompactTextString(m) }
func (*OpusInfo) ProtoMessage()    {}
func (*OpusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{4}
}

func (m *OpusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpusInfo.Unmarshal(m, b)
}
func (m *OpusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpusInfo.Marshal(b, m, deterministic)
}
func (m *OpusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpusInfo.Merge(m, src)
}
func (m *OpusInfo) XXX_Size() int {
	return xxx_messageInfo_OpusInfo.Size(m)
}
func (m *OpusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OpusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OpusInfo proto.InternalMessageInfo

func (m *OpusInfo) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *OpusInfo) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *OpusInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type ReqActivityAdd struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string     `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Remark               string     `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string     `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Type                 int32      `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string     `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Date                 *DateInfo  `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	Place                *PlaceInfo `protobuf:"bytes,8,opt,name=place,proto3" json:"place,omitempty"`
	Organizer            string     `protobuf:"bytes,9,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Require              string     `protobuf:"bytes,10,opt,name=require,proto3" json:"require,omitempty"`
	Limit                uint32     `protobuf:"varint,11,opt,name=limit,proto3" json:"limit,omitempty"`
	Show                 uint32     `protobuf:"varint,12,opt,name=show,proto3" json:"show,omitempty"`
	Template             string     `protobuf:"bytes,13,opt,name=template,proto3" json:"template,omitempty"`
	Prize                *PrizeInfo `protobuf:"bytes,14,opt,name=prize,proto3" json:"prize,omitempty"`
	Status               uint32     `protobuf:"varint,15,opt,name=status,proto3" json:"status,omitempty"`
	Assets               []string   `protobuf:"bytes,31,rep,name=assets,proto3" json:"assets,omitempty"`
	Tags                 []string   `protobuf:"bytes,32,rep,name=tags,proto3" json:"tags,omitempty"`
	Targets              []string   `protobuf:"bytes,33,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReqActivityAdd) Reset()         { *m = ReqActivityAdd{} }
func (m *ReqActivityAdd) String() string { return proto.CompactTextString(m) }
func (*ReqActivityAdd) ProtoMessage()    {}
func (*ReqActivityAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{5}
}

func (m *ReqActivityAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqActivityAdd.Unmarshal(m, b)
}
func (m *ReqActivityAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqActivityAdd.Marshal(b, m, deterministic)
}
func (m *ReqActivityAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqActivityAdd.Merge(m, src)
}
func (m *ReqActivityAdd) XXX_Size() int {
	return xxx_messageInfo_ReqActivityAdd.Size(m)
}
func (m *ReqActivityAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqActivityAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqActivityAdd proto.InternalMessageInfo

func (m *ReqActivityAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqActivityAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqActivityAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqActivityAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqActivityAdd) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqActivityAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqActivityAdd) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ReqActivityAdd) GetPlace() *PlaceInfo {
	if m != nil {
		return m.Place
	}
	return nil
}

func (m *ReqActivityAdd) GetOrganizer() string {
	if m != nil {
		return m.Organizer
	}
	return ""
}

func (m *ReqActivityAdd) GetRequire() string {
	if m != nil {
		return m.Require
	}
	return ""
}

func (m *ReqActivityAdd) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReqActivityAdd) GetShow() uint32 {
	if m != nil {
		return m.Show
	}
	return 0
}

func (m *ReqActivityAdd) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *ReqActivityAdd) GetPrize() *PrizeInfo {
	if m != nil {
		return m.Prize
	}
	return nil
}

func (m *ReqActivityAdd) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReqActivityAdd) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *ReqActivityAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ReqActivityAdd) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqActivityUpdate struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string     `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Cover                string     `protobuf:"bytes,4,opt,name=cover,proto3" json:"cover,omitempty"`
	Operator             string     `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Date                 *DateInfo  `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	Place                *PlaceInfo `protobuf:"bytes,7,opt,name=place,proto3" json:"place,omitempty"`
	Organizer            string     `protobuf:"bytes,8,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Require              string     `protobuf:"bytes,9,opt,name=require,proto3" json:"require,omitempty"`
	Limit                uint32     `protobuf:"varint,10,opt,name=limit,proto3" json:"limit,omitempty"`
	Tags                 []string   `protobuf:"bytes,21,rep,name=tags,proto3" json:"tags,omitempty"`
	Targets              []string   `protobuf:"bytes,22,rep,name=targets,proto3" json:"targets,omitempty"`
	Assets               []string   `protobuf:"bytes,23,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReqActivityUpdate) Reset()         { *m = ReqActivityUpdate{} }
func (m *ReqActivityUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqActivityUpdate) ProtoMessage()    {}
func (*ReqActivityUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{6}
}

func (m *ReqActivityUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqActivityUpdate.Unmarshal(m, b)
}
func (m *ReqActivityUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqActivityUpdate.Marshal(b, m, deterministic)
}
func (m *ReqActivityUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqActivityUpdate.Merge(m, src)
}
func (m *ReqActivityUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqActivityUpdate.Size(m)
}
func (m *ReqActivityUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqActivityUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqActivityUpdate proto.InternalMessageInfo

func (m *ReqActivityUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqActivityUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqActivityUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqActivityUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqActivityUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqActivityUpdate) GetDate() *DateInfo {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *ReqActivityUpdate) GetPlace() *PlaceInfo {
	if m != nil {
		return m.Place
	}
	return nil
}

func (m *ReqActivityUpdate) GetOrganizer() string {
	if m != nil {
		return m.Organizer
	}
	return ""
}

func (m *ReqActivityUpdate) GetRequire() string {
	if m != nil {
		return m.Require
	}
	return ""
}

func (m *ReqActivityUpdate) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReqActivityUpdate) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ReqActivityUpdate) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *ReqActivityUpdate) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ReqActivityPrize struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc                 string      `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Operator             string      `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Ranks                []*RankInfo `protobuf:"bytes,15,rep,name=ranks,proto3" json:"ranks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReqActivityPrize) Reset()         { *m = ReqActivityPrize{} }
func (m *ReqActivityPrize) String() string { return proto.CompactTextString(m) }
func (*ReqActivityPrize) ProtoMessage()    {}
func (*ReqActivityPrize) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{7}
}

func (m *ReqActivityPrize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqActivityPrize.Unmarshal(m, b)
}
func (m *ReqActivityPrize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqActivityPrize.Marshal(b, m, deterministic)
}
func (m *ReqActivityPrize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqActivityPrize.Merge(m, src)
}
func (m *ReqActivityPrize) XXX_Size() int {
	return xxx_messageInfo_ReqActivityPrize.Size(m)
}
func (m *ReqActivityPrize) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqActivityPrize.DiscardUnknown(m)
}

var xxx_messageInfo_ReqActivityPrize proto.InternalMessageInfo

func (m *ReqActivityPrize) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqActivityPrize) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqActivityPrize) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *ReqActivityPrize) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqActivityPrize) GetRanks() []*RankInfo {
	if m != nil {
		return m.Ranks
	}
	return nil
}

type ReqActivityOpuses struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string      `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []*OpusInfo `protobuf:"bytes,13,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReqActivityOpuses) Reset()         { *m = ReqActivityOpuses{} }
func (m *ReqActivityOpuses) String() string { return proto.CompactTextString(m) }
func (*ReqActivityOpuses) ProtoMessage()    {}
func (*ReqActivityOpuses) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{8}
}

func (m *ReqActivityOpuses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqActivityOpuses.Unmarshal(m, b)
}
func (m *ReqActivityOpuses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqActivityOpuses.Marshal(b, m, deterministic)
}
func (m *ReqActivityOpuses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqActivityOpuses.Merge(m, src)
}
func (m *ReqActivityOpuses) XXX_Size() int {
	return xxx_messageInfo_ReqActivityOpuses.Size(m)
}
func (m *ReqActivityOpuses) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqActivityOpuses.DiscardUnknown(m)
}

var xxx_messageInfo_ReqActivityOpuses proto.InternalMessageInfo

func (m *ReqActivityOpuses) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqActivityOpuses) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqActivityOpuses) GetList() []*OpusInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyActivityInfo struct {
	Status               *ReplyStatus  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *ActivityInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyActivityInfo) Reset()         { *m = ReplyActivityInfo{} }
func (m *ReplyActivityInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyActivityInfo) ProtoMessage()    {}
func (*ReplyActivityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{9}
}

func (m *ReplyActivityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyActivityInfo.Unmarshal(m, b)
}
func (m *ReplyActivityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyActivityInfo.Marshal(b, m, deterministic)
}
func (m *ReplyActivityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyActivityInfo.Merge(m, src)
}
func (m *ReplyActivityInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyActivityInfo.Size(m)
}
func (m *ReplyActivityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyActivityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyActivityInfo proto.InternalMessageInfo

func (m *ReplyActivityInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyActivityInfo) GetInfo() *ActivityInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyActivityList struct {
	Status               *ReplyStatus    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32          `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page                 uint32          `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32          `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*ActivityInfo `protobuf:"bytes,16,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ReplyActivityList) Reset()         { *m = ReplyActivityList{} }
func (m *ReplyActivityList) String() string { return proto.CompactTextString(m) }
func (*ReplyActivityList) ProtoMessage()    {}
func (*ReplyActivityList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{10}
}

func (m *ReplyActivityList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyActivityList.Unmarshal(m, b)
}
func (m *ReplyActivityList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyActivityList.Marshal(b, m, deterministic)
}
func (m *ReplyActivityList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyActivityList.Merge(m, src)
}
func (m *ReplyActivityList) XXX_Size() int {
	return xxx_messageInfo_ReplyActivityList.Size(m)
}
func (m *ReplyActivityList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyActivityList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyActivityList proto.InternalMessageInfo

func (m *ReplyActivityList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyActivityList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyActivityList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyActivityList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyActivityList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyActivityList) GetList() []*ActivityInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyPairList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*PairInfo  `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyPairList) Reset()         { *m = ReplyPairList{} }
func (m *ReplyPairList) String() string { return proto.CompactTextString(m) }
func (*ReplyPairList) ProtoMessage()    {}
func (*ReplyPairList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d14c73187376933, []int{11}
}

func (m *ReplyPairList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyPairList.Unmarshal(m, b)
}
func (m *ReplyPairList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyPairList.Marshal(b, m, deterministic)
}
func (m *ReplyPairList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyPairList.Merge(m, src)
}
func (m *ReplyPairList) XXX_Size() int {
	return xxx_messageInfo_ReplyPairList.Size(m)
}
func (m *ReplyPairList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyPairList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyPairList proto.InternalMessageInfo

func (m *ReplyPairList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyPairList) GetList() []*PairInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*ActivityInfo)(nil), "omo.msp.favorite.ActivityInfo")
	proto.RegisterType((*PrizeInfo)(nil), "omo.msp.favorite.PrizeInfo")
	proto.RegisterType((*RecordInfo)(nil), "omo.msp.favorite.RecordInfo")
	proto.RegisterType((*RankInfo)(nil), "omo.msp.favorite.RankInfo")
	proto.RegisterType((*OpusInfo)(nil), "omo.msp.favorite.OpusInfo")
	proto.RegisterType((*ReqActivityAdd)(nil), "omo.msp.favorite.ReqActivityAdd")
	proto.RegisterType((*ReqActivityUpdate)(nil), "omo.msp.favorite.ReqActivityUpdate")
	proto.RegisterType((*ReqActivityPrize)(nil), "omo.msp.favorite.ReqActivityPrize")
	proto.RegisterType((*ReqActivityOpuses)(nil), "omo.msp.favorite.ReqActivityOpuses")
	proto.RegisterType((*ReplyActivityInfo)(nil), "omo.msp.favorite.ReplyActivityInfo")
	proto.RegisterType((*ReplyActivityList)(nil), "omo.msp.favorite.ReplyActivityList")
	proto.RegisterType((*ReplyPairList)(nil), "omo.msp.favorite.ReplyPairList")
}

func init() {
	proto.RegisterFile("proto/favorite/activity.proto", fileDescriptor_3d14c73187376933)
}

var fileDescriptor_3d14c73187376933 = []byte{
	// 1211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x58, 0xdd, 0x72, 0x1b, 0x35,
	0x14, 0x8e, 0x7f, 0x63, 0xcb, 0xb1, 0xe3, 0x8a, 0xb4, 0x68, 0x92, 0xb6, 0x31, 0xdb, 0x9b, 0xf0,
	0xe7, 0x42, 0x18, 0xb8, 0x4f, 0x87, 0x21, 0xd3, 0x92, 0x92, 0xb0, 0xa1, 0x33, 0x0c, 0x77, 0x9b,
	0x5d, 0x25, 0x68, 0xe2, 0x5d, 0x6d, 0xb4, 0xb2, 0x4b, 0x3a, 0x0c, 0x97, 0x5c, 0xf0, 0x08, 0x3c,
	0x18, 0x2f, 0xc2, 0x1d, 0x57, 0xcc, 0x39, 0x5a, 0xd9, 0xb2, 0xe3, 0xb5, 0x9d, 0xa4, 0x77, 0x3a,
	0xda, 0xa3, 0xef, 0xfc, 0x7c, 0xfa, 0x8e, 0x3c, 0x26, 0x4f, 0x52, 0x25, 0xb5, 0x7c, 0x7e, 0x1e,
	0x8c, 0xa4, 0x12, 0x9a, 0x3f, 0x0f, 0x42, 0x2d, 0x46, 0x42, 0x5f, 0xf7, 0x71, 0x9f, 0x76, 0x65,
	0x2c, 0xfb, 0x71, 0x96, 0xf6, 0xad, 0xc3, 0xf6, 0xce, 0xcc, 0x81, 0x50, 0xc6, 0xb1, 0x4c, 0x8c,
	0xbb, 0xf7, 0x67, 0x9d, 0x6c, 0x1c, 0xe4, 0x08, 0x2f, 0x93, 0x73, 0x49, 0xbb, 0xa4, 0x32, 0x14,
	0x11, 0x2b, 0xf5, 0x4a, 0x7b, 0x4d, 0x1f, 0x96, 0xb4, 0x43, 0xca, 0x22, 0x62, 0xe5, 0x5e, 0x69,
	0xaf, 0xea, 0x97, 0x45, 0x44, 0x29, 0xa9, 0xea, 0xeb, 0x94, 0xb3, 0x4a, 0xaf, 0xb4, 0xd7, 0xf6,
	0x71, 0x4d, 0x19, 0x59, 0x0f, 0x15, 0x0f, 0x34, 0x8f, 0x58, 0xb5, 0x57, 0xda, 0xab, 0xf8, 0xd6,
	0x84, 0x2f, 0xc3, 0x34, 0xc2, 0x2f, 0x35, 0xf3, 0x25, 0x37, 0xc7, 0x67, 0xa4, 0x62, 0x75, 0x8c,
	0x66, 0x4d, 0xba, 0x4d, 0x1a, 0x32, 0xe5, 0x0a, 0x3f, 0xad, 0xe3, 0xa7, 0xb1, 0x0d, 0xd1, 0x93,
	0x20, 0xe6, 0xac, 0x81, 0xfb, 0xb8, 0xa6, 0x8f, 0x48, 0x5d, 0xf1, 0x38, 0x50, 0x97, 0xac, 0x89,
	0xbb, 0xb9, 0x05, 0x11, 0x14, 0xbf, 0x1a, 0x0a, 0xc5, 0x19, 0x31, 0x11, 0x72, 0x93, 0x6e, 0x91,
	0x5a, 0x28, 0x47, 0x5c, 0xb1, 0x16, 0xee, 0x1b, 0x03, 0x76, 0xe5, 0xdb, 0x84, 0x2b, 0xb6, 0x61,
	0x76, 0xd1, 0xa0, 0x7d, 0x52, 0x85, 0x84, 0x59, 0xbb, 0x57, 0xda, 0x6b, 0xed, 0x6f, 0xf7, 0x67,
	0x1b, 0xdc, 0xff, 0x36, 0xd0, 0x1c, 0x7a, 0xe7, 0xa3, 0x1f, 0xfd, 0x92, 0xd4, 0xd2, 0x41, 0x10,
	0x72, 0xd6, 0xc1, 0x03, 0x3b, 0x37, 0x0f, 0x9c, 0xc0, 0x67, 0x3c, 0x61, 0x3c, 0xe9, 0x63, 0xd2,
	0x94, 0xea, 0x22, 0x48, 0xc4, 0x3b, 0xae, 0xd8, 0x26, 0x06, 0x9f, 0x6c, 0x40, 0x5a, 0x03, 0x11,
	0x0b, 0xcd, 0xba, 0xd8, 0x71, 0x63, 0x40, 0xd1, 0x99, 0x0e, 0xf4, 0x30, 0x63, 0x0f, 0x70, 0x3b,
	0xb7, 0xa0, 0x41, 0xd9, 0xaf, 0xf2, 0x2d, 0xa3, 0x86, 0x1e, 0x58, 0xd3, 0x1e, 0x69, 0xa5, 0x81,
	0xd2, 0x22, 0x14, 0x69, 0x90, 0x68, 0xf6, 0x01, 0x7e, 0x72, 0xb7, 0xa0, 0xe5, 0x9a, 0xc7, 0xe9,
	0x00, 0x0a, 0xdd, 0x32, 0x2d, 0xb7, 0x36, 0x16, 0xa4, 0xc4, 0x3b, 0xce, 0x1e, 0x16, 0x16, 0x04,
	0x9f, 0xf3, 0x82, 0x60, 0x89, 0x77, 0x24, 0xb8, 0xc8, 0xd8, 0xc7, 0xbd, 0x0a, 0xb0, 0x04, 0x6b,
	0x48, 0x38, 0xc8, 0x32, 0xae, 0x33, 0xf6, 0x09, 0xee, 0xe6, 0x16, 0xb0, 0xa4, 0x03, 0x75, 0x01,
	0x1f, 0x3e, 0xc5, 0x0f, 0xd6, 0xa4, 0xfb, 0xa4, 0x2e, 0xd3, 0x61, 0xc6, 0x33, 0xf6, 0x59, 0xaf,
	0x32, 0xbf, 0xf7, 0xc7, 0xe9, 0x30, 0xc3, 0xc0, 0xb9, 0x27, 0xfd, 0x06, 0x38, 0x0f, 0xa5, 0x8a,
	0x32, 0xf6, 0x39, 0x1e, 0x7a, 0x7c, 0xf3, 0x90, 0x8f, 0x0e, 0x78, 0xcc, 0x3a, 0x7b, 0x9c, 0x34,
	0xc7, 0x55, 0x8c, 0x2f, 0x59, 0xc9, 0xb9, 0x64, 0x94, 0x54, 0x23, 0x9e, 0x85, 0x28, 0x84, 0xa6,
	0x8f, 0x6b, 0xfa, 0x05, 0xa9, 0xa9, 0x20, 0xb9, 0xcc, 0x58, 0xa5, 0x28, 0x3f, 0x3f, 0x48, 0x2e,
	0x4d, 0x63, 0xd0, 0xd1, 0xfb, 0x9d, 0x90, 0x49, 0x74, 0x68, 0x89, 0x4c, 0xb5, 0x90, 0x09, 0x46,
	0x6a, 0xfb, 0xb9, 0x05, 0xb1, 0xce, 0x95, 0x8c, 0x6d, 0x2c, 0x58, 0x83, 0x0c, 0xb5, 0x44, 0xd1,
	0x35, 0xfd, 0xb2, 0x96, 0xce, 0xa5, 0xaf, 0x4e, 0x5d, 0x7a, 0x57, 0x3c, 0xb5, 0x69, 0xf1, 0x78,
	0xaf, 0x48, 0xc3, 0x26, 0x04, 0xb7, 0x4a, 0x24, 0x11, 0xff, 0x2d, 0x0f, 0x6d, 0x8c, 0x71, 0xe5,
	0x65, 0xa7, 0x72, 0x14, 0xcb, 0x30, 0xd1, 0xb9, 0xe2, 0x8d, 0xe1, 0x1d, 0x91, 0x86, 0x6d, 0x3e,
	0x9c, 0x82, 0xf2, 0x72, 0x28, 0x5c, 0xc3, 0x29, 0x24, 0x38, 0x87, 0x32, 0x86, 0x93, 0x75, 0xc5,
	0xcd, 0xda, 0xfb, 0xab, 0x4a, 0x3a, 0x3e, 0xbf, 0xb2, 0xa3, 0xe8, 0x20, 0x8a, 0xe6, 0x92, 0x30,
	0x56, 0x68, 0xd9, 0x55, 0x68, 0x01, 0xe8, 0x44, 0xe5, 0x55, 0x57, 0xe5, 0x76, 0x7e, 0x41, 0x73,
	0x6a, 0xf9, 0xfc, 0x72, 0x9b, 0x56, 0x9f, 0x99, 0x38, 0x56, 0xff, 0xeb, 0xb7, 0xd5, 0x7f, 0xe3,
	0x6e, 0xfa, 0x6f, 0xce, 0xea, 0x7f, 0xe1, 0x18, 0x33, 0x93, 0xa1, 0xe5, 0x4e, 0x06, 0x3b, 0x01,
	0x36, 0x9c, 0x09, 0xe0, 0xea, 0xbb, 0x5d, 0xa4, 0xef, 0xce, 0xca, 0xfa, 0x9e, 0x0c, 0x9f, 0xcd,
	0xa9, 0xe1, 0x33, 0xd1, 0xf8, 0xee, 0x94, 0xc6, 0xed, 0x3c, 0xe8, 0x39, 0xf3, 0xc0, 0xd1, 0xfd,
	0x47, 0x53, 0xba, 0xf7, 0xfe, 0x2b, 0x93, 0x07, 0xce, 0x65, 0x78, 0x83, 0x0f, 0xc6, 0x9c, 0x97,
	0x69, 0xde, 0x65, 0xbd, 0xdd, 0x5d, 0x58, 0x20, 0x96, 0x31, 0xef, 0xf5, 0xdb, 0xf2, 0xbe, 0x7e,
	0x37, 0xde, 0x1b, 0x0b, 0x78, 0x6f, 0x16, 0xf0, 0x4e, 0x66, 0x78, 0xc7, 0x26, 0x3f, 0x9c, 0xdf,
	0xe4, 0x47, 0xd3, 0xc3, 0x75, 0x42, 0xd5, 0x87, 0x2e, 0x55, 0xde, 0xdf, 0x25, 0xd2, 0x75, 0x9a,
	0x8f, 0xd4, 0xaf, 0xd8, 0x7b, 0x3b, 0x22, 0x2b, 0xce, 0x88, 0x74, 0x3b, 0x5c, 0x9d, 0xe9, 0xf0,
	0x78, 0x7c, 0x6e, 0xae, 0x3a, 0x3e, 0xaf, 0xa6, 0x2e, 0xc6, 0xb1, 0x19, 0xf9, 0x37, 0x93, 0x73,
	0x83, 0x96, 0x6f, 0xd2, 0x3a, 0x10, 0x99, 0x66, 0xed, 0xa5, 0x4f, 0x0a, 0xfa, 0x79, 0x7f, 0x40,
	0xc8, 0x74, 0x70, 0x3d, 0xf5, 0x2b, 0xe9, 0xeb, 0xf1, 0xfd, 0x2f, 0x21, 0xd9, 0x4f, 0xe6, 0x3d,
	0x32, 0xe9, 0xe0, 0xfa, 0x14, 0x9d, 0xc6, 0xf2, 0xd8, 0x27, 0x55, 0x91, 0x9c, 0x4b, 0xcc, 0xa9,
	0xb5, 0xff, 0xf4, 0xe6, 0x21, 0x37, 0x88, 0x8f, 0xbe, 0xde, 0x3f, 0xa5, 0x99, 0x04, 0x8e, 0x44,
	0xa6, 0xef, 0x9a, 0xc0, 0x16, 0xa9, 0x69, 0xa9, 0x83, 0x01, 0x66, 0xd0, 0xf6, 0x8d, 0x01, 0xbb,
	0x69, 0x70, 0xc1, 0x33, 0x3b, 0xe0, 0xd1, 0x00, 0x36, 0x61, 0x81, 0xac, 0xb5, 0x7d, 0x5c, 0xc3,
	0xa5, 0x49, 0x86, 0xf1, 0x19, 0x37, 0x6a, 0x69, 0xfb, 0xb9, 0x05, 0x85, 0x61, 0x53, 0xbb, 0xd8,
	0xd4, 0xa5, 0x85, 0x61, 0x63, 0x47, 0xa4, 0x8d, 0x29, 0x9e, 0x04, 0x42, 0xdd, 0xa7, 0x26, 0x4b,
	0x68, 0xb9, 0x88, 0x50, 0x08, 0x30, 0x89, 0xbb, 0xff, 0x2f, 0x21, 0x9b, 0x36, 0x9d, 0x53, 0xae,
	0x46, 0x22, 0xe4, 0xf4, 0x47, 0x52, 0x3f, 0x88, 0xa2, 0xe3, 0x84, 0xd3, 0xde, 0xbc, 0xa0, 0xee,
	0xbb, 0xb4, 0xfd, 0xac, 0x20, 0x2d, 0xb7, 0x44, 0x6f, 0x8d, 0xfe, 0x40, 0xea, 0x87, 0x5c, 0x03,
	0xe4, 0xdc, 0x3a, 0xae, 0x86, 0x3c, 0xd3, 0xe0, 0xba, 0x2a, 0xde, 0x31, 0x59, 0x3f, 0xe4, 0x1a,
	0x1b, 0x75, 0x4f, 0x40, 0xc0, 0xf0, 0xd6, 0xe8, 0x1b, 0xd2, 0x3a, 0xe4, 0xfa, 0xc5, 0xf5, 0x77,
	0x62, 0xa0, 0xb9, 0xa2, 0xbb, 0x85, 0xa0, 0xc6, 0x61, 0x55, 0xd8, 0x53, 0xb2, 0x71, 0xc8, 0x35,
	0x70, 0x24, 0x32, 0x2d, 0xc2, 0xe5, 0xb8, 0xbd, 0x05, 0x34, 0x23, 0x84, 0xb7, 0x46, 0x4f, 0x48,
	0xc7, 0xbc, 0x02, 0x2b, 0xa4, 0x6b, 0x1c, 0xb7, 0x77, 0x0a, 0x60, 0xf3, 0x76, 0xfe, 0x4c, 0x48,
	0x8e, 0x18, 0x64, 0x9c, 0x3e, 0x5b, 0xc8, 0x7a, 0x8e, 0xb8, 0x22, 0x51, 0xaf, 0x2c, 0xf2, 0x4f,
	0x30, 0x80, 0x8b, 0xb9, 0x82, 0x7e, 0x15, 0x66, 0x99, 0x37, 0xf3, 0x88, 0x6c, 0x18, 0xac, 0x03,
	0xf3, 0x8e, 0xde, 0x0f, 0xed, 0x35, 0x69, 0xdb, 0xcc, 0xcc, 0x1b, 0x70, 0x5f, 0xb8, 0x3c, 0x39,
	0x23, 0x48, 0xfa, 0xb4, 0x10, 0x0d, 0x1c, 0x96, 0x32, 0xf2, 0xbd, 0xed, 0xdb, 0x29, 0xfc, 0x60,
	0xb9, 0x27, 0xd8, 0x4b, 0xd2, 0xf4, 0x79, 0x2c, 0x47, 0x7c, 0x05, 0x01, 0x2e, 0x81, 0x7a, 0x4d,
	0x9a, 0x07, 0x69, 0xca, 0x93, 0x68, 0x05, 0xa8, 0xdd, 0x02, 0x28, 0x3b, 0xe3, 0x50, 0xc7, 0xad,
	0xd3, 0xe1, 0x99, 0x56, 0x41, 0xa8, 0xdf, 0x0f, 0xa0, 0x6f, 0x69, 0xc8, 0x9f, 0xc3, 0xc5, 0x77,
	0xd9, 0x38, 0x2d, 0xab, 0xf9, 0x84, 0xb4, 0x0c, 0xa6, 0x79, 0xfe, 0xbd, 0x85, 0x90, 0xe8, 0xb3,
	0x04, 0xf1, 0x45, 0xf7, 0x97, 0xce, 0xf4, 0xff, 0x10, 0x67, 0x75, 0xb4, 0xbf, 0xfa, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0xc8, 0x04, 0xfb, 0xa9, 0xd1, 0x10, 0x00, 0x00,
}
