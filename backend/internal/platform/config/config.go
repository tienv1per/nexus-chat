// Package config loads local service configuration from environment variables.
package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultAppEnv                      = "local"
	defaultLogLevel                    = "debug"
	defaultChatHTTPPort                = 8080
	defaultWSHTTPPort                  = 8081
	defaultChatGRPCPort                = 9080
	defaultWSGRPCPort                  = 9081
	defaultPostgresDSN                 = "postgres://chat:chat@localhost:5432/chat_v1?sslmode=disable"
	defaultRedisAddr                   = "localhost:6379"
	defaultKafkaBrokers                = "localhost:9092"
	defaultMessageCreatedTopic         = "chat.message.created"
	defaultMessageDeliveredTopic       = "chat.message.delivered"
	defaultUploadDir                   = "./data/uploads"
	defaultMediaMaxBytes         int64 = 25 * 1024 * 1024
	defaultSessionTTLSeconds           = 60
	defaultPresenceTTLSeconds          = 60
	defaultHeartbeatSeconds            = 25
	defaultJWTSecret                   = "local-dev-secret-change-me"
)

// LookupFunc returns the environment value for a key.
type LookupFunc func(key string) (string, bool)

// Config contains process configuration shared by local chat services.
type Config struct {
	AppEnv   string
	LogLevel string
	HTTP     HTTPConfig
	GRPC     GRPCConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Kafka    KafkaConfig
	Upload   UploadConfig
	Presence PresenceConfig
	Auth     AuthConfig
}

// HTTPConfig contains public HTTP ports.
type HTTPConfig struct {
	ChatPort int
	WSPort   int
}

// GRPCConfig contains internal gRPC ports.
type GRPCConfig struct {
	ChatPort int
	WSPort   int
}

// PostgresConfig contains PostgreSQL connection settings.
type PostgresConfig struct {
	DSN string
}

// RedisConfig contains Redis connection settings.
type RedisConfig struct {
	Addr string
}

// KafkaConfig contains Kafka connection settings and topic names.
type KafkaConfig struct {
	Brokers               []string
	MessageCreatedTopic   string
	MessageDeliveredTopic string
}

// UploadConfig contains local media upload settings.
type UploadConfig struct {
	Dir      string
	MaxBytes int64
}

// PresenceConfig contains Redis TTL and WebSocket heartbeat settings.
type PresenceConfig struct {
	SessionTTL        time.Duration
	PresenceTTL       time.Duration
	HeartbeatInterval time.Duration
}

// AuthConfig contains local development auth settings.
type AuthConfig struct {
	JWTSecret string
}

// Load reads configuration from process environment variables.
func Load() (Config, error) {
	return LoadFromLookup(os.LookupEnv)
}

