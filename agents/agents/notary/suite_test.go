package notary_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/ethergo/utils/chainwatcher"
	"math/big"
	"testing"
	"time"
)

// NotarySuite tests the notary agent.
type NotarySuite struct {
	*testsuite.TestSuite
	testBackend         backends.TestBackend
	deployManager       *testutil.DeployManager
	originContract      *origin.OriginRef
	attestationContract *attestationcollector.AttestationCollectorRef
	domainClient        domains.DomainClient
	// wallet is the wallet used for the signer
	wallet wallet.Wallet
	signer signer.Signer
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
	_, u.originContract = u.deployManager.GetOrigin(u.GetTestContext(), u.testBackend)
	_, u.attestationContract = u.deployManager.GetAttestationCollector(u.GetTestContext(), u.testBackend)

	var err error
	u.domainClient, err = evm.NewEVM(u.GetTestContext(), "notary", config.DomainConfig{
		DomainID:                   uint32(u.testBackend.Config().ChainID),
		Type:                       types.EVM.String(),
		OriginAddress:              u.originContract.Address().String(),
		AttesationCollectorAddress: u.attestationContract.Address().String(),
		RPCUrl:                     u.testBackend.RPCAddress(),
	})
	Nil(u.T(), err)

	u.wallet, err = wallet.FromRandom()
	Nil(u.T(), err)

	// fund the signer
	_, notaryManager := u.deployManager.GetNotaryManager(u.GetTestContext(), u.testBackend)
	owner, err := notaryManager.Owner(&bind.CallOpts{Context: u.GetTestContext()})
	Nil(u.T(), err)

	u.signer = localsigner.NewSigner(u.wallet.PrivateKey())
	u.testBackend.FundAccount(u.GetTestContext(), u.signer.Address(), *big.NewInt(params.Ether))

	auth := u.testBackend.GetTxContext(u.GetTestContext(), &owner)

	// set the notary on the notary manager
	tx, err := notaryManager.SetNotary(auth.TransactOpts, u.signer.Address())
	Nil(u.T(), err)

	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)

	// set the notary on the attestation collector
	owner, err = u.attestationContract.Owner(&bind.CallOpts{Context: u.GetTestContext()})
	Nil(u.T(), err)

	auth = u.testBackend.GetTxContext(u.GetTestContext(), &owner)

	tx, err = u.attestationContract.AddNotary(auth.TransactOpts, u.domainClient.Config().DomainID, u.signer.Address())
	Nil(u.T(), err)

	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)
}

func TestNotarySuite(t *testing.T) {
	suite.Run(t, NewNotarySuite(t))
}
