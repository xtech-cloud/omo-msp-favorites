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

type ResultStatus int32

const (
	ResultStatus_Success     ResultStatus = 0
	ResultStatus_MaxLimit    ResultStatus = 1
	ResultStatus_Repeated    ResultStatus = 2
	ResultStatus_NotExisted  ResultStatus = 3
	ResultStatus_DBException ResultStatus = 4
	ResultStatus_Empty       ResultStatus = 5
)

var ResultStatus_name = map[int32]string{
	0: "Success",
	1: "MaxLimit",
	2: "Repeated",
	3: "NotExisted",
	4: "DBException",
	5: "Empty",
}

var ResultStatus_value = map[string]int32{
	"Success":     0,
	"MaxLimit":    1,
	"Repeated":    2,
	"NotExisted":  3,
	"DBException": 4,
	"Empty":       5,
}

func (x ResultStatus) String() string {
	return proto.EnumName(ResultStatus_name, int32(x))
}

func (ResultStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{0}
}

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
	return fileDescriptor_fa24438d19a59ad8, []int{1}
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

type RequestInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag                 string   `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Person               bool     `protobuf:"varint,4,opt,name=person,proto3" json:"person,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{3}
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

func (m *RequestInfo) GetPerson() bool {
	if m != nil {
		return m.Person
	}
	return false
}

func (m *RequestInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type RequestFilter struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	List                 []string `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFilter) Reset()         { *m = RequestFilter{} }
func (m *RequestFilter) String() string { return proto.CompactTextString(m) }
func (*RequestFilter) ProtoMessage()    {}
func (*RequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{4}
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

func (m *RequestFilter) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
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
	return fileDescriptor_fa24438d19a59ad8, []int{5}
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
	return fileDescriptor_fa24438d19a59ad8, []int{6}
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
	Person               bool     `protobuf:"varint,3,opt,name=person,proto3" json:"person,omitempty"`
	List                 []string `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestList) Reset()         { *m = RequestList{} }
