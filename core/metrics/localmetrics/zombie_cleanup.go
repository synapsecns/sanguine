package localmetrics

import (
	"fmt"
	"os/exec"
	"strings"
)

// cleanupZombieProcesses attempts to cleanup zombie processes that might be holding ports
func (j *testJaeger) cleanupZombieProcesses(port string) error {
	// Use ps to find zombie processes
	cmd := exec.Command("ps", "-eo", "pid,stat,command")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("ps command failed: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		// Check for zombie processes (status Z)
		if strings.Contains(line, " Z ") {
			fields := strings.Fields(line)
			if len(fields) >= 1 {
				pid := fields[0]
				j.tb.Logf("Found zombie process %s, attempting cleanup", pid)

				// Try to kill the zombie process and its parent
				killCmd := exec.Command("sudo", "kill", "-9", pid)
				if err := killCmd.Run(); err != nil {
					j.tb.Logf("Warning: Failed to kill zombie process %s: %v", pid, err)
				}
			}
		}
	}

	return nil
}
