package config_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/testsuite"
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

	// reset the name suffix after each test
	config.SetNameSuffix(config.GetDefaultNameSuffix())
}

// TestConfigSuite runs the integration test suite.
func TestConfigSuite(t *testing.T) {
	suite.Run(t, NewConfigSuite(t))
}
