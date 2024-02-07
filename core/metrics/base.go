package metrics

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"net/http"
)

const pyroscopeEndpoint = internal.PyroscopeEndpoint

// baseHandler is a base metrics handler that implements the Handler interface.
// this is used to reduce the amount of boilerplate code needed to implement opentracing methods.
type baseHandler struct {
	resource   *resource.Resource
	tp         trace.TracerProvider
	tracer     trace.Tracer
	name       string
	propagator propagation.TextMapPropagator
	meter      MeterProvider
	// handler is an integrated handler for everything exported over http. This includes prometheus
	// or http-based sampling methods for other providers.
	handler http.Handler
}

func (b *baseHandler) Handler() http.Handler {
	return b.handler
}

func (b *baseHandler) Meter(name string, options ...metric.MeterOption) metric.Meter {
	return b.meter.Meter(name, options...)
}

func (b *baseHandler) Start(ctx context.Context) error {
	reader, err := prometheus.New()
	if err != nil {
		logger.Warnf("could not initialize prometheus exporter: %v, using no-op provider", err)
	}

	// TODO: allow this to be customizable, separate from the tracer provider.
	// in a way that's still usable.
	b.meter = sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(b.resource),
		// TODO: figure out how to provide a sample interval here
		// see: https://github.com/open-telemetry/opentelemetry-go/issues/3244
		// for more on the trade-off space
		sdkmetric.WithReader(reader),
	)
	otel.SetMeterProvider(b.meter)
	b.handler = promhttp.Handler()

	go func() {
		<-ctx.Done()
		// shutting down this way will not flush.
		_ = b.meter.Shutdown(ctx)
	}()

	go func() {
		b.startMetricsServer(ctx)
	}()

	return nil
}

const (
	// MetricsPortEnabledEnv is the environment variable that controls whether the metrics server is enabled.
	MetricsPortEnabledEnv = "METRICS_PORT_ENABLED"
	metricsPortEnv        = "METRICS_PORT"
	// MetricsPath is the environment variable that controls the path for the metrics server.
	MetricsPath        = "METRICS_PATH"
	metricsPortDefault = 8080
	// MetricsPathDefault is the default path for the metrics server.
	MetricsPathDefault = "/metrics"
)

// startMetricsServer starts the metrics server on the given port.
// this should be run in a separate goroutine.
func (b *baseHandler) startMetricsServer(ctx context.Context) {
	if !core.GetEnvBool(MetricsPortEnabledEnv, true) {
		return
	}

	port := core.GetEnvInt(metricsPortEnv, metricsPortDefault)
	path := core.GetEnv(MetricsPath, MetricsPathDefault)

	logger.Infof("starting metrics server on port %d at path %s", port, path)

	// create the metrics server
	server := ginhelper.New(logger)
	// note: this is a global setter, so it will affect all gin servers.
	// this is probably not wise, but a better workaround is required.
	gin.SetMode(gin.ReleaseMode)
	server.Use(b.Gin())
	server.GET(path, gin.WrapH(b.handler))

	connection := baseServer.Server{}
	err := connection.ListenAndServe(ctx, fmt.Sprintf(":%d", port), server)
	if err != nil {
		logger.Warnf("running metrics server failed: %v", err)
	}
}

func (b *baseHandler) Gin() gin.HandlerFunc {
	return otelgin.Middleware(b.name, otelgin.WithTracerProvider(b.tp), otelgin.WithPropagators(b.propagator))
}

func (b *baseHandler) Propagator() propagation.TextMapPropagator {
	return b.propagator
}

func (b *baseHandler) ConfigureHTTPClient(client *http.Client, opts ...otelhttp.Option) {
	opts = append([]otelhttp.Option{otelhttp.WithTracerProvider(b.tp), otelhttp.WithPropagators(b.propagator)}, opts...)
	client.Transport = otelhttp.NewTransport(client.Transport, opts...)
}

func (b *baseHandler) AddGormCallbacks(db *gorm.DB) {
	err := db.Use(otelgorm.NewPlugin(otelgorm.WithTracerProvider(b.tp)))
	if err != nil {
		logger.Warn("could not add gorm callbacks", "error", err)
	}
}

func (b *baseHandler) GetTracerProvider() trace.TracerProvider {
	return b.tp
}

// Tracer returns the tracer provider.
func (b *baseHandler) Tracer() trace.Tracer {
	return b.tracer
}

func (b *baseHandler) Type() HandlerType {
	panic("must be overridden by children")
}

func (b *baseHandler) Metrics() Meter {
	return NewOtelMeter(b.meter)
}

func makeResource(buildInfo config.BuildInfo) (*resource.Resource, error) {
	rsr, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(buildInfo.Name()),
			attribute.String("ENVIRONMENT", "default"),
			semconv.ServiceVersion(buildInfo.Version()),
			attribute.String("commit", buildInfo.Commit()),
			attribute.String("library.language", "go"),
		))
	// TODO: handle error or report
	if err != nil {
		return nil, fmt.Errorf("could not merge resources: %w", err)
	}
	return rsr, nil
}

// newBaseHandler creates a new baseHandler for otel.
// this is exported for testing.
func newBaseHandler(buildInfo config.BuildInfo, extraOpts ...tracesdk.TracerProviderOption) *baseHandler {
	// Ensure default SDK resources and the required service name are set.
	rsr, err := makeResource(buildInfo)
	// TODO: handle error or report
	if err != nil {
		logger.Warn("could not merge resources", "error", err)
	}

	opts := append([]tracesdk.TracerProviderOption{tracesdk.WithResource(rsr)}, extraOpts...)

	// TODO: add a way for users to pass in extra pyroscope options
	tp := PyroscopeWrapTracerProvider(tracesdk.NewTracerProvider(opts...), buildInfo)
	// will do nothing if not enabled.
	StartPyroscope(buildInfo)

	propagator := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader | b3.B3SingleHeader))
	return newBaseHandlerWithTracerProvider(rsr, buildInfo, tp, propagator)
}

// newBaseHandlerWithTracerProvider creates a new baseHandler for any opentelemtry tracer.
func newBaseHandlerWithTracerProvider(rsr *resource.Resource, buildInfo config.BuildInfo, tracerProvider trace.TracerProvider, propagator propagation.TextMapPropagator) *baseHandler {
	// default tracer for server.
	otel.SetTracerProvider(tracerProvider)
	tracer := tracerProvider.Tracer(buildInfo.Name())
	otel.SetTextMapPropagator(propagator)

	// note: meter purposely is not registered until startup.
	return &baseHandler{
		resource:   rsr,
		tp:         tracerProvider,
		tracer:     tracer,
		name:       buildInfo.Name(),
		propagator: propagator,
		handler:    promhttp.Handler(),
	}
}

var _ Handler = &baseHandler{}
