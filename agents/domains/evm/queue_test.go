package evm_test

import (
	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	dbMocks "github.com/synapsecns/sanguine/agents/db/mocks"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	chainMocks "github.com/synapsecns/sanguine/ethergo/chain/mocks"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	signerMocks "github.com/synapsecns/sanguine/ethergo/signer/signer/mocks"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"math/big"
)

func (t *TxQueueSuite) TestGetNonce() {
	dbMock := new(dbMocks.TxQueueDB)
	dbMock.On("GetNonceForChainID", mock.Anything, mock.Anything, mock.Anything).Return(uint64(4), nil)

	evmMock := new(chainMocks.Chain)
	evmMock.On("NonceAt", mock.Anything, mock.Anything, mock.Anything).Return(uint64(1), nil)
	evmMock.On("GetBigChainID").Return(big.NewInt(1))

	signerMock := new(signerMocks.Signer)
	signerMock.On("Address").Return(common.Address{})

	txQueue := evm.NewTxQueue(signerMock, dbMock, evmMock)

	nonce, err := txQueue.GetNonce(t.GetTestContext())
	Nil(t.T(), err)
	Equal(t.T(), nonce, uint64(5))
}

func (t *TxQueueSuite) TestGetTransactor() {
	// create a test chain
	chn := simulated.NewSimulatedBackend(t.GetTestContext(), t.T())
	manager := testutil.NewDeployManager(t.T())

	originDeployment, originHarness := manager.GetOriginHarness(t.GetTestContext(), chn)
	notaryManagerDeployment, notaryManager := manager.GetNotaryManager(t.GetTestContext(), chn)

	notaryManagerOpts := chn.GetTxContext(t.GetTestContext(), notaryManagerDeployment.OwnerPtr())

	// setup the origin on the notary contract
	setOriginTx, err := notaryManager.SetOrigin(notaryManagerOpts.TransactOpts, originHarness.Address())
	Nil(t.T(), err)
	chn.WaitForConfirmation(t.GetTestContext(), setOriginTx)

	// setup the notary on the origin contract
	originOwnerTxOpts := chn.GetTxContext(t.GetTestContext(), originDeployment.OwnerPtr())
	Nil(t.T(), err)

	setNotaryManagerTx, err := originHarness.SetNotaryManager(originOwnerTxOpts.TransactOpts, notaryManagerDeployment.Address())
	Nil(t.T(), err)
	chn.WaitForConfirmation(t.GetTestContext(), setNotaryManagerTx)

	// add the notary
	setNotaryTx, err := notaryManager.SetNotary(notaryManagerOpts.TransactOpts, notaryManagerDeployment.Owner())
	Nil(t.T(), err)
	chn.WaitForConfirmation(t.GetTestContext(), setNotaryTx)

	// create a test signer
	wllt, err := wallet.FromRandom()
	Nil(t.T(), err)

	msigner := localsigner.NewSigner(wllt.PrivateKey())
	testDB, err := sqlite.NewSqliteStore(t.GetTestContext(), filet.TmpDir(t.T(), ""))
	Nil(t.T(), err)

	testQueue := evm.NewTxQueue(msigner, testDB, chn)

	testTransactor, err := testQueue.GetTransactor(t.GetTestContext(), chn.GetBigChainID())
	Nil(t.T(), err)

	chn.FundAccount(t.GetTestContext(), msigner.Address(), *big.NewInt(params.Ether))

	encodedTips, err := types.EncodeTips(types.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)))
	Nil(t.T(), err)

	tx, err := originHarness.Dispatch(testTransactor, 1, [32]byte{}, 1, encodedTips, []byte("hello world"))
	Nil(t.T(), err)

	chn.WaitForConfirmation(t.GetTestContext(), tx)

	_, err = originHarness.Dispatch(testTransactor, 1, [32]byte{}, 1, encodedTips, []byte("hello world"))
	Nil(t.T(), err)
	chn.WaitForConfirmation(t.GetTestContext(), tx)
}
