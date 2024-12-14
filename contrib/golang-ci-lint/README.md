# Golang CI Lint Version Manager

A standalone tool for managing and running specific versions of golangci-lint across different repositories. This package can be used in any Go project to ensure consistent linting with version pinning, regardless of where it's installed.

## Features
- **Cross-Repository Usage**: Can be installed and used in any Go project
- **Version Pinning**: Automatically uses the version specified in `.golangci-version`
- **Architecture Support**: Downloads correct binary for your system (AMD64/ARM64)
- **Binary Caching**: Avoids redundant downloads of the same version
- **Secure Downloads**: Verifies binary checksums for security
- **CI/CD Integration**: Easy to integrate with GitHub Actions and other CI systems

## Installation

As a standalone tool:
```bash
go install github.com/synapsecns/sanguine/contrib/golang-ci-lint@latest
```

Or in your project's go.mod:
```go
require github.com/synapsecns/sanguine/contrib/golang-ci-lint v1.0.0
```

## Usage

### 1. Version Configuration
Create a `.golangci-version` file in your repository root:
```
1.61.0
```

### 2. Running the Linter

As a standalone binary:
```bash
golang-ci-lint run --fix --config=.golangci.yml
```

Or using go run:
```bash
go run github.com/synapsecns/sanguine/contrib/golang-ci-lint run --fix --config=.golangci.yml
```

### 3. Makefile Integration

Add to your project's Makefile:
```makefile
.PHONY: lint
lint:
	go run github.com/synapsecns/sanguine/contrib/golang-ci-lint run --fix --config=.golangci.yml
```

### 4. CI/CD Integration

GitHub Actions example:
```yaml
name: Lint
on: [push, pull_request]

jobs:
  golangci:
    name: lint
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - name: Install golang-ci-lint manager
        run: go install github.com/synapsecns/sanguine/contrib/golang-ci-lint@latest
      - name: Run linter
        run: golang-ci-lint run --config=.golangci.yml
```

## Common Use Cases

1. **Multiple Projects**: Use the same linter version across all your repositories
2. **CI/CD Pipelines**: Ensure consistent linting in automated workflows
3. **Team Collaboration**: Maintain consistent code style across development teams
4. **Version Control**: Lock linter version to avoid unexpected behavior changes

## Troubleshooting

1. **Binary Not Found**
   ```bash
   # Clear the cache and redownload
   rm -rf ~/.cache/golangci-lint
   golang-ci-lint run
   ```

2. **Version Mismatch**
   - Ensure `.golangci-version` exists in repository root
   - Check file permissions and format

3. **Architecture Issues**
   - The tool automatically detects and downloads the correct binary
   - Supported architectures: linux-amd64, linux-arm64, darwin-amd64, darwin-arm64, windows-amd64
   - MacOS (/private/var) symlink handling:
     - Properly resolves /private/var to /var for temp directories
     - Maintains secure path validation across symlinked paths
     - Handles platform-specific temp directory structures

## Contributing

Contributions are welcome! Please ensure:
- Cross-platform compatibility:
  - Linux: Standard path resolution
  - MacOS: Handles /private/var symlinks and temp directories
  - Windows: Supports standard Windows paths
- Backward compatibility with existing `.golangci-version` files
- Proper error handling and user feedback

## License

MIT License - See LICENSE file for details
