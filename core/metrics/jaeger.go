package metrics

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/metrics/localserver"
	"go.opentelemetry.io/otel/exporters/jaeger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"os"
	"testing"
)

type jaegerHandler struct {
	*baseHandler
	buildInfo config.BuildInfo
	exporter  *jaeger.Exporter
}

// NewJaegerHandler creates a new jaeger handler for handling jaeger traces.
// the JAEGER_ENDPOINT environment variable must be set for this to work.
// Note: currently, this is only suitable for local runs, because of default options we've put in place
// This can be fixed in a future version through an option builder
func NewJaegerHandler(buildInfo config.BuildInfo) Handler {
	return &jaegerHandler{
		buildInfo: buildInfo,
	}
}

const (
	jaegerEnv         = internal.JAEGER_ENDPOINT
	pyroscopeEndpoint = internal.PYROSCOPE_ENDPOINT
)

func (j *jaegerHandler) Start(ctx context.Context) (err error) {
	endpoint := os.Getenv(jaegerEnv)
	if endpoint == "" {
		return fmt.Errorf("could not get jaeger endpoint from env")
	}
	j.exporter, err = jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)),
	)
	if err != nil {
		return fmt.Errorf("could not create jaeger exporter: %w", err)
	}

	j.baseHandler = newBaseHandler(j.buildInfo, tracesdk.WithSyncer(j.exporter), tracesdk.WithSampler(tracesdk.AlwaysSample()))
	err = j.baseHandler.Start(ctx)
	if err != nil {
		return fmt.Errorf("could not start base handler: %w", err)
	}

	return nil
}

// SetupTestJaeger creates a new test jaeger instance. If the test fails, the instance is kept alive for 5 minutes.
// we also allow a GLOBAL_jaeger env var to be set to a jaeger url to send all traces to in order to avoid having to boot for long running tests.
func SetupTestJaeger(ctx context.Context, tb testing.TB) {
	localserver.StartServer(ctx, tb)
}
