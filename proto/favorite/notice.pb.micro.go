// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/favorite/notice.proto

package favorite

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for NoticeService service

type NoticeService interface {
	AddOne(ctx context.Context, in *ReqNoticeAdd, opts ...client.CallOption) (*ReplyNoticeInfo, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyNoticeInfo, error)
	GetList(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyNoticeList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateBase(ctx context.Context, in *ReqNoticeUpdate, opts ...client.CallOption) (*ReplyNoticeInfo, error)
	UpdateStatus(ctx context.Context, in *RequestState, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateTags(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateTargets(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
}

type noticeService struct {
	c    client.Client
	name string
}

func NewNoticeService(name string, c client.Client) NoticeService {
	return &noticeService{
		c:    c,
		name: name,
	}
}

func (c *noticeService) AddOne(ctx context.Context, in *ReqNoticeAdd, opts ...client.CallOption) (*ReplyNoticeInfo, error) {
	req := c.c.NewRequest(c.name, "NoticeService.AddOne", in)
	out := new(ReplyNoticeInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyNoticeInfo, error) {
	req := c.c.NewRequest(c.name, "NoticeService.GetOne", in)
	out := new(ReplyNoticeInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeService) GetList(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyNoticeList, error) {
	req := c.c.NewRequest(c.name, "NoticeService.GetList", in)
	out := new(ReplyNoticeList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "NoticeService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "NoticeService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeService) UpdateBase(ctx context.Context, in *ReqNoticeUpdate, opts ...client.CallOption) (*ReplyNoticeInfo, error) {
	req := c.c.NewRequest(c.name, "NoticeService.UpdateBase", in)
	out := new(ReplyNoticeInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeService) UpdateStatus(ctx context.Context, in *RequestState, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "NoticeService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeService) UpdateTags(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "NoticeService.UpdateTags", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeService) UpdateTargets(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "NoticeService.UpdateTargets", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noticeService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "NoticeService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NoticeService service

type NoticeServiceHandler interface {
	AddOne(context.Context, *ReqNoticeAdd, *ReplyNoticeInfo) error
	GetOne(context.Context, *RequestInfo, *ReplyNoticeInfo) error
	GetList(context.Context, *RequestFilter, *ReplyNoticeList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	UpdateBase(context.Context, *ReqNoticeUpdate, *ReplyNoticeInfo) error
	UpdateStatus(context.Context, *RequestState, *ReplyInfo) error
	UpdateTags(context.Context, *RequestList, *ReplyInfo) error
	UpdateTargets(context.Context, *RequestList, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
}

func RegisterNoticeServiceHandler(s server.Server, hdlr NoticeServiceHandler, opts ...server.HandlerOption) error {
	type noticeService interface {
		AddOne(ctx context.Context, in *ReqNoticeAdd, out *ReplyNoticeInfo) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyNoticeInfo) error
		GetList(ctx context.Context, in *RequestFilter, out *ReplyNoticeList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		UpdateBase(ctx context.Context, in *ReqNoticeUpdate, out *ReplyNoticeInfo) error
		UpdateStatus(ctx context.Context, in *RequestState, out *ReplyInfo) error
		UpdateTags(ctx context.Context, in *RequestList, out *ReplyInfo) error
		UpdateTargets(ctx context.Context, in *RequestList, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
	}
	type NoticeService struct {
		noticeService
	}
	h := &noticeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&NoticeService{h}, opts...))
}

type noticeServiceHandler struct {
	NoticeServiceHandler
}

func (h *noticeServiceHandler) AddOne(ctx context.Context, in *ReqNoticeAdd, out *ReplyNoticeInfo) error {
	return h.NoticeServiceHandler.AddOne(ctx, in, out)
}

func (h *noticeServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyNoticeInfo) error {
	return h.NoticeServiceHandler.GetOne(ctx, in, out)
}

func (h *noticeServiceHandler) GetList(ctx context.Context, in *RequestFilter, out *ReplyNoticeList) error {
	return h.NoticeServiceHandler.GetList(ctx, in, out)
}

func (h *noticeServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.NoticeServiceHandler.GetStatistic(ctx, in, out)
}

func (h *noticeServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.NoticeServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *noticeServiceHandler) UpdateBase(ctx context.Context, in *ReqNoticeUpdate, out *ReplyNoticeInfo) error {
	return h.NoticeServiceHandler.UpdateBase(ctx, in, out)
}

func (h *noticeServiceHandler) UpdateStatus(ctx context.Context, in *RequestState, out *ReplyInfo) error {
	return h.NoticeServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *noticeServiceHandler) UpdateTags(ctx context.Context, in *RequestList, out *ReplyInfo) error {
	return h.NoticeServiceHandler.UpdateTags(ctx, in, out)
}

func (h *noticeServiceHandler) UpdateTargets(ctx context.Context, in *RequestList, out *ReplyInfo) error {
	return h.NoticeServiceHandler.UpdateTargets(ctx, in, out)
}

func (h *noticeServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.NoticeServiceHandler.RemoveOne(ctx, in, out)
}
