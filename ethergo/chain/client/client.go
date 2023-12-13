package client

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/dwasse/w3"
	"github.com/dwasse/w3/w3types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

// EVMClient is the chain client. It defines everything necessary to create a Chain.
//
// Deprecated: use ethergo/client
//
//go:generate go run github.com/vektra/mockery/v2 --name EVMClient --output ./mocks --case=underscore
type EVMClient interface {
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
	// BatchContext uses w3 as a helper method for batch calls
	BatchContext(ctx context.Context, calls ...w3types.Caller) error
}

// clientImpl is a client implementation for an ethclient.
// it adds the chainconfig to the chain.
type clientImpl struct {
	*ethclient.Client
	// rpcClient is the underlying rpc client
	rpcClient *rpc.Client
	// w3Client is sued for batch calls
	w3Client *w3.Client
	// chainID contains the chain id
	chainID *big.Int
	// config is the chain config
	config *params.ChainConfig
	// wsURL is stored for reconnection attempts
	wsURL string
	// ctx stores the context of the original client
	//nolint: containedctx
	ctx context.Context
}

func (c *clientImpl) BatchContext(ctx context.Context, calls ...w3types.Caller) error {
	//nolint:wrapcheck
	return c.w3Client.CallCtx(ctx, calls...)
}

// connectionResetTimeout is how long the client should wait before rehupping.
var connectionResetTimeout = time.Minute * 5

const (
	// secureWebsocketScheme defines the scheme used to connect to a secure websocket.
	secureWebsocketScheme = "wss"
	// insecureWebsocketScheme defines the scheme used to connect to an insecure websocket.
	insecureWebsocketScheme = "ws"
)

// StartConnectionResetTicker should be fired off as a goroutine upon a baseChain's
// initialization. This is only used for websocket connections
// At every connectionResetTimeout (as of writing, every 5 minutes), the client will
// attempt to reconnect to its RPC client as a way of "rehup"ing the connection, which
// is a workaround for an issue documented in https://github.com/ethereum/go-ethereum/issues/22266
// with websocket connections. This can be removed when that issue is fixed
// If the attempted reconnect fails, it is tried again on the next tick.
func (c *clientImpl) StartConnectionResetTicker(ctx context.Context) {
	parsedURL, err := url.Parse(c.wsURL)
	if err != nil {
		// this should never happen because we make sure the url can be parsed when validating the config. Nonetheless, we want to avoid a panic here
		logger.Warn(err)
		return
	}

	// we don't need to rehup if this isn't a websocket
	if parsedURL.Scheme != secureWebsocketScheme && parsedURL.Scheme != insecureWebsocketScheme {
		return
	}

	select {
	case <-ctx.Done():
		return
	case <-time.After(connectionResetTimeout):
		if err := c.AttemptReconnect(); err != nil {
			logger.Warn(errors.Wrapf(err, "reconnect attempt for Chain client for Chain ID %s failed", c.config.ChainID))
		}
	}
}

// AttemptReconnect attempts to create a new client. Because this replaces
// a pointer, this should be subscription safe.
func (c *clientImpl) AttemptReconnect() error {
	tmpRPCClient, err := rpc.DialContext(c.ctx, c.wsURL)
	if err != nil {
		return fmt.Errorf("could not create new client, continuing to use original: %w", err)
	}

	c.rpcClient = tmpRPCClient
	c.Client = ethclient.NewClient(c.rpcClient)
	c.w3Client = w3.NewClient(c.rpcClient)

	return nil
}

// CallContext exposes the CallContext methods in the underlying ethereum rpc client.
//
//nolint:wrapcheck
func (c *clientImpl) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) (err error) {
	//nolint:asasalint
	return c.rpcClient.CallContext(ctx, result, method, args)
}

// BatchCallContext calls BatchCallContext on the underlying ethereum rpc client.
//
//nolint:wrapcheck
func (c *clientImpl) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return c.rpcClient.BatchCallContext(ctx, b)
}

// NewClient creates a client from a url.
func NewClient(ctx context.Context, url string) (EVMClient, error) {
	rpcClient, err := rpc.DialContext(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("could not connect to rpc server %s. Received error: %w", url, err)
	}

	ethClient := ethclient.NewClient(rpcClient)
	w3Client := w3.NewClient(rpcClient)

	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get chain id: %w", err)
	}

	client := &clientImpl{
		chainID:   chainID,
		rpcClient: rpcClient,
		Client:    ethClient,
		wsURL:     url,
		ctx:       ctx,
		w3Client:  w3Client,
	}

	go client.StartConnectionResetTicker(ctx)

	return client, nil
}

// NewClientFromChainID creates a new client from a chain id.
func NewClientFromChainID(ctx context.Context, url string, chainID *big.Int) (EVMClient, error) {
	rpcClient, err := rpc.DialContext(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("could not connect to rpc server %s. Received error: %w", url, err)
	}

	ethClient := ethclient.NewClient(rpcClient)
	w3Client := w3.NewClient(rpcClient)

	client := &clientImpl{
		chainID:   chainID,
		rpcClient: rpcClient,
		Client:    ethClient,
		wsURL:     url,
		ctx:       ctx,
		w3Client:  w3Client,
	}

	go client.StartConnectionResetTicker(ctx)

	return client, nil
}
