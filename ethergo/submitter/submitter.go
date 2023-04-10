package submitter

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/mapmutex"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
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

// TransactionSubmitter is the interface for submitting transactions to the chain.
type TransactionSubmitter interface {
	// Start starts the transaction submitter.
	Start(ctx context.Context) error
	// SubmitTransaction submits a transaction to the chain.
	// the transaction is not guaranteed to be executed immediately, only at some point in the future.
	// the nonce is returned, and can be used to track the status of the transaction.
	SubmitTransaction(parentCtx context.Context, chainID *big.Int, call ContractCallType) (nonce uint64, err error)
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
}

// ClientFetcher is the interface for fetching a chain client.
type ClientFetcher interface {
	GetClient(ctx context.Context, chainID *big.Int) (client.EVM, error)
}

func NewTransactionSubmitter(metrics metrics.Handler, signer signer.Signer, fetcher ClientFetcher) TransactionSubmitter {
	return &txSubmitterImpl{
		metrics: metrics,
		signer:  signer,
		fetcher: fetcher,
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
	// TODO implement me
	panic("implement me")
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

	err = t.db.PutTX(ctx, tx, status)
	if err != nil {
		return fmt.Errorf("could not put tx: %w", err)
	}

	return nil
}

// ContractCallType is a contract call that can be called safely.
type ContractCallType func(transactor *bind.TransactOpts) (tx *types.Transaction, err error)

func (t *txSubmitterImpl) SubmitTransaction(parentCtx context.Context, chainID *big.Int, call ContractCallType) (nonce uint64, err error) {
	ctx, span := t.metrics.Tracer().Start(parentCtx, "submitter.SubmitTransaction", trace.WithAttributes(
		attribute.Stringer("chainID", chainID),
		attribute.String("caller", runtime.FuncForPC(reflect.ValueOf(call).Pointer()).Name()),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// make sure we have a client for this chain.
	_, err = t.fetcher.GetClient(ctx, chainID)
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

		transaction, err = util.CopyTXWithNonce(transaction, newNonce)
		if err != nil {
			return nil, fmt.Errorf("could not copy tx: %w", err)
		}

		// TODO: gas pricing

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

	go func() {
		t.submitter.SubmitTransaction(ctx, tx)
	}()

	return tx.Nonce(), nil
}

var _ TransactionSubmitter = &txSubmitterImpl{}
