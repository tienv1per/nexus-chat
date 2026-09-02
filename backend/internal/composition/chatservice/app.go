// Package chatservice wires the chat-service binary.
package chatservice

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"mini-grpc/backend/internal/adapters/inbound/httpserver"
	"mini-grpc/backend/internal/adapters/outbound/kafka"
	"mini-grpc/backend/internal/adapters/outbound/localmedia"
	"mini-grpc/backend/internal/adapters/outbound/postgres"
	"mini-grpc/backend/internal/adapters/outbound/redis"
	"mini-grpc/backend/internal/platform/config"
)

// ServiceName is the stable log and health identity for this process.
const ServiceName = "chat-service"

// App is the composition root output for chat-service.
type App struct {
	Config     config.Config
	Logger     *slog.Logger
	StartedAt  time.Time
	Handler    http.Handler
	Postgres   *postgres.Repository
	Redis      *redis.Client
	Kafka      *kafka.Publisher
	MediaStore *localmedia.Storage
}

// New creates the chat-service graph without using package-level clients.
func New(cfg config.Config, logger *slog.Logger) (*App, error) {
	if logger == nil {
		logger = slog.Default()
	}

	startedAt := time.Now().UTC()
	app := &App{
		Config:     cfg,
		Logger:     logger,
		StartedAt:  startedAt,
		Postgres:   postgres.NewRepository(cfg.Postgres.DSN),
		Redis:      redis.NewClient(cfg.Redis.Addr),
		Kafka:      kafka.NewPublisher(cfg.Kafka.Brokers),
		MediaStore: localmedia.NewStorage(cfg.Upload.Dir),
	}
	app.Handler = httpserver.New(httpserver.Options{
		ServiceName: ServiceName,
		Logger:      logger,
		StartedAt:   startedAt,
	})

	return app, nil
}

// HTTPServer returns the public HTTP server for REST and health routes.
func (a *App) HTTPServer() *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", a.Config.HTTP.ChatPort),
		Handler:           a.Handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
