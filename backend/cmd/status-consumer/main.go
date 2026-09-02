package main

import (
	"log/slog"
	"os"

	"mini-grpc/backend/internal/composition/statusconsumer"
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

	log := logger.New(cfg.LogLevel).With("service", statusconsumer.ServiceName)
	slog.SetDefault(log)

	app, err := statusconsumer.New(cfg, log)
	if err != nil {
		log.Error("compose status consumer", "error", err)
		os.Exit(1)
	}

	ctx, stop := service.SignalContext()
	defer stop()

	log.Info(
		"status consumer bootstrap ready",
		"topic",
		cfg.Kafka.MessageDeliveredTopic,
		"storage",
		"postgres",
	)
	if err := app.Run(ctx); err != nil {
		log.Error("run status consumer", "error", err)
		os.Exit(1)
	}
}
