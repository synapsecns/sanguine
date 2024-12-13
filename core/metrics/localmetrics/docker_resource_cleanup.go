package localmetrics

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// cleanupDockerResources attempts to cleanup all Docker resources that might interfere with testing
func (j *testJaeger) cleanupDockerResources() error {
	// First, stop and remove ALL containers using our ports
	ports := []string{"14268", "16686", "14269", "4317", "4318"}
	for _, port := range ports {
		cmd := exec.Command("docker", "ps", "-q", "--filter", fmt.Sprintf("publish=%s", port))
		output, err := cmd.Output()
		if err != nil {
			j.tb.Logf("Warning: Failed to list containers for port %s: %v", port, err)
			continue
		}

		if len(output) > 0 {
			containers := strings.Split(strings.TrimSpace(string(output)), "\n")
			for _, container := range containers {
				j.tb.Logf("Stopping container %s using port %s", container, port)
				stopCmd := exec.Command("docker", "stop", container)
				if err := stopCmd.Run(); err != nil {
					j.tb.Logf("Warning: Failed to stop container %s: %v", container, err)
				}
				rmCmd := exec.Command("docker", "rm", "-f", container)
				if err := rmCmd.Run(); err != nil {
					j.tb.Logf("Warning: Failed to remove container %s: %v", container, err)
				}
			}
		}
	}

	// Then remove any existing jaeger networks
	cmd := exec.Command("docker", "network", "ls", "--format", "{{.Name}}")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to list networks: %w", err)
	}

	if len(output) > 0 {
		networks := strings.Split(strings.TrimSpace(string(output)), "\n")
		for _, network := range networks {
			if network == "bridge" || network == "host" || network == "none" {
				continue // Skip default networks
			}

			// First disconnect all containers from the network
			disconnectCmd := exec.Command("docker", "network", "disconnect", "-f", network)
			if err := disconnectCmd.Run(); err != nil {
				j.tb.Logf("Warning: Failed to disconnect containers from network %s: %v", network, err)
			}

			// Wait a bit for disconnection to complete
			time.Sleep(time.Second)

			// Then remove the network
			j.tb.Logf("Removing network %s", network)
			rmCmd := exec.Command("docker", "network", "rm", network)
			if err := rmCmd.Run(); err != nil {
				j.tb.Logf("Warning: Failed to remove network %s: %v", network, err)
			}
		}
	}

	// Wait for Docker to cleanup
	time.Sleep(time.Second * 3)

	return nil
}
