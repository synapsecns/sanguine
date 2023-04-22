package pyroscope

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/pyroscope-io/client/pyroscope"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	"os"
	"runtime"
)

// Monitor starts the pyroscope client.
func Monitor(info config.BuildInfo) error {
	// disable in ci
	if core.IsTest() || os.Getenv("CI") != "" {
		return nil
	}

	runTag := fmt.Sprintf("%s-%s", gofakeit.Word(), gofakeit.Word())
	fmt.Printf("running with tag: %s \n", runTag)
	// These 2 lines are only required if you're using mutex or block profiling
	// Read the explanation below for how to set these rates:
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: info.Name(),

		// replace this with the address of pyroscope server
		ServerAddress: core.GetEnv("PYROSCOPE_SERVER_ADDRESS", "https://ingest.pyroscope.cloud"),

		// you can disable logging by setting this to nil
		//Logger: pyroscope.StandardLogger,

		// optionally, if authentication is enabled, specify the API key:
		// Note: since this will be deprecated in the future, we hardcode the secret here.
		AuthToken: core.GetEnv("PYROSCOPE_AUTH_TOKEN", "psx-g8D1gMTc3sU-3gLaNkyDd3FVOwuTGAPQtosP-IhtYJyQxB8z5v6vhpE"),

		// you can provide static tags via a map:
		Tags: map[string]string{
			"hostname": os.Getenv("HOSTNAME"),
			"runtag":   runTag,
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

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	if err != nil {
		return fmt.Errorf("could not profile app: %w", err)
	}

	return nil
}
