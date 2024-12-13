package localmetrics

import (
	"fmt"
	"net"
	"time"
)

// verifyPortsAvailable checks if all required ports are actually free
func verifyPortsAvailable() error {
	ports := []string{"14268", "16686", "14269", "4317", "4318"}

	// Try binding to each port
	for _, port := range ports {
		addr := fmt.Sprintf(":%s", port)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			return fmt.Errorf("port %s is not available: %v", port, err)
		}
		listener.Close()

		// Small delay between checks
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
