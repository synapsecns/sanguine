# Promexporter Project Guide

This document contains important information about the promexporter project to help Claude assist with development tasks.

## Project Overview

The promexporter is a Prometheus metrics exporter for the Synapse ecosystem. It exports various metrics related to blockchain bridges, validators, gas balances, and token information across different chains.

## Build Commands

```bash
# Build the project
go build

# Run tests
go test ./...

# Generate code
go generate ./...

# Lint the code
make lint
```

## Project Structure

- **cmd/**: Command-line interface for the exporter
- **config/**: Configuration parsing and management
- **exporters/**: Core metric exporters for various components
  - **bridge.go**: Bridge-related metrics
  - **exporter.go**: Main exporter implementation
  - **otel.go**: OpenTelemetry metrics setup
  - **relayer.go**: Relayer metrics
  - **submitter.go**: Submitter metrics
- **internal/**: Internal packages
  - **decoders/**: Data decoders
  - **gql/**: GraphQL clients
  - **types/**: Type definitions, including ChainID

## Key Concepts

- **ChainID**: Enum type representing supported blockchain networks
- **Metrics**: The project uses OpenTelemetry for metrics collection and reporting
- **TokenConfig**: Configuration for tokens being monitored
- **Submitter**: Components that submit transactions to blockchains
- **Bridge**: Cross-chain bridge components

## Common Tasks

### Adding a New Chain

1. Add the new chain to `internal/types/chainID.go`
2. Run `go generate ./...` to update the string representation
3. Add corresponding bridge addresses to the configuration

### Adding a New Metric

1. Add the metric to the appropriate exporter file
2. Register the metric in the otelRecorder initialization
3. Create a callback function for the metric
4. Register the callback in the otelRecorder initialization

## Important Notes

- The project uses a hashmap-based approach for storing metrics
- Metrics are periodically collected in the `recordMetrics` function
- Always regenerate code after modifying interfaces