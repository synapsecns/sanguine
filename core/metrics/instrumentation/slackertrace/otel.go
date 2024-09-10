package slackertrace

import (
	"context"
	"github.com/hedzr/log"
	"github.com/slack-io/slacker"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"time"
)

const (
	instrumentationName    = "github.com/synapsecns/sanguine/core/metrics/instrumentation/slcakertracer"
	instrumentationVersion = "0.1.0"
)

// assumes method is only called once.
type otelRecorder struct {
	attemptsCounter       metric.Int64UpDownCounter
	totalDuration         metric.Int64Histogram
	activeRequestsCounter metric.Int64UpDownCounter
}

func newOtelRecorder() otelRecorder {
	otr := otelRecorder{}
	meter := otel.Meter(instrumentationName, metric.WithInstrumentationVersion(instrumentationVersion))

	// todo: make an option
	metricName := func(metricName string) string {
		return metricName
	}

	var err error
	otr.attemptsCounter, err = meter.Int64UpDownCounter(metricName("slacker.request_count"), metric.WithDescription("Number of Requests"), metric.WithUnit("Count"))
	if err != nil {
		log.Warnf("failed to create counter: %v", err)
	}
	otr.totalDuration, err = meter.Int64Histogram(metricName("slacker.duration"), metric.WithDescription("Time Taken by request"), metric.WithUnit("Milliseconds"))
	if err != nil {
		log.Warnf("failed to create histogram: %v", err)
	}

	otr.activeRequestsCounter, err = meter.Int64UpDownCounter("slacker.active_requests")
	if err != nil {
		log.Warnf("failed to create histogram: %v", err)
	}

	return otr
}

// AddRequests increments the number of requests being processed.
func (r *otelRecorder) AddRequests(ctx context.Context, quantity int64, attributes []attribute.KeyValue) {
	r.attemptsCounter.Add(ctx, quantity, metric.WithAttributes(attributes...))
}

// AddInflightRequests increments and decrements the number of inflight request being processed.
func (r *otelRecorder) AddInflightRequests(ctx context.Context, quantity int64, attributes []attribute.KeyValue) {
	r.activeRequestsCounter.Add(ctx, quantity, metric.WithAttributes(attributes...))
}

// ObserverCommandDuration measures the duration of an HTTP request.
func (r *otelRecorder) ObserverCommandDuration(ctx context.Context, duration time.Duration, attributes []attribute.KeyValue) {
	r.totalDuration.Record(ctx, int64(duration/time.Millisecond), metric.WithAttributes(attributes...))
}

// MetricsMiddleware returns a slacker middleware that records metrics for each command.
func MetricsMiddleware() slacker.CommandMiddlewareHandler {
	// assumes method is only called once.
	otr := newOtelRecorder()

	return func(next slacker.CommandHandler) slacker.CommandHandler {
		return func(cmdCtx *slacker.CommandContext) {
			attributes := []attribute.KeyValue{
				attribute.String("command", cmdCtx.Definition().Command),
			}
			start := time.Now()
			otr.AddInflightRequests(cmdCtx.Context(), 1, attributes)
			otr.AddRequests(cmdCtx.Context(), 1, attributes)
			defer func() {
				otr.AddInflightRequests(cmdCtx.Context(), -1, attributes)
			}()
			next(cmdCtx)
			otr.ObserverCommandDuration(cmdCtx.Context(), time.Since(start), attributes)
		}
	}
}
