package e2e_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
	"math/big"
	"testing"
	"time"
)

type IntegrationSuite struct {
	*testsuite.TestSuite
	manager       *testutil.DeployManager
	originBackend backends.SimulatedTestBackend
	destBackend   backends.SimulatedTestBackend
	//omniserver is the omnirpc server address
	omniServer    string
	omniClient    omnirpcClient.RPCClient
	metrics       metrics.Handler
	apiServer     string
	relayer       *relayer.Relayer
	relayerWallet wallet.Wallet
	userWallet    wallet.Wallet
}

func NewIntegrationSuite(tb testing.TB) *IntegrationSuite {
	tb.Helper()
	return &IntegrationSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, NewIntegrationSuite(t))
}

const (
	originBackendChainID = 1
	destBackendChainID   = 2
)

// SetupTest sets up each test in the integration suite. We need to do a few things here:
//
// 1. Create the backends
// 2. Create a bunch of different tokens on a bunch of different chains. We do this here so we can pre-generate a lot of
// the tedious configurations for both api and relayer at the same time before testing individual cases.
// 3. Create the omnirpc server: this is used by both the api and the relayer.
func (i *IntegrationSuite) SetupTest() {
	i.TestSuite.SetupTest()

	i.manager = testutil.NewDeployManager(i.T())
	// TODO: consider jaeger
	i.metrics = metrics.NewNullHandler()
	// setup backends for ethereum & omnirpc
	i.setupBackends()

	// setup the api server
	i.setupAPI()
	i.setupRelayer()

}

// getOtherBackend gets the backend that is not the current one. This is a helper
func (i *IntegrationSuite) getOtherBackend(backend backends.SimulatedTestBackend) backends.SimulatedTestBackend {
	allBackends := core.ToSlice(i.originBackend, i.destBackend)
	for _, b := range allBackends {
		if b.GetChainID() != backend.GetChainID() {
			return b
		}
	}
	return nil
}

// TODO:
func (i *IntegrationSuite) TestUSDCtoUSDC() {
	// Before we do anything, we're going to mint ourselves some USDC on the destination chain.
	// 100k should do.
	i.manager.MintToAddress(i.GetTestContext(), i.destBackend, testutil.USDCType, i.relayerWallet.Address(), big.NewInt(100000))
	destUSDC := i.manager.Get(i.GetTestContext(), i.destBackend, testutil.USDCType)
	i.Approve(i.destBackend, destUSDC, i.relayerWallet)

	// let's give the user some money as well, $500 should do.
	const userWantAmount = 500
	i.manager.MintToAddress(i.GetTestContext(), i.originBackend, testutil.USDCType, i.userWallet.Address(), big.NewInt(userWantAmount))
	originUSDC := i.manager.Get(i.GetTestContext(), i.originBackend, testutil.USDCType)
	i.Approve(i.originBackend, originUSDC, i.userWallet)

	// non decimal adjusted user want amount
	realWantAmount, err := testutil.AdjustAmount(i.GetTestContext(), big.NewInt(userWantAmount), destUSDC.ContractHandle())
	i.NoError(err)

	// now our friendly user is going to check the quote and send us some USDC on the origin chain.
	i.Eventually(func() bool {
		// first he's gonna check the quotes.
		userAPIClient, err := client.NewClient(i.apiServer, localsigner.NewSigner(i.userWallet.PrivateKey()))
		i.NoError(err)

		allQuotes, err := userAPIClient.GetAllQuotes()
		i.NoError(err)

		// let's figure out the amount of usdc we need

		for _, quote := range allQuotes {
			if common.HexToAddress(quote.DestTokenAddr) == destUSDC.Address() {

				if quote.DestAmount.BigInt().Cmp(realWantAmount) > 0 {
					// we found our quote!
					// now we can move on
					return true
				}
			}
		}
		return false
	})

	// now we can send the money
	_, originFastBridge := i.manager.GetFastBridge(i.GetTestContext(), i.originBackend)
	auth := i.originBackend.GetTxContext(i.GetTestContext(), i.userWallet.AddressPtr())
	// we want 499 usdc for 500 requested within a day
	tx, err := originFastBridge.Bridge(auth.TransactOpts, fastbridge.IFastBridgeBridgeParams{
		DstChainId:   uint32(i.destBackend.GetChainID()),
		To:           i.userWallet.Address(),
		OriginToken:  originUSDC.Address(),
		DestToken:    destUSDC.Address(),
		OriginAmount: realWantAmount,
		DestAmount:   new(big.Int).Sub(realWantAmount, big.NewInt(1)),
		Deadline:     new(big.Int).SetInt64(time.Now().Add(time.Hour * 24).Unix()),
	})
	i.NoError(err)
	i.originBackend.WaitForConfirmation(i.GetTestContext(), tx)

	time.Sleep(time.Second * 100)
}
