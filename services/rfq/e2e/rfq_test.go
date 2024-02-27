package e2e_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/anvil"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	cctpTest "github.com/synapsecns/sanguine/services/cctp-relayer/testutil"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/service"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
)

type IntegrationSuite struct {
	*testsuite.TestSuite
	manager           *testutil.DeployManager
	cctpDeployManager *cctpTest.DeployManager
	originBackend     backends.SimulatedTestBackend
	destBackend       backends.SimulatedTestBackend
	//omniserver is the omnirpc server address
	omniServer    string
	omniClient    omnirpcClient.RPCClient
	metrics       metrics.Handler
	store         reldb.Service
	apiServer     string
	relayer       *service.Relayer
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
	destBackendChainID   = 43114
)

// SetupTest sets up each test in the integration suite. We need to do a few things here:
//
// 1. Create the backends
// 2. Create a bunch of different tokens on a bunch of different chains. We do this here so we can pre-generate a lot of
// the tedious configurations for both api and relayer at the same time before testing individual cases.
// 3. Create the omnirpc server: this is used by both the api and the relayer.
func (i *IntegrationSuite) SetupTest() {
	i.TestSuite.SetupTest()

	// TODO: no need for this when anvil CI issues are fixed
	if core.GetEnvBool("CI", false) {
		return
	}

	i.manager = testutil.NewDeployManager(i.T())
	i.cctpDeployManager = cctpTest.NewDeployManager(i.T())
	// TODO: consider jaeger
	i.metrics = metrics.NewNullHandler()
	// setup backends for ethereum & omnirpc
	i.setupBackends()

	// setup the api server
	i.setupQuoterAPI()
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

func (i *IntegrationSuite) TestUSDCtoUSDC() {
	if core.GetEnvBool("CI", false) {
		i.T().Skip("skipping until anvil issues are fixed in CI")
	}

	// load token contracts
	const startAmount = 1000
	const rfqAmount = 900
	opts := i.destBackend.GetTxContext(i.GetTestContext(), nil)
	destUSDC, destUSDCHandle := i.cctpDeployManager.GetMockMintBurnTokenType(i.GetTestContext(), i.destBackend)
	realStartAmount, err := testutil.AdjustAmount(i.GetTestContext(), big.NewInt(startAmount), destUSDC.ContractHandle())
	i.NoError(err)
	realRFQAmount, err := testutil.AdjustAmount(i.GetTestContext(), big.NewInt(rfqAmount), destUSDC.ContractHandle())
	i.NoError(err)

	// add initial usdc to relayer on destination
	tx, err := destUSDCHandle.MintPublic(opts.TransactOpts, i.relayerWallet.Address(), realStartAmount)
	i.Nil(err)
	i.destBackend.WaitForConfirmation(i.GetTestContext(), tx)
	i.Approve(i.destBackend, destUSDC, i.relayerWallet)

	// add initial USDC to relayer on origin
	optsOrigin := i.originBackend.GetTxContext(i.GetTestContext(), nil)
	originUSDC, originUSDCHandle := i.cctpDeployManager.GetMockMintBurnTokenType(i.GetTestContext(), i.originBackend)
	tx, err = originUSDCHandle.MintPublic(optsOrigin.TransactOpts, i.relayerWallet.Address(), realStartAmount)
	i.Nil(err)
	i.originBackend.WaitForConfirmation(i.GetTestContext(), tx)
	i.Approve(i.originBackend, originUSDC, i.relayerWallet)

	// add initial USDC to user on origin
	tx, err = originUSDCHandle.MintPublic(optsOrigin.TransactOpts, i.userWallet.Address(), realRFQAmount)
	i.Nil(err)
	i.originBackend.WaitForConfirmation(i.GetTestContext(), tx)
	i.Approve(i.originBackend, originUSDC, i.userWallet)

	// non decimal adjusted user want amount
	// now our friendly user is going to check the quote and send us some USDC on the origin chain.
	i.Eventually(func() bool {
		// first he's gonna check the quotes.
		userAPIClient, err := client.NewAuthenticatedClient(metrics.Get(), i.apiServer, localsigner.NewSigner(i.userWallet.PrivateKey()))
		i.NoError(err)

		allQuotes, err := userAPIClient.GetAllQuotes()
		i.NoError(err)

		// let's figure out the amount of usdc we need
		for _, quote := range allQuotes {
			if common.HexToAddress(quote.DestTokenAddr) == destUSDC.Address() {
				destAmountBigInt, _ := new(big.Int).SetString(quote.DestAmount, 10)
				if destAmountBigInt.Cmp(realRFQAmount) > 0 {
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
	tx, err = originFastBridge.Bridge(auth.TransactOpts, fastbridge.IFastBridgeBridgeParams{
		DstChainId:   uint32(i.destBackend.GetChainID()),
		To:           i.userWallet.Address(),
		OriginToken:  originUSDC.Address(),
		SendChainGas: true,
		DestToken:    destUSDC.Address(),
		OriginAmount: realRFQAmount,
		DestAmount:   new(big.Int).Sub(realRFQAmount, big.NewInt(10_000_000)),
		Deadline:     new(big.Int).SetInt64(time.Now().Add(time.Hour * 24).Unix()),
	})
	i.NoError(err)
	i.originBackend.WaitForConfirmation(i.GetTestContext(), tx)

	// TODO: this, but cleaner
	anvilClient, err := anvil.Dial(i.GetTestContext(), i.originBackend.RPCAddress())
	i.NoError(err)

	go func() {
		for {
			select {
			case <-i.GetTestContext().Done():
				return
			case <-time.After(time.Second * 4):
				// increase time by 30 mintutes every second, should be enough to get us a fastish e2e test
				// we don't need to worry about deadline since we're only doing this on origin
				err = anvilClient.IncreaseTime(i.GetTestContext(), 60*30)
				i.NoError(err)

				// because can claim works on last block timestamp, we need to do something
				err = anvilClient.Mine(i.GetTestContext(), 1)
				i.NoError(err)
			}
		}

	}()

	// since relayer started w/ 0 usdc, once they're offering the inventory up on origin chain we know the workflow completed
	i.Eventually(func() bool {
		// first he's gonna check the quotes.
		relayerAPIClient, err := client.NewAuthenticatedClient(metrics.Get(), i.apiServer, localsigner.NewSigner(i.relayerWallet.PrivateKey()))
		i.NoError(err)

		allQuotes, err := relayerAPIClient.GetAllQuotes()
		i.NoError(err)

		// let's figure out the amount of usdc we need
		for _, quote := range allQuotes {
			if common.HexToAddress(quote.DestTokenAddr) == originUSDC.Address() && quote.DestChainID == originBackendChainID {

				// we should now have some usdc on the origin chain since we claimed
				// this should be offered up as inventory
				destAmountBigInt, _ := new(big.Int).SetString(quote.DestAmount, 10)
				if destAmountBigInt.Cmp(realStartAmount) > 0 {
					// we found our quote!
					// now we can move on
					return true
				}
			}
		}
		return false
	})

	i.Eventually(func() bool {
		// check to see if the USDC balance has decreased on destination due to rebalance
		balance, err := originUSDCHandle.BalanceOf(&bind.CallOpts{Context: i.GetTestContext()}, i.relayerWallet.Address())
		i.NoError(err)
		balanceThresh, _ := new(big.Float).Mul(big.NewFloat(1.5), new(big.Float).SetInt(realStartAmount)).Int(nil)
		if balance.Cmp(balanceThresh) > 0 {
			return false
		}

		// check to see if there is a pending rebalance from the destination back to origin
		// TODO: validate more of the rebalance- expose in db interface just for testing?
		destPending, err := i.store.HasPendingRebalance(i.GetTestContext(), uint64(i.destBackend.GetChainID()))
		i.NoError(err)
		if !destPending {
			return false
		}
		originPending, err := i.store.HasPendingRebalance(i.GetTestContext(), uint64(i.originBackend.GetChainID()))
		i.NoError(err)
		if !originPending {
			return false
		}
		return true
	})
}

// nolint: cyclop
func (i *IntegrationSuite) TestETHtoETH() {
	if core.GetEnvBool("CI", false) {
		i.T().Skip("skipping until anvil issues are fixed in CI")
	}
	// Send ETH to the relayer on destination
	const initialBalance = 10
	i.destBackend.FundAccount(i.GetTestContext(), i.relayerWallet.Address(), *big.NewInt(initialBalance))

	// let's give the user some money as well
	const userWantAmount = 1
	i.originBackend.FundAccount(i.GetTestContext(), i.userWallet.Address(), *big.NewInt(userWantAmount))

	// non decimal adjusted user want amount
	realWantAmount := new(big.Int).Mul(big.NewInt(userWantAmount), big.NewInt(1e18))

	// now our friendly user is going to check the quote and send us some ETH on the origin chain.
	i.Eventually(func() bool {
		// first he's gonna check the quotes.
		userAPIClient, err := client.NewAuthenticatedClient(metrics.Get(), i.apiServer, localsigner.NewSigner(i.userWallet.PrivateKey()))
		i.NoError(err)

		allQuotes, err := userAPIClient.GetAllQuotes()
		i.NoError(err)

		// let's figure out the amount of ETH we need
		for _, quote := range allQuotes {
			if common.HexToAddress(quote.DestTokenAddr) == chain.EthAddress {
				destAmountBigInt, _ := new(big.Int).SetString(quote.DestAmount, 10)
				if destAmountBigInt.Cmp(realWantAmount) > 0 {
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
	auth.TransactOpts.Value = realWantAmount
	// we want 499 ETH for 500 requested within a day
	tx, err := originFastBridge.Bridge(auth.TransactOpts, fastbridge.IFastBridgeBridgeParams{
		DstChainId:   uint32(i.destBackend.GetChainID()),
		To:           i.userWallet.Address(),
		OriginToken:  chain.EthAddress,
		SendChainGas: true,
		DestToken:    chain.EthAddress,
		OriginAmount: realWantAmount,
		DestAmount:   new(big.Int).Sub(realWantAmount, big.NewInt(1e17)),
		Deadline:     new(big.Int).SetInt64(time.Now().Add(time.Hour * 24).Unix()),
	})
	i.NoError(err)
	i.originBackend.WaitForConfirmation(i.GetTestContext(), tx)

	// TODO: this, but cleaner
	for _, rpcAddr := range []string{i.originBackend.RPCAddress(), i.destBackend.RPCAddress()} {
		anvilClient, err := anvil.Dial(i.GetTestContext(), rpcAddr)
		i.NoError(err)

		go func() {
			for {
				select {
				case <-i.GetTestContext().Done():
					return
				case <-time.After(time.Second * 4):
					// increase time by 30 mintutes every second, should be enough to get us a fastish e2e test
					// we don't need to worry about deadline since we're only doing this on origin
					err = anvilClient.IncreaseTime(i.GetTestContext(), 60*30)
					i.NoError(err)

					// because can claim works on last block timestamp, we need to do something
					err = anvilClient.Mine(i.GetTestContext(), 1)
					i.NoError(err)
				}
			}
		}()
	}

	// since relayer started w/ 0 ETH, once they're offering the inventory up on origin chain we know the workflow completed
	i.Eventually(func() bool {
		// first he's gonna check the quotes.
		relayerAPIClient, err := client.NewAuthenticatedClient(metrics.Get(), i.apiServer, localsigner.NewSigner(i.relayerWallet.PrivateKey()))
		i.NoError(err)

		allQuotes, err := relayerAPIClient.GetAllQuotes()
		i.NoError(err)

		// let's figure out the amount of ETH we need
		for _, quote := range allQuotes {
			if common.HexToAddress(quote.DestTokenAddr) == chain.EthAddress && quote.DestChainID == originBackendChainID {
				// we should now have some ETH on the origin chain since we claimed
				// this should be offered up as inventory
				destAmountBigInt, _ := new(big.Int).SetString(quote.DestAmount, 10)
				if destAmountBigInt.Cmp(realWantAmount) > 0 {
					// we found our quote!
					// now we can move on
					return true
				}
			}
		}
		return false
	})
}
