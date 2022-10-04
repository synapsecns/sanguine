package prom

import (
	"fmt"
	"strings"
)

// PrometheusDataSourceName is the name of the data source in grafana to use for prometheus.
const PrometheusDataSourceName = "Prometheus"

// NodeIDVar is the variable used for node id.
const NodeIDVar = "node_id"

// ChainIDVar is the variable used for the chain id.
const ChainIDVar = "chain_id"

// TimeSeriesFormat is the format used for timeseries (in sdk.target).
const TimeSeriesFormat = "time_series"

// TimeSeriesGraphType is the graph type used for timeseries.
const TimeSeriesGraphType = "timeseries"

// MetricToTitle formats a metric as a title.
func MetricToTitle(metric string) string {
	//nolint: staticcheck
	return strings.Title(strings.ReplaceAll(metric, "_", " "))
}

// WrapMetricJobQuery wraps the metric in a query around job.
func WrapMetricJobQuery(metric string) string {
	return fmt.Sprintf("%s{%s=~\"^($%s)$\"}", metric, NodeIDVar, NodeIDVar)
}
