# pkg/db AGENTS.md

Guidance for coding agents working in `pkg/db/`.

## Overview

This package uses:

- `sqlc` for query definitions and generated Go code
- `database/sql` with the `pgx` driver for runtime access
- embedded `golang-migrate` SQL files for schema changes

Do not reintroduce Bun patterns here.

## File Layout

- `sqlc/`: hand-written SQL query inputs for `sqlc`
- `queries/`: generated Go code from `sqlc` output, do not edit by hand
- `migrations/*.up.sql`: forward schema migrations
- `migrations/*.down.sql`: rollback schema migrations
- `migrations/migrations.go`: migration runner and Bun-to-`golang-migrate` reconciliation
- `migrations/create.go`: helper for creating paired migration files

## Adding or Changing Queries

1. Edit or add SQL in `pkg/db/sqlc/`.
2. Prefer explicit named queries and stable column names because service code depends on the generated field names.
3. Regenerate code with:

```bash
./dev/gen-sqlc
```

4. Review the generated files under `pkg/db/queries/`.
5. Update calling code and tests together.

Rules:

- Never hand-edit files in `pkg/db/queries/`.
- Keep query behavior aligned with existing API/service semantics.
- If generated parameter names are poor, prefer improving the SQL with `sqlc.arg(...)` names rather than working around bad generated names in Go.

## Adding Migrations

Create a new migration pair with:

```bash
./dev/create-migration <name>
```

This writes paired `.up.sql` and `.down.sql` files into `pkg/db/migrations/`.

Rules:

- Treat migrations as append-only once merged.
- Keep fresh-database boot and legacy Bun-database upgrade paths in mind.
- Do not change the Bun reconciliation baseline casually. It is intentionally pinned to the Bun handoff version so future `golang-migrate`-only migrations still run on upgraded deployments.

## Legacy Upgrade Model

There are two migration bookkeeping systems you may encounter:

- `bun_migrations`: legacy Bun metadata, no longer used by the server
- `schema_migrations`: `golang-migrate` metadata, current source of truth

For legacy Bun-initialized databases, reconciliation means:

- detect that the legacy application schema already exists
- create `schema_migrations` if needed
- record the fixed Bun handoff version
- leave the application tables untouched

That behavior is implemented in `pkg/db/migrations/migrations.go`. If you change it, update docs and tests in the same change.

## Validation

For DB-layer changes, run:

```bash
./dev/gen-sqlc && git diff --exit-code pkg/db/queries
$(go env GOPATH)/bin/golangci-lint run --timeout=5m --config dev/.golangci.yaml
go test -p 1 ./...
./dev/build
```

If you only changed DB queries or migrations, still prefer running the full serial test suite because DB-backed packages share upgrade assumptions.
