package metrics

import (
	"context"
	"fmt"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/synapsecns/sanguine/core/config"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"os"
)

type googleHandler struct {
	*baseHandler
	buildInfo config.BuildInfo
}

const googleProjectEnv = "GOOGLE_CLOUD_PROJECT"

var errMissingProjectID = fmt.Errorf("missing project ID, please set %s", googleProjectEnv)

func NewGoogleMetricsHandler(buildInfo config.BuildInfo) Handler {
	handler := googleHandler{
		buildInfo: buildInfo,
	}

	// this is a no-op to prevent panics, it gets replaced in start.
	handler.baseHandler = newBaseHandler(buildInfo)

	return &handler
}

func Type() HandlerType {
	return Google
}

func (h *googleHandler) Start(ctx context.Context) error {
	projectID := os.Getenv(googleProjectEnv)
	if projectID == "" {
		return errMissingProjectID
	}

	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err != nil {
		return fmt.Errorf("exporter could not be created: %w", err)
	}

	res, err := resource.New(ctx,
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(h.buildInfo.Name()),
			semconv.ServiceVersionKey.String(h.buildInfo.Version()),
			attribute.String("commit", h.buildInfo.Commit()),
		),
	)

	if err != nil {
		return fmt.Errorf("resource could not be created: %w", err)
	}

	// Create trace provider with the exporter.
	//
	// By default it uses AlwaysSample() which samples all traces.
	// TODO: Setup probabilistic sampling.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// TODO: profiler
	// TOOD: metrics
	// TODO: more stuff in the readme
	propagator := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader | b3.B3SingleHeader))

	h.baseHandler = newBaseHandlerWithTracerProvider(h.buildInfo, tp, propagator)
	return nil
}
