# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Project Does

Ringer is a universal package manager and workstation setup tool. It abstracts platform-specific package managers (Homebrew on macOS/Linux, Winget on Windows) behind a single CLI, and supports **guise files** — YAML declarations of desired packages to install on a system.

## Commands

```bash
# Build
go build -o ringer .

# Lint and static analysis (run before committing)
go vet ./...
golangci-lint run

# Format code
gofmt -s -w .
goimports -w .

# Run the CLI locally
go run . add <package>
go run . remove <package>
go run . guise sample.guise.yml
go run . platforms
```

Pre-commit hooks enforce `gofmt`, `goimports`, `go vet`, and `golangci-lint` automatically.

## Architecture

### Data Flow

```
CLI input (cmd/)
  → Package lookup (io/package.go) — resolves name to YAML definition in data/packages/
  → Platform Broker (system/platforms/platform.go) — selects the right package manager
  → Platform implementation (homebrew.go / winget.go) — executes system commands
```

### Key Packages

- **`cmd/`** — Cobra commands: `add`, `remove`, `guise`, `platforms`. Entry point via `main.go → cmd.Execute()`.
- **`io/`** — Data structures and loading: `Package` (per-platform name mappings, loaded from `data/packages/`), `Guise` (list of package names from a `.yml` config file).
- **`system/`** — OS detection (`GetSystemInfo()`). Darwin/Linux/Windows each have their own file. Returns a `SystemInfo` struct used by the broker.
- **`system/platforms/`** — The `Platform` interface (`AddPackage`, `RemovePackage`, `EnabledForSystem`, etc.) and the `Broker` struct that selects among registered platforms. `NewBroker()` auto-detects and registers applicable platforms.
- **`data/packages/`** — One YAML file per supported package. Maps a canonical name to platform-specific package IDs (e.g., `homebrew.name`, `windows.name`).

### Adding a New Package

Create a YAML file in `data/packages/` following this structure:

```yaml
name: "Display Name"
description: "Brief description"
platforms:
  homebrew:
    name: "brew-formula-or-cask"
  windows:
    name: "Publisher.PackageId"
```

### Adding a New Platform

1. Implement the `Platform` interface in `system/platforms/`.
2. Register it in `NewBroker()` in `platform.go`.

### Known TODOs (from code comments)

- Package validity checking is not yet implemented.
- Linux-native package managers (apt, yum, pacman) are not supported yet; only Homebrew on Linux.
- No state tracking for where packages were installed.
- Platform preference fallback (currently only tries the preferred platform, not alternatives).
