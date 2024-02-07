// Package nodeinterface_test defines the basic test suite.
package nodeinterface_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/sdks/arbitrum/contracts/nodeinterface"
	"github.com/synapsecns/sanguine/ethergo/sdks/arbitrum/internal"
)

// NodeInterfaceSuite defines the basic test suite.
type NodeInterfaceSuite struct {
	*testsuite.TestSuite
	handler metrics.Handler
}

func NewNodeInterfaceTest(tb testing.TB) *NodeInterfaceSuite {
	tb.Helper()
	return &NodeInterfaceSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestNodeInterface(t *testing.T) {
	suite.Run(t, NewNodeInterfaceTest(t))
}

func (n *NodeInterfaceSuite) SetupSuite() {
	n.TestSuite.SetupSuite()
	n.handler = metrics.NewNullHandler()
}

func (n *NodeInterfaceSuite) TestGetGasEstimateComponents() {
	arbClient, err := client.DialBackend(n.GetTestContext(), "https://arb1.arbitrum.io/rpc", n.handler)
	n.Require().NoError(err)

	nodeRef, err := nodeinterface.NewNodeInterfaceRef(internal.GetNodeInterfaceAddress(), arbClient)
	n.Require().NoError(err)

	// TODO: disable this if it's causing too many network related test fails
	err = retry.WithBackoff(n.GetTestContext(), func(ctx context.Context) error {
		gasEstimate, gasEstimateForL1, baseFee, l1BaseFeeEstimate, err := nodeRef.GetGasEstimateComponents(&bind.TransactOpts{
			Nonce:    big.NewInt(0),
			Context:  ctx,
			From:     common.BigToAddress(big.NewInt(0)),
			GasPrice: big.NewInt(params.GWei),
		}, common.BigToAddress(big.NewInt(0)), false, []byte(""))

		assert.NotZero(n.T(), gasEstimate)
		assert.NotZero(n.T(), gasEstimateForL1)
		assert.NotZero(n.T(), baseFee)
		assert.NotZero(n.T(), l1BaseFeeEstimate)

		//nolint: wrapcheck
		return err
	}, retry.WithMaxAttempts(5), retry.WithMaxAttemptTime(time.Second*30), retry.WithMin(time.Second*5))
	n.Require().NoError(err)
}
