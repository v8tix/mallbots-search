package grpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/v8tix/mallbots-search-proto/pb"
	"github.com/v8tix/mallbots-search/internal/application"
)

type server struct {
	app application.Application
	pb.UnimplementedSearchServiceServer
}

func RegisterServer(_ context.Context, app application.Application, registrar grpc.ServiceRegistrar) error {
	pb.RegisterSearchServiceServer(registrar, server{app: app})
	return nil
}

func (s server) SearchOrders(ctx context.Context, request *pb.SearchOrdersRequest) (*pb.SearchOrdersResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (s server) GetOrder(ctx context.Context, request *pb.GetSearchOrderRequest) (*pb.GetSearchOrderResponse, error) {
	// TODO implement me
	panic("implement me")
}
