package config_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/scribe/config"
	"github.com/synapsecns/synapse-node/testutils/utils"
)

func contractConfigFixture() config.ContractConfig {
	return config.ContractConfig{
		Address:    utils.NewMockAddress().String(),
		StartBlock: gofakeit.Uint64(),
	}
}

func (c ConfigSuite) TestAddress() {
	contractConfig := contractConfigFixture()
	contractConfig.Address = ""

	ok, err := contractConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrRequiredField)
}

func (c ConfigSuite) TestContractConfigDuplicateAddress() {
	contractConfigA := contractConfigFixture()
	contractConfigB := contractConfigFixture()

	// manually set these to the same id
	contractConfigB.Address = contractConfigA.Address

	contractConfigs := config.ContractConfigs{
		"a": contractConfigA,
		"b": contractConfigB,
	}

	ok, err := contractConfigs.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrDuplicateAddress)
}
