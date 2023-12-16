package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"math/big"
)

// EVM is the set of functions that the scribe needs from a client.
//
//go:generate go run github.com/vektra/mockery/v2 --name=EVM --output=mocks --case=underscore
type EVM interface {
	// ContractBackend defines the methods needed to work with contracts on a read-write basis.
	// this is used for deploying an interacting with contracts
	bind.ContractBackend
	// ChainReader ethereum.ChainReader for getting transactions
	ethereum.ChainReader
	// TransactionReader is used for reading txes by hash
	ethereum.TransactionReader
	// ChainStateReader gets the chain state reader
	ethereum.ChainStateReader
	// PendingStateReader handles pending state calls
	ethereum.PendingStateReader
	// ChainSyncReader tracks state head
	ethereum.ChainSyncReader
	// PendingContractCaller tracks pending contract calls
	ethereum.PendingContractCaller
	// FeeHistory gets the fee history for a given block
	FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (*ethereum.FeeHistory, error)
	// NetworkID returns the network ID (also known as the chain ID) for this chain.
	NetworkID(ctx context.Context) (*big.Int, error)
	// ChainID gets the chain id from the rpc server
	ChainID(ctx context.Context) (*big.Int, error)
	// CallContext is used for manual overrides
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	// BatchCallContext is used for manual overrides
	BatchCallContext(ctx context.Context, b []rpc.BatchElem) error
	// BlockNumber gets the latest block number
	BlockNumber(ctx context.Context) (uint64, error)
	// BatchWithContext batches multiple w3type calls
	BatchWithContext(ctx context.Context, calls ...w3types.Caller) error
}

type clientImpl struct {
	tracing           metrics.Handler
	captureClient     *captureClient
	rpcClient         *rpc.Client
	endpoint          string
	captureRequestRes bool
	// TODO: consider using sync.Pool for capture clients to improve performance
}

// DialBackend returns a scribe backend.
func DialBackend(ctx context.Context, url string, handler metrics.Handler, opts ...Options) (_ EVM, err error) {
	client := &clientImpl{
		endpoint: url,
		tracing:  handler,
	}

	for _, opt := range opts {
		opt(client)
	}

	client.captureClient, err = newCaptureClient(ctx, url, handler, client.captureRequestRes)
	if err != nil {
		return nil, fmt.Errorf("could not create capture client: %w", err)
	}

	return client, nil
}

const (
	batchAttribute    = "batch"
	methodsAttribute  = "methods"
	endpointAttribute = "endpoint"
)

func (c *clientImpl) getEthClient() *ethclient.Client {
	return c.captureClient.ethClient
}

func (c *clientImpl) getW3Client() *w3.Client {
	return c.captureClient.w3Client
}

// BatchWithContext batches multiple w3 calls.
func (c *clientImpl) BatchWithContext(ctx context.Context, calls ...w3types.Caller) (err error) {
	// Do not create an even if there are no calls
	if len(calls) == 0 {
		return nil
	}

	ctx, span := c.tracing.Tracer().Start(ctx, batchAttribute)
	span.SetAttributes(parseCalls(calls))
	span.SetAttributes(attribute.String(endpointAttribute, c.endpoint))

	defer func() {
		if errors.Is(err, w3.CallErrors{}) {
			var batchErr w3.CallErrors
			_ = errors.As(err, &batchErr)
			for i, callErr := range batchErr {
				rawReq, err := calls[i].CreateRequest()
				// this already happened, so it can't be failing now.
				// just error.
				if err != nil {
					fmt.Println("could not create request: this should never happen", err)
					continue
				}

				params, err := json.Marshal(rawReq.Args)
				if err != nil {
					fmt.Println("could not marshal params: this should never happen", err)
					continue
				}
				span.RecordError(callErr, trace.WithAttributes(attribute.String("method", rawReq.Method), attribute.String("params", string(params))))
			}
			metrics.EndSpan(span)
			return
		}
		metrics.EndSpanWithErr(span, err)
	}()

	//nolint: wrapcheck
	return c.getW3Client().CallCtx(ctx, calls...)
}

// BatchCallContext calls BatchCallContext on the underlying client. Note: this will bypass the rate-limiter.
//
//nolint:wrapcheck
func (c *clientImpl) BatchCallContext(ctx context.Context, b []rpc.BatchElem) (err error) {
	requestCtx, span := c.startSpan(ctx, NetVersionMethod)
	span.SetAttributes(parseBatch(b))
	span.SetAttributes(attribute.String(endpointAttribute, c.endpoint))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.rpcClient.BatchCallContext(requestCtx, b)
}

