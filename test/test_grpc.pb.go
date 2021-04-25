// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package test

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	TestQ(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*OutHello, error)
	TestM(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) TestQ(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*OutHello, error) {
	out := new(OutHello)
	err := c.cc.Invoke(ctx, "/test.Service/testQ", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TestM(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error) {
	out := new(Hello)
	err := c.cc.Invoke(ctx, "/test.Service/testM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations should embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	TestQ(context.Context, *Hello) (*OutHello, error)
	TestM(context.Context, *Hello) (*Hello, error)
}

// UnimplementedServiceServer should be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) TestQ(context.Context, *Hello) (*OutHello, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestQ not implemented")
}
func (UnimplementedServiceServer) TestM(context.Context, *Hello) (*Hello, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestM not implemented")
}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_TestQ_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TestQ(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Service/testQ",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TestQ(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_TestM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TestM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Service/testM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TestM(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "testQ",
			Handler:    _Service_TestQ_Handler,
		},
		{
			MethodName: "testM",
			Handler:    _Service_TestM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test/test.proto",
}

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	Query1(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error)
	Query2(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error)
	Mutate1(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error)
	Subscribe(ctx context.Context, in *Hello, opts ...grpc.CallOption) (Query_SubscribeClient, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Query1(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error) {
	out := new(Hello)
	err := c.cc.Invoke(ctx, "/test.Query/Query1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Query2(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error) {
	out := new(Hello)
	err := c.cc.Invoke(ctx, "/test.Query/Query2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Mutate1(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error) {
	out := new(Hello)
	err := c.cc.Invoke(ctx, "/test.Query/Mutate1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Subscribe(ctx context.Context, in *Hello, opts ...grpc.CallOption) (Query_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Query_ServiceDesc.Streams[0], "/test.Query/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &querySubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_SubscribeClient interface {
	Recv() (*Hello, error)
	grpc.ClientStream
}

type querySubscribeClient struct {
	grpc.ClientStream
}

func (x *querySubscribeClient) Recv() (*Hello, error) {
	m := new(Hello)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServer is the server API for Query service.
// All implementations should embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	Query1(context.Context, *Hello) (*Hello, error)
	Query2(context.Context, *Hello) (*Hello, error)
	Mutate1(context.Context, *Hello) (*Hello, error)
	Subscribe(*Hello, Query_SubscribeServer) error
}

// UnimplementedQueryServer should be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Query1(context.Context, *Hello) (*Hello, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query1 not implemented")
}
func (UnimplementedQueryServer) Query2(context.Context, *Hello) (*Hello, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query2 not implemented")
}
func (UnimplementedQueryServer) Mutate1(context.Context, *Hello) (*Hello, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mutate1 not implemented")
}
func (UnimplementedQueryServer) Subscribe(*Hello, Query_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Query1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Query1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Query/Query1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Query1(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Query2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Query2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Query/Query2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Query2(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Mutate1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Mutate1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.Query/Mutate1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Mutate1(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Hello)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).Subscribe(m, &querySubscribeServer{stream})
}

type Query_SubscribeServer interface {
	Send(*Hello) error
	grpc.ServerStream
}

type querySubscribeServer struct {
	grpc.ServerStream
}

func (x *querySubscribeServer) Send(m *Hello) error {
	return x.ServerStream.SendMsg(m)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "test.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query1",
			Handler:    _Query_Query1_Handler,
		},
		{
			MethodName: "Query2",
			Handler:    _Query_Query2_Handler,
		},
		{
			MethodName: "Mutate1",
			Handler:    _Query_Mutate1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Query_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test/test.proto",
}
