package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// Receipt is the receipt of a transaction.
type Receipt interface {
	// Status gets the status of the transaction
	Status() uint64
	// CumulativeGasUSed gets the total amount of gas used when this transaction was executed in the block
	CumulativeGasUsed() uint64
	// TxHash gets the hash of the transaction
	TxHash() common.Hash
	// ContractAddress gets the address of the contract created by the transaction
	ContractAddress() common.Address
	// GasUsed gets the amount of gas used by this transaction alone
	GasUsed() uint64
	// BlockHash gets the hash of the block in which this transaction was included
	BlockHash() common.Hash
	// BlockNumber gets the number of the block in which this transaction was included
	BlockNumber() uint64
	// TransactionIndex gets the index of the transaction in the block
	TransactionIndex() uint64
}

type receipt struct {
	// status of the transaction
	status uint64
	// cumulative amount of gas used when this transaction was executed in the block
	cumulativeGasUsed uint64
	// hash of the transaction
	txHash common.Hash
	// address of the contract created by the transaction
	contractAddress common.Address
	// amount of gas used by this transaction alone
	gasUsed uint64
	// hash of the block in which this transaction was included
	blockHash common.Hash
	// number of the block in which this transaction was included
	blockNumber uint64
	// index of the transaction in the block
	transactionIndex uint64
}

// NewReceipt creates a new receipt.
func NewReceipt(status uint64, cumulativeGasUsed uint64, txHash common.Hash, contractAddress common.Address, gasUsed uint64, blockHash common.Hash, blockNumber uint64, transactionIndex uint64) Receipt {
	return receipt{
		status:            status,
		cumulativeGasUsed: cumulativeGasUsed,
		txHash:            txHash,
		contractAddress:   contractAddress,
		gasUsed:           gasUsed,
		blockHash:         blockHash,
		blockNumber:       blockNumber,
		transactionIndex:  transactionIndex,
	}
}

func (r receipt) Status() uint64 {
	return r.status
}

func (r receipt) CumulativeGasUsed() uint64 {
	return r.cumulativeGasUsed
}

func (r receipt) TxHash() common.Hash {
	return r.txHash
}

func (r receipt) ContractAddress() common.Address {
	return r.contractAddress
}

func (r receipt) GasUsed() uint64 {
	return r.gasUsed
}

func (r receipt) BlockHash() common.Hash {
	return r.blockHash
}

func (r receipt) BlockNumber() uint64 {
	return r.blockNumber
}

func (r receipt) TransactionIndex() uint64 {
	return r.transactionIndex
}
