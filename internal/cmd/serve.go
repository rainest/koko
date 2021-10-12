package cmd

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"

	"github.com/hbagdi/gang"
	"github.com/kong/koko/internal/config"
	v1 "github.com/kong/koko/internal/gen/grpc/kong/admin/service/v1"
	"github.com/kong/koko/internal/log"
	"github.com/kong/koko/internal/persistence"
	"github.com/kong/koko/internal/server"
	"github.com/kong/koko/internal/server/admin"
	"github.com/kong/koko/internal/server/kong/ws"
	"github.com/kong/koko/internal/store"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var cfgFile string

// serveCmd is 'koko serve' command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Control plane software for Kong Gateway",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := serveMain(cmd.Context())
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	serveCmd.Flags().StringVar(&cfgFile, "config", "koko.yaml",
		"path to configuration file")

	rootCmd.AddCommand(serveCmd)
}

type initOpts struct {
	Config config.Config
	Logger *zap.Logger
}

func serveMain(ctx context.Context) error {
	opts, err := setup()
	if err != nil {
		return err
	}
	logger := opts.Logger
	logger.Debug("setup successful")

	var g gang.Gang

	memory, err := persistence.NewMemory()
	if err != nil {
		return err
	}
	store := store.New(memory, logger.With(zap.String("component",
		"store"))).ForCluster("default")

	adminLogger := logger.With(zap.String("component", "admin-server"))
	h, err := admin.NewHandler(admin.HandlerOpts{
		Logger:        adminLogger,
		StoreInjector: admin.DefaultStoreWrapper{Store: store},
	})
	if err != nil {
		return err
	}

	s, err := server.NewHTTP(server.HTTPOpts{
		Address: ":3000",
		Logger:  adminLogger,
		Handler: h,
	})
	if err != nil {
		return err
	}
	g.AddWithCtxE(s.Run)

	rawGRPCServer := admin.NewGRPC(admin.HandlerOpts{
		Logger:        logger.With(zap.String("component", "admin-server")),
		StoreInjector: admin.DefaultStoreWrapper{Store: store},
	})
	if err != nil {
		return err
	}

	grpcServer, err := server.NewGRPC(server.GRPCOpts{
		Address:    ":3001",
		GRPCServer: rawGRPCServer,
		Logger:     logger,
	})
	if err != nil {
		return err
	}
	g.AddWithCtxE(grpcServer.Run)

	configClient, err := setupConfigClient()
	if err != nil {
		return err
	}
	controlLogger := logger.With(zap.String("component", "control-server"))
	m := ws.NewManager(ws.ManagerOpts{
		Logger: controlLogger,
		Client: configClient,
	})
	handler, err := ws.NewHandler(ws.HandlerOpts{
		Logger:  controlLogger,
		Manager: m,
	})
	if err != nil {
		return err
	}
	cert, err := tls.LoadX509KeyPair("cluster.crt", "cluster.key")
	if err != nil {
		return err
	}

	s, err = server.NewHTTP(server.HTTPOpts{
		Address: ":3100",
		Logger:  controlLogger,
		Handler: handler,
		TLS: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
		},
	})
	if err != nil {
		return err
	}
	g.AddWithCtxE(s.Run)

	errCh := g.Run(ctx)
	var mErr multiErr
	for err := range errCh {
		mErr.Errors = append(mErr.Errors, err)
	}
	return mErr
}

type multiErr struct {
	Errors []error
}

func (m multiErr) Error() string {
	var buf bytes.Buffer
	for i, err := range m.Errors {
		buf.WriteString("- ")
		buf.WriteString(err.Error())
		if i < len(m.Errors)-1 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}

func setup() (initOpts, error) {
	cfg, err := config.Get(cfgFile)
	if err != nil {
		return initOpts{}, err
	}

	errs := config.Validate(cfg)
	if len(errs) > 0 {
		return initOpts{}, multiError{Errors: errs}
	}

	logger, err := setupLogging(cfg.Log)
	if err != nil {
		return initOpts{}, err
	}
	return initOpts{Config: cfg, Logger: logger}, nil
}

type multiError struct {
	Errors []error
}

func (m multiError) Error() string {
	var b bytes.Buffer
	b.WriteString("Configuration errors:\n")
	for _, err := range m.Errors {
		b.WriteString("- " + err.Error() + "\n")
	}
	return b.String()
}

func setupLogging(c config.Log) (*zap.Logger, error) {
	zapConfig := zap.NewProductionConfig()
	level := config.Levels[c.Level]
	zapConfig.Level.SetLevel(level)
	logger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("create logger: %v", err)
	}
	log.Logger = logger
	return logger, nil
}

func setupConfigClient() (ws.ConfigClient, error) {
	cc, err := grpc.Dial("localhost:3001", grpc.WithInsecure())
	if err != nil {
		return ws.ConfigClient{}, err
	}
	return ws.ConfigClient{
		Service: v1.NewServiceServiceClient(cc),
		Route:   v1.NewRouteServiceClient(cc),
	}, nil
}
