# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Covermyass is a post-exploitation security tool written in Go that helps penetration testers cover their tracks by finding and securely erasing log files on Linux, macOS, and Windows systems. The tool uses secure file shredding (multiple overwrites with random data) to make data recovery extremely difficult.

**Important**: This is a defensive security tool for authorized penetration testing. Only assist with legitimate security research, testing on owned systems, or educational purposes.

## Development Commands

### Build

```bash
make build                    # Build binary to ./bin/covermyass
go build -o ./bin/covermyass  # Alternative build command
```

### Testing

```bash
make test                     # Run all tests with coverage
gotestsum --format testname --junitfile unit-tests.xml -- -mod=readonly -race -coverprofile=./c.out -covermode=atomic -coverpkg=.,./... ./...
make coverage                 # Run tests and show coverage report
```

### Code Quality

```bash
make fmt                      # Format code
make lint                     # Run golangci-lint
make all                      # Run fmt, lint, test, build, and go.mod tidy
```

### Development Setup

```bash
make install-tools            # Install required development tools (gotestsum, mockery)
make mocks                    # Generate mocks using mockery
```

### Clean

```bash
make clean                    # Clean build artifacts
```

## Architecture

### Core Components

**Main Entry Point** (`main.go`, `cmd/root.go`):

- CLI built with Cobra framework
- Handles command-line argument parsing and execution flow
- Supports list mode (--list) and destructive write mode (--write)

**Analysis Engine** (`lib/analysis/`):

- `analyzer.go`: Orchestrates the file discovery and analysis process
- `analysis.go`: Data structures for storing scan results and generating reports
- Coordinates checks, filtering, and result aggregation

**Check System** (`lib/check/`):

- Modular check system with pluggable check implementations
- Each check (e.g., `system_check.go`, `sshd_check.go`) defines log file patterns for specific services
- Checks are registered globally and executed during analysis

**File Discovery** (`lib/find/`):

- `finder.go`: Implements filesystem traversal using glob patterns
- Uses doublestar library for advanced glob pattern matching
- Integrates with filter system to exclude specified paths

**Filtering** (`lib/filter/`):

- `filter.go`: Implements path filtering using glob patterns
- Allows users to exclude certain files/paths from analysis via -f flag

**Secure Deletion** (`lib/shred/`):

- `shred.go`: Implements secure file overwriting
- Overwrites files multiple times with random data (default 3, configurable with -n)
- Optional final zero-fill pass (--zero flag) to hide shredding activity

**Output/UI** (`output/`):

- `printer.go`: Handles formatted console output
- Progress bar integration for shredding operations

### Data Flow

1. **Command Parsing**: CLI arguments parsed in `cmd/root.go`
2. **Filter Setup**: Filter rules applied from --filter flags
3. **Check Registration**: All check modules register their log file patterns
4. **Analysis**: Analyzer runs checks through finder to discover matching files
5. **Filtering**: Results filtered based on user-specified exclusion patterns
6. **Output**: Results displayed with file paths, sizes, and permissions
7. **Shredding** (if --write): Files securely overwritten using configurable iterations

### Test Organization

Tests are organized alongside source files (`*_test.go`). Key test areas:

- Unit tests for each component (analysis, checks, filtering, shredding)
- Mock interfaces generated with mockery (`mocks/` directory)
- Test data in `testdata/` subdirectories

## Key Dependencies

- `github.com/spf13/cobra`: CLI framework
- `github.com/bmatcuk/doublestar/v4`: Advanced glob pattern matching
- `github.com/schollz/progressbar/v3`: Progress indication
- `github.com/sirupsen/logrus`: Structured logging
- `github.com/stretchr/testify`: Testing framework

## Development Notes

- Go 1.18+ required
- Uses Go modules for dependency management
- Binary versioning handled via build-time ldflags in Makefile
- Cross-platform support (Linux, macOS, Windows)
- Logs are structured using logrus with debug-level build information
