```go
package localmetrics

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ory/dockertest/v3"
)

// cleanupExistingContainers removes any existing containers with names containing the given pattern
func cleanupExistingContainers(pool *dockertest.Pool, pattern string) error {
	containers, err := pool.Client.ListContainers(dockertest.ListContainersOptions{All: true})
	if err != nil {
		return fmt.Errorf("failed to list containers: %w", err)
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if strings.Contains(name, pattern) {
				// Try to stop the container first
				if err := pool.Client.StopContainer(container.ID, 1); err != nil {
					// Ignore stop errors, try to remove anyway
					fmt.Printf("Warning: failed to stop container %s: %v\n", container.ID, err)
				}

				// Wait a bit for the container to stop
				time.Sleep(time.Second)

				// Try to remove the container
				if err := pool.Client.RemoveContainer(dockertest.RemoveContainerOptions{
					ID:            container.ID,
					Force:         true,
					RemoveVolumes: true,
					Context:       context.Background(),
				}); err != nil {
					return fmt.Errorf("failed to remove container %s: %w", container.ID, err)
				}
			}
		}
	}
	return nil
}
```
