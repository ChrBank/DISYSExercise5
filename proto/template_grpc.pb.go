// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/template.proto

package proto

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

// TimeSyncClient is the client API for TimeSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeSyncClient interface {
	Ping(ctx context.Context, in *Ack, opts ...grpc.CallOption) (*Ack, error)
	GetTime(ctx context.Context, in *Ack, opts ...grpc.CallOption) (*Time, error)
}

type timeSyncClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeSyncClient(cc grpc.ClientConnInterface) TimeSyncClient {
	return &timeSyncClient{cc}
}

func (c *timeSyncClient) Ping(ctx context.Context, in *Ack, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/proto.TimeSync/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeSyncClient) GetTime(ctx context.Context, in *Ack, opts ...grpc.CallOption) (*Time, error) {
	out := new(Time)
	err := c.cc.Invoke(ctx, "/proto.TimeSync/GetTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeSyncServer is the server API for TimeSync service.
// All implementations must embed UnimplementedTimeSyncServer
// for forward compatibility
type TimeSyncServer interface {
	Ping(context.Context, *Ack) (*Ack, error)
	GetTime(context.Context, *Ack) (*Time, error)
	mustEmbedUnimplementedTimeSyncServer()
}

// UnimplementedTimeSyncServer must be embedded to have forward compatible implementations.
type UnimplementedTimeSyncServer struct {
}

func (UnimplementedTimeSyncServer) Ping(context.Context, *Ack) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedTimeSyncServer) GetTime(context.Context, *Ack) (*Time, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTime not implemented")
}
func (UnimplementedTimeSyncServer) mustEmbedUnimplementedTimeSyncServer() {}

// UnsafeTimeSyncServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeSyncServer will
// result in compilation errors.
type UnsafeTimeSyncServer interface {
	mustEmbedUnimplementedTimeSyncServer()
}

func RegisterTimeSyncServer(s grpc.ServiceRegistrar, srv TimeSyncServer) {
	s.RegisterService(&TimeSync_ServiceDesc, srv)
}

func _TimeSync_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ack)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSyncServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TimeSync/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSyncServer).Ping(ctx, req.(*Ack))
	}
	return interceptor(ctx, in, info, handler)
}

func _TimeSync_GetTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ack)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSyncServer).GetTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TimeSync/GetTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSyncServer).GetTime(ctx, req.(*Ack))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeSync_ServiceDesc is the grpc.ServiceDesc for TimeSync service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeSync_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TimeSync",
	HandlerType: (*TimeSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _TimeSync_Ping_Handler,
		},
		{
			MethodName: "GetTime",
			Handler:    _TimeSync_GetTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/template.proto",
}
