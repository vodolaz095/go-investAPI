// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package investapi

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

// StopOrdersServiceClient is the client API for StopOrdersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StopOrdersServiceClient interface {
	// Метод выставления стоп-заявки.
	PostStopOrder(ctx context.Context, in *PostStopOrderRequest, opts ...grpc.CallOption) (*PostStopOrderResponse, error)
	// Метод получения списка активных стоп заявок по счёту.
	GetStopOrders(ctx context.Context, in *GetStopOrdersRequest, opts ...grpc.CallOption) (*GetStopOrdersResponse, error)
	// Метод отмены стоп-заявки.
	CancelStopOrder(ctx context.Context, in *CancelStopOrderRequest, opts ...grpc.CallOption) (*CancelStopOrderResponse, error)
}

type stopOrdersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStopOrdersServiceClient(cc grpc.ClientConnInterface) StopOrdersServiceClient {
	return &stopOrdersServiceClient{cc}
}

func (c *stopOrdersServiceClient) PostStopOrder(ctx context.Context, in *PostStopOrderRequest, opts ...grpc.CallOption) (*PostStopOrderResponse, error) {
	out := new(PostStopOrderResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.StopOrdersService/PostStopOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stopOrdersServiceClient) GetStopOrders(ctx context.Context, in *GetStopOrdersRequest, opts ...grpc.CallOption) (*GetStopOrdersResponse, error) {
	out := new(GetStopOrdersResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.StopOrdersService/GetStopOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stopOrdersServiceClient) CancelStopOrder(ctx context.Context, in *CancelStopOrderRequest, opts ...grpc.CallOption) (*CancelStopOrderResponse, error) {
	out := new(CancelStopOrderResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.StopOrdersService/CancelStopOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StopOrdersServiceServer is the server API for StopOrdersService service.
// All implementations must embed UnimplementedStopOrdersServiceServer
// for forward compatibility
type StopOrdersServiceServer interface {
	// Метод выставления стоп-заявки.
	PostStopOrder(context.Context, *PostStopOrderRequest) (*PostStopOrderResponse, error)
	// Метод получения списка активных стоп заявок по счёту.
	GetStopOrders(context.Context, *GetStopOrdersRequest) (*GetStopOrdersResponse, error)
	// Метод отмены стоп-заявки.
	CancelStopOrder(context.Context, *CancelStopOrderRequest) (*CancelStopOrderResponse, error)
	mustEmbedUnimplementedStopOrdersServiceServer()
}

// UnimplementedStopOrdersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStopOrdersServiceServer struct {
}

func (UnimplementedStopOrdersServiceServer) PostStopOrder(context.Context, *PostStopOrderRequest) (*PostStopOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostStopOrder not implemented")
}
func (UnimplementedStopOrdersServiceServer) GetStopOrders(context.Context, *GetStopOrdersRequest) (*GetStopOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStopOrders not implemented")
}
func (UnimplementedStopOrdersServiceServer) CancelStopOrder(context.Context, *CancelStopOrderRequest) (*CancelStopOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelStopOrder not implemented")
}
func (UnimplementedStopOrdersServiceServer) mustEmbedUnimplementedStopOrdersServiceServer() {}

// UnsafeStopOrdersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StopOrdersServiceServer will
// result in compilation errors.
type UnsafeStopOrdersServiceServer interface {
	mustEmbedUnimplementedStopOrdersServiceServer()
}

func RegisterStopOrdersServiceServer(s grpc.ServiceRegistrar, srv StopOrdersServiceServer) {
	s.RegisterService(&StopOrdersService_ServiceDesc, srv)
}

func _StopOrdersService_PostStopOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostStopOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StopOrdersServiceServer).PostStopOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.StopOrdersService/PostStopOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StopOrdersServiceServer).PostStopOrder(ctx, req.(*PostStopOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StopOrdersService_GetStopOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStopOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StopOrdersServiceServer).GetStopOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.StopOrdersService/GetStopOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StopOrdersServiceServer).GetStopOrders(ctx, req.(*GetStopOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StopOrdersService_CancelStopOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelStopOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StopOrdersServiceServer).CancelStopOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.StopOrdersService/CancelStopOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StopOrdersServiceServer).CancelStopOrder(ctx, req.(*CancelStopOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StopOrdersService_ServiceDesc is the grpc.ServiceDesc for StopOrdersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StopOrdersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.StopOrdersService",
	HandlerType: (*StopOrdersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostStopOrder",
			Handler:    _StopOrdersService_PostStopOrder_Handler,
		},
		{
			MethodName: "GetStopOrders",
			Handler:    _StopOrdersService_GetStopOrders_Handler,
		},
		{
			MethodName: "CancelStopOrder",
			Handler:    _StopOrdersService_CancelStopOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stoporders.proto",
}
