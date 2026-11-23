# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Layout

- `backend/` — the only implemented part: a Go app that embeds [PocketBase](https://pocketbase.io) as a framework (module `homebase`). Run all Go commands from `backend/`.
- `frontend/` — empty placeholder.

The domain is HOA/community management: associations, properties, memberships, assessments, ledger/payments, announcements, documents.

## Commands (run in `backend/`)

- `make run` — `go run . serve --http=localhost:8080` (admin UI at `/_/`)
- `go run . seed` — populate the DB with demo data (`seeds/seed.go`; all users share the password in that file)
- `go run . truncate-all` — delete all records in every non-system collection, keeping schemas (prompts for `yes`; flags `skipConfirm`/`includeSystem` are declared in `main.go`)
- `go build ./...` / `go vet ./...` — no tests or linter config exist yet

## Architecture

- `main.go` wires PocketBase, registers `migratecmd` with `Automigrate` enabled only under `go run` (so collection edits made in the admin UI during dev auto-generate Go files in `migrations/`), and adds the custom `seed` and `truncate-all` cobra commands. Migrations are loaded via the blank import `_ "homebase/migrations"`.
- `migrations/` — the schema is defined here as Go migrations (a full `collections_snapshot` plus incremental `updated_*` files). Don't hand-edit applied migrations; make changes through the admin UI (generates a new migration) or add a new migration file.
- **Authorization lives in PocketBase collection API rules, not Go code.** Every tenant-scoped collection (memberships, properties, assessment_schedules, ledger_entries, payments, announcements, documents, ...) carries an `association` relation, and its list/view/update rules check the caller's row in `memberships` via the `@collection.memberships:mine` alias — requiring `status = "active"` and a role (`admin`, `board`, `resident`) appropriate to the operation. When adding a collection, copy this pattern so data stays isolated per association. Documents additionally filter on a `visibility` field.
- `seeds/seed.go` builds a fixed scenario for exercising those rules: two associations (A, B), users such as outsider, admin, board, residents, former/disabled members, and a `mixed` user (active admin in A, disabled resident in B). It runs in a single transaction and also uploads sample PDFs. Keep it in sync when schema changes.
- `pb_data/` (SQLite DB and uploaded files) is gitignored via `backend/.gitignore`; `.env*` is ignored at the repo root except `.env.example`.
