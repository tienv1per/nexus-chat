// Package httpserver contains shared HTTP inbound adapter helpers.
package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"mini-grpc/backend/internal/application"
	"mini-grpc/backend/internal/platform/correlation"
)

// ReadyCheck verifies whether the service can receive traffic.
type ReadyCheck func(ctx context.Context) error

// Options configures the shared HTTP adapter surface.
type Options struct {
	ServiceName string
	Logger      *slog.Logger
	StartedAt   time.Time
	Ready       ReadyCheck
	Next        http.Handler
}

// New builds an HTTP handler with health endpoints and correlation propagation.
func New(options Options) http.Handler {
	mux := http.NewServeMux()
	startedAt := options.StartedAt
	if startedAt.IsZero() {
		startedAt = time.Now().UTC()
	}

	mux.HandleFunc("/health/live", writeHealth(options.ServiceName, startedAt, nil))
	mux.HandleFunc("/health/ready", writeHealth(options.ServiceName, startedAt, options.Ready))
	mux.HandleFunc("/healthz", writeHealth(options.ServiceName, startedAt, options.Ready))

	if options.Next != nil {
		mux.Handle("/", options.Next)
	}

	return CorrelationMiddleware(options.Logger)(mux)
}

// CorrelationMiddleware reads or creates the correlation ID for one request.
func CorrelationMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := correlation.Normalize(r.Header.Get(correlation.HeaderName))
			ctx := correlation.WithID(r.Context(), id)

			w.Header().Set(correlation.HeaderName, id)
			if logger != nil {
				logger.Debug(
					"http request",
					"method",
					r.Method,
					"path",
					r.URL.Path,
					"correlation_id",
					id,
				)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// StatusCode maps application errors to HTTP status codes.
func StatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	var appErr *application.Error
	if !errors.As(err, &appErr) {
		return http.StatusInternalServerError
	}

	switch appErr.Code {
	case application.ErrorCodeValidation:
		return http.StatusUnprocessableEntity
	case application.ErrorCodeUnauthorized:
		return http.StatusUnauthorized
	case application.ErrorCodeForbidden:
		return http.StatusForbidden
	case application.ErrorCodeNotFound:
		return http.StatusNotFound
	case application.ErrorCodeConflict:
		return http.StatusConflict
	case application.ErrorCodeUnavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

func writeHealth(serviceName string, startedAt time.Time, ready ReadyCheck) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := "ok"
		code := http.StatusOK

		if ready != nil {
			if err := ready(r.Context()); err != nil {
				status = "unavailable"
				code = http.StatusServiceUnavailable
			}
		}

		writeJSON(w, code, map[string]any{
			"service":    serviceName,
			"status":     status,
			"started_at": startedAt.Format(time.RFC3339Nano),
		})
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
