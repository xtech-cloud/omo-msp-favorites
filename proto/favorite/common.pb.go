// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/favorite/common.proto

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

type ActivityType int32

const (
	ActivityType_Gather ActivityType = 0
	ActivityType_SignUp ActivityType = 1
)

var ActivityType_name = map[int32]string{
	0: "Gather",
	1: "SignUp",
}

var ActivityType_value = map[string]int32{
	"Gather": 0,
	"SignUp": 1,
}

func (x ActivityType) String() string {
	return proto.EnumName(ActivityType_name, int32(x))
}

func (ActivityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{0}
}

type DateInfo struct {
	Start                string   `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	Stop                 string   `protobuf:"bytes,2,opt,name=stop,proto3" json:"stop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateInfo) Reset()         { *m = DateInfo{} }
func (m *DateInfo) String() string { return proto.CompactTextString(m) }
func (*DateInfo) ProtoMessage()    {}
func (*DateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{0}
}

func (m *DateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateInfo.Unmarshal(m, b)
}
func (m *DateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateInfo.Marshal(b, m, deterministic)
}
func (m *DateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateInfo.Merge(m, src)
}
func (m *DateInfo) XXX_Size() int {
	return xxx_messageInfo_DateInfo.Size(m)
}
func (m *DateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DateInfo proto.InternalMessageInfo

func (m *DateInfo) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *DateInfo) GetStop() string {
	if m != nil {
		return m.Stop
	}
	return ""
}

type PairInfo struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairInfo) Reset()         { *m = PairInfo{} }
func (m *PairInfo) String() string { return proto.CompactTextString(m) }
func (*PairInfo) ProtoMessage()    {}
func (*PairInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{1}
}

func (m *PairInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairInfo.Unmarshal(m, b)
}
func (m *PairInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairInfo.Marshal(b, m, deterministic)
}
func (m *PairInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairInfo.Merge(m, src)
}
func (m *PairInfo) XXX_Size() int {
	return xxx_messageInfo_PairInfo.Size(m)
}
func (m *PairInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PairInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PairInfo proto.InternalMessageInfo

func (m *PairInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PairInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type PlaceInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlaceInfo) Reset()         { *m = PlaceInfo{} }
func (m *PlaceInfo) String() string { return proto.CompactTextString(m) }
func (*PlaceInfo) ProtoMessage()    {}
func (*PlaceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{2}
}

func (m *PlaceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlaceInfo.Unmarshal(m, b)
}
func (m *PlaceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlaceInfo.Marshal(b, m, deterministic)
}
func (m *PlaceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlaceInfo.Merge(m, src)
}
func (m *PlaceInfo) XXX_Size() int {
	return xxx_messageInfo_PlaceInfo.Size(m)
}
func (m *PlaceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PlaceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PlaceInfo proto.InternalMessageInfo

func (m *PlaceInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlaceInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type RecordInfo struct {
	Option               uint32   `protobuf:"varint,1,opt,name=option,proto3" json:"option,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Content              string   `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Stamp                int64    `protobuf:"varint,7,opt,name=stamp,proto3" json:"stamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordInfo) Reset()         { *m = RecordInfo{} }
func (m *RecordInfo) String() string { return proto.CompactTextString(m) }
func (*RecordInfo) ProtoMessage()    {}
func (*RecordInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{3}
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

func (m *RecordInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *RecordInfo) GetStamp() int64 {
	if m != nil {
		return m.Stamp
	}
	return 0
}

type RequestInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag                 string   `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32   `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{4}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestInfo) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *RequestInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RequestInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type RequestFilter struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Page                 uint32   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	Number               uint32   `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	List                 []string `protobuf:"bytes,7,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFilter) Reset()         { *m = RequestFilter{} }
func (m *RequestFilter) String() string { return proto.CompactTextString(m) }
func (*RequestFilter) ProtoMessage()    {}
func (*RequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{5}
}

func (m *RequestFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFilter.Unmarshal(m, b)
}
func (m *RequestFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFilter.Marshal(b, m, deterministic)
}
func (m *RequestFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFilter.Merge(m, src)
}
func (m *RequestFilter) XXX_Size() int {
	return xxx_messageInfo_RequestFilter.Size(m)
}
func (m *RequestFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFilter.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFilter proto.InternalMessageInfo

func (m *RequestFilter) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestFilter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestFilter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestFilter) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *RequestFilter) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *RequestFilter) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type RequestUpdate struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []string `protobuf:"bytes,20,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUpdate) Reset()         { *m = RequestUpdate{} }
func (m *RequestUpdate) String() string { return proto.CompactTextString(m) }
func (*RequestUpdate) ProtoMessage()    {}
func (*RequestUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{6}
}

func (m *RequestUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUpdate.Unmarshal(m, b)
}
func (m *RequestUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUpdate.Marshal(b, m, deterministic)
}
func (m *RequestUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUpdate.Merge(m, src)
}
func (m *RequestUpdate) XXX_Size() int {
	return xxx_messageInfo_RequestUpdate.Size(m)
}
func (m *RequestUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUpdate proto.InternalMessageInfo

func (m *RequestUpdate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *RequestUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestUpdate) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestUpdate) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestUpdate) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type RequestState struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               uint32   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Parent               string   `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	Remark               string   `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestState) Reset()         { *m = RequestState{} }
func (m *RequestState) String() string { return proto.CompactTextString(m) }
func (*RequestState) ProtoMessage()    {}
func (*RequestState) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{7}
}

func (m *RequestState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestState.Unmarshal(m, b)
}
func (m *RequestState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestState.Marshal(b, m, deterministic)
}
func (m *RequestState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestState.Merge(m, src)
}
func (m *RequestState) XXX_Size() int {
	return xxx_messageInfo_RequestState.Size(m)
}
func (m *RequestState) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestState.DiscardUnknown(m)
}

