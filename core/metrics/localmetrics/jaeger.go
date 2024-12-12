package localmetrics

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dockerutil"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/synapsecns/sanguine/core/retry"
)

// StartJaegerServer starts a new jaeger instance.
// nolint: cyclop
func (j *testJaeger) StartJaegerServer(ctx context.Context) *uiResource {
	if core.HasEnv(internal.JaegerEndpoint) && !core.HasEnv(internal.JaegerUIEndpoint) {
		j.tb.Fatalf("%s is set but %s is not, please remove %s or set %s", internal.JaegerEndpoint, internal.JaegerUIEndpoint, internal.JaegerEndpoint, internal.JaegerUIEndpoint)
	}

	if core.HasEnv(internal.JaegerEndpoint) {
		return &uiResource{
			Resource: nil,
			uiURL:    os.Getenv(internal.JaegerUIEndpoint),
		}
	}

	runOptions := &dockertest.RunOptions{
		Repository:   "jaegertracing/all-in-one",
		Tag:          "latest",
		Hostname:     "jaeger",
		ExposedPorts: []string{"14268/tcp", "16686/tcp"},
		Networks:     j.getNetworks(),
		Labels: map[string]string{
			appLabel:   "jaeger",
			runIDLabel: j.runID,
		},
	}
	resource, err := j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		config.PublishAllPorts = true
	})
	if err != nil {
		j.tb.Logf("Failed to start Jaeger container: %v", err)
		return nil
	}

	var uiEndpoint string
	j.tb.Setenv(internal.JaegerEndpoint, fmt.Sprintf("http://localhost:%s/api/traces", dockerutil.GetPort(resource, "14268/tcp")))
	uiEndpoint = fmt.Sprintf("http://localhost:%s", dockerutil.GetPort(resource, "16686/tcp"))

	if !j.cfg.keepContainers {
		err = resource.Expire(uint(keepAliveOnFailure.Seconds()))
		if err != nil {
			j.tb.Logf("Failed to set container expiry: %v", err)
		}
	}

	logResourceChan := make(chan *uiResource, 1)

	go func() {
		_ = dockerutil.TailContainerLogs(dockerutil.WithContext(ctx), dockerutil.WithPool(j.pool), dockerutil.WithProcessLogOptions(processlog.WithLogDir(j.logDir), processlog.WithLogFileName("jaeger")), dockerutil.WithFollow(true),
			dockerutil.WithResource(resource), dockerutil.WithCallback(func(ctx context.Context, metadata processlog.LogMetadata) {
				select {
				case <-ctx.Done():
					return
				case logResourceChan <- &uiResource{
					Resource: resource,
					uiURL:    uiEndpoint,
				}:
					return
				}
			}))
	}()

	// Wait for Jaeger endpoint to be ready with more lenient retry parameters
	err = retry.WithBackoff(ctx, checkURL(os.Getenv(internal.JaegerEndpoint)),
		retry.WithMax(time.Second*5),    // Increase max retry interval
		retry.WithMaxAttempts(60),       // Increase max attempts
		retry.WithInitial(time.Second*1)) // Start with a longer initial delay
	if err != nil {
		j.tb.Logf("Failed to connect to Jaeger endpoint: %v", err)
		return nil
	}

	select {
	case <-ctx.Done():
		return nil
	case logResource := <-logResourceChan:
		if !j.cfg.enablePyroscopeJaeger && logResource != nil && logResource.uiURL != "" {
			err = os.Setenv(internal.JaegerUIEndpoint, logResource.uiURL)
			if err != nil {
				j.tb.Logf("Failed to set Jaeger UI endpoint: %v", err)
			}
		}

		return logResource
	}
}

// StartJaegerPyroscopeUI starts a new jaeger pyroscope ui instance.
func (j *testJaeger) StartJaegerPyroscopeUI(ctx context.Context) *uiResource {
	if core.HasEnv(internal.JaegerUIEndpoint) || !j.cfg.enablePyroscope {
		return &uiResource{
			uiURL: os.Getenv(internal.JaegerUIEndpoint),
		}
	}

	err := os.Setenv(internal.PyroscopeJaegerUIEnabled, "true")
	if err != nil {
		j.tb.Logf("Failed to enable Pyroscope Jaeger UI: %v", err)
		return nil
	}

	runOptions := &dockertest.RunOptions{
		Repository:   "ghcr.io/synapsecns/jaeger-ui-pyroscope",
		Tag:          "latest",
		ExposedPorts: []string{"80/tcp"},
		Networks:     j.getNetworks(),
		Labels: map[string]string{
			appLabel:   "jaeger-ui",
			runIDLabel: j.runID,
		},
	}
	resource, err := j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		config.PublishAllPorts = true
	})
	if err != nil {
		j.tb.Logf("Failed to start Jaeger Pyroscope UI container: %v", err)
		return nil
	}

	j.tb.Setenv(internal.JaegerUIEndpoint, fmt.Sprintf("http://localhost:%s", dockerutil.GetPort(resource, "80/tcp")))

	if !j.cfg.keepContainers {
		err = resource.Expire(uint(keepAliveOnFailure.Seconds()))
		if err != nil {
			j.tb.Logf("Failed to set container expiry: %v", err)
		}
	}

	logResourceChan := make(chan *uiResource, 1)

	go func() {
		_ = dockerutil.TailContainerLogs(dockerutil.WithContext(ctx), dockerutil.WithPool(j.pool), dockerutil.WithProcessLogOptions(processlog.WithLogDir(j.logDir), processlog.WithLogFileName("jaeger-pyroscope-ui")), dockerutil.WithFollow(true),
			dockerutil.WithResource(resource), dockerutil.WithCallback(func(ctx context.Context, metadata processlog.LogMetadata) {
				select {
				case <-ctx.Done():
					return
				case logResourceChan <- &uiResource{
					Resource: resource,
					uiURL:    os.Getenv(internal.JaegerUIEndpoint),
				}:
					return
				}
			}))
	}()

	err = retry.WithBackoff(ctx, checkURL(os.Getenv(internal.JaegerEndpoint)),
		retry.WithMax(time.Second*5),    // Increase max retry interval
		retry.WithMaxAttempts(60),       // Increase max attempts
		retry.WithInitial(time.Second*1)) // Start with a longer initial delay
	if err != nil {
		j.tb.Logf("Failed to connect to Jaeger endpoint: %v", err)
		return nil
	}

	select {
	case <-ctx.Done():
		return nil
	case logResource := <-logResourceChan:
		return logResource
	}
}
