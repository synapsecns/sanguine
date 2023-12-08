package parser_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/scribe"
	gqlClient "github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/scribe/client"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/token"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/price"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parsers"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/static"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"github.com/synapsecns/sanguine/services/scribe/metadata"

	"math/big"
	"testing"

	"go.uber.org/atomic"
)

type ParserSuite struct {
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

// NewBackfillSuite creates a new backfill test suite.
func NewBackfillSuite(tb testing.TB) *ParserSuite {
	tb.Helper()
	return &ParserSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (p *ParserSuite) SetupSuite() {
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

func (p *ParserSuite) SetupTest() {
	p.TestSuite.SetupTest()

	p.db, p.eventDB, p.gqlClient, p.logIndex, p.cleanup, _, p.deployManager = testutil.NewTestEnvDB(p.GetTestContext(), p.T(), p.metrics)

	chainID := big.NewInt(1)
	p.testBackend = geth.NewEmbeddedBackendForChainID(p.GetTestContext(), p.T(), chainID)

	var deployInfo contracts.DeployedContract
	deployInfo, p.bridgeConfigContract = p.deployManager.GetBridgeConfigV3(p.GetTestContext(), p.testBackend)

	var testERC20Info contracts.DeployedContract
	testERC20Info, _ = p.deployManager.GetERC20A(p.GetTestContext(), p.testBackend)

	for _, token := range testTokens {
		token.TokenAddress = testERC20Info.Address().String()
		auth := p.testBackend.GetTxContext(p.GetTestContext(), deployInfo.OwnerPtr())
		tx, err := token.SetTokenConfig(p.bridgeConfigContract, auth)
		p.Require().NoError(err)
		p.testBackend.WaitForConfirmation(p.GetTestContext(), tx)
	}

	bcf, err := token.NewBridgeConfigFetcher(p.bridgeConfigContract.Address(), p.bridgeConfigContract)
	Nil(p.T(), err)
	tokenSymbolToIDs, err := parser.ParseYaml(static.GetTokenSymbolToTokenIDConfig())
	Nil(p.T(), err)
	tokenDataService, err := token.NewTokenFetcher(bcf, tokenSymbolToIDs)
	Nil(p.T(), err)
	tokenPriceService, err := price.NewPriceFetcher()
	Nil(p.T(), err)
	p.tokenPriceService = tokenPriceService
	p.tokenDataService = tokenDataService
	p.consumerFetcher = scribe.NewFetcher(p.gqlClient, p.metrics)
	arbRPC := "https://arbitrum.llamarpc.com"
	opRPC := "https://optimism.llamarpc.com"

	arbClient, err := client.DialBackend(p.GetTestContext(), arbRPC, p.metrics)
	Nil(p.T(), err)
	opClient, err := client.DialBackend(p.GetTestContext(), opRPC, p.metrics)
	Nil(p.T(), err)

	p.arbClient = arbClient
	p.opClient = opClient

	p.cctpContractArb = common.HexToAddress("0xD359bc471554504f683fbd4f6e36848612349DDF")
	p.cctpContractOp = common.HexToAddress("0x5e69c336661dde70404e3345BA61F9c01DdB4C36")

}

// TestBackfillSuite tests the backfill suite.
func TestBackfillSuite(t *testing.T) {
	suite.Run(t, NewBackfillSuite(t))
}
