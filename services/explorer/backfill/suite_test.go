package backfill_test

import (
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/scribe"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/scribe/client"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/metadata"

	"math/big"
	"testing"

	"go.uber.org/atomic"
)

type BackfillSuite struct {
	*testsuite.TestSuite
	db                   db.ConsumerDB
	eventDB              scribedb.EventDB
	gqlClient            *client.Client
	logIndex             atomic.Int64
	cleanup              func()
	testBackend          backends.SimulatedTestBackend
	deployManager        *testutil.DeployManager
	bridgeConfigContract *bridgeconfig.BridgeConfigRef
	consumerFetcher      scribe.IScribeFetcher
	metrics              metrics.Handler
}

// NewBackfillSuite creates a new backfill test suite.
func NewBackfillSuite(tb testing.TB) *BackfillSuite {
	tb.Helper()
	return &BackfillSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		logIndex:  atomic.Int64{},
	}
}

func (b *BackfillSuite) SetupSuite() {
	b.TestSuite.SetupSuite()
	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(b.GetSuiteContext(), b.T())
		metricsHandler = metrics.Jaeger
	}

	var err error
	b.metrics, err = metrics.NewByType(b.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	b.Require().Nil(err)
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

func (b *BackfillSuite) SetupTest() {
	b.TestSuite.SetupTest()

	b.db, b.eventDB, b.gqlClient, b.logIndex, b.cleanup, _, b.deployManager = testutil.NewTestEnvDB(b.GetTestContext(), b.T(), b.metrics)

	chainID := big.NewInt(1)
	b.testBackend = geth.NewEmbeddedBackendForChainID(b.GetTestContext(), b.T(), chainID)

	b.consumerFetcher = scribe.NewFetcher(b.gqlClient, b.metrics)
	var deployInfo contracts.DeployedContract
	deployInfo, b.bridgeConfigContract = b.deployManager.GetBridgeConfigV3(b.GetTestContext(), b.testBackend)

	var testERC20Info contracts.DeployedContract
	testERC20Info, _ = b.deployManager.GetERC20(b.GetTestContext(), b.testBackend)

	for _, token := range testTokens {
		token.TokenAddress = testERC20Info.Address().String()
		auth := b.testBackend.GetTxContext(b.GetTestContext(), deployInfo.OwnerPtr())
		tx, err := token.SetTokenConfig(b.bridgeConfigContract, auth)
		b.Require().NoError(err)
		b.testBackend.WaitForConfirmation(b.GetTestContext(), tx)
	}
}

// TestBackfillSuite tests the backfill suite.
func TestBackfillSuite(t *testing.T) {
	suite.Run(t, NewBackfillSuite(t))
}
