package evm_test

import (
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	dbMocks "github.com/synapsecns/sanguine/core/db/mocks"
	"github.com/synapsecns/sanguine/core/domains/evm"
	signerMocks "github.com/synapsecns/sanguine/ethergo/signer/signer/mocks"
	chainMocks "github.com/synapsecns/synapse-node/pkg/evm/mocks"
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
