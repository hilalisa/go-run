// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/go-run/proto/run.proto

/*
Package go_micro_run is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-run/proto/run.proto

It has these top-level messages:
	Source
	Binary
	Process
	FetchRequest
	FetchResponse
	BuildRequest
	BuildResponse
	ExecRequest
	ExecResponse
	KillRequest
	KillResponse
	WaitRequest
	WaitResponse
	RunRequest
	RunResponse
	StopRequest
	StopResponse
	StatusRequest
	StatusResponse
*/
package go_micro_run

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Runtime service

type RuntimeService interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...client.CallOption) (*FetchResponse, error)
	Build(ctx context.Context, in *BuildRequest, opts ...client.CallOption) (*BuildResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error)
	Kill(ctx context.Context, in *KillRequest, opts ...client.CallOption) (*KillResponse, error)
	Wait(ctx context.Context, in *WaitRequest, opts ...client.CallOption) (Runtime_WaitService, error)
}

type runtimeService struct {
	c    client.Client
	name string
}

func NewRuntimeService(name string, c client.Client) RuntimeService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.run"
	}
	return &runtimeService{
		c:    c,
		name: name,
	}
}

func (c *runtimeService) Fetch(ctx context.Context, in *FetchRequest, opts ...client.CallOption) (*FetchResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Fetch", in)
	out := new(FetchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Build(ctx context.Context, in *BuildRequest, opts ...client.CallOption) (*BuildResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Build", in)
	out := new(BuildResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Exec(ctx context.Context, in *ExecRequest, opts ...client.CallOption) (*ExecResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Exec", in)
	out := new(ExecResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Kill(ctx context.Context, in *KillRequest, opts ...client.CallOption) (*KillResponse, error) {
	req := c.c.NewRequest(c.name, "Runtime.Kill", in)
	out := new(KillResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runtimeService) Wait(ctx context.Context, in *WaitRequest, opts ...client.CallOption) (Runtime_WaitService, error) {
	req := c.c.NewRequest(c.name, "Runtime.Wait", &WaitRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &runtimeServiceWait{stream}, nil
}

type Runtime_WaitService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*WaitResponse, error)
}

type runtimeServiceWait struct {
	stream client.Stream
}

func (x *runtimeServiceWait) Close() error {
	return x.stream.Close()
}

func (x *runtimeServiceWait) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeServiceWait) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeServiceWait) Recv() (*WaitResponse, error) {
	m := new(WaitResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Runtime service

type RuntimeHandler interface {
	Fetch(context.Context, *FetchRequest, *FetchResponse) error
	Build(context.Context, *BuildRequest, *BuildResponse) error
	Exec(context.Context, *ExecRequest, *ExecResponse) error
	Kill(context.Context, *KillRequest, *KillResponse) error
	Wait(context.Context, *WaitRequest, Runtime_WaitStream) error
}

func RegisterRuntimeHandler(s server.Server, hdlr RuntimeHandler, opts ...server.HandlerOption) {
	type runtime interface {
		Fetch(ctx context.Context, in *FetchRequest, out *FetchResponse) error
		Build(ctx context.Context, in *BuildRequest, out *BuildResponse) error
		Exec(ctx context.Context, in *ExecRequest, out *ExecResponse) error
		Kill(ctx context.Context, in *KillRequest, out *KillResponse) error
		Wait(ctx context.Context, stream server.Stream) error
	}
	type Runtime struct {
		runtime
	}
	h := &runtimeHandler{hdlr}
	s.Handle(s.NewHandler(&Runtime{h}, opts...))
}

type runtimeHandler struct {
	RuntimeHandler
}

func (h *runtimeHandler) Fetch(ctx context.Context, in *FetchRequest, out *FetchResponse) error {
	return h.RuntimeHandler.Fetch(ctx, in, out)
}

func (h *runtimeHandler) Build(ctx context.Context, in *BuildRequest, out *BuildResponse) error {
	return h.RuntimeHandler.Build(ctx, in, out)
}

func (h *runtimeHandler) Exec(ctx context.Context, in *ExecRequest, out *ExecResponse) error {
	return h.RuntimeHandler.Exec(ctx, in, out)
}

func (h *runtimeHandler) Kill(ctx context.Context, in *KillRequest, out *KillResponse) error {
	return h.RuntimeHandler.Kill(ctx, in, out)
}

func (h *runtimeHandler) Wait(ctx context.Context, stream server.Stream) error {
	m := new(WaitRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RuntimeHandler.Wait(ctx, m, &runtimeWaitStream{stream})
}

type Runtime_WaitStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*WaitResponse) error
}

type runtimeWaitStream struct {
	stream server.Stream
}

func (x *runtimeWaitStream) Close() error {
	return x.stream.Close()
}

func (x *runtimeWaitStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runtimeWaitStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runtimeWaitStream) Send(m *WaitResponse) error {
	return x.stream.Send(m)
}

// Client API for Service service

type Service interface {
	Run(ctx context.Context, in *RunRequest, opts ...client.CallOption) (*RunResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...client.CallOption) (*StopResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error)
}

type service struct {
	c    client.Client
	name string
}

func NewService(name string, c client.Client) Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.run"
	}
	return &service{
		c:    c,
		name: name,
	}
}

func (c *service) Run(ctx context.Context, in *RunRequest, opts ...client.CallOption) (*RunResponse, error) {
	req := c.c.NewRequest(c.name, "Service.Run", in)
	out := new(RunResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) Stop(ctx context.Context, in *StopRequest, opts ...client.CallOption) (*StopResponse, error) {
	req := c.c.NewRequest(c.name, "Service.Stop", in)
	out := new(StopResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) Status(ctx context.Context, in *StatusRequest, opts ...client.CallOption) (*StatusResponse, error) {
	req := c.c.NewRequest(c.name, "Service.Status", in)
	out := new(StatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceHandler interface {
	Run(context.Context, *RunRequest, *RunResponse) error
	Stop(context.Context, *StopRequest, *StopResponse) error
	Status(context.Context, *StatusRequest, *StatusResponse) error
}

func RegisterServiceHandler(s server.Server, hdlr ServiceHandler, opts ...server.HandlerOption) {
	type service interface {
		Run(ctx context.Context, in *RunRequest, out *RunResponse) error
		Stop(ctx context.Context, in *StopRequest, out *StopResponse) error
		Status(ctx context.Context, in *StatusRequest, out *StatusResponse) error
	}
	type Service struct {
		service
	}
	h := &serviceHandler{hdlr}
	s.Handle(s.NewHandler(&Service{h}, opts...))
}

type serviceHandler struct {
	ServiceHandler
}

func (h *serviceHandler) Run(ctx context.Context, in *RunRequest, out *RunResponse) error {
	return h.ServiceHandler.Run(ctx, in, out)
}

func (h *serviceHandler) Stop(ctx context.Context, in *StopRequest, out *StopResponse) error {
	return h.ServiceHandler.Stop(ctx, in, out)
}

func (h *serviceHandler) Status(ctx context.Context, in *StatusRequest, out *StatusResponse) error {
	return h.ServiceHandler.Status(ctx, in, out)
}
