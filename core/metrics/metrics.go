package metrics

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/config"
	experimentalLogger "github.com/synapsecns/sanguine/core/metrics/logger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
)

// Handler collects metrics.
type Handler interface {
	Start(ctx context.Context) error
	// Gin gets a gin middleware for tracing.
	Gin() gin.HandlerFunc
	// ConfigureHTTPClient configures tracing on an http client
	ConfigureHTTPClient(client *http.Client, opts ...otelhttp.Option)
	// AddGormCallbacks adds gorm callbacks for tracing.
	AddGormCallbacks(db *gorm.DB)
	// GetTracerProvider returns the tracer provider.
	GetTracerProvider() trace.TracerProvider
	// Tracer returns the tracer provider.
	Tracer() trace.Tracer
	// Propagator returns the propagator.
	Propagator() propagation.TextMapPropagator
	// Type returns the handler type.
	Type() HandlerType
	// Metrics returns a metric provider
	// Deprecated: Will be removed in a future version please use meter.
	Metrics() Meter
	// Meter returns a metric provider
	Meter(name string, options ...metric.MeterOption) metric.Meter
	// Handler returns the http handler for the metrics endpoint.
	// right now, this supports only a single route
	Handler() http.Handler
	// ExperimentalLogger returns an experimental logger.
	ExperimentalLogger() experimentalLogger.ExperimentalLogger
}

// HandlerType is the handler type to use
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=HandlerType -linecomment
type HandlerType uint8

// AllHandlerTypes is a list of all contract types. Since we use stringer and this is a testing library, instead
// of manually copying all these out we pull the names out of stringer. In order to make sure stringer is updated, we panic on
// any method called where the index is higher than the stringer array length.
var AllHandlerTypes []HandlerType

func init() {
	for i := 0; i < len(_HandlerType_index); i++ {
		contractType := HandlerType(i)
		AllHandlerTypes = append(AllHandlerTypes, contractType)
	}
}

const (
	// OTLP is the otlp driver.
	OTLP HandlerType = iota + 1 // OTLP
	// Jaeger is the jaeger driver.
	Jaeger // Jaeger
	// Null is a null data type handler.
	Null // Null
)

// Lower gets the lowercase version of the handler type. Useful for comparison
// in switch.
func (i HandlerType) Lower() string {
	return strings.ToLower(i.String())
}

// HandlerEnv is the driver to use for metrics.
const HandlerEnv = "METRICS_HANDLER"

// NewFromEnv sets up a metrics handler from environment variable.
// this will not set the global and generally, SetupFromEnv should be used instead.
func NewFromEnv(ctx context.Context, buildInfo config.BuildInfo) (handler Handler, err error) {
	metricsHandler := strings.ToLower(os.Getenv(HandlerEnv))
	var ht HandlerType
	//nolint: gocritic
	switch metricsHandler {
	case OTLP.Lower():
		ht = OTLP
	case Jaeger.Lower():
		ht = Jaeger
	case Null.Lower():
		ht = Null
	default:
		ht = Null
	}

	return NewByType(ctx, buildInfo, ht)
}

// NewByType sets up a metrics handler by type.
func NewByType(ctx context.Context, buildInfo config.BuildInfo, ht HandlerType) (handler Handler, err error) {
	//nolint: gocritic
	switch ht {
	case OTLP:
		handler = NewOTLPMetricsHandler(buildInfo)
	case Jaeger:
		handler = NewJaegerHandler(buildInfo)
	case Null:
		handler = NewNullHandler()
	default:
		handler = NewNullHandler()
	}

	if handler != nil {
		err = handler.Start(ctx)
		if err != nil {
			return nil, fmt.Errorf("could not start handler: %w", err)
		}
	}

	return handler, nil
}
