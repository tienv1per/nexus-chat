package main

import (
	"log/slog"
	"os"

	"mini-grpc/backend/internal/composition/deliveryconsumer"
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

	log := logger.New(cfg.LogLevel).With("service", deliveryconsumer.ServiceName)
	slog.SetDefault(log)

	app, err := deliveryconsumer.New(cfg, log)
	if err != nil {
		log.Error("compose delivery consumer", "error", err)
		os.Exit(1)
	}

	ctx, stop := service.SignalContext()
	defer stop()

	log.Info(
		"delivery consumer bootstrap ready",
		"topic",
		cfg.Kafka.MessageCreatedTopic,
		"brokers",
		cfg.Kafka.Brokers,
	)
	if err := app.Run(ctx); err != nil {
		log.Error("run delivery consumer", "error", err)
		os.Exit(1)
	}
}
