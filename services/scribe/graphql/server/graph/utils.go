package graph

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/model"
)

var logger = log.Logger("scribe-graph")

func (r Resolver) receiptsToModelReceipts(receipts []types.Receipt, chainID uint32) []*model.Receipt {
	modelReceipts := make([]*model.Receipt, len(receipts))

	for i, receipt := range receipts {
		modelReceipts[i] = r.receiptToModelReceipt(receipt, chainID)
	}

	return modelReceipts
}

func (r Resolver) receiptToModelReceipt(receipt types.Receipt, chainID uint32) *model.Receipt {
	return &model.Receipt{
		ChainID:           int(chainID),
		Type:              int(receipt.Type),
		PostState:         string(receipt.PostState),
		Status:            int(receipt.Status),
		CumulativeGasUsed: int(receipt.CumulativeGasUsed),
		Bloom:             common.Bytes2Hex(receipt.Bloom.Bytes()),
		TxHash:            receipt.TxHash.String(),
		ContractAddress:   receipt.ContractAddress.String(),
		GasUsed:           int(receipt.GasUsed),
		BlockNumber:       int(receipt.BlockNumber.Int64()),
		TransactionIndex:  int(receipt.TransactionIndex),
	}
}

func (r Resolver) logsToModelLogs(logs []*types.Log, chainID uint32) []*model.Log {
	modelLogs := make([]*model.Log, len(logs))
	for i, log := range logs {
		modelLogs[i] = r.logToModelLog(log, chainID)
	}

	return modelLogs
}

func (r Resolver) logToModelLog(log *types.Log, chainID uint32) *model.Log {
	topicsList := make([]string, len(log.Topics))

	for i, topic := range log.Topics {
		topicsList[i] = topic.String()
	}

	return &model.Log{
		ContractAddress: log.Address.String(),
		ChainID:         int(chainID),
		Topics:          topicsList,
		Data:            common.Bytes2Hex(log.Data),
		BlockNumber:     int(log.BlockNumber),
		TxHash:          log.TxHash.String(),
		TxIndex:         int(log.TxIndex),
		BlockHash:       log.BlockHash.String(),
		Index:           int(log.Index),
		Removed:         log.Removed,
	}
}

func (r Resolver) ethTxsToModelTransactions(ctx context.Context, ethTxs []db.TxWithBlockNumber, chainID uint32) []*model.Transaction {
	modelTxs := make([]*model.Transaction, len(ethTxs))

	for i := range ethTxs {
		ethTx := ethTxs[i]
		modelTxs[i] = r.ethTxToModelTransaction(ethTx.Tx, chainID)

		// Return empty sender if that this operation errors (will only occur in tests or invalid txs).
		msgFrom, _ := ethTx.Tx.AsMessage(types.LatestSignerForChainID(ethTx.Tx.ChainId()), big.NewInt(1))
		modelTxs[i].Sender = msgFrom.From().String()

		timestamp, err := r.DB.RetrieveBlockTime(ctx, chainID, ethTx.BlockNumber)
		if err != nil || timestamp == 0 {
			newBlockTime, err := r.getBlockTime(ctx, chainID, ethTx.BlockNumber)
			if err != nil {
				continue
			}

			timestamp = *newBlockTime
			err = r.DB.StoreBlockTime(ctx, chainID, ethTx.BlockNumber, *newBlockTime)
			if err != nil {
				continue
			}
		}

		modelTxs[i].Timestamp = int(timestamp)
	}

	return modelTxs
}

func (r Resolver) ethTxToModelTransaction(ethTx types.Transaction, chainID uint32) *model.Transaction {
	protected := ethTx.Protected()
	return &model.Transaction{
		ChainID:   int(chainID),
		TxHash:    ethTx.Hash().String(),
		Protected: protected,
		Type:      int(ethTx.Type()),
		Data:      common.Bytes2Hex(ethTx.Data()),
		Gas:       int(ethTx.Gas()),
		GasPrice:  int(ethTx.GasPrice().Uint64()),
		GasTipCap: ethTx.GasFeeCap().String(),
		GasFeeCap: ethTx.GasTipCap().String(),
		Value:     ethTx.Value().String(),
		Nonce:     int(ethTx.Nonce()),
		To:        ethTx.To().String(),
	}
}

// getBlockTime retrieves a singular blocktime.
//
//nolint:gocognit,cyclop
func (r Resolver) getBlockTime(ctx context.Context, chainID uint32, blockNumber uint64) (*uint64, error) {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    30 * time.Millisecond,
		Max:    5 * time.Second,
	}

	timeout := time.Duration(0)
	var backendClient backend.ScribeBackend
	backendClient, err := backend.DialBackend(ctx, fmt.Sprintf("%s/%d", r.OmniRPCURL, chainID), r.Metrics)
	if err != nil {
		return nil, fmt.Errorf("could not create backend client: %w", err)
	}

	for {
		select {
		case <-ctx.Done():

			return nil, fmt.Errorf("context canceled: %w", ctx.Err())
		case <-time.After(timeout):
			block, err := backendClient.HeaderByNumber(ctx, big.NewInt(int64(blockNumber)))

			if err != nil {
				timeout = b.Duration()
				fmt.Println("TESTING--", fmt.Sprintf("%s/%d", r.OmniRPCURL, chainID), err)

				continue
			}
			blockTime := block.Time
			return &blockTime, nil
		}
	}
}
