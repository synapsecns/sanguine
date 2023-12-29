package submitter_test

import (
	"fmt"
	"math/big"

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
	"github.com/synapsecns/sanguine/ethergo/util"
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

	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Twice().Return(gasPrice, nil)
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
	client.On(testsuite.GetFunctionName(client.SuggestGasTipCap), mock.Anything).Once().Return(tipCap, nil)

	err = ts.SetGasPrice(s.GetTestContext(), client, transactor, chainID, nil)
	s.Require().NoError(err)

	s.Equal(tipCap, transactor.GasTipCap, testsuite.BigIntComparer())
	s.Equal(maxPrice, transactor.GasFeeCap, testsuite.BigIntComparer())

	// 4. Test with zero gas price, should return base gas price
	cfg.SetGlobalEIP1559Support(false)
	baseGasPrice := new(big.Int).SetUint64(uint64(gofakeit.Uint32()))
	cfg.SetBaseGasPrice(baseGasPrice)
	gasPrice = big.NewInt(0)
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Return(gasPrice, nil)

	err = ts.SetGasPrice(s.GetTestContext(), client, transactor, chainID, nil)
	s.Require().NoError(err)

	s.Equal(baseGasPrice, transactor.GasPrice, testsuite.BigIntComparer())

	// 5. Test with zero gas price with EIP1559, should return base gas price
	cfg.SetGlobalEIP1559Support(true)
	gasPrice = big.NewInt(0)
	client.On(testsuite.GetFunctionName(client.SuggestGasTipCap), mock.Anything).Return(gasPrice, nil)

	err = ts.SetGasPrice(s.GetTestContext(), client, transactor, chainID, nil)
	s.Require().NoError(err)

	s.Equal(baseGasPrice, transactor.GasTipCap, testsuite.BigIntComparer())

	// 5. Test with bump (TODO)
	// 6. Test with bump and max (TODO)
}

func (s *SubmitterSuite) TestGetGasBlock() {
	wall, err := wallet.FromRandom()
	s.Require().NoError(err)

	signer := localsigner.NewSigner(wall.PrivateKey())

	chainID := s.testBackends[0].GetBigChainID()
	client := new(clientMocks.EVM)

	cfg := &config.Config{}
	ts := submitter.NewTestTransactionSubmitter(s.metrics, signer, s, s.store, cfg)
	currentHeader := &types.Header{Number: big.NewInt(1)}

	// 1. Test with failed HeaderByNumber RPC call; Error is expected.
	mockErrMsg := "mock error"
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Times(5).Return(nil, fmt.Errorf(mockErrMsg))
	gasBlock, err := ts.GetGasBlock(s.GetTestContext(), client, int(chainID.Int64()))
	s.Nil(gasBlock)
	s.NotNil(err)

	// 2. Test with successful HeaderByNumber RPC call.
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(currentHeader, nil)
	gasBlock, err = ts.GetGasBlock(s.GetTestContext(), client, int(chainID.Int64()))
	s.Require().NoError(err)
	s.Equal(gasBlock.Number.String(), currentHeader.Number.String())

	// 3. Test with failed HeaderByNumber RPC call; the cached value should be used.
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Times(5).Return(nil, fmt.Errorf(mockErrMsg))
	gasBlock, err = ts.GetGasBlock(s.GetTestContext(), client, int(chainID.Int64()))
	s.Require().NoError(err)
	s.Equal(gasBlock.Number.String(), currentHeader.Number.String())
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

	go func() {
		// now we'll start a new submitter with a new signer and submit the tx
		err = ts.Start(s.GetTestContext())
		s.Require().NoError(err)
	}()

	// wait for the tx to be submitted
	s.Eventually(func() bool {
		currentCounter, err = cntr.GetCount(&bind.CallOpts{
			Context: s.GetTestContext(),
		})
		s.Require().NoError(err)

		return currentCounter.Uint64() > ogCounter.Uint64()
	})
}

func (s *SubmitterSuite) TestCheckAndSetConfirmation() {
	cfg := &config.Config{}
	ts := submitter.NewTestTransactionSubmitter(s.metrics, s.signer, s, s.store, cfg)

	tb := s.testBackends[0]
	confirmedTx := ethMocks.MockTx(s.GetTestContext(), s.T(), tb, s.localAccount, types.LegacyTxType)
	allTxes := []db.TX{{
		Transaction: confirmedTx,
		Status:      db.Pending,
	}}

	chainClient, err := s.GetClient(s.GetTestContext(), tb.GetBigChainID())
	s.Require().NoError(err)

	const duplicateCount = 15
	for i := 0; i < duplicateCount; i++ {
		copiedTX, err := util.CopyTX(confirmedTx, util.WithGasPrice(big.NewInt(int64(i))))
		s.Require().NoError(err)

		transactor, err := s.signer.GetTransactor(s.GetTestContext(), tb.GetBigChainID())
		s.Require().NoError(err)

		copiedTX, err = transactor.Signer(s.signer.Address(), copiedTX)
		s.Require().NoError(err)

		allTxes = append(allTxes, db.TX{
			Transaction: copiedTX,
			Status:      db.ReplacedOrConfirmed,
		})
	}

	err = ts.CheckAndSetConfirmation(s.GetTestContext(), chainClient, allTxes)
	s.Require().NoError(err)

	txs, err := s.store.GetAllTXAttemptByStatus(s.GetTestContext(), s.signer.Address(), tb.GetBigChainID(), db.ReplacedOrConfirmed, db.Confirmed, db.Replaced)
	s.Require().NoError(err)

	var replacedCount int
	for _, tx := range txs {
		//nolint: exhaustive
		switch tx.Status {
		case db.Replaced:
			replacedCount++
		case db.Confirmed:
			s.Require().Equal(tx.Hash(), confirmedTx.Hash())
			// make sure submission status is congruent
			status, err := ts.GetSubmissionStatus(s.GetTestContext(), tb.GetBigChainID(), tx.Nonce())
			s.Require().NoError(err)
			s.Require().Equal(submitter.Confirmed, status.State())
			s.Require().Equal(confirmedTx.Hash(), status.TxHash())

		default:
			s.Failf("unexpected status: %s", tx.Status.String())
		}
	}

	s.Require().Equal(duplicateCount, replacedCount)
}

func (s *SubmitterSuite) TestCheckAndSetConfirmationSingleTx() {
	cfg := &config.Config{}
	ts := submitter.NewTestTransactionSubmitter(s.metrics, s.signer, s, s.store, cfg)

	tb := s.testBackends[0]
	confirmedTx := ethMocks.MockTx(s.GetTestContext(), s.T(), tb, s.localAccount, types.LegacyTxType)
	allTxes := []db.TX{{
		Transaction: confirmedTx,
		Status:      db.Pending,
	}}

	chainClient, err := s.GetClient(s.GetTestContext(), tb.GetBigChainID())
	s.Require().NoError(err)

	err = ts.CheckAndSetConfirmation(s.GetTestContext(), chainClient, allTxes)
	s.Require().NoError(err)

	txs, err := s.store.GetAllTXAttemptByStatus(s.GetTestContext(), s.signer.Address(), tb.GetBigChainID(), db.ReplacedOrConfirmed, db.Confirmed, db.Replaced)
	s.Require().NoError(err)

	for _, tx := range txs {
		//nolint: exhaustive
		switch tx.Status {
		case db.Confirmed:
			s.Require().Equal(tx.Hash(), confirmedTx.Hash())
		default:
			s.Failf("unexpected status: %s", tx.Status.String())
		}
	}
}
