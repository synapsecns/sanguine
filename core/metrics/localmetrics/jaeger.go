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
	"github.com/synapsecns/sanguine/core/retry"
)

// StartJaegerServer starts a new jaeger instance.
// nolint: cyclop
func (j *testJaeger) StartJaegerServer(ctx context.Context) *uiResource {
	// Handle environment variables
	if core.HasEnv(internal.JaegerEndpoint) {
		// If JaegerEndpoint is set but JaegerUIEndpoint is empty, fail
		if !core.HasEnv(internal.JaegerUIEndpoint) {
			if j.tb != nil {
				j.tb.Error("JaegerUIEndpoint must be set when JaegerEndpoint is set")
				j.tb.Fail()
			}
			return nil
		}
		return &uiResource{
			Resource: nil,
			uiURL:    os.Getenv(internal.JaegerUIEndpoint),
		}
	}

	runOptions := &dockertest.RunOptions{
		Repository:   "jaegertracing/all-in-one",
		Tag:          "latest",
		Hostname:     "jaeger",
		ExposedPorts: []string{"14268", "16686"},
		Env: []string{
			"COLLECTOR_OTLP_ENABLED=true",
			"LOG_LEVEL=debug",
		},
		Networks: j.getNetworks(),
		Labels: map[string]string{
			appLabel:   "jaeger",
			runIDLabel: j.runID,
		},
	}

	var resource *dockertest.Resource
	var err error

	// Create container with improved retry logic
	err = retry.WithBackoff(ctx, func(ctx context.Context) error {
		// Wait for network stability
		time.Sleep(time.Second * 2)

		resource, err = j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			config.PublishAllPorts = true
			config.PortBindings = make(map[docker.Port][]docker.PortBinding)
		})
		if err != nil {
			j.tb.Logf("Failed to start Jaeger container: %v", err)
			return err
		}

		// Validate ports
		tracePort := dockerutil.GetPort(resource, "14268/tcp")
		uiPort := dockerutil.GetPort(resource, "16686/tcp")
		if tracePort == "" || uiPort == "" {
			return fmt.Errorf("failed to get Jaeger ports")
		}

		// Set environment variables using os.Setenv instead of tb.Setenv
		endpoint := fmt.Sprintf("http://localhost:%s", tracePort)
		uiEndpoint := fmt.Sprintf("http://localhost:%s", uiPort)
		if err := os.Setenv(internal.JaegerEndpoint, fmt.Sprintf("%s/api/traces", endpoint)); err != nil {
			return fmt.Errorf("failed to set Jaeger endpoint: %v", err)
		}
		if err := os.Setenv(internal.JaegerUIEndpoint, uiEndpoint); err != nil {
			return fmt.Errorf("failed to set Jaeger UI endpoint: %v", err)
		}

		// Wait for endpoints with increased timeout
		return retry.WithBackoff(ctx, func(ctx context.Context) error {
			if err := checkURL(endpoint)(ctx); err != nil {
				return fmt.Errorf("jaeger endpoint not ready: %v", err)
			}
			if err := checkURL(uiEndpoint)(ctx); err != nil {
				return fmt.Errorf("jaeger UI endpoint not ready: %v", err)
			}
			return nil
		},
			retry.WithMax(time.Second*5),
			retry.WithMaxAttempts(15))
	},
		retry.WithMax(time.Second*5),
		retry.WithMaxAttempts(3))

	if err != nil {
		j.tb.Logf("Failed to start Jaeger container after retries: %v", err)
		return nil
	}

	if !j.cfg.keepContainers {
		if err = resource.Expire(uint(keepAliveOnFailure.Seconds())); err != nil {
			j.tb.Logf("Failed to set container expiry: %v", err)
		}
	}

	// PLACEHOLDER: log resource handling and return

	return &uiResource{
		Resource: resource,
		uiURL:    os.Getenv(internal.JaegerUIEndpoint),
	}
}

// StartJaegerPyroscopeUI starts a new jaeger pyroscope ui instance.
func (j *testJaeger) StartJaegerPyroscopeUI(ctx context.Context) *uiResource {
	// Handle environment variables
	if core.HasEnv(internal.JaegerUIEndpoint) || !j.cfg.enablePyroscope {
		return &uiResource{
			Resource: nil,
			uiURL:    os.Getenv(internal.JaegerUIEndpoint),
		}
	}

	// Set required environment variables
	if err := os.Setenv(internal.PyroscopeJaegerUIEnabled, "true"); err != nil {
		j.tb.Logf("Failed to enable Pyroscope Jaeger UI: %v", err)
		return nil
	}

	runOptions := &dockertest.RunOptions{
		Repository:   "ghcr.io/synapsecns/jaeger-ui-pyroscope",
		Tag:          "latest",
		ExposedPorts: []string{"80"},
		Env: []string{
			"LOG_LEVEL=debug",
		},
		Networks: j.getNetworks(),
		Labels: map[string]string{
			appLabel:   "jaeger-ui",
			runIDLabel: j.runID,
		},
	}

	var resource *dockertest.Resource
	var err error

	// Create container with retry logic
	err = retry.WithBackoff(ctx, func(ctx context.Context) error {
		// Wait for network stability
		time.Sleep(time.Second)

		resource, err = j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			config.PublishAllPorts = true
			config.PortBindings = make(map[docker.Port][]docker.PortBinding)
		})
		if err != nil {
			j.tb.Logf("Failed to start Jaeger Pyroscope UI container: %v", err)
			return err
		}

		// Validate port
		uiPort := dockerutil.GetPort(resource, "80/tcp")
		if uiPort == "" {
			return fmt.Errorf("failed to get Jaeger UI port")
		}

		// Set environment variable for endpoint
		uiEndpoint := fmt.Sprintf("http://localhost:%s", uiPort)
		j.tb.Setenv(internal.JaegerUIEndpoint, uiEndpoint)

		// Wait for UI endpoint with increased timeout
		return retry.WithBackoff(ctx, checkURL(uiEndpoint),
			retry.WithMax(time.Second*5),
			retry.WithMaxAttempts(15))
	},
		retry.WithMax(time.Second*5),
		retry.WithMaxAttempts(3))

	if err != nil {
		j.tb.Logf("Failed to start Jaeger Pyroscope UI container after retries: %v", err)
		return nil
	}

	if !j.cfg.keepContainers {
		if err = resource.Expire(uint(keepAliveOnFailure.Seconds())); err != nil {
			j.tb.Logf("Failed to set container expiry: %v", err)
		}
	}

	return &uiResource{
		Resource: resource,
		uiURL:    os.Getenv(internal.JaegerUIEndpoint),
	}
}
