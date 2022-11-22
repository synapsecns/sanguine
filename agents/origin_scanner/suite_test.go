package origin_scanner_test

import (
	"math/big"
	"testing"
	"time"

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
	"github.com/synapsecns/sanguine/ethergo/chain/chainwatcher"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

// OriginScannerSuite tests the origin scanner.
type OriginScannerSuite struct {
	*testsuite.TestSuite
	testBackend         backends.TestBackend
	deployManager       *testutil.DeployManager
	originContracts     []*origin.OriginRef
	attestationContract *attestationcollector.AttestationCollectorRef
	domainClients       []domains.DomainClient
	// wallet is the wallet used for the signer
	wallet wallet.Wallet
	signer signer.Signer
}

// NewOriginScannerSuite creates a new origin scanner suite.
func NewOriginScannerSuite(tb testing.TB) *OriginScannerSuite {
	tb.Helper()

	return &OriginScannerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (o *OriginScannerSuite) SetupTest() {
	chainwatcher.PollInterval = time.Second

	o.TestSuite.SetupTest()

	o.testBackend = preset.GetRinkeby().Geth(o.GetTestContext(), o.T())
	o.deployManager = testutil.NewDeployManager(o.T())

	o.originContracts = []*origin.OriginRef{}
	o.domainClients = []domains.DomainClient{}
	_, o.attestationContract = o.deployManager.GetAttestationCollector(o.GetTestContext(), o.testBackend)
	for i := 0; i < 10; i++ {
		_, originContract := o.deployManager.GetOrigin(o.GetTestContext(), o.testBackend)
		o.originContracts = append(o.originContracts, originContract)

		var err error
		domainClient, err := evm.NewEVM(o.GetTestContext(), "notary", config.DomainConfig{
			DomainID:                   uint32(i),
			Type:                       types.EVM.String(),
			OriginAddress:              originContract.Address().String(),
			AttesationCollectorAddress: o.attestationContract.Address().String(),
			RPCUrl:                     o.testBackend.RPCAddress(),
		})
		Nil(o.T(), err)
		o.domainClients = append(o.domainClients, domainClient)
	}

	var err error
	o.wallet, err = wallet.FromRandom()
	Nil(o.T(), err)

	// fund the signer
	/*_, notaryManager := u.deployManager.GetNotaryManager(u.GetTestContext(), u.testBackend)
	owner, err := notaryManager.Owner(&bind.CallOpts{Context: u.GetTestContext()})
	Nil(u.T(), err)*/

	o.signer = localsigner.NewSigner(o.wallet.PrivateKey())
	o.testBackend.FundAccount(o.GetTestContext(), o.signer.Address(), *big.NewInt(params.Ether))

	/*auth := u.testBackend.GetTxContext(u.GetTestContext(), &owner)

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

	o.testBackend.WaitForConfirmation(o.GetTestContext(), tx)*/
}

func TestOriginScannerSuite(t *testing.T) {
	suite.Run(t, NewOriginScannerSuite(t))
}
