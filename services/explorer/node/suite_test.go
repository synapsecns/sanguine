package node_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/metadata"
	"github.com/synapsecns/sanguine/services/explorer/testutil/testcontracts"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	scribeMetadata "github.com/synapsecns/sanguine/services/scribe/metadata"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"

	"github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"go.uber.org/atomic"
)

// NodeSuite is the config test suite.
type NodeSuite struct {
	*testsuite.TestSuite
	db                db.ConsumerDB
	eventDB           scribedb.EventDB
	gqlClient         *client.Client
	logIndex          atomic.Int64
	cleanup           func()
	testBackends      map[uint32]backends.SimulatedTestBackend
	deployManager     *testutil.DeployManager
	testDeployManager *testcontracts.DeployManager
	// blockConfigChainID is the chain ID of the block config.
	blockConfigChainID uint32
	scribeMetrics      metrics.Handler
	explorerMetrics    metrics.Handler
}

// NewConsumerSuite creates an end-to-end test suite.
func NewConsumerSuite(tb testing.TB) *NodeSuite {
	tb.Helper()
	return &NodeSuite{
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

func (c *NodeSuite) SetupTest() {
	c.TestSuite.SetupTest()
	backends := make(map[uint32]backends.SimulatedTestBackend)
	c.blockConfigChainID = uint32(10)
	c.db, c.eventDB, c.gqlClient, c.logIndex, c.cleanup, _, c.deployManager = testutil.NewTestEnvDB(c.GetTestContext(), c.T(), c.scribeMetrics)
	backend1 := geth.NewEmbeddedBackendForChainID(c.GetTestContext(), c.T(), big.NewInt(int64(1)))
	backend2 := geth.NewEmbeddedBackendForChainID(c.GetTestContext(), c.T(), big.NewInt(int64(2)))
	backend3 := geth.NewEmbeddedBackendForChainID(c.GetTestContext(), c.T(), big.NewInt(int64(3)))
	backend4 := geth.NewEmbeddedBackendForChainID(c.GetTestContext(), c.T(), big.NewInt(int64(4))) // for the bridge config contract

	backends[uint32(1)] = backend1
	backends[uint32(2)] = backend2
	backends[uint32(3)] = backend3
	backends[c.blockConfigChainID] = backend4

	c.testBackends = backends
	c.testDeployManager = testcontracts.NewDeployManager(c.T())
}

func (c *NodeSuite) SetupSuite() {
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
	c.scribeMetrics, err = metrics.NewByType(c.GetSuiteContext(), scribeMetadata.BuildInfo(), metricsHandler)
	c.Require().Nil(err)
	c.explorerMetrics, err = metrics.NewByType(c.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	c.Require().Nil(err)
}

// TestConsumerSuite runs the integration test suite.
func TestConsumerSuite(t *testing.T) {
	suite.Run(t, NewConsumerSuite(t))
}
