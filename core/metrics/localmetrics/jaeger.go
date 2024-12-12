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

// purgeAllResources cleans up all Docker resources associated with Jaeger
func (j *testJaeger) purgeAllResources() error {
	j.tb.Log("Cleaning up all Docker resources...")

	// Clean up Jaeger containers
	if err := cleanupExistingContainers(j.pool, "jaeger"); err != nil {
		return fmt.Errorf("failed to clean up Jaeger containers: %w", err)
	}

	// Clean up UI containers
	if err := cleanupExistingContainers(j.pool, "jaeger-ui"); err != nil {
		return fmt.Errorf("failed to clean up UI containers: %w", err)
	}

	return nil
}

// StartJaegerServer starts a new jaeger instance.
// nolint: cyclop
func (j *testJaeger) StartJaegerServer(ctx context.Context) *uiResource {
	j.tb.Log("Starting Jaeger server...")

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

	// Clean up any existing resources before starting
	if err := j.purgeAllResources(); err != nil {
		j.tb.Logf("Warning: Failed to clean up resources: %v", err)
	}

	j.tb.Log("Setting up container options...")
	runOptions := &dockertest.RunOptions{
		Repository:   "jaegertracing/all-in-one",
		Tag:          "latest",
		Hostname:     "jaeger",
		ExposedPorts: []string{"14268", "16686"},
		Env: []string{
			"COLLECTOR_OTLP_ENABLED=true",
			"LOG_LEVEL=debug",
			"COLLECTOR_HTTP_PORT=0",
			"COLLECTOR_OTLP_HTTP_PORT=0",
			"QUERY_HTTP_PORT=0",
		},
		Networks: j.getNetworks(),
		Labels: map[string]string{
			appLabel:   "jaeger",
			runIDLabel: j.runID,
		},
	}

	var resource *dockertest.Resource
	var err error

	j.tb.Log("Starting container with retry logic...")
	// Create container with improved retry logic and shorter timeouts
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err = retry.WithBackoff(ctx, func(ctx context.Context) error {
		j.tb.Log("Attempting to start container...")

		// Create a context with timeout for this attempt
		attemptCtx, attemptCancel := context.WithTimeout(ctx, 5*time.Second)
		defer attemptCancel()

		// Wait for network stability (shorter wait)
		time.Sleep(time.Second)

		j.tb.Log("Running container with options...")
		resource, err = j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
			j.tb.Log("Configuring container host settings...")
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			config.PublishAllPorts = true
			config.PortBindings = make(map[docker.Port][]docker.PortBinding)
			config.NetworkMode = "bridge"
		})
		if err != nil {
			j.tb.Logf("Failed to start container: %v", err)
			return fmt.Errorf("container start failed: %w", err)
		}

		j.tb.Log("Container started, validating ports...")
		// Validate ports with timeout
		portCtx, portCancel := context.WithTimeout(attemptCtx, 2*time.Second)
		defer portCancel()

		portChan := make(chan error, 1)
		go func() {
			tracePort := dockerutil.GetPort(resource, "14268/tcp")
			uiPort := dockerutil.GetPort(resource, "16686/tcp")
			if tracePort == "" || uiPort == "" {
				portChan <- fmt.Errorf("failed to get container ports")
				return
			}

			j.tb.Logf("Setting up endpoints with ports - trace: %s, ui: %s", tracePort, uiPort)
			// Set environment variables
			endpoint := fmt.Sprintf("http://localhost:%s", tracePort)
			uiEndpoint := fmt.Sprintf("http://localhost:%s", uiPort)

			if err := os.Setenv(internal.JaegerEndpoint, fmt.Sprintf("%s/api/traces", endpoint)); err != nil {
				portChan <- fmt.Errorf("failed to set endpoint: %w", err)
				return
			}
			if err := os.Setenv(internal.JaegerUIEndpoint, uiEndpoint); err != nil {
				portChan <- fmt.Errorf("failed to set UI endpoint: %w", err)
				return
			}
			portChan <- nil
		}()

		select {
		case err := <-portChan:
			if err != nil {
				j.tb.Logf("Port validation failed: %v", err)
				return err
			}
		case <-portCtx.Done():
			return fmt.Errorf("port validation timed out")
		}

		j.tb.Log("Waiting for endpoints to be ready...")
		// Wait for endpoints with shorter timeout
		return retry.WithBackoff(ctx, func(ctx context.Context) error {
			endpoint := os.Getenv(internal.JaegerEndpoint)
			uiEndpoint := os.Getenv(internal.JaegerUIEndpoint)

			if err := checkURL(endpoint)(ctx); err != nil {
				j.tb.Logf("Endpoint not ready: %v", err)
				return err
			}
			if err := checkURL(uiEndpoint)(ctx); err != nil {
				j.tb.Logf("UI endpoint not ready: %v", err)
				return err
			}
			j.tb.Log("Endpoints are ready")
			return nil
		},
			retry.WithMax(time.Second*2),
			retry.WithMaxAttempts(5))
	},
		retry.WithMax(time.Second*3),
		retry.WithMaxAttempts(5))

	if err != nil {
		j.tb.Logf("Failed to start container after retries: %v", err)
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

// StartJaegerPyroscopeUI starts a new jaeger pyroscope ui instance.
func (j *testJaeger) StartJaegerPyroscopeUI(ctx context.Context) *uiResource {
	// Handle environment variables
	if core.HasEnv(internal.JaegerUIEndpoint) || !j.cfg.enablePyroscope {
		return &uiResource{
			Resource: nil,
			uiURL:    os.Getenv(internal.JaegerUIEndpoint),
		}
	}

	// Clean up any existing containers before starting
	if err := cleanupExistingContainers(j.pool, "jaeger-ui"); err != nil {
		j.tb.Logf("Warning: Failed to clean up existing containers: %v", err)
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
			"HTTP_PORT=0",
		},
		Networks: j.getNetworks(),
		Labels: map[string]string{
			appLabel:   "jaeger-ui",
			runIDLabel: j.runID,
		},
	}

	var resource *dockertest.Resource
	var err error

	// Create container with retry logic and shorter timeouts
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err = retry.WithBackoff(ctx, func(ctx context.Context) error {
		j.tb.Log("Attempting to start Jaeger Pyroscope UI container...")

		// Create a context with timeout for this attempt
		attemptCtx, attemptCancel := context.WithTimeout(ctx, 5*time.Second)
		defer attemptCancel()

		// Wait for network stability (shorter wait)
		time.Sleep(time.Second)

		j.tb.Log("Running container with options...")
		resource, err = j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
			j.tb.Log("Configuring container host settings...")
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			config.PublishAllPorts = true
			config.PortBindings = make(map[docker.Port][]docker.PortBinding)
			config.NetworkMode = "bridge"
		})
		if err != nil {
			j.tb.Logf("Failed to start Jaeger Pyroscope UI container: %v", err)
			return fmt.Errorf("container start failed: %w", err)
		}

		j.tb.Log("Container started, validating port...")
		// Validate port with timeout
		portCtx, portCancel := context.WithTimeout(attemptCtx, 2*time.Second)
		defer portCancel()

		portChan := make(chan error, 1)
		go func() {
			uiPort := dockerutil.GetPort(resource, "80/tcp")
			if uiPort == "" {
				portChan <- fmt.Errorf("failed to get Jaeger UI port")
				return
			}

			// Set environment variable for endpoint
			uiEndpoint := fmt.Sprintf("http://localhost:%s", uiPort)
			if err := os.Setenv(internal.JaegerUIEndpoint, uiEndpoint); err != nil {
				portChan <- fmt.Errorf("failed to set UI endpoint: %w", err)
				return
			}
			portChan <- nil
		}()

		select {
		case err := <-portChan:
			if err != nil {
				j.tb.Logf("Port validation failed: %v", err)
				return err
			}
		case <-portCtx.Done():
			return fmt.Errorf("port validation timed out")
		}

		j.tb.Log("Waiting for UI endpoint to be ready...")
		// Wait for UI endpoint with shorter timeout
		return retry.WithBackoff(ctx, checkURL(os.Getenv(internal.JaegerUIEndpoint)),
			retry.WithMax(time.Second*2),
			retry.WithMaxAttempts(5))
	},
		retry.WithMax(time.Second*3),
		retry.WithMaxAttempts(5))

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
