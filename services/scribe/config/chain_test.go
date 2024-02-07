package config_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/scribe/config"
)

func chainConfigFixture() config.ChainConfig {
	return config.ChainConfig{
		ChainID: gofakeit.Uint32(),
		Contracts: config.ContractConfigs{
			config.ContractConfig{
				Address:    mocks.MockAddress().String(),
				StartBlock: gofakeit.Uint64(),
			},
		},
	}
}

func (c ConfigSuite) TestChainID() {
	chainConfig := chainConfigFixture()
	chainConfig.ChainID = 0

	ok, err := chainConfig.IsValid()
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrInvalidChainID)
}

func (c ConfigSuite) TestChainConfigsDuplicateChainID() {
	chainConfigA := chainConfigFixture()
	chainConfigB := chainConfigFixture()

	// manually set these to the same id
	chainConfigB.ChainID = chainConfigA.ChainID
	chainConfigs := config.ChainConfigs{
		chainConfigA,
		chainConfigB,
	}

	ok, err := chainConfigs.IsValid()
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrDuplicateChainID)
}
