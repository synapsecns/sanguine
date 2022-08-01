package notary_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/contracts/xappconfig"
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

// UpdaterSuite tests the updater agent.
type UpdaterSuite struct {
	*testutils.TestSuite
	testBackend         backends.TestBackend
	deployManager       *testutil.DeployManager
	xappConfig          *xappconfig.XAppConfigRef
	homeContract        *home.HomeRef
	attestationContract *attestationcollector.AttestationCollectorRef
	domainClient        domains.DomainClient
	// wallet is the wallet used for the signer
	wallet wallet.Wallet
	signer signer.Signer
}

// NewUpdaterSuite creates a new updater suite.
func NewUpdaterSuite(tb testing.TB) *UpdaterSuite {
	tb.Helper()

	return &UpdaterSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func (u *UpdaterSuite) SetupTest() {
	chainwatcher.PollInterval = time.Second

	u.TestSuite.SetupTest()

	u.testBackend = preset.GetRinkeby().Geth(u.GetTestContext(), u.T())
	u.deployManager = testutil.NewDeployManager(u.T())
	_, u.xappConfig = u.deployManager.GetXAppConfig(u.GetTestContext(), u.testBackend)
	_, u.homeContract = u.deployManager.GetHome(u.GetTestContext(), u.testBackend)
	_, u.attestationContract = u.deployManager.GetAttestationCollector(u.GetTestContext(), u.testBackend)

	var err error
	u.domainClient, err = evm.NewEVM(u.GetTestContext(), "updater", config.DomainConfig{
		DomainID:                   1,
		Type:                       types.EVM.String(),
		XAppConfigAddress:          u.xappConfig.Address().String(),
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

	transactOpts := u.testBackend.GetTxContext(u.GetTestContext(), &owner)

	// set the updater
	tx, err := updaterManager.SetUpdater(transactOpts.TransactOpts, u.signer.Address())
	Nil(u.T(), err)

	u.testBackend.WaitForConfirmation(u.GetTestContext(), tx)
}

func TestUpdaterSuite(t *testing.T) {
	suite.Run(t, NewUpdaterSuite(t))
}
