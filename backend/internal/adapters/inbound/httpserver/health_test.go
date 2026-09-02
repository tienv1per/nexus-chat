package httpserver

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"mini-grpc/backend/internal/platform/correlation"
)

func TestHealthLiveAddsCorrelationID(t *testing.T) {
	handler := New(Options{
		ServiceName: "chat-service",
		StartedAt:   time.Date(2026, 6, 24, 8, 0, 0, 0, time.UTC),
	})

	request := httptest.NewRequest(http.MethodGet, "/health/live", nil)
	request.Header.Set(correlation.HeaderName, "corr_test")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d; want %d", response.Code, http.StatusOK)
	}
	if got := response.Header().Get(correlation.HeaderName); got != "corr_test" {
		t.Fatalf("correlation header = %q; want corr_test", got)
	}
}

func TestHealthReadyReportsUnavailable(t *testing.T) {
	handler := New(Options{
		ServiceName: "chat-service",
		Ready: func(context.Context) error {
			return errors.New("database unavailable")
		},
	})

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/health/ready", nil)

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d; want %d", response.Code, http.StatusServiceUnavailable)
	}
}
