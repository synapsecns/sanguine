package bridgeconfig_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"

	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"math/big"
	"testing"
)

// ConfigSuite is the config test suite.
type ConfigSuite struct {
	*testsuite.TestSuite
	testBackend          backends.SimulatedTestBackend
	deployManager        *testutil.DeployManager
	bridgeConfigContract *bridgeconfig.BridgeConfigRef
}

// NewConfigSuite creates a end-to-end test suite.
func NewConfigSuite(tb testing.TB) *ConfigSuite {
	tb.Helper()
	return &ConfigSuite{
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

func (c *ConfigSuite) SetupTest() {
	c.TestSuite.SetupTest()

	c.testBackend = simulated.NewSimulatedBackend(c.GetTestContext(), c.T())
	c.deployManager = testutil.NewDeployManager(c.T())

	var deployInfo contracts.DeployedContract
	deployInfo, c.bridgeConfigContract = c.deployManager.GetBridgeConfigV3(c.GetTestContext(), c.testBackend)

	for _, token := range testTokens {
		auth := c.testBackend.GetTxContext(c.GetTestContext(), deployInfo.OwnerPtr())
		tx, err := token.SetTokenConfig(c.bridgeConfigContract, auth)
		Nil(c.T(), err)

		c.testBackend.WaitForConfirmation(c.GetTestContext(), tx)
	}
}

// TestConfigSuite runs the integration test suite.
func TestConfigSuite(t *testing.T) {
	suite.Run(t, NewConfigSuite(t))
}
