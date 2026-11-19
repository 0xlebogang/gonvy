.DEFAULT_GOAL := help

MODULE      := gonvy.internal
CMD_SERVER  := ./cmd/server
BIN_DIR     := bin
BIN_SERVER  := $(BIN_DIR)/server

ENT_SCHEMA  := ./ent/schema
MIGRATE_DIR := ent/migrate/migrations
MIGRATE_CMD := ./ent/migrate/main.go

WWW_DIR     := www
WWW_DIST    := $(WWW_DIR)/dist
WWW_EMBED   := $(WWW_DIR)/embed.go

GO          ?= go
DOCKER      ?= docker
COMPOSE     ?= docker compose
GOOSE       ?= goose
NPM         ?= npm
GOLINT 	    ?= golangci-lint

GOFLAGS     ?=
LDFLAGS     ?=

# Load .env into the Make environment when present (PORT, DATABASE_URL, …).
ifneq (,$(wildcard .env))
    include .env
    export
endif

DATABASE_URL     ?= postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable
DEV_DATABASE_URL ?= postgresql://postgres:postgres@localhost:5432/dev?sslmode=disable
PORT             ?= 8080

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n" } \
		/^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)


.PHONY: infa-up infa-down infa-ps infa-reset
infra-up: ## Start Postgres (docker compose)
	$(COMPOSE) up -d postgres

infra-down: ## Stop Postgres
	$(COMPOSE) down

infra-ps: ## Show compose service status
	$(COMPOSE) ps

infra-reset: ## Wipe Postgres volume and start fresh (DESTRUCTIVE)
	$(COMPOSE) down -v
	$(COMPOSE) up -d


.PHONY: generate ent tidy fmt lint vet test check
generate:
	$(GO) generate ./...

ent:
	$(GO) generate ./ent

tidy:
	$(GO) mod tidy

fmt:
	$(GO) fmt ./...

lint:
	$(GOLINT) run ./...

vet:
	$(GO) vet ./...

test:
	$(GO) test $(GOFLAGS) ./...

check: fmt vet lint test


.PHONY: migrate-diff migrate-up migrate-down migrate-status migrate-reset
migrate-diff:
	@test -n "$(name)" || (echo "usage: make migrate-diff name=<migration_name>"; exit 1)
	@echo ">> NamedDiff against DEV_DATABASE_URL (clean replay DB)"
	DATABASE_URL="$(DEV_DATABASE_URL)" $(GO) run -mod=mod $(MIGRATE_CMD) "$(name)"

migrate-up:
	$(GOOSE) -dir $(MIGRATE_DIR) postgres "$(DATABASE_URL)" up

migrate-down:
	$(GOOSE) -dir $(MIGRATE_DIR) postgres "$(DATABASE_URL)" down

migrate-status:
	$(GOOSE) -dir $(MIGRATE_DIR) postgres "$(DATABASE_URL)" status

migrate-reset:
	$(GOOSE) -dir $(MIGRATE_DIR) postgres "$(DATABASE_URL)" reset


.PHONY: run build build-release clean
run:
	$(GO) run $(GOFLAGS) $(CMD_SERVER)

build:
	@mkdir -p $(BIN_DIR)
	$(GO) build $(GOFLAGS) -o $(BIN_SERVER) $(CMD_SERVER)

build-release: web-build
	@mkdir -p $(BIN_DIR)
	GIN_MODE=release $(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" -o $(BIN_SERVER) $(CMD_SERVER)

clean:
	rm -rf $(BIN_DIR) $(WWW_DIST)


.PHONY: web-install web-dev web-build web-preview www-init
web-install: ## npm install in $(WEB_DIR)
	@test -d $(WEB_DIR) || (echo "missing $(WEB_DIR)/ — scaffold Vite first"; exit 1)
	cd $(WEB_DIR) && $(NPM) install

web-dev: ## Vite dev server (proxy API separately / via Caddy)
	@test -d $(WEB_DIR) || (echo "missing $(WEB_DIR)/ — scaffold Vite first"; exit 1)
	cd $(WEB_DIR) && $(NPM) run dev

web-build: ## Build SPA into $(WWW_DIST) for go:embed
	@if [ ! -d $(WEB_DIR) ]; then \
		echo ">> skip web-build: $(WEB_DIR)/ not present yet"; \
		exit 0; \
	fi
	@test -f $(WWW_EMBED) || (echo "missing $(WWW_EMBED) — run: make www-init"; exit 1)
	cd $(WEB_DIR) && $(NPM) run build
	@test -d $(WWW_DIST) || (echo "expected Vite outDir $(WWW_DIST)"; exit 1)

web-preview:
	@test -d $(WEB_DIR) || (echo "missing $(WEB_DIR)/ — scaffold Vite first"; exit 1)
	cd $(WEB_DIR) && $(NPM) run preview

www-init:
	@mkdir -p $(WWW_DIST)
	@if [ ! -f $(WWW_EMBED) ]; then \
		printf '%s\n' \
			'package www' \
			'' \
			'import "embed"' \
			'' \
			'// FS is the Vite production build (see Makefile WEB_DIR / WWW_DIST).' \
			'//' \
			'//go:embed all:dist' \
			'var FS embed.FS' \
			> $(WWW_EMBED); \
		echo "wrote $(WWW_EMBED)"; \
	else \
		echo "$(WWW_EMBED) already exists"; \
	fi
	@touch $(WWW_DIST)/.gitkeep
	@echo "Configure Vite: build.outDir = '../www/dist' (emptyOutDir: true)"


.PHONY: up down dev fresh
up: infra-up

down: infra-down

dev: infra-up
	@$(MAKE) run

fresh: infra-reset
	@echo ">> waiting for Postgres…"
	@sleep 3
	@$(MAKE) migrate-up
	@$(MAKE) run
