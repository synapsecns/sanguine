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
			Resource: nil,
			uiURL:    core.GetEnv(internal.PyroscopeEndpoint, "pyroscope not enabled"),
		}
	}
	if core.HasEnv(internal.PyroscopeEndpoint) {
		return &uiResource{
			Resource: nil,
			uiURL:    os.Getenv(internal.PyroscopeEndpoint),
		}
	}

	runOptions := &dockertest.RunOptions{
		Hostname:   "pyroscope",
		Repository: "pyroscope/pyroscope",
		Env: []string{
			fmt.Sprintf("PYROSCOPE_CONFIG=%s", pyroscopePath),
			"PYROSCOPE_LOG_LEVEL=debug",
		},
		Tag:          "latest",
		Cmd:          []string{"server"},
		ExposedPorts: []string{"4040"},
		Networks: j.getNetworks(),
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

	var resource *dockertest.Resource
	var err error

	// Create container with improved retry logic
	err = retry.WithBackoff(ctx, func(ctx context.Context) error {
		// Wait for network stability
		time.Sleep(time.Second * 5)

		// Create temporary config file
		tmpFile := filet.TmpFile(j.tb, "", pyroscopeConfig)
		if tmpFile == nil {
			return fmt.Errorf("failed to create temporary pyroscope config file")
		}
		j.tb.Logf("Created temporary pyroscope config file at: %s", tmpFile.Name())

		// Start container
		resource, err = j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
			config.Mounts = []docker.HostMount{
				{
					Type:     string(mount.TypeBind),
					Target:   pyroscopePath,
					Source:   tmpFile.Name(),
					ReadOnly: true,
				},
			}
			config.VolumesFrom = []string{}
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
			config.PublishAllPorts = true
		})
		if err != nil {
			j.tb.Logf("Failed to start Pyroscope container: %v", err)
			return err
		}

		// Wait for container to initialize
		time.Sleep(time.Second * 2)

		// Validate port with retries
		var port string
		err = retry.WithBackoff(ctx, func(ctx context.Context) error {
			port = dockerutil.GetPort(resource, "4040/tcp")
			if port == "" {
				return fmt.Errorf("port not available")
			}
			return nil
		}, retry.WithMax(time.Second*5),
			retry.WithMaxAttempts(10))
		if err != nil {
			return fmt.Errorf("failed to get Pyroscope port: %v", err)
		}

		// Set environment variable
		endpoint := fmt.Sprintf("http://localhost:%s", port)
		if err := os.Setenv(internal.PyroscopeEndpoint, endpoint); err != nil {
			return fmt.Errorf("failed to set Pyroscope endpoint: %v", err)
		}

		// Wait for endpoint with increased timeout and better error handling
		return retry.WithBackoff(ctx, func(ctx context.Context) error {
			err := checkURL(endpoint)(ctx)
			if err != nil {
				return fmt.Errorf("pyroscope not ready: %v", err)
			}
			return nil
		},
			retry.WithMax(time.Second*10),
			retry.WithMaxAttempts(20))
	},
		retry.WithMax(time.Second*15),
		retry.WithMaxAttempts(5))

	if err != nil {
		j.tb.Logf("Failed to start Pyroscope container after retries: %v", err)
		return nil
	}

	if !j.cfg.keepContainers {
		if err = resource.Expire(uint(keepAliveOnFailure.Seconds())); err != nil {
			j.tb.Logf("Failed to set container expiry: %v", err)
		}
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

	select {
	case <-ctx.Done():
		return nil
	case logResource := <-logResourceChan:
		return logResource
	}
}
