package localmetrics

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

// waitForContainerHealth waits for the container to be healthy
func (j *testJaeger) waitForContainerHealth(resource *dockertest.Resource) error {
	startTime := time.Now()
	j.tb.Log("Starting container health check...")

	// Use pool's retry mechanism with timeout and detailed logging
	if err := j.pool.Retry(func() error {
		// Check container status first
		container := resource.Container
		if container == nil {
			return fmt.Errorf("container reference is nil")
		}

		j.tb.Logf("Container status: %s (running for %v)", container.State.Status, time.Since(startTime))
		if container.State.Status != "running" {
			return fmt.Errorf("container is not running, status: %s", container.State.Status)
		}

		// Try to connect to all endpoints with increased timeouts
		collectorEndpoint := fmt.Sprintf("http://localhost:%s/api/traces", resource.GetPort("14268/tcp"))
		queryEndpoint := fmt.Sprintf("http://localhost:%s", resource.GetPort("16686/tcp"))
		healthEndpoint := fmt.Sprintf("http://localhost:%s", resource.GetPort("14269/tcp"))
		otlpGrpcEndpoint := fmt.Sprintf("http://localhost:%s", resource.GetPort("4317/tcp"))
		otlpHttpEndpoint := fmt.Sprintf("http://localhost:%s", resource.GetPort("4318/tcp"))

		j.tb.Logf("Checking endpoints - collector: %s, query: %s, health: %s, otlp-grpc: %s, otlp-http: %s",
			collectorEndpoint, queryEndpoint, healthEndpoint, otlpGrpcEndpoint, otlpHttpEndpoint)

		// Check health endpoint first
		healthReady := isEndpointReady(healthEndpoint)
		if !healthReady {
			j.tb.Log("Health endpoint not ready")
			return fmt.Errorf("health endpoint not ready (waited %v)", time.Since(startTime))
		}

		// Quick check all endpoints with detailed logging
		collectorReady := isEndpointReady(collectorEndpoint)
		if !collectorReady {
			j.tb.Log("Collector endpoint not ready")
		}

		queryReady := isEndpointReady(queryEndpoint)
		if !queryReady {
			j.tb.Log("Query endpoint not ready")
		}

		otlpGrpcReady := isEndpointReady(otlpGrpcEndpoint)
		if !otlpGrpcReady {
			j.tb.Log("OTLP gRPC endpoint not ready")
		}

		otlpHttpReady := isEndpointReady(otlpHttpEndpoint)
		if !otlpHttpReady {
			j.tb.Log("OTLP HTTP endpoint not ready")
		}

		if !collectorReady || !queryReady || !otlpGrpcReady || !otlpHttpReady {
			// Get container logs on failure using Docker API
			if resource.Container != nil {
				var buf bytes.Buffer
				err := j.pool.Client.Logs(docker.LogsOptions{
					Container:    resource.Container.ID,
					OutputStream: &buf,
					Follow:       false,
					Stdout:       true,
					Stderr:       true,
				})
				if err == nil && buf.Len() > 0 {
					j.tb.Logf("Container logs: %s", buf.String())
				} else if err != nil {
					j.tb.Logf("Failed to get container logs: %v", err)
				}
			}
			return fmt.Errorf("endpoints not ready - collector: %v, query: %v, otlp-grpc: %v, otlp-http: %v (waited %v)",
				collectorReady, queryReady, otlpGrpcReady, otlpHttpReady, time.Since(startTime))
		}

		j.tb.Logf("Container health check passed after %v", time.Since(startTime))
		return nil
	}); err != nil {
		return fmt.Errorf("container health check failed: %v", err)
	}

	return nil
}

// isEndpointReady performs a quick check if an endpoint is responding
func isEndpointReady(endpoint string) bool {
	client := &http.Client{
		Timeout: time.Second * 10, // Increased timeout for endpoint checks
	}
	resp, err := client.Get(endpoint)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode < 500 // Accept 2xx, 3xx, and 4xx as "ready"
}
