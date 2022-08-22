package config_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/scribe/config"
	"github.com/synapsecns/synapse-node/testutils/utils"
)

func chainConfigFixture() config.ChainConfig {
	return config.ChainConfig{
		ChainID:               gofakeit.Uint32(),
		RPCUrl:                gofakeit.URL(),
		ConfirmationThreshold: gofakeit.Uint32(),
		Contracts: config.ContractConfigs{
			"a": config.ContractConfig{
				Address:    utils.NewMockAddress().String(),
				StartBlock: gofakeit.Uint64(),
			},
		},
	}
}

func (c ConfigSuite) TestChainID() {
	chainConfig := chainConfigFixture()
	chainConfig.ChainID = 0

	ok, err := chainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrInvalidChainID)
}

func (c ConfigSuite) TestRPCUrl() {
	chainConfig := chainConfigFixture()
	chainConfig.RPCUrl = ""

	ok, err := chainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrRequiredField)
}

func (c ConfigSuite) TestChainConfigsDuplicateChainID() {
	chainConfigA := chainConfigFixture()
	chainConfigB := chainConfigFixture()

	// manually set these to the same id
	chainConfigB.ChainID = chainConfigA.ChainID
	chainConfigs := config.ChainConfigs{
		"a": chainConfigA,
		"b": chainConfigB,
	}

	ok, err := chainConfigs.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrDuplicateChainID)
}
