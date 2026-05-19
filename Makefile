SHELL := /bin/sh
GO_CACHE ?= $(CURDIR)/.cache/go-build

.PHONY: infra-up infra-down infra-ps infra-logs topics migrate seed test test-backend test-backend-race vet-backend build-web run-chat-service run-ws-server run-delivery-consumer run-status-consumer

infra-up:
	docker compose up -d

infra-down:
	docker compose down

infra-ps:
	docker compose ps

infra-logs:
	docker compose logs -f

topics:
	docker compose up kafka-init

migrate:
	@if [ -x ./infra/scripts/migrate.sh ]; then \
		./infra/scripts/migrate.sh; \
	else \
		echo "Phase 2 will add database migration scripts."; \
	fi

seed:
	@if [ -x ./infra/scripts/seed.sh ]; then \
		./infra/scripts/seed.sh; \
	else \
		echo "Phase 2 will add local seed scripts."; \
	fi

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
