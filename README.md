# Homebase

HOA / community management backend built on [PocketBase](https://pocketbase.io) as a Go framework. It models associations, properties, memberships, assessments, ledger entries and payments, announcements, and documents, with per-association data isolation enforced by PocketBase API rules.

## Repository layout

| Path | Description |
| --- | --- |
| `backend/` | Go app embedding PocketBase (module `homebase`). The only implemented part. |
| `frontend/` | Empty placeholder. |

## Getting started

Prerequisite: Go (see `backend/go.mod` for the required version).

All commands run from `backend/`:

```sh
cd backend
make run                # go run . serve --http=localhost:8080
```

- API: <http://localhost:8080/api/>
- Admin UI: <http://localhost:8080/_/> (create a superuser on first launch)

### Other commands

```sh
go run . seed           # populate the DB with demo data (seeds/seed.go)
go run . truncate-all   # delete all records in non-system collections, keeping schemas (asks for "yes")
go build ./...
go vet ./...
```

`truncate-all` accepts `skipConfirm` and `includeSystem` flags (declared in `main.go`). There are no tests or linter config yet.

## Architecture

- **`main.go`** wires up PocketBase, registers `migratecmd` (auto-migrate is on only under `go run`), and adds the `seed` and `truncate-all` commands.
- **`migrations/`** defines the schema as Go migrations: a full `collections_snapshot` plus incremental `updated_*` files. Don't hand-edit applied migrations. Change collections in the admin UI during dev (which generates a new migration) or add a new migration file.
- **Authorization** lives in PocketBase collection API rules, not Go code. Every tenant-scoped collection has an `association` relation. Its list/view/update rules look up the caller's row in `memberships` (via `@collection.memberships:mine`), requiring `status = "active"` and a suitable role (`admin`, `board`, `resident`). Documents also filter on `visibility`. Copy this pattern when adding a collection.
- **`seeds/seed.go`** builds a fixed scenario for exercising those rules: two associations (A and B) and users such as an outsider, admin, board member, residents, former/disabled members, and a `mixed` user (active admin in A, disabled resident in B). It runs in one transaction and uploads sample PDFs. Keep it in sync with schema changes. All seeded users share the password defined in that file.
- **`pb_data/`** (SQLite DB and uploaded files) is gitignored.

## Domain model

Associations, properties, memberships (with roles), residents, assessment schedules, ledger entries, payments, notes, announcements, and documents.
