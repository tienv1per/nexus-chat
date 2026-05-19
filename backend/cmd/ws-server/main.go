package main

import (
	"log/slog"
	"os"

	"mini-grpc/backend/internal/platform/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("load config", "error", err)
		os.Exit(1)
	}

	slog.Info(
		"ws server bootstrap ready",
		"http_port",
		cfg.HTTP.WSPort,
		"grpc_port",
		cfg.GRPC.WSPort,
	)
}
