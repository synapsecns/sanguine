package localmetrics

import (
	"os/exec"
	"strings"
)

// networkExists checks if a Docker network exists
func networkExists(name string) bool {
	cmd := exec.Command("docker", "network", "ls", "--format", "{{.Name}}")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	networks := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, network := range networks {
		if network == name {
			return true
		}
	}
	return false
}