// LoadFromLookup reads configuration using a caller-provided lookup function.
func LoadFromLookup(lookup LookupFunc) (Config, error) {
	chatHTTPPort, err := readInt(lookup, "CHAT_HTTP_PORT", defaultChatHTTPPort)
	if err != nil {
		return Config{}, err
	}

	wsHTTPPort, err := readInt(lookup, "WS_HTTP_PORT", defaultWSHTTPPort)
	if err != nil {
		return Config{}, err
	}

	chatGRPCPort, err := readInt(lookup, "CHAT_GRPC_PORT", defaultChatGRPCPort)
	if err != nil {
		return Config{}, err
	}

	wsGRPCPort, err := readInt(lookup, "WS_GRPC_PORT", defaultWSGRPCPort)
	if err != nil {
		return Config{}, err
	}

	mediaMaxBytes, err := readInt64(lookup, "MEDIA_MAX_BYTES", defaultMediaMaxBytes)
	if err != nil {
		return Config{}, err
	}

	sessionTTL, err := readSeconds(lookup, "SESSION_TTL_SECONDS", defaultSessionTTLSeconds)
	if err != nil {
		return Config{}, err
	}

	presenceTTL, err := readSeconds(lookup, "PRESENCE_TTL_SECONDS", defaultPresenceTTLSeconds)
	if err != nil {
		return Config{}, err
	}

	heartbeatInterval, err := readSeconds(lookup, "HEARTBEAT_INTERVAL_SECONDS", defaultHeartbeatSeconds)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		AppEnv:   readString(lookup, "APP_ENV", defaultAppEnv),
		LogLevel: readString(lookup, "LOG_LEVEL", defaultLogLevel),
		HTTP: HTTPConfig{
			ChatPort: chatHTTPPort,
			WSPort:   wsHTTPPort,
		},
		GRPC: GRPCConfig{
			ChatPort: chatGRPCPort,
			WSPort:   wsGRPCPort,
		},
		Postgres: PostgresConfig{
			DSN: readString(lookup, "POSTGRES_DSN", defaultPostgresDSN),
		},
		Redis: RedisConfig{
			Addr: readString(lookup, "REDIS_ADDR", defaultRedisAddr),
		},
		Kafka: KafkaConfig{
			Brokers:               readCSV(lookup, "KAFKA_BROKERS", defaultKafkaBrokers),
			MessageCreatedTopic:   readString(lookup, "KAFKA_TOPIC_MESSAGE_CREATED", defaultMessageCreatedTopic),
			MessageDeliveredTopic: readString(lookup, "KAFKA_TOPIC_MESSAGE_DELIVERED", defaultMessageDeliveredTopic),
		},
		Upload: UploadConfig{
			Dir:      readString(lookup, "UPLOAD_DIR", defaultUploadDir),
			MaxBytes: mediaMaxBytes,
		},
		Presence: PresenceConfig{
			SessionTTL:        sessionTTL,
			PresenceTTL:       presenceTTL,
			HeartbeatInterval: heartbeatInterval,
		},
		Auth: AuthConfig{
			JWTSecret: readString(lookup, "JWT_SECRET", defaultJWTSecret),
		},
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// Validate verifies required local configuration invariants.
func (c Config) Validate() error {
	ports := map[string]int{
		"CHAT_HTTP_PORT": c.HTTP.ChatPort,
		"WS_HTTP_PORT":   c.HTTP.WSPort,
		"CHAT_GRPC_PORT": c.GRPC.ChatPort,
		"WS_GRPC_PORT":   c.GRPC.WSPort,
	}

	seenPorts := map[int]string{}
	for key, port := range ports {
		if err := validatePort(key, port); err != nil {
			return err
		}

		if existingKey, ok := seenPorts[port]; ok {
			return fmt.Errorf("%s duplicates %s on port %d", key, existingKey, port)
		}
		seenPorts[port] = key
	}

	if strings.TrimSpace(c.Postgres.DSN) == "" {
		return fmt.Errorf("POSTGRES_DSN is required")
	}
	if strings.TrimSpace(c.Redis.Addr) == "" {
		return fmt.Errorf("REDIS_ADDR is required")
	}
	if len(c.Kafka.Brokers) == 0 {
		return fmt.Errorf("KAFKA_BROKERS is required")
	}
	if strings.TrimSpace(c.Kafka.MessageCreatedTopic) == "" {
		return fmt.Errorf("KAFKA_TOPIC_MESSAGE_CREATED is required")
	}
	if strings.TrimSpace(c.Kafka.MessageDeliveredTopic) == "" {
		return fmt.Errorf("KAFKA_TOPIC_MESSAGE_DELIVERED is required")
	}
	if strings.TrimSpace(c.Upload.Dir) == "" {
		return fmt.Errorf("UPLOAD_DIR is required")
	}
	if c.Upload.MaxBytes <= 0 {
		return fmt.Errorf("MEDIA_MAX_BYTES must be positive")
	}
	if c.Presence.SessionTTL <= 0 {
		return fmt.Errorf("SESSION_TTL_SECONDS must be positive")
	}
	if c.Presence.PresenceTTL <= 0 {
		return fmt.Errorf("PRESENCE_TTL_SECONDS must be positive")
	}
	if c.Presence.HeartbeatInterval <= 0 {
		return fmt.Errorf("HEARTBEAT_INTERVAL_SECONDS must be positive")
	}
	if c.Presence.HeartbeatInterval >= c.Presence.PresenceTTL {
		return fmt.Errorf("HEARTBEAT_INTERVAL_SECONDS must be lower than PRESENCE_TTL_SECONDS")
	}
	if strings.TrimSpace(c.Auth.JWTSecret) == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}

	return nil
}

func readString(lookup LookupFunc, key string, fallback string) string {
	value, ok := lookup(key)
	if !ok {
		return fallback
	}

	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return fallback
	}

	return trimmed
}

func readCSV(lookup LookupFunc, key string, fallback string) []string {
	value := readString(lookup, key, fallback)
	parts := strings.Split(value, ",")
	items := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		items = append(items, trimmed)
	}

	return items
}

func readInt(lookup LookupFunc, key string, fallback int) (int, error) {
	value := readString(lookup, key, strconv.Itoa(fallback))
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("parsing %s=%q as int: %w", key, value, err)
	}

	return parsed, nil
}

func readInt64(lookup LookupFunc, key string, fallback int64) (int64, error) {
	value := readString(lookup, key, strconv.FormatInt(fallback, 10))
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing %s=%q as int64: %w", key, value, err)
	}

	return parsed, nil
}

func readSeconds(lookup LookupFunc, key string, fallback int) (time.Duration, error) {
	seconds, err := readInt(lookup, key, fallback)
	if err != nil {
		return 0, err
	}

	return time.Duration(seconds) * time.Second, nil
}

func validatePort(key string, port int) error {
	if port <= 0 || port > 65535 {
		return fmt.Errorf("%s must be between 1 and 65535", key)
	}

	return nil
}
