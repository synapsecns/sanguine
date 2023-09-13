package localmetrics

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dockerutil"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/synapsecns/sanguine/core/retry"
)

// StartJaegerServer starts a new jaeger instance.
// nolint: cyclop
func (j *testJaeger) StartJaegerServer(ctx context.Context) *uiResource {
	if core.HasEnv(internal.JaegerEndpoint) && !core.HasEnv(internal.JaegerUIEndpoint) {
		j.tb.Fatalf("%s is set but %s is not, please remove %s or set %s", internal.JaegerEndpoint, internal.JaegerUIEndpoint, internal.JaegerEndpoint, internal.JaegerUIEndpoint)
	}

	if core.HasEnv(internal.JaegerEndpoint) {
		return &uiResource{
			Resource: nil,
			uiURL:    os.Getenv(internal.JaegerUIEndpoint),
		}
	}

	runOptions := &dockertest.RunOptions{
		Repository:   "jaegertracing/all-in-one",
		Tag:          "latest",
		Hostname:     "jaeger",
		ExposedPorts: []string{"14268", "16686"},
		Networks:     j.getNetworks(),
		Labels: map[string]string{
			appLabel:   "jaeger",
			runIDLabel: j.runID,
		},
	}
	resource, err := j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	assert.Nil(j.tb, err)

	j.tb.Setenv(internal.JaegerEndpoint, fmt.Sprintf("http://localhost:%s/api/traces", dockerutil.GetPort(resource, "14268/tcp")))
	// uiEndpoint is the jaeger endpoint, we want to instead use the pyroscope endpoint
	uiEndpoint := fmt.Sprintf("http://localhost:%s", dockerutil.GetPort(resource, "16686/tcp"))

	if !j.cfg.keepContainers {
		err = resource.Expire(uint(keepAliveOnFailure.Seconds()))
		assert.Nil(j.tb, err)
	}

	logResourceChan := make(chan *uiResource, 1)

	go func() {
		_ = dockerutil.TailContainerLogs(dockerutil.WithContext(ctx), dockerutil.WithPool(j.pool), dockerutil.WithProcessLogOptions(processlog.WithLogDir(j.logDir), processlog.WithLogFileName("jaeger")), dockerutil.WithFollow(true),
			dockerutil.WithResource(resource), dockerutil.WithCallback(func(ctx context.Context, metadata processlog.LogMetadata) {
				select {
				case <-ctx.Done():
					return
				case logResourceChan <- &uiResource{
					Resource: resource,
					uiURL:    uiEndpoint,
				}:
					return
				}
			}))
	}()

	// make sure client is alive
	err = retry.WithBackoff(ctx, checkURL(os.Getenv(internal.JaegerEndpoint)), retry.WithMax(time.Millisecond*10), retry.WithMax(time.Minute))
	if err != nil {
		return nil
	}

	select {
	case <-ctx.Done():
		return nil
	case logResource := <-logResourceChan:
		// if pyroscope jaeger is enabled, we'll use that ui otherwise we'll use this one
		if !j.cfg.enablePyroscopeJaeger {
			err = os.Setenv(internal.JaegerUIEndpoint, logResource.uiURL)
			assert.Nil(j.tb, err)
		}

		return logResource
	}
}

// StartJaegerPyroscopeUI starts a new jaeger pyroscope ui instance.
func (j *testJaeger) StartJaegerPyroscopeUI(ctx context.Context) *uiResource {
	// can't enable if pyroscope is disabled
	// TODO: add a warning here.
	if core.HasEnv(internal.JaegerUIEndpoint) || !j.cfg.enablePyroscope {
		return &uiResource{
			uiURL: os.Getenv(internal.JaegerUIEndpoint),
		}
	}

	// we use this to  let pyroscope no to include profiles as span tags
	err := os.Setenv(internal.PyroscopeJaegerUIEnabled, "true")
	assert.Nil(j.tb, err)

	runOptions := &dockertest.RunOptions{
		Repository:   "ghcr.io/synapsecns/jaeger-ui-pyroscope",
		Tag:          "latest",
		ExposedPorts: []string{"80"},
		Networks:     j.getNetworks(),
		Labels: map[string]string{
			appLabel:   "jaeger-ui",
			runIDLabel: j.runID,
		},
	}
	resource, err := j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	assert.Nil(j.tb, err)

	// must only be done after the container is started
	j.tb.Setenv(internal.JaegerUIEndpoint, fmt.Sprintf("http://localhost:%s", dockerutil.GetPort(resource, "80/tcp")))

	if !j.cfg.keepContainers {
		err = resource.Expire(uint(keepAliveOnFailure.Seconds()))
		assert.Nil(j.tb, err)
	}

	logResourceChan := make(chan *uiResource, 1)

	go func() {
		_ = dockerutil.TailContainerLogs(dockerutil.WithContext(ctx), dockerutil.WithPool(j.pool), dockerutil.WithProcessLogOptions(processlog.WithLogDir(j.logDir), processlog.WithLogFileName("jaeger-pyroscope-ui")), dockerutil.WithFollow(true),
			dockerutil.WithResource(resource), dockerutil.WithCallback(func(ctx context.Context, metadata processlog.LogMetadata) {
				select {
				case <-ctx.Done():
					return
				case logResourceChan <- &uiResource{
					Resource: resource,
					uiURL:    os.Getenv(internal.JaegerUIEndpoint),
				}:
					return
				}
			}))
	}()

	// make sure client is alive
	err = retry.WithBackoff(ctx, checkURL(os.Getenv(internal.JaegerEndpoint)), retry.WithMax(time.Millisecond*10), retry.WithMax(time.Minute))
	if err != nil {
		return nil
	}

	select {
	case <-ctx.Done():
		return nil
	case logResource := <-logResourceChan:
		return logResource
	}
}
