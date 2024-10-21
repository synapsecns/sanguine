package metrics

import (
	"context"
	"go.opentelemetry.io/otel/metric"
	"time"
)

// standardMetrics records metrics across any service using the metrics handler.
type standardMetrics struct {
	metrics     Handler
	meter       metric.Meter
	uptimeGauge metric.Float64ObservableGauge
	startTime   time.Time
}

const processUptimeSecondsMetric = "process_uptime_seconds"

func newStandardMetrics(ctx context.Context, handler Handler) {
	str := standardMetrics{
		metrics:   handler,
		meter:     handler.Meter("standard_metrics"),
		startTime: time.Now(),
	}

	var err error
	if str.uptimeGauge, err = str.meter.Float64ObservableGauge(processUptimeSecondsMetric, metric.WithDescription("The uptime of the process in seconds"), metric.WithUnit("seconds")); err != nil {
		handler.ExperimentalLogger().Errorf(ctx, "failed to create %s gauge: %v", processUptimeSecondsMetric, err)
	}

	// Register callback
	if _, err = str.meter.RegisterCallback(str.uptimeCallback, str.uptimeGauge); err != nil {
		handler.ExperimentalLogger().Warnf(ctx, "failed to register callback: %v", err)
	}
}

func (str *standardMetrics) uptimeCallback(_ context.Context, observer metric.Observer) error {
	uptimeDuration := time.Since(str.startTime).Seconds()
	observer.ObserveFloat64(str.uptimeGauge, uptimeDuration)
	return nil
}
