package search

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"github.com/v8tix/eda/ddd"
	"github.com/v8tix/eda/di"
	"github.com/v8tix/eda/jetstream"
	pg "github.com/v8tix/eda/postgres"
	"github.com/v8tix/eda/registry"
	"github.com/v8tix/eda/tm"
	customerspb "github.com/v8tix/mallbots-customers-proto/pb"
	orderingpb "github.com/v8tix/mallbots-ordering-proto/pb"
	"github.com/v8tix/mallbots-search-proto/rest"
	"github.com/v8tix/mallbots-search/internal/application"
	"github.com/v8tix/mallbots-search/internal/grpc"
	"github.com/v8tix/mallbots-search/internal/handlers"
	"github.com/v8tix/mallbots-search/internal/logging"
	"github.com/v8tix/mallbots-search/internal/monolith"
	"github.com/v8tix/mallbots-search/internal/postgres"
	storespb "github.com/v8tix/mallbots-stores-proto/pb"
)

type Module struct{}

func (m Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton("registry", func(c di.Container) (any, error) {
		reg := registry.New()
		if err := orderingpb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := customerspb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := storespb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	container.AddSingleton("logger", func(c di.Container) (any, error) {
		return mono.Logger(), nil
	})
	container.AddSingleton("stream", func(c di.Container) (any, error) {
		return jetstream.NewStream(mono.Config().Nats.Stream, mono.JS(), c.Get("logger").(zerolog.Logger)), nil
	})
	container.AddSingleton("db", func(c di.Container) (any, error) {
		return mono.DB(), nil
	})
	container.AddSingleton("conn", func(c di.Container) (any, error) {
		return grpc.Dial(ctx, mono.Config().Rpc.Address())
	})
	container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*sql.DB)
		return db.Begin()
	})
	container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		inboxStore := pg.NewInboxStore("search.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	container.AddScoped("customers", func(c di.Container) (any, error) {
		return postgres.NewCustomerCacheRepository(
			"search.customers_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewCustomerRepository(c.Get("conn").(*grpc.ClientConn)),
		), nil
	})
	container.AddScoped("stores", func(c di.Container) (any, error) {
		return postgres.NewStoreCacheRepository(
			"search.stores_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewStoreRepository(c.Get("conn").(*grpc.ClientConn)),
		), nil
	})
	container.AddScoped("products", func(c di.Container) (any, error) {
		return postgres.NewProductCacheRepository(
			"search.products_cache",
			c.Get("tx").(*sql.Tx),
			grpc.NewProductRepository(c.Get("conn").(*grpc.ClientConn)),
		), nil
	})
	container.AddScoped("orders", func(c di.Container) (any, error) {
		return postgres.NewOrderRepository("search.orders", c.Get("tx").(*sql.Tx)), nil
	})

	// setup application
	container.AddScoped("app", func(c di.Container) (any, error) {
		return logging.LogApplicationAccess(
			application.New(
				c.Get("orders").(application.OrderRepository),
			),
			c.Get("logger").(zerolog.Logger),
		), nil
	})
	container.AddScoped("integrationEventHandlers", func(c di.Container) (any, error) {
		return logging.LogEventHandlerAccess[ddd.Event](
			handlers.NewIntegrationEventHandlers(
				c.Get("orders").(application.OrderRepository),
				c.Get("customers").(application.CustomerCacheRepository),
				c.Get("stores").(application.StoreCacheRepository),
				c.Get("products").(application.ProductCacheRepository),
			),
			"IntegrationEvents", c.Get("logger").(zerolog.Logger),
		), nil
	})

	// setup Driver adapters
	if err = grpc.RegisterServerTx(container, mono.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}

	return nil
}
