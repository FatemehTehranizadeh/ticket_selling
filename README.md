# Ticket Selling (worker-pool demo)

A small Go project that simulates a ticket selling / buying system using a worker-pool pattern.

This repo is intentionally simple: it focuses on concurrency, clean architecture, and testable in-memory stores.
It is designed to be extended later with a persistent database (MySQL), an HTTP API (Echo), or CI pipelines.

---

## Features

- Domain-first design: `domain/` contains pure business types (Event, Order, User).
- `repo/` defines storage interfaces (`EventStore`, `OrderStore`).
- `repo/mem` provides in-memory, concurrency-safe implementations.
- `service/` contains business logic (`MarketService`) that depends on storage interfaces.
- `workerpool/` is a reusable concurrency primitive to process jobs.
- Structured logging with Zap.
- Small CLI demo in `cmd/ticketsvc`.

---

## Requirements

- Go 1.21+ (or latest stable Go)
- (optional) `golangci-lint` for linting
- (optional) GitHub Actions for CI

---

## Quick start (local)

```bash
# clone
git clone https://github.com/<you>/ticket_selling.git
cd ticket_selling

# build & run demo
go mod tidy
go run ./cmd/ticketsvc
