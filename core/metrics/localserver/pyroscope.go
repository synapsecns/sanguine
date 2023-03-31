package localserver

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dockerutil"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/synapsecns/sanguine/core/retry"
	"os"
	"time"
)

// StartPyroscopeServer starts a new pyroscope instance.
func (j *testJaeger) StartPyroscopeServer(ctx context.Context) *uiResource {
	if core.HasEnv(internal.PYROSCOPE_ENDPOINT) {
		return &uiResource{
			uiURL: os.Getenv(internal.PYROSCOPE_ENDPOINT),
		}
	}

	runOptions := &dockertest.RunOptions{
		Repository:   "pyroscope/pyroscope",
		Tag:          "latest",
		Cmd:          []string{"server"},
		ExposedPorts: []string{"4040"},
		Labels: map[string]string{
			appLabel:   "pyroscope",
			runIDLabel: j.runID,
		},
	}

	resource, err := j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	assert.Nil(j.tb, err)

	j.tb.Setenv(internal.PYROSCOPE_ENDPOINT, fmt.Sprintf("http://localhost:%s", resource.GetPort("4040/tcp")))

	if !debugLocal {
		err = resource.Expire(uint(keepAliveOnFailure.Seconds()))
		assert.Nil(j.tb, err)
	}

	// make sure client is alive
	err = retry.WithBackoff(ctx, checkURL(os.Getenv(internal.PYROSCOPE_ENDPOINT)), retry.WithMax(time.Millisecond*10), retry.WithMax(time.Minute), retry.WithMaxAttempts(100))
	if err != nil {
		return nil
	}

	logResourceChan := make(chan *uiResource, 1)
	go func() {
		_ = dockerutil.TailContainerLogs(dockerutil.WithContext(ctx), dockerutil.WithPool(j.pool), dockerutil.WithProcessLogOptions(processlog.WithLogDir(j.logDir), processlog.WithLogFileName("pyroscope")), dockerutil.WithFollow(true),
			dockerutil.WithResource(resource), dockerutil.WithCallback(func(ctx context.Context, metadata processlog.LogMetadata) {
				select {
				case <-ctx.Done():
					return
				case logResourceChan <- &uiResource{
					Resource: resource,
					uiURL:    resource.GetPort("4040/tcp"),
				}:
					return
				}
			}))
	}()
	// make sure client is alive
	err = retry.WithBackoff(ctx, checkURL(os.Getenv(internal.PYROSCOPE_ENDPOINT)), retry.WithMax(time.Millisecond*10), retry.WithMax(time.Minute))
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
