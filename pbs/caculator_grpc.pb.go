// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: caculator.proto

package pbs

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
	CaculatorService_Hello_FullMethodName           = "/protos.CaculatorService/Hello"
	CaculatorService_Sum_FullMethodName             = "/protos.CaculatorService/Sum"
	CaculatorService_SumWithDeadline_FullMethodName = "/protos.CaculatorService/SumWithDeadline"
	CaculatorService_ToPrimeNumber_FullMethodName   = "/protos.CaculatorService/ToPrimeNumber"
	CaculatorService_Average_FullMethodName         = "/protos.CaculatorService/Average"
	CaculatorService_FindMax_FullMethodName         = "/protos.CaculatorService/FindMax"
	CaculatorService_Square_FullMethodName          = "/protos.CaculatorService/Square"
)

// CaculatorServiceClient is the client API for CaculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaculatorServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	SumWithDeadline(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	ToPrimeNumber(ctx context.Context, in *ToPrimeNumberRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ToPrimeNumberResponse], error)
	Average(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AverageRequest, AverageResponse], error)
	FindMax(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[FindMaxRequest, FindMaxResponse], error)
	Square(ctx context.Context, in *SquareRequest, opts ...grpc.CallOption) (*SquareResponse, error)
}

type caculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaculatorServiceClient(cc grpc.ClientConnInterface) CaculatorServiceClient {
	return &caculatorServiceClient{cc}
}

func (c *caculatorServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, CaculatorService_Hello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, CaculatorService_Sum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caculatorServiceClient) SumWithDeadline(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, CaculatorService_SumWithDeadline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caculatorServiceClient) ToPrimeNumber(ctx context.Context, in *ToPrimeNumberRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ToPrimeNumberResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CaculatorService_ServiceDesc.Streams[0], CaculatorService_ToPrimeNumber_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ToPrimeNumberRequest, ToPrimeNumberResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CaculatorService_ToPrimeNumberClient = grpc.ServerStreamingClient[ToPrimeNumberResponse]

func (c *caculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[AverageRequest, AverageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CaculatorService_ServiceDesc.Streams[1], CaculatorService_Average_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AverageRequest, AverageResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CaculatorService_AverageClient = grpc.ClientStreamingClient[AverageRequest, AverageResponse]

func (c *caculatorServiceClient) FindMax(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[FindMaxRequest, FindMaxResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CaculatorService_ServiceDesc.Streams[2], CaculatorService_FindMax_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FindMaxRequest, FindMaxResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CaculatorService_FindMaxClient = grpc.BidiStreamingClient[FindMaxRequest, FindMaxResponse]

func (c *caculatorServiceClient) Square(ctx context.Context, in *SquareRequest, opts ...grpc.CallOption) (*SquareResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SquareResponse)
	err := c.cc.Invoke(ctx, CaculatorService_Square_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaculatorServiceServer is the server API for CaculatorService service.
// All implementations should embed UnimplementedCaculatorServiceServer
// for forward compatibility.
type CaculatorServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	SumWithDeadline(context.Context, *SumRequest) (*SumResponse, error)
	ToPrimeNumber(*ToPrimeNumberRequest, grpc.ServerStreamingServer[ToPrimeNumberResponse]) error
	Average(grpc.ClientStreamingServer[AverageRequest, AverageResponse]) error
	FindMax(grpc.BidiStreamingServer[FindMaxRequest, FindMaxResponse]) error
	Square(context.Context, *SquareRequest) (*SquareResponse, error)
}

// UnimplementedCaculatorServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCaculatorServiceServer struct{}

func (UnimplementedCaculatorServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedCaculatorServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCaculatorServiceServer) SumWithDeadline(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumWithDeadline not implemented")
}
func (UnimplementedCaculatorServiceServer) ToPrimeNumber(*ToPrimeNumberRequest, grpc.ServerStreamingServer[ToPrimeNumberResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ToPrimeNumber not implemented")
}
func (UnimplementedCaculatorServiceServer) Average(grpc.ClientStreamingServer[AverageRequest, AverageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedCaculatorServiceServer) FindMax(grpc.BidiStreamingServer[FindMaxRequest, FindMaxResponse]) error {
	return status.Errorf(codes.Unimplemented, "method FindMax not implemented")
}
func (UnimplementedCaculatorServiceServer) Square(context.Context, *SquareRequest) (*SquareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Square not implemented")
}
func (UnimplementedCaculatorServiceServer) testEmbeddedByValue() {}

// UnsafeCaculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaculatorServiceServer will
// result in compilation errors.
type UnsafeCaculatorServiceServer interface {
	mustEmbedUnimplementedCaculatorServiceServer()
}

func RegisterCaculatorServiceServer(s grpc.ServiceRegistrar, srv CaculatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedCaculatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CaculatorService_ServiceDesc, srv)
}

func _CaculatorService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaculatorServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaculatorService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaculatorServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaculatorService_Sum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaculatorService_SumWithDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaculatorServiceServer).SumWithDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaculatorService_SumWithDeadline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaculatorServiceServer).SumWithDeadline(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaculatorService_ToPrimeNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ToPrimeNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CaculatorServiceServer).ToPrimeNumber(m, &grpc.GenericServerStream[ToPrimeNumberRequest, ToPrimeNumberResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CaculatorService_ToPrimeNumberServer = grpc.ServerStreamingServer[ToPrimeNumberResponse]

func _CaculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CaculatorServiceServer).Average(&grpc.GenericServerStream[AverageRequest, AverageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CaculatorService_AverageServer = grpc.ClientStreamingServer[AverageRequest, AverageResponse]

func _CaculatorService_FindMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CaculatorServiceServer).FindMax(&grpc.GenericServerStream[FindMaxRequest, FindMaxResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CaculatorService_FindMaxServer = grpc.BidiStreamingServer[FindMaxRequest, FindMaxResponse]

func _CaculatorService_Square_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaculatorServiceServer).Square(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaculatorService_Square_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaculatorServiceServer).Square(ctx, req.(*SquareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CaculatorService_ServiceDesc is the grpc.ServiceDesc for CaculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.CaculatorService",
	HandlerType: (*CaculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _CaculatorService_Hello_Handler,
		},
		{
			MethodName: "Sum",
			Handler:    _CaculatorService_Sum_Handler,
		},
		{
			MethodName: "SumWithDeadline",
			Handler:    _CaculatorService_SumWithDeadline_Handler,
		},
		{
			MethodName: "Square",
			Handler:    _CaculatorService_Square_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ToPrimeNumber",
			Handler:       _CaculatorService_ToPrimeNumber_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CaculatorService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMax",
			Handler:       _CaculatorService_FindMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "caculator.proto",
}
