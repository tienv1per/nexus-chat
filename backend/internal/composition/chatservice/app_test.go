package chatservice

import (
	"log/slog"
	"testing"

	"mini-grpc/backend/internal/platform/config"
)

func TestNewWiresChatServiceAdapters(t *testing.T) {
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

	if app.Postgres.DSN() != cfg.Postgres.DSN {
		t.Fatalf("postgres DSN = %q; want %q", app.Postgres.DSN(), cfg.Postgres.DSN)
	}
	if app.Redis.Addr() != cfg.Redis.Addr {
		t.Fatalf("redis addr = %q; want %q", app.Redis.Addr(), cfg.Redis.Addr)
	}
	if app.MediaStore.RootDir() != cfg.Upload.Dir {
		t.Fatalf("media root = %q; want %q", app.MediaStore.RootDir(), cfg.Upload.Dir)
	}
	if app.HTTPServer().Addr != ":8080" {
		t.Fatalf("server addr = %q; want :8080", app.HTTPServer().Addr)
	}
}
