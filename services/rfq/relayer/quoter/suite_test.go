package quoter_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
)

// Server suite is the main API server test suite.
type QuoterSuite struct {
	*testsuite.TestSuite
	manager quoter.Manager
}

// NewServerSuite creates a end-to-end test suite.
func NewQuoterSuite(tb testing.TB) *QuoterSuite {
	tb.Helper()
	return &QuoterSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (c *QuoterSuite) SetupTest() {
	c.TestSuite.SetupTest()
	// Setup
	manager := quoter.Manager{}
	c.manager = manager
	c.manager.SetQuotableTokens(map[string][]string{
		"42161-0xaf88d065e77c8cc2239327c5edb3a432268e5831": {"1-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "10-0x0b2c639c533813f4aa9d7837caf62653d097ff85"},
		"1-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48":     {"42161-0xaf88d065e77c8cc2239327c5edb3a432268e5831", "10-0x0b2c639c533813f4aa9d7837caf62653d097ff85"},
		"10-0x0b2c639c533813f4aa9d7837caf62653d097ff85":    {"1-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "42161-0xaf88d065e77c8cc2239327c5edb3a432268e5831"},
		// Add more mock quotableTokens if needed
	})
}

func (c *QuoterSuite) SetupSuite() {
	c.TestSuite.SetupSuite()
}

// TestConfigSuite runs the integration test suite.
func TestQuoterSuite(t *testing.T) {
	suite.Run(t, NewQuoterSuite(t))
}
