package executor_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/config/executor"
)

func chainConfigFixture() executor.ChainConfig {
	return executor.ChainConfig{
		ChainID:            gofakeit.Uint32(),
		OriginAddress:      gofakeit.Word(),
		DestinationAddress: gofakeit.Word(),
	}
}

func (c ConfigSuite) TestChainID() {
	chainConfig := chainConfigFixture()
	chainConfig.ChainID = 0

	ok, err := chainConfig.IsValid(c.GetTestContext())
	NotNil(c.T(), err)
	False(c.T(), ok)
}

func (c ConfigSuite) TestChainConfigsDuplicateChainID() {
	chainConfigA := chainConfigFixture()
	chainConfigB := chainConfigFixture()

	// manually set these to the same id
	chainConfigB.ChainID = chainConfigA.ChainID
	chainConfigs := executor.ChainConfigs{
		chainConfigA,
		chainConfigB,
	}

	ok, err := chainConfigs.IsValid(c.GetTestContext())
	NotNil(c.T(), err)
	False(c.T(), ok)
}

func (c ConfigSuite) TestInvalidAddresses() {
	chainConfig := chainConfigFixture()
	chainConfig.OriginAddress = ""

	ok, err := chainConfig.IsValid(c.GetTestContext())
	NotNil(c.T(), err)
	False(c.T(), ok)

	chainConfig = chainConfigFixture()
	chainConfig.DestinationAddress = ""

	ok, err = chainConfig.IsValid(c.GetTestContext())
	NotNil(c.T(), err)
	False(c.T(), ok)
}
