package dockerutil_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/dockerutil"
	"github.com/synapsecns/sanguine/core/mocktesting"
	"github.com/synapsecns/sanguine/core/processlog"
	"os"
	"testing"
)

const stdOut = "hello world\n"

func (c *DockerSuite) TestTailContainerLogs() {
	out := testContainer(c.GetTestContext(), c.T())
	assert.FileExists(c.T(), out.CombinedFile())
	assert.FileExists(c.T(), out.StdOutFile())
	assert.FileExists(c.T(), out.StdErrFile())
}

func ExampleTailContainerLogs() {
	ctx, cancel := context.WithCancel(context.Background())

	t := mocktesting.NewMockTester("TestTailContainerLogs")
	out := testContainer(ctx, t)
	cancel()

	// get the log dir
	_ = out.LogDir()
	// get the stdout file
	_ = out.StdOutFile()
	// get the stderr file
	_ = out.StdErrFile()
	// get the combined file
	_ = out.CombinedFile()
	// Output: hello world
}

// testContainer creates a test alpine container that prints hello world.
// the return will be empty if this fails.
func testContainer(ctx context.Context, tb testing.TB) processlog.LogMetadata {
	tb.Helper()

	pool, err := dockertest.NewPool("")
	assert.Nil(tb, err)

	runOptions := &dockertest.RunOptions{
		Repository: "alpine",
		Tag:        "latest",
		Cmd:        []string{"sh", "-c", fmt.Sprintf("echo %s", stdOut)},
	}

	resource, err := pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	assert.Nil(tb, err)

	assert.Nil(tb, resource.Expire(60))

	logInfoChan := make(chan processlog.LogMetadata)
	go func() {
		// we can't check the error here
		_ = dockerutil.TailContainerLogs(dockerutil.WithContext(ctx), dockerutil.WithResource(resource), dockerutil.WithPool(pool), dockerutil.WithFollow(true), dockerutil.WithCallback(func(ctx context.Context, metadata processlog.LogMetadata) {
			select {
			case <-ctx.Done():
				return
			case logInfoChan <- metadata:
			}
		}))
	}()

	if err := pool.Retry(func() error {
		container, err := pool.Client.InspectContainer(resource.Container.ID)
		assert.Nil(tb, err)

		if !container.State.FinishedAt.IsZero() {
			return nil
		}
		return errors.New("container is still running")
	}); err != nil {
		tb.Fatal(err)
	}

	select {
	case <-ctx.Done():
		tb.Errorf("context canceled before container started")
	case logInfo := <-logInfoChan:
		fileContents := waitForFileContents(ctx, tb, logInfo.CombinedFile())

		assert.Nil(tb, err)

		// this is what's tested in the example
		fmt.Println(string(fileContents))
		return logInfo
	}
	return nil
}

func waitForFileContents(ctx context.Context, tb testing.TB, file string) []byte {
	tb.Helper()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			//nolint: gosec
			fileContents, err := os.ReadFile(file)
			assert.Nil(tb, err)

			if len(fileContents) == 0 {
				continue
			}

			return fileContents
		}
	}
}
