// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.4
// source: bizserver/v1/bizserver.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BizServer_SayHello_FullMethodName = "/bizserver.v1.BizServer/SayHello"
)

// BizServerClient is the client API for BizServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BizServerClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type bizServerClient struct {
	cc grpc.ClientConnInterface
}

func NewBizServerClient(cc grpc.ClientConnInterface) BizServerClient {
	return &bizServerClient{cc}
}

func (c *bizServerClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, BizServer_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BizServerServer is the server API for BizServer service.
// All implementations must embed UnimplementedBizServerServer
// for forward compatibility.
type BizServerServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedBizServerServer()
}

// UnimplementedBizServerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBizServerServer struct{}

func (UnimplementedBizServerServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedBizServerServer) mustEmbedUnimplementedBizServerServer() {}
func (UnimplementedBizServerServer) testEmbeddedByValue()                   {}

// UnsafeBizServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BizServerServer will
// result in compilation errors.
type UnsafeBizServerServer interface {
	mustEmbedUnimplementedBizServerServer()
}

func RegisterBizServerServer(s grpc.ServiceRegistrar, srv BizServerServer) {
	// If the following call pancis, it indicates UnimplementedBizServerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BizServer_ServiceDesc, srv)
}

func _BizServer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BizServer_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServerServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BizServer_ServiceDesc is the grpc.ServiceDesc for BizServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BizServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bizserver.v1.BizServer",
	HandlerType: (*BizServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _BizServer_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bizserver/v1/bizserver.proto",
}
