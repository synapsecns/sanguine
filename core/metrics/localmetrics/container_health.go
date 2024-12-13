package localmetrics

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ory/dockertest/v3"
)

// waitForContainerHealth waits for the container to be healthy
func (j *testJaeger) waitForContainerHealth(resource *dockertest.Resource) error {
	maxWait := time.Second * 30
	startTime := time.Now()

	for {
		if time.Since(startTime) > maxWait {
			return fmt.Errorf("container health check timed out after %v", maxWait)
		}

		// Use the Docker client to inspect the container
		container, err := j.pool.Client.InspectContainer(resource.Container.ID)
		if err != nil {
			j.tb.Logf("Failed to inspect container: %v", err)
			return err
		}

		if container.State.Health != nil && container.State.Health.Status != "" {
			if container.State.Health.Status == "healthy" {
				return nil
			}
			j.tb.Logf("Container health status: %s", container.State.Health.Status)
		} else {
			// If no health check is configured, check if container is running
			if container.State.Running {
				return nil
			}
			j.tb.Logf("Container state: %s", container.State.Status)
		}

		time.Sleep(time.Second)
	}
}

// verifyEndpoint attempts to connect to an endpoint with retries
func (j *testJaeger) verifyEndpoint(endpoint string, maxRetries int) error {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	for i := 0; i < maxRetries; i++ {
		resp, err := client.Get(endpoint)
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode < 500 {  // Accept 2xx, 3xx, and 4xx as "ready"
				return nil
			}
		}
		time.Sleep(time.Second)
	}

	return fmt.Errorf("endpoint %s not ready after %d attempts", endpoint, maxRetries)
}
