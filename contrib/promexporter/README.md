# Prom Exporter

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/contrib/promexporter.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/contrib/promexporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/contrib/promexporter)](https://goreportcard.com/report/github.com/synapsecns/sanguine/contrib/promexporter)


The prometheus exporter is a simple HTTP server that exposes metrics derived from several public apis. The metrics are exposed in the prometheus format and can be scraped by a prometheus server.

Grafana dashbaords are in the `/dashboards` directory.

## Running Locally

Make sure `METRICS_HANDLER` is set to `OTLP` otherwise this will no-op

## Building

### Using the build script

For convenience, a build script is provided that includes the necessary flags to handle the memsize compatibility issue with Go 1.22+:

```bash
./build.sh ./...
```

### Manual build with flags

If you encounter the error `link: github.com/fjl/memsize: invalid reference to runtime.stopTheWorld`, you need to use the following build command:

```bash
go build -ldflags=-checklinkname=0 ./...
```

This is because the `github.com/fjl/memsize` package attempts to access internal runtime functions that are restricted in Go 1.22+ (specifically `runtime.stopTheWorld`). The `-checklinkname=0` flag disables this restriction.

