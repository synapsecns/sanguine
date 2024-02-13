package metrics

import (
	"context"
	"github.com/gin-gonic/gin"
	experimentalLogger "github.com/synapsecns/sanguine/core/metrics/logger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"net/http"
)

// nullHandler is a metrics handler that does nothing.
// it is used to allow metrics collection to be skipped.
type nullHandler struct {
	tracer     trace.Tracer
	propagator nullPropogator
	meter      Meter
}

func (n nullHandler) ExperimentalLogger() experimentalLogger.ExperimentalLogger {
	return experimentalLogger.NewNullLogger()
}

func (n nullHandler) Meter(name string, options ...metric.MeterOption) metric.Meter {
	return noop.Meter{}
}

func (n nullHandler) Handler() http.Handler {
	return noopHandler{}
}

func (n nullHandler) Type() HandlerType {
	return Null
}

func (n nullHandler) Propagator() propagation.TextMapPropagator {
	return n.propagator
}

func (n nullHandler) GetTracerProvider() trace.TracerProvider {
	return trace.NewNoopTracerProvider()
}

func (n nullHandler) Tracer() trace.Tracer {
	return n.tracer
}

func (n nullHandler) AddGormCallbacks(db *gorm.DB) {
	// Do nothing
}

func (n nullHandler) ConfigureHTTPClient(client *http.Client, opts ...otelhttp.Option) {
	// Do nothing
}

func (n nullHandler) Gin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (n nullHandler) Start(_ context.Context) error {
	return nil
}
func (n nullHandler) Metrics() Meter {
	return n.meter
}

// NewNullHandler creates a new null transaction handler.
func NewNullHandler() Handler {
	return &nullHandler{
		tracer:     trace.NewNoopTracerProvider().Tracer(""),
		propagator: nullPropogator{},
		meter:      &NullMeterImpl{},
	}
}

var _ Handler = &nullHandler{}

// nullPropogator is a metrics propagator that does nothing.
type nullPropogator struct{}

func (n nullPropogator) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
}

func (n nullPropogator) Extract(ctx context.Context, _ propagation.TextMapCarrier) context.Context {
	return ctx
}

func (n nullPropogator) Fields() []string {
	return []string{}
}

var _ propagation.TextMapPropagator = &nullPropogator{}

// noopHandler is an http handler that does nothing.
type noopHandler struct{}

func (n noopHandler) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	// Do nothing
}

var _ http.Handler = noopHandler{}
