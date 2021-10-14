package admin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	model "github.com/kong/koko/internal/gen/grpc/kong/admin/model/v1"
	v1 "github.com/kong/koko/internal/gen/grpc/kong/admin/service/v1"
	"github.com/kong/koko/internal/server"
	"github.com/kong/koko/internal/store"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// HandlerWrapper is used to wrap a http.Handler with another http.Handler.
type HandlerWrapper interface {
	Wrap(http.Handler) http.Handler
	server.GrpcInterceptorInjector
}

// ContextKey type must be used to manipulate the context of a request.
type ContextKey struct{}

type HandlerOpts struct {
	Logger *zap.Logger

	StoreLoader StoreLoader
}

type CommonOpts struct {
	logger *zap.Logger

	storeLoader StoreLoader
}

func (c CommonOpts) getDB(ctx context.Context,
	cluster *model.RequestCluster) (store.Store, error) {
	store, err := c.storeLoader.Load(ctx, cluster)
	if err != nil {
		if storeLoadErr, ok := err.(StoreLoadErr); ok {
			return nil, status.Error(storeLoadErr.Code, storeLoadErr.Message)
		}
		return nil, err
	}
	return store, nil
}

func NewHandler(opts HandlerOpts) (http.Handler, error) {
	err := validateOpts(opts)
	if err != nil {
		return nil, err
	}
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.
			MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
		}),
		runtime.WithForwardResponseOption(setHTTPStatus),
	)

	err = v1.RegisterMetaServiceHandlerServer(context.Background(),
		mux, &MetaService{})
	if err != nil {
		return nil, err
	}

	err = v1.RegisterServiceServiceHandlerServer(context.Background(),
		mux, &ServiceService{
			CommonOpts: CommonOpts{
				storeLoader: opts.StoreLoader,
				logger: opts.Logger.With(zap.String("admin-service",
					"service")),
			},
		})
	if err != nil {
		return nil, err
	}

	err = v1.RegisterRouteServiceHandlerServer(context.Background(),
		mux, &RouteService{
			CommonOpts: CommonOpts{
				storeLoader: opts.StoreLoader,
				logger: opts.Logger.With(zap.String("admin-service",
					"route")),
			},
		})
	if err != nil {
		return nil, err
	}
	var res http.Handler = mux
	return res, nil
}

func validateOpts(opts HandlerOpts) error {
	if opts.StoreLoader == nil {
		return fmt.Errorf("opts.StoreLoader is required")
	}
	if opts.Logger == nil {
		return fmt.Errorf("opts.Logger is required")
	}
	return nil
}

func NewGRPC(opts HandlerOpts) *grpc.Server {
	server := grpc.NewServer()
	v1.RegisterMetaServiceServer(server, &MetaService{})
	v1.RegisterServiceServiceServer(server, &ServiceService{
		CommonOpts: CommonOpts{
			logger: opts.Logger.With(zap.String("admin-service",
				"service")),
		},
	})
	v1.RegisterRouteServiceServer(server, &RouteService{
		CommonOpts: CommonOpts{
			logger: opts.Logger.With(zap.String("admin-service",
				"route")),
		},
	})
	return server
}
