package client

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/ethergo/chain/client/near"
)

// Permitter handles permit acquires/releases for a lifecycle client.
// this is useful for implementing generic pre/post request logic on
// different kinds of clients that conform to the evm client.
//
//go:generate go run github.com/vektra/mockery/v2 --name Permitter --output ./mocks --case=underscore
type Permitter interface {
	AcquirePermit(ctx context.Context) (err error)
	// ReleasePermit releases a permit
	ReleasePermit()
}

// LifecycleClient is an evm client that acquires a permit upon request start and releases after the request.
// this logic can be implemented however the caller needs.
type LifecycleClient struct {
	underlyingClient EVMClient
	Permitter
	// chainID is the chain id used for the client
	chainID *big.Int
	// requestTimeout gets the request timeout.
	requestTimeout time.Duration
}

// NewLifecycleClient creates a new lifecyle client from an underlying client and a permitter.
// requestTimeout is the maximum amount of time to wait for any request. This does not include
// the acquirePermit/releasePermit time.
func NewLifecycleClient(client EVMClient, chainID *big.Int, permitter Permitter, requestTimeout time.Duration) LifecycleClient {
	return LifecycleClient{
		underlyingClient: client,
		Permitter:        permitter,
		chainID:          chainID,
		requestTimeout:   requestTimeout,
	}
}

// CallContract calls contract on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) (contractResponse []byte, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return contractResponse, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.CallContract(requestCtx, call, blockNumber)
}

// PendingCallContract calls contract on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) PendingCallContract(ctx context.Context, call ethereum.CallMsg) (contractResponse []byte, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return contractResponse, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.PendingCallContract(requestCtx, call)
}

// PendingCodeAt calls PendingCodeAt on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) PendingCodeAt(ctx context.Context, account common.Address) (codeResponse []byte, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return codeResponse, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.PendingCodeAt(requestCtx, account)
}

// PendingBalanceAt calls PendingBalanceAt on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) PendingBalanceAt(ctx context.Context, account common.Address) (pendingBalance *big.Int, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return pendingBalance, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.PendingBalanceAt(requestCtx, account)
}

// PendingStorageAt calls PendingStorageAt on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) (pendingStorage []byte, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return pendingStorage, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.PendingStorageAt(requestCtx, account, key)
}

// PendingNonceAt calls PendingNonceAt on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) PendingNonceAt(ctx context.Context, account common.Address) (pendingNonce uint64, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return pendingNonce, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.PendingNonceAt(requestCtx, account)
}

// PendingTransactionCount calls PendingTransactionCount on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) PendingTransactionCount(ctx context.Context) (count uint, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return count, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.PendingTransactionCount(requestCtx)
}

// NetworkID calls NetworkID on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) NetworkID(ctx context.Context) (id *big.Int, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return id, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.NetworkID(requestCtx)
}

// SyncProgress calls SyncProgress on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) SyncProgress(ctx context.Context) (syncProgress *ethereum.SyncProgress, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return nil, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.SyncProgress(requestCtx)
}

// SuggestGasPrice calls SuggestGasPrice on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) SuggestGasPrice(ctx context.Context) (gasPrice *big.Int, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return gasPrice, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.SuggestGasPrice(requestCtx)
}

// EstimateGas calls EstimateGas on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return gas, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.EstimateGas(requestCtx, call)
}

// SendTransaction calls SendTransaction on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) SendTransaction(ctx context.Context, tx *types.Transaction) (err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.SendTransaction(requestCtx, tx)
}

// FilterLogs calls FilterLogs on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) (logs []types.Log, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return logs, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.FilterLogs(requestCtx, query)
}

// SubscribeFilterLogs calls SubscribeFilterLogs on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (sub ethereum.Subscription, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return sub, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.SubscribeFilterLogs(requestCtx, query, ch)
}

// BlockByHash calls BlockByHash on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) BlockByHash(ctx context.Context, hash common.Hash) (block *types.Block, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return block, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	if m.chainID.Cmp(AuroraMainnet.ChainID) == 0 {
		return near.BlockByHash(ctx, m, hash)
	}

	return m.underlyingClient.BlockByHash(requestCtx, hash)
}

