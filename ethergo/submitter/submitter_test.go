package submitter_test

import (
	"fmt"
	"github.com/synapsecns/sanguine/ethergo/examples/contracttests"
	"github.com/synapsecns/sanguine/ethergo/examples/contracttests/counter"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core/testsuite"
	clientMocks "github.com/synapsecns/sanguine/ethergo/client/mocks"
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

	legacyChainID := s.testBackends[0].GetBigChainID()
	dynamicChainID := s.testBackends[1].GetBigChainID()

	client := new(clientMocks.EVM)
	legacyTransactor, err := signer.GetTransactor(s.GetTestContext(), legacyChainID)
	s.Require().NoError(err)

	dynamicTransactor, err := signer.GetTransactor(s.GetTestContext(), dynamicChainID)
	s.Require().NoError(err)

	maxGasPrice := big.NewInt(1000 * params.GWei)
	minGasPrice := big.NewInt(1 * params.GWei)
	cfg := &config.Config{
		Chains: map[int]config.ChainConfig{
			int(legacyChainID.Int64()): {
				MinGasPrice:     minGasPrice,
				MaxGasPrice:     maxGasPrice,
				SupportsEIP1559: false,
			},
			int(dynamicChainID.Int64()): {
				MinGasPrice:     minGasPrice,
				MaxGasPrice:     maxGasPrice,
				SupportsEIP1559: true,
			},
		},
	}
	ts := submitter.NewTestTransactionSubmitter(s.metrics, signer, s, s.store, cfg)

	resetTransactors := func() {
		legacyTransactor.GasPrice = nil
		dynamicTransactor.GasFeeCap = nil
		dynamicTransactor.GasTipCap = nil
	}

	getLegacyTx := func(gasPrice *big.Int) *types.Transaction {
		return types.NewTx(&types.LegacyTx{
			GasPrice: gasPrice,
		})
	}

	getDynamicTx := func(gasFeeCap, gasTipCap *big.Int) *types.Transaction {
		return types.NewTx(&types.DynamicFeeTx{
			GasFeeCap: gasFeeCap,
			GasTipCap: gasTipCap,
		})
	}

	assertGasValues := func(transactor *bind.TransactOpts, gasPrice, gasFeeCap, gasTipCap *big.Int) {
		s.Equal(gasPrice, transactor.GasPrice, testsuite.BigIntComparer())
		s.Equal(gasFeeCap, transactor.GasFeeCap, testsuite.BigIntComparer())
		s.Equal(gasTipCap, transactor.GasTipCap, testsuite.BigIntComparer())
	}

	s.Run("LegacyTx:FromOracle", func() {
		resetTransactors()
		gasPrice := big.NewInt(100 * params.GWei)
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(gasPrice, nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, legacyTransactor, legacyChainID, nil)
		s.Require().NoError(err)
		assertGasValues(legacyTransactor, gasPrice, nil, nil)
	})

	s.Run("DynamicTx:FromOracle", func() {
		resetTransactors()
		gasPrice := big.NewInt(120 * params.GWei)
		gasTipCap := big.NewInt(50 * params.GWei)
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(gasPrice, nil)
		client.On(testsuite.GetFunctionName(client.SuggestGasTipCap), mock.Anything).Once().Return(gasTipCap, nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, dynamicTransactor, dynamicChainID, nil)
		s.Require().NoError(err)
		assertGasValues(dynamicTransactor, nil, gasPrice, gasTipCap)
	})

	s.Run("LegacyTx:BelowMin", func() {
		resetTransactors()
		gasPrice := big.NewInt(1 * params.GWei)
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(gasPrice, nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, legacyTransactor, legacyChainID, nil)
		s.Require().NoError(err)
		assertGasValues(legacyTransactor, minGasPrice, nil, nil)
	})

	s.Run("DynamicTx:BelowMin", func() {
		resetTransactors()
		gasPrice := big.NewInt(0.5 * params.GWei)
		gasTipCap := big.NewInt(0)
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(gasPrice, nil)
		client.On(testsuite.GetFunctionName(client.SuggestGasTipCap), mock.Anything).Once().Return(gasTipCap, nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, dynamicTransactor, dynamicChainID, nil)
		s.Require().NoError(err)
		assertGasValues(dynamicTransactor, nil, minGasPrice, big.NewInt(10*params.Wei))
	})

	s.Run("LegacyTx:AboveMax", func() {
		resetTransactors()
		gasPrice := big.NewInt(10000 * params.GWei)
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(gasPrice, nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, legacyTransactor, legacyChainID, nil)
		s.NotNil(err)
	})

	s.Run("DynamicTx:AboveMax", func() {
		resetTransactors()
		gasPrice := big.NewInt(20000 * params.GWei)
		gasTipCap := big.NewInt(10000 * params.GWei)
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(gasPrice, nil)
		client.On(testsuite.GetFunctionName(client.SuggestGasTipCap), mock.Anything).Once().Return(gasTipCap, nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, dynamicTransactor, dynamicChainID, nil)
		s.NotNil(err)
	})

	s.Run("LegacyTx:SimpleBump", func() {
		resetTransactors()
		prevTx := getLegacyTx(big.NewInt(100 * params.GWei))
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(big.NewInt(50*params.GWei), nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, legacyTransactor, legacyChainID, prevTx)
		s.Require().NoError(err)
		assertGasValues(legacyTransactor, big.NewInt(110*params.GWei), nil, nil)
	})

	s.Run("DynamicTx:SimpleBump", func() {
		resetTransactors()
		prevTx := getDynamicTx(big.NewInt(150*params.GWei), big.NewInt(110*params.GWei))
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(big.NewInt(70*params.GWei), nil)
		client.On(testsuite.GetFunctionName(client.SuggestGasTipCap), mock.Anything).Once().Return(big.NewInt(60*params.GWei), nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, dynamicTransactor, dynamicChainID, prevTx)
		s.Require().NoError(err)
		assertGasValues(dynamicTransactor, nil, big.NewInt(165*params.GWei), big.NewInt(121*params.GWei))
	})

	s.Run("LegacyTx:BumpWithOracleOverride", func() {
		resetTransactors()
		prevTx := getLegacyTx(big.NewInt(100 * params.GWei))
		gasPrice := big.NewInt(200 * params.GWei)
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(gasPrice, nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, legacyTransactor, legacyChainID, prevTx)
		s.Require().NoError(err)
		assertGasValues(legacyTransactor, gasPrice, nil, nil)
	})

	s.Run("DynamicTx:BumpWithOracleOverride", func() {
		resetTransactors()
		prevTx := getDynamicTx(big.NewInt(150*params.GWei), big.NewInt(110*params.GWei))
		gasPrice := big.NewInt(200 * params.GWei)
		gasTipCap := big.NewInt(150 * params.GWei)
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(gasPrice, nil)
		client.On(testsuite.GetFunctionName(client.SuggestGasTipCap), mock.Anything).Once().Return(gasTipCap, nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, dynamicTransactor, dynamicChainID, prevTx)
		s.Require().NoError(err)
		assertGasValues(dynamicTransactor, nil, gasPrice, gasTipCap)
	})

	s.Run("LegacyTx:BumpWithPrevDynamicTx", func() {
		resetTransactors()
		prevTx := getDynamicTx(big.NewInt(100*params.GWei), big.NewInt(80*params.GWei))
		gasPrice := big.NewInt(50 * params.GWei)
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(gasPrice, nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, legacyTransactor, legacyChainID, prevTx)
		s.Require().NoError(err)
		assertGasValues(legacyTransactor, big.NewInt(110*params.GWei), nil, nil)
	})

	s.Run("DynamicTx:BumpWithPrevLegacyTx", func() {
		resetTransactors()
		prevTx := getLegacyTx(big.NewInt(100 * params.GWei))
		gasPrice := big.NewInt(50 * params.GWei)
		gasTipCap := big.NewInt(25 * params.GWei)
		client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(gasPrice, nil)
		client.On(testsuite.GetFunctionName(client.SuggestGasTipCap), mock.Anything).Once().Return(gasTipCap, nil)
		err = ts.SetGasPrice(s.GetTestContext(), client, dynamicTransactor, dynamicChainID, prevTx)
		s.Require().NoError(err)
		assertGasValues(dynamicTransactor, nil, big.NewInt(110*params.GWei), big.NewInt(110*params.GWei))
	})
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
		s.deployer, s.testBackends[0], contracttests.CounterType)

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

	txs, err := s.store.GetTXS(s.GetTestContext(), s.signer.Address(), chainID, db.WithStatuses(db.Stored))
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
	s.T().Skip("TODO: fix me")
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

	txs, err := s.store.GetAllTXAttemptByStatus(s.GetTestContext(), s.signer.Address(), tb.GetBigChainID(), db.WithStatuses(db.ReplacedOrConfirmed, db.Confirmed, db.Replaced))
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

	txs, err := s.store.GetAllTXAttemptByStatus(s.GetTestContext(), s.signer.Address(), tb.GetBigChainID(), db.WithStatuses(db.ReplacedOrConfirmed, db.Confirmed, db.Replaced))
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
