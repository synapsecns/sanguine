package dockerutil

import (
	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"
)

// GetPort returns the port of a container.
// unlike resource.GetPort this returns the ports in reverse order.
// this can be useful in ci enviornments(?)
func GetPort(resource *dockertest.Resource, port string) string {
	if resource.Container == nil || resource.Container.NetworkSettings == nil {
		return ""
	}

	m, ok := resource.Container.NetworkSettings.Ports[dc.Port(port)]
	if !ok || len(m) == 0 {
		return ""
	}

	return m[len(m)-1].HostPort
}
