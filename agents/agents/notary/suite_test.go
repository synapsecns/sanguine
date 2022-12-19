package notary_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

// NotarySuite tests the notary agent.
type NotarySuite struct {
	*testsuite.TestSuite
	testBackend         backends.TestBackend
	deployManager       *testutil.DeployManager
	originContract      *originharness.OriginHarnessRef
	attestationContract *attestationcollector.AttestationCollectorRef
	domainClient        domains.DomainClient
	// wallet is the wallet used for the signer
	wallet        wallet.Wallet
	signer        signer.Signer
	destinationID uint32
}

// NewNotarySuite creates a new notary suite.
func NewNotarySuite(tb testing.TB) *NotarySuite {
	tb.Helper()

	return &NotarySuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (u *NotarySuite) SetupTest() {
	chainwatcher.PollInterval = time.Second

	u.TestSuite.SetupTest()

	u.testBackend = preset.GetRinkeby().Geth(u.GetTestContext(), u.T())
	u.deployManager = testutil.NewDeployManager(u.T())
	_, u.originContract = u.deployManager.GetOriginHarness(u.GetTestContext(), u.testBackend)
	_, u.attestationContract = u.deployManager.GetAttestationCollector(u.GetTestContext(), u.testBackend)

	u.destinationID = uint32(u.testBackend.GetBigChainID().Uint64()) + 1
	var err error
	u.domainClient, err = evm.NewEVM(u.GetTestContext(), "notary", config.DomainConfig{
		DomainID:                    uint32(u.testBackend.GetBigChainID().Uint64()),
		Type:                        types.EVM.String(),
		OriginAddress:               u.originContract.Address().String(),
		AttestationCollectorAddress: u.attestationContract.Address().String(),
		RPCUrl:                      u.testBackend.RPCAddress(),
	})
	Nil(u.T(), err)

	u.wallet, err = wallet.FromRandom()
	Nil(u.T(), err)

	u.signer = localsigner.NewSigner(u.wallet.PrivateKey())
	u.testBackend.FundAccount(u.GetTestContext(), u.signer.Address(), *big.NewInt(params.Ether))

	// set the notary on the attestation collector
	owner, err := u.attestationContract.Owner(&bind.CallOpts{Context: u.GetTestContext()})
	Nil(u.T(), err)

	auth := u.testBackend.GetTxContext(u.GetTestContext(), &owner)

	tx, err := u.attestationContract.AddAgent(auth.TransactOpts, u.domainClient.Config().DomainID, u.signer.Address())
	Nil(u.T(), err)

	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)

	ownerPtr, err := u.originContract.OriginHarnessCaller.Owner(&bind.CallOpts{Context: u.GetTestContext()})
	Nil(u.T(), err)

	originOwnerAuth := u.testBackend.GetTxContext(u.GetTestContext(), &ownerPtr)
	tx, err = u.originContract.AddAgent(originOwnerAuth.TransactOpts, u.destinationID, u.signer.Address())
	Nil(u.T(), err)
	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)

	notaries, err := u.originContract.AllAgents(&bind.CallOpts{Context: u.GetTestContext()}, u.destinationID)
	Nil(u.T(), err)
	Len(u.T(), notaries, 1)
}

func TestNotarySuite(t *testing.T) {
	suite.Run(t, NewNotarySuite(t))
}
