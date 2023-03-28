package client

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/w3types"
	"math/big"
)

// ETHClient is the set of functions that the scribe needs from a client.
//
//go:generate go run github.com/vektra/mockery/v2 --name EVMClient --output ./mocks --case=underscore
type ETHClient interface {
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
