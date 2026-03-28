# Task Tracker CLI

A small Go command-line tool to manage tasks. Tasks are stored in **`task_master.json`** in the **current working directory** (the file is created automatically).

## Index

| Section / doc | Description |
|-----------------|-------------|
| [Requirements](#requirements) & [Setup](#setup) | Go version; setup for **this repo** vs **starting from scratch** |
| [PROBLEM_DEFINITION.md](PROBLEM_DEFINITION.md) | Original assignment / problem statement and requirements |
| [CHANGELOG.md](CHANGELOG.md) | Project changes and history |

## Requirements

- [Go](https://go.dev/dl/) **1.21 or newer** (this module uses `go 1.26.1`; use a matching or newer toolchain).

## Setup

### Using this repository

Follow these steps to run or build **this** project as it exists today.

1. **Install Go** from the [official downloads](https://go.dev/dl/) and confirm it is on your `PATH`:
   ```bash
   go version
   ```
2. **Clone or unpack** the repo and `cd` to the **project root** (the folder that contains `go.mod`):
   ```bash
   cd /path/to/task_tracker_cli
   ```
3. **Dependencies** — versions are pinned in **`go.mod`** / **`go.sum`**. You do **not** run `cobra-cli` or `go get github.com/spf13/cobra` just to use this codebase.
   - Optional: `go mod download` (e.g. for CI or offline builds). Otherwise **`go run`** / **`go build`** fetch modules on first use.
4. **Run or build** from the project root:
   ```bash
   go run . --help
   go run ./cmd/task-stdlib help
   ```
   ```bash
   go build -o task-cobra .
   go build -o task-stdlib ./cmd/task-stdlib
   ```
5. **Pick a working directory** for your data file — `task_master.json` is created in the **current working directory** when you invoke the binary:
   ```bash
   cd ~/my-tasks
   /path/to/task-cobra list
   ```

To bump Cobra later (optional): `go get github.com/spf13/cobra@latest` → `go mod tidy` → re-test.

### Starting from scratch (new project)

Use this path only if you are **creating a new CLI** (for learning or a different app), not when cloning this repo.

1. **Install Go** (same as above).
2. **Create a module** in a new empty directory:
   ```bash
   mkdir mycli && cd mycli
   go mod init example.com/mycli
   ```
3. **Optional — Cobra scaffolding** — install the generator and lay out commands:
   ```bash
   go install github.com/spf13/cobra-cli@latest
   cobra-cli init
   cobra-cli add serve   # example; repeat for each subcommand
   ```
   Add Cobra to the module if the scaffold did not pull it:
   ```bash
   go get github.com/spf13/cobra@latest
   ```
4. **Implement** your logic (e.g. JSON storage, handlers), wire commands, then `go build` or `go run .`.

This repository already has **`cobracli`**, **`cmd/task-cobra`**, **`cmd/task-stdlib`**, and **`internal/handlers`**; you would only repeat the “from scratch” flow to practice or to bootstrap a **different** repo. See the [Cobra user guide](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md) for more detail.

## Two ways to run the CLI

This repository ships **two entrypoints** that share the same behavior (same subcommands and JSON storage):

| Mode | Description | Build / run |
|------|-------------|-------------|
| **Standard library** | Parsing with `os.Args` only — **no Cobra** in this binary | `go run ./cmd/task-stdlib …` or `go build -o task-stdlib ./cmd/task-stdlib` |
| **Cobra** | Subcommands and help via [Cobra](https://github.com/spf13/cobra) | `go run . …`, `go run ./cmd/task-cobra …`, or `go build -o task-cobra .` |

Shared logic lives in **`internal/handlers`**. The default **`go run .`** and **`main.go`** at the repo root use the **Cobra** stack (`cobracli`).

## Commands

Run these after any of: `go run .`, `go run ./cmd/task-cobra`, or `go run ./cmd/task-stdlib` (replace the prefix with your built binary name if you use `go build`).

| Action | Command |
|--------|---------|
| Add a task | `<binary> add "Description"` |
| List all tasks | `<binary> list` |
| List by status | `<binary> list` `done` \| `todo` \| `in-progress` |
| Update description | `<binary> update <id> "New description"` |
| Delete a task | `<binary> delete <id>` |
| Mark in progress | `<binary> mark-in-progress <id>` |
| Mark done | `<binary> mark-done <id>` |
| Help (stdlib) | `<binary> help` or `<binary> -h` or `<binary> --help` |
| Help (Cobra) | `<binary> --help` or `<binary> <command> --help` |

### Examples

```bash
# From repo root — Cobra (default main)
go run . add "Buy groceries"
go run . list
go run . list done
go run . update 0 "Buy groceries and cook dinner"
go run . mark-in-progress 0
go run . mark-done 0
go run . delete 0

# Standard library only
go run ./cmd/task-stdlib add "Buy groceries"
go run ./cmd/task-stdlib list
```

Use a **quoted** description when it contains spaces. Task **IDs** start at **0** in this implementation (see `task_master.json`).
