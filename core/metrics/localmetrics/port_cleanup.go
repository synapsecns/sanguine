package localmetrics

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// cleanupPorts kills any processes using our specific ports
func (j *testJaeger) cleanupPorts() error {
	// Define specific ports we want to use
	ports := []string{"14268", "16686", "14269", "4317", "4318"}

	// First, try to cleanup Docker-specific processes
	if err := j.cleanupDockerPorts(ports); err != nil {
		j.tb.Logf("Warning: Docker port cleanup failed: %v", err)
	}

	// Then cleanup system-wide processes
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
	}

	// Wait for ports to be fully released
	time.Sleep(time.Second * 2)

	return nil
}

// cleanupDockerPorts attempts to cleanup Docker-specific port bindings
func (j *testJaeger) cleanupDockerPorts(ports []string) error {
	// Find containers using our ports
	for _, port := range ports {
		cmd := exec.Command("docker", "ps", "-q", "--filter", fmt.Sprintf("publish=%s", port))
		output, err := cmd.Output()
		if err != nil {
			continue // No containers found is fine
		}

		if len(output) > 0 {
			containers := strings.Split(strings.TrimSpace(string(output)), "\n")
			for _, container := range containers {
				j.tb.Logf("Stopping Docker container %s using port %s", container, port)
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
	return nil
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