// BlockByNumber calls BlockByNumber on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) BlockByNumber(ctx context.Context, number *big.Int) (block *types.Block, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return block, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	if m.chainID.Cmp(AuroraMainnet.ChainID) == 0 {
		return near.BlockByNumber(ctx, m, number)
	}

	return m.underlyingClient.BlockByNumber(requestCtx, number)
}

// HeaderByHash calls HeaderByHash on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) HeaderByHash(ctx context.Context, hash common.Hash) (header *types.Header, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return header, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.HeaderByHash(requestCtx, hash)
}

// HeaderByNumber calls HeaderByNumber on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) HeaderByNumber(ctx context.Context, number *big.Int) (header *types.Header, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return header, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.HeaderByNumber(requestCtx, number)
}

// TransactionCount calls TransactionCount on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) TransactionCount(ctx context.Context, blockHash common.Hash) (txCount uint, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return txCount, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.TransactionCount(requestCtx, blockHash)
}

// TransactionInBlock calls TransactionInBlock on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (tx *types.Transaction, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return tx, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.TransactionInBlock(requestCtx, blockHash, index)
}

// SubscribeNewHead calls SubscribeNewHead on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (sub ethereum.Subscription, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return sub, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.SubscribeNewHead(requestCtx, ch)
}

// TransactionByHash calls TransactionByHash on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return tx, isPending, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.TransactionByHash(requestCtx, txHash)
}

// TransactionReceipt calls TransactionReceipt on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return receipt, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.TransactionReceipt(requestCtx, txHash)
}

// BalanceAt calls BalanceAt on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (balance *big.Int, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return balance, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.BalanceAt(requestCtx, account, blockNumber)
}

// StorageAt calls StorageAt on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) (storage []byte, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return storage, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.StorageAt(requestCtx, account, key, blockNumber)
}

// BlockNumber gets the latest block number
//
//nolint:wrapcheck
func (m LifecycleClient) BlockNumber(ctx context.Context) (_ uint64, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return 0, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.BlockNumber(requestCtx)
}

// CodeAt calls CodeAt on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) (codeAt []byte, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return codeAt, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.CodeAt(requestCtx, account, blockNumber)
}

// SuggestGasTipCap gets the suggested gas tip for a chain.
//
//nolint:wrapcheck
func (m LifecycleClient) SuggestGasTipCap(ctx context.Context) (tip *big.Int, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return tip, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.SuggestGasTipCap(requestCtx)
}

// CallContext calls CallContext on the underlying client. Note: this will bypass the rate-limiter.
//
//nolint:wrapcheck
func (m LifecycleClient) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) (err error) {
	return m.underlyingClient.CallContext(ctx, result, method, args...)
}

// BatchCallContext calls BatchCallContext on the underlying client. Note: this will bypass the rate-limiter.
//
//nolint:wrapcheck
func (m LifecycleClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return m.underlyingClient.BatchCallContext(ctx, b)
}

// NonceAt calls NonceAt on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (nonce uint64, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return nonce, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.NonceAt(requestCtx, account, blockNumber)
}

// ChainConfig calls ChainConfig on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) ChainConfig() *params.ChainConfig {
	return ConfigFromID(m.chainID)
}

// ChainID calls ChainID on the underlying client.
//
//nolint:wrapcheck
func (m LifecycleClient) ChainID(ctx context.Context) (chainID *big.Int, err error) {
	err = m.AcquirePermit(ctx)
	if err != nil {
		return nil, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.ChainID(requestCtx)
}

// BatchContext calls BatchContext on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) BatchContext(ctx context.Context, calls ...w3types.Caller) error {
	err := m.AcquirePermit(ctx)
	if err != nil {
		return err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.BatchContext(requestCtx, calls...)
}

// FeeHistory calls FeeHistory on the underlying client
//
//nolint:wrapcheck
func (m LifecycleClient) FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error) {
	err := m.AcquirePermit(ctx)
	if err != nil {
		return nil, err
	}
	defer m.ReleasePermit()

	requestCtx, cancel := context.WithTimeout(ctx, m.requestTimeout)
	defer cancel()

	return m.underlyingClient.FeeHistory(requestCtx, blockCount, lastBlock, rewardPercentiles)
}
