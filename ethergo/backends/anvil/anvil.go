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
	"github.com/synapsecns/sanguine/ethergo/chain"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"math/big"
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
	// containers should be removed on their own, but this is a safety net.
	// to prevent old containers from piling up, we set a timeout to remove the container.
	const resourceLifetime = uint(4000)

	assert.Nil(t, resource.Expire(resourceLifetime))

	address := fmt.Sprintf("%s:%s", "http://localhost", resource.GetPort("8545/tcp"))

	var chainID *big.Int
	if err := pool.Retry(func() error {
		rpcClient, err := ethclient.DialContext(ctx, address)
		if err != nil {
			return fmt.Errorf("failed to connect")
		}
		chainID, err = rpcClient.ChainID(ctx)
		if err != nil {
			return fmt.Errorf("failed to get chain id: %w", err)
		}
		return nil
	}); err != nil {
		assert.Nil(t, err)
	}

	chainConfig := args.GetHardfork().ToChainConfig(chainID)

	chn, err := chain.New(ctx, &client.Config{
		RPCUrl:  []string{address},
		ChainID: int(chainConfig.ChainID.Int64()),
	})
	chn.SetChainConfig(chainConfig)
	assert.Nil(t, err)

	baseBackend, err := base.NewBaseBackend(ctx, t, chn)
	assert.Nil(t, err)

	backend := Backend{
		baseBackend,
	}

	go func() {
		<-ctx.Done()
		_ = pool.Purge(resource)
	}()
	return &backend
}