func (c *clientImpl) startSpan(parentCtx context.Context, method RPCMethod) (context.Context, trace.Span) {
	ctx, span := c.tracing.Tracer().Start(parentCtx, method.String())
	span.SetAttributes(attribute.String("endpoint", c.endpoint))

	return ctx, span
}

// CallContract calls contract on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) (contractResponse []byte, err error) {
	requestCtx, span := c.startSpan(ctx, CallMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().CallContract(requestCtx, call, blockNumber)
}

// PendingCallContract calls contract on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) PendingCallContract(ctx context.Context, call ethereum.CallMsg) (contractResponse []byte, err error) {
	requestCtx, span := c.startSpan(ctx, CallMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().PendingCallContract(requestCtx, call)
}

// PendingCodeAt calls PendingCodeAt on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) PendingCodeAt(ctx context.Context, account common.Address) (codeResponse []byte, err error) {
	requestCtx, span := c.startSpan(ctx, GetCodeMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().PendingCodeAt(requestCtx, account)
}

// PendingBalanceAt calls PendingBalanceAt on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) PendingBalanceAt(ctx context.Context, account common.Address) (pendingBalance *big.Int, err error) {
	requestCtx, span := c.startSpan(ctx, GetBalanceMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().PendingBalanceAt(requestCtx, account)
}

// PendingStorageAt calls PendingStorageAt on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) (pendingStorage []byte, err error) {
	requestCtx, span := c.startSpan(ctx, StorageAtMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().PendingStorageAt(requestCtx, account, key)
}

// PendingNonceAt calls PendingNonceAt on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) PendingNonceAt(ctx context.Context, account common.Address) (pendingNonce uint64, err error) {
	requestCtx, span := c.startSpan(ctx, TransactionCountMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().PendingNonceAt(requestCtx, account)
}

// PendingTransactionCount calls PendingTransactionCount on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) PendingTransactionCount(ctx context.Context) (count uint, err error) {
	requestCtx, span := c.startSpan(ctx, PendingTransactionCountMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().PendingTransactionCount(requestCtx)
}

// NetworkID calls NetworkID on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) NetworkID(ctx context.Context) (id *big.Int, err error) {
	requestCtx, span := c.startSpan(ctx, NetVersionMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().NetworkID(requestCtx)
}

// SyncProgress calls SyncProgress on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) SyncProgress(ctx context.Context) (syncProgress *ethereum.SyncProgress, err error) {
	requestCtx, span := c.startSpan(ctx, SyncProgressMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()
	return c.getEthClient().SyncProgress(requestCtx)
}

// SuggestGasPrice calls SuggestGasPrice on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) SuggestGasPrice(ctx context.Context) (gasPrice *big.Int, err error) {
	requestCtx, span := c.startSpan(ctx, GasPriceMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().SuggestGasPrice(requestCtx)
}

// EstimateGas calls EstimateGas on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	requestCtx, span := c.startSpan(ctx, EstimateGasMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().EstimateGas(requestCtx, call)
}

// SendTransaction calls SendTransaction on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) SendTransaction(ctx context.Context, tx *types.Transaction) (err error) {
	requestCtx, span := c.startSpan(ctx, SendRawTransactionMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().SendTransaction(requestCtx, tx)
}

// FilterLogs calls FilterLogs on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) FilterLogs(ctx context.Context, query ethereum.FilterQuery) (logs []types.Log, err error) {
	requestCtx, span := c.startSpan(ctx, GetLogsMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().FilterLogs(requestCtx, query)
}

// SubscribeFilterLogs calls SubscribeFilterLogs on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (sub ethereum.Subscription, err error) {
	requestCtx, span := c.startSpan(ctx, SubscribeMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()
	return c.getEthClient().SubscribeFilterLogs(requestCtx, query, ch)
}

// BlockByHash calls BlockByHash on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) BlockByHash(ctx context.Context, hash common.Hash) (block *types.Block, err error) {
	requestCtx, span := c.startSpan(ctx, BlockByHashMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().BlockByHash(requestCtx, hash)
}

// BlockByNumber calls BlockByNumber on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) BlockByNumber(ctx context.Context, number *big.Int) (block *types.Block, err error) {
	requestCtx, span := c.startSpan(ctx, BlockByNumberMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().BlockByNumber(requestCtx, number)
}

// HeaderByHash calls HeaderByHash on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) HeaderByHash(ctx context.Context, hash common.Hash) (header *types.Header, err error) {
	requestCtx, span := c.startSpan(ctx, BlockByHashMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().HeaderByHash(requestCtx, hash)
}

