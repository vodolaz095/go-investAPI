// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: signals.proto

package investapi

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
	SignalService_GetStrategies_FullMethodName = "/tinkoff.public.invest.api.contract.v1.SignalService/GetStrategies"
	SignalService_GetSignals_FullMethodName    = "/tinkoff.public.invest.api.contract.v1.SignalService/GetSignals"
)

// SignalServiceClient is the client API for SignalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignalServiceClient interface {
	// GetStrategies — стратегии
	GetStrategies(ctx context.Context, in *GetStrategiesRequest, opts ...grpc.CallOption) (*GetStrategiesResponse, error)
	// GetSignals — сигналы
	GetSignals(ctx context.Context, in *GetSignalsRequest, opts ...grpc.CallOption) (*GetSignalsResponse, error)
}

type signalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignalServiceClient(cc grpc.ClientConnInterface) SignalServiceClient {
	return &signalServiceClient{cc}
}

func (c *signalServiceClient) GetStrategies(ctx context.Context, in *GetStrategiesRequest, opts ...grpc.CallOption) (*GetStrategiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStrategiesResponse)
	err := c.cc.Invoke(ctx, SignalService_GetStrategies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) GetSignals(ctx context.Context, in *GetSignalsRequest, opts ...grpc.CallOption) (*GetSignalsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSignalsResponse)
	err := c.cc.Invoke(ctx, SignalService_GetSignals_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignalServiceServer is the server API for SignalService service.
// All implementations must embed UnimplementedSignalServiceServer
// for forward compatibility.
type SignalServiceServer interface {
	// GetStrategies — стратегии
	GetStrategies(context.Context, *GetStrategiesRequest) (*GetStrategiesResponse, error)
	// GetSignals — сигналы
	GetSignals(context.Context, *GetSignalsRequest) (*GetSignalsResponse, error)
	mustEmbedUnimplementedSignalServiceServer()
}

// UnimplementedSignalServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSignalServiceServer struct{}

func (UnimplementedSignalServiceServer) GetStrategies(context.Context, *GetStrategiesRequest) (*GetStrategiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStrategies not implemented")
}
func (UnimplementedSignalServiceServer) GetSignals(context.Context, *GetSignalsRequest) (*GetSignalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignals not implemented")
}
func (UnimplementedSignalServiceServer) mustEmbedUnimplementedSignalServiceServer() {}
func (UnimplementedSignalServiceServer) testEmbeddedByValue()                       {}

// UnsafeSignalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignalServiceServer will
// result in compilation errors.
type UnsafeSignalServiceServer interface {
	mustEmbedUnimplementedSignalServiceServer()
}

func RegisterSignalServiceServer(s grpc.ServiceRegistrar, srv SignalServiceServer) {
	// If the following call pancis, it indicates UnimplementedSignalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SignalService_ServiceDesc, srv)
}

func _SignalService_GetStrategies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStrategiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).GetStrategies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignalService_GetStrategies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).GetStrategies(ctx, req.(*GetStrategiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_GetSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).GetSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SignalService_GetSignals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).GetSignals(ctx, req.(*GetSignalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignalService_ServiceDesc is the grpc.ServiceDesc for SignalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.SignalService",
	HandlerType: (*SignalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStrategies",
			Handler:    _SignalService_GetStrategies_Handler,
		},
		{
			MethodName: "GetSignals",
			Handler:    _SignalService_GetSignals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signals.proto",
}
