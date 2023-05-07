package origin_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

// OriginSuite is the origin test suite.
type OriginSuite struct {
	*testsuite.TestSuite
	originContract *originharness.OriginHarnessRef
	testBackend    backends.SimulatedTestBackend
	notarySigner   *localsigner.Signer
	destinationID  uint32
}

// NewOriginSuite creates a end-to-end test suite.
func NewOriginSuite(tb testing.TB) *OriginSuite {
	tb.Helper()
	return &OriginSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (h *OriginSuite) SetupTest() {
	h.TestSuite.SetupTest()

	deployManager := testutil.NewDeployManager(h.T())

	h.testBackend = simulated.NewSimulatedBackendWithChainID(h.GetTestContext(), h.T(), big.NewInt(2))
	_, h.originContract = deployManager.GetOriginHarness(h.GetTestContext(), h.testBackend)

	h.destinationID = uint32(453)

	wllt, err := wallet.FromRandom()
	Nil(h.T(), err)

	h.testBackend.FundAccount(h.GetTestContext(), wllt.Address(), *big.NewInt(params.Ether))

	h.notarySigner = localsigner.NewSigner(wllt.PrivateKey())

	// ownerPtr, err := h.originContract.OriginHarnessCaller.Owner(&bind.CallOpts{Context: h.GetTestContext()})
	// Nil(h.T(), err)

	// TODO (joeallen): FIX ME
	// originOwnerAuth := h.testBackend.GetTxContext(h.GetTestContext(), &ownerPtr)
	// tx, err := h.originContract.AddAgent(originOwnerAuth.TransactOpts, h.destinationID, h.notarySigner.Address())
	// Nil(h.T(), err)
	// h.testBackend.WaitForConfirmation(h.GetTestContext(), tx)
}

// TestOriginSuite runs the integration test suite.
func TestOriginSuite(t *testing.T) {
	// TODO (joeallen): FIX ME
	// t.Skip()
	suite.Run(t, NewOriginSuite(t))
}
