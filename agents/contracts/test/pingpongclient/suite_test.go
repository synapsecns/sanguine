package pingpongclient_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

// PingPongClientSuite is the ping pong test client test suite.
type PingPongClientSuite struct {
	*testsuite.TestSuite
	originContract         *originharness.OriginHarnessRef
	destinationContract    *destinationharness.DestinationHarnessRef
	pingPongClientContract *pingpongclient.PingPongClientRef
	pingPongClientMetadata contracts.DeployedContract
	testBackend            backends.SimulatedTestBackend
	notarySigner           *localsigner.Signer
	destinationID          uint32
}

// NewPingPongClientSuite creates a end-to-end test suite.
func NewPingPongClientSuite(tb testing.TB) *PingPongClientSuite {
	tb.Helper()
	return &PingPongClientSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (h *PingPongClientSuite) SetupTest() {
	h.TestSuite.SetupTest()

	deployManager := testutil.NewDeployManager(h.T())

	h.testBackend = simulated.NewSimulatedBackendWithChainID(h.GetTestContext(), h.T(), big.NewInt(2))
	_, h.originContract = deployManager.GetOriginHarness(h.GetTestContext(), h.testBackend)
	_, h.destinationContract = deployManager.GetDestinationHarness(h.GetTestContext(), h.testBackend)
	h.pingPongClientMetadata, h.pingPongClientContract = deployManager.GetPingPongClient(h.GetTestContext(), h.testBackend)

	h.destinationID = uint32(453)

	wllt, err := wallet.FromRandom()
	Nil(h.T(), err)

	h.testBackend.FundAccount(h.GetTestContext(), wllt.Address(), *big.NewInt(params.Ether))

	h.notarySigner = localsigner.NewSigner(wllt.PrivateKey())

	// TODO (joeallen): FIX ME
	// ownerPtr, err := h.originContract.OriginHarnessCaller.Owner(&bind.CallOpts{Context: h.GetTestContext()})
	// Nil(h.T(), err)

	// TODO (joeallen): FIX ME
	// originOwnerAuth := h.testBackend.GetTxContext(h.GetTestContext(), &ownerPtr)
	// tx, err := h.originContract.AddAgent(originOwnerAuth.TransactOpts, h.destinationID, h.notarySigner.Address())
	// Nil(h.T(), err)
	// h.testBackend.WaitForConfirmation(h.GetTestContext(), tx)
}

// TestPingPongClientSuite runs the integration test suite.
func TestPingPongClientSuite(t *testing.T) {
	// TODO (joeallen): FIX ME
	t.Skip()
	suite.Run(t, NewPingPongClientSuite(t))
}
