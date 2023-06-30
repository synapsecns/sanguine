package metrics

import (
	"cloud.google.com/go/compute/metadata"
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
	"google.golang.org/api/gkehub/v1"
	"net/http"
	"os"
)

type googleHandler struct {
	*baseHandler
	buildInfo config.BuildInfo
}

const googleProjectEnv = "GOOGLE_CLOUD_PROJECT"

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

func (h *googleHandler) Start(ctx context.Context) (err error) {
	projectID := os.Getenv(googleProjectEnv)
	// if project id can't be detected, try fetching it from the compute metadata server
	if projectID == "" {
		metadataClient := metadata.NewClient(http.DefaultClient)
		projectID, err = metadataClient.ProjectID()
		if err != nil {
			return fmt.Errorf("could not get project ID from metadata server: %w. If you cannot get this to work, please set %s", err, googleProjectEnv)
		}
		fmt.Println(projectID)
		svc, err := gkehub.NewService(ctx)
		if err != nil {
			return fmt.Errorf("could not create gkehub service: %w", err)
		}

		_, err = svc.Projects.Locations.Memberships.List("test").Do()
		if err != nil {
			return fmt.Errorf("could not list memberships: %w", err)
		}

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