var xxx_messageInfo_RequestState proto.InternalMessageInfo

func (m *RequestState) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestState) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RequestState) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestState) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *RequestState) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type ReplyStatus struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyStatus) Reset()         { *m = ReplyStatus{} }
func (m *ReplyStatus) String() string { return proto.CompactTextString(m) }
func (*ReplyStatus) ProtoMessage()    {}
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{8}
}

func (m *ReplyStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatus.Unmarshal(m, b)
}
func (m *ReplyStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatus.Marshal(b, m, deterministic)
}
func (m *ReplyStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatus.Merge(m, src)
}
func (m *ReplyStatus) XXX_Size() int {
	return xxx_messageInfo_ReplyStatus.Size(m)
}
func (m *ReplyStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatus proto.InternalMessageInfo

func (m *ReplyStatus) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReplyStatus) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type ReplyInfo struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Parent               string       `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyInfo) Reset()         { *m = ReplyInfo{} }
func (m *ReplyInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyInfo) ProtoMessage()    {}
func (*ReplyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{9}
}

func (m *ReplyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyInfo.Unmarshal(m, b)
}
func (m *ReplyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyInfo.Marshal(b, m, deterministic)
}
func (m *ReplyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyInfo.Merge(m, src)
}
func (m *ReplyInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyInfo.Size(m)
}
func (m *ReplyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyInfo proto.InternalMessageInfo

func (m *ReplyInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyInfo) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReplyInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type RequestList struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []string `protobuf:"bytes,21,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestList) Reset()         { *m = RequestList{} }
func (m *RequestList) String() string { return proto.CompactTextString(m) }
func (*RequestList) ProtoMessage()    {}
func (*RequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{10}
}

func (m *RequestList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestList.Unmarshal(m, b)
}
func (m *RequestList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestList.Marshal(b, m, deterministic)
}
func (m *RequestList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestList.Merge(m, src)
}
func (m *RequestList) XXX_Size() int {
	return xxx_messageInfo_RequestList.Size(m)
}
func (m *RequestList) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestList.DiscardUnknown(m)
}