// HeaderByNumber calls HeaderByNumber on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) HeaderByNumber(ctx context.Context, number *big.Int) (header *types.Header, err error) {
	requestCtx, span := c.startSpan(ctx, BlockByNumberMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().HeaderByNumber(requestCtx, number)
}

// TransactionCount calls TransactionCount on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) TransactionCount(ctx context.Context, blockHash common.Hash) (txCount uint, err error) {
	requestCtx, span := c.startSpan(ctx, TransactionCountByHashMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().TransactionCount(requestCtx, blockHash)
}

// TransactionInBlock calls TransactionInBlock on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (tx *types.Transaction, err error) {
	requestCtx, span := c.startSpan(ctx, TransactionByBlockHashAndIndexMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().TransactionInBlock(requestCtx, blockHash, index)
}

// SubscribeNewHead calls SubscribeNewHead on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (sub ethereum.Subscription, err error) {
	requestCtx, span := c.startSpan(ctx, SubscribeMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().SubscribeNewHead(requestCtx, ch)
}

// TransactionByHash calls TransactionByHash on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	requestCtx, span := c.startSpan(ctx, TransactionByHashMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().TransactionByHash(requestCtx, txHash)
}

// TransactionReceipt calls TransactionReceipt on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) TransactionReceipt(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error) {
	requestCtx, span := c.startSpan(ctx, TransactionReceiptByHashMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().TransactionReceipt(requestCtx, txHash)
}

// BalanceAt calls BalanceAt on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (balance *big.Int, err error) {
	requestCtx, span := c.startSpan(ctx, GetBalanceMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().BalanceAt(requestCtx, account, blockNumber)
}

// StorageAt calls StorageAt on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) (storage []byte, err error) {
	requestCtx, span := c.startSpan(ctx, StorageAtMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().StorageAt(requestCtx, account, key, blockNumber)
}

// BlockNumber gets the latest block number
//
//nolint:wrapcheck
func (c *clientImpl) BlockNumber(ctx context.Context) (_ uint64, err error) {
	requestCtx, span := c.startSpan(ctx, BlockNumberMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().BlockNumber(requestCtx)
}

// CodeAt calls CodeAt on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) (codeAt []byte, err error) {
	requestCtx, span := c.startSpan(ctx, GetCodeMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().CodeAt(requestCtx, account, blockNumber)
}

// SuggestGasTipCap gets the suggested gas tip for a chain.
//
//nolint:wrapcheck
func (c *clientImpl) SuggestGasTipCap(ctx context.Context) (tip *big.Int, err error) {
	requestCtx, span := c.startSpan(ctx, MaxPriorityMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().SuggestGasTipCap(requestCtx)
}

// CallContext calls CallContext on the underlying client. Note: this will bypass the rate-limiter.
//
//nolint:wrapcheck
func (c *clientImpl) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) (err error) {
	requestCtx, span := c.startSpan(ctx, RPCMethod(method))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.rpcClient.CallContext(requestCtx, result, method, args...)
}

// NonceAt calls NonceAt on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (nonce uint64, err error) {
	requestCtx, span := c.startSpan(ctx, TransactionCountMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()
	return c.getEthClient().NonceAt(requestCtx, account, blockNumber)
}

// ChainID calls ChainID on the underlying client.
//
//nolint:wrapcheck
func (c *clientImpl) ChainID(ctx context.Context) (chainID *big.Int, err error) {
	requestCtx, span := c.startSpan(ctx, ChainIDMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().ChainID(requestCtx)
}

// FeeHistory calls FeeHistory on the underlying client
//
//nolint:wrapcheck
func (c *clientImpl) FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (_ *ethereum.FeeHistory, err error) {
	requestCtx, span := c.startSpan(ctx, FeeHistoryMethod)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	return c.getEthClient().FeeHistory(requestCtx, blockCount, lastBlock, rewardPercentiles)
}

// parseCalls parses out calls from w3types.Caller.
func parseCalls(calls []w3types.Caller) attribute.KeyValue {
	res := make([]string, len(calls))

	for i, call := range calls {
		req, err := call.CreateRequest()
		if err != nil {
			res[i] = fmt.Sprintf("unknown: %v", err)
			continue
		}
		res[i] = req.Method
	}

	return attribute.StringSlice(methodsAttribute, res)
}

func parseBatch(batchElem []rpc.BatchElem) attribute.KeyValue {
	res := make([]string, len(batchElem))

	for i, elem := range batchElem {
		res[i] = elem.Method
	}

	return attribute.StringSlice(methodsAttribute, res)
}
