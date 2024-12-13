package localmetrics

import (
	"fmt"
	"net"
	"strconv"
)

// getPortsInUse checks if the Jaeger ports are currently in use
func (j *testJaeger) getPortsInUse() []string {
	ports := []string{"14268", "16686"}
	inUse := make([]string, 0)

	for _, port := range ports {
		portNum, _ := strconv.Atoi(port)
		addr := fmt.Sprintf(":%d", portNum)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			// Port is in use
			inUse = append(inUse, port)
		} else {
			// Port is free, close the listener
			_ = listener.Close()
		}
	}

	return inUse
}
