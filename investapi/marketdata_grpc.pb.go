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

// MarketDataServiceClient is the client API for MarketDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketDataServiceClient interface {
	// Метод запроса исторических свечей по инструменту.
	GetCandles(ctx context.Context, in *GetCandlesRequest, opts ...grpc.CallOption) (*GetCandlesResponse, error)
	// Метод запроса цен последних сделок по инструментам.
	GetLastPrices(ctx context.Context, in *GetLastPricesRequest, opts ...grpc.CallOption) (*GetLastPricesResponse, error)
	// Метод получения стакана по инструменту.
	GetOrderBook(ctx context.Context, in *GetOrderBookRequest, opts ...grpc.CallOption) (*GetOrderBookResponse, error)
	// Метод запроса статуса торгов по инструментам.
	GetTradingStatus(ctx context.Context, in *GetTradingStatusRequest, opts ...grpc.CallOption) (*GetTradingStatusResponse, error)
	// Метод запроса статуса торгов по инструментам.
	GetTradingStatuses(ctx context.Context, in *GetTradingStatusesRequest, opts ...grpc.CallOption) (*GetTradingStatusesResponse, error)
	// Метод запроса обезличенных сделок за последний час.
	GetLastTrades(ctx context.Context, in *GetLastTradesRequest, opts ...grpc.CallOption) (*GetLastTradesResponse, error)
	// Метод запроса цен закрытия торговой сессии по инструментам.
	GetClosePrices(ctx context.Context, in *GetClosePricesRequest, opts ...grpc.CallOption) (*GetClosePricesResponse, error)
}

type marketDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketDataServiceClient(cc grpc.ClientConnInterface) MarketDataServiceClient {
	return &marketDataServiceClient{cc}
}