var xxx_messageInfo_RequestList proto.InternalMessageInfo

func (m *RequestList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RequestList) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *RequestList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyList struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	List                 []string     `protobuf:"bytes,21,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyList) Reset()         { *m = ReplyList{} }
func (m *ReplyList) String() string { return proto.CompactTextString(m) }
func (*ReplyList) ProtoMessage()    {}
func (*ReplyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{11}
}

func (m *ReplyList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyList.Unmarshal(m, b)
}
func (m *ReplyList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyList.Marshal(b, m, deterministic)
}
func (m *ReplyList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyList.Merge(m, src)
}
func (m *ReplyList) XXX_Size() int {
	return xxx_messageInfo_ReplyList.Size(m)
}
func (m *ReplyList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyList proto.InternalMessageInfo

func (m *ReplyList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReplyList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReplyStatistic struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Key                  string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Owner                string       `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Count                uint32       `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyStatistic) Reset()         { *m = ReplyStatistic{} }
func (m *ReplyStatistic) String() string { return proto.CompactTextString(m) }
func (*ReplyStatistic) ProtoMessage()    {}
func (*ReplyStatistic) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{12}
}

func (m *ReplyStatistic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyStatistic.Unmarshal(m, b)
}
func (m *ReplyStatistic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyStatistic.Marshal(b, m, deterministic)
}
func (m *ReplyStatistic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyStatistic.Merge(m, src)
}
func (m *ReplyStatistic) XXX_Size() int {
	return xxx_messageInfo_ReplyStatistic.Size(m)
}
func (m *ReplyStatistic) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyStatistic.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyStatistic proto.InternalMessageInfo

func (m *ReplyStatistic) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyStatistic) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReplyStatistic) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReplyStatistic) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterEnum("omo.msp.favorite.ActivityType", ActivityType_name, ActivityType_value)
	proto.RegisterType((*DateInfo)(nil), "omo.msp.favorite.DateInfo")
	proto.RegisterType((*PairInfo)(nil), "omo.msp.favorite.PairInfo")
	proto.RegisterType((*PlaceInfo)(nil), "omo.msp.favorite.PlaceInfo")
	proto.RegisterType((*RecordInfo)(nil), "omo.msp.favorite.RecordInfo")
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.favorite.RequestInfo")
	proto.RegisterType((*RequestFilter)(nil), "omo.msp.favorite.RequestFilter")
	proto.RegisterType((*RequestUpdate)(nil), "omo.msp.favorite.RequestUpdate")
	proto.RegisterType((*RequestState)(nil), "omo.msp.favorite.RequestState")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.favorite.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.favorite.ReplyInfo")
	proto.RegisterType((*RequestList)(nil), "omo.msp.favorite.RequestList")
	proto.RegisterType((*ReplyList)(nil), "omo.msp.favorite.ReplyList")
	proto.RegisterType((*ReplyStatistic)(nil), "omo.msp.favorite.ReplyStatistic")
}

func init() {
	proto.RegisterFile("proto/favorite/common.proto", fileDescriptor_fa24438d19a59ad8)
}

var fileDescriptor_fa24438d19a59ad8 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0x1f, 0xbf, 0xc4, 0x69, 0xa6, 0x75, 0x14, 0x59, 0x79, 0x22, 0x0b, 0x84, 0x14, 0xf9,
	0x80, 0x2a, 0x0e, 0xa9, 0x54, 0xa8, 0x38, 0x70, 0x02, 0x21, 0x10, 0x12, 0x12, 0x95, 0x43, 0x2f,
	0xdc, 0xb6, 0xce, 0x26, 0x5d, 0xd5, 0xf6, 0x2e, 0xeb, 0x71, 0x50, 0xee, 0x70, 0x44, 0x88, 0xef,
	0xc1, 0x87, 0x44, 0xbb, 0x5e, 0x27, 0x9b, 0xe2, 0xf0, 0x72, 0x9b, 0xff, 0x7a, 0x67, 0xe6, 0x37,
	0x2f, 0x6b, 0xb8, 0x2f, 0x24, 0x47, 0x7e, 0xb6, 0x24, 0x6b, 0x2e, 0x19, 0xd2, 0xb3, 0x8c, 0x17,
	0x05, 0x2f, 0x67, 0xfa, 0x34, 0x1a, 0xf1, 0x82, 0xcf, 0x8a, 0x4a, 0xcc, 0xda, 0xcf, 0xc9, 0x13,
	0x38, 0x7a, 0x49, 0x90, 0xbe, 0x29, 0x97, 0x3c, 0x1a, 0x43, 0xaf, 0x42, 0x22, 0x31, 0x76, 0xa6,
	0xce, 0xe9, 0x20, 0x6d, 0x44, 0x14, 0x81, 0x5f, 0x21, 0x17, 0xb1, 0xab, 0x0f, 0xb5, 0x9d, 0x9c,
	0xc3, 0xd1, 0x25, 0x61, 0x52, 0x7b, 0x8d, 0xc0, 0xbb, 0xa5, 0x1b, 0xe3, 0xa3, 0x4c, 0x15, 0x67,
	0x4d, 0xf2, 0x9a, 0x1a, 0x97, 0x46, 0x24, 0xcf, 0x60, 0x70, 0x99, 0x93, 0xac, 0x49, 0x15, 0x81,
	0x5f, 0x92, 0x82, 0x1a, 0x2f, 0x6d, 0x47, 0xf7, 0xe0, 0x28, 0xe7, 0x19, 0x41, 0xc6, 0x4b, 0xe3,
	0xb9, 0xd5, 0xc9, 0x0f, 0x07, 0x20, 0xa5, 0x19, 0x97, 0x0b, 0xed, 0x3e, 0x81, 0x80, 0x0b, 0x7d,
	0x51, 0x05, 0x08, 0x53, 0xa3, 0x54, 0xd8, 0xa5, 0xe4, 0x45, 0xcb, 0xaa, 0xec, 0x68, 0x08, 0x2e,
	0xf2, 0xd8, 0xd3, 0x27, 0x2e, 0x6a, 0x5f, 0x49, 0x0b, 0x22, 0x6f, 0x63, 0x5f, 0x9f, 0x19, 0xa5,
	0xd2, 0x73, 0x41, 0x25, 0x41, 0x2e, 0xe3, 0x5e, 0x93, 0xbe, 0xd5, 0x51, 0x0c, 0xfd, 0x8c, 0x97,
	0x48, 0x4b, 0x8c, 0x03, 0xfd, 0xa9, 0x95, 0xa6, 0x67, 0x85, 0x88, 0xfb, 0x53, 0xe7, 0xd4, 0x4b,
	0x1b, 0x91, 0x7c, 0x77, 0xe0, 0x38, 0xa5, 0x1f, 0x6b, 0x5a, 0x61, 0xdb, 0xa3, 0x9a, 0x2d, 0xda,
	0x1e, 0xd5, 0x6c, 0xa1, 0x49, 0x73, 0xb2, 0xda, 0x92, 0xe6, 0x64, 0xa5, 0x62, 0xf1, 0x4f, 0x25,
	0x95, 0x06, 0xb6, 0x11, 0xbf, 0xe5, 0x8a, 0xc0, 0xc7, 0x8d, 0xa0, 0x1a, 0x2a, 0x4c, 0xb5, 0x6d,
	0xd5, 0xd7, 0xb7, 0xeb, 0x4b, 0xbe, 0x3a, 0x10, 0x1a, 0xa6, 0x57, 0x2c, 0x47, 0x2a, 0x77, 0xf9,
	0x1c, 0x3b, 0x9f, 0x99, 0xa7, 0xdb, 0x31, 0x4f, 0xcf, 0x9a, 0xa7, 0xca, 0x2d, 0xc8, 0x8a, 0xea,
	0x2e, 0x86, 0xa9, 0xb6, 0x55, 0xee, 0xb2, 0x2e, 0xae, 0x69, 0x43, 0x1a, 0xa6, 0x46, 0xa9, 0xbb,
	0x39, 0xab, 0x30, 0xee, 0x4f, 0x3d, 0x55, 0xad, 0xb2, 0x93, 0x6f, 0x3b, 0x9e, 0x2b, 0xb1, 0x20,
	0x48, 0x0f, 0xf3, 0xa8, 0xde, 0xb9, 0xbb, 0xde, 0x19, 0x42, 0xaf, 0x83, 0xd0, 0xb7, 0x09, 0xff,
	0xd0, 0x39, 0x4d, 0x34, 0xb6, 0x88, 0x3e, 0x3b, 0x70, 0x62, 0x88, 0xe6, 0xa8, 0x80, 0x7e, 0x1d,
	0xdb, 0x04, 0x82, 0x0a, 0x09, 0xd6, 0x95, 0xe6, 0x09, 0x53, 0xa3, 0xf6, 0x52, 0x79, 0x77, 0x52,
	0x4d, 0x20, 0x10, 0x44, 0xaa, 0xdd, 0x31, 0x0b, 0xd7, 0x28, 0x6b, 0x50, 0xbd, 0xbd, 0x41, 0x3d,
	0x55, 0xbb, 0x23, 0xf2, 0xcd, 0xbc, 0x09, 0x1d, 0x81, 0x9f, 0xf1, 0x05, 0x35, 0x9b, 0xae, 0x6d,
	0x55, 0x2f, 0x95, 0x92, 0xcb, 0xf6, 0x85, 0x69, 0x91, 0xe4, 0x30, 0xd0, 0x8e, 0x07, 0x56, 0x6e,
	0xc7, 0xe1, 0xee, 0x71, 0x5c, 0x6c, 0x6b, 0x52, 0xe4, 0xc7, 0xe7, 0x0f, 0x66, 0x77, 0xff, 0x12,
	0x33, 0x8b, 0xa7, 0x2d, 0x39, 0x79, 0xb7, 0x5d, 0xf1, 0xb7, 0xac, 0xc2, 0x8e, 0x7c, 0x76, 0x4f,
	0xdc, 0x03, 0xed, 0xff, 0xdf, 0x6a, 0xff, 0x8d, 0xc1, 0x3f, 0x10, 0xee, 0x62, 0xaf, 0xf5, 0x7f,
	0x8b, 0xd9, 0x99, 0xe9, 0x8b, 0x03, 0xc3, 0xed, 0x5d, 0x56, 0x21, 0xcb, 0xac, 0xe8, 0xce, 0xbf,
	0x44, 0xef, 0x7c, 0x2c, 0x1d, 0x8f, 0x78, 0x0c, 0xbd, 0x8c, 0xd7, 0x66, 0x05, 0xc2, 0xb4, 0x11,
	0x8f, 0x1e, 0xc2, 0xc9, 0xf3, 0x0c, 0xd9, 0x9a, 0xe1, 0xe6, 0xbd, 0x7a, 0xba, 0x00, 0xc1, 0x6b,
	0x82, 0x37, 0x54, 0x8e, 0xfe, 0x53, 0xf6, 0x9c, 0xad, 0xca, 0x2b, 0x31, 0x72, 0x5e, 0x8c, 0x3e,
	0x0c, 0xf7, 0xff, 0xea, 0xd7, 0x81, 0xd6, 0x8f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xfd,
	0x8d, 0x6c, 0xee, 0x05, 0x00, 0x00,
}
