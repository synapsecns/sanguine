package evm_test

import (
	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core/db/datastore/sql/sqlite"
	dbMocks "github.com/synapsecns/sanguine/core/db/mocks"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	signerMocks "github.com/synapsecns/sanguine/ethergo/signer/signer/mocks"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	chainMocks "github.com/synapsecns/synapse-node/pkg/evm/mocks"
	"github.com/synapsecns/synapse-node/testutils/backends/simulated"
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

	_, homeHarness := manager.GetHomeHarness(t.GetTestContext(), chn)

	// create a test signer
	wllt, err := wallet.FromRandom()
	Nil(t.T(), err)

	msigner := localsigner.NewSigner(wllt.PrivateKey())
	testDb, err := sqlite.NewSqliteStore(t.GetTestContext(), filet.TmpDir(t.T(), ""))
	Nil(t.T(), err)

	testQueue := evm.NewTxQueue(msigner, testDb, chn)

	testTransactor, err := testQueue.GetTransactor(t.GetTestContext(), chn.GetBigChainID())
	Nil(t.T(), err)

	chn.FundAccount(t.GetTestContext(), msigner.Address(), *big.NewInt(params.Ether))

	tx, err := homeHarness.Dispatch(testTransactor, 1, [32]byte{}, 1, []byte("hello world"))
	Nil(t.T(), err)

	chn.WaitForConfirmation(t.GetTestContext(), tx)

	_, err = homeHarness.Dispatch(testTransactor, 1, [32]byte{}, 1, []byte("hello world"))
	Nil(t.T(), err)
	chn.WaitForConfirmation(t.GetTestContext(), tx)
}
