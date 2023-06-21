// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: countPores.proto

package grpc

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

// CountPoresClient is the client API for CountPores service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountPoresClient interface {
	Count(ctx context.Context, in *CountPoresRequest, opts ...grpc.CallOption) (*CountPoresResponse, error)
}

type countPoresClient struct {
	cc grpc.ClientConnInterface
}

func NewCountPoresClient(cc grpc.ClientConnInterface) CountPoresClient {
	return &countPoresClient{cc}
}

func (c *countPoresClient) Count(ctx context.Context, in *CountPoresRequest, opts ...grpc.CallOption) (*CountPoresResponse, error) {
	out := new(CountPoresResponse)
	err := c.cc.Invoke(ctx, "/CountPores/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountPoresServer is the server API for CountPores service.
// All implementations must embed UnimplementedCountPoresServer
// for forward compatibility
type CountPoresServer interface {
	Count(context.Context, *CountPoresRequest) (*CountPoresResponse, error)
	mustEmbedUnimplementedCountPoresServer()
}

// UnimplementedCountPoresServer must be embedded to have forward compatible implementations.
type UnimplementedCountPoresServer struct {
}

func (UnimplementedCountPoresServer) Count(context.Context, *CountPoresRequest) (*CountPoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedCountPoresServer) mustEmbedUnimplementedCountPoresServer() {}

// UnsafeCountPoresServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountPoresServer will
// result in compilation errors.
type UnsafeCountPoresServer interface {
	mustEmbedUnimplementedCountPoresServer()
}

func RegisterCountPoresServer(s grpc.ServiceRegistrar, srv CountPoresServer) {
	s.RegisterService(&CountPores_ServiceDesc, srv)
}

func _CountPores_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountPoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountPoresServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CountPores/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountPoresServer).Count(ctx, req.(*CountPoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CountPores_ServiceDesc is the grpc.ServiceDesc for CountPores service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CountPores_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CountPores",
	HandlerType: (*CountPoresServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Count",
			Handler:    _CountPores_Count_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "countPores.proto",
}