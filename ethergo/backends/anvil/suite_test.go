package anvil_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"testing"
)

type AnvilSuite struct {
	*testsuite.TestSuite
	backend *anvil.Backend
	options *anvil.OptionBuilder
}

// NewAnvilSuite creates a end-to-end test suite.
func NewAnvilSuite(tb testing.TB) *AnvilSuite {
	tb.Helper()
	return &AnvilSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (a *AnvilSuite) SetupSuite() {
	a.TestSuite.SetupSuite()

	backendRPCAddress := core.GetEnv("ETH_URL", "https://rpc.ankr.com/eth")
	options := anvil.NewAnvilOptionBuilder()
	err := options.SetForkURL(backendRPCAddress)
	Nil(a.T(), err)

	a.backend = anvil.NewAnvilBackend(a.GetSuiteContext(), a.T(), options)
	a.options = options
}

func TestTestUtilSuite(t *testing.T) {
	suite.Run(t, NewAnvilSuite(t))
}
