package util

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbModel "github.com/kong/koko/internal/gen/grpc/kong/admin/model/v1"
	"github.com/kong/koko/internal/model/json/validation"
	"github.com/kong/koko/internal/store"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	proto2 "google.golang.org/protobuf/proto"
)

const (
	StatusCodeKey = "koko-status-code"
)

type spanKey int

var SpanKey spanKey

type SpanValue interface {
	TraceIDLogKey() string
	SpanIDLogKey() string
	Resource() string
	TraceID() string
	SpanID() string
	SetResource(name string)
}

type loggerKey int

var LoggerKey loggerKey

func HandlerWithLogger(handler http.Handler, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, RequestContextWithLogger(r, logger))
	})
}

// HandlerWithRecovery is http handler middleware that gracefully handles panics by calling a deferred recover
// for all wrapped handlers. When a panic is encountered, a generic error message and response code 500 are returned
// to the client. More detailed error information is logged if the service log level is set to Error or higher.
func HandlerWithRecovery(handler http.Handler, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logPanic(r, err, logger)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				errMsg := []byte(`{"message":"internal server error"}`)
				if _, err := w.Write(errMsg); err != nil {
					logger.Error("internal server error response after recovering from panic", zap.Error(err))
				}
			}
		}()
		handler.ServeHTTP(w, r)
	})
}

func logPanic(req *http.Request, err interface{}, logger *zap.Logger) {
	fields := []zap.Field{
		zap.String("method", req.Method),
		zap.String("path", req.URL.Path),
		zap.String("host", req.Host),
		zap.String("content-type", req.Header.Get("Content-Type")),
		zap.Any("panic-message", err),
	}
	logger.Error("recovered from panic", fields...)
}

func RequestContextWithLogger(req *http.Request, logger *zap.Logger) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), LoggerKey,
		LoggerWithSpan(req.Context(), logger)))
}

func LoggerFromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(LoggerKey).(*zap.Logger); ok {
		return logger
	}
	panic(errors.New("logger not set in context"))
}

func LoggerWithSpan(ctx context.Context, l *zap.Logger) *zap.Logger {
	if span, ok := ctx.Value(SpanKey).(SpanValue); ok {
		return l.With(zap.String(span.TraceIDLogKey(), span.TraceID()),
			zap.String(span.SpanIDLogKey(), span.SpanID()))
	}
	return l
}

type ErrClient struct {
	Message string
}

func (e ErrClient) Error() string {
	return e.Message
}

func ErrorHandler(ctx context.Context,
	mux *runtime.ServeMux, m runtime.Marshaler,
	w http.ResponseWriter, r *http.Request, err error,
) {
	if _, ok := status.FromError(err); !ok {
		if log, ok := r.Context().Value(LoggerKey).(*zap.Logger); ok {
			log.With(zap.Error(err)).Error("grpc service error")
		}
	}

	SetSpanResource(ctx)
	runtime.DefaultHTTPErrorHandler(ctx, mux, m,
		w, r, err)
}

func FinishTrace(ctx context.Context,
	_ http.ResponseWriter, _ proto2.Message,
) error {
	SetSpanResource(ctx)
	return nil
}

func SetHeader(ctx context.Context, code int) {
	err := grpc.SetHeader(ctx, metadata.Pairs(StatusCodeKey,
		fmt.Sprintf("%d", code)))
	if err != nil {
		panic(err)
	}
}

func SetHTTPStatus(ctx context.Context, w http.ResponseWriter,
	_ proto2.Message,
) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return fmt.Errorf("no server metadata in context")
	}
	values := md.HeaderMD.Get(StatusCodeKey)
	if len(values) > 0 {
		code, err := strconv.Atoi(values[0])
		if err != nil {
			return err
		}
		defer w.WriteHeader(code)
	}
	w.Header().Del("grpc-metadata-" + StatusCodeKey)
	return nil
}

func SetSpanResource(ctx context.Context) {
	if span, ok := ctx.Value(SpanKey).(SpanValue); ok {
		if path, ok := runtime.HTTPPathPattern(ctx); ok {
			span.SetResource(path)
		}
	}
}

func HandleErr(ctx context.Context, logger *zap.Logger, err error) error {
	if err == nil {
		return nil
	}

	logger = LoggerWithSpan(ctx, logger)
	if errors.Is(err, store.ErrNotFound) {
		return status.Error(codes.NotFound, "")
	}

	// Do not include types that are aliased to the `error`
	// interface, as it will leak all errors to clients.
	switch e := err.(type) {
	case store.ErrConstraint:
		s := status.New(codes.InvalidArgument, "data constraint error")
		errDetail := &pbModel.ErrorDetail{
			Type:     pbModel.ErrorType_ERROR_TYPE_REFERENCE,
			Field:    e.Index.FieldName,
			Messages: []string{e.Error()},
		}
		s, err := s.WithDetails(errDetail)
		if err != nil {
			panic(err)
		}
		return s.Err()
	case validation.Error:
		s := status.New(codes.InvalidArgument, "validation error")
		var errs []Message
		for _, err := range e.Errs {
			errs = append(errs, err)
		}
		s, err := s.WithDetails(errs...)
		if err != nil {
			panic(err)
		}
		return s.Err()
	case ErrClient:
		s := status.New(codes.InvalidArgument, e.Message)
		return s.Err()
	case store.ErrUnsupportedListOpts:
		return status.New(codes.FailedPrecondition, e.Error()).Err()
	default:
		logger.With(zap.Error(err)).Error("error in service")
		return status.Error(codes.Internal, "")
	}
}

func LoggerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		_ *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (interface{}, error) {
		if _, ok := ctx.Value(LoggerKey).(*zap.Logger); !ok {
			ctx = context.WithValue(ctx, LoggerKey, LoggerWithSpan(ctx, logger))
			return handler(ctx, req)
		}
		return handler(ctx, req)
	}
}

// PanicInterceptor wraps the panic recovery handler grpcRecoveryHandler for use as a UnaryServerInterceptor.
func PanicInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return grpcRecovery.UnaryServerInterceptor(grpcRecoveryHandler(logger))
}

// PanicStreamInterceptor wraps the panic recovery handler grpcRecoveryHandler for use as a StreamServerInterceptor.
func PanicStreamInterceptor(logger *zap.Logger) grpc.StreamServerInterceptor {
	return grpcRecovery.StreamServerInterceptor(grpcRecoveryHandler(logger))
}

// grpcRecoveryHandler facilitates graceful handling of panics for use in both Unary and Stream grpc servers.
// When a panic is encountered, a generic error message and response code 500 are returned to the client.
// More detailed error information is logged if the service log level is set to Error or higher.
func grpcRecoveryHandler(logger *zap.Logger) grpcRecovery.Option {
	return grpcRecovery.WithRecoveryHandler(func(p interface{}) (err error) {
		logger.Error(zap.Any("panic-message", p).String)
		return status.Errorf(codes.Internal, "internal server error")
	})
}
