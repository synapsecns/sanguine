package config_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/synapse-node/testutils/utils"
	"os"
	"testing"
)

func domainConfigFixture(t testing.TB) config.DomainConfig {
	return config.DomainConfig{
		DomainID:              gofakeit.Uint32(),
		Type:                  types.AllChainTypes()[0].String(),
		RequiredConfirmations: gofakeit.Uint32(),
		XAppConfigAddress:     utils.NewMockAddress().String(),
		RPCUrl:                gofakeit.URL(),
		KeyFile:               filet.TmpFile(t, "", gofakeit.Password(true, true, true, true, true, 10)).Name(),
	}
}

func (c ConfigSuite) TestDomainConfigChainType() {
	domainConfig := domainConfigFixture(c.T())
	domainConfig.Type = gofakeit.StreetName()

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrInvalidChainType)
}

func (c ConfigSuite) TestDomainConfigID() {
	domainConfig := domainConfigFixture(c.T())
	domainConfig.DomainID = 0

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrInvalidDomainID)
}

func (c ConfigSuite) TestXappConfigAddressBlank() {
	domainConfig := domainConfigFixture(c.T())
	domainConfig.XAppConfigAddress = ""

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrRequiredField)
}

func (c ConfigSuite) TestXappRPCddressBlank() {
	domainConfig := domainConfigFixture(c.T())
	domainConfig.RPCUrl = ""

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrRequiredField)
}

func (c ConfigSuite) TestDomainConfigsDuplicateDomainID() {
	domainConfigA := domainConfigFixture(c.T())
	domainConfigB := domainConfigFixture(c.T())

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

func (c ConfigSuite) TestInvalidKeyFile() {
	domainConfig := domainConfigFixture(c.T())
	domainConfig.KeyFile = os.TempDir() + string(os.PathSeparator) + "noexist.key"

	ok, err := domainConfig.IsValid(c.GetTestContext())
	False(c.T(), ok)
	ErrorIs(c.T(), err, config.ErrDoesNotExist)
}