func (c *marketDataServiceClient) GetCandles(ctx context.Context, in *GetCandlesRequest, opts ...grpc.CallOption) (*GetCandlesResponse, error) {
	out := new(GetCandlesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetCandles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) GetLastPrices(ctx context.Context, in *GetLastPricesRequest, opts ...grpc.CallOption) (*GetLastPricesResponse, error) {
	out := new(GetLastPricesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) GetOrderBook(ctx context.Context, in *GetOrderBookRequest, opts ...grpc.CallOption) (*GetOrderBookResponse, error) {
	out := new(GetOrderBookResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetOrderBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) GetTradingStatus(ctx context.Context, in *GetTradingStatusRequest, opts ...grpc.CallOption) (*GetTradingStatusResponse, error) {
	out := new(GetTradingStatusResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTradingStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) GetTradingStatuses(ctx context.Context, in *GetTradingStatusesRequest, opts ...grpc.CallOption) (*GetTradingStatusesResponse, error) {
	out := new(GetTradingStatusesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTradingStatuses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) GetLastTrades(ctx context.Context, in *GetLastTradesRequest, opts ...grpc.CallOption) (*GetLastTradesResponse, error) {
	out := new(GetLastTradesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) GetClosePrices(ctx context.Context, in *GetClosePricesRequest, opts ...grpc.CallOption) (*GetClosePricesResponse, error) {
	out := new(GetClosePricesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetClosePrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketDataServiceServer is the server API for MarketDataService service.
// All implementations must embed UnimplementedMarketDataServiceServer
// for forward compatibility
type MarketDataServiceServer interface {
	// Метод запроса исторических свечей по инструменту.
	GetCandles(context.Context, *GetCandlesRequest) (*GetCandlesResponse, error)
	// Метод запроса цен последних сделок по инструментам.
	GetLastPrices(context.Context, *GetLastPricesRequest) (*GetLastPricesResponse, error)
	// Метод получения стакана по инструменту.
	GetOrderBook(context.Context, *GetOrderBookRequest) (*GetOrderBookResponse, error)
	// Метод запроса статуса торгов по инструментам.
	GetTradingStatus(context.Context, *GetTradingStatusRequest) (*GetTradingStatusResponse, error)
	// Метод запроса статуса торгов по инструментам.
	GetTradingStatuses(context.Context, *GetTradingStatusesRequest) (*GetTradingStatusesResponse, error)
	// Метод запроса обезличенных сделок за последний час.
	GetLastTrades(context.Context, *GetLastTradesRequest) (*GetLastTradesResponse, error)
	// Метод запроса цен закрытия торговой сессии по инструментам.
	GetClosePrices(context.Context, *GetClosePricesRequest) (*GetClosePricesResponse, error)
	mustEmbedUnimplementedMarketDataServiceServer()
}

// UnimplementedMarketDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMarketDataServiceServer struct {
}

func (UnimplementedMarketDataServiceServer) GetCandles(context.Context, *GetCandlesRequest) (*GetCandlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCandles not implemented")
}
func (UnimplementedMarketDataServiceServer) GetLastPrices(context.Context, *GetLastPricesRequest) (*GetLastPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastPrices not implemented")
}
func (UnimplementedMarketDataServiceServer) GetOrderBook(context.Context, *GetOrderBookRequest) (*GetOrderBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderBook not implemented")
}
func (UnimplementedMarketDataServiceServer) GetTradingStatus(context.Context, *GetTradingStatusRequest) (*GetTradingStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTradingStatus not implemented")
}
func (UnimplementedMarketDataServiceServer) GetTradingStatuses(context.Context, *GetTradingStatusesRequest) (*GetTradingStatusesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTradingStatuses not implemented")
}
func (UnimplementedMarketDataServiceServer) GetLastTrades(context.Context, *GetLastTradesRequest) (*GetLastTradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastTrades not implemented")
}
func (UnimplementedMarketDataServiceServer) GetClosePrices(context.Context, *GetClosePricesRequest) (*GetClosePricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClosePrices not implemented")
}
func (UnimplementedMarketDataServiceServer) mustEmbedUnimplementedMarketDataServiceServer() {}

// UnsafeMarketDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketDataServiceServer will
// result in compilation errors.
type UnsafeMarketDataServiceServer interface {
	mustEmbedUnimplementedMarketDataServiceServer()
}

func RegisterMarketDataServiceServer(s grpc.ServiceRegistrar, srv MarketDataServiceServer) {
	s.RegisterService(&MarketDataService_ServiceDesc, srv)
}

func _MarketDataService_GetCandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetCandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetCandles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetCandles(ctx, req.(*GetCandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_GetLastPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetLastPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetLastPrices(ctx, req.(*GetLastPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_GetOrderBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetOrderBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetOrderBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetOrderBook(ctx, req.(*GetOrderBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_GetTradingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradingStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetTradingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTradingStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetTradingStatus(ctx, req.(*GetTradingStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_GetTradingStatuses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradingStatusesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetTradingStatuses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTradingStatuses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetTradingStatuses(ctx, req.(*GetTradingStatusesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_GetLastTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastTradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetLastTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetLastTrades(ctx, req.(*GetLastTradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_GetClosePrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClosePricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetClosePrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetClosePrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetClosePrices(ctx, req.(*GetClosePricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketDataService_ServiceDesc is the grpc.ServiceDesc for MarketDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.MarketDataService",
	HandlerType: (*MarketDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCandles",
			Handler:    _MarketDataService_GetCandles_Handler,
		},
		{
			MethodName: "GetLastPrices",
			Handler:    _MarketDataService_GetLastPrices_Handler,
		},
		{
			MethodName: "GetOrderBook",
			Handler:    _MarketDataService_GetOrderBook_Handler,
		},
		{
			MethodName: "GetTradingStatus",
			Handler:    _MarketDataService_GetTradingStatus_Handler,
		},
		{
			MethodName: "GetTradingStatuses",
			Handler:    _MarketDataService_GetTradingStatuses_Handler,
		},
		{
			MethodName: "GetLastTrades",
			Handler:    _MarketDataService_GetLastTrades_Handler,
		},
		{
			MethodName: "GetClosePrices",
			Handler:    _MarketDataService_GetClosePrices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketdata.proto",
}

// MarketDataStreamServiceClient is the client API for MarketDataStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketDataStreamServiceClient interface {
	// Bi-directional стрим предоставления биржевой информации.
	MarketDataStream(ctx context.Context, opts ...grpc.CallOption) (MarketDataStreamService_MarketDataStreamClient, error)
	// Server-side стрим предоставления биржевой информации.
	MarketDataServerSideStream(ctx context.Context, in *MarketDataServerSideStreamRequest, opts ...grpc.CallOption) (MarketDataStreamService_MarketDataServerSideStreamClient, error)
}

type marketDataStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketDataStreamServiceClient(cc grpc.ClientConnInterface) MarketDataStreamServiceClient {
	return &marketDataStreamServiceClient{cc}
}

func (c *marketDataStreamServiceClient) MarketDataStream(ctx context.Context, opts ...grpc.CallOption) (MarketDataStreamService_MarketDataStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketDataStreamService_ServiceDesc.Streams[0], "/tinkoff.public.invest.api.contract.v1.MarketDataStreamService/MarketDataStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketDataStreamServiceMarketDataStreamClient{stream}
	return x, nil
}

type MarketDataStreamService_MarketDataStreamClient interface {
	Send(*MarketDataRequest) error
	Recv() (*MarketDataResponse, error)
	grpc.ClientStream
}

type marketDataStreamServiceMarketDataStreamClient struct {
	grpc.ClientStream
}

func (x *marketDataStreamServiceMarketDataStreamClient) Send(m *MarketDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *marketDataStreamServiceMarketDataStreamClient) Recv() (*MarketDataResponse, error) {
	m := new(MarketDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *marketDataStreamServiceClient) MarketDataServerSideStream(ctx context.Context, in *MarketDataServerSideStreamRequest, opts ...grpc.CallOption) (MarketDataStreamService_MarketDataServerSideStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketDataStreamService_ServiceDesc.Streams[1], "/tinkoff.public.invest.api.contract.v1.MarketDataStreamService/MarketDataServerSideStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketDataStreamServiceMarketDataServerSideStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MarketDataStreamService_MarketDataServerSideStreamClient interface {
	Recv() (*MarketDataResponse, error)
	grpc.ClientStream
}

type marketDataStreamServiceMarketDataServerSideStreamClient struct {
	grpc.ClientStream
}

func (x *marketDataStreamServiceMarketDataServerSideStreamClient) Recv() (*MarketDataResponse, error) {
	m := new(MarketDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MarketDataStreamServiceServer is the server API for MarketDataStreamService service.
// All implementations must embed UnimplementedMarketDataStreamServiceServer
// for forward compatibility
type MarketDataStreamServiceServer interface {
	// Bi-directional стрим предоставления биржевой информации.
	MarketDataStream(MarketDataStreamService_MarketDataStreamServer) error
	// Server-side стрим предоставления биржевой информации.
	MarketDataServerSideStream(*MarketDataServerSideStreamRequest, MarketDataStreamService_MarketDataServerSideStreamServer) error
	mustEmbedUnimplementedMarketDataStreamServiceServer()
}

// UnimplementedMarketDataStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMarketDataStreamServiceServer struct {
}

func (UnimplementedMarketDataStreamServiceServer) MarketDataStream(MarketDataStreamService_MarketDataStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MarketDataStream not implemented")
}
func (UnimplementedMarketDataStreamServiceServer) MarketDataServerSideStream(*MarketDataServerSideStreamRequest, MarketDataStreamService_MarketDataServerSideStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MarketDataServerSideStream not implemented")
}
func (UnimplementedMarketDataStreamServiceServer) mustEmbedUnimplementedMarketDataStreamServiceServer() {
}

// UnsafeMarketDataStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketDataStreamServiceServer will
// result in compilation errors.
type UnsafeMarketDataStreamServiceServer interface {
	mustEmbedUnimplementedMarketDataStreamServiceServer()
}

func RegisterMarketDataStreamServiceServer(s grpc.ServiceRegistrar, srv MarketDataStreamServiceServer) {
	s.RegisterService(&MarketDataStreamService_ServiceDesc, srv)
}

func _MarketDataStreamService_MarketDataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MarketDataStreamServiceServer).MarketDataStream(&marketDataStreamServiceMarketDataStreamServer{stream})
}

type MarketDataStreamService_MarketDataStreamServer interface {
	Send(*MarketDataResponse) error
	Recv() (*MarketDataRequest, error)
	grpc.ServerStream
}

type marketDataStreamServiceMarketDataStreamServer struct {
	grpc.ServerStream
}

func (x *marketDataStreamServiceMarketDataStreamServer) Send(m *MarketDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *marketDataStreamServiceMarketDataStreamServer) Recv() (*MarketDataRequest, error) {
	m := new(MarketDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MarketDataStreamService_MarketDataServerSideStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MarketDataServerSideStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MarketDataStreamServiceServer).MarketDataServerSideStream(m, &marketDataStreamServiceMarketDataServerSideStreamServer{stream})
}

type MarketDataStreamService_MarketDataServerSideStreamServer interface {
	Send(*MarketDataResponse) error
	grpc.ServerStream
}

type marketDataStreamServiceMarketDataServerSideStreamServer struct {
	grpc.ServerStream
}

func (x *marketDataStreamServiceMarketDataServerSideStreamServer) Send(m *MarketDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MarketDataStreamService_ServiceDesc is the grpc.ServiceDesc for MarketDataStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketDataStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.MarketDataStreamService",
	HandlerType: (*MarketDataStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MarketDataStream",
			Handler:       _MarketDataStreamService_MarketDataStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MarketDataServerSideStream",
			Handler:       _MarketDataStreamService_MarketDataServerSideStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "marketdata.proto",
}
