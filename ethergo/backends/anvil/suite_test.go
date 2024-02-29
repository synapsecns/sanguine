package anvil_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"github.com/synapsecns/sanguine/ethergo/example"
	"github.com/synapsecns/sanguine/ethergo/example/counter"
	"github.com/synapsecns/sanguine/ethergo/manager"
)

var vitalik = common.HexToAddress("0xd8da6bf26964af9d7eed9e03e53415d37aa96045")

type AnvilSuite struct {
	*testsuite.TestSuite
	backend     *anvil.Backend
	options     *anvil.OptionBuilder
	client      *anvil.Client
	forkAddress string
	counter     *counter.CounterRef
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

	a.forkAddress = core.GetEnv("ETHEREUM_RPC_URI", "https://1rpc.io/eth")
	options := anvil.NewAnvilOptionBuilder()
	err := options.SetForkURL(a.forkAddress)
	Nil(a.T(), err)

	// enable otterscan
	options.OtterscanEnabled(true)

	a.backend = anvil.NewAnvilBackend(a.GetSuiteContext(), a.T(), options)
	a.options = options
	a.client, err = anvil.Dial(a.GetSuiteContext(), a.backend.RPCAddress())
	Nil(a.T(), err)

	deployer := manager.NewDeployerManager(a.T(), example.NewCounterDeployer)
	deployedContract := deployer.Get(a.GetSuiteContext(), a.backend, example.CounterType)

	var ok bool
	a.counter, ok = deployedContract.ContractHandle().(*counter.CounterRef)
	True(a.T(), ok)
}

func TestAnvilSuite(t *testing.T) {
	suite.Run(t, NewAnvilSuite(t))
}
