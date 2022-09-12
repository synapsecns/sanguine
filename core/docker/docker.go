package docker

import (
	"github.com/ory/dockertest/v3"
	"net/url"
	"os"
)

// GetHostPort gets the host port of a running docker image
func GetHostPort(resource *dockertest.Resource, id string) string {
	dockerURL := os.Getenv("DOCKER_HOST")
	if dockerURL == "" {
		return resource.GetHostPort(id)
	}
	u, err := url.Parse(dockerURL)
	if err != nil {
		panic(err)
	}
	return u.Hostname() + ":" + resource.GetPort(id)
}