func (m *RequestList) String() string { return proto.CompactTextString(m) }
func (*RequestList) ProtoMessage()    {}
func (*RequestList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{7}
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

func (m *RequestList) GetPerson() bool {
	if m != nil {
		return m.Person
	}
	return false
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
	List                 []string     `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyList) Reset()         { *m = ReplyList{} }
func (m *ReplyList) String() string { return proto.CompactTextString(m) }
func (*ReplyList) ProtoMessage()    {}
func (*ReplyList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa24438d19a59ad8, []int{8}
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

func init() {
	proto.RegisterEnum("omo.msp.favorite.ResultStatus", ResultStatus_name, ResultStatus_value)
	proto.RegisterEnum("omo.msp.favorite.ActivityType", ActivityType_name, ActivityType_value)
	proto.RegisterType((*DateInfo)(nil), "omo.msp.favorite.DateInfo")
	proto.RegisterType((*PairInfo)(nil), "omo.msp.favorite.PairInfo")
	proto.RegisterType((*PlaceInfo)(nil), "omo.msp.favorite.PlaceInfo")
	proto.RegisterType((*RequestInfo)(nil), "omo.msp.favorite.RequestInfo")
	proto.RegisterType((*RequestFilter)(nil), "omo.msp.favorite.RequestFilter")
	proto.RegisterType((*ReplyStatus)(nil), "omo.msp.favorite.ReplyStatus")
	proto.RegisterType((*ReplyInfo)(nil), "omo.msp.favorite.ReplyInfo")
	proto.RegisterType((*RequestList)(nil), "omo.msp.favorite.RequestList")
	proto.RegisterType((*ReplyList)(nil), "omo.msp.favorite.ReplyList")
}

func init() {
	proto.RegisterFile("proto/favorite/common.proto", fileDescriptor_fa24438d19a59ad8)
}

var fileDescriptor_fa24438d19a59ad8 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x5d, 0x9a, 0xb6, 0xb4, 0x5f, 0xbb, 0x11, 0x59, 0x08, 0x45, 0x20, 0xa4, 0x2a, 0x07, 0x34,
	0xed, 0xd0, 0x49, 0x03, 0xc4, 0x81, 0x13, 0xd3, 0x0a, 0x42, 0x1a, 0x68, 0x4a, 0xe1, 0xc2, 0xcd,
	0xa4, 0x5f, 0x3b, 0x8b, 0x24, 0x36, 0xf6, 0x97, 0xd2, 0x1e, 0xf8, 0xef, 0xc8, 0x8e, 0xd3, 0xa6,
	0xa8, 0x3b, 0x70, 0x7b, 0xcf, 0xf5, 0xfb, 0xde, 0xfb, 0x5e, 0x1d, 0x78, 0xae, 0xb4, 0x24, 0x79,
	0xb9, 0xe4, 0x6b, 0xa9, 0x05, 0xe1, 0x65, 0x26, 0x8b, 0x42, 0x96, 0x53, 0x77, 0xca, 0x22, 0x59,
	0xc8, 0x69, 0x61, 0xd4, 0xb4, 0xf9, 0x39, 0x79, 0x0d, 0x83, 0x1b, 0x4e, 0xf8, 0xa9, 0x5c, 0x4a,
	0xf6, 0x04, 0x7a, 0x86, 0xb8, 0xa6, 0x38, 0x98, 0x04, 0xe7, 0xc3, 0xb4, 0x26, 0x8c, 0x41, 0xd7,
	0x90, 0x54, 0x71, 0xc7, 0x1d, 0x3a, 0x9c, 0x5c, 0xc1, 0xe0, 0x8e, 0x0b, 0xed, 0x54, 0x11, 0x84,
	0x3f, 0x71, 0xeb, 0x35, 0x16, 0xda, 0x39, 0x6b, 0x9e, 0x57, 0xe8, 0x25, 0x35, 0x49, 0xde, 0xc1,
	0xf0, 0x2e, 0xe7, 0x59, 0x6d, 0xc5, 0xa0, 0x5b, 0xf2, 0x02, 0xbd, 0xca, 0x61, 0xf6, 0x0c, 0x06,
	0xb9, 0xcc, 0x38, 0x09, 0x59, 0x7a, 0xe5, 0x8e, 0x27, 0x7f, 0x60, 0x94, 0xe2, 0xaf, 0x0a, 0x0d,
	0x35, 0x9e, 0x95, 0x58, 0x34, 0x9e, 0x95, 0x58, 0xd8, 0x81, 0xcb, 0x9c, 0xaf, 0x9a, 0x94, 0x16,
	0xdb, 0x1c, 0xf2, 0x77, 0x89, 0x3a, 0x0e, 0xeb, 0x1c, 0x8e, 0xb0, 0xa7, 0xd0, 0x57, 0xa8, 0x8d,
	0x2c, 0xe3, 0xee, 0x24, 0x38, 0x1f, 0xa4, 0x9e, 0x59, 0x7b, 0xa9, 0x50, 0x73, 0x92, 0x3a, 0xee,
	0xd5, 0xf6, 0x0d, 0x4f, 0x38, 0x9c, 0x7a, 0xfb, 0x0f, 0x22, 0x27, 0xd4, 0xfb, 0xd1, 0x41, 0x7b,
	0xb4, 0xaf, 0xa2, 0x73, 0xa4, 0x8a, 0xb0, 0x55, 0x85, 0x0d, 0x9b, 0x0b, 0x43, 0x71, 0x77, 0x12,
	0xda, 0xb0, 0x16, 0x27, 0x6f, 0xed, 0x86, 0x2a, 0xdf, 0xce, 0x89, 0x53, 0x65, 0xec, 0x95, 0x4c,
	0x2e, 0xea, 0x82, 0x4e, 0x53, 0x87, 0xed, 0x30, 0xd4, 0x5a, 0xea, 0xa6, 0x57, 0x47, 0x92, 0x1c,
	0x86, 0x4e, 0xf8, 0x40, 0x31, 0x76, 0x5d, 0xae, 0xb1, 0x24, 0xaf, 0xf2, 0x8c, 0xbd, 0x81, 0xbe,
	0x71, 0x56, 0x2e, 0xda, 0xe8, 0xea, 0xc5, 0xf4, 0xdf, 0xb7, 0x31, 0x6d, 0xe5, 0x49, 0xfd, 0xe5,
	0x64, 0xb5, 0xfb, 0x23, 0x6e, 0x85, 0xa1, 0x23, 0x7e, 0xed, 0x1a, 0x3b, 0x87, 0x35, 0xb6, 0xaa,
	0x0f, 0x0f, 0xaa, 0x3f, 0xd6, 0xc7, 0xbd, 0x5f, 0xeb, 0x01, 0x9b, 0x7d, 0xfc, 0xce, 0x7f, 0xc4,
	0xdf, 0x39, 0x85, 0x7b, 0xa7, 0x8b, 0x0c, 0xc6, 0x29, 0x9a, 0x2a, 0x27, 0x5f, 0xfd, 0x08, 0x1e,
	0xcd, 0xab, 0x2c, 0x43, 0x63, 0xa2, 0x13, 0x36, 0x86, 0xc1, 0x67, 0xbe, 0xb9, 0x15, 0x85, 0xa0,
	0x28, 0xb0, 0x2c, 0x45, 0x85, 0x9c, 0x70, 0x11, 0x75, 0xd8, 0x19, 0xc0, 0x17, 0x49, 0xb3, 0x8d,
	0x30, 0x96, 0x87, 0xec, 0x31, 0x8c, 0x6e, 0xae, 0x67, 0x9b, 0x0c, 0x95, 0x7d, 0xb3, 0x51, 0x97,
	0x0d, 0xa1, 0x37, 0x2b, 0x14, 0x6d, 0xa3, 0xde, 0xc5, 0x4b, 0x18, 0xbf, 0xcf, 0x48, 0xac, 0x05,
	0x6d, 0xbf, 0x6e, 0x15, 0x32, 0x80, 0xfe, 0x47, 0x4e, 0xf7, 0xa8, 0xa3, 0x13, 0x8b, 0xe7, 0x62,
	0x55, 0x7e, 0x53, 0x51, 0x70, 0x1d, 0x7d, 0x3f, 0x3b, 0xfc, 0x80, 0x7f, 0xf4, 0x1d, 0x7f, 0xf5,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x95, 0xc2, 0x65, 0xd9, 0x03, 0x00, 0x00,
}
