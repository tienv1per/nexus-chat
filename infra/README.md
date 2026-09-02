# Local Infrastructure: Chat V1

This folder owns local-only infrastructure for the real-time chat learning system.

## Services

| Service | Port | Purpose |
|---|---:|---|
| Redis | `6379` | Sessions, presence, sequence counters, dedup keys. |
| Kafka | `9092` | Async event handoff for delivery fan-out. |

PostgreSQL is external by default for Phase 3. Use Neon by setting `POSTGRES_DSN`
in `.env`:

```bash
POSTGRES_DSN="postgresql://role:password@ep-example-pooler.region.aws.neon.tech/chat_v1?sslmode=require&channel_binding=require"
```

Neon requires SSL/TLS, so keep `sslmode=require` in the connection string.
The database scripts read `.env`, `.env.local`, `POSTGRES_DSN`, or
`DATABASE_URL`; `.env.example` is documentation only.
If you want a local PostgreSQL container instead, use the optional compose
override:

```bash
make infra-up-local-postgres
```

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
make migrate
make seed
make db-smoke
make infra-ps
make infra-logs
make infra-down
```

Kafka topics are created by `infra/kafka/create-topics.sh` with a one-day local retention policy.

Database schema, seed data, and smoke validation live under `infra/postgres/`.
