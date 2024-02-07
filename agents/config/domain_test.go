package config_test

import (
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/mocks"
)

func domainConfigFixture() config.DomainConfig {
	return config.DomainConfig{
		DomainID:              gofakeit.Uint32(),
		Type:                  types.AllChainTypes()[0].String(),
		RequiredConfirmations: gofakeit.Uint32(),
		OriginAddress:         mocks.MockAddress().String(),
		DestinationAddress:    mocks.MockAddress().String(),
		LightManagerAddress:   mocks.MockAddress().String(),
	}
}

func (c ConfigSuite) TestDomainConfigChainType() {
	domainConfig := domainConfigFixture()
	domainConfig.Type = gofakeit.StreetName()

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrInvalidChainType)
}

func (c ConfigSuite) TestDomainConfigID() {
	domainConfig := domainConfigFixture()
	domainConfig.DomainID = 0

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrInvalidDomainID)
}

func (c ConfigSuite) TestXappConfigAddressBlank() {
	domainConfig := domainConfigFixture()
	domainConfig.OriginAddress = ""

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrRequiredField)
}

func (c ConfigSuite) TestDomainConfigsDuplicateDomainID() {
	domainConfigA := domainConfigFixture()
	domainConfigB := domainConfigFixture()

	// manually set these to the same id
	domainConfigB.DomainID = domainConfigA.DomainID

	domainConfigs := config.DomainConfigs{
		"a": domainConfigA,
		"b": domainConfigB,
	}

	ok, err := domainConfigs.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrInvalidDomainID)
}
