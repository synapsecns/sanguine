package db

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

// TxQueueDB contains an interface for storing transactions currently being processed.
//
//go:generate go run github.com/vektra/mockery/v2 --name TxQueueDB --output ./mocks --case=underscore
type TxQueueDB interface {
	// StoreRawTx stores a raw transaction
	StoreRawTx(ctx context.Context, tx *ethTypes.Transaction, chainID *big.Int, from common.Address) error
	// StoreProcessedTx stores a tx that has already been processed
	StoreProcessedTx(ctx context.Context, tx *ethTypes.Transaction) error
	// GetNonceForChainID gets the latest nonce for a sender
	GetNonceForChainID(ctx context.Context, fromAddress common.Address, chainID *big.Int) (nonce uint64, err error)
}
