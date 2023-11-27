package consumer_test

import (
	"fmt"
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
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db"

	"math/big"
	"testing"

	"github.com/synapsecns/sanguine/services/explorer/testutil"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"go.uber.org/atomic"
)

// ConsumerSuite is the config test suite.
type ConsumerSuite struct {
	*testsuite.TestSuite
	db                   db.ConsumerDB
	eventDB              scribedb.EventDB
	gqlClient            *client.Client
	logIndex             atomic.Int64
	cleanup              func()
	testBackend          backends.SimulatedTestBackend
	deployManager        *testutil.DeployManager
	bridgeConfigContract *bridgeconfig.BridgeConfigRef
	scribeMetrics        metrics.Handler
}

// NewConsumerSuite creates an end-to-end test suite.
func NewConsumerSuite(tb testing.TB) *ConsumerSuite {
	tb.Helper()
	return &ConsumerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
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
		ChainId:       big.NewInt(int64(gofakeit.Uint32())),
		TokenAddress:  mocks.MockAddress().String(),
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

func (c *ConsumerSuite) SetupTest() {
	c.TestSuite.SetupTest()

	c.db, c.eventDB, c.gqlClient, c.logIndex, c.cleanup, c.testBackend, c.deployManager = testutil.NewTestEnvDB(c.GetTestContext(), c.T(), c.scribeMetrics)

	var deployInfo contracts.DeployedContract
	deployInfo, c.bridgeConfigContract = c.deployManager.GetBridgeConfigV3(c.GetTestContext(), c.testBackend)
	for _, token := range testTokens {
		auth := c.testBackend.GetTxContext(c.GetTestContext(), deployInfo.OwnerPtr())
		tx, err := token.SetTokenConfig(c.bridgeConfigContract, auth)
		c.Require().NoError(err)
		c.testBackend.WaitForConfirmation(c.GetTestContext(), tx)
	}
}

func (c *ConsumerSuite) SetupSuite() {
	c.TestSuite.SetupSuite()

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(c.GetSuiteContext(), c.T())
		metricsHandler = metrics.Jaeger
	}
	var err error
	c.scribeMetrics, err = metrics.NewByType(c.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	c.Require().Nil(err)
}

// TestConsumerSuite runs the integration test suite.
func TestConsumerSuite(t *testing.T) {
	suite.Run(t, NewConsumerSuite(t))
}
