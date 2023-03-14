package anvil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/ipfs/go-log"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"strings"
	"testing"
	"time"
)

type Backend struct {
	*base.Backend
}

var logger = log.Logger("anvil-backend")

// NewAnvilBackend creates a test anvil backend.
func NewAnvilBackend(ctx context.Context, t *testing.T, args *OptionBuilder) *Backend {
	t.Helper()

	pool, err := dockertest.NewPool("")
	assert.Nil(t, err)

	pool.MaxWait = time.Minute * 2
	if err != nil {
		assert.Nil(t, err)
	}

	commandArgs, err := args.Build()
	assert.Nil(t, err)

	runOptions := &dockertest.RunOptions{
		Repository: "ghcr.io/foundry-rs/foundry",
		Tag:        "latest",
		Cmd:        []string{strings.Join(append([]string{"anvil"}, commandArgs...), " ")},
		Labels: map[string]string{
			"test-id": uuid.New().String(),
		},
		ExposedPorts: []string{"8545"},
	}

	resource, err := pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	assert.Nil(t, err)

	// Docker will hard kill the container in 4000 seconds (this is a test env).
	const resourceLifetime = uint(4000)

	assert.Nil(t, resource.Expire(resourceLifetime))

	address := fmt.Sprintf("%s:%s", "http://localhost", resource.GetPort("8545/tcp"))

	if err := pool.Retry(func() error {
		rpcClient, err := ethclient.DialContext(ctx, address)
		if err != nil {
			return fmt.Errorf("failed to connect")
		}
		_, err = rpcClient.ChainID(ctx)
		if err != nil {
			return fmt.Errorf("failed to get chain id: %w", err)
		}
		return err
	}); err != nil {
		assert.Nil(t, err)
	}

	go func() {
		<-ctx.Done()
		_ = pool.Purge(resource)
	}()
	return nil
}
