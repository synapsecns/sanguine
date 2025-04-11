# CLAUDE.md

This file contains instructions and automated commands for Claude to run when working with this codebase.

## Commands to Run After Code Changes

### Building

```
go build
```

### Linting

```
go mod tidy
go fmt ./...
```

Note: The full linting system requires additional dependencies outside this directory. 
For comprehensive linting, run:

```
make lint
```

### Testing

```
go test ./...
```

## Codebase Overview

OpBot is a Slack bot that provides RFQ (Request for Quote) functionality via Slack commands.

Main components:
- `botmd/`: Bot server implementation
- `cmd/`: Command line interface
- `config/`: Configuration management
- `internal/`: Internal utilities and clients
- `sql/`: Database implementations

## Common Tasks

### Adding a New Command

1. Create a command function in `botmd/commands.go`
2. Register the command in `botmd/botmd.go`
3. Update README.md with usage details