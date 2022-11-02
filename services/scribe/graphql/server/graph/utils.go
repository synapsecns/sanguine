package graph

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/model"
	scribetypes "github.com/synapsecns/sanguine/services/scribe/types"
)

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

func (r Resolver) ethTxsToModelTransactions(ethTxs []types.Transaction, chainID uint32) []*model.Transaction {
	modelTxs := make([]*model.Transaction, len(ethTxs))
	for i, ethTx := range ethTxs {
		modelTxs[i] = r.ethTxToModelTransaction(ethTx, chainID)
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

func (r Resolver) failedLogsToModelFailedLogs(failedLogs []*scribetypes.FailedLog) []*model.FailedLog {
	modelFailedLogs := make([]*model.FailedLog, len(failedLogs))
	for i, failedLog := range failedLogs {
		modelFailedLogs[i] = r.failedLogToModelFailedLog(failedLog)
	}
	return modelFailedLogs
}

func (r Resolver) failedLogToModelFailedLog(failedLog *scribetypes.FailedLog) *model.FailedLog {
	return &model.FailedLog{
		ChainID:         int(failedLog.ChainID),
		ContractAddress: failedLog.ContractAddress,
		TxHash:          failedLog.TxHash,
		BlockIndex:      int(failedLog.BlockIndex),
		BlockNumber:     int(failedLog.BlockNumber),
		FailedAttempts:  int(failedLog.FailedAttempts),
	}
}
