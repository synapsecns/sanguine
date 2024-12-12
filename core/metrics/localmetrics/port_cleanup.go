package localmetrics

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// cleanupPorts kills any processes using our port range
func (j *testJaeger) cleanupPorts() error {
	// Find processes using ports in our range (expanded range to catch all possible ports)
	cmd := exec.Command("sudo", "lsof", "-t", "-i", ":32768-65535")
	output, err := cmd.Output()
	if err != nil {
		// lsof returns error if no processes found, which is fine
		if strings.Contains(err.Error(), "exit status 1") {
			return nil
		}
		return fmt.Errorf("failed to list processes using ports: %v", err)
	}

	if len(output) == 0 {
		return nil
	}

	// Kill each process with sudo
	pids := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, pid := range pids {
		j.tb.Logf("Killing process %s using ports in our range", pid)
		killCmd := exec.Command("sudo", "kill", "-9", pid)
		if err := killCmd.Run(); err != nil {
			j.tb.Logf("Warning: Failed to kill process %s: %v", pid, err)
		}
	}

	// Wait longer for processes to be fully killed
	time.Sleep(time.Second * 5)

	return nil
}
