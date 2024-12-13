package localmetrics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ory/dockertest/v3"
)

// waitForContainerHealth waits for the container to be healthy
func (j *testJaeger) waitForContainerHealth(resource *dockertest.Resource) error {
	// Set container expiration
	if err := resource.Expire(uint(time.Second * 30)); err != nil {
		return fmt.Errorf("failed to set container expiration: %v", err)
	}

	j.tb.Log("Starting container health check...")

	// Use dockertest's built-in wait strategy
	err := resource.Wait(func() (bool, error) {
		// Try to connect to both endpoints
		collectorEndpoint := fmt.Sprintf("http://localhost:%s/api/traces", resource.GetPort("14268/tcp"))
		queryEndpoint := fmt.Sprintf("http://localhost:%s", resource.GetPort("16686/tcp"))

		j.tb.Logf("Checking endpoints - collector: %s, query: %s", collectorEndpoint, queryEndpoint)

		// Quick check both endpoints
		collectorReady := isEndpointReady(collectorEndpoint)
		queryReady := isEndpointReady(queryEndpoint)

		if !collectorReady || !queryReady {
			j.tb.Logf("Endpoints not ready - collector: %v, query: %v", collectorReady, queryReady)
			return false, nil
		}

		j.tb.Log("Container health check passed")
		return true, nil
	}, 30*time.Second)

	if err != nil {
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
