package swap_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/price"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/scribe"
	gqlClient "github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/scribe/client"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/token"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/metadata"

	"math/big"
	"testing"

	"go.uber.org/atomic"
)

type SwapFetcherSuite struct {
	*testsuite.TestSuite
	db                   db.ConsumerDB
	eventDB              scribedb.EventDB
	logIndex             atomic.Int64
	gqlClient            *gqlClient.Client
	cleanup              func()
	testBackend          backends.SimulatedTestBackend
	deployManager        *testutil.DeployManager
	bridgeConfigContract *bridgeconfig.BridgeConfigRef
	metrics              metrics.Handler
	cctpContractOp       common.Address
	cctpContractArb      common.Address
	tokenDataService     token.ITokenFetcher
	tokenPriceService    price.IPriceFetcher
	arbClient            bind.ContractBackend
	opClient             bind.ContractBackend
	consumerFetcher      scribe.IScribeFetcher
}

// NewSwapFetcherSuite creates a new swap fetcher
func NewSwapFetcherSuite(tb testing.TB) *SwapFetcherSuite {
	tb.Helper()
	return &SwapFetcherSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (p *SwapFetcherSuite) SetupSuite() {
	p.TestSuite.SetupSuite()
	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(p.GetSuiteContext(), p.T())
		metricsHandler = metrics.Jaeger
	}

	var err error
	p.metrics, err = metrics.NewByType(p.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)

	p.Require().Nil(err)
}

type TestToken struct {
	tokenID string
	bridgeconfig.BridgeConfigV3Token
}

func (c *TestToken) SetTokenConfig(bridgeConfigContract *bridgeconfig.BridgeConfigRef, opts backends.AuthType) (*types.Transaction, error) {
	tx, err := bridgeConfigContract.SetTokenConfig(opts.TransactOpts, c.tokenID, c.ChainId, common.HexToAddress(c.TokenAddress),
		c.TokenDecimals, c.MaxSwap, c.MinSwap, c.SwapFee, c.MaxSwapFee, c.MinSwapFee, c.HasUnderlying, c.IsUnderlying)
	if err != nil {
		return nil, fmt.Errorf("could not set token config: %w", err)
	}
	return tx, nil
}

var testTokens = []TestToken{{
	tokenID: gofakeit.FirstName(),
	BridgeConfigV3Token: bridgeconfig.BridgeConfigV3Token{
		ChainId:       big.NewInt(1337),
		TokenAddress:  common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
		TokenDecimals: gofakeit.Uint8(),
		MaxSwap:       new(big.Int).SetUint64(gofakeit.Uint64()),
		// TODO: this should probably be smaller than maxswap
		MinSwap:       new(big.Int).SetUint64(gofakeit.Uint64()),
		SwapFee:       new(big.Int).SetUint64(gofakeit.Uint64()),
		MaxSwapFee:    new(big.Int).SetUint64(gofakeit.Uint64()),
		MinSwapFee:    new(big.Int).SetUint64(gofakeit.Uint64()),
		HasUnderlying: gofakeit.Bool(),
		IsUnderlying:  gofakeit.Bool(),
	},
},
}

func (p *SwapFetcherSuite) SetupTest() {
	p.TestSuite.SetupTest()

	p.db, p.eventDB, p.gqlClient, p.logIndex, p.cleanup, _, p.deployManager = testutil.NewTestEnvDB(p.GetTestContext(), p.T(), p.metrics)

	chainID := big.NewInt(1)
	p.testBackend = geth.NewEmbeddedBackendForChainID(p.GetTestContext(), p.T(), chainID)

	testERC20AInfo, _ := p.deployManager.GetERC20A(p.GetTestContext(), p.testBackend)
	testERC20BInfo, _ := p.deployManager.GetERC20B(p.GetTestContext(), p.testBackend)
	testLPTokenInfo, _ := p.deployManager.GetLPToken(p.GetTestContext(), p.testBackend)

	testSwapInfo, handle := p.deployManager.GetSwapFlashLoan(p.GetTestContext(), p.testBackend)
	handleOwner := testSwapInfo.Owner()
	fmt.Println("handleOwner", handleOwner)
	callOpts := &bind.CallOpts{}
	testOwner, err := handle.Owner(callOpts)
	ops := bind.TransactOpts{
		From:      testSwapInfo.Owner(),
		GasTipCap: big.NewInt(10000000000000),
	}
	bind.SignerFn()
	fmt.Println("testOwnertestOwnertestOwner", testOwner, err)

	test, err := handle.TransferOwnership(&ops, testOwner)
	fmt.Println("TransferOwnership", test, err)

	initTx, err := handle.Initialize(&ops, []common.Address{testERC20AInfo.Address(), testERC20BInfo.Address()}, []uint8{18, 18}, "TESTLP", "TSTLP", big.NewInt(2000), big.NewInt(0), big.NewInt(4000000), testLPTokenInfo.Address())
	Nil(p.T(), err)
	fmt.Println("SSS", initTx, err, testSwapInfo.Address(), testLPTokenInfo.Address())
	//
	//for _, token := range testTokens {
	//	token.TokenAddress = testERC20Info.Address().String()
	//	auth := p.testBackend.GetTxContext(p.GetTestContext(), deployInfo.OwnerPtr())
	//	tx, err := token.SetTokenConfig(p.bridgeConfigContract, auth)
	//	p.Require().NoError(err)
	//	p.testBackend.WaitForConfirmation(p.GetTestContext(), tx)
	//}

}

// TestSwapFetcherSuite tests the Swap Fetcher suite.
func TestSwapFetcherSuite(t *testing.T) {
	suite.Run(t, NewSwapFetcherSuite(t))
}
