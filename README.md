# Gonv

Secure environment-variable management for development teams.

## Overview

**Gonv** is a web platform for storing, organizing, and sharing development environment variables in a managed way—instead of scattering secrets across personal `.env` files, chat threads, and password managers.

The goal is a single place where you can:

- Keep env vars **centralized** per project / environment (local, staging, etc.)
- Control **who can read or edit** values
- Reduce accidental leaks from copy-paste and outdated local files
- Pull a consistent set of variables into your local workflow when you need them

This repository is the Gonv backend (and, soon, the embedded SPA). Auth-oriented data models (`User`, `Account`, `Session`) are already in place; the HTTP API currently exposes health/readiness while the product surface grows.

## Stack

| Layer | Choice |
|--------|--------|
| API | Go, [Gin](https://github.com/gin-gonic/gin), [Huma](https://github.com/danielgtaylor/huma) |
| Data | PostgreSQL, [Ent](https://entgo.io/) (versioned migrations → goose SQL) |
| Config | `.env` via `godotenv` (non-release) |
| Frontend (planned) | React + Vite, built into `www/dist` and embedded with `www/embed.go` |
| Local edge (optional) | Caddy (`Caddyfile` → `gonvy.localhost` → `:8080`) |

Module path: `gonvy.internal`.

## Status

Early scaffolding:

- HTTP server with `healthz` / `readyz`
- Ent schemas for users, OAuth/credential accounts, and sessions
- Docker Compose Postgres, Makefile workflow, migration tooling
- React/Vite UI and env-var domain APIs are **not** shipped yet

## Prerequisites

- Go **1.26+** (see `go.mod` / toolchain)
- Docker (for Postgres)
- [goose](https://github.com/pressly/goose) (apply migrations)
- Optional: [golangci-lint](https://golangci-lint.run/), Node/npm (when `web/` exists), [Caddy](https://caddyserver.com/)

## Quick start

```bash
cp .env.example .env
make up          # start Postgres
make ent         # regenerate Ent code from schemas
make migrate-up  # apply SQL migrations (once you have some)
make run         # API on :8080 (or PORT from .env)
```

Check liveness:

```bash
curl -s http://localhost:8080/healthz
curl -s http://localhost:8080/readyz
```

Optional: point Caddy at the API with the repo `Caddyfile` (`gonvy.localhost` → `localhost:8080`).

Reset DB volume, migrate, and run:

```bash
make fresh
```

## Configuration

Copy `.env.example` → `.env`. Relevant variables:

| Variable | Purpose |
|----------|---------|
| `PORT` | HTTP listen port (default `8080`) |
| `DATABASE_URL` | App + goose target database |
| `DEV_DATABASE_URL` | Clean DB used by Ent `NamedDiff` when generating migrations |

Never commit real secrets; `.env` is gitignored.

## Common Make targets

Infrastructure and day-to-day:

| Target | What it does |
|--------|----------------|
| `make up` / `make down` | Start / stop Compose Postgres |
| `make infra-reset` | Wipe the Postgres volume (destructive) |
| `make run` | Run `./cmd/server` |
| `make build` | Binary → `bin/server` |
| `make dev` | `infra-up` then `run` |
| `make fresh` | Reset DB, migrate up, run |

Codegen and quality:

| Target | What it does |
|--------|----------------|
| `make ent` / `make generate` | Ent (and other) `go generate` |
| `make fmt` `vet` `lint` `test` | Formatting and checks |
| `make check` | `fmt` + `vet` + `lint` + `test` |

Migrations:

| Target | What it does |
|--------|----------------|
| `make migrate-diff name=add_foo` | Generate goose SQL via Ent (uses `DEV_DATABASE_URL`) |
| `make migrate-up` / `down` / `status` | Apply / roll back / inspect with goose |
| `make migrate-reset` | Roll back all migrations (destructive) |

Frontend (when scaffolded under `web/`):

| Target | What it does |
|--------|----------------|
| `make www-init` | Create `www/embed.go` + `www/dist` placeholder |
| `make web-install` / `web-dev` / `web-build` | npm install, Vite dev, production build into `www/dist` |
| `make build-release` | `web-build` then release-mode Go binary |

Run `make help` for the annotated list.

## Migrations

Schemas live under `ent/schema/`. Codegen:

```bash
make ent
```

Generate a new versioned migration (Atlas/Ent replay against the **dev** database, goose-compatible SQL under `ent/migrate/migrations/`):

```bash
make migrate-diff name=describe_your_change
```

Apply to the app database:

```bash
make migrate-up
```

SQL files and `embed.go` in `ent/migrate/migrations/` are what get committed; generated Ent packages under `ent/` are largely gitignored and regenerated locally.

## Project layout

```
cmd/server/          API entrypoint
internal/
  config/            Env loading
  di/                App wiring
  httpapi/           Huma/Gin routes & middleware
  httperr/           API error helpers
  persistance/       Postgres / Ent access
ent/
  schema/            Source of truth for the data model
  migrate/           Diff tool + goose migrations
www/                 (planned) embed.go + Vite dist for go:embed
web/                 (planned) React + Vite sources
docker-compose.yaml  Local Postgres
Caddyfile            Optional local reverse proxy
Makefile             Day-to-day workflow
```

## Frontend embedding (planned)

Production will serve a React SPA from the Go binary:

1. Vite app in `web/` with `build.outDir` → `../www/dist` (`emptyOutDir: true`)
2. `www/embed.go` embeds `dist` via `//go:embed all:dist`
3. `make web-build` then `make build-release`

Keeping `embed.go` outside the wiped `dist/` directory avoids losing the embed package on every frontend build.

## License

Proprietary / unset — add a license file when you publish the project.
