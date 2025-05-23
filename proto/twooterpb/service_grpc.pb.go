// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.3
// source: service.proto

package twooterpb

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

// TwooterClient is the client API for Twooter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwooterClient interface {
	// Creates a Twoot.
	CreateTwoot(ctx context.Context, in *CreateTwootRequest, opts ...grpc.CallOption) (*Twoot, error)
	// List all Twoots.
	ListTwoots(ctx context.Context, in *ListTwootsRequest, opts ...grpc.CallOption) (*ListTwootsResponse, error)
}

type twooterClient struct {
	cc grpc.ClientConnInterface
}

func NewTwooterClient(cc grpc.ClientConnInterface) TwooterClient {
	return &twooterClient{cc}
}

func (c *twooterClient) CreateTwoot(ctx context.Context, in *CreateTwootRequest, opts ...grpc.CallOption) (*Twoot, error) {
	out := new(Twoot)
	err := c.cc.Invoke(ctx, "/twooter.Twooter/CreateTwoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twooterClient) ListTwoots(ctx context.Context, in *ListTwootsRequest, opts ...grpc.CallOption) (*ListTwootsResponse, error) {
	out := new(ListTwootsResponse)
	err := c.cc.Invoke(ctx, "/twooter.Twooter/ListTwoots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwooterServer is the server API for Twooter service.
// All implementations must embed UnimplementedTwooterServer
// for forward compatibility
type TwooterServer interface {
	// Creates a Twoot.
	CreateTwoot(context.Context, *CreateTwootRequest) (*Twoot, error)
	// List all Twoots.
	ListTwoots(context.Context, *ListTwootsRequest) (*ListTwootsResponse, error)
	mustEmbedUnimplementedTwooterServer()
}

// UnimplementedTwooterServer must be embedded to have forward compatible implementations.
type UnimplementedTwooterServer struct {
}

func (UnimplementedTwooterServer) CreateTwoot(context.Context, *CreateTwootRequest) (*Twoot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTwoot not implemented")
}
func (UnimplementedTwooterServer) ListTwoots(context.Context, *ListTwootsRequest) (*ListTwootsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTwoots not implemented")
}
func (UnimplementedTwooterServer) mustEmbedUnimplementedTwooterServer() {}

// UnsafeTwooterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwooterServer will
// result in compilation errors.
type UnsafeTwooterServer interface {
	mustEmbedUnimplementedTwooterServer()
}

func RegisterTwooterServer(s grpc.ServiceRegistrar, srv TwooterServer) {
	s.RegisterService(&Twooter_ServiceDesc, srv)
}

func _Twooter_CreateTwoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTwootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwooterServer).CreateTwoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twooter.Twooter/CreateTwoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwooterServer).CreateTwoot(ctx, req.(*CreateTwootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twooter_ListTwoots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTwootsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwooterServer).ListTwoots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twooter.Twooter/ListTwoots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwooterServer).ListTwoots(ctx, req.(*ListTwootsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Twooter_ServiceDesc is the grpc.ServiceDesc for Twooter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Twooter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "twooter.Twooter",
	HandlerType: (*TwooterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTwoot",
			Handler:    _Twooter_CreateTwoot_Handler,
		},
		{
			MethodName: "ListTwoots",
			Handler:    _Twooter_ListTwoots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
