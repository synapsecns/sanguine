package localmetrics

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"

	dc "github.com/ory/dockertest/v3/docker"
)

// cleanupPorts kills any processes using our specific ports
func (j *testJaeger) cleanupPorts() error {
	// Define specific ports we want to use
	ports := []string{"14268", "16686", "14269", "4317", "4318"}

	// First, cleanup all Docker resources
	containers, err := j.pool.Client.ListContainers(dc.ListContainersOptions{
		All: true,
		Filters: map[string][]string{
			"label": {"app=jaeger"},
		},
	})
	if err != nil {
		j.tb.Logf("Warning: Failed to list containers: %v", err)
	} else {
		for _, container := range containers {
			// Check if container is using any of our ports
			for _, port := range ports {
				for _, p := range container.Ports {
					if fmt.Sprint(p.PublicPort) == port {
						j.tb.Logf("Found container %s using port %s, stopping...", container.ID[:12], port)
						// Remove the container using dockertest's purge method
						err := j.pool.Client.RemoveContainer(dc.RemoveContainerOptions{
							ID:            container.ID,
							Force:         true,
							RemoveVolumes: true,
						})
						if err != nil {
							j.tb.Logf("Warning: Failed to remove container %s: %v", container.ID[:12], err)
						}
					}
				}
			}
		}
	}

	// Then cleanup system-wide processes
	maxRetries := 3
	for retry := 0; retry < maxRetries; retry++ {
		allPortsFree := true
		for _, port := range ports {
			// Try lsof first
			if err := j.killProcessesWithLsof(port); err != nil {
				j.tb.Logf("Warning: lsof cleanup for port %s failed: %v", port, err)
			}

			// Then try fuser
			if err := j.killProcessesWithFuser(port); err != nil {
				j.tb.Logf("Warning: fuser cleanup for port %s failed: %v", port, err)
			}

			// Double-check with netstat
			if err := j.killProcessesWithNetstat(port); err != nil {
				j.tb.Logf("Warning: netstat cleanup for port %s failed: %v", port, err)
			}

			// Verify port is free by attempting to bind to it
			if !j.isPortFree(port) {
				allPortsFree = false
				j.tb.Logf("Port %s is still in use after cleanup attempt %d", port, retry+1)
			}
		}

		if allPortsFree {
			j.tb.Logf("All ports are free after cleanup attempt %d", retry+1)
			return nil
		}

		// Wait between retries
		time.Sleep(time.Second * 5)
	}

	// Final verification
	for _, port := range ports {
		if !j.isPortFree(port) {
			return fmt.Errorf("port %s still in use after %d cleanup attempts", port, maxRetries)
		}
	}

	return nil
}

// isPortFree checks if a port is available by attempting to bind to it
func (j *testJaeger) isPortFree(port string) bool {
	addr := fmt.Sprintf(":%s", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}
	ln.Close()
	return true
}

// killProcessesWithLsof uses lsof to find and kill processes
func (j *testJaeger) killProcessesWithLsof(port string) error {
	cmd := exec.Command("sudo", "lsof", "-t", "-i", ":"+port)
	output, err := cmd.Output()
	if err != nil {
		return err // No processes is fine
	}

	if len(output) > 0 {
		pids := strings.Split(strings.TrimSpace(string(output)), "\n")
		for _, pid := range pids {
			j.tb.Logf("Killing process %s using port %s (lsof)", pid, port)
			killCmd := exec.Command("sudo", "kill", "-9", pid)
			if err := killCmd.Run(); err != nil {
				j.tb.Logf("Warning: Failed to kill process %s: %v", pid, err)
			}
		}
	}
	return nil
}

// killProcessesWithFuser uses fuser to find and kill processes
func (j *testJaeger) killProcessesWithFuser(port string) error {
	cmd := exec.Command("sudo", "fuser", "-k", port+"/tcp")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("fuser failed: %w", err)
	}
	return nil
}

// killProcessesWithNetstat uses netstat to find and kill processes
func (j *testJaeger) killProcessesWithNetstat(port string) error {
	cmd := exec.Command("sudo", "netstat", "-tlpn")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("netstat failed: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, ":"+port) {
			fields := strings.Fields(line)
			if len(fields) >= 7 {
				pidField := fields[6]
				if i := strings.Index(pidField, "/"); i != -1 {
					pid := pidField[:i]
					j.tb.Logf("Killing process %s using port %s (netstat)", pid, port)
					killCmd := exec.Command("sudo", "kill", "-9", pid)
					if err := killCmd.Run(); err != nil {
						j.tb.Logf("Warning: Failed to kill process %s: %v", pid, err)
					}
				}
			}
		}
	}
	return nil
}
