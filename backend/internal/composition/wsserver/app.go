// Package wsserver wires the ws-server binary.
package wsserver

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"mini-grpc/backend/internal/adapters/inbound/httpserver"
	"mini-grpc/backend/internal/adapters/outbound/redis"
	"mini-grpc/backend/internal/platform/config"
)

// ServiceName is the stable log and health identity for this process.
const ServiceName = "ws-server"

// App is the composition root output for ws-server.
type App struct {
	Config    config.Config
	Logger    *slog.Logger
	StartedAt time.Time
	Handler   http.Handler
	Redis     *redis.Client
}

// New creates the ws-server graph without using package-level clients.
func New(cfg config.Config, logger *slog.Logger) (*App, error) {
	if logger == nil {
		logger = slog.Default()
	}

	startedAt := time.Now().UTC()
	app := &App{
		Config:    cfg,
		Logger:    logger,
		StartedAt: startedAt,
		Redis:     redis.NewClient(cfg.Redis.Addr),
	}
	app.Handler = httpserver.New(httpserver.Options{
		ServiceName: ServiceName,
		Logger:      logger,
		StartedAt:   startedAt,
	})

	return app, nil
}

// HTTPServer returns the public HTTP server for WebSocket and health routes.
func (a *App) HTTPServer() *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", a.Config.HTTP.WSPort),
		Handler:           a.Handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
