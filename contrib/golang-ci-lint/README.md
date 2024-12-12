# Golang CI Lint Version Manager

This tool manages the installation of golangci-lint versions specified in the root .golangci-version file.

## Features
- Architecture-specific binary downloads (supports both AMD64 and ARM64)
- Version pinning via .golangci-version file
- Binary caching to avoid redundant downloads
- Automatic version management through make lint

## Usage
The tool is automatically invoked by `make lint`. To manually run:
```go
go run contrib/golang-ci-lint
```
