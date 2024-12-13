package localmetrics

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/ory/dockertest/v3"
)

// waitForContainerHealth waits for the container to be healthy
func (j *testJaeger) waitForContainerHealth(resource *dockertest.Resource) error {
	// Set container expiration
	if err := resource.Expire(uint(time.Second * 60)); err != nil {
		return fmt.Errorf("failed to set container expiration: %v", err)
	}

	j.tb.Log("Starting container health check...")

	// Use pool's retry mechanism with timeout and detailed logging
	if err := j.pool.Retry(func() error {
		// Check container status first
		container := resource.Container
		if container == nil {
			return fmt.Errorf("container reference is nil")
		}

		j.tb.Logf("Container status: %s", container.State.Status)
		if container.State.Status != "running" {
			return fmt.Errorf("container is not running, status: %s", container.State.Status)
		}

		// Try to connect to both endpoints
		collectorEndpoint := fmt.Sprintf("http://localhost:%s/api/traces", resource.GetPort("14268/tcp"))
		queryEndpoint := fmt.Sprintf("http://localhost:%s", resource.GetPort("16686/tcp"))

		j.tb.Logf("Checking endpoints - collector: %s, query: %s", collectorEndpoint, queryEndpoint)

		// Quick check both endpoints with detailed logging
		collectorReady := isEndpointReady(collectorEndpoint)
		if !collectorReady {
			j.tb.Log("Collector endpoint not ready")
		}

		queryReady := isEndpointReady(queryEndpoint)
		if !queryReady {
			j.tb.Log("Query endpoint not ready")
		}

		if !collectorReady || !queryReady {
			// Get container logs on failure
			if resource.Container != nil {
				logsOptions := types.ContainerLogsOptions{
					ShowStdout: true,
					ShowStderr: true,
					Follow:     false,
				}
				reader, err := j.pool.Client.ContainerLogs(context.Background(), resource.Container.ID, logsOptions)
				if err == nil {
					logs, _ := io.ReadAll(reader)
					reader.Close()
					j.tb.Logf("Container logs: %s", string(logs))
				}
			}
			return fmt.Errorf("endpoints not ready - collector: %v, query: %v", collectorReady, queryReady)
		}

		j.tb.Log("Container health check passed")
		return nil
	}); err != nil {
		return fmt.Errorf("container health check failed: %v", err)
	}

	return nil
}

// isEndpointReady performs a quick check if an endpoint is responding
func isEndpointReady(endpoint string) bool {
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	resp, err := client.Get(endpoint)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode < 500 // Accept 2xx, 3xx, and 4xx as "ready"
}
