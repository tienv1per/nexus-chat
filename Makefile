SHELL := /bin/sh
GO_CACHE ?= $(CURDIR)/.cache/go-build

.PHONY: infra-up infra-up-local-postgres infra-down infra-down-local-postgres infra-ps infra-logs topics migrate seed db-smoke test test-backend test-backend-race vet-backend build-web run-chat-service run-ws-server run-delivery-consumer run-status-consumer

infra-up:
	docker compose up -d

infra-up-local-postgres:
	docker compose -f docker-compose.yml -f docker-compose.postgres.yml up -d

infra-down:
	docker compose down

infra-down-local-postgres:
	docker compose -f docker-compose.yml -f docker-compose.postgres.yml down

infra-ps:
	docker compose ps

infra-logs:
	docker compose logs -f

topics:
	docker compose up kafka-init

migrate:
	./infra/scripts/migrate.sh

seed:
	./infra/scripts/seed.sh

db-smoke:
	./infra/scripts/db-smoke-test.sh

test: test-backend

test-backend:
	mkdir -p $(GO_CACHE)
	cd backend && GOCACHE=$(GO_CACHE) go test ./...

test-backend-race:
	mkdir -p $(GO_CACHE)
	cd backend && GOCACHE=$(GO_CACHE) go test -race ./...

vet-backend:
	mkdir -p $(GO_CACHE)
	cd backend && GOCACHE=$(GO_CACHE) go vet ./...

build-web:
	@if [ -f ./web/package.json ]; then \
		cd web && npm run build; \
	else \
		echo "Phase 9 will scaffold the Next.js app."; \
	fi

run-chat-service:
	mkdir -p $(GO_CACHE)
	cd backend && GOCACHE=$(GO_CACHE) go run ./cmd/chat-service

run-ws-server:
	mkdir -p $(GO_CACHE)
	cd backend && GOCACHE=$(GO_CACHE) go run ./cmd/ws-server

run-delivery-consumer:
	mkdir -p $(GO_CACHE)
	cd backend && GOCACHE=$(GO_CACHE) go run ./cmd/delivery-consumer

run-status-consumer:
	mkdir -p $(GO_CACHE)
	cd backend && GOCACHE=$(GO_CACHE) go run ./cmd/status-consumer
