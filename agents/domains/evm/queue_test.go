package evm_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	dbMocks "github.com/synapsecns/sanguine/agents/db/mocks"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	chainMocks "github.com/synapsecns/sanguine/ethergo/chain/mocks"
	signerMocks "github.com/synapsecns/sanguine/ethergo/signer/signer/mocks"
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
	destinationDomain := uint32(t.TestBackendDestination.GetChainID())

	auth := t.TestBackendOrigin.GetTxContext(t.GetTestContext(), nil)

	paddedRequest := big.NewInt(0)
	tx, err := t.OriginContract.SendBaseMessage(auth.TransactOpts, destinationDomain, [32]byte{}, 1, paddedRequest, []byte("hello world"))
	Nil(t.T(), err)

	t.TestBackendOrigin.WaitForConfirmation(t.GetTestContext(), tx)
}
