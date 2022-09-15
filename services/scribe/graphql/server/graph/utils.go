package graph

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/model"
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

func (r Resolver) buildLogFilter(contractAddress *string, blockNumber *int, txHash *string, txIndex *int, blockHash *string, index *int, confirmed *bool) db.LogFilter {
	logFilter := db.LogFilter{}
	if contractAddress != nil {
		logFilter.ContractAddress = *contractAddress
	}
	if blockNumber != nil {
		logFilter.BlockNumber = uint64(*blockNumber)
	}
	if txHash != nil {
		logFilter.TxHash = *txHash
	}
	if txIndex != nil {
		logFilter.TxIndex = uint64(*txIndex)
	}
	if blockHash != nil {
		logFilter.BlockHash = *blockHash
	}
	if index != nil {
		logFilter.Index = uint64(*index)
	}
	if confirmed != nil {
		logFilter.Confirmed = *confirmed
	}
	return logFilter
}

func (r Resolver) buildReceiptFilter(txHash *string, contractAddress *string, blockHash *string, blockNumber *int, transactionIndex *int, confirmed *bool) db.ReceiptFilter {
	receiptFilter := db.ReceiptFilter{}
	if txHash != nil {
		receiptFilter.TxHash = *txHash
	}
	if contractAddress != nil {
		receiptFilter.ContractAddress = *contractAddress
	}
	if blockHash != nil {
		receiptFilter.BlockHash = *blockHash
	}
	if blockNumber != nil {
		receiptFilter.BlockNumber = uint64(*blockNumber)
	}
	if transactionIndex != nil {
		receiptFilter.TransactionIndex = uint64(*transactionIndex)
	}
	if confirmed != nil {
		receiptFilter.Confirmed = *confirmed
	}
	return receiptFilter
}

func (r Resolver) buildEthTxFilter(txHash *string, blockNumber *int, blockHash *string, confirmed *bool) db.EthTxFilter {
	ethTxFilter := db.EthTxFilter{}
	if txHash != nil {
		ethTxFilter.TxHash = *txHash
	}
	if blockNumber != nil {
		ethTxFilter.BlockNumber = uint64(*blockNumber)
	}
	if blockHash != nil {
		ethTxFilter.BlockHash = *blockHash
	}
	if confirmed != nil {
		ethTxFilter.Confirmed = *confirmed
	}
	return ethTxFilter
}
