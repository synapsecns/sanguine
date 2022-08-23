package config_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

// ConfigSuite is the config test suite.
type ConfigSuite struct {
	*testutils.TestSuite
}

// NewConfigSuite creates a end-to-end test suite.
func NewConfigSuite(tb testing.TB) *ConfigSuite {
	tb.Helper()
	return &ConfigSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func (c ConfigSuite) SetupTest() {
	c.TestSuite.SetupTest()
}

// TestConfigSuite runs the integration test suite.
func TestConfigSuite(t *testing.T) {
	suite.Run(t, NewConfigSuite(t))
}
