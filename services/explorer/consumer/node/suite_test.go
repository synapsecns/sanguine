package node_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/testutil/testcontracts"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	"math/big"
	"testing"

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
	db                   db.ConsumerDB
	eventDB              scribedb.EventDB
	gqlClient            *client.Client
	logIndex             atomic.Int64
	cleanup              func()
	testBackends         map[uint32]backends.SimulatedTestBackend
	deployManager        *testutil.DeployManager
	testDeployManager    *testcontracts.DeployManager
	bridgeConfigContract *bridgeconfig.BridgeConfigRef
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

func (c *NodeSuite) SetupTest() {
	c.TestSuite.SetupTest()
	backends := make(map[uint32]backends.SimulatedTestBackend)
	c.db, c.eventDB, c.gqlClient, c.logIndex, c.cleanup, _, c.deployManager = testutil.NewTestEnvDB(c.GetTestContext(), c.T())
	backend1 := geth.NewEmbeddedBackendForChainID(c.GetTestContext(), c.T(), big.NewInt(int64(1)))
	backend2 := geth.NewEmbeddedBackendForChainID(c.GetTestContext(), c.T(), big.NewInt(int64(2)))
	backend3 := geth.NewEmbeddedBackendForChainID(c.GetTestContext(), c.T(), big.NewInt(int64(3)))

	backends[uint32(1)] = backend1
	backends[uint32(2)] = backend2
	backends[uint32(3)] = backend3
	c.testBackends = backends
	c.testDeployManager = testcontracts.NewDeployManager(c.T())

	var deployInfo contracts.DeployedContract
	deployInfo, c.bridgeConfigContract = c.deployManager.GetBridgeConfigV3(c.GetTestContext(), backend1)
	for _, token := range testTokens {
		auth := backend1.GetTxContext(c.GetTestContext(), deployInfo.OwnerPtr())
		tx, err := token.SetTokenConfig(c.bridgeConfigContract, auth)
		c.Require().NoError(err)
		backend1.WaitForConfirmation(c.GetTestContext(), tx)
	}
}

// TestConsumerSuite runs the integration test suite.
func TestConsumerSuite(t *testing.T) {
	suite.Run(t, NewConsumerSuite(t))
}
