package prom

import (
	"context"
	"github.com/grafana-tools/sdk"
	"github.com/prometheus/client_golang/prometheus"
)

// Config contains the configuration for the metrics handler.
type Config struct {
	// PushGateway url to use (if any)
	PushGateway string `toml:"PushGateway"`
	// Enabled Whether or not to enable metrics
	Enabled bool `toml:"Enabled"`
	// GrafanaHost to push dashboards to (none will be pushed if disabled)
	GrafanaHost string `toml:"GrafanaHost"`
	// GrafanaKey to use to push dashboards
	GrafanaKey string `toml:"GrafanaKey"`
}

// IsValid determines whether or not the metrics config is valid.
func (c *Config) IsValid(ctx context.Context) (ok bool, err error) {
	// TODO implement
	return true, nil
}

// Instrumentable is a type that exports some (or none as an array) prometheus metrics.
type Instrumentable interface {
	// GetMetrics gets the metrics associated with this object.
	GetMetrics(labels map[string]string) []prometheus.Collector
}

// Graphable is an interface that exposes a way for metrics to be registered
// with a the metric handler. It allows the responsibility. of placing the graphs to be used by the client.
type Graphable interface {
	Instrumentable
	// GetGraphs gets the graphs associated with this object.
	GetGraphs() []*sdk.Panel
}
