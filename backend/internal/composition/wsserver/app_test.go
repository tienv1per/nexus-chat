package wsserver

import (
	"log/slog"
	"testing"

	"mini-grpc/backend/internal/platform/config"
)

func TestNewWiresWSServerAdapters(t *testing.T) {
	cfg, err := config.LoadFromLookup(func(string) (string, bool) {
		return "", false
	})
	if err != nil {
		t.Fatalf("LoadFromLookup returned error: %v", err)
	}

	app, err := New(cfg, slog.Default())
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}

	if app.Redis.Addr() != cfg.Redis.Addr {
		t.Fatalf("redis addr = %q; want %q", app.Redis.Addr(), cfg.Redis.Addr)
	}
	if app.HTTPServer().Addr != ":8081" {
		t.Fatalf("server addr = %q; want :8081", app.HTTPServer().Addr)
	}
}
