package main

import (
	"log/slog"
	"os"

	"mini-grpc/backend/internal/composition/chatservice"
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

	log := logger.New(cfg.LogLevel).With("service", chatservice.ServiceName)
	slog.SetDefault(log)

	app, err := chatservice.New(cfg, log)
	if err != nil {
		log.Error("compose chat service", "error", err)
		os.Exit(1)
	}

	ctx, stop := service.SignalContext()
	defer stop()

	log.Info("chat service bootstrap ready", "http_port", cfg.HTTP.ChatPort, "grpc_port", cfg.GRPC.ChatPort)
	if err := service.ServeHTTP(ctx, log, chatservice.ServiceName, app.HTTPServer()); err != nil {
		log.Error("run chat service", "error", err)
		os.Exit(1)
	}
}
