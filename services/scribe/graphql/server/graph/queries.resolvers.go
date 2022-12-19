package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/model"
	resolvers "github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/resolver"
)

// Logs is the resolver for the logs field.
func (r *queryResolver) Logs(ctx context.Context, contractAddress *string, chainID int, blockNumber *int, txHash *string, txIndex *int, blockHash *string, index *int, confirmed *bool, page int) ([]*model.Log, error) {
	logsFilter := db.BuildLogFilter(contractAddress, blockNumber, txHash, txIndex, blockHash, index, confirmed)
	logsFilter.ChainID = uint32(chainID)
	logs, err := r.DB.RetrieveLogsWithFilter(ctx, logsFilter, page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving logs: %w", err)
	}

	return r.logsToModelLogs(logs, logsFilter.ChainID), nil
}

// LogsRange is the resolver for the logsRange field.
func (r *queryResolver) LogsRange(ctx context.Context, contractAddress *string, chainID int, blockNumber *int, txHash *string, txIndex *int, blockHash *string, index *int, confirmed *bool, startBlock int, endBlock int, page int) ([]*model.Log, error) {
	logsFilter := db.BuildLogFilter(contractAddress, blockNumber, txHash, txIndex, blockHash, index, confirmed)
	logsFilter.ChainID = uint32(chainID)
	logs, err := r.DB.RetrieveLogsInRange(ctx, logsFilter, uint64(startBlock), uint64(endBlock), page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving logs: %w", err)
	}

	return r.logsToModelLogs(logs, logsFilter.ChainID), nil
}

// Receipts is the resolver for the receipts field.
func (r *queryResolver) Receipts(ctx context.Context, chainID int, txHash *string, contractAddress *string, blockHash *string, blockNumber *int, txIndex *int, confirmed *bool, page int) ([]*model.Receipt, error) {
	receiptsFilter := db.BuildReceiptFilter(txHash, contractAddress, blockHash, blockNumber, txIndex, confirmed)
	receiptsFilter.ChainID = uint32(chainID)
	receipts, err := r.DB.RetrieveReceiptsWithFilter(ctx, receiptsFilter, page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipts: %w", err)
	}

	return r.receiptsToModelReceipts(receipts, receiptsFilter.ChainID), nil
}

// ReceiptsRange is the resolver for the receiptsRange field.
func (r *queryResolver) ReceiptsRange(ctx context.Context, chainID int, txHash *string, contractAddress *string, blockHash *string, blockNumber *int, txIndex *int, confirmed *bool, startBlock int, endBlock int, page int) ([]*model.Receipt, error) {
	receiptsFilter := db.BuildReceiptFilter(txHash, contractAddress, blockHash, blockNumber, txIndex, confirmed)
	receiptsFilter.ChainID = uint32(chainID)
	receipts, err := r.DB.RetrieveReceiptsInRange(ctx, receiptsFilter, uint64(startBlock), uint64(endBlock), page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipts: %w", err)
	}

	return r.receiptsToModelReceipts(receipts, receiptsFilter.ChainID), nil
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context, txHash *string, chainID int, blockNumber *int, blockHash *string, confirmed *bool, page int) ([]*model.Transaction, error) {
	transactionsFilter := db.BuildEthTxFilter(txHash, blockNumber, blockHash, confirmed)
	transactionsFilter.ChainID = uint32(chainID)
	transactions, err := r.DB.RetrieveEthTxsWithFilter(ctx, transactionsFilter, page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving transactions: %w", err)
	}

	return r.ethTxsToModelTransactions(ctx, transactions, transactionsFilter.ChainID), nil
}

// TransactionsRange is the resolver for the transactionsRange field.
func (r *queryResolver) TransactionsRange(ctx context.Context, txHash *string, chainID int, blockNumber *int, blockHash *string, confirmed *bool, startBlock int, endBlock int, page int) ([]*model.Transaction, error) {
	transactionsFilter := db.BuildEthTxFilter(txHash, blockNumber, blockHash, confirmed)
	transactionsFilter.ChainID = uint32(chainID)
	transactions, err := r.DB.RetrieveEthTxsInRange(ctx, transactionsFilter, uint64(startBlock), uint64(endBlock), page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving transactions: %w", err)
	}

	return r.ethTxsToModelTransactions(ctx, transactions, transactionsFilter.ChainID), nil
}

