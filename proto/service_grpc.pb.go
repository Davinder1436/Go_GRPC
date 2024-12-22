// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/service.proto

package proto

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
	ExampleService_UnaryCall_FullMethodName                  = "/example.ExampleService/UnaryCall"
	ExampleService_ClientStreamingCall_FullMethodName        = "/example.ExampleService/ClientStreamingCall"
	ExampleService_ServerStreamingCall_FullMethodName        = "/example.ExampleService/ServerStreamingCall"
	ExampleService_BidirectionalStreamingCall_FullMethodName = "/example.ExampleService/BidirectionalStreamingCall"
)

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	UnaryCall(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error)
	ClientStreamingCall(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Input, Output], error)
	ServerStreamingCall(ctx context.Context, in *Input, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Output], error)
	BidirectionalStreamingCall(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Input, Output], error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) UnaryCall(ctx context.Context, in *Input, opts ...grpc.CallOption) (*Output, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Output)
	err := c.cc.Invoke(ctx, ExampleService_UnaryCall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) ClientStreamingCall(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[Input, Output], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[0], ExampleService_ClientStreamingCall_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Input, Output]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_ClientStreamingCallClient = grpc.ClientStreamingClient[Input, Output]

func (c *exampleServiceClient) ServerStreamingCall(ctx context.Context, in *Input, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Output], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[1], ExampleService_ServerStreamingCall_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Input, Output]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_ServerStreamingCallClient = grpc.ServerStreamingClient[Output]

func (c *exampleServiceClient) BidirectionalStreamingCall(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Input, Output], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[2], ExampleService_BidirectionalStreamingCall_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Input, Output]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_BidirectionalStreamingCallClient = grpc.BidiStreamingClient[Input, Output]

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility.
type ExampleServiceServer interface {
	UnaryCall(context.Context, *Input) (*Output, error)
	ClientStreamingCall(grpc.ClientStreamingServer[Input, Output]) error
	ServerStreamingCall(*Input, grpc.ServerStreamingServer[Output]) error
	BidirectionalStreamingCall(grpc.BidiStreamingServer[Input, Output]) error
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExampleServiceServer struct{}

func (UnimplementedExampleServiceServer) UnaryCall(context.Context, *Input) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryCall not implemented")
}
func (UnimplementedExampleServiceServer) ClientStreamingCall(grpc.ClientStreamingServer[Input, Output]) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamingCall not implemented")
}
func (UnimplementedExampleServiceServer) ServerStreamingCall(*Input, grpc.ServerStreamingServer[Output]) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingCall not implemented")
}
func (UnimplementedExampleServiceServer) BidirectionalStreamingCall(grpc.BidiStreamingServer[Input, Output]) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamingCall not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}
func (UnimplementedExampleServiceServer) testEmbeddedByValue()                        {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	// If the following call pancis, it indicates UnimplementedExampleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_UnaryCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).UnaryCall(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_ClientStreamingCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).ClientStreamingCall(&grpc.GenericServerStream[Input, Output]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_ClientStreamingCallServer = grpc.ClientStreamingServer[Input, Output]

func _ExampleService_ServerStreamingCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Input)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServiceServer).ServerStreamingCall(m, &grpc.GenericServerStream[Input, Output]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_ServerStreamingCallServer = grpc.ServerStreamingServer[Output]

func _ExampleService_BidirectionalStreamingCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).BidirectionalStreamingCall(&grpc.GenericServerStream[Input, Output]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ExampleService_BidirectionalStreamingCallServer = grpc.BidiStreamingServer[Input, Output]

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _ExampleService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStreamingCall",
			Handler:       _ExampleService_ClientStreamingCall_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStreamingCall",
			Handler:       _ExampleService_ServerStreamingCall_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamingCall",
			Handler:       _ExampleService_BidirectionalStreamingCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}