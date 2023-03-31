package metrics

import (
	otelpyroscope "github.com/pyroscope-io/otel-profiling-go"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	"go.opentelemetry.io/otel/trace"
	"os"
)

// PyroscopeWrapTracerProvider wraps the tracer provider with pyroscope.
// The traceProvider is not affected if the pyroscope endpoint is not set.
func PyroscopeWrapTracerProvider(provider trace.TracerProvider, buildInfo config.BuildInfo, extraOpts ...otelpyroscope.Option) trace.TracerProvider {
	if !core.HasEnv(pyroscopeEndpoint) {
		return provider
	}

	opts := append([]otelpyroscope.Option{
		otelpyroscope.WithAppName(buildInfo.Name()),
		otelpyroscope.WithPyroscopeURL(os.Getenv(pyroscopeEndpoint)),
		otelpyroscope.WithRootSpanOnly(false),
		otelpyroscope.WithProfileURL(true),
		otelpyroscope.WithProfileBaselineURL(true),
	}, extraOpts...)

	return otelpyroscope.NewTracerProvider(provider, opts...)
}
