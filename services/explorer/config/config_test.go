package config_test

import (
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"math/big"
	"testing"
)

// ConfigSuite is the config test suite.
type ConfigSuite struct {
	*testsuite.TestSuite
}

// NewConfigSuite creates a end-to-end test suite.
func NewConfigSuite(tb testing.TB) *ConfigSuite {
	tb.Helper()
	return &ConfigSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (c ConfigSuite) SetupTest() {
	c.TestSuite.SetupTest()
}

// TestConfigSuite runs the integration test suite.
func TestConfigSuite(t *testing.T) {
	suite.Run(t, NewConfigSuite(t))
}

func (c ConfigSuite) TestConfig() {
	testConfig := config.Config{
		SynapseBridgeAddress:  common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		BridgeConfigV3Address: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		SwapFlashLoanAddress:  common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
	}

	encodedConfig, err := testConfig.Encode()
	Nil(c.T(), err)

	file := filet.TmpFile(c.T(), "", string(encodedConfig))
	decodedConfig, err := config.DecodeConfig(file.Name())
	Nil(c.T(), err)

	ok, err := decodedConfig.IsValid(c.GetTestContext())
	True(c.T(), ok)
	Nil(c.T(), err)
}
