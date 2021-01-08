// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/favorite/repertory.proto

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

// Client API for RepertoryService service

type RepertoryService interface {
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRepertoryInfo, error)
	AppendOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRepertoryInfo, error)
	SubtractOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRepertoryInfo, error)
	UpdateList(ctx context.Context, in *ReqRepertoryBags, opts ...client.CallOption) (*ReplyRepertoryInfo, error)
}

type repertoryService struct {
	c    client.Client
	name string
}

func NewRepertoryService(name string, c client.Client) RepertoryService {
	return &repertoryService{
		c:    c,
		name: name,
	}
}

func (c *repertoryService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRepertoryInfo, error) {
	req := c.c.NewRequest(c.name, "RepertoryService.GetOne", in)
	out := new(ReplyRepertoryInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repertoryService) AppendOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRepertoryInfo, error) {
	req := c.c.NewRequest(c.name, "RepertoryService.AppendOne", in)
	out := new(ReplyRepertoryInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repertoryService) SubtractOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyRepertoryInfo, error) {
	req := c.c.NewRequest(c.name, "RepertoryService.SubtractOne", in)
	out := new(ReplyRepertoryInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repertoryService) UpdateList(ctx context.Context, in *ReqRepertoryBags, opts ...client.CallOption) (*ReplyRepertoryInfo, error) {
	req := c.c.NewRequest(c.name, "RepertoryService.UpdateList", in)
	out := new(ReplyRepertoryInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RepertoryService service

type RepertoryServiceHandler interface {
	GetOne(context.Context, *RequestInfo, *ReplyRepertoryInfo) error
	AppendOne(context.Context, *RequestInfo, *ReplyRepertoryInfo) error
	SubtractOne(context.Context, *RequestInfo, *ReplyRepertoryInfo) error
	UpdateList(context.Context, *ReqRepertoryBags, *ReplyRepertoryInfo) error
}

func RegisterRepertoryServiceHandler(s server.Server, hdlr RepertoryServiceHandler, opts ...server.HandlerOption) error {
	type repertoryService interface {
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyRepertoryInfo) error
		AppendOne(ctx context.Context, in *RequestInfo, out *ReplyRepertoryInfo) error
		SubtractOne(ctx context.Context, in *RequestInfo, out *ReplyRepertoryInfo) error
		UpdateList(ctx context.Context, in *ReqRepertoryBags, out *ReplyRepertoryInfo) error
	}
	type RepertoryService struct {
		repertoryService
	}
	h := &repertoryServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RepertoryService{h}, opts...))
}

type repertoryServiceHandler struct {
	RepertoryServiceHandler
}

func (h *repertoryServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyRepertoryInfo) error {
	return h.RepertoryServiceHandler.GetOne(ctx, in, out)
}

func (h *repertoryServiceHandler) AppendOne(ctx context.Context, in *RequestInfo, out *ReplyRepertoryInfo) error {
	return h.RepertoryServiceHandler.AppendOne(ctx, in, out)
}

func (h *repertoryServiceHandler) SubtractOne(ctx context.Context, in *RequestInfo, out *ReplyRepertoryInfo) error {
	return h.RepertoryServiceHandler.SubtractOne(ctx, in, out)
}

func (h *repertoryServiceHandler) UpdateList(ctx context.Context, in *ReqRepertoryBags, out *ReplyRepertoryInfo) error {
	return h.RepertoryServiceHandler.UpdateList(ctx, in, out)
}
