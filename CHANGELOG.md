# Changelog

All notable changes to this project are documented in this file.

## [Unreleased]

### Added

- **Dual CLI entrypoints**
  - **Standard library (`cmd/task-stdlib`)** — command dispatch with `os.Args` via `internal/handlers`; **no Cobra** in this binary.
  - **Cobra (`main.go`, `cmd/task-cobra`, package `cobracli`)** — subcommands, built-in help, and the same handler layer.
- Shared **`internal/handlers`** package so both flows implement identical behavior (add, list, update, delete, mark-in-progress, mark-done).
- **`CHANGELOG.md`** and **`PROBLEM_DEFINITION.md`** — original problem statement moved out of the main README; **`README.md`** now focuses on usage, commands, and indexing.

### Changed

- **`updatedAt`** is set when a task’s description or status is updated (see `models`).

### Documentation

- **README.md** — index table, how to run stdlib vs Cobra, full command list and examples.
- **PROBLEM_DEFINITION.md** — preserved assignment text from the previous README.

### Notes

- Data file: **`task_master.json`** in the process working directory.
- The Cobra-based tools add a **`require`** on `github.com/spf13/cobra` in `go.mod`; the **`cmd/task-stdlib`** binary does not import Cobra.
