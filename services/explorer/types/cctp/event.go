package cctp

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EventLog is the interface for all bridge events.
//
//nolint:interfacebloat
type EventLog interface {
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetBlockNumber returns the block number of the log.
	GetBlockNumber() uint64
	// GetTxHash returns the transaction hash of the log.
	GetTxHash() common.Hash

	GetSender() *big.Int
	GetDestinationChainID()

	//	uint256 chainId,
	//address indexed sender,
	//uint64 nonce,
	//address token,
	//uint256 amount,
	//uint32 requestVersion,
	//bytes formattedRequest,
	//bytes32 requestID
}
