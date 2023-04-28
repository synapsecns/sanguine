package metrics

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	"net/http"
)

// baseHandler is a base metrics handler that implements the Handler interface.
// this is used to reduce the amount of boilerplate code needed to implement opentracing methods.
type baseHandler struct {
	tp         trace.TracerProvider
	tracer     trace.Tracer
	name       string
	propagator propagation.TextMapPropagator
}

func (b *baseHandler) Start(ctx context.Context) error {
	// do nothing
	return nil
}

func (b *baseHandler) Gin() gin.HandlerFunc {
	return otelgin.Middleware(b.name, otelgin.WithTracerProvider(b.tp))
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

// newBaseHandler creates a new baseHandler for otel.
// this is exported for testing.
func newBaseHandler(buildInfo config.BuildInfo, extraOpts ...tracesdk.TracerProviderOption) *baseHandler {
	// Ensure default SDK resources and the required service name are set.
	rsr, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(buildInfo.Name()),
			attribute.String("ENVIRONMENT", "default"),
			semconv.ServiceVersion(buildInfo.Version()),
			attribute.String("commit", buildInfo.Commit()),
		))
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
	return newBaseHandlerWithTracerProvider(buildInfo, tp, propagator)
}

// newBaseHandlerWithTracerProvider creates a new baseHandler for any opentelemtry tracer.
func newBaseHandlerWithTracerProvider(buildInfo config.BuildInfo, tracerProvider trace.TracerProvider, propagator propagation.TextMapPropagator) *baseHandler {
	// default tracer for server
	otel.SetTracerProvider(tracerProvider)
	tracer := tracerProvider.Tracer(buildInfo.Name())
	otel.SetTextMapPropagator(propagator)

	return &baseHandler{
		tp:         tracerProvider,
		tracer:     tracer,
		name:       buildInfo.Name(),
		propagator: propagator,
	}
}

var _ Handler = &baseHandler{}
