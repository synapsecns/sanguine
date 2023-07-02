package submitter

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/mapmutex"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/ethergo/chain/gas"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter/config"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"math"
	"math/big"
	"reflect"
	"runtime"
	"sync"
	"time"
)

var logger = log.Logger("ethergo-submitter")

// TransactionSubmitter is the interface for submitting transactions to the chain.
type TransactionSubmitter interface {
	// Start starts the transaction submitter.
	Start(ctx context.Context) error
	// SubmitTransaction submits a transaction to the chain.
	// the transaction is not guaranteed to be executed immediately, only at some point in the future.
	// the nonce is returned, and can be used to track the status of the transaction.
	SubmitTransaction(ctx context.Context, chainID *big.Int, call ContractCallType) (nonce uint64, err error)
	// GetSubmissionStatus returns the status of a transaction and any metadata associated with it if it is complete.
	GetSubmissionStatus(ctx context.Context, chainID *big.Int, nonce uint64) (status SubmissionStatus, err error)
}

// txSubmitterImpl is the implementation of the transaction submitter.
type txSubmitterImpl struct {
	metrics metrics.Handler
	// signer is the signer for signing transactions.
	signer signer.Signer
	// nonceMux is the mutex for the nonces. It is keyed by chain.
	nonceMux mapmutex.StringerMapMutex
	// statusMux is the mutex for the status of a tx. It is keyed by tx hash.
	statusMux mapmutex.StringMapMutex
	// fetcher is used to fetch the chain client for a given chain id.
	fetcher ClientFetcher
	// db is the database for storing transactions.
	db db.Service
	// retryOnce is used to return 0 on the first call to GetRetryInterval.
	retryOnce sync.Once
	// retryNow is used to trigger a retry immediately.
	// it circumvents the retry interval.
	// to prevent memory leaks, this has a buffer of 1.
	// callers adding to this channel should not block.
	retryNow chan bool
	// config is the config for the transaction submitter.
	config config.IConfig
}

// ClientFetcher is the interface for fetching a chain client.
//
//go:generate go run github.com/vektra/mockery/v2 --name ClientFetcher --output ./mocks --case=underscore
type ClientFetcher interface {
	GetClient(ctx context.Context, chainID *big.Int) (client.EVM, error)
}

// NewTransactionSubmitter creates a new transaction submitter.
func NewTransactionSubmitter(metrics metrics.Handler, signer signer.Signer, fetcher ClientFetcher, db db.Service, config config.IConfig) TransactionSubmitter {
	return &txSubmitterImpl{
		db:        db,
		config:    config,
		metrics:   metrics,
		signer:    signer,
		fetcher:   fetcher,
		nonceMux:  mapmutex.NewStringerMapMutex(),
		statusMux: mapmutex.NewStringMapMutex(),
		retryNow:  make(chan bool, 1),
	}
}

// GetRetryInterval returns the retry interval for the transaction submitter.
func (t *txSubmitterImpl) GetRetryInterval() time.Duration {
	retryInterval := time.Second * 10
	t.retryOnce.Do(func() {
		retryInterval = time.Duration(0)
	})
	return retryInterval
}

func (t *txSubmitterImpl) Start(ctx context.Context) error {
	i := 0
	for {
		i++
		shouldExit, err := t.runSelector(ctx, i)
		if err != nil {
			logger.Warn(err)
		}
		if shouldExit {
			return nil
		}
	}
}

func (t *txSubmitterImpl) GetSubmissionStatus(ctx context.Context, chainID *big.Int, nonce uint64) (status SubmissionStatus, err error) {
	nonceStatus, err := t.db.GetNonceStatus(ctx, t.signer.Address(), chainID, nonce)
	if err != nil {
		if errors.Is(err, db.ErrNonceNotExist) {
			return submissionStatusImpl{
				state: NotFound,
			}, nil
		}

		return nil, fmt.Errorf("could not get nonce status: %w", err)
	}

	if nonceStatus == db.ReplacedOrConfirmed {
		return submissionStatusImpl{
			state: Confirming,
		}, nil
	}

	if nonceStatus == db.Confirmed {
		txs, err := t.db.GetNonceAttemptsByStatus(ctx, t.signer.Address(), chainID, nonce, db.Confirmed)
		if err != nil {
			return nil, fmt.Errorf("could not get nonce attempts by status: %w", err)
		}

		if len(txs) == 0 {
			return nil, fmt.Errorf("unexpected error: no transactions found for nonce %d", nonce)
		}

		return submissionStatusImpl{
			state:  Confirmed,
			txHash: txs[0].Hash(),
		}, nil
	}

	return submissionStatusImpl{
		state: Pending,
	}, nil
}

