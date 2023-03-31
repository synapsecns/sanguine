package localserver

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pyroscope-io/client/pyroscope"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dockerutil"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/synapsecns/sanguine/core/retry"
	"os"
	"runtime"
	"sync"
	"time"
)

// StartPyroscopeServer starts a new pyroscope instance.
func (j *testJaeger) StartPyroscopeServer(ctx context.Context) *uiResource {
	defer func() {
		j.StartPyroscope()
	}()

	if core.HasEnv(internal.PyroscopeEndpoint) {
		return &uiResource{
			uiURL: os.Getenv(internal.PyroscopeEndpoint),
		}
	}

	network := j.getNetwork()

	runOptions := &dockertest.RunOptions{
		Hostname:     "pyroscope",
		Repository:   "pyroscope/pyroscope",
		Tag:          "latest",
		Cmd:          []string{"server"},
		ExposedPorts: []string{"4040"},
		Networks:     []*dockertest.Network{network},
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

	j.tb.Setenv(internal.PyroscopeEndpoint, fmt.Sprintf("http://localhost:%s", resource.GetPort("4040/tcp")))

	if !debugLocal {
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

// pyroscopeMux is used to prevent multiple pyroscope instances from being started.
var pyroscopeMux sync.Mutex

// StartPyroscope starts the pyroscope profiler.
// this will not run if the pyroscope endpoint is not set.
func (j *testJaeger) StartPyroscope() {
	pyroscopeMux.Lock()
	defer pyroscopeMux.Unlock()

	if !core.HasEnv(internal.PyroscopeEndpoint) {
		return
	}

	runtime.SetMutexProfileFraction(5)
	runtime.SetBlockProfileRate(5)

	// note: because pyroscope is a profiler, we cannot use buildinfo for the application name.
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "sanguine",

		// replace this with the address of pyroscope server
		ServerAddress: os.Getenv(internal.PyroscopeEndpoint),

		// you can disable logging by setting this to nil
		Logger: pyroscopeLogger{},

		// you can provide static tags via a map:
		Tags: map[string]string{"hostname": os.Getenv("HOSTNAME")},

		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	if err != nil {
		j.tb.Error(err)
	}
}

type pyroscopeLogger struct{}

func (p pyroscopeLogger) Infof(_ string, _ ...interface{}) {
	// do nothing
}

func (p pyroscopeLogger) Debugf(_ string, _ ...interface{}) {
	// do nothing
}

func (p pyroscopeLogger) Errorf(str string, args ...interface{}) {
	logger.Warnf(str, args...)
}

var _ pyroscope.Logger = &pyroscopeLogger{}
