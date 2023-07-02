package metrics

import (
	"context"
	"fmt"
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/synapsecns/sanguine/core/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"os"
)

type newRelicHandler struct {
	*baseHandler
	app       *newrelic.Application
	buildInfo config.BuildInfo
}

// see: https://docs.newrelic.com/docs/more-integrations/open-source-telemetry-integrations/opentelemetry/get-started/opentelemetry-set-up-your-app/
const (
	newRelicOTLPEndpoint = "otlp.nr-data.net"
	newrelicOtlpHeader   = "api-key"
	newRelicLicenseEnv   = "NEW_RELIC_LICENSE_KEY"
)

// NewRelicMetricsHandler creates a new newrelic metrics handler.
func NewRelicMetricsHandler(buildInfo config.BuildInfo) Handler {
	return &newRelicHandler{
		buildInfo:   buildInfo,
		baseHandler: newBaseHandler(buildInfo),
	}
}

func (n *newRelicHandler) Start(ctx context.Context) (err error) {
	// TODO: consider grpc
	client := otlptracehttp.NewClient(otlptracehttp.WithEndpoint(newRelicOTLPEndpoint), otlptracehttp.WithHeaders(map[string]string{
		newrelicOtlpHeader: os.Getenv(newRelicLicenseEnv),
	}))
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to create new relic otlp exporter: %w", err)
	}
	// TODO: allow customizable sampling
	n.baseHandler = newBaseHandler(n.buildInfo, tracesdk.WithBatcher(exporter), tracesdk.WithSampler(tracesdk.AlwaysSample()))

	if n.app == nil {
		n.app, err = newrelic.NewApplication(
			newrelic.ConfigAppName(n.buildInfo.Name()),
			newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
			newrelic.ConfigAppLogForwardingEnabled(true),
			newrelic.ConfigAppLogEnabled(true),
			newrelic.ConfigCodeLevelMetricsEnabled(true),
			nrzap.ConfigLogger(logger.Desugar()),
			func(c *newrelic.Config) {
				c.Labels = map[string]string{
					"version": n.buildInfo.Version(),
					"commit":  n.buildInfo.Commit(),
				}
			},
			// optional overrides
			newrelic.ConfigFromEnvironment(),
		)
		if err != nil {
			return fmt.Errorf("could not create new relic application: %w", err)
		}
	}

	return nil
}

func (n *newRelicHandler) Type() HandlerType {
	return NewRelic
}
