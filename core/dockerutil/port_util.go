package dockerutil

import (
	"github.com/ory/dockertest/v3"
	dc "github.com/ory/dockertest/v3/docker"
)

// GetPort returns the port of a container.
// We intentionally get the last port due to quirky behavior of docker-for-mac,
// which does not enforce uniqueness on 0.0.0.0:xxx ports. Not entirely sure why
// docker-for-mac behaves this way.
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
