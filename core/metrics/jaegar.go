package metrics

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"os"
	"testing"
	"time"
)

type jaegerHandler struct {
	*baseHandler
	buildInfo config.BuildInfo
}

// NewJaegerHandler creates a new jaeger handler for handling jaeger traces.
// the JAEGER_ENDPOINT environment variable must be set for this to work.
func NewJaegerHandler(buildInfo config.BuildInfo) Handler {
	return &jaegerHandler{
		buildInfo: buildInfo,
	}
}

const jaegerEnv = "JAEGER_ENDPOINT"

func (j *jaegerHandler) Start(_ context.Context) error {
	endpoint := os.Getenv(jaegerEnv)
	if endpoint == "" {
		return fmt.Errorf("could not get jaeger endpoint from env")
	}
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return fmt.Errorf("could not create jaeger exporter: %w", err)
	}

	j.baseHandler = newBaseHandler(exp, j.buildInfo)
	return nil
}

var keepAliveOnFailure = time.Minute * 5

// SetupTestJaeger creates a new test jaegar instance. If the test fails, the instance is kept alive for 5 minutes.
// we also allow a GLOBAL_JAEGAR env var to be set to a jaegar url to send all traces to in order to avoid having to boot for long running tests.
func SetupTestJaeger(tb testing.TB) {
	tb.Helper()

	pool, err := dockertest.NewPool("")
	assert.Nil(tb, err)

	runOptions := &dockertest.RunOptions{
		Repository:   "jaegertracing/all-in-one",
		Tag:          "latest",
		ExposedPorts: []string{"6831/udp", "16686"},
	}
	resource, err := pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	assert.Nil(tb, err)

	assert.Nil(tb, os.Setenv("JAEGER_ENDPOINT", fmt.Sprintf("http://localhost:%s", resource.GetPort("16686/tcp"))))

	err = resource.Expire(uint(keepAliveOnFailure.Seconds()))
	assert.Nil(tb, err)

	logger.Errorf("serving jaegar instance on http://localhost:%s", resource.GetPort("16686/tcp"))

	tb.Cleanup(func() {
		if tb.Failed() {
			logger.Warnf("Test failed, will continue serving jaegar instance on http://localhost:%s", resource.GetPort("16686/tcp"))
			err := pool.Purge(resource)
			assert.Nil(tb, err)
		}
	})
}
