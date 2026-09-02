// Package deliveryconsumer wires the delivery-consumer binary.
package deliveryconsumer

import (
	"context"
	"log/slog"
	"time"

	"mini-grpc/backend/internal/adapters/outbound/kafka"
	"mini-grpc/backend/internal/adapters/outbound/redis"
	"mini-grpc/backend/internal/platform/config"
	"mini-grpc/backend/internal/platform/service"
)

// ServiceName is the stable log identity for this process.
const ServiceName = "delivery-consumer"

// App is the composition root output for delivery-consumer.
type App struct {
	Config    config.Config
	Logger    *slog.Logger
	StartedAt time.Time
	Kafka     *kafka.Publisher
	Redis     *redis.Client
}

// New creates the delivery-consumer graph without using package-level clients.
func New(cfg config.Config, logger *slog.Logger) (*App, error) {
	if logger == nil {
		logger = slog.Default()
	}

	return &App{
		Config:    cfg,
		Logger:    logger,
		StartedAt: time.Now().UTC(),
		Kafka:     kafka.NewPublisher(cfg.Kafka.Brokers),
		Redis:     redis.NewClient(cfg.Redis.Addr),
	}, nil
}

// Run keeps the worker process alive until real Kafka consumption lands.
func (a *App) Run(ctx context.Context) error {
	return service.Wait(ctx, a.Logger, ServiceName)
}
