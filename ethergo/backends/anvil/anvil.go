package anvil

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mcuadros/go-defaults"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"math/big"
	"testing"
	"time"
)

type Backend struct {
	*base.Backend
	// wsURL of the backend
	wsURL string
	// chainName is the name of the chain
	chainName string
}

type Config struct {
	BlockTimeSeconds *uint
	Accounts         uint   `default:"10"`
	BalanceETH       uint   `default:"10000"`
	DerivationPath   string `default:"m/44'/60'/0'/0/"`
	// TODO: Default needs to be set explicitly here
	Hardfork     Hardfork `default:"14"`
	Mnemonic     string   `default:"test test test test test test test test test test test junk"`
	NoMining     bool
	Port         uint16
	ForkSettings *ForkConfig
	EnvConfig    *EnvConfig
	Host         string `default:"0.0.0.0"`
}

type ForkConfig struct {
	URL              string
	BlockNumber      uint32
	NoStorageCaching bool
	Retries          uint   `default:"5"`
	TimeoutMS        string `default:"45000"`
}

type EnvConfig struct {
	BlockBaseFeePerGas *big.Int
	ChainID            *uint32
	GasLimit           *uint32
	GasPrice           *uint32
}

func NewAnvilBackend(ctx context.Context, t *testing.T, cfg *Config) *Backend {
	defaults.SetDefaults(cfg)
	if cfg.Port == 0 {
		cfg.Port = uint16(freeport.GetPort())
	}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool(core.GetEnv("DOCKER_SOCKET", ""))
	assert.Nil(t, err)

	res, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:       fmt.Sprintf("foundry-%s", gofakeit.UUID()),
		Repository: core.GetEnv("FOUNDRY_IMAGE", "ghcr.io/foundry-rs/foundry"),
		Tag:        core.GetEnv("FOUNDRY_TAG", "latest"),
		// TODO: build options
		Cmd: []string{"anvil", "--host", "0.0.0.0", "--port", fmt.Sprintf("%d", cfg.Port)},
	}, func(c *docker.HostConfig) {
		c.AutoRemove = true
		c.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	assert.Nil(t, err)

	_ = res

	if err := pool.Retry(func() error {
		ctx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		testClient, err := ethclient.DialContext(ctx, fmt.Sprintf("http://0.0.0.0:%d", cfg.Port))
		if err != nil {
			return fmt.Errorf("could not connect to anvil: %w", err)
		}

		_, err = testClient.ChainID(ctx)
		if err != nil {
			return fmt.Errorf("could not get chain id from anvil: %w", err)
		}

		return nil
	}); err != nil {
		assert.Nil(t, err)
	}

	return nil
}
