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
- `internal/domain`: framework-free chat IDs, enums, and entities.
- `internal/application`: inbound use-case contracts, outbound ports, and application errors.
- `internal/adapters`: inbound/outbound adapter boundaries for HTTP, gRPC, WebSocket, PostgreSQL, Redis, Kafka, and local media storage.
- `internal/composition`: per-binary wiring roots with no package-level database or client globals.
- `internal/platform`: shared platform concerns such as configuration, logging, correlation IDs, and process runtime.

## Health

`chat-service` and `ws-server` expose:

- `GET /health/live`
- `GET /health/ready`
- `GET /healthz`
