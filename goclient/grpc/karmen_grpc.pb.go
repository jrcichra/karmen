// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package karmen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KarmenClient is the client API for Karmen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KarmenClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	EmitEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	ActionDispatcher(ctx context.Context, opts ...grpc.CallOption) (Karmen_ActionDispatcherClient, error)
	PingPong(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type karmenClient struct {
	cc grpc.ClientConnInterface
}

func NewKarmenClient(cc grpc.ClientConnInterface) KarmenClient {
	return &karmenClient{cc}
}

func (c *karmenClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/Karmen/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karmenClient) EmitEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/Karmen/EmitEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *karmenClient) ActionDispatcher(ctx context.Context, opts ...grpc.CallOption) (Karmen_ActionDispatcherClient, error) {
	stream, err := c.cc.NewStream(ctx, &Karmen_ServiceDesc.Streams[0], "/Karmen/ActionDispatcher", opts...)
	if err != nil {
		return nil, err
	}
	x := &karmenActionDispatcherClient{stream}
	return x, nil
}

type Karmen_ActionDispatcherClient interface {
	Send(*ActionResponse) error
	Recv() (*ActionRequest, error)
	grpc.ClientStream
}

type karmenActionDispatcherClient struct {
	grpc.ClientStream
}

func (x *karmenActionDispatcherClient) Send(m *ActionResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *karmenActionDispatcherClient) Recv() (*ActionRequest, error) {
	m := new(ActionRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *karmenClient) PingPong(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/Karmen/PingPong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KarmenServer is the server API for Karmen service.
// All implementations must embed UnimplementedKarmenServer
// for forward compatibility
type KarmenServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	EmitEvent(context.Context, *EventRequest) (*EventResponse, error)
	ActionDispatcher(Karmen_ActionDispatcherServer) error
	PingPong(context.Context, *Ping) (*Pong, error)
	mustEmbedUnimplementedKarmenServer()
}

// UnimplementedKarmenServer must be embedded to have forward compatible implementations.
type UnimplementedKarmenServer struct {
}

func (UnimplementedKarmenServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedKarmenServer) EmitEvent(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitEvent not implemented")
}
func (UnimplementedKarmenServer) ActionDispatcher(Karmen_ActionDispatcherServer) error {
	return status.Errorf(codes.Unimplemented, "method ActionDispatcher not implemented")
}
func (UnimplementedKarmenServer) PingPong(context.Context, *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}
func (UnimplementedKarmenServer) mustEmbedUnimplementedKarmenServer() {}

// UnsafeKarmenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KarmenServer will
// result in compilation errors.
type UnsafeKarmenServer interface {
	mustEmbedUnimplementedKarmenServer()
}

func RegisterKarmenServer(s grpc.ServiceRegistrar, srv KarmenServer) {
	s.RegisterService(&Karmen_ServiceDesc, srv)
}

func _Karmen_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KarmenServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Karmen/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KarmenServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Karmen_EmitEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KarmenServer).EmitEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Karmen/EmitEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KarmenServer).EmitEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Karmen_ActionDispatcher_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KarmenServer).ActionDispatcher(&karmenActionDispatcherServer{stream})
}

type Karmen_ActionDispatcherServer interface {
	Send(*ActionRequest) error
	Recv() (*ActionResponse, error)
	grpc.ServerStream
}

type karmenActionDispatcherServer struct {
	grpc.ServerStream
}

func (x *karmenActionDispatcherServer) Send(m *ActionRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *karmenActionDispatcherServer) Recv() (*ActionResponse, error) {
	m := new(ActionResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Karmen_PingPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KarmenServer).PingPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Karmen/PingPong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KarmenServer).PingPong(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

// Karmen_ServiceDesc is the grpc.ServiceDesc for Karmen service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Karmen_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Karmen",
	HandlerType: (*KarmenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Karmen_Register_Handler,
		},
		{
			MethodName: "EmitEvent",
			Handler:    _Karmen_EmitEvent_Handler,
		},
		{
			MethodName: "PingPong",
			Handler:    _Karmen_PingPong_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ActionDispatcher",
			Handler:       _Karmen_ActionDispatcher_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/karmen.proto",
}