// BlockTime is the resolver for the blockTime field.
func (r *queryResolver) BlockTime(ctx context.Context, chainID int, blockNumber int) (*int, error) {
	blockTime, err := r.DB.RetrieveBlockTime(ctx, uint32(chainID), uint64(blockNumber))
	if err != nil {
		blockTimeRaw, err := r.getBlockTime(ctx, uint32(chainID), uint64(blockNumber))
		if err != nil {
			return nil, fmt.Errorf("error retrieving block time: %w", err)
		}
		blockTime = *blockTimeRaw

		go func() {
			// we create a new context here to allow async storage and allow for quick returns
			// TODO: this is a hotfix and should be undone once reindexed
			storeCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			storeErr := r.DB.StoreBlockTime(storeCtx, uint32(chainID), uint64(blockNumber), blockTime)

			if storeErr != nil {
				logger.Error(storeErr)
			}
		}()
	}
	blockTimeInt := int(blockTime)

	return &blockTimeInt, nil
}

// LastStoredBlockNumber is the resolver for the lastStoredBlockNumber field.
func (r *queryResolver) LastStoredBlockNumber(ctx context.Context, chainID int) (*int, error) {
	blockNumber, err := r.DB.RetrieveLastBlockStored(ctx, uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving last block: %w", err)
	}

	blockNumberInt := int(blockNumber)

	return &blockNumberInt, nil
}

// FirstStoredBlockNumber is the resolver for the firstStoredBlockNumber field.
func (r *queryResolver) FirstStoredBlockNumber(ctx context.Context, chainID int) (*int, error) {
	blockNumber, err := r.DB.RetrieveFirstBlockStored(ctx, uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving first block: %w", err)
	}

	blockNumberInt := int(blockNumber)

	return &blockNumberInt, nil
}

// LastConfirmedBlockNumber is the resolver for the lastConfirmedBlockNumber field.
func (r *queryResolver) LastConfirmedBlockNumber(ctx context.Context, chainID int) (*int, error) {
	blockNumber, err := r.DB.RetrieveLastConfirmedBlock(ctx, uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving first block: %w", err)
	}

	blockNumberInt := int(blockNumber)

	return &blockNumberInt, nil
}

// TxSender is the resolver for the txSender field.
func (r *queryResolver) TxSender(ctx context.Context, txHash string, chainID int) (*string, error) {
	filter := db.EthTxFilter{
		TxHash:  txHash,
		ChainID: uint32(chainID),
	}

	ethTx, err := r.DB.RetrieveEthTxsWithFilter(ctx, filter, 1)
	if err != nil || len(ethTx) == 0 {
		return nil, fmt.Errorf("error retrieving transaction: %w", err)
	}

	msgFrom, err := ethTx[0].Tx.AsMessage(types.LatestSignerForChainID(ethTx[0].Tx.ChainId()), big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("error retrieving ethtx: %w", err)
	}

	sender := msgFrom.From().String()

	return &sender, nil
}

// LastIndexed is the resolver for the lastIndexed field.
func (r *queryResolver) LastIndexed(ctx context.Context, contractAddress string, chainID int) (*int, error) {
	blockNumber, err := r.DB.RetrieveLastIndexed(ctx, common.HexToAddress(contractAddress), uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving contract last block: %w", err)
	}

	blockNumberInt := int(blockNumber)

	return &blockNumberInt, nil
}

// LogCount is the resolver for the logCount field.
func (r *queryResolver) LogCount(ctx context.Context, contractAddress string, chainID int) (*int, error) {
	logCount, err := r.DB.RetrieveLogCountForContract(ctx, common.HexToAddress(contractAddress), uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving log count: %w", err)
	}

	logCountInt := int(logCount)

	return &logCountInt, nil
}

// ReceiptCount is the resolver for the receiptCount field.
func (r *queryResolver) ReceiptCount(ctx context.Context, chainID int) (*int, error) {
	receiptCount, err := r.DB.RetrieveReceiptCountForChain(ctx, uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipt count: %w", err)
	}

	logCountInt := int(receiptCount)

	return &logCountInt, nil
}

// BlockTimeCount is the resolver for the blockTimeCount field.
func (r *queryResolver) BlockTimeCount(ctx context.Context, chainID int) (*int, error) {
	blockTimesCount, err := r.DB.RetrieveBlockTimesCountForChain(ctx, uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving contract last block: %w", err)
	}

	blockTimesCountInt := int(blockTimesCount)

	return &blockTimesCountInt, nil
}

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
