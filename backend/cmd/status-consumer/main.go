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
		"status consumer bootstrap ready",
		"topic",
		cfg.Kafka.MessageDeliveredTopic,
		"keyspace",
		cfg.Cassandra.Keyspace,
	)
}
