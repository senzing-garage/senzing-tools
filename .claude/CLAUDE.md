# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`senzing-tools` is a Go CLI application that aggregates multiple Senzing tools as subcommands. It uses the Cobra/Viper framework for command-line interface handling. Each subcommand (check-self, demo-entity-search, explain, init-database, load, move, observe, playground, serve-grpc, serve-http, validate) is imported from a separate `senzing-garage` repository and registered via `init()` functions in `cmd/`.

## Build Commands

```bash
# Install development dependencies (one-time)
make dependencies-for-development

# Update Go dependencies
make dependencies

# Build (output to target/<os>-<arch>/senzing-tools)
make clean build

# Run tests
make clean setup test

# Run linting (golangci-lint + govulncheck + cspell)
make lint

# Run individual linters
make golangci-lint
make govulncheck
make cspell

# Fix linting issues
make fix

# Code coverage
make clean setup coverage

# Build Docker image
make docker-build

# Build RPM/DEB packages
make package
```

## Running the Binary

Requires the Senzing C library. Set the library path before running:

```bash
# Linux
export LD_LIBRARY_PATH=/opt/senzing/er/lib/
./target/linux-amd64/senzing-tools [command]

# macOS
export DYLD_LIBRARY_PATH=/opt/senzing/er/lib/:/opt/senzing/er/lib/macos
./target/darwin-amd64/senzing-tools [command]
```

## Running Tests

```bash
# Run all tests
go test ./...

# Run a specific test
go test -run TestName ./...

# Run tests with verbose output
go test -v ./...

# Run tests in a specific package
go test ./cmd/...
```

## Architecture

- **main.go**: Entry point, calls `cmd.Execute()`
- **cmd/root.go**: Root Cobra command setup, configuration loading via Viper (searches `~/.senzing-tools/`, `~/`, `/etc/senzing-tools/` for `senzing-tools.yaml`)
- **cmd/*.go**: Each file registers a subcommand from an external `senzing-garage` repository by calling `RootCmd.AddCommand()` in `init()`
- **cmd/context_*.go**: Platform-specific build constraints (linux, darwin, windows)
- **helper/**: Utility functions (e.g., version string formatting)
- **makefiles/**: Platform-specific Make includes (`linux_x86_64.mk`, `darwin_arm64.mk`, etc.)

## Code Style

- Uses golangci-lint with extensive linters enabled (see `.github/linters/.golangci.yaml`)
- Max line length: 120 characters
- Uses gofumpt for formatting
- Tests use `testify/require` for assertions
- Test functions follow pattern: `Test_FunctionName_scenario`

## Environment Variables

- `SENZING_TOOLS_COMMAND`: Override the subcommand to run
- `LD_LIBRARY_PATH` (Linux) / `DYLD_LIBRARY_PATH` (macOS): Path to Senzing libraries
- `SENZING_TOOLS_DATABASE_URL`: Database connection string for database-related commands
- `SENZING_TOOLS_ENGINE_CONFIGURATION_JSON`: Full Senzing engine configuration (overrides DATABASE_URL)
