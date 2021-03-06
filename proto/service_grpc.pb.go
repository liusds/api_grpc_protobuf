// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// CreateServerClient is the client API for CreateServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateServerClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type createServerClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateServerClient(cc grpc.ClientConnInterface) CreateServerClient {
	return &createServerClient{cc}
}

func (c *createServerClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CreateServer/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *createServerClient) Multiply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.CreateServer/Multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateServerServer is the server API for CreateServer service.
// All implementations must embed UnimplementedCreateServerServer
// for forward compatibility
type CreateServerServer interface {
	Add(context.Context, *Request) (*Response, error)
	Multiply(context.Context, *Request) (*Response, error)
	// mustEmbedUnimplementedCreateServerServer()
}

// UnimplementedCreateServerServer must be embedded to have forward compatible implementations.
type UnimplementedCreateServerServer struct {
}

func (UnimplementedCreateServerServer) Add(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCreateServerServer) Multiply(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedCreateServerServer) mustEmbedUnimplementedCreateServerServer() {}

// UnsafeCreateServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateServerServer will
// result in compilation errors.
type UnsafeCreateServerServer interface {
	mustEmbedUnimplementedCreateServerServer()
}

func RegisterCreateServerServer(s grpc.ServiceRegistrar, srv CreateServerServer) {
	s.RegisterService(&CreateServer_ServiceDesc, srv)
}

func _CreateServer_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateServerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CreateServer/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateServerServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreateServer_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateServerServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CreateServer/Multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateServerServer).Multiply(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateServer_ServiceDesc is the grpc.ServiceDesc for CreateServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CreateServer",
	HandlerType: (*CreateServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CreateServer_Add_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _CreateServer_Multiply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
