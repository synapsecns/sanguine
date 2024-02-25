package metrics

import (
	otelpyroscope "github.com/grafana/otel-profiling-go"
	pyroscope "github.com/grafana/pyroscope-go"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"go.opentelemetry.io/otel/trace"
	"os"
)

// StartPyroscope starts the pyroscope profiler.
// this will not run if the pyroscope endpoint is not set.
func StartPyroscope(info config.BuildInfo) *pyroscope.Profiler {
	if !core.HasEnv(internal.PyroscopeEndpoint) {
		return nil
	}

	// note: because pyroscope is a profiler, we cannot use buildinfo for the application name.
	pf, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: info.Name(),

		// replace this with the address of pyroscope server
		ServerAddress: os.Getenv(internal.PyroscopeEndpoint),

		// you can disable logging by setting this to nil
		Logger: pyroscopeLogger{},
		// Uncomment this line if you're having issues:
		//Logger:     pyroscope.StandardLogger,

		// you can provide static tags via a map:
		Tags: map[string]string{
			"hostname": os.Getenv("HOSTNAME"),
			"version":  info.Version(),
			"commit":   info.Commit(),
		},

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
	})
	if err != nil {
		logger.Warn(err)
	}
	return pf
}

type pyroscopeLogger struct{}

func (p pyroscopeLogger) Infof(_ string, _ ...interface{}) {
	// do nothing
}

func (p pyroscopeLogger) Debugf(_ string, _ ...interface{}) {
	// do nothing
}

func (p pyroscopeLogger) Errorf(str string, args ...interface{}) {
	logger.Warnf(str, args...)
}

var _ pyroscope.Logger = &pyroscopeLogger{}

// PyroscopeWrapTracerProvider wraps the tracer provider with pyroscope.
// The traceProvider is not affected if the pyroscope endpoint is not set.
func PyroscopeWrapTracerProvider(provider trace.TracerProvider, buildInfo config.BuildInfo, extraOpts ...otelpyroscope.Option) trace.TracerProvider {
	if !core.HasEnv(pyroscopeEndpoint) {
		return provider
	}

	opts := append([]otelpyroscope.Option{
		otelpyroscope.WithAppName(buildInfo.Name()),
		otelpyroscope.WithPyroscopeURL(os.Getenv(pyroscopeEndpoint)),
		otelpyroscope.WithRootSpanOnly(true),
		// we only need this sort of linkinf for the jaeger ui.
		otelpyroscope.WithProfileURL(core.HasEnv(internal.PyroscopeJaegerUIEnabled)),
		otelpyroscope.WithProfileBaselineURL(core.HasEnv(internal.PyroscopeJaegerUIEnabled)),
	}, extraOpts...)

	return otelpyroscope.NewTracerProvider(provider, opts...)
}
