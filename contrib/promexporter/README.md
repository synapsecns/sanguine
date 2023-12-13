# Prom Exporter

[![Go Reference](https://pkg.go.dev/badge/github.com/synapsecns/sanguine/contrib/promexporter.svg)](https://pkg.go.dev/github.com/synapsecns/sanguine/contrib/promexporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/synapsecns/sanguine/contrib/promexporter)](https://goreportcard.com/report/github.com/synapsecns/sanguine/contrib/promexporter)


The prometheus exporter is a simple HTTP server that exposes metrics derived from several public apis. The metrics are exposed in the prometheus format and can be scraped by a prometheus server.

Grafana dashbaords are in the `/dashboards` directory.

## Running Locally

Make sure `METRICS_PROVIDER` is set to `OTLP` otherwise this will no-op

