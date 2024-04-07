package submitter

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/puzpuzpuz/xsync/v2"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
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
	// lastGasBlockCache is used to cache the last gas block for a given chain. A new block should still be fetched, if possible.
	lastGasBlockCache *xsync.MapOf[int, *types.Header]
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
		db:                db,
		config:            config,
		metrics:           metrics,
		signer:            signer,
		fetcher:           fetcher,
		nonceMux:          mapmutex.NewStringerMapMutex(),
		statusMux:         mapmutex.NewStringMapMutex(),
		retryNow:          make(chan bool, 1),
		lastGasBlockCache: xsync.NewIntegerMapOf[int, *types.Header](),
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
			logger.Warn("exiting transaction submitter")
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
	// onChainNonce is the latest nonce from eth_transactionCount. db nonce is latest nonce from db + 1
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

func (t *txSubmitterImpl) storeTX(ctx context.Context, tx *types.Transaction, status db.Status, UUID string) (err error) {
	ctx, span := t.metrics.Tracer().Start(ctx, "submitter.StoreTX", trace.WithAttributes(
		append(txToAttributes(tx, UUID), attribute.String("status", status.String()))...))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	err = t.db.PutTXS(ctx, db.TX{
		UUID:        UUID,
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

		txType := t.txTypeForChain(chainID)

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
	err = t.storeTX(ctx, tx, db.Stored, uuid.New().String())
	if err != nil {
		return 0, fmt.Errorf("could not store transaction: %w", err)
	}

	span.AddEvent("trigger reprocess")
	t.triggerProcessQueue(ctx)

	return tx.Nonce(), nil
}

func (t *txSubmitterImpl) txTypeForChain(chainID *big.Int) (txType uint8) {
	if t.config.SupportsEIP1559(int(chainID.Uint64())) {
		txType = types.DynamicFeeTxType
	} else {
		txType = types.LegacyTxType
	}
	return txType
}

// setGasPrice sets the gas price for the transaction.
// If a prevTx is specified, a bump will be attempted; otherwise values will be
// set from the gas oracle.
// If gas values exceed the configured max, an error will be returned.
func (t *txSubmitterImpl) setGasPrice(ctx context.Context, client client.EVM,
	transactor *bind.TransactOpts, bigChainID *big.Int, prevTx *types.Transaction) (err error) {
	ctx, span := t.metrics.Tracer().Start(ctx, "submitter.setGasPrice")

	chainID := int(bigChainID.Uint64())
	useDynamic := t.config.SupportsEIP1559(chainID)

	defer func() {
		span.SetAttributes(
			attribute.Int(metrics.ChainID, chainID),
			attribute.Bool("use_dynamic", useDynamic),
			attribute.String("gas_price", bigPtrToString(transactor.GasPrice)),
			attribute.String("gas_fee_cap", bigPtrToString(transactor.GasFeeCap)),
			attribute.String("gas_tip_cap", bigPtrToString(transactor.GasTipCap)),
		)
		metrics.EndSpanWithErr(span, err)
	}()

	t.bumpGasFromPrevTx(ctx, transactor, prevTx, chainID, useDynamic)
	t.applyGasFloor(ctx, transactor, chainID, useDynamic)

	err = t.applyGasFromOracle(ctx, transactor, client, useDynamic)
	if err != nil {
		return fmt.Errorf("could not populate gas from oracle: %w", err)
	}

	err = t.applyGasCeil(ctx, transactor, chainID, useDynamic)
	if err != nil {
		return fmt.Errorf("could not apply gas ceil: %w", err)
	}
	return nil
}

// bumpGasFromPrevTx populates the gas fields from the previous transaction and bumps
// the appropriate values corresponding to the configured GasBumpPercentage.
// Note that in the event of a tx type mismatch, gasFeeCap is copied to gasPrice,
// and gasPrice is copied to both gasFeeCap and gasTipCap in the opposite scenario.
//
//nolint:nestif
func (t *txSubmitterImpl) bumpGasFromPrevTx(ctx context.Context, transactor *bind.TransactOpts, prevTx *types.Transaction, chainID int, currentDynamic bool) {
	if prevTx == nil {
		return
	}

	_, span := t.metrics.Tracer().Start(ctx, "submitter.bumpGasFromPrevTx")

	defer func() {
		span.SetAttributes(
			attribute.String("gas_price", bigPtrToString(transactor.GasPrice)),
			attribute.String("gas_fee_cap", bigPtrToString(transactor.GasFeeCap)),
			attribute.String("gas_tip_cap", bigPtrToString(transactor.GasTipCap)),
		)
		metrics.EndSpan(span)
	}()

	prevDynamic := prevTx.Type() == types.DynamicFeeTxType
	bumpPct := t.config.GetGasBumpPercentage(chainID)
	if currentDynamic {
		if prevDynamic {
			transactor.GasFeeCap = gas.BumpByPercent(core.CopyBigInt(prevTx.GasFeeCap()), bumpPct)
			transactor.GasTipCap = gas.BumpByPercent(core.CopyBigInt(prevTx.GasTipCap()), bumpPct)
		} else {
			transactor.GasFeeCap = gas.BumpByPercent(core.CopyBigInt(prevTx.GasPrice()), bumpPct)
			transactor.GasTipCap = gas.BumpByPercent(core.CopyBigInt(prevTx.GasPrice()), bumpPct)
		}
	} else {
		if prevDynamic {
			transactor.GasPrice = gas.BumpByPercent(core.CopyBigInt(prevTx.GasFeeCap()), bumpPct)
		} else {
			transactor.GasPrice = gas.BumpByPercent(core.CopyBigInt(prevTx.GasPrice()), bumpPct)
		}
	}
}

var minTipCap = big.NewInt(10 * params.Wei)

// applyGasFloor applies the min gas price from the config if values are unset.
//
//nolint:cyclop,nestif
func (t *txSubmitterImpl) applyGasFloor(ctx context.Context, transactor *bind.TransactOpts, chainID int, useDynamic bool) {
	_, span := t.metrics.Tracer().Start(ctx, "submitter.applyGasFloor")

	defer func() {
		span.SetAttributes(
			attribute.String("gas_price", bigPtrToString(transactor.GasPrice)),
			attribute.String("gas_fee_cap", bigPtrToString(transactor.GasFeeCap)),
			attribute.String("gas_tip_cap", bigPtrToString(transactor.GasTipCap)),
		)
		metrics.EndSpan(span)
	}()

	gasFloor := t.config.GetMinGasPrice(chainID)
	if useDynamic {
		if transactor.GasFeeCap == nil || transactor.GasFeeCap.Cmp(gasFloor) < 0 {
			transactor.GasFeeCap = gasFloor
		}
		if transactor.GasTipCap == nil || transactor.GasTipCap.Cmp(minTipCap) < 0 {
			transactor.GasTipCap = minTipCap
		}
	} else if transactor.GasPrice == nil || transactor.GasPrice.Cmp(gasFloor) < 0 {
		transactor.GasPrice = gasFloor
	}
}

// applyGasFromOracle fetches gas values from a RPC endpoint and attempts to set them.
// If values are already specified, they will be overridden if the oracle values are higher.
func (t *txSubmitterImpl) applyGasFromOracle(ctx context.Context, transactor *bind.TransactOpts, client client.EVM, useDynamic bool) (err error) {
	ctx, span := t.metrics.Tracer().Start(ctx, "submitter.applyGasFromOracle")

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	if useDynamic {
		suggestedGasFeeCap, err := client.SuggestGasPrice(ctx)
		if err != nil {
			return fmt.Errorf("could not get gas fee cap: %w", err)
		}
		transactor.GasFeeCap = maxOfBig(transactor.GasFeeCap, suggestedGasFeeCap)
		suggestedGasTipCap, err := client.SuggestGasTipCap(ctx)
		if err != nil {
			return fmt.Errorf("could not get gas tip cap: %w", err)
		}
		transactor.GasTipCap = maxOfBig(transactor.GasTipCap, suggestedGasTipCap)
		span.SetAttributes(
			attribute.String("suggested_gas_fee_cap", bigPtrToString(suggestedGasFeeCap)),
			attribute.String("suggested_gas_tip_cap", bigPtrToString(suggestedGasTipCap)),
			attribute.String("gas_fee_cap", bigPtrToString(transactor.GasFeeCap)),
			attribute.String("gas_tip_cap", bigPtrToString(transactor.GasTipCap)),
		)
	} else {
		suggestedGasPrice, err := client.SuggestGasPrice(ctx)
		if err != nil {
			return fmt.Errorf("could not get gas price: %w", err)
		}
		transactor.GasPrice = maxOfBig(transactor.GasPrice, suggestedGasPrice)
		span.SetAttributes(
			attribute.String("suggested_gas_price", bigPtrToString(suggestedGasPrice)),
			attribute.String("gas_price", bigPtrToString(transactor.GasPrice)),
		)
	}
	return nil
}

// applyGasCeil evaluates current gas values versus the configured maximum, and
// returns an error if they exceed the maximum.
func (t *txSubmitterImpl) applyGasCeil(ctx context.Context, transactor *bind.TransactOpts, chainID int, useDynamic bool) (err error) {
	_, span := t.metrics.Tracer().Start(ctx, "submitter.applyGasCeil")

	maxPrice := t.config.GetMaxGasPrice(chainID)

	defer func() {
		span.SetAttributes(attribute.String("max_price", bigPtrToString(maxPrice)))
		metrics.EndSpanWithErr(span, err)
	}()

	if useDynamic {
		if transactor.GasFeeCap.Cmp(maxPrice) > 0 {
			return fmt.Errorf("gas fee cap %s exceeds max price %s", transactor.GasFeeCap, maxPrice)
		}
		if transactor.GasTipCap.Cmp(transactor.GasFeeCap) > 0 {
			transactor.GasTipCap = core.CopyBigInt(transactor.GasFeeCap)
			span.AddEvent("tip cap exceeds fee cap; setting tip cap to fee cap")
		}
	} else {
		if transactor.GasPrice.Cmp(maxPrice) > 0 {
			return fmt.Errorf("gas price %s exceeds max price %s", transactor.GasPrice, maxPrice)
		}
	}
	return nil
}

func maxOfBig(a, b *big.Int) *big.Int {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Cmp(b) > 0 {
		return a
	}
	return b
}

// getGasBlock gets the gas block for the given chain.
func (t *txSubmitterImpl) getGasBlock(ctx context.Context, chainClient client.EVM, chainID int) (gasBlock *types.Header, err error) {
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

	// if we can't get the current gas block, attempt to load it from the cache
	if err != nil {
		var ok bool
		gasBlock, ok = t.lastGasBlockCache.Load(chainID)
		if ok {
			span.AddEvent("could not get gas block; using cached value", trace.WithAttributes(
				attribute.String("error", err.Error()),
				attribute.String("blockNumber", bigPtrToString(gasBlock.Number)),
			))
		} else {
			return nil, fmt.Errorf("could not get gas block: %w", err)
		}
	}

	// cache the latest gas block
	t.lastGasBlockCache.Store(chainID, gasBlock)

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

	// since we checked for dynamic gas estimate above, we can fetch the gas estimate here
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

	return gasEstimate, nil
}

var _ TransactionSubmitter = &txSubmitterImpl{}
