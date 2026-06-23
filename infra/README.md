# Local Infrastructure: Chat V1

This folder owns local-only infrastructure for the real-time chat learning system.

## Services

| Service | Port | Purpose |
|---|---:|---|
| PostgreSQL | `5432` | Metadata, message history, and delivery status. |
| Redis | `6379` | Sessions, presence, sequence counters, dedup keys. |
| Kafka | `9092` | Async event handoff for delivery fan-out. |

Application ports are reserved for later phases:

| App Process | Port |
|---|---:|
| Chat Service HTTP | `8080` |
| WS Server HTTP | `8081` |
| Chat Service gRPC | `9080` |
| WS Push gRPC | `9081` |

## Commands

```bash
make infra-up
make topics
make infra-ps
make infra-logs
make infra-down
```

Kafka topics are created by `infra/kafka/create-topics.sh` with a one-day local retention policy.

Database migrations and seed scripts are intentionally deferred to Phase 2.
