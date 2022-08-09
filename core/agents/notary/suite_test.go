package notary_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/domains"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/sanguine/core/types"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/synapse-node/pkg/chainwatcher"
	"github.com/synapsecns/synapse-node/testutils"
	"github.com/synapsecns/synapse-node/testutils/backends"
	"github.com/synapsecns/synapse-node/testutils/backends/preset"
	"math/big"
	"testing"
	"time"
)

// NotarySuite tests the updater agent.
type NotarySuite struct {
	*testutils.TestSuite
	testBackend         backends.TestBackend
	deployManager       *testutil.DeployManager
	homeContract        *home.HomeRef
	attestationContract *attestationcollector.AttestationCollectorRef
	domainClient        domains.DomainClient
	// wallet is the wallet used for the signer
	wallet wallet.Wallet
	signer signer.Signer
}

// NewNotarySuite creates a new updater suite.
func NewNotarySuite(tb testing.TB) *NotarySuite {
	tb.Helper()

	return &NotarySuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func (u *NotarySuite) SetupTest() {
	chainwatcher.PollInterval = time.Second

	u.TestSuite.SetupTest()

	u.testBackend = preset.GetRinkeby().Geth(u.GetTestContext(), u.T())
	u.deployManager = testutil.NewDeployManager(u.T())
	_, u.homeContract = u.deployManager.GetHome(u.GetTestContext(), u.testBackend)
	_, u.attestationContract = u.deployManager.GetAttestationCollector(u.GetTestContext(), u.testBackend)

	var err error
	u.domainClient, err = evm.NewEVM(u.GetTestContext(), "updater", config.DomainConfig{
		DomainID:                   uint32(u.testBackend.Config().ChainID),
		Type:                       types.EVM.String(),
		HomeAddress:                u.homeContract.Address().String(),
		AttesationCollectorAddress: u.attestationContract.Address().String(),
		RPCUrl:                     u.testBackend.RPCAddress(),
	})
	Nil(u.T(), err)

	u.wallet, err = wallet.FromRandom()
	Nil(u.T(), err)

	// fund the signer
	_, updaterManager := u.deployManager.GetUpdaterManager(u.GetTestContext(), u.testBackend)
	owner, err := updaterManager.Owner(&bind.CallOpts{Context: u.GetTestContext()})
	Nil(u.T(), err)

	u.signer = localsigner.NewSigner(u.wallet.PrivateKey())
	u.testBackend.FundAccount(u.GetTestContext(), u.signer.Address(), *big.NewInt(params.Ether))

	auth := u.testBackend.GetTxContext(u.GetTestContext(), &owner)

	// set the updater on the update manager
	tx, err := updaterManager.SetUpdater(auth.TransactOpts, u.signer.Address())
	Nil(u.T(), err)

	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)

	// set the updater on the attestation collector
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
