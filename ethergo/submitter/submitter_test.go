package submitter_test

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core/testsuite"
	clientMocks "github.com/synapsecns/sanguine/ethergo/client/mocks"
	"github.com/synapsecns/sanguine/ethergo/example"
	"github.com/synapsecns/sanguine/ethergo/example/counter"
	"github.com/synapsecns/sanguine/ethergo/manager"
	ethMocks "github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/ethergo/submitter/config"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	dbMocks "github.com/synapsecns/sanguine/ethergo/submitter/db/mocks"
	submitterMocks "github.com/synapsecns/sanguine/ethergo/submitter/mocks"
	"math/big"
)

func (s *SubmitterSuite) TestSetGasPrice() {
	wall, err := wallet.FromRandom()
	s.Require().NoError(err)

	signer := localsigner.NewSigner(wall.PrivateKey())

	chainID := s.testBackends[0].GetBigChainID()
	client := new(clientMocks.EVM)

	transactor, err := signer.GetTransactor(s.GetTestContext(), chainID)
	s.Require().NoError(err)

	cfg := &config.Config{}
	ts := submitter.NewTestTransactionSubmitter(s.metrics, signer, s, s.store, cfg)

	// 1. Test with gas price set, but not one that exceeds max (not eip-1559)
	gasPrice := new(big.Int).SetUint64(gofakeit.Uint64())
	maxPrice := new(big.Int).Add(gasPrice, new(big.Int).SetUint64(1))
	cfg.SetGlobalMaxGasPrice(maxPrice)

	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Return(gasPrice, nil)
	err = ts.SetGasPrice(s.GetTestContext(), client, transactor, chainID, nil)
	s.Require().NoError(err)

	s.Equal(gasPrice, transactor.GasPrice, testsuite.BigIntComparer())

	// 2. Test with gas price set, but one that exceeds max, should return max (not eip-1559)
	maxPrice = new(big.Int).Sub(gasPrice, new(big.Int).SetUint64(1))
	cfg.SetGlobalMaxGasPrice(maxPrice)

	err = ts.SetGasPrice(s.GetTestContext(), client, transactor, chainID, nil)
	s.Require().NoError(err)
	s.Equal(maxPrice, transactor.GasPrice, testsuite.BigIntComparer())

	// 3. Test with gas price set, but one that exceeds max, should return max (legacy tx)
	cfg.SetGlobalEIP1559Support(true)
	tipCap := new(big.Int).SetUint64(uint64(gofakeit.Uint32()))
	client.On(testsuite.GetFunctionName(client.SuggestGasTipCap), mock.Anything).Return(tipCap, nil)

	err = ts.SetGasPrice(s.GetTestContext(), client, transactor, chainID, nil)
	s.Require().NoError(err)

	s.Equal(tipCap, transactor.GasTipCap, testsuite.BigIntComparer())
	s.Equal(maxPrice, transactor.GasFeeCap, testsuite.BigIntComparer())

	// 4. Test with bump (TODO)
	// 5. Test with bump and max (TODO)
}

func (s *SubmitterSuite) TestGetNonce() {
	chainID := s.testBackends[0].GetBigChainID()

	chainMock := new(clientMocks.EVM)
	clientFetcherMock := new(submitterMocks.ClientFetcher)
	dbMock := new(dbMocks.Service)

	clientFetcherMock.On(testsuite.GetFunctionName(clientFetcherMock.GetClient), mock.Anything, mock.Anything).Return(chainMock, nil)

	cfg := &config.Config{}
	ts := submitter.NewTestTransactionSubmitter(s.metrics, s.signer, clientFetcherMock, dbMock, cfg)
	testAddress := ethMocks.MockAddress()

	// 1. Test with db nonce > on chain nonce. Should return db nonce + 1
	dbMock.On(testsuite.GetFunctionName(dbMock.GetNonceForChainID), mock.Anything, mock.Anything, mock.Anything).Once().Return(uint64(4), nil)
	chainMock.On(testsuite.GetFunctionName(chainMock.NonceAt), mock.Anything, mock.Anything, mock.Anything).Once().Return(uint64(2), nil)

	nonce, err := ts.GetNonce(s.GetTestContext(), chainID, testAddress)
	s.Require().NoError(err)
	s.Equal(uint64(5), nonce)

	// 2. Test with chain nonce > db nonce. Should return db nonce + 1
	dbMock.On(testsuite.GetFunctionName(dbMock.GetNonceForChainID), mock.Anything, mock.Anything, mock.Anything).Once().Return(uint64(2), nil)
	chainMock.On(testsuite.GetFunctionName(chainMock.NonceAt), mock.Anything, mock.Anything, mock.Anything).Once().Return(uint64(4), nil)

	nonce, err = ts.GetNonce(s.GetTestContext(), chainID, testAddress)
	s.Require().NoError(err)
	s.Equal(uint64(4), nonce)
}

func (s *SubmitterSuite) TestSubmitTransaction() {
	_, cntr := manager.GetContract[*counter.CounterRef](s.GetTestContext(), s.T(),
		s.deployer, s.testBackends[0], example.CounterType)

	cfg := &config.Config{}
	chainID := s.testBackends[0].GetBigChainID()

	ogCounter, err := cntr.GetCount(&bind.CallOpts{
		Context: s.GetTestContext(),
	})
	s.Require().NoError(err)

	ts := submitter.NewTestTransactionSubmitter(s.metrics, s.signer, s, s.store, cfg)
	_, err = ts.SubmitTransaction(s.GetTestContext(), chainID, func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = cntr.IncrementCounter(transactor)
		if err != nil {
			return nil, fmt.Errorf("failed to increment counter: %w", err)
		}

		return tx, nil
	})
	s.Require().NoError(err)

	currentCounter, err := cntr.GetCount(&bind.CallOpts{
		Context: s.GetTestContext(),
	})
	s.Require().NoError(err)

	// make sure the tx wasn't submitted
	s.Equal(ogCounter.Uint64(), currentCounter.Uint64())

	txs, err := s.store.GetTXS(s.GetTestContext(), s.signer.Address(), chainID, db.Stored)
	s.Require().NoError(err)

	s.Require().NotNil(txs[0])
}
