// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chatgrpc

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

// ChatterInterfaceClient is the client API for ChatterInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatterInterfaceClient interface {
	ChatSender(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Status, error)
	ChatListener(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Status, error)
}

type chatterInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatterInterfaceClient(cc grpc.ClientConnInterface) ChatterInterfaceClient {
	return &chatterInterfaceClient{cc}
}

func (c *chatterInterfaceClient) ChatSender(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/chatgrpc.ChatterInterface/ChatSender", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatterInterfaceClient) ChatListener(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/chatgrpc.ChatterInterface/ChatListener", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatterInterfaceServer is the server API for ChatterInterface service.
// All implementations must embed UnimplementedChatterInterfaceServer
// for forward compatibility
type ChatterInterfaceServer interface {
	ChatSender(context.Context, *Msg) (*Status, error)
	ChatListener(context.Context, *Msg) (*Status, error)
}

// UnimplementedChatterInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedChatterInterfaceServer struct {
}

func (UnimplementedChatterInterfaceServer) ChatSender(context.Context, *Msg) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatSender not implemented")
}
func (UnimplementedChatterInterfaceServer) ChatListener(context.Context, *Msg) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatListener not implemented")
}

func RegisterChatterInterfaceServer(s grpc.ServiceRegistrar, srv ChatterInterfaceServer) {
	s.RegisterService(&ChatterInterface_ServiceDesc, srv)
}

func _ChatterInterface_ChatSender_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatterInterfaceServer).ChatSender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatgrpc.ChatterInterface/ChatSender",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatterInterfaceServer).ChatSender(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatterInterface_ChatListener_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatterInterfaceServer).ChatListener(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatgrpc.ChatterInterface/ChatListener",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatterInterfaceServer).ChatListener(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatterInterface_ServiceDesc is the grpc.ServiceDesc for ChatterInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatterInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatgrpc.ChatterInterface",
	HandlerType: (*ChatterInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChatSender",
			Handler:    _ChatterInterface_ChatSender_Handler,
		},
		{
			MethodName: "ChatListener",
			Handler:    _ChatterInterface_ChatListener_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatgrpc/msg.proto",
}
