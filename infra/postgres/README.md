# PostgreSQL

Phase 3 owns the PostgreSQL schema, seed data, and query-path documentation for
the chat V1 backend. PostgreSQL can be Neon, local Postgres, or any compatible
Postgres 16+ instance. The scripts only require `POSTGRES_DSN` and a local
`psql` client. They can find `psql` from `PATH`, Homebrew `libpq`, or Homebrew
`postgresql@16`.

## Neon setup

1. Copy your Neon connection string from the Neon console.
2. Put it in `.env` as `POSTGRES_DSN`.
3. Keep the value quoted because Neon strings commonly include `&`.

Example:

```bash
POSTGRES_DSN="postgresql://role:password@ep-example-pooler.region.aws.neon.tech/chat_v1?sslmode=require&channel_binding=require"
```

Neon requires SSL/TLS, so the connection string should include `sslmode=require`.

## Commands

```bash
make migrate
make seed
make db-smoke
```

The scripts are intentionally idempotent for local/dev reruns. `migrate` applies
the schema files in `infra/postgres/migrations/`, `seed` applies deterministic
seed data from `infra/postgres/seeds/`, and `db-smoke` verifies tables, indexes,
seed users, conversations, and a message-history query.

The scripts read real runtime values from `.env`, `.env.local`, `POSTGRES_DSN`,
or `DATABASE_URL`. `.env.example` is documentation only and is not used as an
execution fallback.

## Schema files

Phase 3 adds SQL for:

- `users`
- `conversations`
- `conversation_members`
- `media_objects`
- `messages`
- `message_deliveries`

Use parameterized queries from Go adapters. Do not couple domain/use-case packages to SQL rows.

## Local Postgres fallback

The main `docker-compose.yml` starts Redis and Kafka only. To start a local
Postgres container as well:

```bash
make infra-up-local-postgres
```
