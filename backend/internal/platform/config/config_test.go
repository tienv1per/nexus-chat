package config

import (
	"strings"
	"testing"
	"time"
)

func TestLoadFromLookupDefaults(t *testing.T) {
	cfg, err := LoadFromLookup(emptyLookup)
	if err != nil {
		t.Fatalf("LoadFromLookup returned error: %v", err)
	}

	if cfg.AppEnv != "local" {
		t.Fatalf("AppEnv = %q; want local", cfg.AppEnv)
	}
	if cfg.HTTP.ChatPort != 8080 {
		t.Fatalf("HTTP.ChatPort = %d; want 8080", cfg.HTTP.ChatPort)
	}
	if cfg.HTTP.WSPort != 8081 {
		t.Fatalf("HTTP.WSPort = %d; want 8081", cfg.HTTP.WSPort)
	}
	if cfg.GRPC.ChatPort != 9080 {
		t.Fatalf("GRPC.ChatPort = %d; want 9080", cfg.GRPC.ChatPort)
	}
	if cfg.GRPC.WSPort != 9081 {
		t.Fatalf("GRPC.WSPort = %d; want 9081", cfg.GRPC.WSPort)
	}
	if cfg.Kafka.MessageCreatedTopic != "chat.message.created" {
		t.Fatalf("MessageCreatedTopic = %q; want chat.message.created", cfg.Kafka.MessageCreatedTopic)
	}
	if cfg.Kafka.MessageDeliveredTopic != "chat.message.delivered" {
		t.Fatalf("MessageDeliveredTopic = %q; want chat.message.delivered", cfg.Kafka.MessageDeliveredTopic)
	}
	if cfg.Upload.MaxBytes != 25*1024*1024 {
		t.Fatalf("Upload.MaxBytes = %d; want %d", cfg.Upload.MaxBytes, 25*1024*1024)
	}
	if cfg.Presence.HeartbeatInterval != 25*time.Second {
		t.Fatalf("HeartbeatInterval = %s; want 25s", cfg.Presence.HeartbeatInterval)
	}
	if cfg.Auth.JWTSecret != "local-dev-secret-change-me" {
		t.Fatalf("Auth.JWTSecret = %q; want default local secret", cfg.Auth.JWTSecret)
	}
}

func TestLoadFromLookupOverrides(t *testing.T) {
	values := map[string]string{
		"APP_ENV":                       "test",
		"LOG_LEVEL":                     "info",
		"CHAT_HTTP_PORT":                "18080",
		"WS_HTTP_PORT":                  "18081",
		"CHAT_GRPC_PORT":                "19080",
		"WS_GRPC_PORT":                  "19081",
		"POSTGRES_DSN":                  "postgres://example",
		"REDIS_ADDR":                    "redis:6379",
		"KAFKA_BROKERS":                 "kafka1:9092,kafka2:9092",
		"KAFKA_TOPIC_MESSAGE_CREATED":   "test.created",
		"KAFKA_TOPIC_MESSAGE_DELIVERED": "test.delivered",
		"UPLOAD_DIR":                    "/tmp/uploads",
		"MEDIA_MAX_BYTES":               "1024",
		"SESSION_TTL_SECONDS":           "45",
		"PRESENCE_TTL_SECONDS":          "45",
		"HEARTBEAT_INTERVAL_SECONDS":    "10",
		"JWT_SECRET":                    "test-secret",
	}

	cfg, err := LoadFromLookup(mapLookup(values))
	if err != nil {
		t.Fatalf("LoadFromLookup returned error: %v", err)
	}

	if cfg.AppEnv != "test" {
		t.Fatalf("AppEnv = %q; want test", cfg.AppEnv)
	}
	if cfg.Postgres.DSN != "postgres://example" {
		t.Fatalf("Postgres.DSN = %q; want postgres://example", cfg.Postgres.DSN)
	}
	if cfg.Kafka.Brokers[1] != "kafka2:9092" {
		t.Fatalf("Kafka.Brokers[1] = %q; want kafka2:9092", cfg.Kafka.Brokers[1])
	}
	if cfg.Upload.Dir != "/tmp/uploads" {
		t.Fatalf("Upload.Dir = %q; want /tmp/uploads", cfg.Upload.Dir)
	}
	if cfg.Presence.SessionTTL != 45*time.Second {
		t.Fatalf("SessionTTL = %s; want 45s", cfg.Presence.SessionTTL)
	}
	if cfg.Auth.JWTSecret != "test-secret" {
		t.Fatalf("Auth.JWTSecret = %q; want test-secret", cfg.Auth.JWTSecret)
	}
}

func TestLoadFromLookupInvalidNumbers(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		value   string
		wantErr string
	}{
		{
			name:    "chat HTTP port",
			key:     "CHAT_HTTP_PORT",
			value:   "not-a-port",
			wantErr: "parsing CHAT_HTTP_PORT",
		},
		{
			name:    "media max bytes",
			key:     "MEDIA_MAX_BYTES",
			value:   "large",
			wantErr: "parsing MEDIA_MAX_BYTES",
		},
		{
			name:    "heartbeat interval",
			key:     "HEARTBEAT_INTERVAL_SECONDS",
			value:   "soon",
			wantErr: "parsing HEARTBEAT_INTERVAL_SECONDS",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadFromLookup(mapLookup(map[string]string{
				tt.key: tt.value,
			}))
			if err == nil {
				t.Fatal("LoadFromLookup returned nil error")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("error = %q; want to contain %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestValidateRejectsInvalidConfig(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*Config)
		wantErr string
	}{
		{
			name: "invalid port",
			mutate: func(cfg *Config) {
				cfg.HTTP.ChatPort = 70000
			},
			wantErr: "CHAT_HTTP_PORT",
		},
		{
			name: "duplicate app port",
			mutate: func(cfg *Config) {
				cfg.GRPC.ChatPort = cfg.HTTP.ChatPort
			},
			wantErr: "duplicates",
		},
		{
			name: "empty Kafka brokers",
			mutate: func(cfg *Config) {
				cfg.Kafka.Brokers = []string{}
			},
			wantErr: "KAFKA_BROKERS",
		},
		{
			name: "non-positive media limit",
			mutate: func(cfg *Config) {
				cfg.Upload.MaxBytes = 0
			},
			wantErr: "MEDIA_MAX_BYTES",
		},
		{
			name: "heartbeat equals presence TTL",
			mutate: func(cfg *Config) {
				cfg.Presence.HeartbeatInterval = cfg.Presence.PresenceTTL
			},
			wantErr: "HEARTBEAT_INTERVAL_SECONDS",
		},
		{
			name: "empty JWT secret",
			mutate: func(cfg *Config) {
				cfg.Auth.JWTSecret = " "
			},
			wantErr: "JWT_SECRET",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := LoadFromLookup(emptyLookup)
			if err != nil {
				t.Fatalf("LoadFromLookup returned error: %v", err)
			}

			tt.mutate(&cfg)
			err = cfg.Validate()
			if err == nil {
				t.Fatal("Validate returned nil error")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("error = %q; want to contain %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func emptyLookup(string) (string, bool) {
	return "", false
}

func mapLookup(values map[string]string) LookupFunc {
	return func(key string) (string, bool) {
		value, ok := values[key]
		return value, ok
	}
}
