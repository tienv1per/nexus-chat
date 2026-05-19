# Backend

Go services for the Local Real-Time Chat System V1.

## Commands

```bash
go test ./...
go vet ./...
go run ./cmd/chat-service
go run ./cmd/ws-server
go run ./cmd/delivery-consumer
go run ./cmd/status-consumer
```

## Layout

- `cmd/chat-service`: REST and internal gRPC command/query service.
- `cmd/ws-server`: WebSocket endpoint and internal push gRPC service.
- `cmd/delivery-consumer`: Kafka consumer for realtime fan-out.
- `cmd/status-consumer`: Kafka consumer for delivery status writes.
- `internal/platform`: shared platform concerns such as configuration.
