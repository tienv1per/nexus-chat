package main

import (
	"log/slog"
	"os"

	"mini-grpc/backend/internal/composition/wsserver"
	"mini-grpc/backend/internal/platform/config"
	"mini-grpc/backend/internal/platform/logger"
	"mini-grpc/backend/internal/platform/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("load config", "error", err)
		os.Exit(1)
	}

	log := logger.New(cfg.LogLevel).With("service", wsserver.ServiceName)
	slog.SetDefault(log)

	app, err := wsserver.New(cfg, log)
	if err != nil {
		log.Error("compose ws server", "error", err)
		os.Exit(1)
	}

	ctx, stop := service.SignalContext()
	defer stop()

	log.Info("ws server bootstrap ready", "http_port", cfg.HTTP.WSPort, "grpc_port", cfg.GRPC.WSPort)
	if err := service.ServeHTTP(ctx, log, wsserver.ServiceName, app.HTTPServer()); err != nil {
		log.Error("run ws server", "error", err)
		os.Exit(1)
	}
}
