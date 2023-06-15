// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: pb/ordering.api.proto

package pb

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

const (
	OrderingService_CreateOrder_FullMethodName   = "/pb.OrderingService/CreateOrder"
	OrderingService_GetOrder_FullMethodName      = "/pb.OrderingService/GetOrder"
	OrderingService_CancelOrder_FullMethodName   = "/pb.OrderingService/CancelOrder"
	OrderingService_ReadyOrder_FullMethodName    = "/pb.OrderingService/ReadyOrder"
	OrderingService_CompleteOrder_FullMethodName = "/pb.OrderingService/CompleteOrder"
)

// OrderingServiceClient is the client API for OrderingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderingServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	ReadyOrder(ctx context.Context, in *ReadyOrderRequest, opts ...grpc.CallOption) (*ReadyOrderResponse, error)
	CompleteOrder(ctx context.Context, in *CompleteOrderRequest, opts ...grpc.CallOption) (*CompleteOrderResponse, error)
}

type orderingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderingServiceClient(cc grpc.ClientConnInterface) OrderingServiceClient {
	return &orderingServiceClient{cc}
}

func (c *orderingServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, OrderingService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderingServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderResponse, error) {
	out := new(GetOrderResponse)
	err := c.cc.Invoke(ctx, OrderingService_GetOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderingServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, OrderingService_CancelOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderingServiceClient) ReadyOrder(ctx context.Context, in *ReadyOrderRequest, opts ...grpc.CallOption) (*ReadyOrderResponse, error) {
	out := new(ReadyOrderResponse)
	err := c.cc.Invoke(ctx, OrderingService_ReadyOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderingServiceClient) CompleteOrder(ctx context.Context, in *CompleteOrderRequest, opts ...grpc.CallOption) (*CompleteOrderResponse, error) {
	out := new(CompleteOrderResponse)
	err := c.cc.Invoke(ctx, OrderingService_CompleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderingServiceServer is the server API for OrderingService service.
// All implementations must embed UnimplementedOrderingServiceServer
// for forward compatibility
type OrderingServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	ReadyOrder(context.Context, *ReadyOrderRequest) (*ReadyOrderResponse, error)
	CompleteOrder(context.Context, *CompleteOrderRequest) (*CompleteOrderResponse, error)
	mustEmbedUnimplementedOrderingServiceServer()
}

// UnimplementedOrderingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderingServiceServer struct {
}

func (UnimplementedOrderingServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderingServiceServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderingServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderingServiceServer) ReadyOrder(context.Context, *ReadyOrderRequest) (*ReadyOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadyOrder not implemented")
}
func (UnimplementedOrderingServiceServer) CompleteOrder(context.Context, *CompleteOrderRequest) (*CompleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteOrder not implemented")
}
func (UnimplementedOrderingServiceServer) mustEmbedUnimplementedOrderingServiceServer() {}

// UnsafeOrderingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderingServiceServer will
// result in compilation errors.
type UnsafeOrderingServiceServer interface {
	mustEmbedUnimplementedOrderingServiceServer()
}

func RegisterOrderingServiceServer(s grpc.ServiceRegistrar, srv OrderingServiceServer) {
	s.RegisterService(&OrderingService_ServiceDesc, srv)
}

func _OrderingService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderingService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderingService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderingService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderingService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderingService_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderingService_ReadyOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadyOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).ReadyOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderingService_ReadyOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).ReadyOrder(ctx, req.(*ReadyOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderingService_CompleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServiceServer).CompleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderingService_CompleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServiceServer).CompleteOrder(ctx, req.(*CompleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderingService_ServiceDesc is the grpc.ServiceDesc for OrderingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.OrderingService",
	HandlerType: (*OrderingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderingService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderingService_GetOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderingService_CancelOrder_Handler,
		},
		{
			MethodName: "ReadyOrder",
			Handler:    _OrderingService_ReadyOrder_Handler,
		},
		{
			MethodName: "CompleteOrder",
			Handler:    _OrderingService_CompleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/ordering.api.proto",
}
