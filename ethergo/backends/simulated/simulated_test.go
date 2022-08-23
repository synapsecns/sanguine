package simulated_test

import (
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"github.com/synapsecns/synapse-node/testutils/contracts"
	"github.com/synapsecns/synapse-node/testutils/contracts/deployers"
	"github.com/synapsecns/synapse-node/testutils/utils"
	"math/big"
)

// TestGetMockBakcend tests getting a mock backend.
// TODO this should test all backends in a backend agnostic way.
func (g *SimulatedSuite) TestGetSimulatedBackend() {
	be := simulated.NewSimulatedBackendWithChainID(g.GetTestContext(), g.T(), client.MaticMainnetConfig.ChainID)
	NotNil(g.T(), be)
	NotNil(g.T(), be.ChainConfig())
	Equal(g.T(), uint(be.ChainConfig().ChainID.Uint64()), be.GetChainID())
	Equal(g.T(), be.BackendName(), simulated.BackendName)
	False(g.T(), be.EnableTenderly())
	Nil(g.T(), be.VerifyContract(contracts.BridgeType, &deployers.DeployedContract{}))

	// generate a new mock address
	testAddress := utils.NewMockAddress()

	// deposit 50 eth
	funding := big.NewInt(0).Mul(big.NewInt(params.Ether), big.NewInt(50))
	be.FundAccount(g.GetTestContext(), testAddress, *funding)

	// get the balance and make sure it equals the funding amount
	balance, err := be.BalanceAt(g.GetTestContext(), testAddress, nil)
	Nil(g.T(), err)
	Equal(g.T(), balance, funding)

	// make sure suggest gas price reflects rpc behavior (adds in the base fee
	suggestedPrice, err := be.SuggestGasPrice(g.GetSuiteContext())
	Nil(g.T(), err)
	NotEqual(g.T(), suggestedPrice.Uint64(), 1)
}
