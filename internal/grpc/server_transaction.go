package grpc

import (
	"context"
	"database/sql"

	"google.golang.org/grpc"

	"github.com/v8tix/eda/di"
	"github.com/v8tix/mallbots-search-proto/pb"
	"github.com/v8tix/mallbots-search/internal/application"
)

type serverTx struct {
	c di.Container
	pb.UnimplementedSearchServiceServer
}

var _ pb.SearchServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	pb.RegisterSearchServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) SearchOrders(ctx context.Context, request *pb.SearchOrdersRequest) (resp *pb.SearchOrdersResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.Application)}

	return next.SearchOrders(ctx, request)
}

func (s serverTx) GetOrder(ctx context.Context, request *pb.GetSearchOrderRequest) (resp *pb.GetSearchOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, "tx").(*sql.Tx))

	next := server{app: di.Get(ctx, "app").(application.Application)}

	return next.GetOrder(ctx, request)
}

func (s serverTx) closeTx(tx *sql.Tx, err error) error {
	if p := recover(); p != nil {
		_ = tx.Rollback()
		panic(p)
	} else if err != nil {
		_ = tx.Rollback()
		return err
	} else {
		return tx.Commit()
	}
}
