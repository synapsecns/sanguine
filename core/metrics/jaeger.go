package metrics

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"go.opentelemetry.io/otel/exporters/jaeger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"os"
)

type jaegerHandler struct {
	*baseHandler
	buildInfo config.BuildInfo
	exporter  *jaeger.Exporter
}

// NewJaegerHandler creates a new jaeger handler for handling jaeger traces.
// the JAEGER_ENDPOINT environment variable must be set for this to work.
// Note: currently, this is only suitable for local runs, because of default options we've put in place
// This can be fixed in a future version through an option builder.
// TODO: this should be replaced w/ the otlp exporter.
func NewJaegerHandler(buildInfo config.BuildInfo) Handler {
	return &jaegerHandler{
		buildInfo: buildInfo,
	}
}

const (
	jaegerEnv = internal.JaegerEndpoint
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

func (j *jaegerHandler) Type() HandlerType {
	return Jaeger
}