func (t *txSubmitterImpl) getNonce(parentCtx context.Context, chainID *big.Int, address common.Address) (_ uint64, err error) {
	ctx, span := t.metrics.Tracer().Start(parentCtx, "submitter.GetNonce", trace.WithAttributes(
		attribute.Stringer("chainID", chainID),
		attribute.Stringer("address", address),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	g, ctx := errgroup.WithContext(ctx)
	// onChainNonce is the latest nonce from eth_transactionCount. DB nonce is latest nonce from db + 1
	// locks are not built into this method or the insertion level of the db
	var onChainNonce, dbNonce uint64

	chainClient, err := t.fetcher.GetClient(ctx, chainID)
	if err != nil {
		return 0, fmt.Errorf("could not get client: %w", err)
	}

	g.Go(func() error {
		onChainNonce, err = chainClient.NonceAt(ctx, address, nil)
		if err != nil {
			return fmt.Errorf("could not get nonce from chain: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		dbNonce, err = t.db.GetNonceForChainID(ctx, address, chainID)
		if errors.Is(err, db.ErrNoNonceForChain) {
			dbNonce = 0
			return nil
		}
		if err != nil {
			return fmt.Errorf("could not get nonce from db: %w", err)
		}

		dbNonce++

		return nil
	})

	err = g.Wait()
	if err != nil {
		return 0, fmt.Errorf("could not get nonce: %w", err)
	}

	if onChainNonce > dbNonce {
		return onChainNonce, nil
	}

	return dbNonce, nil
}

func (t *txSubmitterImpl) storeTX(ctx context.Context, tx *types.Transaction, status db.Status) (err error) {
	ctx, span := t.metrics.Tracer().Start(ctx, "submitter.StoreTX", trace.WithAttributes(
		append(txToAttributes(tx), attribute.String("status", status.String()))...))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	err = t.db.PutTXS(ctx, db.TX{
		Transaction: tx,
		Status:      status,
	})
	if err != nil {
		return fmt.Errorf("could not put tx: %w", err)
	}

	return nil
}

// ContractCallType is a contract call that can be called safely.
type ContractCallType func(transactor *bind.TransactOpts) (tx *types.Transaction, err error)

// triggerProcessQueue triggers the process queue.
// will not block if the channel is full (the tx will be processed on the next retry).
func (t *txSubmitterImpl) triggerProcessQueue(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	// trigger the process queue now if we can.
	case t.retryNow <- true:
	default:
		// do nothing
		return
	}
}

// nolint: cyclop
func (t *txSubmitterImpl) SubmitTransaction(parentCtx context.Context, chainID *big.Int, call ContractCallType) (nonce uint64, err error) {
	ctx, span := t.metrics.Tracer().Start(parentCtx, "submitter.SubmitTransaction", trace.WithAttributes(
		attribute.Stringer("chainID", chainID),
		attribute.String("caller", runtime.FuncForPC(reflect.ValueOf(call).Pointer()).Name()),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// make sure we have a client for this chain.
	chainClient, err := t.fetcher.GetClient(ctx, chainID)
	if err != nil {
		return 0, fmt.Errorf("could not get client: %w", err)
	}

	// get the underlying transactor
	parentTransactor, err := t.signer.GetTransactor(ctx, core.CopyBigInt(chainID))
	if err != nil {
		return 0, fmt.Errorf("could not get transactor: %w", err)
	}

	// then we copy the transactor, this is the one we'll modify w/ no send.
	transactor := copyTransactOpts(parentTransactor)

	var locker mapmutex.Unlocker

	// this should not be modified, we need to modify this only after we set the nonce
	transactor.NoSend = true
	// we set the nonce to the max uint64 + 1. This allows the contract call to not get the nonce
	// since it's set, while allowing us to make sure the tx won't execute until we set a valid nonce.
	// this also prevents a bug in the caller from breaking our lock
	transactor.Nonce = new(big.Int).Add(new(big.Int).SetUint64(math.MaxUint64), big.NewInt(1))

	err = t.setGasPrice(ctx, chainClient, transactor, chainID, nil)
	if err != nil {
		span.AddEvent("could not set gas price", trace.WithAttributes(attribute.String("error", err.Error())))
	}
	if !t.config.GetDynamicGasEstimate(int(chainID.Uint64())) {
		transactor.GasLimit = t.config.GetGasEstimate(int(chainID.Uint64()))
	}

	transactor.Signer = func(address common.Address, transaction *types.Transaction) (_ *types.Transaction, err error) {
		locker = t.nonceMux.Lock(chainID)
		// it's important that we unlock the nonce if we fail to sign the transaction.
		// this is why we use a defer here. The second defer should only be called if the first defer is not called.
		defer func() {
			if err != nil {
				locker.Unlock()
			}
		}()

		newNonce, err := t.getNonce(ctx, chainID, address)
		if err != nil {
			return nil, fmt.Errorf("could not sign tx: %w", err)
		}

		txType := transaction.Type()
		if t.config.SupportsEIP1559(int(chainID.Uint64())) {
			txType = types.DynamicFeeTxType
		}

		transaction, err = util.CopyTX(transaction, util.WithNonce(newNonce), util.WithTxType(txType))
		if err != nil {
			return nil, fmt.Errorf("could not copy tx: %w", err)
		}

		//nolint: wrapcheck
		return parentTransactor.Signer(address, transaction)
	}
	tx, err := call(transactor)
	if err != nil {
		return 0, fmt.Errorf("could not call contract: %w", err)
	}
	defer locker.Unlock()

	// now that we've stored the tx
	err = t.storeTX(ctx, tx, db.Stored)
	if err != nil {
		return 0, fmt.Errorf("could not store transaction: %w", err)
	}

	span.AddEvent("trigger reprocess")
	t.triggerProcessQueue(ctx)

	return tx.Nonce(), nil
}

// setGasPrice sets the gas price for the transaction.
// it bumps if prevtx is set
// nolint: cyclop
// TODO: use options.
func (t *txSubmitterImpl) setGasPrice(ctx context.Context, client client.EVM,
	transactor *bind.TransactOpts, bigChainID *big.Int, prevTx *types.Transaction) (err error) {
	ctx, span := t.metrics.Tracer().Start(ctx, "submitter.setGasPrice")

	chainID := int(bigChainID.Uint64())
	maxPrice := t.config.GetMaxGasPrice(chainID)

	defer func() {
		if transactor.GasPrice != nil && maxPrice.Cmp(transactor.GasPrice) < 0 {
			transactor.GasPrice = maxPrice
		}
		if transactor.GasFeeCap != nil && maxPrice.Cmp(transactor.GasFeeCap) < 0 {
			transactor.GasFeeCap = maxPrice
		}

		metrics.EndSpanWithErr(span, err)
	}()

	// TODO: cache both of these values
	if t.config.SupportsEIP1559(int(bigChainID.Uint64())) {
		transactor.GasFeeCap = t.config.GetMaxGasPrice(chainID)

		transactor.GasTipCap, err = client.SuggestGasTipCap(ctx)
		if err != nil {
			return fmt.Errorf("could not get gas tip cap: %w", err)
		}
	} else {
		transactor.GasPrice, err = client.SuggestGasPrice(ctx)
		if err != nil {
			return fmt.Errorf("could not get gas price: %w", err)
		}
	}

	//nolint: nestif
	if prevTx != nil {
		// TODO: cache
		gasBlock, err := t.getGasBlock(ctx, client)
		if err != nil {
			span.AddEvent("could not get gas block", trace.WithAttributes(attribute.String("error", err.Error())))
		}

		// if the prev tx was greater than this one, we should bump the gas price from that point
		comparison := gas.CompareGas(prevTx, gas.OptsToComparableTx(transactor), gasBlock.BaseFee)
		if comparison > 0 {
			if prevTx.Type() == types.LegacyTxType {
				transactor.GasPrice = core.CopyBigInt(prevTx.GasPrice())
			} else {
				transactor.GasTipCap = core.CopyBigInt(prevTx.GasTipCap())
				transactor.GasFeeCap = core.CopyBigInt(prevTx.GasFeeCap())
			}
		}
		gas.BumpGasFees(transactor, t.config.GetGasBumpPercentage(chainID), gasBlock.BaseFee, maxPrice)
	}
	return nil
}

// getGasBlock gets the gas block for the given chain.
func (t *txSubmitterImpl) getGasBlock(ctx context.Context, chainClient client.EVM) (gasBlock *types.Header, err error) {
	ctx, span := t.metrics.Tracer().Start(ctx, "submitter.getGasBlock")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	err = retry.WithBackoff(ctx, func(ctx context.Context) (err error) {
		gasBlock, err = chainClient.HeaderByNumber(ctx, nil)
		if err != nil {
			return fmt.Errorf("could not get gas block: %w", err)
		}

		return nil
	}, retry.WithMin(time.Millisecond*50), retry.WithMax(time.Second*3), retry.WithMaxAttempts(4))

	if err != nil {
		return nil, fmt.Errorf("could not get gas block: %w", err)
	}

	return gasBlock, nil
}

// getGasEstimate gets the gas estimate for the given transaction.
// TODO: handle l2s w/ custom gas pricing through contracts.
func (t *txSubmitterImpl) getGasEstimate(ctx context.Context, chainClient client.EVM, chainID int, tx *types.Transaction) (gasEstimate uint64, err error) {
	if !t.config.GetDynamicGasEstimate(chainID) {
		return t.config.GetGasEstimate(chainID), nil
	}

	ctx, span := t.metrics.Tracer().Start(ctx, "submitter.getGasEstimate", trace.WithAttributes(
		attribute.Int(metrics.ChainID, chainID),
		attribute.String(metrics.TxHash, tx.Hash().String()),
	))

	defer func() {
		span.AddEvent("estimated_gas", trace.WithAttributes(attribute.Int64("gas", int64(gasEstimate))))
		metrics.EndSpanWithErr(span, err)
	}()

	// if it needs a dynamic gas estimate, we'll get it.
	if t.config.GetDynamicGasEstimate(chainID) {
		call, err := util.TxToCall(tx)
		if err != nil {
			return 0, fmt.Errorf("could not convert tx to call: %w", err)
		}

		gasEstimate, err = chainClient.EstimateGas(ctx, *call)
		if err != nil {
			span.AddEvent("could not estimate gas", trace.WithAttributes(attribute.String("error", err.Error())))
			// fallback to default
			return t.config.GetGasEstimate(chainID), nil
		}
	}

	return gasEstimate, nil
}

var _ TransactionSubmitter = &txSubmitterImpl{}
