package db

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/scribe/types"
)

// EventDB stores events.
type EventDB interface {
	// StoreLog stores a log
	StoreLog(ctx context.Context, log ethTypes.Log, chainID uint32) error
	// RetrieveLogByTxHash retrieves a log by tx hash
	RetrieveLogByTxHash(ctx context.Context, txHash common.Hash) (log types.Log, err error)
	// StoreReceipt stores a receipt
	StoreReceipt(ctx context.Context, receipt ethTypes.Receipt) error
	// RetrieveReceiptByTxHash retrieves a receipt by tx hash
	RetrieveReceiptByTxHash(ctx context.Context, txHash common.Hash) (receipt types.Receipt, err error)
}
