package localmetrics

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/ory/dockertest/v3/docker"
)

// cleanupPorts kills any processes using our specific ports
func (j *testJaeger) cleanupPorts() error {
	// Define specific ports we want to use
	ports := []string{"14268", "16686"}

	// Check and kill processes using our specific ports
	for _, port := range ports {
		cmd := exec.Command("sudo", "lsof", "-t", "-i", ":"+port)
		output, err := cmd.Output()
		if err != nil {
			// lsof returns error if no processes found, which is fine
			if strings.Contains(err.Error(), "exit status 1") {
				continue
			}
			j.tb.Logf("Warning: Failed to check port %s: %v", port, err)
			continue
		}

		if len(output) > 0 {
			pids := strings.Split(strings.TrimSpace(string(output)), "\n")
			for _, pid := range pids {
				j.tb.Logf("Killing process %s using port %s", pid, port)
				killCmd := exec.Command("sudo", "kill", "-9", pid)
				if err := killCmd.Run(); err != nil {
					j.tb.Logf("Warning: Failed to kill process %s: %v", pid, err)
				}
			}
		}
	}

	// Wait for ports to be fully released
	time.Sleep(time.Second * 2)

	return nil
}
