package localmetrics

import (
	"context"

	"github.com/Flaque/filet"

	// embeds the pyroscope config file.
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/ory/dockertest/v3/docker/types/mount"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dockerutil"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/synapsecns/sanguine/core/retry"
)

//go:embed pyroscope.yaml
var pyroscopeConfig string

// pyroscopePath is the local path to the pyroscope config file.
const pyroscopePath = "/pyroscope.yaml"

// StartPyroscopeServer starts a new pyroscope instance.
// See:  https://github.com/grafana/pyroscope/blob/main/examples/tracing/jaeger/docker-compose.yml
// for details.
// nolint: cyclop
func (j *testJaeger) StartPyroscopeServer(ctx context.Context) *uiResource {
	if !j.cfg.enablePyroscope {
		return &uiResource{
			uiURL: core.GetEnv(internal.PyroscopeEndpoint, fmt.Sprintf("%s not found", internal.PyroscopeEndpoint)),
		}
	}
	if core.HasEnv(internal.PyroscopeEndpoint) {
		return &uiResource{
			uiURL: os.Getenv(internal.PyroscopeEndpoint),
		}
	}

	runOptions := &dockertest.RunOptions{
		Hostname:   "pyroscope",
		Repository: "pyroscope/pyroscope",
		Env: []string{
			fmt.Sprintf("PYROSCOPE_CONFIG=%s", pyroscopePath),
		},
		Tag:          "latest",
		Cmd:          []string{"server"},
		ExposedPorts: []string{"4040"},
		Networks:     j.getNetworks(),
		Labels: map[string]string{
			appLabel:   "pyroscope",
			runIDLabel: j.runID,
		},
	}

	// github actions functions on a bridge so the host mount happens on the machine host
	// rather than in the isolated container environment. This causes the pyroscope config
	// to be inaccessible. To get around this we remove the environment variables from the
	// container.
	// See:  https://stackoverflow.com/a/60202672 for details.
	if core.HasEnv("CI") {
		runOptions.Env = []string{}
	}

	resource, err := j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.Mounts = []docker.HostMount{
			{
				Type:     string(mount.TypeBind),
				Target:   pyroscopePath,
				Source:   filet.TmpFile(j.tb, "", pyroscopeConfig).Name(),
				ReadOnly: true,
			},
		}
		config.VolumesFrom = []string{}
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	assert.Nil(j.tb, err)

	j.tb.Setenv(internal.PyroscopeEndpoint, fmt.Sprintf("http://localhost:%s", dockerutil.GetPort(resource, "4040/tcp")))

	if !j.cfg.keepContainers {
		err = resource.Expire(uint(keepAliveOnFailure.Seconds()))
		assert.Nil(j.tb, err)
	}

	// make sure client is alive
	err = retry.WithBackoff(ctx, checkURL(os.Getenv(internal.PyroscopeEndpoint)), retry.WithMax(time.Millisecond*10), retry.WithMax(time.Minute), retry.WithMaxAttempts(100))
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
					uiURL:    os.Getenv(internal.PyroscopeEndpoint),
				}:
					return
				}
			}))
	}()
	// make sure client is alive
	err = retry.WithBackoff(ctx, checkURL(os.Getenv(internal.PyroscopeEndpoint)), retry.WithMax(time.Millisecond*10), retry.WithMax(time.Minute))
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
