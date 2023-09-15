package config_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/scribe/config"
)

func contractConfigFixture() config.ContractConfig {
	return config.ContractConfig{
		Address:     mocks.MockAddress().String(),
		StartBlock:  gofakeit.Uint64(),
		RefreshRate: gofakeit.Uint64(),
	}
}

func (c ConfigSuite) TestAddress() {
	contractConfig := contractConfigFixture()
	contractConfig.Address = ""

	ok, err := contractConfig.IsValid()
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrRequiredField)
}

func (c ConfigSuite) TestRefreshRate() {
	contractConfig := contractConfigFixture()
	contractConfig.Address = ""

	ok, err := contractConfig.IsValid()
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrRequiredField)
}

func (c ConfigSuite) TestContractConfigDuplicateAddress() {
	contractConfigA := contractConfigFixture()
	contractConfigB := contractConfigFixture()

	// manually set these to the same id
	contractConfigB.Address = contractConfigA.Address

	contractConfigs := config.ContractConfigs{
		contractConfigA,
		contractConfigB,
	}

	ok, err := contractConfigs.IsValid()
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrDuplicateAddress)
}
